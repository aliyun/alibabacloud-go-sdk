// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *DescribeSlowLogRecordsRequest
	GetAsc() *bool
	SetEndTime(v int64) *DescribeSlowLogRecordsRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeSlowLogRecordsRequestFilters) *DescribeSlowLogRecordsRequest
	GetFilters() []*DescribeSlowLogRecordsRequestFilters
	SetInstanceId(v string) *DescribeSlowLogRecordsRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeSlowLogRecordsRequest
	GetNodeId() *string
	SetOrderBy(v string) *DescribeSlowLogRecordsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeSlowLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSlowLogRecordsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeSlowLogRecordsRequest
	GetStartTime() *int64
}

type DescribeSlowLogRecordsRequest struct {
	// Specifies whether to sort results in ascending order. Default value: **true**.
	//
	// - **true**: ascending order.
	//
	// - **false**: descending order.
	//
	// example:
	//
	// true
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1634972640000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of filter conditions.
	Filters []*DescribeSlowLogRecordsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-8vbk4xz99su8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pi-d9j9fe7wq7t9i****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// - **MySQL*	-
	//
	//   - QueryTimeSeconds: query duration in seconds.
	//
	//   - LockTimeSeconds: lock time in seconds.
	//
	//   - RowsSent: rows sent.
	//
	//   - RowsExamined: rows examined.
	//
	// - **Redis**
	//
	//   - QueryTime: query duration.
	//
	//   - Timestamp: execution end time.
	//
	// - **MongoDB**
	//
	//   - QueryTime: query duration.
	//
	//   - Timestamp: execution end time.
	//
	//   - KeysExamined: keys examined.
	//
	//   - DocExamined: documents examined.
	//
	//   - ReturnNum: rows returned.
	//
	// <notice>RDS PostgreSQL, PolarDB for PostgreSQL, and SQL Server do not support sorting.</notice>
	//
	// example:
	//
	// QueryTimeSeconds
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) GetAsc() *bool {
	return s.Asc
}

func (s *DescribeSlowLogRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSlowLogRecordsRequest) GetFilters() []*DescribeSlowLogRecordsRequestFilters {
	return s.Filters
}

func (s *DescribeSlowLogRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlowLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeSlowLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsRequest) SetAsc(v bool) *DescribeSlowLogRecordsRequest {
	s.Asc = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v int64) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetFilters(v []*DescribeSlowLogRecordsRequestFilters) *DescribeSlowLogRecordsRequest {
	s.Filters = v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetNodeId(v string) *DescribeSlowLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrderBy(v string) *DescribeSlowLogRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v int64) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) Validate() error {
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

type DescribeSlowLogRecordsRequestFilters struct {
	// The filter parameter.
	//
	// > For more information, refer to the supplementary description.
	//
	// example:
	//
	// None
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter parameter.
	//
	// example:
	//
	// None
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSlowLogRecordsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSlowLogRecordsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSlowLogRecordsRequestFilters) SetKey(v string) *DescribeSlowLogRecordsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeSlowLogRecordsRequestFilters) SetValue(v string) *DescribeSlowLogRecordsRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeSlowLogRecordsRequestFilters) Validate() error {
	return dara.Validate(s)
}
