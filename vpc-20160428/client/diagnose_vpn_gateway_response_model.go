// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiagnoseVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiagnoseVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DiagnoseVpnGatewayResponseBody) *DiagnoseVpnGatewayResponse
	GetBody() *DiagnoseVpnGatewayResponseBody
}

type DiagnoseVpnGatewayResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiagnoseVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiagnoseVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiagnoseVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiagnoseVpnGatewayResponse) GetBody() *DiagnoseVpnGatewayResponseBody {
	return s.Body
}

func (s *DiagnoseVpnGatewayResponse) SetHeaders(v map[string]*string) *DiagnoseVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *DiagnoseVpnGatewayResponse) SetStatusCode(v int32) *DiagnoseVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DiagnoseVpnGatewayResponse) SetBody(v *DiagnoseVpnGatewayResponseBody) *DiagnoseVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *DiagnoseVpnGatewayResponse) Validate() error {
	return dara.Validate(s)
}
