// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *DescribeSlowLogStatisticRequest
	GetAsc() *bool
	SetEndTime(v int64) *DescribeSlowLogStatisticRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeSlowLogStatisticRequestFilters) *DescribeSlowLogStatisticRequest
	GetFilters() []*DescribeSlowLogStatisticRequestFilters
	SetInstanceId(v string) *DescribeSlowLogStatisticRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeSlowLogStatisticRequest
	GetNodeId() *string
	SetOrderBy(v string) *DescribeSlowLogStatisticRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeSlowLogStatisticRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSlowLogStatisticRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeSlowLogStatisticRequest
	GetStartTime() *int64
	SetTemplateId(v string) *DescribeSlowLogStatisticRequest
	GetTemplateId() *string
	SetType(v string) *DescribeSlowLogStatisticRequest
	GetType() *string
}

type DescribeSlowLogStatisticRequest struct {
	// example:
	//
	// true
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters []*DescribeSlowLogStatisticRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1568269711000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 04ea3310df40c3fa8a6b4854db49f79a
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticRequest) GetAsc() *bool {
	return s.Asc
}

func (s *DescribeSlowLogStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSlowLogStatisticRequest) GetFilters() []*DescribeSlowLogStatisticRequestFilters {
	return s.Filters
}

func (s *DescribeSlowLogStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlowLogStatisticRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogStatisticRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeSlowLogStatisticRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogStatisticRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSlowLogStatisticRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSlowLogStatisticRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSlowLogStatisticRequest) SetAsc(v bool) *DescribeSlowLogStatisticRequest {
	s.Asc = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetEndTime(v int64) *DescribeSlowLogStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetFilters(v []*DescribeSlowLogStatisticRequestFilters) *DescribeSlowLogStatisticRequest {
	s.Filters = v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetInstanceId(v string) *DescribeSlowLogStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetNodeId(v string) *DescribeSlowLogStatisticRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetOrderBy(v string) *DescribeSlowLogStatisticRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetPageNumber(v int32) *DescribeSlowLogStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetPageSize(v int32) *DescribeSlowLogStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetStartTime(v int64) *DescribeSlowLogStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetTemplateId(v string) *DescribeSlowLogStatisticRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) SetType(v string) *DescribeSlowLogStatisticRequest {
	s.Type = &v
	return s
}

func (s *DescribeSlowLogStatisticRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogStatisticRequestFilters struct {
	// example:
	//
	// KeyWords
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// select
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSlowLogStatisticRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSlowLogStatisticRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSlowLogStatisticRequestFilters) SetKey(v string) *DescribeSlowLogStatisticRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeSlowLogStatisticRequestFilters) SetValue(v string) *DescribeSlowLogStatisticRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeSlowLogStatisticRequestFilters) Validate() error {
	return dara.Validate(s)
}
