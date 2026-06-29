// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostname(v string) *ListCustomHostnamesRequest
	GetHostname() *string
	SetNameMatchType(v string) *ListCustomHostnamesRequest
	GetNameMatchType() *string
	SetPageNumber(v int32) *ListCustomHostnamesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomHostnamesRequest
	GetPageSize() *int32
	SetRecordId(v int64) *ListCustomHostnamesRequest
	GetRecordId() *int64
	SetSiteId(v int64) *ListCustomHostnamesRequest
	GetSiteId() *int64
	SetStatus(v string) *ListCustomHostnamesRequest
	GetStatus() *string
}

type ListCustomHostnamesRequest struct {
	// The SaaS domain name.
	//
	// example:
	//
	// custom.site.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The search match mode for the SaaS domain name. Default value: exact. Valid values:
	//
	// - **prefix**: prefix match.
	//
	// - **suffix**: suffix match.
	//
	// - **exact**: exact match.
	//
	// - **fuzzy**: fuzzy match.
	//
	// example:
	//
	// exact
	NameMatchType *string `json:"NameMatchType,omitempty" xml:"NameMatchType,omitempty"`
	// The page number for a paged query. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Valid values: 1 to 500.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The associated record ID. You can call the [ListRecords](https://help.aliyun.com/document_detail/2850265.html) operation to obtain the record ID.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 744571165985008
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The SaaS domain name status. Valid values:
	//
	// - **pending**: pending verification.
	//
	// - **active**: activated.
	//
	// - **conflicted**: occupied.
	//
	// - **offline**: offline.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCustomHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomHostnamesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomHostnamesRequest) GetHostname() *string {
	return s.Hostname
}

func (s *ListCustomHostnamesRequest) GetNameMatchType() *string {
	return s.NameMatchType
}

func (s *ListCustomHostnamesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomHostnamesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomHostnamesRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListCustomHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCustomHostnamesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCustomHostnamesRequest) SetHostname(v string) *ListCustomHostnamesRequest {
	s.Hostname = &v
	return s
}

func (s *ListCustomHostnamesRequest) SetNameMatchType(v string) *ListCustomHostnamesRequest {
	s.NameMatchType = &v
	return s
}

func (s *ListCustomHostnamesRequest) SetPageNumber(v int32) *ListCustomHostnamesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomHostnamesRequest) SetPageSize(v int32) *ListCustomHostnamesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomHostnamesRequest) SetRecordId(v int64) *ListCustomHostnamesRequest {
	s.RecordId = &v
	return s
}

func (s *ListCustomHostnamesRequest) SetSiteId(v int64) *ListCustomHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *ListCustomHostnamesRequest) SetStatus(v string) *ListCustomHostnamesRequest {
	s.Status = &v
	return s
}

func (s *ListCustomHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
