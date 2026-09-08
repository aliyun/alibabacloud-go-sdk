// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdList(v string) *ListHistoricalAgentReportRequest
	GetAgentIdList() *string
	SetInstanceId(v string) *ListHistoricalAgentReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListHistoricalAgentReportRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListHistoricalAgentReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalAgentReportRequest
	GetPageSize() *int32
	SetSkillGroupId(v string) *ListHistoricalAgentReportRequest
	GetSkillGroupId() *string
	SetStartTime(v int64) *ListHistoricalAgentReportRequest
	GetStartTime() *int64
	SetStopTime(v int64) *ListHistoricalAgentReportRequest
	GetStopTime() *int64
}

type ListHistoricalAgentReportRequest struct {
	// The list of agent IDs. The list can contain 0 to 100 agent IDs.
	//
	// example:
	//
	// ["agent1@ccc-test", "agent2@ccc-test"]
	AgentIdList *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The media type. Default value: Audio. Other valid values: Chat and Video.
	//
	// example:
	//
	// VIDEO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The page number. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The skill group ID.
	//
	// example:
	//
	// wwtest@test_yunhu
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The start time of the historical data to retrieve. Default value: 00:00 of the current day. The earliest allowed value is 180 days before the current time. The statistical time precision is hour-level, rounded down to the nearest hour. This is a closed interval. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The end time of the historical data to retrieve. Default value: the current time. The statistical time precision is hour-level, rounded up to the nearest hour. This is an open interval. For example, if the start time is 11:12:20 and the end time is 11:45:50, the aligned time range is [11:00:00, 12:00:00), which means greater than or equal to 11:00 and less than 12:00. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1532707199000
	StopTime *int64 `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s ListHistoricalAgentReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportRequest) GetAgentIdList() *string {
	return s.AgentIdList
}

func (s *ListHistoricalAgentReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalAgentReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListHistoricalAgentReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalAgentReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalAgentReportRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListHistoricalAgentReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalAgentReportRequest) GetStopTime() *int64 {
	return s.StopTime
}

func (s *ListHistoricalAgentReportRequest) SetAgentIdList(v string) *ListHistoricalAgentReportRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetInstanceId(v string) *ListHistoricalAgentReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetMediaType(v string) *ListHistoricalAgentReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetPageNumber(v int32) *ListHistoricalAgentReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetPageSize(v int32) *ListHistoricalAgentReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetSkillGroupId(v string) *ListHistoricalAgentReportRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetStartTime(v int64) *ListHistoricalAgentReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetStopTime(v int64) *ListHistoricalAgentReportRequest {
	s.StopTime = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) Validate() error {
	return dara.Validate(s)
}
