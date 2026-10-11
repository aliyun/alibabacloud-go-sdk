// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpAddressesToEgressProxyIpPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddIpAddressesToEgressProxyIpPoolResponseBody
	GetRequestId() *string
}

type AddIpAddressesToEgressProxyIpPoolResponseBody struct {
	// example:
	//
	// C79118A9-CA6D-5359-B53E-111BE8A03DCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpAddressesToEgressProxyIpPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIpAddressesToEgressProxyIpPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpAddressesToEgressProxyIpPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIpAddressesToEgressProxyIpPoolResponseBody) SetRequestId(v string) *AddIpAddressesToEgressProxyIpPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
