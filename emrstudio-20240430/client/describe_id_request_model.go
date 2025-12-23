// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeIdRequest
	GetId() *string
	SetWorkspaceId(v string) *DescribeIdRequest
	GetWorkspaceId() *string
}

type DescribeIdRequest struct {
	// id
	//
	// This parameter is required.
	//
	// example:
	//
	// p-123****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DescribeIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdRequest) GetId() *string {
	return s.Id
}

func (s *DescribeIdRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeIdRequest) SetId(v string) *DescribeIdRequest {
	s.Id = &v
	return s
}

func (s *DescribeIdRequest) SetWorkspaceId(v string) *DescribeIdRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeIdRequest) Validate() error {
	return dara.Validate(s)
}
