// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaListExportPagedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QuotaListExportPagedRequest
	GetCurrentPage() *int32
	SetLanguage(v string) *QuotaListExportPagedRequest
	GetLanguage() *string
	SetPageSize(v int32) *QuotaListExportPagedRequest
	GetPageSize() *int32
}

type QuotaListExportPagedRequest struct {
	// Pagination, current page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Multilingual Parameters, the default language is English.</br>
	//
	// en: English</br>
	//
	// zh: Chinese</br>
	//
	// ja: Japanese </br>
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Pagination, record number on each page, maximum 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QuotaListExportPagedRequest) String() string {
	return dara.Prettify(s)
}

func (s QuotaListExportPagedRequest) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QuotaListExportPagedRequest) GetLanguage() *string {
	return s.Language
}

func (s *QuotaListExportPagedRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuotaListExportPagedRequest) SetCurrentPage(v int32) *QuotaListExportPagedRequest {
	s.CurrentPage = &v
	return s
}

func (s *QuotaListExportPagedRequest) SetLanguage(v string) *QuotaListExportPagedRequest {
	s.Language = &v
	return s
}

func (s *QuotaListExportPagedRequest) SetPageSize(v int32) *QuotaListExportPagedRequest {
	s.PageSize = &v
	return s
}

func (s *QuotaListExportPagedRequest) Validate() error {
	return dara.Validate(s)
}
