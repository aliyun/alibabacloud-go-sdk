// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteWorkspaceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteWorkspaceRequest
	GetWorkspaceId() *string
}

type DeleteWorkspaceRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceRequest) SetRegionId(v string) *DeleteWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWorkspaceRequest) SetWorkspaceId(v string) *DeleteWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
