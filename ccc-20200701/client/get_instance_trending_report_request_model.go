// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrendingReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetInstanceTrendingReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetInstanceTrendingReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *GetInstanceTrendingReportRequest
	GetMediaType() *string
	SetStartTime(v int64) *GetInstanceTrendingReportRequest
	GetStartTime() *int64
}

type GetInstanceTrendingReportRequest struct {
	// End UNIX timestamp. The default value is the current time. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1604725528000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Media type. The default value is Audio. Other valid values include Chat and Video.
	//
	// example:
	//
	// Audio
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Start UNIX timestamp. The default value is the start time of the current day. The earliest allowed time is 180 days before the current time. The interval between the start time and end time cannot exceed 7 days. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1604639129000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInstanceTrendingReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceTrendingReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceTrendingReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *GetInstanceTrendingReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceTrendingReportRequest) SetEndTime(v int64) *GetInstanceTrendingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetInstanceId(v string) *GetInstanceTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetMediaType(v string) *GetInstanceTrendingReportRequest {
	s.MediaType = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetStartTime(v int64) *GetInstanceTrendingReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) Validate() error {
	return dara.Validate(s)
}
