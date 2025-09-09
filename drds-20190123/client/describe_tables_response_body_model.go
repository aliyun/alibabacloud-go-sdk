// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeTablesResponseBodyList) *DescribeTablesResponseBody
	GetList() []*DescribeTablesResponseBodyList
	SetPageNumber(v int32) *DescribeTablesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTablesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeTablesResponseBody
	GetTotal() *int32
}

type DescribeTablesResponseBody struct {
	// The list of returned tables.
	List []*DescribeTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tables returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 83AC3D7E-461C-4D87-8ACD-6CC295******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned tables.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) GetList() []*DescribeTablesResponseBodyList {
	return s.List
}

func (s *DescribeTablesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTablesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeTablesResponseBody) SetList(v []*DescribeTablesResponseBodyList) *DescribeTablesResponseBody {
	s.List = v
	return s
}

func (s *DescribeTablesResponseBody) SetPageNumber(v int32) *DescribeTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetSuccess(v bool) *DescribeTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotal(v int32) *DescribeTablesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTablesResponseBodyList struct {
	// Indicates whether full table scanning is allowed.
	//
	// example:
	//
	// false
	AllowFullTableScan *bool `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	// Indicates whether the table is a replicated table.
	//
	// example:
	//
	// false
	Broadcast *bool `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	// The type of the PolarDB-X 1.0 instance. Valid values:
	//
	// 	- 0: The instance is a dedicated instance.
	//
	// 	- 1: The instance is a shard instance.
	//
	// example:
	//
	// 0
	DbInstType *int32 `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates whether the table is locked.
	//
	// example:
	//
	// false
	IsLocked *bool `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	// Indicates whether the table is sharded.
	//
	// example:
	//
	// false
	IsShard *bool `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	// The shard key of the table.
	//
	// example:
	//
	// null
	ShardKey *string `json:"ShardKey,omitempty" xml:"ShardKey,omitempty"`
	// Indicates whether sharding tasks are performed on the table. Valid values:
	//
	// 	- 0: No sharding task is performed on the table.
	//
	// 	- 1: Sharding tasks are performed on the table.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTablesResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyList) GetAllowFullTableScan() *bool {
	return s.AllowFullTableScan
}

func (s *DescribeTablesResponseBodyList) GetBroadcast() *bool {
	return s.Broadcast
}

func (s *DescribeTablesResponseBodyList) GetDbInstType() *int32 {
	return s.DbInstType
}

func (s *DescribeTablesResponseBodyList) GetIsLocked() *bool {
	return s.IsLocked
}

func (s *DescribeTablesResponseBodyList) GetIsShard() *bool {
	return s.IsShard
}

func (s *DescribeTablesResponseBodyList) GetShardKey() *string {
	return s.ShardKey
}

func (s *DescribeTablesResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeTablesResponseBodyList) GetTable() *string {
	return s.Table
}

func (s *DescribeTablesResponseBodyList) SetAllowFullTableScan(v bool) *DescribeTablesResponseBodyList {
	s.AllowFullTableScan = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetBroadcast(v bool) *DescribeTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetDbInstType(v int32) *DescribeTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsLocked(v bool) *DescribeTablesResponseBodyList {
	s.IsLocked = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsShard(v bool) *DescribeTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetShardKey(v string) *DescribeTablesResponseBodyList {
	s.ShardKey = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetStatus(v int32) *DescribeTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetTable(v string) *DescribeTablesResponseBodyList {
	s.Table = &v
	return s
}

func (s *DescribeTablesResponseBodyList) Validate() error {
	return dara.Validate(s)
}
