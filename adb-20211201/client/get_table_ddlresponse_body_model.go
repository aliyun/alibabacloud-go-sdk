// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDDLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTableDDLResponseBody
	GetRequestId() *string
	SetSQL(v string) *GetTableDDLResponseBody
	GetSQL() *string
}

type GetTableDDLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 863D51B7-5321-41D8-A0B6-A088B0******
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

func (s GetTableDDLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableDDLResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDDLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableDDLResponseBody) GetSQL() *string {
	return s.SQL
}

func (s *GetTableDDLResponseBody) SetRequestId(v string) *GetTableDDLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDDLResponseBody) SetSQL(v string) *GetTableDDLResponseBody {
	s.SQL = &v
	return s
}

func (s *GetTableDDLResponseBody) Validate() error {
	return dara.Validate(s)
}
