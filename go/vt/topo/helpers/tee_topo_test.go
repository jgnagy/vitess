/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helpers

import (
	"testing"

	"vitess.io/vitess/go/test/utils"
	"vitess.io/vitess/go/vt/topo"
	"vitess.io/vitess/go/vt/topo/memorytopo"
	"vitess.io/vitess/go/vt/topo/test"
)

func TestTeeTopo(t *testing.T) {
	ctx := utils.LeakCheckContext(t)
	test.TopoServerTestSuite(t, ctx, func() *topo.Server {
		s1 := memorytopo.NewServer(ctx, test.LocalCellName)
		s2 := memorytopo.NewServer(ctx, test.LocalCellName)
		tee, err := NewTee(s1, s2, false)
		if err != nil {
			t.Fatalf("NewTee() failed: %v", err)
		}
		return tee
	}, []string{"checkTryLock", "checkShardWithLock"})
}
