// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResolverEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResolverEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResolverEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResolverEndpointResponseBody) *DeleteResolverEndpointResponse
	GetBody() *DeleteResolverEndpointResponseBody
}

type DeleteResolverEndpointResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResolverEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteResolverEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResolverEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResolverEndpointResponse) GetBody() *DeleteResolverEndpointResponseBody {
	return s.Body
}

func (s *DeleteResolverEndpointResponse) SetHeaders(v map[string]*string) *DeleteResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteResolverEndpointResponse) SetStatusCode(v int32) *DeleteResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResolverEndpointResponse) SetBody(v *DeleteResolverEndpointResponseBody) *DeleteResolverEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteResolverEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
