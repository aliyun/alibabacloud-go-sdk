// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSqlLogRecordsRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeSqlLogRecordsRequestFilters) *DescribeSqlLogRecordsRequest
	GetFilters() []*DescribeSqlLogRecordsRequestFilters
	SetInstanceId(v string) *DescribeSqlLogRecordsRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeSqlLogRecordsRequest
	GetNodeId() *string
	SetPageNo(v int32) *DescribeSqlLogRecordsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeSqlLogRecordsRequest
	GetPageSize() *int32
	SetRole(v string) *DescribeSqlLogRecordsRequest
	GetRole() *string
	SetStartTime(v int64) *DescribeSqlLogRecordsRequest
	GetStartTime() *int64
}

type DescribeSqlLogRecordsRequest struct {
	// The end of the time range to query. This is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A list of filter conditions.
	Filters []*DescribeSqlLogRecordsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// - This parameter is applicable only to ApsaraDB RDS for MySQL and PolarDB for MySQL cluster instances. If this parameter is omitted, the log details of the primary node are returned by default.
	//
	// - For PolarDB-X 2.0, set this parameter to **polarx_cn*	- (compute node) or **polarx_dn*	- (data node).
	//
	// example:
	//
	// pi-uf6k5f6g3912i****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role of the node in the PolarDB-X 2.0 database instance.
	//
	// - **polarx_cn**: compute node.
	//
	// - **polarx_dn**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The beginning of the time range to query. This is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSqlLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSqlLogRecordsRequest) GetFilters() []*DescribeSqlLogRecordsRequestFilters {
	return s.Filters
}

func (s *DescribeSqlLogRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSqlLogRecordsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeSqlLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlLogRecordsRequest) GetRole() *string {
	return s.Role
}

func (s *DescribeSqlLogRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSqlLogRecordsRequest) SetEndTime(v int64) *DescribeSqlLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetFilters(v []*DescribeSqlLogRecordsRequestFilters) *DescribeSqlLogRecordsRequest {
	s.Filters = v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetInstanceId(v string) *DescribeSqlLogRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetNodeId(v string) *DescribeSqlLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetPageNo(v int32) *DescribeSqlLogRecordsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetPageSize(v int32) *DescribeSqlLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetRole(v string) *DescribeSqlLogRecordsRequest {
	s.Role = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) SetStartTime(v int64) *DescribeSqlLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlLogRecordsRequest) Validate() error {
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

type DescribeSqlLogRecordsRequestFilters struct {
	// The filter key.
	//
	// > For details on supported filter keys and their values, see **Additional information about request parameters**.
	//
	// example:
	//
	// keyWords
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value.
	//
	// example:
	//
	// select
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSqlLogRecordsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSqlLogRecordsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSqlLogRecordsRequestFilters) SetKey(v string) *DescribeSqlLogRecordsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeSqlLogRecordsRequestFilters) SetValue(v string) *DescribeSqlLogRecordsRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeSqlLogRecordsRequestFilters) Validate() error {
	return dara.Validate(s)
}
