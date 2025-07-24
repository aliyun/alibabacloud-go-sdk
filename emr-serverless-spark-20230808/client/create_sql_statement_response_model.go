// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSqlStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSqlStatementResponse
	GetStatusCode() *int32
	SetBody(v *CreateSqlStatementResponseBody) *CreateSqlStatementResponse
	GetBody() *CreateSqlStatementResponseBody
}

type CreateSqlStatementResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSqlStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSqlStatementResponse) GetBody() *CreateSqlStatementResponseBody {
	return s.Body
}

func (s *CreateSqlStatementResponse) SetHeaders(v map[string]*string) *CreateSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlStatementResponse) SetStatusCode(v int32) *CreateSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlStatementResponse) SetBody(v *CreateSqlStatementResponseBody) *CreateSqlStatementResponse {
	s.Body = v
	return s
}

func (s *CreateSqlStatementResponse) Validate() error {
	return dara.Validate(s)
}
