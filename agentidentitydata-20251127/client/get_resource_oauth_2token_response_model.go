// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOAuth2TokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceOAuth2TokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceOAuth2TokenResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceOAuth2TokenResponseBody) *GetResourceOAuth2TokenResponse
	GetBody() *GetResourceOAuth2TokenResponseBody
}

type GetResourceOAuth2TokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceOAuth2TokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceOAuth2TokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOAuth2TokenResponse) GoString() string {
	return s.String()
}

func (s *GetResourceOAuth2TokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceOAuth2TokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceOAuth2TokenResponse) GetBody() *GetResourceOAuth2TokenResponseBody {
	return s.Body
}

func (s *GetResourceOAuth2TokenResponse) SetHeaders(v map[string]*string) *GetResourceOAuth2TokenResponse {
	s.Headers = v
	return s
}

func (s *GetResourceOAuth2TokenResponse) SetStatusCode(v int32) *GetResourceOAuth2TokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceOAuth2TokenResponse) SetBody(v *GetResourceOAuth2TokenResponseBody) *GetResourceOAuth2TokenResponse {
	s.Body = v
	return s
}

func (s *GetResourceOAuth2TokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
