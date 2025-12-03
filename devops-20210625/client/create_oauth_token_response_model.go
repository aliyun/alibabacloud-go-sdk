// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuthTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOAuthTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOAuthTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateOAuthTokenResponseBody) *CreateOAuthTokenResponse
	GetBody() *CreateOAuthTokenResponseBody
}

type CreateOAuthTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOAuthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOAuthTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOAuthTokenResponse) GetBody() *CreateOAuthTokenResponseBody {
	return s.Body
}

func (s *CreateOAuthTokenResponse) SetHeaders(v map[string]*string) *CreateOAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateOAuthTokenResponse) SetStatusCode(v int32) *CreateOAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOAuthTokenResponse) SetBody(v *CreateOAuthTokenResponseBody) *CreateOAuthTokenResponse {
	s.Body = v
	return s
}

func (s *CreateOAuthTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
