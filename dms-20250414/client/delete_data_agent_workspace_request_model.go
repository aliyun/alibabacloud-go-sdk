// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *DeleteDataAgentWorkspaceRequest
	GetDMSUnit() *string
	SetWorkspaceId(v string) *DeleteDataAgentWorkspaceRequest
	GetWorkspaceId() *string
}

type DeleteDataAgentWorkspaceRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataAgentWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentWorkspaceRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *DeleteDataAgentWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataAgentWorkspaceRequest) SetDMSUnit(v string) *DeleteDataAgentWorkspaceRequest {
	s.DMSUnit = &v
	return s
}

func (s *DeleteDataAgentWorkspaceRequest) SetWorkspaceId(v string) *DeleteDataAgentWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataAgentWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
