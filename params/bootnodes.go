// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the mian WhaleCoin network.
var MainnetBootnodes = []string{

  "enode://ff82427382f1c4b0c6b0113e742cb0b06264385ff2c074e65780a667325358fd327f9f7b3fd65e74e89a2c8081448533a80324d52da25c781a224d749d629441@34.235.197.115:30371",
  "enode://690db490f15a366b15126f67221bbcb0d4a6e6089e96db13787e805d3d273f3ffeccabc8bd149e6bbf6a4ab4744ed51ab7723c129007348dd7e0ff9e91df064e@213.32.53.181:30371",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// WhaleCoin test network.

var TestnetBootnodes = []string{

}



// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.

var DiscoveryV5Bootnodes = []string{}
