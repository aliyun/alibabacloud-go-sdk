// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayAuthConsumerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayAuthConsumerResponseBody) *UpdateGatewayAuthConsumerResponse
	GetBody() *UpdateGatewayAuthConsumerResponseBody
}

type UpdateGatewayAuthConsumerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayAuthConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayAuthConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayAuthConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayAuthConsumerResponse) GetBody() *UpdateGatewayAuthConsumerResponseBody {
	return s.Body
}

func (s *UpdateGatewayAuthConsumerResponse) SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayAuthConsumerResponse) SetStatusCode(v int32) *UpdateGatewayAuthConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponse) SetBody(v *UpdateGatewayAuthConsumerResponseBody) *UpdateGatewayAuthConsumerResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayAuthConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
