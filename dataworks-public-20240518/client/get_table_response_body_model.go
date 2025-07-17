// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableResponseBody
	GetSuccess() *bool
	SetTable(v *Table) *GetTableResponseBody
	GetTable() *Table
}

type GetTableResponseBody struct {
	// example:
	//
	// 7B3435F4-2D91-XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool  `json:"Success,omitempty" xml:"Success,omitempty"`
	Table   *Table `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s GetTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableResponseBody) GetTable() *Table {
	return s.Table
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetSuccess(v bool) *GetTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableResponseBody) SetTable(v *Table) *GetTableResponseBody {
	s.Table = v
	return s
}

func (s *GetTableResponseBody) Validate() error {
	return dara.Validate(s)
}
