// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v int64) *GetWorkspaceRequest
	GetWorkspaceId() *int64
}

type GetWorkspaceRequest struct {
	// The ID of the DMS workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetWorkspaceRequest) SetWorkspaceId(v int64) *GetWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
