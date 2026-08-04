// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *GrantApiKeyResponseBody) *GrantApiKeyResponse
	GetBody() *GrantApiKeyResponseBody
}

type GrantApiKeyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantApiKeyResponse) GoString() string {
	return s.String()
}

func (s *GrantApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantApiKeyResponse) GetBody() *GrantApiKeyResponseBody {
	return s.Body
}

func (s *GrantApiKeyResponse) SetHeaders(v map[string]*string) *GrantApiKeyResponse {
	s.Headers = v
	return s
}

func (s *GrantApiKeyResponse) SetStatusCode(v int32) *GrantApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantApiKeyResponse) SetBody(v *GrantApiKeyResponseBody) *GrantApiKeyResponse {
	s.Body = v
	return s
}

func (s *GrantApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
