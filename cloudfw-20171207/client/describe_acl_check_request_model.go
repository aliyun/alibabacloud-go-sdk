// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAclCheckRequest
	GetLang() *string
	SetPageNo(v int32) *DescribeAclCheckRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAclCheckRequest
	GetPageSize() *int32
	SetTaskId(v string) *DescribeAclCheckRequest
	GetTaskId() *string
}

type DescribeAclCheckRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 132
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAclCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclCheckRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAclCheckRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAclCheckRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAclCheckRequest) SetLang(v string) *DescribeAclCheckRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclCheckRequest) SetPageNo(v int32) *DescribeAclCheckRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAclCheckRequest) SetPageSize(v int32) *DescribeAclCheckRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAclCheckRequest) SetTaskId(v string) *DescribeAclCheckRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeAclCheckRequest) Validate() error {
	return dara.Validate(s)
}
