// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpnGatewayWithCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateVpnGatewayWithCertificateResponseBody
	GetRequestId() *string
}

type AssociateVpnGatewayWithCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateVpnGatewayWithCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpnGatewayWithCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateVpnGatewayWithCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateVpnGatewayWithCertificateResponseBody) SetRequestId(v string) *AssociateVpnGatewayWithCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
