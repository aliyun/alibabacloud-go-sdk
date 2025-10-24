// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsPartitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsPartitionsResponseBodyData) *ListMmsPartitionsResponseBody
	GetData() *ListMmsPartitionsResponseBodyData
	SetRequestId(v string) *ListMmsPartitionsResponseBody
	GetRequestId() *string
}

type ListMmsPartitionsResponseBody struct {
	Data *ListMmsPartitionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// D9F872FD-5DDE-30A6-8C8A-1B8C6A81059F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsPartitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponseBody) GetData() *ListMmsPartitionsResponseBodyData {
	return s.Data
}

func (s *ListMmsPartitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsPartitionsResponseBody) SetData(v *ListMmsPartitionsResponseBodyData) *ListMmsPartitionsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsPartitionsResponseBody) SetRequestId(v string) *ListMmsPartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsPartitionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsPartitionsResponseBodyData struct {
	ObjectList []*ListMmsPartitionsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
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
	// 1000
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsPartitionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponseBodyData) GetObjectList() []*ListMmsPartitionsResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsPartitionsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsPartitionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsPartitionsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsPartitionsResponseBodyData) SetObjectList(v []*ListMmsPartitionsResponseBodyDataObjectList) *ListMmsPartitionsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) SetPageNum(v int32) *ListMmsPartitionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) SetPageSize(v int32) *ListMmsPartitionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) SetTotal(v int32) *ListMmsPartitionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) Validate() error {
	if s.ObjectList != nil {
		for _, item := range s.ObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsPartitionsResponseBodyDataObjectList struct {
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// mms_test
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// example:
	//
	// default
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	DstValue *string `json:"dstValue,omitempty" xml:"dstValue,omitempty"`
	// example:
	//
	// 2323
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// lastDdlTime
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// 2323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// 23223
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 200018
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 23
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsPartitionsResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetDbName() *string {
	return s.DbName
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetDstTableName() *string {
	return s.DstTableName
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetDstValue() *string {
	return s.DstValue
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetNumRows() *int64 {
	return s.NumRows
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetSize() *int64 {
	return s.Size
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetSourceName() *string {
	return s.SourceName
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetTableId() *int64 {
	return s.TableId
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetTableName() *string {
	return s.TableName
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetUpdated() *bool {
	return s.Updated
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) GetValue() *string {
	return s.Value
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDbId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDbName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DbName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDstProjectName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DstProjectName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDstSchemaName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DstSchemaName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDstTableName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DstTableName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDstValue(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DstValue = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetLastDdlTime(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetNumRows(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.NumRows = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetSize(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetSourceName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetStatus(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetTableId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.TableId = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetTableName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.TableName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetUpdated(v bool) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Updated = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetValue(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Value = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}
