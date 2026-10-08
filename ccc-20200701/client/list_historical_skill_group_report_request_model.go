// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalSkillGroupReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListHistoricalSkillGroupReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListHistoricalSkillGroupReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListHistoricalSkillGroupReportRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListHistoricalSkillGroupReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalSkillGroupReportRequest
	GetPageSize() *int32
	SetSkillGroupIdList(v string) *ListHistoricalSkillGroupReportRequest
	GetSkillGroupIdList() *string
	SetStartTime(v int64) *ListHistoricalSkillGroupReportRequest
	GetStartTime() *int64
	SetSummarizeByInstanceId(v bool) *ListHistoricalSkillGroupReportRequest
	GetSummarizeByInstanceId() *bool
}

type ListHistoricalSkillGroupReportRequest struct {
	// The end time of the historical data to retrieve. Specify a UNIX timestamp in milliseconds. This parameter is optional. Default value: the current time. The statistical time precision is in hours. The end time is rounded up to the nearest hour, and the interval is open. For example, if the start time is 11:12:20 and the end time is 11:45:50, the aligned time range is [11:00:00, 12:00:00), which means greater than or equal to 11:00:00 and less than 12:00:00.
	//
	// example:
	//
	// 1532707199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The media type. Default value: Audio. Valid values: Audio, Chat, and Video.
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
	// The number of entries per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of skill group IDs to query. The value is a character string in the JSON array format, where each array element is a skill group ID. This parameter is optional. Default value: empty. An empty value indicates that all skill groups in the current paging are queried.
	//
	// example:
	//
	// ["skillgroup1@ccc-test", "skillgroup2@ccc-test2"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	// The start time of the historical data to retrieve. Specify a UNIX timestamp in milliseconds. This parameter is optional. Default value: 00:00:00 on the current day. The earliest allowed time is 180 days before the current time. The statistical time precision is in hours. The start time is rounded down to the nearest hour, and the interval is closed.
	//
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether to aggregate data by instance ID.
	SummarizeByInstanceId *bool `json:"SummarizeByInstanceId,omitempty" xml:"SummarizeByInstanceId,omitempty"`
}

func (s ListHistoricalSkillGroupReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHistoricalSkillGroupReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalSkillGroupReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListHistoricalSkillGroupReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalSkillGroupReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalSkillGroupReportRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *ListHistoricalSkillGroupReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalSkillGroupReportRequest) GetSummarizeByInstanceId() *bool {
	return s.SummarizeByInstanceId
}

func (s *ListHistoricalSkillGroupReportRequest) SetEndTime(v int64) *ListHistoricalSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetInstanceId(v string) *ListHistoricalSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetMediaType(v string) *ListHistoricalSkillGroupReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetPageNumber(v int32) *ListHistoricalSkillGroupReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetPageSize(v int32) *ListHistoricalSkillGroupReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetSkillGroupIdList(v string) *ListHistoricalSkillGroupReportRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetStartTime(v int64) *ListHistoricalSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetSummarizeByInstanceId(v bool) *ListHistoricalSkillGroupReportRequest {
	s.SummarizeByInstanceId = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) Validate() error {
	return dara.Validate(s)
}
