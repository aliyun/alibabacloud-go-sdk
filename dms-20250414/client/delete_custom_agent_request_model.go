// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *DeleteCustomAgentRequest
	GetCustomAgentId() *string
	SetWorkspaceId(v string) *DeleteCustomAgentRequest
	GetWorkspaceId() *string
}

type DeleteCustomAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DeleteCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCustomAgentRequest) SetCustomAgentId(v string) *DeleteCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *DeleteCustomAgentRequest) SetWorkspaceId(v string) *DeleteCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
