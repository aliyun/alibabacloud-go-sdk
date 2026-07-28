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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating a DNAT entry. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the check succeeds, a 2xx HTTP status code is returned and the DNAT entry is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// - If you add a DNAT entry for an Internet NAT gateway, this parameter specifies the elastic IP address (EIP) that provides public network access.
	//
	// - If you add a DNAT entry for a VPC NAT gateway, this parameter specifies the NAT IP address that provides external network access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 116.28.XX.XX
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// - If you add a DNAT entry for an Internet NAT gateway, this parameter specifies the external port or port range for port forwarding.
	//
	//
	//
	//     - Valid port values: **1*	- to **65535**.
	//
	//     - To specify a port range, separate the start and end ports with a forward slash (/), such as `10/20`.
	//
	//     - If **ExternalPort*	- is set to a port range, **InternalPort*	- must also be set to a port range with the same number of ports. For example, if **ExternalPort*	- is set to `10/20`, **InternalPort*	- can be set to `80/90`.
	//
	// - If you add a DNAT entry for a VPC NAT gateway, this parameter specifies the port on the NAT IP address that is accessed by the external network. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The name of the DNAT rule.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It cannot start with `http://` or `https://`.
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
	// - If you add a DNAT entry for an Internet NAT gateway, this parameter specifies the private IP address of the ECS instance that needs to communicate over the Internet. The private IP address must meet the following conditions:
	//
	//
	//
	//     - It must belong to the CIDR block of the VPC in which the NAT gateway resides.
	//
	//     - The DNAT entry takes effect only when the IP address is used by an ECS instance that is not associated with an EIP.
	//
	// - If you add a DNAT entry for a VPC NAT gateway, this parameter specifies the private IP address that communicates through the DNAT rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// - If you add a DNAT entry for an Internet NAT gateway, this parameter specifies the internal port or port range for port forwarding. Valid values: **1*	- to **65535**.
	//
	// - If you add a DNAT entry for a VPC NAT gateway, this parameter specifies the destination port of the ECS instance to which traffic is mapped. Valid values: **1*	- to **65535**.
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
	// > If a DNAT entry and an SNAT entry use the same public IP address, and you want to configure a port number greater than 1024, set **PortBreak*	- to **true**.
	//
	// example:
	//
	// false
	PortBreak *bool `json:"PortBreak,omitempty" xml:"PortBreak,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
