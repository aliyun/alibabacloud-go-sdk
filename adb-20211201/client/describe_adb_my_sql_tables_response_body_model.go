// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeAdbMySqlTablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdbMySqlTablesResponseBody
	GetRequestId() *string
	SetSchema(v string) *DescribeAdbMySqlTablesResponseBody
	GetSchema() *string
	SetSuccess(v bool) *DescribeAdbMySqlTablesResponseBody
	GetSuccess() *bool
	SetTables(v []*string) *DescribeAdbMySqlTablesResponseBody
	GetTables() []*string
}

type DescribeAdbMySqlTablesResponseBody struct {
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
	// 7A7D49E3-5585-5DF8-B62C-75C46B4991DC
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
	// The names of tables.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeAdbMySqlTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdbMySqlTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdbMySqlTablesResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdbMySqlTablesResponseBody) GetTables() []*string {
	return s.Tables
}

func (s *DescribeAdbMySqlTablesResponseBody) SetMessage(v string) *DescribeAdbMySqlTablesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetRequestId(v string) *DescribeAdbMySqlTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetSchema(v string) *DescribeAdbMySqlTablesResponseBody {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetSuccess(v bool) *DescribeAdbMySqlTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetTables(v []*string) *DescribeAdbMySqlTablesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
