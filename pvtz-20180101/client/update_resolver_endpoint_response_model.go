// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResolverEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResolverEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResolverEndpointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResolverEndpointResponseBody) *UpdateResolverEndpointResponse
	GetBody() *UpdateResolverEndpointResponseBody
}

type UpdateResolverEndpointResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResolverEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResolverEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResolverEndpointResponse) GetBody() *UpdateResolverEndpointResponseBody {
	return s.Body
}

func (s *UpdateResolverEndpointResponse) SetHeaders(v map[string]*string) *UpdateResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateResolverEndpointResponse) SetStatusCode(v int32) *UpdateResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResolverEndpointResponse) SetBody(v *UpdateResolverEndpointResponseBody) *UpdateResolverEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateResolverEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
