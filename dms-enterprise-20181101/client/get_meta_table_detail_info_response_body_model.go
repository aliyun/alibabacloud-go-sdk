// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableDetailInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetailInfo(v *GetMetaTableDetailInfoResponseBodyDetailInfo) *GetMetaTableDetailInfoResponseBody
	GetDetailInfo() *GetMetaTableDetailInfoResponseBodyDetailInfo
	SetErrorCode(v string) *GetMetaTableDetailInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableDetailInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetMetaTableDetailInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableDetailInfoResponseBody
	GetSuccess() *bool
}

type GetMetaTableDetailInfoResponseBody struct {
	// The details of the table.
	DetailInfo *GetMetaTableDetailInfoResponseBodyDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E881CB2F-DE42-42E5-90EB-8B3173DCB9B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableDetailInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBody) GetDetailInfo() *GetMetaTableDetailInfoResponseBodyDetailInfo {
	return s.DetailInfo
}

func (s *GetMetaTableDetailInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableDetailInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableDetailInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableDetailInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableDetailInfoResponseBody) SetDetailInfo(v *GetMetaTableDetailInfoResponseBodyDetailInfo) *GetMetaTableDetailInfoResponseBody {
	s.DetailInfo = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetErrorCode(v string) *GetMetaTableDetailInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetErrorMessage(v string) *GetMetaTableDetailInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetRequestId(v string) *GetMetaTableDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetSuccess(v bool) *GetMetaTableDetailInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableDetailInfoResponseBodyDetailInfo struct {
	// The columns in the table.
	ColumnList []*GetMetaTableDetailInfoResponseBodyDetailInfoColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// The list of indexes.
	IndexList []*GetMetaTableDetailInfoResponseBodyDetailInfoIndexList `json:"IndexList,omitempty" xml:"IndexList,omitempty" type:"Repeated"`
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfo) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) GetColumnList() []*GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	return s.ColumnList
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) GetIndexList() []*GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	return s.IndexList
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) SetColumnList(v []*GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) *GetMetaTableDetailInfoResponseBodyDetailInfo {
	s.ColumnList = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) SetIndexList(v []*GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) *GetMetaTableDetailInfoResponseBodyDetailInfo {
	s.IndexList = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableDetailInfoResponseBodyDetailInfoColumnList struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// 	- true: The column is an auto-increment column.
	//
	// 	- false: The column is not an auto-increment column.
	//
	// example:
	//
	// true
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The ID of the column.
	//
	// example:
	//
	// 191234849
	ColumnId *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the column. Examples: Bigint, Int, and Varchar.
	//
	// example:
	//
	// bigint(20) unsigned
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The length of the field.
	//
	// example:
	//
	// 0
	DataLength *int64 `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
	// The precision of the field.
	//
	// example:
	//
	// 0
	DataPrecision *int32 `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	// The scale of the column.
	//
	// example:
	//
	// 0
	DataScale *int32 `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	// The description of the column.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the column is nullable. Valid values:
	//
	// 	- true: The column is nullable.
	//
	// 	- false: The column is not nullable.
	//
	// example:
	//
	// false
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// The position of the field in the table.
	//
	// example:
	//
	// 1
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetColumnId() *string {
	return s.ColumnId
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetDataLength() *int64 {
	return s.DataLength
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetDataPrecision() *int32 {
	return s.DataPrecision
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetDataScale() *int32 {
	return s.DataScale
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetDescription() *string {
	return s.Description
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetNullable() *bool {
	return s.Nullable
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GetPosition() *string {
	return s.Position
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetAutoIncrement(v bool) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.AutoIncrement = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetColumnId(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.ColumnId = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetColumnName(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetColumnType(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.ColumnType = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDataLength(v int64) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.DataLength = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDataPrecision(v int32) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.DataPrecision = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDataScale(v int32) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.DataScale = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDescription(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.Description = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetNullable(v bool) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.Nullable = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetPosition(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.Position = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableDetailInfoResponseBodyDetailInfoIndexList struct {
	// The index column.
	IndexColumns []*string `json:"IndexColumns,omitempty" xml:"IndexColumns,omitempty" type:"Repeated"`
	// The ID of the index.
	//
	// example:
	//
	// 123
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// PRIMARY
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The type of the index. Examples: Primary, Unique, and Normal.
	//
	// example:
	//
	// Primary
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// Indicates whether the index is unique. Valid values:
	//
	// 	- true: The index is unique.
	//
	// 	- false: The index is not unique.
	//
	// example:
	//
	// false
	Unique *bool `json:"Unique,omitempty" xml:"Unique,omitempty"`
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GetIndexColumns() []*string {
	return s.IndexColumns
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GetIndexId() *string {
	return s.IndexId
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GetIndexName() *string {
	return s.IndexName
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GetIndexType() *string {
	return s.IndexType
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GetUnique() *bool {
	return s.Unique
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexColumns(v []*string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexColumns = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexId(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexId = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexName(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexName = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexType(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexType = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetUnique(v bool) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.Unique = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) Validate() error {
	return dara.Validate(s)
}
