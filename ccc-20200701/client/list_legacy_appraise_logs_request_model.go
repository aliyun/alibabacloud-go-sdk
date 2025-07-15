// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAppraiseLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListLegacyAppraiseLogsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListLegacyAppraiseLogsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListLegacyAppraiseLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLegacyAppraiseLogsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListLegacyAppraiseLogsRequest
	GetStartTime() *int64
}

type ListLegacyAppraiseLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1620273600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1604638129000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLegacyAppraiseLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAppraiseLogsRequest) GoString() string {
	return s.String()
}

func (s *ListLegacyAppraiseLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLegacyAppraiseLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLegacyAppraiseLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyAppraiseLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyAppraiseLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListLegacyAppraiseLogsRequest) SetEndTime(v int64) *ListLegacyAppraiseLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLegacyAppraiseLogsRequest) SetInstanceId(v string) *ListLegacyAppraiseLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLegacyAppraiseLogsRequest) SetPageNumber(v int32) *ListLegacyAppraiseLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyAppraiseLogsRequest) SetPageSize(v int32) *ListLegacyAppraiseLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLegacyAppraiseLogsRequest) SetStartTime(v int64) *ListLegacyAppraiseLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLegacyAppraiseLogsRequest) Validate() error {
	return dara.Validate(s)
}
