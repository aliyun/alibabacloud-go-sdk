// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnCount(v int32) *DescribeAdbMySqlColumnsResponseBody
	GetColumnCount() *int32
	SetColumns(v []*DescribeAdbMySqlColumnsResponseBodyColumns) *DescribeAdbMySqlColumnsResponseBody
	GetColumns() []*DescribeAdbMySqlColumnsResponseBodyColumns
	SetMessage(v string) *DescribeAdbMySqlColumnsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdbMySqlColumnsResponseBody
	GetRequestId() *string
	SetSchema(v string) *DescribeAdbMySqlColumnsResponseBody
	GetSchema() *string
	SetSuccess(v bool) *DescribeAdbMySqlColumnsResponseBody
	GetSuccess() *bool
	SetTableName(v string) *DescribeAdbMySqlColumnsResponseBody
	GetTableName() *string
}

type DescribeAdbMySqlColumnsResponseBody struct {
	// The total number of columns.
	//
	// example:
	//
	// 1
	ColumnCount *int32 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// Details of the columns.
	Columns []*DescribeAdbMySqlColumnsResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The message returned for the operation. Valid values:
	//
	// 	- **Success*	- is returned if the operation is successful.
	//
	// 	- An error message is returned if the operation fails.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A9F013CD-0222-595E-8157-445969B97F03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- **true**: The operation is successful.
	//
	// 	- **false**: The operation fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetColumnCount() *int32 {
	return s.ColumnCount
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetColumns() []*DescribeAdbMySqlColumnsResponseBodyColumns {
	return s.Columns
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdbMySqlColumnsResponseBody) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetColumnCount(v int32) *DescribeAdbMySqlColumnsResponseBody {
	s.ColumnCount = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetColumns(v []*DescribeAdbMySqlColumnsResponseBodyColumns) *DescribeAdbMySqlColumnsResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetMessage(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetRequestId(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetSchema(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetSuccess(v bool) *DescribeAdbMySqlColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetTableName(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.TableName = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAdbMySqlColumnsResponseBodyColumns struct {
	// The comments of the column.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAdbMySqlColumnsResponseBodyColumns) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlColumnsResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) GetComment() *string {
	return s.Comment
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) GetName() *string {
	return s.Name
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) GetType() *string {
	return s.Type
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) SetComment(v string) *DescribeAdbMySqlColumnsResponseBodyColumns {
	s.Comment = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) SetName(v string) *DescribeAdbMySqlColumnsResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) SetType(v string) *DescribeAdbMySqlColumnsResponseBodyColumns {
	s.Type = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) Validate() error {
	return dara.Validate(s)
}
