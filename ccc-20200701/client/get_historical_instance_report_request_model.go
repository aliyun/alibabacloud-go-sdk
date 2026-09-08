// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalInstanceReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetHistoricalInstanceReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetHistoricalInstanceReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *GetHistoricalInstanceReportRequest
	GetMediaType() *string
	SetStartTime(v int64) *GetHistoricalInstanceReportRequest
	GetStartTime() *int64
}

type GetHistoricalInstanceReportRequest struct {
	// The end time of the Historical Data to retrieve, in UNIX timestamp format with millisecond precision. This parameter is optional. The default value is the current time. The time precision for statistics is hourly, snapped to the next full hour, and the interval is open. For example, if the start time is 11:12:20 and the end time is 11:45:50, the snapped input parameter Time Range becomes [11:00:00, 12:00:00), meaning greater than or equal to 11:00:00 and less than 12:00:00.
	//
	// example:
	//
	// 1532707199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Media Type. The default value is Audio. Other valid values include Chat and Video.
	//
	// example:
	//
	// VIDEO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The start time of the Historical Data to retrieve, in UNIX timestamp format with millisecond precision. This parameter is optional. The default value is 00:00:00 of the current day. The earliest allowed time is 180 days before the current time. The time precision for statistics is hourly, snapped to the previous full hour, and the interval is closed.
	//
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHistoricalInstanceReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetHistoricalInstanceReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHistoricalInstanceReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *GetHistoricalInstanceReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetHistoricalInstanceReportRequest) SetEndTime(v int64) *GetHistoricalInstanceReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetInstanceId(v string) *GetHistoricalInstanceReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetMediaType(v string) *GetHistoricalInstanceReportRequest {
	s.MediaType = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetStartTime(v int64) *GetHistoricalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) Validate() error {
	return dara.Validate(s)
}
