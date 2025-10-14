// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomEndpointResponseBody) *DeleteCustomEndpointResponse
	GetBody() *DeleteCustomEndpointResponseBody
}

type DeleteCustomEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomEndpointResponse) GetBody() *DeleteCustomEndpointResponseBody {
	return s.Body
}

func (s *DeleteCustomEndpointResponse) SetHeaders(v map[string]*string) *DeleteCustomEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomEndpointResponse) SetStatusCode(v int32) *DeleteCustomEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomEndpointResponse) SetBody(v *DeleteCustomEndpointResponseBody) *DeleteCustomEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
