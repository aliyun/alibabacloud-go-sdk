// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableLineageResponseBodyData) *GetMetaTableLineageResponseBody
	GetData() *GetMetaTableLineageResponseBodyData
	SetErrorCode(v string) *GetMetaTableLineageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableLineageResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableLineageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableLineageResponseBody
	GetSuccess() *bool
}

type GetMetaTableLineageResponseBody struct {
	// The business data.
	Data *GetMetaTableLineageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableLineageResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableLineageResponseBody) GetData() *GetMetaTableLineageResponseBodyData {
	return s.Data
}

func (s *GetMetaTableLineageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableLineageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableLineageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableLineageResponseBody) SetData(v *GetMetaTableLineageResponseBodyData) *GetMetaTableLineageResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableLineageResponseBody) SetErrorCode(v string) *GetMetaTableLineageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableLineageResponseBody) SetErrorMessage(v string) *GetMetaTableLineageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableLineageResponseBody) SetHttpStatusCode(v int32) *GetMetaTableLineageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableLineageResponseBody) SetRequestId(v string) *GetMetaTableLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableLineageResponseBody) SetSuccess(v bool) *GetMetaTableLineageResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableLineageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableLineageResponseBodyData struct {
	// The information about the table.
	DataEntityList []*GetMetaTableLineageResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// Indicates whether the next page exists.
	//
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// The logic of paging. If the value true is returned for the HasNext parameter and a value is returned for the NextPrimaryKey parameter in the response of the previous request, you must use the value of the NextPrimaryKey parameter for the next request.
	//
	// example:
	//
	// odps | retail_e_commerce_2 | retail_e_commerce_2 | dws_ec_trd__cate_commodity_gmv_kpy_fy
	NextPrimaryKey *string `json:"NextPrimaryKey,omitempty" xml:"NextPrimaryKey,omitempty"`
}

func (s GetMetaTableLineageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableLineageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableLineageResponseBodyData) GetDataEntityList() []*GetMetaTableLineageResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *GetMetaTableLineageResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *GetMetaTableLineageResponseBodyData) GetNextPrimaryKey() *string {
	return s.NextPrimaryKey
}

func (s *GetMetaTableLineageResponseBodyData) SetDataEntityList(v []*GetMetaTableLineageResponseBodyDataDataEntityList) *GetMetaTableLineageResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *GetMetaTableLineageResponseBodyData) SetHasNext(v bool) *GetMetaTableLineageResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *GetMetaTableLineageResponseBodyData) SetNextPrimaryKey(v string) *GetMetaTableLineageResponseBodyData {
	s.NextPrimaryKey = &v
	return s
}

func (s *GetMetaTableLineageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableLineageResponseBodyDataDataEntityList struct {
	// The time when the table was created.
	//
	// example:
	//
	// 1638720736000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// db1
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The unique identifier of the table.
	//
	// example:
	//
	// odps.tt.name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTableLineageResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableLineageResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) SetCreateTimestamp(v int64) *GetMetaTableLineageResponseBodyDataDataEntityList {
	s.CreateTimestamp = &v
	return s
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) SetDatabaseName(v string) *GetMetaTableLineageResponseBodyDataDataEntityList {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) SetTableGuid(v string) *GetMetaTableLineageResponseBodyDataDataEntityList {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) SetTableName(v string) *GetMetaTableLineageResponseBodyDataDataEntityList {
	s.TableName = &v
	return s
}

func (s *GetMetaTableLineageResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
