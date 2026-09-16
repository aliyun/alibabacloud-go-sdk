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
	// The product API. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification
	//
	// - **ID_PERIOD**: ID card validity period verification
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition)
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition)
	//
	// - **MOBILE_2_META**: mobile number two-element verification
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition)
	//
	// - **MOBILE_DETECT**: phone number detection
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition)
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query
	//
	// - **VEHICLE_CHECK**: vehicle element verification
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_PERIOD
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. The value is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1737561599999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query. The timestamp is in milliseconds.
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
