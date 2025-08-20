// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViewDDLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetViewDDLResponseBody
	GetRequestId() *string
	SetSQL(v string) *GetViewDDLResponseBody
	GetSQL() *string
}

type GetViewDDLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 421794A3-72A5-5D27-9E8B-A75A4C503E17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// CREATE VIEW `test`.`test_view` AS SELECT
	//
	//   `id`
	//
	// , `name`
	//
	// FROM
	//
	//   `test_tbl_adb`
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s GetViewDDLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetViewDDLResponseBody) GoString() string {
	return s.String()
}

func (s *GetViewDDLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetViewDDLResponseBody) GetSQL() *string {
	return s.SQL
}

func (s *GetViewDDLResponseBody) SetRequestId(v string) *GetViewDDLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetViewDDLResponseBody) SetSQL(v string) *GetViewDDLResponseBody {
	s.SQL = &v
	return s
}

func (s *GetViewDDLResponseBody) Validate() error {
	return dara.Validate(s)
}
