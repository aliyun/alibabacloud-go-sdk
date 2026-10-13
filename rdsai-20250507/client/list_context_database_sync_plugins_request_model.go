// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *ListContextDatabaseSyncPluginsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncPluginsRequest struct {
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseSyncPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncPluginsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncPluginsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncPluginsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsRequest) Validate() error {
	return dara.Validate(s)
}
