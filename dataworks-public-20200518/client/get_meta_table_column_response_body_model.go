// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableColumnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableColumnResponseBodyData) *GetMetaTableColumnResponseBody
	GetData() *GetMetaTableColumnResponseBodyData
	SetErrorCode(v string) *GetMetaTableColumnResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableColumnResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableColumnResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableColumnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableColumnResponseBody
	GetSuccess() *bool
}

type GetMetaTableColumnResponseBody struct {
	// The business data.
	Data *GetMetaTableColumnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
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

func (s GetMetaTableColumnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponseBody) GetData() *GetMetaTableColumnResponseBodyData {
	return s.Data
}

func (s *GetMetaTableColumnResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableColumnResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableColumnResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableColumnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableColumnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableColumnResponseBody) SetData(v *GetMetaTableColumnResponseBodyData) *GetMetaTableColumnResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetErrorCode(v string) *GetMetaTableColumnResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetErrorMessage(v string) *GetMetaTableColumnResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetHttpStatusCode(v int32) *GetMetaTableColumnResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetRequestId(v string) *GetMetaTableColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetSuccess(v bool) *GetMetaTableColumnResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableColumnResponseBodyData struct {
	// The information about fields.
	ColumnList []*GetMetaTableColumnResponseBodyDataColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
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
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaTableColumnResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponseBodyData) GetColumnList() []*GetMetaTableColumnResponseBodyDataColumnList {
	return s.ColumnList
}

func (s *GetMetaTableColumnResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaTableColumnResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableColumnResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaTableColumnResponseBodyData) SetColumnList(v []*GetMetaTableColumnResponseBodyDataColumnList) *GetMetaTableColumnResponseBodyData {
	s.ColumnList = v
	return s
}

func (s *GetMetaTableColumnResponseBodyData) SetPageNum(v int32) *GetMetaTableColumnResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyData) SetPageSize(v int32) *GetMetaTableColumnResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyData) SetTotalCount(v int64) *GetMetaTableColumnResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableColumnResponseBodyDataColumnList struct {
	// The description of the field.
	//
	// example:
	//
	// data column
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The GUID of the field.
	//
	// example:
	//
	// odps.engine_name.table_name.name
	ColumnGuid *string `json:"ColumnGuid,omitempty" xml:"ColumnGuid,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the field.
	//
	// example:
	//
	// string
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The remarks of the field.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether the field is a foreign key. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsForeignKey *bool `json:"IsForeignKey,omitempty" xml:"IsForeignKey,omitempty"`
	// Indicates whether the field is a partition field. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsPartitionColumn *bool `json:"IsPartitionColumn,omitempty" xml:"IsPartitionColumn,omitempty"`
	// Indicates whether the field is a primary key. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsPrimaryKey *bool `json:"IsPrimaryKey,omitempty" xml:"IsPrimaryKey,omitempty"`
	// The sequence number of the field.
	//
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	// The number of times the field is read.
	//
	// example:
	//
	// 2
	RelationCount *int64 `json:"RelationCount,omitempty" xml:"RelationCount,omitempty"`
}

func (s GetMetaTableColumnResponseBodyDataColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnResponseBodyDataColumnList) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetCaption() *string {
	return s.Caption
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetColumnGuid() *string {
	return s.ColumnGuid
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetComment() *string {
	return s.Comment
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetIsForeignKey() *bool {
	return s.IsForeignKey
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetIsPartitionColumn() *bool {
	return s.IsPartitionColumn
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetIsPrimaryKey() *bool {
	return s.IsPrimaryKey
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetPosition() *int32 {
	return s.Position
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) GetRelationCount() *int64 {
	return s.RelationCount
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetCaption(v string) *GetMetaTableColumnResponseBodyDataColumnList {
	s.Caption = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetColumnGuid(v string) *GetMetaTableColumnResponseBodyDataColumnList {
	s.ColumnGuid = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetColumnName(v string) *GetMetaTableColumnResponseBodyDataColumnList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetColumnType(v string) *GetMetaTableColumnResponseBodyDataColumnList {
	s.ColumnType = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetComment(v string) *GetMetaTableColumnResponseBodyDataColumnList {
	s.Comment = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetIsForeignKey(v bool) *GetMetaTableColumnResponseBodyDataColumnList {
	s.IsForeignKey = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetIsPartitionColumn(v bool) *GetMetaTableColumnResponseBodyDataColumnList {
	s.IsPartitionColumn = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetIsPrimaryKey(v bool) *GetMetaTableColumnResponseBodyDataColumnList {
	s.IsPrimaryKey = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetPosition(v int32) *GetMetaTableColumnResponseBodyDataColumnList {
	s.Position = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) SetRelationCount(v int64) *GetMetaTableColumnResponseBodyDataColumnList {
	s.RelationCount = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyDataColumnList) Validate() error {
	return dara.Validate(s)
}
