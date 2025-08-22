// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDeletedDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeDcdnDeletedDomainsResponseBodyDomains) *DescribeDcdnDeletedDomainsResponseBody
	GetDomains() *DescribeDcdnDeletedDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeDcdnDeletedDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnDeletedDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDcdnDeletedDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDcdnDeletedDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDcdnDeletedDomainsResponseBody struct {
	// The information about the accelerated domain name.
	Domains *DescribeDcdnDeletedDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
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
	// The ID of the request.
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

func (s DescribeDcdnDeletedDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponseBody) GetDomains() *DescribeDcdnDeletedDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnDeletedDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnDeletedDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDeletedDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDeletedDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetDomains(v *DescribeDcdnDeletedDomainsResponseBodyDomains) *DescribeDcdnDeletedDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetPageNumber(v int64) *DescribeDcdnDeletedDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetPageSize(v int64) *DescribeDcdnDeletedDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetRequestId(v string) *DescribeDcdnDeletedDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetTotalCount(v int64) *DescribeDcdnDeletedDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDeletedDomainsResponseBodyDomains struct {
	PageData []*DescribeDcdnDeletedDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomains) GetPageData() []*DescribeDcdnDeletedDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomains) SetPageData(v []*DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) *DescribeDcdnDeletedDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDeletedDomainsResponseBodyDomainsPageData struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when the accelerated domain name was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-10-28T11:05:52Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}
