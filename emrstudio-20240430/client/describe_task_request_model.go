// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeTaskRequest
	GetWorkspaceId() *string
}

type DescribeTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeTaskRequest) SetWorkspaceId(v string) *DescribeTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeTaskRequest) Validate() error {
	return dara.Validate(s)
}
