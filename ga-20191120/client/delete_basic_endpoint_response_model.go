// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBasicEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBasicEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBasicEndpointResponseBody) *DeleteBasicEndpointResponse
	GetBody() *DeleteBasicEndpointResponseBody
}

type DeleteBasicEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBasicEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBasicEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBasicEndpointResponse) GetBody() *DeleteBasicEndpointResponseBody {
	return s.Body
}

func (s *DeleteBasicEndpointResponse) SetHeaders(v map[string]*string) *DeleteBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicEndpointResponse) SetStatusCode(v int32) *DeleteBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicEndpointResponse) SetBody(v *DeleteBasicEndpointResponseBody) *DeleteBasicEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteBasicEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
