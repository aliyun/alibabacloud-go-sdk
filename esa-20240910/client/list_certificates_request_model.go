// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListCertificatesRequest
	GetKeyword() *string
	SetPageNumber(v int64) *ListCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCertificatesRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListCertificatesRequest
	GetSiteId() *int64
	SetValidOnly(v bool) *ListCertificatesRequest
	GetValidOnly() *bool
}

type ListCertificatesRequest struct {
	// The keyword that is used for the search.
	//
	// example:
	//
	// example
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 3
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to return only valid certificates.
	//
	// example:
	//
	// 1
	ValidOnly *bool `json:"ValidOnly,omitempty" xml:"ValidOnly,omitempty"`
}

func (s ListCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListCertificatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCertificatesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCertificatesRequest) GetValidOnly() *bool {
	return s.ValidOnly
}

func (s *ListCertificatesRequest) SetKeyword(v string) *ListCertificatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListCertificatesRequest) SetPageNumber(v int64) *ListCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCertificatesRequest) SetPageSize(v int64) *ListCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCertificatesRequest) SetSiteId(v int64) *ListCertificatesRequest {
	s.SiteId = &v
	return s
}

func (s *ListCertificatesRequest) SetValidOnly(v bool) *ListCertificatesRequest {
	s.ValidOnly = &v
	return s
}

func (s *ListCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
