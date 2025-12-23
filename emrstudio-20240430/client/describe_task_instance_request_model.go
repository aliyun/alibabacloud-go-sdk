// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeTaskInstanceRequest
	GetWorkspaceId() *string
}

type DescribeTaskInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeTaskInstanceRequest) SetWorkspaceId(v string) *DescribeTaskInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
