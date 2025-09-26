// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpretersInput interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterName(v string) *ListCodeInterpretersInput
	GetCodeInterpreterName() *string
	SetPageNumber(v int) *ListCodeInterpretersInput
	GetPageNumber() *int
	SetPageSize(v int) *ListCodeInterpretersInput
	GetPageSize() *int
}

type ListCodeInterpretersInput struct {
	// 按代码解释器名称过滤
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	PageNumber          *int    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize            *int    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListCodeInterpretersInput) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpretersInput) GoString() string {
	return s.String()
}

func (s *ListCodeInterpretersInput) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *ListCodeInterpretersInput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListCodeInterpretersInput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListCodeInterpretersInput) SetCodeInterpreterName(v string) *ListCodeInterpretersInput {
	s.CodeInterpreterName = &v
	return s
}

func (s *ListCodeInterpretersInput) SetPageNumber(v int) *ListCodeInterpretersInput {
	s.PageNumber = &v
	return s
}

func (s *ListCodeInterpretersInput) SetPageSize(v int) *ListCodeInterpretersInput {
	s.PageSize = &v
	return s
}

func (s *ListCodeInterpretersInput) Validate() error {
	return dara.Validate(s)
}
