// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyQueueEventLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListLegacyQueueEventLogsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListLegacyQueueEventLogsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListLegacyQueueEventLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLegacyQueueEventLogsRequest
	GetPageSize() *int32
	SetSkillGroupId(v string) *ListLegacyQueueEventLogsRequest
	GetSkillGroupId() *string
	SetStartTime(v int64) *ListLegacyQueueEventLogsRequest
	GetStartTime() *int64
}

type ListLegacyQueueEventLogsRequest struct {
	// End time, in UNIX timestamp format, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1658026180018
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number for paging, ranging from 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filter by skill group ID. This parameter is optional and defaults to empty. An empty value means no filtering is applied.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Start Time, in UNIX timestamp format, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1657939540015
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLegacyQueueEventLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyQueueEventLogsRequest) GoString() string {
	return s.String()
}

func (s *ListLegacyQueueEventLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLegacyQueueEventLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLegacyQueueEventLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyQueueEventLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyQueueEventLogsRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListLegacyQueueEventLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListLegacyQueueEventLogsRequest) SetEndTime(v int64) *ListLegacyQueueEventLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLegacyQueueEventLogsRequest) SetInstanceId(v string) *ListLegacyQueueEventLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLegacyQueueEventLogsRequest) SetPageNumber(v int32) *ListLegacyQueueEventLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyQueueEventLogsRequest) SetPageSize(v int32) *ListLegacyQueueEventLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLegacyQueueEventLogsRequest) SetSkillGroupId(v string) *ListLegacyQueueEventLogsRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListLegacyQueueEventLogsRequest) SetStartTime(v int64) *ListLegacyQueueEventLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLegacyQueueEventLogsRequest) Validate() error {
	return dara.Validate(s)
}
