// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateForwardEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateForwardEntryRequest
	GetDryRun() *bool
	SetExternalIp(v string) *CreateForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *CreateForwardEntryRequest
	GetExternalPort() *string
	SetForwardEntryName(v string) *CreateForwardEntryRequest
	GetForwardEntryName() *string
	SetForwardTableId(v string) *CreateForwardEntryRequest
	GetForwardTableId() *string
	SetInternalIp(v string) *CreateForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *CreateForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *CreateForwardEntryRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *CreateForwardEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateForwardEntryRequest
	GetOwnerId() *int64
	SetPortBreak(v bool) *CreateForwardEntryRequest
	GetPortBreak() *bool
	SetRegionId(v string) *CreateForwardEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateForwardEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateForwardEntryRequest
	GetResourceOwnerId() *int64
}

type CreateForwardEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this request. Valid values:
	//
	// - **true**: performs a dry run without creating the DNAT entry. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the request passes the dry run, a 2xx HTTP status code is returned and the DNAT entry is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// - If you configure a DNAT entry for an Internet NAT gateway, specify an elastic IP address (EIP) for public network access.
	//
	// - If you configure a DNAT entry for a VPC NAT gateway, specify a NAT IP address for external network access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 116.28.XX.XX
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// - If you configure a DNAT entry for an Internet NAT gateway, specify the external port or port range for port forwarding.
	//
	//
	//
	//     - The port range must be from **1*	- to **65535**.
	//
	//     - To forward traffic within a port range, separate the start port and the end port with a forward slash (/). Example: `10/20`.
	//
	//     - If **ExternalPort*	- is set to a port range, **InternalPort*	- must also be set to a port range with the same number of ports. For example, if **ExternalPort*	- is set to `10/20`, **InternalPort*	- must be set to `80/90`.
	//
	// - If you configure a DNAT entry for a VPC NAT gateway, specify the port on the NAT IP address that is accessed by the external network. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The name of the DNAT rule.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ForwardEntry-1
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The ID of the DNAT table.
	//
	// This parameter is required.
	//
	// example:
	//
	// ftb-bp1mbjubq34hlcqpa****
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// - If you configure a DNAT entry for an Internet NAT gateway, specify the private IP address of the ECS instance that requires Internet communication. The private IP address must meet the following requirements:
	//
	//
	//
	//     - The private IP address must belong to the CIDR block of the VPC where the NAT gateway is deployed.
	//
	//     - The private IP address must be used by an ECS instance, and the ECS instance must not be associated with an EIP. Otherwise, the DNAT entry does not take effect.
	//
	// - If you configure a DNAT entry for a VPC NAT gateway, specify the private IP address that requires communication through the DNAT rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// - If you configure a DNAT entry for an Internet NAT gateway, specify the internal port or port range for port forwarding. Valid values: **1*	- to **65535**.
	//
	// - If you configure a DNAT entry for a VPC NAT gateway, specify the destination port of the ECS instance to be mapped. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol type. Valid values:
	//
	// - **TCP**: forwards TCP packets.
	//
	// - **UDP**: forwards UDP packets.
	//
	// - **Any**: forwards packets of all protocols. If **IpProtocol*	- is set to **Any**, **ExternalPort*	- and **InternalPort*	- must also be set to **Any*	- to implement DNAT IP mapping.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable port breaking. Valid values:
	//
	// - **true**: enables port breaking.
	//
	// - **false*	- (default): disables port breaking.
	//
	// > If a DNAT entry and an SNAT entry use the same public IP address and you need to configure a port number greater than 1024, you must set **PortBreak*	- to **true**.
	//
	// example:
	//
	// false
	PortBreak *bool `json:"PortBreak,omitempty" xml:"PortBreak,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateForwardEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *CreateForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *CreateForwardEntryRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *CreateForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *CreateForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *CreateForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *CreateForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateForwardEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateForwardEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateForwardEntryRequest) GetPortBreak() *bool {
	return s.PortBreak
}

func (s *CreateForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateForwardEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateForwardEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateForwardEntryRequest) SetClientToken(v string) *CreateForwardEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateForwardEntryRequest) SetDryRun(v bool) *CreateForwardEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateForwardEntryRequest) SetExternalIp(v string) *CreateForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetExternalPort(v string) *CreateForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetForwardEntryName(v string) *CreateForwardEntryRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *CreateForwardEntryRequest) SetForwardTableId(v string) *CreateForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalIp(v string) *CreateForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalPort(v string) *CreateForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetIpProtocol(v string) *CreateForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateForwardEntryRequest) SetOwnerAccount(v string) *CreateForwardEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateForwardEntryRequest) SetOwnerId(v int64) *CreateForwardEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetPortBreak(v bool) *CreateForwardEntryRequest {
	s.PortBreak = &v
	return s
}

func (s *CreateForwardEntryRequest) SetRegionId(v string) *CreateForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetResourceOwnerAccount(v string) *CreateForwardEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateForwardEntryRequest) SetResourceOwnerId(v int64) *CreateForwardEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
