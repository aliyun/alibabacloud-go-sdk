// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyForwardEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyForwardEntryRequest
	GetDryRun() *bool
	SetExternalIp(v string) *ModifyForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *ModifyForwardEntryRequest
	GetExternalPort() *string
	SetForwardEntryId(v string) *ModifyForwardEntryRequest
	GetForwardEntryId() *string
	SetForwardEntryName(v string) *ModifyForwardEntryRequest
	GetForwardEntryName() *string
	SetForwardTableId(v string) *ModifyForwardEntryRequest
	GetForwardTableId() *string
	SetInternalIp(v string) *ModifyForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *ModifyForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *ModifyForwardEntryRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *ModifyForwardEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyForwardEntryRequest
	GetOwnerId() *int64
	SetPortBreak(v bool) *ModifyForwardEntryRequest
	GetPortBreak() *bool
	SetRegionId(v string) *ModifyForwardEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyForwardEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyForwardEntryRequest
	GetResourceOwnerId() *int64
}

type ModifyForwardEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the DNAT entry. The system checks the required parameters, request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the DNAT entry is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// - If you modify a DNAT entry of an Internet NAT gateway, this parameter specifies the public IP address used to provide public network access.
	//
	// - If you modify a DNAT entry of a VPC NAT gateway, this parameter specifies the NAT IP address accessed by the external network.
	//
	// example:
	//
	// 116.85.XX.XX
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// - If you modify a DNAT entry of an Internet NAT gateway, this parameter specifies the external port or port range used for port forwarding in the DNAT entry.
	//
	//     - The port range must be within **1*	- to **65535**.
	//
	//     - To specify a port range, separate the start and end ports with a forward slash (/), such as `10/20`.
	//
	//     - If you modify both **ExternalPort*	- and **InternalPort**, and **ExternalPort*	- is set to a port range, **InternalPort*	- must also be set to a port range with the same number of ports. For example, if **ExternalPort*	- is set to `10/20`, **InternalPort*	- must be set to `80/90`.
	//
	// - If you modify a DNAT entry of a VPC NAT gateway, this parameter specifies the port accessed by the external network. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The ID of the DNAT entry to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// fwd-8vbn3bc8roygjp0gy****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// The new name of the DNAT entry.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The ID of the DNAT table to which the DNAT entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ftb-8vbx8xu2lqj9qb334****
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// - If you modify a DNAT entry of an Internet NAT gateway, this parameter specifies the private IP address of the ECS instance that communicates with the Internet through the DNAT entry.
	//
	// - If you modify a DNAT entry of a VPC NAT gateway, this parameter specifies the private IP address that needs to communicate through the DNAT rule.
	//
	// example:
	//
	// 10.0.0.78
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// - If you modify a DNAT entry of an Internet NAT gateway, this parameter specifies the internal port or port range used for port forwarding in the DNAT entry. Valid values: **1*	- to **65535**.
	//
	// - If you modify a DNAT entry of a VPC NAT gateway, this parameter specifies the port of the destination ECS instance to be mapped. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol type. Valid values:
	//
	//
	//
	// - **TCP**: forwards TCP packets.
	//
	//
	//
	// - **UDP**: forwards UDP packets.
	//
	//
	//
	// - **Any**: forwards packets of all protocols.
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
	// - **false**: does not enable port breaking. If a DNAT entry and an SNAT entry use the same public IP address and you want to configure a port number greater than `1024`, set `PortBreak` to `true`.
	//
	// example:
	//
	// false
	PortBreak *bool `json:"PortBreak,omitempty" xml:"PortBreak,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
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

func (s ModifyForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyForwardEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyForwardEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *ModifyForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ModifyForwardEntryRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *ModifyForwardEntryRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *ModifyForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *ModifyForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *ModifyForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *ModifyForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyForwardEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyForwardEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyForwardEntryRequest) GetPortBreak() *bool {
	return s.PortBreak
}

func (s *ModifyForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyForwardEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyForwardEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyForwardEntryRequest) SetClientToken(v string) *ModifyForwardEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetDryRun(v bool) *ModifyForwardEntryRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetExternalIp(v string) *ModifyForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetExternalPort(v string) *ModifyForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardEntryId(v string) *ModifyForwardEntryRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardEntryName(v string) *ModifyForwardEntryRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardTableId(v string) *ModifyForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetInternalIp(v string) *ModifyForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetInternalPort(v string) *ModifyForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetIpProtocol(v string) *ModifyForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetOwnerAccount(v string) *ModifyForwardEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetOwnerId(v int64) *ModifyForwardEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetPortBreak(v bool) *ModifyForwardEntryRequest {
	s.PortBreak = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetRegionId(v string) *ModifyForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetResourceOwnerAccount(v string) *ModifyForwardEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetResourceOwnerId(v int64) *ModifyForwardEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
