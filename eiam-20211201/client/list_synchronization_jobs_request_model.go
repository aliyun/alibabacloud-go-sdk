// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSynchronizationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListSynchronizationJobsRequest
	GetDirection() *string
	SetEndTime(v int64) *ListSynchronizationJobsRequest
	GetEndTime() *int64
	SetFilters(v []*ListSynchronizationJobsRequestFilters) *ListSynchronizationJobsRequest
	GetFilters() []*ListSynchronizationJobsRequestFilters
	SetInstanceId(v string) *ListSynchronizationJobsRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListSynchronizationJobsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListSynchronizationJobsRequest
	GetNextToken() *string
	SetPageNumber(v int64) *ListSynchronizationJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSynchronizationJobsRequest
	GetPageSize() *int64
	SetStartTime(v int64) *ListSynchronizationJobsRequest
	GetStartTime() *int64
	SetStatus(v string) *ListSynchronizationJobsRequest
	GetStatus() *string
	SetTargetIds(v []*string) *ListSynchronizationJobsRequest
	GetTargetIds() []*string
	SetTargetType(v string) *ListSynchronizationJobsRequest
	GetTargetType() *string
}

type ListSynchronizationJobsRequest struct {
	// 同步方向[ingress,egress]
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 同步结束时间
	//
	// example:
	//
	// 1649830226000
	EndTime *int64                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters []*ListSynchronizationJobsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 当前查询的列表页码，默认为1。
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 当前查询的列表页码，默认为20。
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 同步开始时间
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 同步状态[pending,running,suspending,failed,partial_success,success]
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 同步目标ID
	//
	// example:
	//
	// target_001
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
	// 同步目标类型[identity_provider,organizational_unit,application,user]
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListSynchronizationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListSynchronizationJobsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSynchronizationJobsRequest) GetFilters() []*ListSynchronizationJobsRequestFilters {
	return s.Filters
}

func (s *ListSynchronizationJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSynchronizationJobsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListSynchronizationJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSynchronizationJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSynchronizationJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSynchronizationJobsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSynchronizationJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSynchronizationJobsRequest) GetTargetIds() []*string {
	return s.TargetIds
}

func (s *ListSynchronizationJobsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListSynchronizationJobsRequest) SetDirection(v string) *ListSynchronizationJobsRequest {
	s.Direction = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetEndTime(v int64) *ListSynchronizationJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetFilters(v []*ListSynchronizationJobsRequestFilters) *ListSynchronizationJobsRequest {
	s.Filters = v
	return s
}

func (s *ListSynchronizationJobsRequest) SetInstanceId(v string) *ListSynchronizationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetMaxResults(v int64) *ListSynchronizationJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetNextToken(v string) *ListSynchronizationJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetPageNumber(v int64) *ListSynchronizationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetPageSize(v int64) *ListSynchronizationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetStartTime(v int64) *ListSynchronizationJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetStatus(v string) *ListSynchronizationJobsRequest {
	s.Status = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetTargetIds(v []*string) *ListSynchronizationJobsRequest {
	s.TargetIds = v
	return s
}

func (s *ListSynchronizationJobsRequest) SetTargetType(v string) *ListSynchronizationJobsRequest {
	s.TargetType = &v
	return s
}

func (s *ListSynchronizationJobsRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSynchronizationJobsRequestFilters struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListSynchronizationJobsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListSynchronizationJobsRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *ListSynchronizationJobsRequestFilters) SetKey(v string) *ListSynchronizationJobsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListSynchronizationJobsRequestFilters) SetValues(v []*string) *ListSynchronizationJobsRequestFilters {
	s.Values = v
	return s
}

func (s *ListSynchronizationJobsRequestFilters) Validate() error {
	return dara.Validate(s)
}
