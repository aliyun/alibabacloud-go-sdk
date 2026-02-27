// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullNatEntryAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDomain(v string) *ModifyFullNatEntryAttributeRequest
	GetAccessDomain() *string
	SetAccessIp(v string) *ModifyFullNatEntryAttributeRequest
	GetAccessIp() *string
	SetAccessPort(v string) *ModifyFullNatEntryAttributeRequest
	GetAccessPort() *string
	SetClientToken(v string) *ModifyFullNatEntryAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyFullNatEntryAttributeRequest
	GetDryRun() *bool
	SetFullNatEntryDescription(v string) *ModifyFullNatEntryAttributeRequest
	GetFullNatEntryDescription() *string
	SetFullNatEntryId(v string) *ModifyFullNatEntryAttributeRequest
	GetFullNatEntryId() *string
	SetFullNatEntryName(v string) *ModifyFullNatEntryAttributeRequest
	GetFullNatEntryName() *string
	SetFullNatTableId(v string) *ModifyFullNatEntryAttributeRequest
	GetFullNatTableId() *string
	SetIpProtocol(v string) *ModifyFullNatEntryAttributeRequest
	GetIpProtocol() *string
	SetNatIp(v string) *ModifyFullNatEntryAttributeRequest
	GetNatIp() *string
	SetNatIpPort(v string) *ModifyFullNatEntryAttributeRequest
	GetNatIpPort() *string
	SetNetworkInterfaceId(v string) *ModifyFullNatEntryAttributeRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *ModifyFullNatEntryAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyFullNatEntryAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyFullNatEntryAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyFullNatEntryAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFullNatEntryAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyFullNatEntryAttributeRequest struct {
	// The backend domain name of the FULLNAT address translation that needs to be modified.
	//
	// example:
	//
	// xxx.com
	AccessDomain *string `json:"AccessDomain,omitempty" xml:"AccessDomain,omitempty"`
	// The backend IP address to be modified in FULLNAT address translation.
	//
	// example:
	//
	// 192.168.XX.XX
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The backend port to be modified in FULLNAT port mapping. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description of the FULLNAT entry.
	//
	// You can leave this parameter empty or enter a description. If you enter a description, the description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abcd
	FullNatEntryDescription *string `json:"FullNatEntryDescription,omitempty" xml:"FullNatEntryDescription,omitempty"`
	// The ID of the FULLNAT entry to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// fullnat-gw8fz23jezpbblf1j****
	FullNatEntryId *string `json:"FullNatEntryId,omitempty" xml:"FullNatEntryId,omitempty"`
	// The new name of the FULLNAT entry.
	//
	// The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// modify
	FullNatEntryName *string `json:"FullNatEntryName,omitempty" xml:"FullNatEntryName,omitempty"`
	// The ID of the FULLNAT table to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// fulltb-gw88z7hhlv43rmb26****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
	// The protocol of the packets that are forwarded by the port. Valid values:
	//
	// 	- **TCP**: TCP
	//
	// 	- **UDP**
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The NAT IP address to be modified.
	//
	// example:
	//
	// 192.168.XX.XX
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The frontend port to be modified in FULLNAT port mapping. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	NatIpPort *string `json:"NatIpPort,omitempty" xml:"NatIpPort,omitempty"`
	// The ID of the elastic network interface (ENI) to be modified.
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Virtual Private Cloud (VPC) NAT gateway to which the FULLNAT entry to be modified belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyFullNatEntryAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullNatEntryAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullNatEntryAttributeRequest) GetAccessDomain() *string {
	return s.AccessDomain
}

func (s *ModifyFullNatEntryAttributeRequest) GetAccessIp() *string {
	return s.AccessIp
}

func (s *ModifyFullNatEntryAttributeRequest) GetAccessPort() *string {
	return s.AccessPort
}

func (s *ModifyFullNatEntryAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyFullNatEntryAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyFullNatEntryAttributeRequest) GetFullNatEntryDescription() *string {
	return s.FullNatEntryDescription
}

func (s *ModifyFullNatEntryAttributeRequest) GetFullNatEntryId() *string {
	return s.FullNatEntryId
}

func (s *ModifyFullNatEntryAttributeRequest) GetFullNatEntryName() *string {
	return s.FullNatEntryName
}

func (s *ModifyFullNatEntryAttributeRequest) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *ModifyFullNatEntryAttributeRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyFullNatEntryAttributeRequest) GetNatIp() *string {
	return s.NatIp
}

func (s *ModifyFullNatEntryAttributeRequest) GetNatIpPort() *string {
	return s.NatIpPort
}

func (s *ModifyFullNatEntryAttributeRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ModifyFullNatEntryAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyFullNatEntryAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFullNatEntryAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyFullNatEntryAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFullNatEntryAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFullNatEntryAttributeRequest) SetAccessDomain(v string) *ModifyFullNatEntryAttributeRequest {
	s.AccessDomain = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetAccessIp(v string) *ModifyFullNatEntryAttributeRequest {
	s.AccessIp = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetAccessPort(v string) *ModifyFullNatEntryAttributeRequest {
	s.AccessPort = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetClientToken(v string) *ModifyFullNatEntryAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetDryRun(v bool) *ModifyFullNatEntryAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetFullNatEntryDescription(v string) *ModifyFullNatEntryAttributeRequest {
	s.FullNatEntryDescription = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetFullNatEntryId(v string) *ModifyFullNatEntryAttributeRequest {
	s.FullNatEntryId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetFullNatEntryName(v string) *ModifyFullNatEntryAttributeRequest {
	s.FullNatEntryName = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetFullNatTableId(v string) *ModifyFullNatEntryAttributeRequest {
	s.FullNatTableId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetIpProtocol(v string) *ModifyFullNatEntryAttributeRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetNatIp(v string) *ModifyFullNatEntryAttributeRequest {
	s.NatIp = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetNatIpPort(v string) *ModifyFullNatEntryAttributeRequest {
	s.NatIpPort = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetNetworkInterfaceId(v string) *ModifyFullNatEntryAttributeRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetOwnerAccount(v string) *ModifyFullNatEntryAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetOwnerId(v int64) *ModifyFullNatEntryAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetRegionId(v string) *ModifyFullNatEntryAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetResourceOwnerAccount(v string) *ModifyFullNatEntryAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) SetResourceOwnerId(v int64) *ModifyFullNatEntryAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeRequest) Validate() error {
	return dara.Validate(s)
}
