// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSqlStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateSqlStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateSqlStatementResponse
	GetStatusCode() *int32
	SetBody(v *ValidateSqlStatementResponseBody) *ValidateSqlStatementResponse
	GetBody() *ValidateSqlStatementResponseBody
}

type ValidateSqlStatementResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateSqlStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateSqlStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateSqlStatementResponse) GetBody() *ValidateSqlStatementResponseBody {
	return s.Body
}

func (s *ValidateSqlStatementResponse) SetHeaders(v map[string]*string) *ValidateSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *ValidateSqlStatementResponse) SetStatusCode(v int32) *ValidateSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateSqlStatementResponse) SetBody(v *ValidateSqlStatementResponseBody) *ValidateSqlStatementResponse {
	s.Body = v
	return s
}

func (s *ValidateSqlStatementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
