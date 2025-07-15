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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 	- The EIP that can be accessed over the Internet when you configure a DNAT entry for an Internet NAT gateway.
	//
	// 	- The NAT IP address that can be accessed by external networks when you configure a DNAT entry for a VPC NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// 116.28.XX.XX
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// 	- The external port range that is used for port forwarding when you configure a DNAT entry for an Internet NAT gateway.
	//
	//     	- Valid values: **1*	- to **65535**.
	//
	//     	- To specify a port range, separate the first port and the last port with a forward slash (/), for example, `10/20`.
	//
	//     	- If you set **ExternalPort*	- to a port range, you must also set **InternalPort*	- to a port range, and the number of ports specified by these parameters must be the same. For example, if you set **ExternalPort*	- to `10/20`, you can set **InternalPort*	- to `80/90`.
	//
	// 	- The port that can be accessed by external networks when you configure a DNAT entry for a VPC NAT gateway. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The name of the DNAT entry.
	//
	// The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
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
	// 	- The private IP address of the ECS instance that needs to communicate with the Internet when you configure a DNAT entry for an Internet NAT gateway. The private IP address must meet the following requirements:
	//
	//     	- It must belong to the CIDR block of the VPC where the NAT gateway is deployed.
	//
	//     	- The DNAT entry takes effect only if the private IP address is assigned to an ECS instance and the ECS instance is not associated with an EIP.
	//
	// 	- The private IP address that uses DNAT when you add a DNAT entry to a VPC NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// 	- The internal port or port range that is used for port forwarding when you configure a DNAT entry for an Internet NAT gateway. Valid values: **1*	- to **65535**.
	//
	// 	- The port of the destination ECS instance to be mapped when you configure a DNAT entry for a VPC NAT gateway. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **Any*	- If you set **IpProtocol*	- to **Any**, you must also set **ExternalPort*	- and **InternalPort*	- to **Any*	- to implement DNAT IP mapping.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to remove limits on the port range. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  If a DNAT entry and an SNAT entry have the same public IP address, ou must specify a port that is larger that 1024, and set **PortBreak*	- to **true**.
	//
	// example:
	//
	// false
	PortBreak *bool `json:"PortBreak,omitempty" xml:"PortBreak,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
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
