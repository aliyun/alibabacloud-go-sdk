// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersForNetworkAccessEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentityProvidersForNetworkAccessEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentityProvidersForNetworkAccessEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ListIdentityProvidersForNetworkAccessEndpointResponseBody) *ListIdentityProvidersForNetworkAccessEndpointResponse
	GetBody() *ListIdentityProvidersForNetworkAccessEndpointResponseBody
}

type ListIdentityProvidersForNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentityProvidersForNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityProvidersForNetworkAccessEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersForNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) GetBody() *ListIdentityProvidersForNetworkAccessEndpointResponseBody {
	return s.Body
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *ListIdentityProvidersForNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) SetStatusCode(v int32) *ListIdentityProvidersForNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) SetBody(v *ListIdentityProvidersForNetworkAccessEndpointResponseBody) *ListIdentityProvidersForNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
