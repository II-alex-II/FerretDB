// Copyright 2021 FerretDB Inc.
//
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

package testutil

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/FerretDB/FerretDB/internal/fjson"
	"github.com/FerretDB/FerretDB/internal/types"
)

// Dump returns string representation for debugging.
func Dump[T types.Type](tb testing.TB, o T) string {
	b, err := fjson.Marshal(o)
	require.NoError(tb, err)

	dst := bytes.NewBuffer(make([]byte, 0, len(b)))
	err = json.Indent(dst, b, "", "  ")
	require.NoError(tb, err)
	return dst.String()
}
