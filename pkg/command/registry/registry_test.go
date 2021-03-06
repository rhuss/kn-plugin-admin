// Copyright 2020 The Knative Authors
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

package registry

import (
	"testing"

	"knative.dev/kn-plugin-admin/pkg/testutil"

	"gotest.tools/assert"
)

func TestNewPrivateRegistryCmd(t *testing.T) {
	p, client := testutil.NewTestAdminParams()
	assert.Check(t, client != nil)
	cmd := NewPrivateRegistryCmd(p)
	assert.Check(t, cmd.HasSubCommands(), "cmd registry should have subcommands")
	assert.Equal(t, 3, len(cmd.Commands()), "registry command should have 3 subcommands")

	_, _, err := cmd.Find([]string{"add"})
	assert.NilError(t, err, "registry command should have add subcommand")

	_, _, err = cmd.Find([]string{"remove"})
	assert.NilError(t, err, "registry command should have remove subcommand")

	_, _, err = cmd.Find([]string{"list"})
	assert.NilError(t, err, "registry command should have list subcommand")
}
