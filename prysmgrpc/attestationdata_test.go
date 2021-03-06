// Copyright © 2020 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prysmgrpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/attestantio/go-eth2-client/prysmgrpc"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/stretchr/testify/require"
)

func TestAttestationData(t *testing.T) {
	tests := []struct {
		name           string
		slot           int64
		committeeIndex uint64
	}{
		{
			name: "Good",
			slot: -1,
		},
	}

	service, err := prysmgrpc.New(context.Background(),
		prysmgrpc.WithAddress(os.Getenv("PRYSMGRPC_ADDRESS")),
		prysmgrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var slot uint64
			var err error
			if test.slot == -1 {
				slot, err = service.CurrentSlot(context.Background())
				require.NoError(t, err)
			} else {
				slot = uint64(test.slot)
			}
			attestationData, err := service.AttestationData(context.Background(), spec.Slot(slot), spec.CommitteeIndex(test.committeeIndex))
			require.NoError(t, err)
			require.NotNil(t, attestationData)
		})
	}
}
