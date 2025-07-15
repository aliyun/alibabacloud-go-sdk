// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpnGatewayDiagnoseResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpnGatewayDiagnoseResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpnGatewayDiagnoseResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVpnGatewayDiagnoseResultResponseBody) *GetVpnGatewayDiagnoseResultResponse
	GetBody() *GetVpnGatewayDiagnoseResultResponseBody
}

type GetVpnGatewayDiagnoseResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpnGatewayDiagnoseResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpnGatewayDiagnoseResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpnGatewayDiagnoseResultResponse) GoString() string {
	return s.String()
}

func (s *GetVpnGatewayDiagnoseResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpnGatewayDiagnoseResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpnGatewayDiagnoseResultResponse) GetBody() *GetVpnGatewayDiagnoseResultResponseBody {
	return s.Body
}

func (s *GetVpnGatewayDiagnoseResultResponse) SetHeaders(v map[string]*string) *GetVpnGatewayDiagnoseResultResponse {
	s.Headers = v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponse) SetStatusCode(v int32) *GetVpnGatewayDiagnoseResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponse) SetBody(v *GetVpnGatewayDiagnoseResultResponseBody) *GetVpnGatewayDiagnoseResultResponse {
	s.Body = v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponse) Validate() error {
	return dara.Validate(s)
}
