// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohDomainStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDohDomainStatisticsSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDohDomainStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDohDomainStatisticsSummaryResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics) *DescribeDohDomainStatisticsSummaryResponseBody
	GetStatistics() []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics
	SetTotalItems(v int32) *DescribeDohDomainStatisticsSummaryResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDohDomainStatisticsSummaryResponseBody
	GetTotalPages() *int32
}

type DescribeDohDomainStatisticsSummaryResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics list.
	Statistics []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 300
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 50
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDohDomainStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) GetStatistics() []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetStatistics(v []*DescribeDohDomainStatisticsSummaryResponseBodyStatistics) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDohDomainStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDohDomainStatisticsSummaryResponseBodyStatistics struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The number of HTTP requests.
	//
	// example:
	//
	// 3141592653
	HttpCount *int64 `json:"HttpCount,omitempty" xml:"HttpCount,omitempty"`
	// The number of HTTPS requests.
	//
	// example:
	//
	// 3141592653
	HttpsCount *int64 `json:"HttpsCount,omitempty" xml:"HttpsCount,omitempty"`
	// The number of IP addresses.
	//
	// example:
	//
	// 20
	IpCount *int64 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 14141592653
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of IPv4-based HTTP requests.
	//
	// example:
	//
	// 3141592653
	V4HttpCount *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	// The number of IPv4-based HTTPS requests.
	//
	// example:
	//
	// 3141592653
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	// The number of IPv6-based HTTP requests.
	//
	// example:
	//
	// 3141592653
	V6HttpCount *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	// The number of IPv6-based HTTPS requests.
	//
	// example:
	//
	// 3141592653
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohDomainStatisticsSummaryResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetHttpCount() *int64 {
	return s.HttpCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetHttpsCount() *int64 {
	return s.HttpsCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetIpCount() *int64 {
	return s.IpCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetDomainName(v string) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.DomainName = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetHttpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetHttpsCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetIpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.IpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
