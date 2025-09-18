// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2RevocationEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OAuth2RevocationEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OAuth2RevocationEndpointResponse
	GetStatusCode() *int32
	SetBody(v *OAuth2RevocationEndpointResponseBody) *OAuth2RevocationEndpointResponse
	GetBody() *OAuth2RevocationEndpointResponseBody
}

type OAuth2RevocationEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OAuth2RevocationEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OAuth2RevocationEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s OAuth2RevocationEndpointResponse) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OAuth2RevocationEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OAuth2RevocationEndpointResponse) GetBody() *OAuth2RevocationEndpointResponseBody {
	return s.Body
}

func (s *OAuth2RevocationEndpointResponse) SetHeaders(v map[string]*string) *OAuth2RevocationEndpointResponse {
	s.Headers = v
	return s
}

func (s *OAuth2RevocationEndpointResponse) SetStatusCode(v int32) *OAuth2RevocationEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *OAuth2RevocationEndpointResponse) SetBody(v *OAuth2RevocationEndpointResponseBody) *OAuth2RevocationEndpointResponse {
	s.Body = v
	return s
}

func (s *OAuth2RevocationEndpointResponse) Validate() error {
	return dara.Validate(s)
}
