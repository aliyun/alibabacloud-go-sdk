// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMvRecommendTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionInner(v string) *DescribeMvRecommendTasksRequest
	GetActionInner() *string
	SetDBClusterId(v string) *DescribeMvRecommendTasksRequest
	GetDBClusterId() *string
	SetFrom(v string) *DescribeMvRecommendTasksRequest
	GetFrom() *string
	SetPageNumber(v int32) *DescribeMvRecommendTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMvRecommendTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMvRecommendTasksRequest
	GetRegionId() *string
	SetTaskName(v string) *DescribeMvRecommendTasksRequest
	GetTaskName() *string
}

type DescribeMvRecommendTasksRequest struct {
	// Fixed system value (non-modifiable).
	ActionInner *string `json:"ActionInner,omitempty" xml:"ActionInner,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Fixed system value (non-modifiable).
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
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
	// The name of the recommendation task.
	//
	// example:
	//
	// my_task_1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeMvRecommendTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendTasksRequest) GetActionInner() *string {
	return s.ActionInner
}

func (s *DescribeMvRecommendTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMvRecommendTasksRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeMvRecommendTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMvRecommendTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMvRecommendTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMvRecommendTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeMvRecommendTasksRequest) SetActionInner(v string) *DescribeMvRecommendTasksRequest {
	s.ActionInner = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) SetDBClusterId(v string) *DescribeMvRecommendTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) SetFrom(v string) *DescribeMvRecommendTasksRequest {
	s.From = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) SetPageNumber(v int32) *DescribeMvRecommendTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) SetPageSize(v int32) *DescribeMvRecommendTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) SetRegionId(v string) *DescribeMvRecommendTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) SetTaskName(v string) *DescribeMvRecommendTasksRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeMvRecommendTasksRequest) Validate() error {
	return dara.Validate(s)
}
