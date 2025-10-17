// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlStatementResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlStatementResponseBody) *GetSqlStatementResponse
	GetBody() *GetSqlStatementResponseBody
}

type GetSqlStatementResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlStatementResponse) GetBody() *GetSqlStatementResponseBody {
	return s.Body
}

func (s *GetSqlStatementResponse) SetHeaders(v map[string]*string) *GetSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *GetSqlStatementResponse) SetStatusCode(v int32) *GetSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlStatementResponse) SetBody(v *GetSqlStatementResponseBody) *GetSqlStatementResponse {
	s.Body = v
	return s
}

func (s *GetSqlStatementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
