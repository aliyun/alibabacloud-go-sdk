// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayAuthConsumerResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayAuthConsumerResourceResponseBody) *UpdateGatewayAuthConsumerResourceResponse
	GetBody() *UpdateGatewayAuthConsumerResourceResponseBody
}

type UpdateGatewayAuthConsumerResourceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayAuthConsumerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayAuthConsumerResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayAuthConsumerResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayAuthConsumerResourceResponse) GetBody() *UpdateGatewayAuthConsumerResourceResponseBody {
	return s.Body
}

func (s *UpdateGatewayAuthConsumerResourceResponse) SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponse) SetStatusCode(v int32) *UpdateGatewayAuthConsumerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponse) SetBody(v *UpdateGatewayAuthConsumerResourceResponseBody) *UpdateGatewayAuthConsumerResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponse) Validate() error {
	return dara.Validate(s)
}
