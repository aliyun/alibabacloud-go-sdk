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
}

type ListHistoricalSkillGroupReportRequest struct {
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
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ["skillgroup1@ccc-test", "skillgroup2@ccc-test2"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *ListHistoricalSkillGroupReportRequest) Validate() error {
	return dara.Validate(s)
}
