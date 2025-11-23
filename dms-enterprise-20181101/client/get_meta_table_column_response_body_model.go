// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableColumnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnList(v []*GetMetaTableColumnResponseBodyColumnList) *GetMetaTableColumnResponseBody
	GetColumnList() []*GetMetaTableColumnResponseBodyColumnList
	SetErrorCode(v string) *GetMetaTableColumnResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableColumnResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetMetaTableColumnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableColumnResponseBody
	GetSuccess() *bool
}

type GetMetaTableColumnResponseBody struct {
	// The details about fields in the table.
	ColumnList []*GetMetaTableColumnResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// The error code returned.
	//
	// example:
	//
	// MissingTableGuid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// TableGuid is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 087DFBA1-378B-5D25-B13B-31F6409F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
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

func (s *GetMetaTableColumnResponseBody) GetColumnList() []*GetMetaTableColumnResponseBodyColumnList {
	return s.ColumnList
}

func (s *GetMetaTableColumnResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableColumnResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableColumnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableColumnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableColumnResponseBody) SetColumnList(v []*GetMetaTableColumnResponseBodyColumnList) *GetMetaTableColumnResponseBody {
	s.ColumnList = v
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

func (s *GetMetaTableColumnResponseBody) SetRequestId(v string) *GetMetaTableColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetSuccess(v bool) *GetMetaTableColumnResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) Validate() error {
	if s.ColumnList != nil {
		for _, item := range s.ColumnList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMetaTableColumnResponseBodyColumnList struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// 	- **true**: The column is an auto-increment column.
	//
	// 	- **false**: The column is not an auto-increment column.
	//
	// example:
	//
	// false
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The ID of the column.
	//
	// example:
	//
	// 63513****
	ColumnId *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// has_promotion
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the column.
	//
	// > The return value of a column is not unique, such as **bigint*	- or **int**.
	//
	// example:
	//
	// bigint(1)
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
	// 19
	DataPrecision *int32 `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	// The number of decimal places for the field.
	//
	// example:
	//
	// 0
	DataScale *int32 `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	// The description of the column.
	//
	// example:
	//
	// Whether discounts are provided
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the field can be empty. Valid values:
	//
	// 	- **true**: The field can be empty.
	//
	// 	- **false**: The field cannot be empty.
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
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Indicates whether the field is the primary key. Valid values:
	//
	// 	- **true**: The field is the primary key.
	//
	// 	- **false**: The field is not the primary key.
	//
	// example:
	//
	// true
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The sensitivity level of the column. Valid values:
	//
	// 	- **INNER**: The column is not sensitive.
	//
	// 	- **SENSITIVE**: The column is sensitive.
	//
	// 	- **CONFIDENTIAL**: The column is confidential.
	//
	// > For more information, see [Sensitivity levels of columns](https://help.aliyun.com/document_detail/66091.html).
	//
	// example:
	//
	// INNER
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetMetaTableColumnResponseBodyColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetColumnId() *string {
	return s.ColumnId
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetDataLength() *int64 {
	return s.DataLength
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetDataPrecision() *int32 {
	return s.DataPrecision
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetDataScale() *int32 {
	return s.DataScale
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetDescription() *string {
	return s.Description
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetNullable() *bool {
	return s.Nullable
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetPosition() *int32 {
	return s.Position
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetMetaTableColumnResponseBodyColumnList) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetAutoIncrement(v bool) *GetMetaTableColumnResponseBodyColumnList {
	s.AutoIncrement = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetColumnId(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.ColumnId = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetColumnName(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetColumnType(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.ColumnType = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDataLength(v int64) *GetMetaTableColumnResponseBodyColumnList {
	s.DataLength = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDataPrecision(v int32) *GetMetaTableColumnResponseBodyColumnList {
	s.DataPrecision = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDataScale(v int32) *GetMetaTableColumnResponseBodyColumnList {
	s.DataScale = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDescription(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.Description = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetNullable(v bool) *GetMetaTableColumnResponseBodyColumnList {
	s.Nullable = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetPosition(v int32) *GetMetaTableColumnResponseBodyColumnList {
	s.Position = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetPrimaryKey(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.PrimaryKey = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetSecurityLevel(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.SecurityLevel = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) Validate() error {
	return dara.Validate(s)
}
