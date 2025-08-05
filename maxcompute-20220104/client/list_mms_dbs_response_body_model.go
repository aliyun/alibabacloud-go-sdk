// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsDbsResponseBodyData) *ListMmsDbsResponseBody
	GetData() *ListMmsDbsResponseBodyData
	SetRequestId(v string) *ListMmsDbsResponseBody
	GetRequestId() *string
}

type ListMmsDbsResponseBody struct {
	Data *ListMmsDbsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// CF3F9978-260F-5204-94BE-30A4E6B54443
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsDbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponseBody) GetData() *ListMmsDbsResponseBodyData {
	return s.Data
}

func (s *ListMmsDbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsDbsResponseBody) SetData(v *ListMmsDbsResponseBodyData) *ListMmsDbsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsDbsResponseBody) SetRequestId(v string) *ListMmsDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsDbsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMmsDbsResponseBodyData struct {
	ObjectList []*ListMmsDbsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 13
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsDbsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponseBodyData) GetObjectList() []*ListMmsDbsResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsDbsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsDbsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsDbsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsDbsResponseBodyData) SetObjectList(v []*ListMmsDbsResponseBodyDataObjectList) *ListMmsDbsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsDbsResponseBodyData) SetPageNum(v int32) *ListMmsDbsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsDbsResponseBodyData) SetPageSize(v int32) *ListMmsDbsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsDbsResponseBodyData) SetTotal(v int32) *ListMmsDbsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsDbsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListMmsDbsResponseBodyDataObjectList struct {
	// example:
	//
	// 2024-12-17 15:44:42
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// for mms test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 1530
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Last DDL Time
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// hdfs://master-1-1.c-6fc187819ed6bae0.cn-shanghai.emr.aliyuncs.com:9000/user/hive/warehouse
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// mms_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 23232
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// xxx@yy.com
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 1000
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 400
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 200
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 200
	PartitionsFailed *int32 `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	// example:
	//
	// 2342342
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 100
	Tables *int32 `json:"tables,omitempty" xml:"tables,omitempty"`
	// example:
	//
	// 20
	TablesDoing *int32 `json:"tablesDoing,omitempty" xml:"tablesDoing,omitempty"`
	// example:
	//
	// 20
	TablesDone *int32 `json:"tablesDone,omitempty" xml:"tablesDone,omitempty"`
	// example:
	//
	// 20
	TablesFailed *int32 `json:"tablesFailed,omitempty" xml:"tablesFailed,omitempty"`
	// example:
	//
	// 20
	TablesPartDone *int32 `json:"tablesPartDone,omitempty" xml:"tablesPartDone,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListMmsDbsResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetDescription() *string {
	return s.Description
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetExtra() *string {
	return s.Extra
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetLocation() *string {
	return s.Location
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetName() *string {
	return s.Name
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetNumRows() *int64 {
	return s.NumRows
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetOwner() *string {
	return s.Owner
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetPartitions() *int32 {
	return s.Partitions
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetPartitionsDoing() *int32 {
	return s.PartitionsDoing
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetPartitionsDone() *int32 {
	return s.PartitionsDone
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetPartitionsFailed() *int32 {
	return s.PartitionsFailed
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetSize() *int64 {
	return s.Size
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetSourceName() *string {
	return s.SourceName
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetTables() *int32 {
	return s.Tables
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetTablesDoing() *int32 {
	return s.TablesDoing
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetTablesDone() *int32 {
	return s.TablesDone
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetTablesFailed() *int32 {
	return s.TablesFailed
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetTablesPartDone() *int32 {
	return s.TablesPartDone
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMmsDbsResponseBodyDataObjectList) GetUpdated() *bool {
	return s.Updated
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetDeleted(v bool) *ListMmsDbsResponseBodyDataObjectList {
	s.Deleted = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetDescription(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Description = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetExtra(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Extra = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetId(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetLastDdlTime(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetLocation(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Location = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetName(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetNumRows(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.NumRows = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetOwner(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Owner = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitions(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.Partitions = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitionsDoing(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.PartitionsDoing = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitionsDone(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.PartitionsDone = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitionsFailed(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.PartitionsFailed = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetSize(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetSourceName(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetStatus(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTables(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.Tables = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesDoing(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesDoing = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesDone(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesDone = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesFailed(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesFailed = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesPartDone(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesPartDone = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetUpdateTime(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.UpdateTime = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetUpdated(v bool) *ListMmsDbsResponseBodyDataObjectList {
	s.Updated = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}
