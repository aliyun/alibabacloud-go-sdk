// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *DescribeForwardTableEntriesRequest
	GetExternalIp() *string
	SetExternalPort(v string) *DescribeForwardTableEntriesRequest
	GetExternalPort() *string
	SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest
	GetForwardEntryId() *string
	SetForwardEntryName(v string) *DescribeForwardTableEntriesRequest
	GetForwardEntryName() *string
	SetForwardTableId(v string) *DescribeForwardTableEntriesRequest
	GetForwardTableId() *string
	SetInternalIp(v string) *DescribeForwardTableEntriesRequest
	GetInternalIp() *string
	SetInternalPort(v string) *DescribeForwardTableEntriesRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *DescribeForwardTableEntriesRequest
	GetIpProtocol() *string
	SetNatGatewayId(v string) *DescribeForwardTableEntriesRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *DescribeForwardTableEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeForwardTableEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeForwardTableEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeForwardTableEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeForwardTableEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeForwardTableEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeForwardTableEntriesRequest
	GetResourceOwnerId() *int64
}

type DescribeForwardTableEntriesRequest struct {
	// - If you query DNAT entries of an Internet NAT gateway, this parameter specifies the elastic IP address (EIP) that provides public network access in the DNAT entry.
	//
	// - If you query DNAT entries of a VPC NAT gateway, this parameter specifies the NAT IP address that is accessed by external networks.
	//
	// example:
	//
	// 116.28.XX.XX
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// - If you query DNAT entries of an Internet NAT gateway, this parameter specifies the external port or port range used for port forwarding.
	//
	//     - The port range must be within **1*	- to **65535**.
	//
	//     - To query a port range, separate the start and end ports with a forward slash (/), such as `10/20`.
	//
	//     - If **ExternalPort*	- is set to a port range, **InternalPort*	- must also be set to a port range with the same number of ports. For example, if **ExternalPort*	- is set to `10/20`, **InternalPort*	- must be set to `80/90`.
	//
	// - If you query DNAT entries of a VPC NAT gateway, this parameter specifies the port on the NAT IP address that is accessed by external networks. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The ID of the DNAT entry.
	//
	// example:
	//
	// fwd-8vbn3bc8roygjp0gy****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// The name of the DNAT entry.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// ForwardEntry-1
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The ID of the DNAT table.
	//
	// > You must specify at least one of **ForwardTableId*	- and **NatGatewayId**.
	//
	// example:
	//
	// ftb-bp1mbjubq34hlcqpa****
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// The private IP address.
	//
	// - If you query DNAT entries of an Internet NAT gateway, this parameter specifies the private IP address of the ECS instance that communicates with the Internet through the DNAT entry.
	//
	// - If you query DNAT entries of a VPC NAT gateway, this parameter specifies the private IP address that needs to communicate through the DNAT rule.
	//
	// example:
	//
	// 192.168.XX.XX
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// - If you query DNAT entries of an Internet NAT gateway, this parameter specifies the internal port or port range used for port forwarding. Valid values: **1*	- to **65535**.
	//
	// - If you query DNAT entries of a VPC NAT gateway, this parameter specifies the port of the destination ECS instance to be mapped. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol type. Valid values:
	//
	// - **tcp**: forwards TCP packets.
	//
	// - **udp**: forwards UDP packets.
	//
	// - **any**: forwards packets of all protocols.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The ID of the NAT gateway to query.
	//
	// > You must specify at least one of **ForwardTableId*	- and **NatGatewayId**.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s DescribeForwardTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeForwardTableEntriesRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeForwardTableEntriesRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *DescribeForwardTableEntriesRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DescribeForwardTableEntriesRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeForwardTableEntriesRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeForwardTableEntriesRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeForwardTableEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeForwardTableEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeForwardTableEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeForwardTableEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeForwardTableEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeForwardTableEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeForwardTableEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeForwardTableEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeForwardTableEntriesRequest) SetExternalIp(v string) *DescribeForwardTableEntriesRequest {
	s.ExternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetExternalPort(v string) *DescribeForwardTableEntriesRequest {
	s.ExternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardEntryName(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardTableId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardTableId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetInternalIp(v string) *DescribeForwardTableEntriesRequest {
	s.InternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetInternalPort(v string) *DescribeForwardTableEntriesRequest {
	s.InternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetIpProtocol(v string) *DescribeForwardTableEntriesRequest {
	s.IpProtocol = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetNatGatewayId(v string) *DescribeForwardTableEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetOwnerAccount(v string) *DescribeForwardTableEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetOwnerId(v int64) *DescribeForwardTableEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetPageNumber(v int32) *DescribeForwardTableEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetPageSize(v int32) *DescribeForwardTableEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetRegionId(v string) *DescribeForwardTableEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetResourceOwnerAccount(v string) *DescribeForwardTableEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetResourceOwnerId(v int64) *DescribeForwardTableEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
