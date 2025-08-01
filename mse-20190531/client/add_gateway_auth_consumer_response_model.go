// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayAuthConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayAuthConsumerResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayAuthConsumerResponseBody) *AddGatewayAuthConsumerResponse
	GetBody() *AddGatewayAuthConsumerResponseBody
}

type AddGatewayAuthConsumerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayAuthConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayAuthConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthConsumerResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayAuthConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayAuthConsumerResponse) GetBody() *AddGatewayAuthConsumerResponseBody {
	return s.Body
}

func (s *AddGatewayAuthConsumerResponse) SetHeaders(v map[string]*string) *AddGatewayAuthConsumerResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayAuthConsumerResponse) SetStatusCode(v int32) *AddGatewayAuthConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayAuthConsumerResponse) SetBody(v *AddGatewayAuthConsumerResponseBody) *AddGatewayAuthConsumerResponse {
	s.Body = v
	return s
}

func (s *AddGatewayAuthConsumerResponse) Validate() error {
	return dara.Validate(s)
}
