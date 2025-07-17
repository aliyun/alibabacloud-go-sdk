// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *GetSyntheticTaskListRequest
	GetDirection() *string
	SetOrder(v string) *GetSyntheticTaskListRequest
	GetOrder() *string
	SetPageNum(v int64) *GetSyntheticTaskListRequest
	GetPageNum() *int64
	SetPageSize(v int64) *GetSyntheticTaskListRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetSyntheticTaskListRequest
	GetRegionId() *string
	SetTaskName(v string) *GetSyntheticTaskListRequest
	GetTaskName() *string
	SetTaskStatus(v string) *GetSyntheticTaskListRequest
	GetTaskStatus() *string
	SetTaskType(v string) *GetSyntheticTaskListRequest
	GetTaskType() *string
	SetUrl(v string) *GetSyntheticTaskListRequest
	GetUrl() *string
}

type GetSyntheticTaskListRequest struct {
	// The order by which the queried tasks are sorted. Valid values:
	//
	// 	- **asc**: ascending
	//
	// 	- **desc**: descending
	//
	// example:
	//
	// asc
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The condition by which the queried tasks are sorted.
	//
	// example:
	//
	// CreateTime
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task name.
	//
	// example:
	//
	// net-test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: The task is stopped.
	//
	// 	- **1**: The task is started.
	//
	// 	- **9**: The task is ended.
	//
	// example:
	//
	// 1
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. Valid values:
	//
	// 1.  3: web page performance - IE
	//
	// 2.  34: web page performance - Chrome
	//
	// 3.  0: network quality
	//
	// 4.  40: file download
	//
	// 5.  7: API performance
	//
	// example:
	//
	// 0
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The URL for synthetic monitoring.
	//
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSyntheticTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskListRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetSyntheticTaskListRequest) GetOrder() *string {
	return s.Order
}

func (s *GetSyntheticTaskListRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetSyntheticTaskListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSyntheticTaskListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSyntheticTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetSyntheticTaskListRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetSyntheticTaskListRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetSyntheticTaskListRequest) GetUrl() *string {
	return s.Url
}

func (s *GetSyntheticTaskListRequest) SetDirection(v string) *GetSyntheticTaskListRequest {
	s.Direction = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetOrder(v string) *GetSyntheticTaskListRequest {
	s.Order = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetPageNum(v int64) *GetSyntheticTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetPageSize(v int64) *GetSyntheticTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetRegionId(v string) *GetSyntheticTaskListRequest {
	s.RegionId = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetTaskName(v string) *GetSyntheticTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetTaskStatus(v string) *GetSyntheticTaskListRequest {
	s.TaskStatus = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetTaskType(v string) *GetSyntheticTaskListRequest {
	s.TaskType = &v
	return s
}

func (s *GetSyntheticTaskListRequest) SetUrl(v string) *GetSyntheticTaskListRequest {
	s.Url = &v
	return s
}

func (s *GetSyntheticTaskListRequest) Validate() error {
	return dara.Validate(s)
}
