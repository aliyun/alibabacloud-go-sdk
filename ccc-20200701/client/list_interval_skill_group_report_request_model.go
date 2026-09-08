// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalSkillGroupReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListIntervalSkillGroupReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListIntervalSkillGroupReportRequest
	GetInstanceId() *string
	SetInterval(v string) *ListIntervalSkillGroupReportRequest
	GetInterval() *string
	SetMediaType(v string) *ListIntervalSkillGroupReportRequest
	GetMediaType() *string
	SetSkillGroupId(v string) *ListIntervalSkillGroupReportRequest
	GetSkillGroupId() *string
	SetStartTime(v int64) *ListIntervalSkillGroupReportRequest
	GetStartTime() *int64
}

type ListIntervalSkillGroupReportRequest struct {
	// End Time, formatted as a UNIX timestamp in milliseconds. This parameter is optional. The default value is the current time. If Interval is Daily, the maximum interval between StartTime and EndTime is 180 days. If Interval is Hourly, the maximum interval is 10 days. The time precision for statistics is hourly, snapped backward to the nearest hour, using an open interval. For example, if the Start Time is 11:12:20 and the End Time is 11:45:50, the aligned input parameter Time Range becomes [11:00:00, 12:00:00), meaning greater than or equal to 11:00:00 and less than 12:00:00.
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
	// Segment statistics type. The default is Daily (aggregated by day).
	//
	// example:
	//
	// Hourly
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Media type. The default is Audio. Other valid values include Chat and Video.
	//
	// example:
	//
	// VIDEO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// skg-default@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Start Time, formatted as a UNIX timestamp in milliseconds. This parameter is optional. The default value is 00:00 of the current day. Statistics are aggregated by hour, rounded down to the nearest hour, and the interval is closed.
	//
	// example:
	//
	// 1604639129000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalSkillGroupReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntervalSkillGroupReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntervalSkillGroupReportRequest) GetInterval() *string {
	return s.Interval
}

func (s *ListIntervalSkillGroupReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListIntervalSkillGroupReportRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListIntervalSkillGroupReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIntervalSkillGroupReportRequest) SetEndTime(v int64) *ListIntervalSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetInstanceId(v string) *ListIntervalSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetInterval(v string) *ListIntervalSkillGroupReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetMediaType(v string) *ListIntervalSkillGroupReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetSkillGroupId(v string) *ListIntervalSkillGroupReportRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetStartTime(v int64) *ListIntervalSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) Validate() error {
	return dara.Validate(s)
}
