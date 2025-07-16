// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeletedDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeCdnDeletedDomainsResponseBodyDomains) *DescribeCdnDeletedDomainsResponseBody
	GetDomains() *DescribeCdnDeletedDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeCdnDeletedDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnDeletedDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCdnDeletedDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCdnDeletedDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeCdnDeletedDomainsResponseBody struct {
	// The list of accelerated domain names and the time each domain name was last modified.
	Domains *DescribeCdnDeletedDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The page number of the returned page, which is the same as the **PageNumber*	- parameter in request parameters.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names returned per page, which is the same as the **PageSize*	- parameter in request parameters.
	//
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA75AADB-5E25-4970-B480-EAA1F5658483
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names returned.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCdnDeletedDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeletedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeletedDomainsResponseBody) GetDomains() *DescribeCdnDeletedDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeCdnDeletedDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnDeletedDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDeletedDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDeletedDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCdnDeletedDomainsResponseBody) SetDomains(v *DescribeCdnDeletedDomainsResponseBodyDomains) *DescribeCdnDeletedDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBody) SetPageNumber(v int64) *DescribeCdnDeletedDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBody) SetPageSize(v int64) *DescribeCdnDeletedDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBody) SetRequestId(v string) *DescribeCdnDeletedDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBody) SetTotalCount(v int64) *DescribeCdnDeletedDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDeletedDomainsResponseBodyDomains struct {
	PageData []*DescribeCdnDeletedDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeCdnDeletedDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeletedDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomains) GetPageData() []*DescribeCdnDeletedDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomains) SetPageData(v []*DescribeCdnDeletedDomainsResponseBodyDomainsPageData) *DescribeCdnDeletedDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDeletedDomainsResponseBodyDomainsPageData struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when the accelerated domain name was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-10-28T11:05:52Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeCdnDeletedDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeletedDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeCdnDeletedDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeCdnDeletedDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}
