// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaColumnLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaColumnLineageResponseBodyData) *GetMetaColumnLineageResponseBody
	GetData() *GetMetaColumnLineageResponseBodyData
	SetErrorCode(v string) *GetMetaColumnLineageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaColumnLineageResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaColumnLineageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaColumnLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaColumnLineageResponseBody
	GetSuccess() *bool
}

type GetMetaColumnLineageResponseBody struct {
	// The business data.
	Data *GetMetaColumnLineageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetMetaColumnLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaColumnLineageResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaColumnLineageResponseBody) GetData() *GetMetaColumnLineageResponseBodyData {
	return s.Data
}

func (s *GetMetaColumnLineageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaColumnLineageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaColumnLineageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaColumnLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaColumnLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaColumnLineageResponseBody) SetData(v *GetMetaColumnLineageResponseBodyData) *GetMetaColumnLineageResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaColumnLineageResponseBody) SetErrorCode(v string) *GetMetaColumnLineageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaColumnLineageResponseBody) SetErrorMessage(v string) *GetMetaColumnLineageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaColumnLineageResponseBody) SetHttpStatusCode(v int32) *GetMetaColumnLineageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaColumnLineageResponseBody) SetRequestId(v string) *GetMetaColumnLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaColumnLineageResponseBody) SetSuccess(v bool) *GetMetaColumnLineageResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaColumnLineageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaColumnLineageResponseBodyData struct {
	// The returned result.
	DataEntityList []*GetMetaColumnLineageResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of fields.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaColumnLineageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaColumnLineageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaColumnLineageResponseBodyData) GetDataEntityList() []*GetMetaColumnLineageResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *GetMetaColumnLineageResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaColumnLineageResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaColumnLineageResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaColumnLineageResponseBodyData) SetDataEntityList(v []*GetMetaColumnLineageResponseBodyDataDataEntityList) *GetMetaColumnLineageResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *GetMetaColumnLineageResponseBodyData) SetPageNum(v int32) *GetMetaColumnLineageResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyData) SetPageSize(v int32) *GetMetaColumnLineageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyData) SetTotalCount(v int64) *GetMetaColumnLineageResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaColumnLineageResponseBodyDataDataEntityList struct {
	// The EMR cluster ID.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The unique identifier of the field.
	//
	// example:
	//
	// odps.engine_name.table_name.1
	ColumnGuid *string `json:"ColumnGuid,omitempty" xml:"ColumnGuid,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// 1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaColumnLineageResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaColumnLineageResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) GetColumnGuid() *string {
	return s.ColumnGuid
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) SetClusterId(v string) *GetMetaColumnLineageResponseBodyDataDataEntityList {
	s.ClusterId = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) SetColumnGuid(v string) *GetMetaColumnLineageResponseBodyDataDataEntityList {
	s.ColumnGuid = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) SetColumnName(v string) *GetMetaColumnLineageResponseBodyDataDataEntityList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) SetDatabaseName(v string) *GetMetaColumnLineageResponseBodyDataDataEntityList {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) SetTableName(v string) *GetMetaColumnLineageResponseBodyDataDataEntityList {
	s.TableName = &v
	return s
}

func (s *GetMetaColumnLineageResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
