// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DescribeProjectRequest
	GetWorkspaceId() *string
}

type DescribeProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeProjectRequest) SetWorkspaceId(v string) *DescribeProjectRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeProjectRequest) Validate() error {
	return dara.Validate(s)
}
