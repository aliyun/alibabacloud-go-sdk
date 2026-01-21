// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteOriginClientCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListSiteOriginClientCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSiteOriginClientCertificatesRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListSiteOriginClientCertificatesRequest
	GetSiteId() *int64
}

type ListSiteOriginClientCertificatesRequest struct {
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
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListSiteOriginClientCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSiteOriginClientCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListSiteOriginClientCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSiteOriginClientCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSiteOriginClientCertificatesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListSiteOriginClientCertificatesRequest) SetPageNumber(v int64) *ListSiteOriginClientCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSiteOriginClientCertificatesRequest) SetPageSize(v int64) *ListSiteOriginClientCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSiteOriginClientCertificatesRequest) SetSiteId(v int64) *ListSiteOriginClientCertificatesRequest {
	s.SiteId = &v
	return s
}

func (s *ListSiteOriginClientCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
