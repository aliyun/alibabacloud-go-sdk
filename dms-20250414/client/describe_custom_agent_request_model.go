// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *DescribeCustomAgentRequest
	GetCustomAgentId() *string
	SetWorkspaceId(v string) *DescribeCustomAgentRequest
	GetWorkspaceId() *string
}

type DescribeCustomAgentRequest struct {
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DescribeCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeCustomAgentRequest) SetCustomAgentId(v string) *DescribeCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *DescribeCustomAgentRequest) SetWorkspaceId(v string) *DescribeCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
