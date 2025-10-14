// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResolveStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDomainResolveStatisticsSummaryResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) *DescribeDomainResolveStatisticsSummaryResponseBody
	GetStatistics() []*DescribeDomainResolveStatisticsSummaryResponseBodyStatistics
	SetTotalItems(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody
	GetTotalPages() *int32
}

type DescribeDomainResolveStatisticsSummaryResponseBody struct {
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics.
	Statistics []*DescribeDomainResolveStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDomainResolveStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) GetStatistics() []*DescribeDomainResolveStatisticsSummaryResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDomainResolveStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) SetStatistics(v []*DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) *DescribeDomainResolveStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDomainResolveStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBody) Validate() error {
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

type DescribeDomainResolveStatisticsSummaryResponseBodyStatistics struct {
	// The number of DNS requests.
	//
	// example:
	//
	// 35509014
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- PUBLIC: hosted public domain name
	//
	// 	- CACHE: cache-accelerated domain name
	//
	// example:
	//
	// CACHE
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}

func (s DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) GetCount() *string {
	return s.Count
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) SetCount(v string) *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics {
	s.Count = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) SetDomainName(v string) *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) SetDomainType(v string) *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
