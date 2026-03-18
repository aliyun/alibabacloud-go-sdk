// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdMemberListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeRdMemberListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeRdMemberListRequest
	GetPageSize() *int32
	SetResourceDirectoryId(v string) *DescribeRdMemberListRequest
	GetResourceDirectoryId() *string
}

type DescribeRdMemberListRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-x9bLhd
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s DescribeRdMemberListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdMemberListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeRdMemberListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRdMemberListRequest) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *DescribeRdMemberListRequest) SetPageNo(v int32) *DescribeRdMemberListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRdMemberListRequest) SetPageSize(v int32) *DescribeRdMemberListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRdMemberListRequest) SetResourceDirectoryId(v string) *DescribeRdMemberListRequest {
	s.ResourceDirectoryId = &v
	return s
}

func (s *DescribeRdMemberListRequest) Validate() error {
	return dara.Validate(s)
}
