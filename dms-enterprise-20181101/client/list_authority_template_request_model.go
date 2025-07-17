// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAuthorityTemplateRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorityTemplateRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListAuthorityTemplateRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorityTemplateRequest
	GetTid() *int64
}

type ListAuthorityTemplateRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAuthorityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorityTemplateRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorityTemplateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorityTemplateRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorityTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorityTemplateRequest) SetPageNumber(v int32) *ListAuthorityTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorityTemplateRequest) SetPageSize(v int32) *ListAuthorityTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorityTemplateRequest) SetSearchKey(v string) *ListAuthorityTemplateRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorityTemplateRequest) SetTid(v int64) *ListAuthorityTemplateRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
