// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListIntervalAgentReportRequest
	GetAgentId() *string
	SetEndTime(v int64) *ListIntervalAgentReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListIntervalAgentReportRequest
	GetInstanceId() *string
	SetInterval(v string) *ListIntervalAgentReportRequest
	GetInterval() *string
	SetMediaType(v string) *ListIntervalAgentReportRequest
	GetMediaType() *string
	SetStartTime(v int64) *ListIntervalAgentReportRequest
	GetStartTime() *int64
}

type ListIntervalAgentReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1532707199000
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
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalAgentReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListIntervalAgentReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntervalAgentReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntervalAgentReportRequest) GetInterval() *string {
	return s.Interval
}

func (s *ListIntervalAgentReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListIntervalAgentReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIntervalAgentReportRequest) SetAgentId(v string) *ListIntervalAgentReportRequest {
	s.AgentId = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetEndTime(v int64) *ListIntervalAgentReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetInstanceId(v string) *ListIntervalAgentReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetInterval(v string) *ListIntervalAgentReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetMediaType(v string) *ListIntervalAgentReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetStartTime(v int64) *ListIntervalAgentReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalAgentReportRequest) Validate() error {
	return dara.Validate(s)
}
