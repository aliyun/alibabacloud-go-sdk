// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceTrafficPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayServiceTrafficPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayServiceTrafficPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayServiceTrafficPolicyResponseBody) *UpdateGatewayServiceTrafficPolicyResponse
	GetBody() *UpdateGatewayServiceTrafficPolicyResponseBody
}

type UpdateGatewayServiceTrafficPolicyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayServiceTrafficPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayServiceTrafficPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceTrafficPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) GetBody() *UpdateGatewayServiceTrafficPolicyResponseBody {
	return s.Body
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) SetHeaders(v map[string]*string) *UpdateGatewayServiceTrafficPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) SetStatusCode(v int32) *UpdateGatewayServiceTrafficPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) SetBody(v *UpdateGatewayServiceTrafficPolicyResponseBody) *UpdateGatewayServiceTrafficPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
