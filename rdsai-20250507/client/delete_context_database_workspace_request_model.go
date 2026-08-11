// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteContextDatabaseWorkspaceRequest
	GetWorkspaceId() *string
}

type DeleteContextDatabaseWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteContextDatabaseWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteContextDatabaseWorkspaceRequest) SetWorkspaceId(v string) *DeleteContextDatabaseWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
