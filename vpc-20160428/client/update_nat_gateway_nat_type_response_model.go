// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNatGatewayNatTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNatGatewayNatTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNatGatewayNatTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNatGatewayNatTypeResponseBody) *UpdateNatGatewayNatTypeResponse
	GetBody() *UpdateNatGatewayNatTypeResponseBody
}

type UpdateNatGatewayNatTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNatGatewayNatTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNatGatewayNatTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNatGatewayNatTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateNatGatewayNatTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNatGatewayNatTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNatGatewayNatTypeResponse) GetBody() *UpdateNatGatewayNatTypeResponseBody {
	return s.Body
}

func (s *UpdateNatGatewayNatTypeResponse) SetHeaders(v map[string]*string) *UpdateNatGatewayNatTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateNatGatewayNatTypeResponse) SetStatusCode(v int32) *UpdateNatGatewayNatTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNatGatewayNatTypeResponse) SetBody(v *UpdateNatGatewayNatTypeResponseBody) *UpdateNatGatewayNatTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateNatGatewayNatTypeResponse) Validate() error {
	return dara.Validate(s)
}
