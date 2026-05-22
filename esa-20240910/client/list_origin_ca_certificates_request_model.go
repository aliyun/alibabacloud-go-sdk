// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginCaCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListOriginCaCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOriginCaCertificatesRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListOriginCaCertificatesRequest
	GetSiteId() *int64
}

type ListOriginCaCertificatesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**. Valid values: 1 to 500.
	//
	// example:
	//
	// 10
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

func (s ListOriginCaCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOriginCaCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListOriginCaCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOriginCaCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOriginCaCertificatesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginCaCertificatesRequest) SetPageNumber(v int64) *ListOriginCaCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOriginCaCertificatesRequest) SetPageSize(v int64) *ListOriginCaCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOriginCaCertificatesRequest) SetSiteId(v int64) *ListOriginCaCertificatesRequest {
	s.SiteId = &v
	return s
}

func (s *ListOriginCaCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
