// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceEndpointResponseBody) *DeleteServiceEndpointResponse
	GetBody() *DeleteServiceEndpointResponseBody
}

type DeleteServiceEndpointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceEndpointResponse) GetBody() *DeleteServiceEndpointResponseBody {
	return s.Body
}

func (s *DeleteServiceEndpointResponse) SetHeaders(v map[string]*string) *DeleteServiceEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceEndpointResponse) SetStatusCode(v int32) *DeleteServiceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceEndpointResponse) SetBody(v *DeleteServiceEndpointResponseBody) *DeleteServiceEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
