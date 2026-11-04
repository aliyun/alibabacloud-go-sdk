// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTimeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryTimeTemplateRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryTimeTemplateRequest
	GetPageSize() *int32
}

type QueryTimeTemplateRequest struct {
	// Specifies the page number to return. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return per page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTimeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryTimeTemplateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTimeTemplateRequest) SetCurrentPage(v int32) *QueryTimeTemplateRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTimeTemplateRequest) SetPageSize(v int32) *QueryTimeTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTimeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
