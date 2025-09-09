// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBroadcastTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsShard(v bool) *DescribeBroadcastTablesResponseBody
	GetIsShard() *bool
	SetList(v []*DescribeBroadcastTablesResponseBodyList) *DescribeBroadcastTablesResponseBody
	GetList() []*DescribeBroadcastTablesResponseBodyList
	SetPageNumber(v int32) *DescribeBroadcastTablesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBroadcastTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBroadcastTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBroadcastTablesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeBroadcastTablesResponseBody
	GetTotal() *int32
}

type DescribeBroadcastTablesResponseBody struct {
	// Indicates whether the database is sharded.
	//
	// example:
	//
	// true
	IsShard *bool `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	// Indicates information about broadcast tables.
	List []*DescribeBroadcastTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned per page.
	//
	// example:
	//
	// 40
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 86E420ED-43F2-4788-A58C-921849******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBroadcastTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBody) GetIsShard() *bool {
	return s.IsShard
}

func (s *DescribeBroadcastTablesResponseBody) GetList() []*DescribeBroadcastTablesResponseBodyList {
	return s.List
}

func (s *DescribeBroadcastTablesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBroadcastTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBroadcastTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBroadcastTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBroadcastTablesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeBroadcastTablesResponseBody) SetIsShard(v bool) *DescribeBroadcastTablesResponseBody {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetList(v []*DescribeBroadcastTablesResponseBodyList) *DescribeBroadcastTablesResponseBody {
	s.List = v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetPageNumber(v int32) *DescribeBroadcastTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetPageSize(v int32) *DescribeBroadcastTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetRequestId(v string) *DescribeBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetSuccess(v bool) *DescribeBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetTotal(v int32) *DescribeBroadcastTablesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBroadcastTablesResponseBodyList struct {
	// Indicates whether a table is a broadcast table.
	//
	// example:
	//
	// true
	Broadcast *bool `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	// Indicates the type of the broadcast table. Valid values:
	//
	// 	- **1**: multi-write mode
	//
	// 	- **2**: synchronous mode
	//
	// example:
	//
	// 1
	BroadcastType *string `json:"BroadcastType,omitempty" xml:"BroadcastType,omitempty"`
	// Indicates the storage type of the database. Valid values:
	//
	// 	- **0**: RDS
	//
	// 	- **4**: PolarDB
	//
	// example:
	//
	// 0
	DbInstType *int32 `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates whether the broadcast table was sharded.
	//
	// example:
	//
	// false
	IsShard *bool `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	// Indicates the activation state of the broadcast table. Valid values:
	//
	// 	- **1**: The broadcast table is activated.
	//
	// 	- **2**: The broadcast table is being activated.
	//
	// 	- **3**: An exception occurs when the broadcast table is being activated.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates the name of the table.
	//
	// example:
	//
	// nation
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBroadcastTablesResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBodyList) GetBroadcast() *bool {
	return s.Broadcast
}

func (s *DescribeBroadcastTablesResponseBodyList) GetBroadcastType() *string {
	return s.BroadcastType
}

func (s *DescribeBroadcastTablesResponseBodyList) GetDbInstType() *int32 {
	return s.DbInstType
}

func (s *DescribeBroadcastTablesResponseBodyList) GetIsShard() *bool {
	return s.IsShard
}

func (s *DescribeBroadcastTablesResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeBroadcastTablesResponseBodyList) GetTable() *string {
	return s.Table
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcast(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcastType(v string) *DescribeBroadcastTablesResponseBodyList {
	s.BroadcastType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetDbInstType(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetIsShard(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetStatus(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetTable(v string) *DescribeBroadcastTablesResponseBodyList {
	s.Table = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) Validate() error {
	return dara.Validate(s)
}
