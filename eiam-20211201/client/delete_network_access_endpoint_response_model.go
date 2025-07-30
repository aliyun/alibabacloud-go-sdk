// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAccessEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkAccessEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkAccessEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkAccessEndpointResponseBody) *DeleteNetworkAccessEndpointResponse
	GetBody() *DeleteNetworkAccessEndpointResponseBody
}

type DeleteNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkAccessEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAccessEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkAccessEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkAccessEndpointResponse) GetBody() *DeleteNetworkAccessEndpointResponseBody {
	return s.Body
}

func (s *DeleteNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *DeleteNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkAccessEndpointResponse) SetStatusCode(v int32) *DeleteNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkAccessEndpointResponse) SetBody(v *DeleteNetworkAccessEndpointResponseBody) *DeleteNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkAccessEndpointResponse) Validate() error {
	return dara.Validate(s)
}
