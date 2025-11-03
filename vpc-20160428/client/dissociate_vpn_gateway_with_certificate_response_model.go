// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateVpnGatewayWithCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateVpnGatewayWithCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateVpnGatewayWithCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DissociateVpnGatewayWithCertificateResponseBody) *DissociateVpnGatewayWithCertificateResponse
	GetBody() *DissociateVpnGatewayWithCertificateResponseBody
}

type DissociateVpnGatewayWithCertificateResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateVpnGatewayWithCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateVpnGatewayWithCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateVpnGatewayWithCertificateResponse) GoString() string {
	return s.String()
}

func (s *DissociateVpnGatewayWithCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateVpnGatewayWithCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateVpnGatewayWithCertificateResponse) GetBody() *DissociateVpnGatewayWithCertificateResponseBody {
	return s.Body
}

func (s *DissociateVpnGatewayWithCertificateResponse) SetHeaders(v map[string]*string) *DissociateVpnGatewayWithCertificateResponse {
	s.Headers = v
	return s
}

func (s *DissociateVpnGatewayWithCertificateResponse) SetStatusCode(v int32) *DissociateVpnGatewayWithCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateResponse) SetBody(v *DissociateVpnGatewayWithCertificateResponseBody) *DissociateVpnGatewayWithCertificateResponse {
	s.Body = v
	return s
}

func (s *DissociateVpnGatewayWithCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
