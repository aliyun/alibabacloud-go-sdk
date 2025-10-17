// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSqlStatementResponseBodyData) *CreateSqlStatementResponseBody
	GetData() *CreateSqlStatementResponseBodyData
	SetRequestId(v string) *CreateSqlStatementResponseBody
	GetRequestId() *string
}

type CreateSqlStatementResponseBody struct {
	// The data returned.
	Data *CreateSqlStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSqlStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementResponseBody) GetData() *CreateSqlStatementResponseBodyData {
	return s.Data
}

func (s *CreateSqlStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSqlStatementResponseBody) SetData(v *CreateSqlStatementResponseBodyData) *CreateSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *CreateSqlStatementResponseBody) SetRequestId(v string) *CreateSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlStatementResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSqlStatementResponseBodyData struct {
	// The interactive query ID.
	//
	// example:
	//
	// st-1231dfafadfa***
	StatementId *string `json:"statementId,omitempty" xml:"statementId,omitempty"`
}

func (s CreateSqlStatementResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementResponseBodyData) GetStatementId() *string {
	return s.StatementId
}

func (s *CreateSqlStatementResponseBodyData) SetStatementId(v string) *CreateSqlStatementResponseBodyData {
	s.StatementId = &v
	return s
}

func (s *CreateSqlStatementResponseBodyData) Validate() error {
	return dara.Validate(s)
}
