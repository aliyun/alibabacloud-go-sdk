// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeWorkflowInstanceRequest
	GetWorkspaceId() *string
}

type DescribeWorkflowInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeWorkflowInstanceRequest) SetWorkspaceId(v string) *DescribeWorkflowInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
