// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthConsumerResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayAuthConsumerResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayAuthConsumerResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayAuthConsumerResourceResponseBody) *ListGatewayAuthConsumerResourceResponse
	GetBody() *ListGatewayAuthConsumerResourceResponseBody
}

type ListGatewayAuthConsumerResourceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayAuthConsumerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayAuthConsumerResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResourceResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayAuthConsumerResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayAuthConsumerResourceResponse) GetBody() *ListGatewayAuthConsumerResourceResponseBody {
	return s.Body
}

func (s *ListGatewayAuthConsumerResourceResponse) SetHeaders(v map[string]*string) *ListGatewayAuthConsumerResourceResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponse) SetStatusCode(v int32) *ListGatewayAuthConsumerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponse) SetBody(v *ListGatewayAuthConsumerResourceResponseBody) *ListGatewayAuthConsumerResourceResponse {
	s.Body = v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponse) Validate() error {
	return dara.Validate(s)
}
