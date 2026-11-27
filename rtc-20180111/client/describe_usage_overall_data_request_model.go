// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageOverallDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeUsageOverallDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeUsageOverallDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeUsageOverallDataRequest
	GetStartDate() *int64
	SetTypes(v string) *DescribeUsageOverallDataRequest
	GetTypes() *string
}

type DescribeUsageOverallDataRequest struct {
	// The app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2hz****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The metric types to query. Separate multiple values with commas (,). Valid values:
	//
	// - **TOTAL_CALL_DURATION**: total duration.
	//
	// - **VIDEO_CALL_DURATION**: video communication duration.
	//
	// - **AUDIO_CALL_DURATION**: audio communication duration.
	//
	// - **CALL_CHANNEL_COUNT**: communication channel count.
	//
	// - **HIGHLY_CONCURRENT_CHANNEL_COUNT**: highly concurrent communication channel count.
	//
	// - **CHANNEL_CONCURRENT_PEAK**: peak concurrent channel count.
	//
	// - **ONLINE_USER_PEAK**: peak online user count.
	//
	// - **TOTAL_CALL_USER**: total call users.
	//
	// - **TOTAL_INOUT_NUM**: total inbound and outbound visits.
	//
	// This parameter is required.
	//
	// example:
	//
	// ONLINE_USER_PEAK
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeUsageOverallDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeUsageOverallDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeUsageOverallDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeUsageOverallDataRequest) GetTypes() *string {
	return s.Types
}

func (s *DescribeUsageOverallDataRequest) SetAppId(v string) *DescribeUsageOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetEndDate(v int64) *DescribeUsageOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetStartDate(v int64) *DescribeUsageOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetTypes(v string) *DescribeUsageOverallDataRequest {
	s.Types = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) Validate() error {
	return dara.Validate(s)
}
