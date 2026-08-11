// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeContextDatabaseApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeContextDatabaseApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeContextDatabaseApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *RevokeContextDatabaseApiKeyResponseBody) *RevokeContextDatabaseApiKeyResponse
	GetBody() *RevokeContextDatabaseApiKeyResponseBody
}

type RevokeContextDatabaseApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeContextDatabaseApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeContextDatabaseApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeContextDatabaseApiKeyResponse) GoString() string {
	return s.String()
}

func (s *RevokeContextDatabaseApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeContextDatabaseApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeContextDatabaseApiKeyResponse) GetBody() *RevokeContextDatabaseApiKeyResponseBody {
	return s.Body
}

func (s *RevokeContextDatabaseApiKeyResponse) SetHeaders(v map[string]*string) *RevokeContextDatabaseApiKeyResponse {
	s.Headers = v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponse) SetStatusCode(v int32) *RevokeContextDatabaseApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponse) SetBody(v *RevokeContextDatabaseApiKeyResponseBody) *RevokeContextDatabaseApiKeyResponse {
	s.Body = v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
