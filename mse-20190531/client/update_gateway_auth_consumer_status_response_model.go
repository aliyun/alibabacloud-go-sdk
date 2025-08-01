// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayAuthConsumerStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayAuthConsumerStatusResponseBody) *UpdateGatewayAuthConsumerStatusResponse
	GetBody() *UpdateGatewayAuthConsumerStatusResponseBody
}

type UpdateGatewayAuthConsumerStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayAuthConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayAuthConsumerStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayAuthConsumerStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayAuthConsumerStatusResponse) GetBody() *UpdateGatewayAuthConsumerStatusResponseBody {
	return s.Body
}

func (s *UpdateGatewayAuthConsumerStatusResponse) SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponse) SetStatusCode(v int32) *UpdateGatewayAuthConsumerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponse) SetBody(v *UpdateGatewayAuthConsumerStatusResponseBody) *UpdateGatewayAuthConsumerStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponse) Validate() error {
	return dara.Validate(s)
}
