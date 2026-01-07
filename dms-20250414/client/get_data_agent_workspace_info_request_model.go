// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentWorkspaceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *GetDataAgentWorkspaceInfoRequest
	GetDMSUnit() *string
	SetWorkspaceId(v string) *GetDataAgentWorkspaceInfoRequest
	GetWorkspaceId() *string
}

type GetDataAgentWorkspaceInfoRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDataAgentWorkspaceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentWorkspaceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentWorkspaceInfoRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetDataAgentWorkspaceInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDataAgentWorkspaceInfoRequest) SetDMSUnit(v string) *GetDataAgentWorkspaceInfoRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoRequest) SetWorkspaceId(v string) *GetDataAgentWorkspaceInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoRequest) Validate() error {
	return dara.Validate(s)
}
