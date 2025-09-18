// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2TokenEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OAuth2TokenEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OAuth2TokenEndpointResponse
	GetStatusCode() *int32
	SetBody(v *OAuth2TokenEndpointResponseBody) *OAuth2TokenEndpointResponse
	GetBody() *OAuth2TokenEndpointResponseBody
}

type OAuth2TokenEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OAuth2TokenEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OAuth2TokenEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s OAuth2TokenEndpointResponse) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OAuth2TokenEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OAuth2TokenEndpointResponse) GetBody() *OAuth2TokenEndpointResponseBody {
	return s.Body
}

func (s *OAuth2TokenEndpointResponse) SetHeaders(v map[string]*string) *OAuth2TokenEndpointResponse {
	s.Headers = v
	return s
}

func (s *OAuth2TokenEndpointResponse) SetStatusCode(v int32) *OAuth2TokenEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *OAuth2TokenEndpointResponse) SetBody(v *OAuth2TokenEndpointResponseBody) *OAuth2TokenEndpointResponse {
	s.Body = v
	return s
}

func (s *OAuth2TokenEndpointResponse) Validate() error {
	return dara.Validate(s)
}
