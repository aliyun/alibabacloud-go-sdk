// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetWorkspaceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *GetWorkspaceRequest
	GetWorkspaceId() *string
}

type GetWorkspaceRequest struct {
	// The region ID of the workspace.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace where the service resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceRequest) SetRegionId(v string) *GetWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceRequest) SetWorkspaceId(v string) *GetWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
