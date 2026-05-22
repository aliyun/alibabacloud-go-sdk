// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListClientCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListClientCertificatesRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListClientCertificatesRequest
	GetSiteId() *int64
}

type ListClientCertificatesRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListClientCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListClientCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListClientCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClientCertificatesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListClientCertificatesRequest) SetPageNumber(v int64) *ListClientCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClientCertificatesRequest) SetPageSize(v int64) *ListClientCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClientCertificatesRequest) SetSiteId(v int64) *ListClientCertificatesRequest {
	s.SiteId = &v
	return s
}

func (s *ListClientCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
