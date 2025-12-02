// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMVRecommendResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionInner(v string) *DescribeMVRecommendResultsRequest
	GetActionInner() *string
	SetDBClusterId(v string) *DescribeMVRecommendResultsRequest
	GetDBClusterId() *string
	SetFrom(v string) *DescribeMVRecommendResultsRequest
	GetFrom() *string
	SetOrderBy(v string) *DescribeMVRecommendResultsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeMVRecommendResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMVRecommendResultsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMVRecommendResultsRequest
	GetRegionId() *string
	SetSubQueryId(v int64) *DescribeMVRecommendResultsRequest
	GetSubQueryId() *int64
	SetSubtaskId(v int64) *DescribeMVRecommendResultsRequest
	GetSubtaskId() *int64
	SetTaskName(v string) *DescribeMVRecommendResultsRequest
	GetTaskName() *string
}

type DescribeMVRecommendResultsRequest struct {
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
	// The sorting field.
	//
	// example:
	//
	// {"Type": "Desc","Field": "AcceleratedQueriesCount"}
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
	// example:
	//
	// 123
	SubQueryId *int64 `json:"SubQueryId,omitempty" xml:"SubQueryId,omitempty"`
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
	// task_n1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeMVRecommendResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMVRecommendResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMVRecommendResultsRequest) GetActionInner() *string {
	return s.ActionInner
}

func (s *DescribeMVRecommendResultsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMVRecommendResultsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeMVRecommendResultsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeMVRecommendResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMVRecommendResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMVRecommendResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMVRecommendResultsRequest) GetSubQueryId() *int64 {
	return s.SubQueryId
}

func (s *DescribeMVRecommendResultsRequest) GetSubtaskId() *int64 {
	return s.SubtaskId
}

func (s *DescribeMVRecommendResultsRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeMVRecommendResultsRequest) SetActionInner(v string) *DescribeMVRecommendResultsRequest {
	s.ActionInner = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetDBClusterId(v string) *DescribeMVRecommendResultsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetFrom(v string) *DescribeMVRecommendResultsRequest {
	s.From = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetOrderBy(v string) *DescribeMVRecommendResultsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetPageNumber(v int32) *DescribeMVRecommendResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetPageSize(v int32) *DescribeMVRecommendResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetRegionId(v string) *DescribeMVRecommendResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetSubQueryId(v int64) *DescribeMVRecommendResultsRequest {
	s.SubQueryId = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetSubtaskId(v int64) *DescribeMVRecommendResultsRequest {
	s.SubtaskId = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) SetTaskName(v string) *DescribeMVRecommendResultsRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeMVRecommendResultsRequest) Validate() error {
	return dara.Validate(s)
}
