// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayPolicyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayPolicyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayPolicyConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayPolicyConfigResponseBody) *UpdateGatewayPolicyConfigResponse
	GetBody() *UpdateGatewayPolicyConfigResponseBody
}

type UpdateGatewayPolicyConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayPolicyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayPolicyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayPolicyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayPolicyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayPolicyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayPolicyConfigResponse) GetBody() *UpdateGatewayPolicyConfigResponseBody {
	return s.Body
}

func (s *UpdateGatewayPolicyConfigResponse) SetHeaders(v map[string]*string) *UpdateGatewayPolicyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayPolicyConfigResponse) SetStatusCode(v int32) *UpdateGatewayPolicyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayPolicyConfigResponse) SetBody(v *UpdateGatewayPolicyConfigResponseBody) *UpdateGatewayPolicyConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayPolicyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
