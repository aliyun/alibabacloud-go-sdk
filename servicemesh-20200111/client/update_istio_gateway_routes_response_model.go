// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioGatewayRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIstioGatewayRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIstioGatewayRoutesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIstioGatewayRoutesResponseBody) *UpdateIstioGatewayRoutesResponse
	GetBody() *UpdateIstioGatewayRoutesResponseBody
}

type UpdateIstioGatewayRoutesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIstioGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIstioGatewayRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIstioGatewayRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIstioGatewayRoutesResponse) GetBody() *UpdateIstioGatewayRoutesResponseBody {
	return s.Body
}

func (s *UpdateIstioGatewayRoutesResponse) SetHeaders(v map[string]*string) *UpdateIstioGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *UpdateIstioGatewayRoutesResponse) SetStatusCode(v int32) *UpdateIstioGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesResponse) SetBody(v *UpdateIstioGatewayRoutesResponseBody) *UpdateIstioGatewayRoutesResponse {
	s.Body = v
	return s
}

func (s *UpdateIstioGatewayRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
