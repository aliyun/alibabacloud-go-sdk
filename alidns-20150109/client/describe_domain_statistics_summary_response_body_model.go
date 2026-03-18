// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDomainStatisticsSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDomainStatisticsSummaryResponseBody
	GetRequestId() *string
	SetStatistics(v *DescribeDomainStatisticsSummaryResponseBodyStatistics) *DescribeDomainStatisticsSummaryResponseBody
	GetStatistics() *DescribeDomainStatisticsSummaryResponseBodyStatistics
	SetTotalItems(v int32) *DescribeDomainStatisticsSummaryResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDomainStatisticsSummaryResponseBody
	GetTotalPages() *int32
}

type DescribeDomainStatisticsSummaryResponseBody struct {
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
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CC625C21-8832-4683-BF10-C3CFB1A4FA13
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics *DescribeDomainStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The total number of data records.
	//
	// example:
	//
	// 68
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 14
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDomainStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainStatisticsSummaryResponseBody) GetStatistics() *DescribeDomainStatisticsSummaryResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDomainStatisticsSummaryResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDomainStatisticsSummaryResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeDomainStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetStatistics(v *DescribeDomainStatisticsSummaryResponseBodyStatistics) *DescribeDomainStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeDomainStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBody) Validate() error {
	if s.Statistics != nil {
		if err := s.Statistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainStatisticsSummaryResponseBodyStatistics struct {
	Statistic []*DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatisticsSummaryResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatistics) GetStatistic() []*DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic {
	return s.Statistic
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatistics) SetStatistic(v []*DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) *DescribeDomainStatisticsSummaryResponseBodyStatistics {
	s.Statistic = v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatistics) Validate() error {
	if s.Statistic != nil {
		for _, item := range s.Statistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic struct {
	Count                 *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainType            *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	ResolveAnalysisStatus *string `json:"resolveAnalysisStatus,omitempty" xml:"resolveAnalysisStatus,omitempty"`
}

func (s DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) GetResolveAnalysisStatus() *string {
	return s.ResolveAnalysisStatus
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) SetCount(v int64) *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic {
	s.Count = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) SetDomainName(v string) *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) SetDomainType(v string) *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) SetResolveAnalysisStatus(v string) *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic {
	s.ResolveAnalysisStatus = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponseBodyStatisticsStatistic) Validate() error {
	return dara.Validate(s)
}
