/*
 *  TVPN: A Peer-to-Peer VPN solution for traversing NAT firewalls
 *  Copyright (C) 2013  Joshua Chase <jcjoshuachase@gmail.com>
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along
 *  with this program; if not, write to the Free Software Foundation, Inc.,
 *  51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/

package dh

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha512"
	"math/big"
	"encoding/base64"
)

type Params struct {
	Priv []byte
	X    *big.Int
	Y    *big.Int
}

// Encodes the X parameter to a base64 string
func (p Params) XS() string {
	return base64.StdEncoding.EncodeToString(p.X.Bytes())
}

// Encodes the Y parameter to a base64 string
func (p Params) YS() string {
	return base64.StdEncoding.EncodeToString(p.Y.Bytes())
}

var curve elliptic.Curve

func init() {
	curve = elliptic.P521()
}

func GenParams() Params {
	priv, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	return Params{priv, x, y}
}

func GenMutSecret(local, remote Params) *big.Int {
	secret, _ := curve.ScalarMult(remote.X, remote.Y, local.Priv)
	return secret
}

func GenKey(local, remote Params) [64]byte {
	secret := GenMutSecret(local,remote)
	return sha512.Sum512(secret.Bytes())
}
