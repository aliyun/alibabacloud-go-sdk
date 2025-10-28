// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecureTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecureTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecureTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetSecureTokenResponseBody) *GetSecureTokenResponse
	GetBody() *GetSecureTokenResponseBody
}

type GetSecureTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecureTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecureTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecureTokenResponse) GoString() string {
	return s.String()
}

func (s *GetSecureTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecureTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecureTokenResponse) GetBody() *GetSecureTokenResponseBody {
	return s.Body
}

func (s *GetSecureTokenResponse) SetHeaders(v map[string]*string) *GetSecureTokenResponse {
	s.Headers = v
	return s
}

func (s *GetSecureTokenResponse) SetStatusCode(v int32) *GetSecureTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecureTokenResponse) SetBody(v *GetSecureTokenResponseBody) *GetSecureTokenResponse {
	s.Body = v
	return s
}

func (s *GetSecureTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
