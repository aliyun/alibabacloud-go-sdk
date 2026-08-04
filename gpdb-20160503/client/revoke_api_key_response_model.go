// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *RevokeApiKeyResponseBody) *RevokeApiKeyResponse
	GetBody() *RevokeApiKeyResponseBody
}

type RevokeApiKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeApiKeyResponse) GoString() string {
	return s.String()
}

func (s *RevokeApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeApiKeyResponse) GetBody() *RevokeApiKeyResponseBody {
	return s.Body
}

func (s *RevokeApiKeyResponse) SetHeaders(v map[string]*string) *RevokeApiKeyResponse {
	s.Headers = v
	return s
}

func (s *RevokeApiKeyResponse) SetStatusCode(v int32) *RevokeApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApiKeyResponse) SetBody(v *RevokeApiKeyResponseBody) *RevokeApiKeyResponse {
	s.Body = v
	return s
}

func (s *RevokeApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
