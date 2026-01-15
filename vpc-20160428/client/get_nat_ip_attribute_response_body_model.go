// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatIpAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *GetNatIpAttributeResponseBody
	GetCreationTime() *string
	SetNatGatewayId(v string) *GetNatIpAttributeResponseBody
	GetNatGatewayId() *string
	SetNatIp(v string) *GetNatIpAttributeResponseBody
	GetNatIp() *string
	SetNatIpCidr(v string) *GetNatIpAttributeResponseBody
	GetNatIpCidr() *string
	SetNatIpDescription(v string) *GetNatIpAttributeResponseBody
	GetNatIpDescription() *string
	SetNatIpId(v string) *GetNatIpAttributeResponseBody
	GetNatIpId() *string
	SetNatIpName(v string) *GetNatIpAttributeResponseBody
	GetNatIpName() *string
	SetNatIpStatus(v string) *GetNatIpAttributeResponseBody
	GetNatIpStatus() *string
	SetRequestId(v string) *GetNatIpAttributeResponseBody
	GetRequestId() *string
}

type GetNatIpAttributeResponseBody struct {
	// The creation time of the queried NAT IP address.
	//
	// example:
	//
	// 2021-07-16T16:47Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the VPC NAT gateway instance to which the queried NAT IP address belongs.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The queried NAT IP address.
	//
	// example:
	//
	// 192.168.0.34
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The address range where the queried NAT IP address is located.
	//
	// example:
	//
	// 192.168.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// Description of the queried NAT IP address.
	//
	// example:
	//
	// 123
	NatIpDescription *string `json:"NatIpDescription,omitempty" xml:"NatIpDescription,omitempty"`
	// The ID of the queried NAT IP address instance.
	//
	// example:
	//
	// vpcnatip-gw8y7q3cpk3fggs87****
	NatIpId *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	// Name of the queried NAT IP address.
	//
	// example:
	//
	// test
	NatIpName *string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty"`
	// The status of the queried NAT IP address. Values:
	//
	// - **Available**: Available.
	//
	//  - **Deleting**: Deleting.
	//
	//  - **Creating**: Creating.
	//
	// example:
	//
	// Available
	NatIpStatus *string `json:"NatIpStatus,omitempty" xml:"NatIpStatus,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 9D55DB90-0D57-46AB-841D-1D4238514AC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNatIpAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNatIpAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNatIpAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetNatIpAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatIpAttributeResponseBody) GetNatIp() *string {
	return s.NatIp
}

func (s *GetNatIpAttributeResponseBody) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *GetNatIpAttributeResponseBody) GetNatIpDescription() *string {
	return s.NatIpDescription
}

func (s *GetNatIpAttributeResponseBody) GetNatIpId() *string {
	return s.NatIpId
}

func (s *GetNatIpAttributeResponseBody) GetNatIpName() *string {
	return s.NatIpName
}

func (s *GetNatIpAttributeResponseBody) GetNatIpStatus() *string {
	return s.NatIpStatus
}

func (s *GetNatIpAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNatIpAttributeResponseBody) SetCreationTime(v string) *GetNatIpAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatGatewayId(v string) *GetNatIpAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatIp(v string) *GetNatIpAttributeResponseBody {
	s.NatIp = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatIpCidr(v string) *GetNatIpAttributeResponseBody {
	s.NatIpCidr = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatIpDescription(v string) *GetNatIpAttributeResponseBody {
	s.NatIpDescription = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatIpId(v string) *GetNatIpAttributeResponseBody {
	s.NatIpId = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatIpName(v string) *GetNatIpAttributeResponseBody {
	s.NatIpName = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetNatIpStatus(v string) *GetNatIpAttributeResponseBody {
	s.NatIpStatus = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) SetRequestId(v string) *GetNatIpAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNatIpAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
