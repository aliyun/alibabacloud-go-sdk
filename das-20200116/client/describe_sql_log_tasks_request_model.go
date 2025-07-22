// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSqlLogTasksRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeSqlLogTasksRequestFilters) *DescribeSqlLogTasksRequest
	GetFilters() []*DescribeSqlLogTasksRequestFilters
	SetInstanceId(v string) *DescribeSqlLogTasksRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeSqlLogTasksRequest
	GetNodeId() *string
	SetPageNo(v int32) *DescribeSqlLogTasksRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeSqlLogTasksRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeSqlLogTasksRequest
	GetStartTime() *int64
}

type DescribeSqlLogTasksRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter parameters.
	Filters []*DescribeSqlLogTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The ID of the database instance.
	//
	// example:
	//
	// r-bp1nti25tc7bq5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  This parameter is available only for instances that are deployed in the cluster architecture. You can specify this parameter to query the tasks of a specific node. If this parameter is not specified, the tasks of the primary node are returned by default.
	//
	// example:
	//
	// pi-bp1o58x3ib7e6z496
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSqlLogTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSqlLogTasksRequest) GetFilters() []*DescribeSqlLogTasksRequestFilters {
	return s.Filters
}

func (s *DescribeSqlLogTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlLogTasksRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSqlLogTasksRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeSqlLogTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlLogTasksRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSqlLogTasksRequest) SetEndTime(v int64) *DescribeSqlLogTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSqlLogTasksRequest) SetFilters(v []*DescribeSqlLogTasksRequestFilters) *DescribeSqlLogTasksRequest {
	s.Filters = v
	return s
}

func (s *DescribeSqlLogTasksRequest) SetInstanceId(v string) *DescribeSqlLogTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlLogTasksRequest) SetNodeId(v string) *DescribeSqlLogTasksRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSqlLogTasksRequest) SetPageNo(v int32) *DescribeSqlLogTasksRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeSqlLogTasksRequest) SetPageSize(v int32) *DescribeSqlLogTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlLogTasksRequest) SetStartTime(v int64) *DescribeSqlLogTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlLogTasksRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTasksRequestFilters struct {
	// The name of the filter parameter.
	//
	// >  For more information about the filter parameters, see the **Valid values of Key*	- section of this topic.
	//
	// example:
	//
	// delimiter
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter parameter.
	//
	// example:
	//
	// ,
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSqlLogTasksRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSqlLogTasksRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSqlLogTasksRequestFilters) SetKey(v string) *DescribeSqlLogTasksRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeSqlLogTasksRequestFilters) SetValue(v string) *DescribeSqlLogTasksRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeSqlLogTasksRequestFilters) Validate() error {
	return dara.Validate(s)
}
