// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerResourceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayAuthConsumerResourceStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayAuthConsumerResourceStatusResponseBody) *UpdateGatewayAuthConsumerResourceStatusResponse
	GetBody() *UpdateGatewayAuthConsumerResourceStatusResponseBody
}

type UpdateGatewayAuthConsumerResourceStatusResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayAuthConsumerResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayAuthConsumerResourceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) GetBody() *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	return s.Body
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) SetHeaders(v map[string]*string) *UpdateGatewayAuthConsumerResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) SetStatusCode(v int32) *UpdateGatewayAuthConsumerResourceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) SetBody(v *UpdateGatewayAuthConsumerResourceStatusResponseBody) *UpdateGatewayAuthConsumerResourceStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
