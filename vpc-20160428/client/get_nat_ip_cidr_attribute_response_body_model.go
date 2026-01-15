// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatIpCidrAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *GetNatIpCidrAttributeResponseBody
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *GetNatIpCidrAttributeResponseBody
	GetNatIpCidr() *string
	SetNatIpCidrDescription(v string) *GetNatIpCidrAttributeResponseBody
	GetNatIpCidrDescription() *string
	SetNatIpCidrId(v string) *GetNatIpCidrAttributeResponseBody
	GetNatIpCidrId() *string
	SetNatIpCidrName(v string) *GetNatIpCidrAttributeResponseBody
	GetNatIpCidrName() *string
	SetNatIpCidrStatus(v string) *GetNatIpCidrAttributeResponseBody
	GetNatIpCidrStatus() *string
	SetRequestId(v string) *GetNatIpCidrAttributeResponseBody
	GetRequestId() *string
}

type GetNatIpCidrAttributeResponseBody struct {
	// The ID of the VPC NAT gateway instance to which the queried NAT IP address range belongs.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The queried NAT IP address range.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// Description of the queried NAT IP address range.
	//
	// example:
	//
	// default
	NatIpCidrDescription *string `json:"NatIpCidrDescription,omitempty" xml:"NatIpCidrDescription,omitempty"`
	// The instance ID of the queried NAT IP address range.
	//
	// example:
	//
	// vpcnatcidr-gw8lhqtvdn4qnea****
	NatIpCidrId *string `json:"NatIpCidrId,omitempty" xml:"NatIpCidrId,omitempty"`
	// The name of the queried NAT IP address range.
	//
	// example:
	//
	// default
	NatIpCidrName *string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty"`
	// The status of the queried NAT IP address segment. Values:
	//
	// - Available: Available status.
	//
	// - Deleting: In the process of being deleted.
	//
	// - Creating: In the process of being created.
	//
	// example:
	//
	// Available
	NatIpCidrStatus *string `json:"NatIpCidrStatus,omitempty" xml:"NatIpCidrStatus,omitempty"`
	// Request ID.
	//
	// example:
	//
	// E9AD97A0-5338-43F8-8A80-5E274CCBA11B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNatIpCidrAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNatIpCidrAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNatIpCidrAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatIpCidrAttributeResponseBody) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *GetNatIpCidrAttributeResponseBody) GetNatIpCidrDescription() *string {
	return s.NatIpCidrDescription
}

func (s *GetNatIpCidrAttributeResponseBody) GetNatIpCidrId() *string {
	return s.NatIpCidrId
}

func (s *GetNatIpCidrAttributeResponseBody) GetNatIpCidrName() *string {
	return s.NatIpCidrName
}

func (s *GetNatIpCidrAttributeResponseBody) GetNatIpCidrStatus() *string {
	return s.NatIpCidrStatus
}

func (s *GetNatIpCidrAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNatIpCidrAttributeResponseBody) SetNatGatewayId(v string) *GetNatIpCidrAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) SetNatIpCidr(v string) *GetNatIpCidrAttributeResponseBody {
	s.NatIpCidr = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) SetNatIpCidrDescription(v string) *GetNatIpCidrAttributeResponseBody {
	s.NatIpCidrDescription = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) SetNatIpCidrId(v string) *GetNatIpCidrAttributeResponseBody {
	s.NatIpCidrId = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) SetNatIpCidrName(v string) *GetNatIpCidrAttributeResponseBody {
	s.NatIpCidrName = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) SetNatIpCidrStatus(v string) *GetNatIpCidrAttributeResponseBody {
	s.NatIpCidrStatus = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) SetRequestId(v string) *GetNatIpCidrAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNatIpCidrAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
