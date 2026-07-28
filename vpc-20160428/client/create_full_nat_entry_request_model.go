// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFullNatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDomain(v string) *CreateFullNatEntryRequest
	GetAccessDomain() *string
	SetAccessIp(v string) *CreateFullNatEntryRequest
	GetAccessIp() *string
	SetAccessPort(v string) *CreateFullNatEntryRequest
	GetAccessPort() *string
	SetClientToken(v string) *CreateFullNatEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateFullNatEntryRequest
	GetDryRun() *bool
	SetFullNatEntryDescription(v string) *CreateFullNatEntryRequest
	GetFullNatEntryDescription() *string
	SetFullNatEntryName(v string) *CreateFullNatEntryRequest
	GetFullNatEntryName() *string
	SetFullNatTableId(v string) *CreateFullNatEntryRequest
	GetFullNatTableId() *string
	SetIpProtocol(v string) *CreateFullNatEntryRequest
	GetIpProtocol() *string
	SetNatIp(v string) *CreateFullNatEntryRequest
	GetNatIp() *string
	SetNatIpPort(v string) *CreateFullNatEntryRequest
	GetNatIpPort() *string
	SetNetworkInterfaceId(v string) *CreateFullNatEntryRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *CreateFullNatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateFullNatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateFullNatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateFullNatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFullNatEntryRequest
	GetResourceOwnerId() *int64
}

type CreateFullNatEntryRequest struct {
	// The backend domain name for FULLNAT address translation.
	//
	// example:
	//
	// xxx.com
	AccessDomain *string `json:"AccessDomain,omitempty" xml:"AccessDomain,omitempty"`
	// The backend IP address for FULLNAT address translation.
	//
	// example:
	//
	// 192.168.XX.XX
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The backend port for FULLNAT port mapping. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a dry run request. The system checks the request for potential issues, including missing required parameters, invalid parameter values, and whether the Resource Access Management (RAM) user is granted the required authorization. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the FULLNAT entry is added.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the FULLNAT entry.
	//
	// The description can be empty or 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	FullNatEntryDescription *string `json:"FullNatEntryDescription,omitempty" xml:"FullNatEntryDescription,omitempty"`
	// The name of the FULLNAT entry. The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It cannot start with http:// or https://.
	//
	// example:
	//
	// test
	FullNatEntryName *string `json:"FullNatEntryName,omitempty" xml:"FullNatEntryName,omitempty"`
	// The ID of the FULLNAT table to which the FULLNAT entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// fulltb-gw88z7hhlv43rmb26****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
	// The protocol type of the Redirection Port. Valid values:
	//
	// -  **TCP**: forwards TCP packets.
	//
	// -  **UDP**: forwards UDP packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The NAT IP address that provides address translation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The frontend port for FULLNAT port mapping. Valid values: **1*	- to **65535**.
	//
	// > If you do not specify this parameter, the system randomly assigns an available port.
	//
	// example:
	//
	// 80
	NatIpPort *string `json:"NatIpPort,omitempty" xml:"NatIpPort,omitempty"`
	// The ID of the network interface controller (NIC).
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC NAT gateway to which the FULLNAT entry belongs.
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

func (s CreateFullNatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFullNatEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateFullNatEntryRequest) GetAccessDomain() *string {
	return s.AccessDomain
}

func (s *CreateFullNatEntryRequest) GetAccessIp() *string {
	return s.AccessIp
}

func (s *CreateFullNatEntryRequest) GetAccessPort() *string {
	return s.AccessPort
}

func (s *CreateFullNatEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFullNatEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFullNatEntryRequest) GetFullNatEntryDescription() *string {
	return s.FullNatEntryDescription
}

func (s *CreateFullNatEntryRequest) GetFullNatEntryName() *string {
	return s.FullNatEntryName
}

func (s *CreateFullNatEntryRequest) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *CreateFullNatEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateFullNatEntryRequest) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateFullNatEntryRequest) GetNatIpPort() *string {
	return s.NatIpPort
}

func (s *CreateFullNatEntryRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *CreateFullNatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateFullNatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFullNatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFullNatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFullNatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFullNatEntryRequest) SetAccessDomain(v string) *CreateFullNatEntryRequest {
	s.AccessDomain = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetAccessIp(v string) *CreateFullNatEntryRequest {
	s.AccessIp = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetAccessPort(v string) *CreateFullNatEntryRequest {
	s.AccessPort = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetClientToken(v string) *CreateFullNatEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetDryRun(v bool) *CreateFullNatEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetFullNatEntryDescription(v string) *CreateFullNatEntryRequest {
	s.FullNatEntryDescription = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetFullNatEntryName(v string) *CreateFullNatEntryRequest {
	s.FullNatEntryName = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetFullNatTableId(v string) *CreateFullNatEntryRequest {
	s.FullNatTableId = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetIpProtocol(v string) *CreateFullNatEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetNatIp(v string) *CreateFullNatEntryRequest {
	s.NatIp = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetNatIpPort(v string) *CreateFullNatEntryRequest {
	s.NatIpPort = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetNetworkInterfaceId(v string) *CreateFullNatEntryRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetOwnerAccount(v string) *CreateFullNatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetOwnerId(v int64) *CreateFullNatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetRegionId(v string) *CreateFullNatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetResourceOwnerAccount(v string) *CreateFullNatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFullNatEntryRequest) SetResourceOwnerId(v int64) *CreateFullNatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFullNatEntryRequest) Validate() error {
	return dara.Validate(s)
}
