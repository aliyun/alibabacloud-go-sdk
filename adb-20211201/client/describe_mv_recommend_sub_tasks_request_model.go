// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMvRecommendSubTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionInner(v string) *DescribeMvRecommendSubTasksRequest
	GetActionInner() *string
	SetDBClusterId(v string) *DescribeMvRecommendSubTasksRequest
	GetDBClusterId() *string
	SetFrom(v string) *DescribeMvRecommendSubTasksRequest
	GetFrom() *string
	SetOrderBy(v string) *DescribeMvRecommendSubTasksRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeMvRecommendSubTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMvRecommendSubTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMvRecommendSubTasksRequest
	GetRegionId() *string
	SetSubtaskId(v int64) *DescribeMvRecommendSubTasksRequest
	GetSubtaskId() *int64
	SetTaskName(v string) *DescribeMvRecommendSubTasksRequest
	GetTaskName() *string
}

type DescribeMvRecommendSubTasksRequest struct {
	// Fixed system value (non-modifiable).
	ActionInner *string `json:"ActionInner,omitempty" xml:"ActionInner,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Fixed system value (non-modifiable).
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The sorting field. Valid values for Type:
	//
	// 	- Asc.
	//
	// 	- Desc.
	//
	// Valid values for Field:
	//
	// 	- StartTime;
	//
	// 	- EndTime;
	//
	// example:
	//
	// {\\"Type\\": \\"ASC\\", \\"Field\\": \\"StartTime\\"}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subtask ID.
	//
	// example:
	//
	// 123
	SubtaskId *int64 `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
	// The name of the recommendation task.
	//
	// example:
	//
	// my_task_1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeMvRecommendSubTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendSubTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendSubTasksRequest) GetActionInner() *string {
	return s.ActionInner
}

func (s *DescribeMvRecommendSubTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMvRecommendSubTasksRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeMvRecommendSubTasksRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeMvRecommendSubTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMvRecommendSubTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMvRecommendSubTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMvRecommendSubTasksRequest) GetSubtaskId() *int64 {
	return s.SubtaskId
}

func (s *DescribeMvRecommendSubTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeMvRecommendSubTasksRequest) SetActionInner(v string) *DescribeMvRecommendSubTasksRequest {
	s.ActionInner = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetDBClusterId(v string) *DescribeMvRecommendSubTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetFrom(v string) *DescribeMvRecommendSubTasksRequest {
	s.From = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetOrderBy(v string) *DescribeMvRecommendSubTasksRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetPageNumber(v int32) *DescribeMvRecommendSubTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetPageSize(v int32) *DescribeMvRecommendSubTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetRegionId(v string) *DescribeMvRecommendSubTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetSubtaskId(v int64) *DescribeMvRecommendSubTasksRequest {
	s.SubtaskId = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) SetTaskName(v string) *DescribeMvRecommendSubTasksRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeMvRecommendSubTasksRequest) Validate() error {
	return dara.Validate(s)
}
