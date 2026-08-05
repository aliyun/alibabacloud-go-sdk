// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskErrorLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int32) *ListOfflineTaskErrorLogsRequest
	GetEndTime() *int32
	SetPageNum(v int32) *ListOfflineTaskErrorLogsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListOfflineTaskErrorLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListOfflineTaskErrorLogsRequest
	GetRegionId() *string
	SetStartTime(v int32) *ListOfflineTaskErrorLogsRequest
	GetStartTime() *int32
}

type ListOfflineTaskErrorLogsRequest struct {
	// The end timestamp in seconds. If not specified, the current time is used by default.
	//
	// example:
	//
	// 1760530000
	EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The start timestamp in seconds. If not specified, the time one hour before the current time is used by default.
	//
	// example:
	//
	// 1762946698
	StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListOfflineTaskErrorLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskErrorLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskErrorLogsRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ListOfflineTaskErrorLogsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListOfflineTaskErrorLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOfflineTaskErrorLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfflineTaskErrorLogsRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ListOfflineTaskErrorLogsRequest) SetEndTime(v int32) *ListOfflineTaskErrorLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListOfflineTaskErrorLogsRequest) SetPageNum(v int32) *ListOfflineTaskErrorLogsRequest {
	s.PageNum = &v
	return s
}

func (s *ListOfflineTaskErrorLogsRequest) SetPageSize(v int32) *ListOfflineTaskErrorLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOfflineTaskErrorLogsRequest) SetRegionId(v string) *ListOfflineTaskErrorLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfflineTaskErrorLogsRequest) SetStartTime(v int32) *ListOfflineTaskErrorLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListOfflineTaskErrorLogsRequest) Validate() error {
	return dara.Validate(s)
}
