// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayAuthConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayAuthConsumerResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayAuthConsumerResponseBody) *ListGatewayAuthConsumerResponse
	GetBody() *ListGatewayAuthConsumerResponseBody
}

type ListGatewayAuthConsumerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayAuthConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayAuthConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayAuthConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayAuthConsumerResponse) GetBody() *ListGatewayAuthConsumerResponseBody {
	return s.Body
}

func (s *ListGatewayAuthConsumerResponse) SetHeaders(v map[string]*string) *ListGatewayAuthConsumerResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayAuthConsumerResponse) SetStatusCode(v int32) *ListGatewayAuthConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResponse) SetBody(v *ListGatewayAuthConsumerResponseBody) *ListGatewayAuthConsumerResponse {
	s.Body = v
	return s
}

func (s *ListGatewayAuthConsumerResponse) Validate() error {
	return dara.Validate(s)
}
