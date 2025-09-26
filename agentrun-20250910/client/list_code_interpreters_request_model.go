// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpretersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterName(v string) *ListCodeInterpretersRequest
	GetCodeInterpreterName() *string
	SetPageNumber(v int32) *ListCodeInterpretersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCodeInterpretersRequest
	GetPageSize() *int32
}

type ListCodeInterpretersRequest struct {
	// 根据代码解释器实例名称进行模糊匹配过滤
	//
	// example:
	//
	// code
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// 当前页码，从1开始计数
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的记录数量
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListCodeInterpretersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpretersRequest) GoString() string {
	return s.String()
}

func (s *ListCodeInterpretersRequest) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *ListCodeInterpretersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCodeInterpretersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCodeInterpretersRequest) SetCodeInterpreterName(v string) *ListCodeInterpretersRequest {
	s.CodeInterpreterName = &v
	return s
}

func (s *ListCodeInterpretersRequest) SetPageNumber(v int32) *ListCodeInterpretersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCodeInterpretersRequest) SetPageSize(v int32) *ListCodeInterpretersRequest {
	s.PageSize = &v
	return s
}

func (s *ListCodeInterpretersRequest) Validate() error {
	return dara.Validate(s)
}
