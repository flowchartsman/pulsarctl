// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ledger

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCmd(t *testing.T) {
	o := doListCmdTest(t)
	assert.True(t, strings.Contains(o, "0"))
}

func doListCmdTest(t *testing.T) string {
	args := []string{"list"}
	out, execErr, nameErr, err := testLedgerCommands(listCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, nameErr)
	assert.Nil(t, execErr)
	return out.String()
}
