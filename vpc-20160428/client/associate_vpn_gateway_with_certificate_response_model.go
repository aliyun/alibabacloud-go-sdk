// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpnGatewayWithCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateVpnGatewayWithCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateVpnGatewayWithCertificateResponse
	GetStatusCode() *int32
	SetBody(v *AssociateVpnGatewayWithCertificateResponseBody) *AssociateVpnGatewayWithCertificateResponse
	GetBody() *AssociateVpnGatewayWithCertificateResponseBody
}

type AssociateVpnGatewayWithCertificateResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateVpnGatewayWithCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateVpnGatewayWithCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpnGatewayWithCertificateResponse) GoString() string {
	return s.String()
}

func (s *AssociateVpnGatewayWithCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateVpnGatewayWithCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateVpnGatewayWithCertificateResponse) GetBody() *AssociateVpnGatewayWithCertificateResponseBody {
	return s.Body
}

func (s *AssociateVpnGatewayWithCertificateResponse) SetHeaders(v map[string]*string) *AssociateVpnGatewayWithCertificateResponse {
	s.Headers = v
	return s
}

func (s *AssociateVpnGatewayWithCertificateResponse) SetStatusCode(v int32) *AssociateVpnGatewayWithCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateResponse) SetBody(v *AssociateVpnGatewayWithCertificateResponseBody) *AssociateVpnGatewayWithCertificateResponse {
	s.Body = v
	return s
}

func (s *AssociateVpnGatewayWithCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
