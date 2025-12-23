// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManualTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeManualTaskRequest
	GetWorkspaceId() *string
}

type DescribeManualTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeManualTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeManualTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeManualTaskRequest) SetWorkspaceId(v string) *DescribeManualTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeManualTaskRequest) Validate() error {
	return dara.Validate(s)
}
