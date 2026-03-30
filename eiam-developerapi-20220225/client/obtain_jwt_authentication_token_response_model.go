// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainJwtAuthenticationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainJwtAuthenticationTokenResponse
	GetStatusCode() *int32
	SetBody(v *ObtainJwtAuthenticationTokenResponseBody) *ObtainJwtAuthenticationTokenResponse
	GetBody() *ObtainJwtAuthenticationTokenResponseBody
}

type ObtainJwtAuthenticationTokenResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainJwtAuthenticationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainJwtAuthenticationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenResponse) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainJwtAuthenticationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainJwtAuthenticationTokenResponse) GetBody() *ObtainJwtAuthenticationTokenResponseBody {
	return s.Body
}

func (s *ObtainJwtAuthenticationTokenResponse) SetHeaders(v map[string]*string) *ObtainJwtAuthenticationTokenResponse {
	s.Headers = v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponse) SetStatusCode(v int32) *ObtainJwtAuthenticationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponse) SetBody(v *ObtainJwtAuthenticationTokenResponseBody) *ObtainJwtAuthenticationTokenResponse {
	s.Body = v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
