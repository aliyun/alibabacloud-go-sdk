// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateVpnGatewayWithCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateVpnGatewayWithCertificateResponseBody
	GetRequestId() *string
}

type DissociateVpnGatewayWithCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 611CB80C-B6A9-43DB-9E38-0B0AC3D9B58F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateVpnGatewayWithCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateVpnGatewayWithCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateVpnGatewayWithCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateVpnGatewayWithCertificateResponseBody) SetRequestId(v string) *DissociateVpnGatewayWithCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
