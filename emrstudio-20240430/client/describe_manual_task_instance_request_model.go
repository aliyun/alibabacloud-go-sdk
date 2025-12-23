// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManualTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeManualTaskInstanceRequest
	GetWorkspaceId() *string
}

type DescribeManualTaskInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeManualTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeManualTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeManualTaskInstanceRequest) SetWorkspaceId(v string) *DescribeManualTaskInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeManualTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
