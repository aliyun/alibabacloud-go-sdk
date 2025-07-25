// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRecordStatisticsSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRecordStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRecordStatisticsSummaryResponseBody
	GetRequestId() *string
	SetStatistics(v *DescribeRecordStatisticsSummaryResponseBodyStatistics) *DescribeRecordStatisticsSummaryResponseBody
	GetStatistics() *DescribeRecordStatisticsSummaryResponseBodyStatistics
	SetTotalItems(v int32) *DescribeRecordStatisticsSummaryResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeRecordStatisticsSummaryResponseBody
	GetTotalPages() *int32
}

type DescribeRecordStatisticsSummaryResponseBody struct {
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E49F0023-4A98-486F-8BA3-6003D5664105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The DNS requests.
	Statistics *DescribeRecordStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeRecordStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRecordStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecordStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordStatisticsSummaryResponseBody) GetStatistics() *DescribeRecordStatisticsSummaryResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeRecordStatisticsSummaryResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeRecordStatisticsSummaryResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeRecordStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetStatistics(v *DescribeRecordStatisticsSummaryResponseBodyStatistics) *DescribeRecordStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeRecordStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordStatisticsSummaryResponseBodyStatistics struct {
	Statistic []*DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeRecordStatisticsSummaryResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatistics) GetStatistic() []*DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic {
	return s.Statistic
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatistics) SetStatistic(v []*DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) *DescribeRecordStatisticsSummaryResponseBodyStatistics {
	s.Statistic = v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic struct {
	// The number of DNS requests.
	//
	// example:
	//
	// 838711553
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The subdomain.
	//
	// example:
	//
	// t1.alitest2.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) SetCount(v int64) *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic {
	s.Count = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) SetSubDomain(v string) *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic {
	s.SubDomain = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponseBodyStatisticsStatistic) Validate() error {
	return dara.Validate(s)
}
