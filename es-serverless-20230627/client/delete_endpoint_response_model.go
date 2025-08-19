// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEndpointResponseBody) *DeleteEndpointResponse
	GetBody() *DeleteEndpointResponseBody
}

type DeleteEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEndpointResponse) GetBody() *DeleteEndpointResponseBody {
	return s.Body
}

func (s *DeleteEndpointResponse) SetHeaders(v map[string]*string) *DeleteEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointResponse) SetStatusCode(v int32) *DeleteEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointResponse) SetBody(v *DeleteEndpointResponseBody) *DeleteEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteEndpointResponse) Validate() error {
	return dara.Validate(s)
}
