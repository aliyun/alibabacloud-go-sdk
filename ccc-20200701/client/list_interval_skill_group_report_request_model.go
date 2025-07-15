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
	// example:
	//
	// 1604725528000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Hourly
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skg-default@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
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
