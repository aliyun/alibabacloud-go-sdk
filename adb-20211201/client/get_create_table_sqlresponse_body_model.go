// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateTableSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetCreateTableSQLResponseBody
	GetRequestId() *string
	SetSQL(v string) *GetCreateTableSQLResponseBody
	GetSQL() *string
}

type GetCreateTableSQLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 109462AF-B5FA-3D5A-9377-B27E5B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// create table (
	//
	//  id varchar(32)
	//
	// );
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s GetCreateTableSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCreateTableSQLResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateTableSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCreateTableSQLResponseBody) GetSQL() *string {
	return s.SQL
}

func (s *GetCreateTableSQLResponseBody) SetRequestId(v string) *GetCreateTableSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateTableSQLResponseBody) SetSQL(v string) *GetCreateTableSQLResponseBody {
	s.SQL = &v
	return s
}

func (s *GetCreateTableSQLResponseBody) Validate() error {
	return dara.Validate(s)
}
