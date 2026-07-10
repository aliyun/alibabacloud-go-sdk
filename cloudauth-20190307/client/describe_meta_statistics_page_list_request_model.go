// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *DescribeMetaStatisticsPageListRequest
	GetApi() *string
	SetCurrentPage(v int32) *DescribeMetaStatisticsPageListRequest
	GetCurrentPage() *int32
	SetEndDate(v int64) *DescribeMetaStatisticsPageListRequest
	GetEndDate() *int64
	SetPageSize(v int32) *DescribeMetaStatisticsPageListRequest
	GetPageSize() *int32
	SetStartDate(v int64) *DescribeMetaStatisticsPageListRequest
	GetStartDate() *int64
}

type DescribeMetaStatisticsPageListRequest struct {
	// Product API:
	//
	// - **ID_CARD_2_META**: ID Card Two-Element Verification
	//
	// - **ID_PERIOD**: ID Card Validity Verification Period
	//
	// - **MOBILE_ONLINE_LENGTH**: Mobile Online Duration
	//
	// - **MOBILE_ONLINE_STATUS**: Mobile Online Status
	//
	// - **MOBILE_3_META_SIMPLE**: Mobile Number Three-Element Verification (Simple)
	//
	// - **MOBILE_3_META**: Mobile Number Three-Element Verification (Detailed)
	//
	// - **MOBILE_2_META**: Mobile Number Two-Element Verification
	//
	// - **BANK_CARD_N_META**: Bank Card Verification (Detailed)
	//
	// - **MOBILE_DETECT**: Number Detection
	//
	// - **VEHICLE_N_META**: Vehicle Element Verification (Enhanced)
	//
	// - **VEHICLE_PENTA_INFO**: Vehicle Five-Element Information Recognition
	//
	// - **VEHICLE_LICENSE_INFO**: Vehicle Information Recognition
	//
	// - **VEHICLE_INSURE_DATE**: Vehicle Insurance Date Query
	//
	// - **VEHICLE_CHECK**: Vehicle Element Verification
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_PERIOD
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Query end time. Unix timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1737561599999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Number of data entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Query start time. The timestamp is in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760112000000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeMetaStatisticsPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsPageListRequest) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaStatisticsPageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeMetaStatisticsPageListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeMetaStatisticsPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetaStatisticsPageListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeMetaStatisticsPageListRequest) SetApi(v string) *DescribeMetaStatisticsPageListRequest {
	s.Api = &v
	return s
}

func (s *DescribeMetaStatisticsPageListRequest) SetCurrentPage(v int32) *DescribeMetaStatisticsPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeMetaStatisticsPageListRequest) SetEndDate(v int64) *DescribeMetaStatisticsPageListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeMetaStatisticsPageListRequest) SetPageSize(v int32) *DescribeMetaStatisticsPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaStatisticsPageListRequest) SetStartDate(v int64) *DescribeMetaStatisticsPageListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMetaStatisticsPageListRequest) Validate() error {
	return dara.Validate(s)
}
