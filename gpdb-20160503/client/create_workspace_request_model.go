// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateWorkspaceRequest
	GetRegionId() *string
	SetWorkspaceName(v string) *CreateWorkspaceRequest
	GetWorkspaceName() *string
}

type CreateWorkspaceRequest struct {
	// The region ID.
	//
	// > Currently, only four units are supported.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workspace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-first-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateWorkspaceRequest) SetRegionId(v string) *CreateWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
