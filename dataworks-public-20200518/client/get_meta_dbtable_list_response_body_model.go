// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaDBTableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaDBTableListResponseBodyData) *GetMetaDBTableListResponseBody
	GetData() *GetMetaDBTableListResponseBodyData
	SetRequestId(v string) *GetMetaDBTableListResponseBody
	GetRequestId() *string
}

type GetMetaDBTableListResponseBody struct {
	// The metatable information in a compute engine instance.
	Data *GetMetaDBTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMetaDBTableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBTableListResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaDBTableListResponseBody) GetData() *GetMetaDBTableListResponseBodyData {
	return s.Data
}

func (s *GetMetaDBTableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaDBTableListResponseBody) SetData(v *GetMetaDBTableListResponseBodyData) *GetMetaDBTableListResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaDBTableListResponseBody) SetRequestId(v string) *GetMetaDBTableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaDBTableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaDBTableListResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of metatables in the compute engine instance.
	TableEntityList []*GetMetaDBTableListResponseBodyDataTableEntityList `json:"TableEntityList,omitempty" xml:"TableEntityList,omitempty" type:"Repeated"`
	// The total number of compute engine instances returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaDBTableListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBTableListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaDBTableListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaDBTableListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaDBTableListResponseBodyData) GetTableEntityList() []*GetMetaDBTableListResponseBodyDataTableEntityList {
	return s.TableEntityList
}

func (s *GetMetaDBTableListResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaDBTableListResponseBodyData) SetPageNumber(v int32) *GetMetaDBTableListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetMetaDBTableListResponseBodyData) SetPageSize(v int32) *GetMetaDBTableListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaDBTableListResponseBodyData) SetTableEntityList(v []*GetMetaDBTableListResponseBodyDataTableEntityList) *GetMetaDBTableListResponseBodyData {
	s.TableEntityList = v
	return s
}

func (s *GetMetaDBTableListResponseBodyData) SetTotalCount(v int64) *GetMetaDBTableListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaDBTableListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaDBTableListResponseBodyDataTableEntityList struct {
	// The name of the metadatabase.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The GUID of the metatable.
	//
	// example:
	//
	// odps.engine_name.tname
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable.
	//
	// example:
	//
	// tname
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaDBTableListResponseBodyDataTableEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBTableListResponseBodyDataTableEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) SetDatabaseName(v string) *GetMetaDBTableListResponseBodyDataTableEntityList {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) SetTableGuid(v string) *GetMetaDBTableListResponseBodyDataTableEntityList {
	s.TableGuid = &v
	return s
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) SetTableName(v string) *GetMetaDBTableListResponseBodyDataTableEntityList {
	s.TableName = &v
	return s
}

func (s *GetMetaDBTableListResponseBodyDataTableEntityList) Validate() error {
	return dara.Validate(s)
}
