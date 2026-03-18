// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenByDerivedShortTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse
	GetStatusCode() *int32
	SetBody(v *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse
	GetBody() *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
}

type ObtainJwtAuthenticationTokenByDerivedShortTokenResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) GetBody() *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	return s.Body
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) SetHeaders(v map[string]*string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse {
	s.Headers = v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) SetStatusCode(v int32) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) SetBody(v *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse {
	s.Body = v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
