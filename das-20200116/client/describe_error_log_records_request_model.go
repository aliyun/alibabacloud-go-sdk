// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeErrorLogRecordsRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeErrorLogRecordsRequestFilters) *DescribeErrorLogRecordsRequest
	GetFilters() []*DescribeErrorLogRecordsRequestFilters
	SetInstanceId(v string) *DescribeErrorLogRecordsRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeErrorLogRecordsRequest
	GetNodeId() *string
	SetPageNumber(v int32) *DescribeErrorLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeErrorLogRecordsRequest
	GetPageSize() *int32
	SetRole(v string) *DescribeErrorLogRecordsRequest
	GetRole() *string
	SetStartTime(v int64) *DescribeErrorLogRecordsRequest
	GetStartTime() *int64
}

type DescribeErrorLogRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1732069466000
	EndTime *int64                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters []*DescribeErrorLogRecordsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// db
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1731983066000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeErrorLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeErrorLogRecordsRequest) GetFilters() []*DescribeErrorLogRecordsRequestFilters {
	return s.Filters
}

func (s *DescribeErrorLogRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeErrorLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeErrorLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeErrorLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeErrorLogRecordsRequest) GetRole() *string {
	return s.Role
}

func (s *DescribeErrorLogRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeErrorLogRecordsRequest) SetEndTime(v int64) *DescribeErrorLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetFilters(v []*DescribeErrorLogRecordsRequestFilters) *DescribeErrorLogRecordsRequest {
	s.Filters = v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetInstanceId(v string) *DescribeErrorLogRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetNodeId(v string) *DescribeErrorLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetPageNumber(v int32) *DescribeErrorLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetPageSize(v int32) *DescribeErrorLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetRole(v string) *DescribeErrorLogRecordsRequest {
	s.Role = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetStartTime(v int64) *DescribeErrorLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) Validate() error {
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

type DescribeErrorLogRecordsRequestFilters struct {
	// example:
	//
	// filters
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// deadlock
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeErrorLogRecordsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeErrorLogRecordsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeErrorLogRecordsRequestFilters) SetKey(v string) *DescribeErrorLogRecordsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeErrorLogRecordsRequestFilters) SetValue(v string) *DescribeErrorLogRecordsRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeErrorLogRecordsRequestFilters) Validate() error {
	return dara.Validate(s)
}
