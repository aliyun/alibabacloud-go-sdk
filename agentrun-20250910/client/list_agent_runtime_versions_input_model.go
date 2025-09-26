// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeVersionsInput interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int) *ListAgentRuntimeVersionsInput
	GetPageNumber() *int
	SetPageSize(v int) *ListAgentRuntimeVersionsInput
	GetPageSize() *int
}

type ListAgentRuntimeVersionsInput struct {
	// 页码
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListAgentRuntimeVersionsInput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeVersionsInput) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeVersionsInput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListAgentRuntimeVersionsInput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListAgentRuntimeVersionsInput) SetPageNumber(v int) *ListAgentRuntimeVersionsInput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimeVersionsInput) SetPageSize(v int) *ListAgentRuntimeVersionsInput {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimeVersionsInput) Validate() error {
	return dara.Validate(s)
}
