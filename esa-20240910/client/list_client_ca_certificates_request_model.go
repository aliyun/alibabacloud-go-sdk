// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCaCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListClientCaCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListClientCaCertificatesRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListClientCaCertificatesRequest
	GetSiteId() *int64
}

type ListClientCaCertificatesRequest struct {
	// The page number. Valid values: 1 to 500.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListClientCaCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientCaCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListClientCaCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListClientCaCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClientCaCertificatesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListClientCaCertificatesRequest) SetPageNumber(v int64) *ListClientCaCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClientCaCertificatesRequest) SetPageSize(v int64) *ListClientCaCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClientCaCertificatesRequest) SetSiteId(v int64) *ListClientCaCertificatesRequest {
	s.SiteId = &v
	return s
}

func (s *ListClientCaCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
