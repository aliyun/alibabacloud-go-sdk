// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *DescribeMetaStatisticsListRequest
	GetApi() *string
	SetEndDate(v int64) *DescribeMetaStatisticsListRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeMetaStatisticsListRequest
	GetStartDate() *int64
}

type DescribeMetaStatisticsListRequest struct {
	// The commodity code. Valid values:
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
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// MOBILE_ONLINE_LENGTH
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The end time of the query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1750694399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time of the query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760112000000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeMetaStatisticsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsListRequest) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaStatisticsListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeMetaStatisticsListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeMetaStatisticsListRequest) SetApi(v string) *DescribeMetaStatisticsListRequest {
	s.Api = &v
	return s
}

func (s *DescribeMetaStatisticsListRequest) SetEndDate(v int64) *DescribeMetaStatisticsListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeMetaStatisticsListRequest) SetStartDate(v int64) *DescribeMetaStatisticsListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMetaStatisticsListRequest) Validate() error {
	return dara.Validate(s)
}
