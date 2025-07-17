// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInsightsEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListInsightsEventsRequest
	GetEndTime() *string
	SetInsightsTypes(v string) *ListInsightsEventsRequest
	GetInsightsTypes() *string
	SetPid(v string) *ListInsightsEventsRequest
	GetPid() *string
	SetRegionId(v string) *ListInsightsEventsRequest
	GetRegionId() *string
	SetStartTime(v string) *ListInsightsEventsRequest
	GetStartTime() *string
}

type ListInsightsEventsRequest struct {
	// The end of the time range to query. The value is a timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1480607940000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The types of the events that you want to query. Separate multiple event types with commas (,). If you do not specify this parameter, all events are queried.
	//
	// 	- errorIncrease: API error-rate spike events. Examples: HTTP API error-rate spike events and Dubbo API error-rate spike events.
	//
	// 	- topErrorIncrease: the top five API error-rate spike events with the highest traffic.
	//
	// 	- topRtIncrease: API response-time spike events. Examples: HTTP API response-time spike events and Dubbo API response-time spike events.
	//
	// 	- rtIncrease: the top five API response-time spike events with the highest traffic.
	//
	// example:
	//
	// errorIncrease,topErrorIncrease,topExceptionIncrease,topRtIncrease,rtIncrease
	InsightsTypes *string `json:"InsightsTypes,omitempty" xml:"InsightsTypes,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// aokcdqn3ly@a195c6d6421****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. The value is a timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595174400000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListInsightsEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInsightsEventsRequest) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListInsightsEventsRequest) GetInsightsTypes() *string {
	return s.InsightsTypes
}

func (s *ListInsightsEventsRequest) GetPid() *string {
	return s.Pid
}

func (s *ListInsightsEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInsightsEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInsightsEventsRequest) SetEndTime(v string) *ListInsightsEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListInsightsEventsRequest) SetInsightsTypes(v string) *ListInsightsEventsRequest {
	s.InsightsTypes = &v
	return s
}

func (s *ListInsightsEventsRequest) SetPid(v string) *ListInsightsEventsRequest {
	s.Pid = &v
	return s
}

func (s *ListInsightsEventsRequest) SetRegionId(v string) *ListInsightsEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInsightsEventsRequest) SetStartTime(v string) *ListInsightsEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListInsightsEventsRequest) Validate() error {
	return dara.Validate(s)
}
