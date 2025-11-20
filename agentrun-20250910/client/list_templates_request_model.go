// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesRequest
	GetPageSize() *int32
	SetTemplateType(v string) *ListTemplatesRequest
	GetTemplateType() *string
}

type ListTemplatesRequest struct {
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
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按模板类型过滤
	//
	// example:
	//
	// Browser
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
