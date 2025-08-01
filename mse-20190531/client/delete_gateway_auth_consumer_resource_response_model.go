// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthConsumerResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayAuthConsumerResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayAuthConsumerResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayAuthConsumerResourceResponseBody) *DeleteGatewayAuthConsumerResourceResponse
	GetBody() *DeleteGatewayAuthConsumerResourceResponseBody
}

type DeleteGatewayAuthConsumerResourceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayAuthConsumerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayAuthConsumerResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthConsumerResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthConsumerResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayAuthConsumerResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayAuthConsumerResourceResponse) GetBody() *DeleteGatewayAuthConsumerResourceResponseBody {
	return s.Body
}

func (s *DeleteGatewayAuthConsumerResourceResponse) SetHeaders(v map[string]*string) *DeleteGatewayAuthConsumerResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponse) SetStatusCode(v int32) *DeleteGatewayAuthConsumerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponse) SetBody(v *DeleteGatewayAuthConsumerResourceResponseBody) *DeleteGatewayAuthConsumerResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponse) Validate() error {
	return dara.Validate(s)
}
