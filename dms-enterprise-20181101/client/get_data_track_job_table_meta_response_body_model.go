// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackJobTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataTrackJobTableMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataTrackJobTableMetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataTrackJobTableMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataTrackJobTableMetaResponseBody
	GetSuccess() *bool
	SetTableMetaList(v []*GetDataTrackJobTableMetaResponseBodyTableMetaList) *GetDataTrackJobTableMetaResponseBody
	GetTableMetaList() []*GetDataTrackJobTableMetaResponseBodyTableMetaList
}

type GetDataTrackJobTableMetaResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
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
	// The metadata of tables.
	TableMetaList []*GetDataTrackJobTableMetaResponseBodyTableMetaList `json:"TableMetaList,omitempty" xml:"TableMetaList,omitempty" type:"Repeated"`
}

func (s GetDataTrackJobTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobTableMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataTrackJobTableMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataTrackJobTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataTrackJobTableMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataTrackJobTableMetaResponseBody) GetTableMetaList() []*GetDataTrackJobTableMetaResponseBodyTableMetaList {
	return s.TableMetaList
}

func (s *GetDataTrackJobTableMetaResponseBody) SetErrorCode(v string) *GetDataTrackJobTableMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBody) SetErrorMessage(v string) *GetDataTrackJobTableMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBody) SetRequestId(v string) *GetDataTrackJobTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBody) SetSuccess(v bool) *GetDataTrackJobTableMetaResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBody) SetTableMetaList(v []*GetDataTrackJobTableMetaResponseBodyTableMetaList) *GetDataTrackJobTableMetaResponseBody {
	s.TableMetaList = v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBody) Validate() error {
	if s.TableMetaList != nil {
		for _, item := range s.TableMetaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataTrackJobTableMetaResponseBodyTableMetaList struct {
	// The information about columns.
	Columns []*GetDataTrackJobTableMetaResponseBodyTableMetaListColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The name of the database.
	//
	// example:
	//
	// DB165
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// live_stat
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetDataTrackJobTableMetaResponseBodyTableMetaList) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobTableMetaResponseBodyTableMetaList) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) GetColumns() []*GetDataTrackJobTableMetaResponseBodyTableMetaListColumns {
	return s.Columns
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) GetTableName() *string {
	return s.TableName
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) SetColumns(v []*GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) *GetDataTrackJobTableMetaResponseBodyTableMetaList {
	s.Columns = v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) SetSchemaName(v string) *GetDataTrackJobTableMetaResponseBodyTableMetaList {
	s.SchemaName = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) SetTableName(v string) *GetDataTrackJobTableMetaResponseBodyTableMetaList {
	s.TableName = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaList) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataTrackJobTableMetaResponseBodyTableMetaListColumns struct {
	// The name of the character set.
	//
	// example:
	//
	// utf8mb4
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// claimantno
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The position of the column.
	//
	// example:
	//
	// 1
	ColumnPosition *int32 `json:"ColumnPosition,omitempty" xml:"ColumnPosition,omitempty"`
	// The data type of the column. Examples: BIGINT, INT, and VARCHAR.
	//
	// example:
	//
	// BIGINT
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// Indicates whether the column is a virtual column. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fictive *bool `json:"Fictive,omitempty" xml:"Fictive,omitempty"`
}

func (s GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) GetCharset() *string {
	return s.Charset
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) GetColumnPosition() *int32 {
	return s.ColumnPosition
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) GetFictive() *bool {
	return s.Fictive
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) SetCharset(v string) *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns {
	s.Charset = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) SetColumnName(v string) *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns {
	s.ColumnName = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) SetColumnPosition(v int32) *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns {
	s.ColumnPosition = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) SetColumnType(v string) *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns {
	s.ColumnType = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) SetFictive(v bool) *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns {
	s.Fictive = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponseBodyTableMetaListColumns) Validate() error {
	return dara.Validate(s)
}
