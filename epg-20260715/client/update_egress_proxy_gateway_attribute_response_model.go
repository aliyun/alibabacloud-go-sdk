// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEgressProxyGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEgressProxyGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEgressProxyGatewayAttributeResponseBody) *UpdateEgressProxyGatewayAttributeResponse
	GetBody() *UpdateEgressProxyGatewayAttributeResponseBody
}

type UpdateEgressProxyGatewayAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEgressProxyGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEgressProxyGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEgressProxyGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEgressProxyGatewayAttributeResponse) GetBody() *UpdateEgressProxyGatewayAttributeResponseBody {
	return s.Body
}

func (s *UpdateEgressProxyGatewayAttributeResponse) SetHeaders(v map[string]*string) *UpdateEgressProxyGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeResponse) SetStatusCode(v int32) *UpdateEgressProxyGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeResponse) SetBody(v *UpdateEgressProxyGatewayAttributeResponseBody) *UpdateEgressProxyGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
