// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohSubDomainStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDohSubDomainStatisticsSummaryResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) *DescribeDohSubDomainStatisticsSummaryResponseBody
	GetStatistics() []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics
	SetTotalItems(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody
	GetTotalPages() *int32
}

type DescribeDohSubDomainStatisticsSummaryResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics list.
	Statistics []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
	// Total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages returned.
	//
	// example:
	//
	// 50
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) GetStatistics() []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetStatistics(v []*DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDohSubDomainStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBody) Validate() error {
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics struct {
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
	// The subdomain.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// Total number of requests.
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

func (s DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetHttpCount() *int64 {
	return s.HttpCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetHttpsCount() *int64 {
	return s.HttpsCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetIpCount() *int64 {
	return s.IpCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetHttpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetHttpsCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetIpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.IpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetSubDomain(v string) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.SubDomain = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
