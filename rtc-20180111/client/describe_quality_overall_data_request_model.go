// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOverallDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeQualityOverallDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeQualityOverallDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeQualityOverallDataRequest
	GetStartDate() *int64
	SetTypes(v string) *DescribeQualityOverallDataRequest
	GetTypes() *string
}

type DescribeQualityOverallDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The End Time of the query, represented as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The Start Time of the query, represented as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The Metric Types to query, separated by commas (,). Valid values:
	//
	// - **JOIN_CHANNEL_SUC_RATE**: Channel Join Success Rate.
	//
	// - **JOIN_CHANNEL_SUC_FIVE_SEC_RATE**: Five-second Channel Join Success Rate.
	//
	// - **VIDEO_FIRST_PIC_DUR**: Time to First Video Frame.
	//
	// - **AUDIO_STUCK_RATE**: Audio Lag Rate.
	//
	// - **VIDEO_STUCK_RATE**: Video Lag Rate.
	//
	// - **AUDIO_DELAY**: Audio Delay.
	//
	// - **VIDEO_DELAY**: Video Delay.
	//
	// - **AUDIO_HIGH_QUALITY_TRANSMISSION_RATE**: Audio High-quality Transmission Rate.
	//
	// - **VIDEO_HIGH_QUALITY_TRANSMISSION_RATE**: Video High-quality Transmission Rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// JOIN_CHANNEL_SUC_RATE
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeQualityOverallDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeQualityOverallDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeQualityOverallDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeQualityOverallDataRequest) GetTypes() *string {
	return s.Types
}

func (s *DescribeQualityOverallDataRequest) SetAppId(v string) *DescribeQualityOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetEndDate(v int64) *DescribeQualityOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetStartDate(v int64) *DescribeQualityOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetTypes(v string) *DescribeQualityOverallDataRequest {
	s.Types = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) Validate() error {
	return dara.Validate(s)
}
