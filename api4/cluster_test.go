// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"testing"
)

func TestGetClusterStatus(t *testing.T) {
	th := Setup(t)
	defer th.TearDown()
	th.InitBasic().InitSystemAdmin()

	_, resp := th.Client.GetClusterStatus()
	CheckForbiddenStatus(t, resp)

	infos, resp := th.SystemAdminClient.GetClusterStatus()
	CheckNoError(t, resp)

	if infos == nil {
		t.Fatal("should not be nil")
	}
}
