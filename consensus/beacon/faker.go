// Copyright 2023 The gotfx Authors
// This file is part of the gotfx library.
//
// The gotfx library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gotfx library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gotfx library. If not, see <http://www.gnu.org/licenses/>.

package beacon

import (
	"math/big"

	"github.com/gitshock-labs/gotfx/consensus"
	"github.com/gitshock-labs/gotfx/core/types"
)

// NewFaker creates a fake consensus engine for testing.
// The fake engine simulates a merged network.
// It can not be used to test the merge transition.
// This type is needed since the fakeChainReader can not be used with
// a normal beacon consensus engine.
func NewFaker() consensus.Engine {
	return new(faker)
}

type faker struct {
	Beacon
}

func (f *faker) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return beaconDifficulty
}
