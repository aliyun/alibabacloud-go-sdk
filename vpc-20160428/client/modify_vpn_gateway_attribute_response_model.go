// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnGatewayAttributeResponseBody) *ModifyVpnGatewayAttributeResponse
	GetBody() *ModifyVpnGatewayAttributeResponseBody
}

type ModifyVpnGatewayAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnGatewayAttributeResponse) GetBody() *ModifyVpnGatewayAttributeResponseBody {
	return s.Body
}

func (s *ModifyVpnGatewayAttributeResponse) SetHeaders(v map[string]*string) *ModifyVpnGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnGatewayAttributeResponse) SetStatusCode(v int32) *ModifyVpnGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponse) SetBody(v *ModifyVpnGatewayAttributeResponseBody) *ModifyVpnGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
