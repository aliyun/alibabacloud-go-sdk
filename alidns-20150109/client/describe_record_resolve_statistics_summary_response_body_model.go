// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordResolveStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRecordResolveStatisticsSummaryResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) *DescribeRecordResolveStatisticsSummaryResponseBody
	GetStatistics() []*DescribeRecordResolveStatisticsSummaryResponseBodyStatistics
	SetTotalItems(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody
	GetTotalPages() *int32
}

type DescribeRecordResolveStatisticsSummaryResponseBody struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 500**. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics.
	Statistics []*DescribeRecordResolveStatisticsSummaryResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
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

func (s DescribeRecordResolveStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordResolveStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) GetStatistics() []*DescribeRecordResolveStatisticsSummaryResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) SetPageNumber(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeRecordResolveStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) SetStatistics(v []*DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) *DescribeRecordResolveStatisticsSummaryResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) SetTotalItems(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) SetTotalPages(v int32) *DescribeRecordResolveStatisticsSummaryResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBody) Validate() error {
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

type DescribeRecordResolveStatisticsSummaryResponseBodyStatistics struct {
	// The number of DNS requests.
	//
	// example:
	//
	// 330
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The subdomain name.
	//
	// example:
	//
	// tes.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. The parameter value is not case-sensitive. Valid values:
	//
	// 	- PUBLIC (default): hosted public domain name
	//
	// 	- CACHE: cache-accelerated domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The subdomain.
	//
	// example:
	//
	// test.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) GetCount() *string {
	return s.Count
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) SetCount(v string) *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics {
	s.Count = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) SetDomainName(v string) *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) SetDomainType(v string) *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics {
	s.DomainType = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) SetSubDomain(v string) *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics {
	s.SubDomain = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
