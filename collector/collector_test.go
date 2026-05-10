// Copyright 2024 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsNoDataError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "ErrNoData",
			err:  ErrNoData,
			want: true,
		},
		{
			name: "wrapped ErrNoData",
			err:  fmt.Errorf("some context: %w", ErrNoData),
			want: true,
		},
		{
			name: "other error",
			err:  errors.New("other error"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsNoDataError(tt.err); got != tt.want {
				t.Errorf("IsNoDataError() = %v, want %v", got, tt.want)
			}
		})
	}
}
