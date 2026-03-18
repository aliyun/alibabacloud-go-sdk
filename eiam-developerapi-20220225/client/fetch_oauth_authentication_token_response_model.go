// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchOAuthAuthenticationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchOAuthAuthenticationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchOAuthAuthenticationTokenResponse
	GetStatusCode() *int32
	SetBody(v *FetchOAuthAuthenticationTokenResponseBody) *FetchOAuthAuthenticationTokenResponse
	GetBody() *FetchOAuthAuthenticationTokenResponseBody
}

type FetchOAuthAuthenticationTokenResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchOAuthAuthenticationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchOAuthAuthenticationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchOAuthAuthenticationTokenResponse) GoString() string {
	return s.String()
}

func (s *FetchOAuthAuthenticationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchOAuthAuthenticationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchOAuthAuthenticationTokenResponse) GetBody() *FetchOAuthAuthenticationTokenResponseBody {
	return s.Body
}

func (s *FetchOAuthAuthenticationTokenResponse) SetHeaders(v map[string]*string) *FetchOAuthAuthenticationTokenResponse {
	s.Headers = v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponse) SetStatusCode(v int32) *FetchOAuthAuthenticationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponse) SetBody(v *FetchOAuthAuthenticationTokenResponseBody) *FetchOAuthAuthenticationTokenResponse {
	s.Body = v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
