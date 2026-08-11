// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseApiKeyResponseBody) *UpdateContextDatabaseApiKeyResponse
	GetBody() *UpdateContextDatabaseApiKeyResponseBody
}

type UpdateContextDatabaseApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseApiKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseApiKeyResponse) GetBody() *UpdateContextDatabaseApiKeyResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseApiKeyResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseApiKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponse) SetStatusCode(v int32) *UpdateContextDatabaseApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponse) SetBody(v *UpdateContextDatabaseApiKeyResponseBody) *UpdateContextDatabaseApiKeyResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
