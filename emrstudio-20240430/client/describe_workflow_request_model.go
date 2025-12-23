// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeWorkflowRequest
	GetWorkspaceId() *string
}

type DescribeWorkflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeWorkflowRequest) SetWorkspaceId(v string) *DescribeWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
