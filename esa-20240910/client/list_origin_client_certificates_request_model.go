// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginClientCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListOriginClientCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOriginClientCertificatesRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListOriginClientCertificatesRequest
	GetSiteId() *int64
}

type ListOriginClientCertificatesRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
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
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListOriginClientCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOriginClientCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListOriginClientCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOriginClientCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOriginClientCertificatesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginClientCertificatesRequest) SetPageNumber(v int64) *ListOriginClientCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOriginClientCertificatesRequest) SetPageSize(v int64) *ListOriginClientCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOriginClientCertificatesRequest) SetSiteId(v int64) *ListOriginClientCertificatesRequest {
	s.SiteId = &v
	return s
}

func (s *ListOriginClientCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
