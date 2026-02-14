// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *OperateCustomAgentRequest
	GetCustomAgentId() *string
	SetOperateType(v string) *OperateCustomAgentRequest
	GetOperateType() *string
	SetWorkspaceId(v string) *OperateCustomAgentRequest
	GetWorkspaceId() *string
}

type OperateCustomAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s OperateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *OperateCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *OperateCustomAgentRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *OperateCustomAgentRequest) SetCustomAgentId(v string) *OperateCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *OperateCustomAgentRequest) SetOperateType(v string) *OperateCustomAgentRequest {
	s.OperateType = &v
	return s
}

func (s *OperateCustomAgentRequest) SetWorkspaceId(v string) *OperateCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *OperateCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
