// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsOptimizationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListApsOptimizationTasksRequest
	GetDBClusterId() *string
	SetEndTime(v string) *ListApsOptimizationTasksRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListApsOptimizationTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApsOptimizationTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListApsOptimizationTasksRequest
	GetRegionId() *string
	SetStartTime(v string) *ListApsOptimizationTasksRequest
	GetStartTime() *string
	SetStrategyType(v string) *ListApsOptimizationTasksRequest
	GetStrategyType() *string
}

type ListApsOptimizationTasksRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-09-30T00:15Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-23T02:18Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the lifecycle management policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// StrategyValue
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s ListApsOptimizationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationTasksRequest) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsOptimizationTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListApsOptimizationTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApsOptimizationTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApsOptimizationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApsOptimizationTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListApsOptimizationTasksRequest) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListApsOptimizationTasksRequest) SetDBClusterId(v string) *ListApsOptimizationTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) SetEndTime(v string) *ListApsOptimizationTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) SetPageNumber(v int32) *ListApsOptimizationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) SetPageSize(v int32) *ListApsOptimizationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) SetRegionId(v string) *ListApsOptimizationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) SetStartTime(v string) *ListApsOptimizationTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) SetStrategyType(v string) *ListApsOptimizationTasksRequest {
	s.StrategyType = &v
	return s
}

func (s *ListApsOptimizationTasksRequest) Validate() error {
	return dara.Validate(s)
}
