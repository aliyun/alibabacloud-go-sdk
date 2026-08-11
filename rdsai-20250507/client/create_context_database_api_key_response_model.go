// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDatabaseApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDatabaseApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDatabaseApiKeyResponseBody) *CreateContextDatabaseApiKeyResponse
	GetBody() *CreateContextDatabaseApiKeyResponseBody
}

type CreateContextDatabaseApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDatabaseApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDatabaseApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseApiKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDatabaseApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDatabaseApiKeyResponse) GetBody() *CreateContextDatabaseApiKeyResponseBody {
	return s.Body
}

func (s *CreateContextDatabaseApiKeyResponse) SetHeaders(v map[string]*string) *CreateContextDatabaseApiKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDatabaseApiKeyResponse) SetStatusCode(v int32) *CreateContextDatabaseApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponse) SetBody(v *CreateContextDatabaseApiKeyResponseBody) *CreateContextDatabaseApiKeyResponse {
	s.Body = v
	return s
}

func (s *CreateContextDatabaseApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
