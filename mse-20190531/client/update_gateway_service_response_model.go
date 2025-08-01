// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayServiceResponseBody) *UpdateGatewayServiceResponse
	GetBody() *UpdateGatewayServiceResponseBody
}

type UpdateGatewayServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayServiceResponse) GetBody() *UpdateGatewayServiceResponseBody {
	return s.Body
}

func (s *UpdateGatewayServiceResponse) SetHeaders(v map[string]*string) *UpdateGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayServiceResponse) SetStatusCode(v int32) *UpdateGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayServiceResponse) SetBody(v *UpdateGatewayServiceResponseBody) *UpdateGatewayServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayServiceResponse) Validate() error {
	return dara.Validate(s)
}
