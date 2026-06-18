// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBranchSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetBranchSchemaResponseBody
	GetRequestId() *string
	SetSql(v string) *GetBranchSchemaResponseBody
	GetSql() *string
}

type GetBranchSchemaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The schema SQL content.
	//
	// example:
	//
	// CREATE TABLE public.example(id int);
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s GetBranchSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBranchSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBranchSchemaResponseBody) GetSql() *string {
	return s.Sql
}

func (s *GetBranchSchemaResponseBody) SetRequestId(v string) *GetBranchSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBranchSchemaResponseBody) SetSql(v string) *GetBranchSchemaResponseBody {
	s.Sql = &v
	return s
}

func (s *GetBranchSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
