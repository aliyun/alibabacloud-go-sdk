// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSqlExecuteCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgenticSqlExecuteCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgenticSqlExecuteCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgenticSqlExecuteCredentialResponseBody) *CreateAgenticSqlExecuteCredentialResponse
	GetBody() *CreateAgenticSqlExecuteCredentialResponseBody
}

type CreateAgenticSqlExecuteCredentialResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgenticSqlExecuteCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgenticSqlExecuteCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSqlExecuteCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateAgenticSqlExecuteCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgenticSqlExecuteCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgenticSqlExecuteCredentialResponse) GetBody() *CreateAgenticSqlExecuteCredentialResponseBody {
	return s.Body
}

func (s *CreateAgenticSqlExecuteCredentialResponse) SetHeaders(v map[string]*string) *CreateAgenticSqlExecuteCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponse) SetStatusCode(v int32) *CreateAgenticSqlExecuteCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponse) SetBody(v *CreateAgenticSqlExecuteCredentialResponseBody) *CreateAgenticSqlExecuteCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
