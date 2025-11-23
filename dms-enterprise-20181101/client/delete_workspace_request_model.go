// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v int64) *DeleteWorkspaceRequest
	GetWorkspaceId() *int64
}

type DeleteWorkspaceRequest struct {
	// The ID of the DMS workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceRequest) SetWorkspaceId(v int64) *DeleteWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
