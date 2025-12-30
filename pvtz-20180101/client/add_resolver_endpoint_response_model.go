// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddResolverEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddResolverEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddResolverEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AddResolverEndpointResponseBody) *AddResolverEndpointResponse
	GetBody() *AddResolverEndpointResponseBody
}

type AddResolverEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddResolverEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s AddResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddResolverEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddResolverEndpointResponse) GetBody() *AddResolverEndpointResponseBody {
	return s.Body
}

func (s *AddResolverEndpointResponse) SetHeaders(v map[string]*string) *AddResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *AddResolverEndpointResponse) SetStatusCode(v int32) *AddResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResolverEndpointResponse) SetBody(v *AddResolverEndpointResponseBody) *AddResolverEndpointResponse {
	s.Body = v
	return s
}

func (s *AddResolverEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
