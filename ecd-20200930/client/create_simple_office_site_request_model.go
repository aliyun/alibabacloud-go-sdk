// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimpleOfficeSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v string) *CreateSimpleOfficeSiteRequest
	GetAccountType() *string
	SetAuthorityHost(v string) *CreateSimpleOfficeSiteRequest
	GetAuthorityHost() *string
	SetBandwidth(v int32) *CreateSimpleOfficeSiteRequest
	GetBandwidth() *int32
	SetCenId(v string) *CreateSimpleOfficeSiteRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *CreateSimpleOfficeSiteRequest
	GetCenOwnerId() *int64
	SetCidrBlock(v string) *CreateSimpleOfficeSiteRequest
	GetCidrBlock() *string
	SetClientId(v string) *CreateSimpleOfficeSiteRequest
	GetClientId() *string
	SetClientSecret(v string) *CreateSimpleOfficeSiteRequest
	GetClientSecret() *string
	SetCloudBoxOfficeSite(v bool) *CreateSimpleOfficeSiteRequest
	GetCloudBoxOfficeSite() *bool
	SetDesktopAccessType(v string) *CreateSimpleOfficeSiteRequest
	GetDesktopAccessType() *string
	SetDomainName(v string) *CreateSimpleOfficeSiteRequest
	GetDomainName() *string
	SetEid(v string) *CreateSimpleOfficeSiteRequest
	GetEid() *string
	SetEnableAdminAccess(v bool) *CreateSimpleOfficeSiteRequest
	GetEnableAdminAccess() *bool
	SetEnableInternetAccess(v bool) *CreateSimpleOfficeSiteRequest
	GetEnableInternetAccess() *bool
	SetNeedVerifyZeroDevice(v bool) *CreateSimpleOfficeSiteRequest
	GetNeedVerifyZeroDevice() *bool
	SetOfficeSiteName(v string) *CreateSimpleOfficeSiteRequest
	GetOfficeSiteName() *string
	SetRegionId(v string) *CreateSimpleOfficeSiteRequest
	GetRegionId() *string
	SetTenantId(v string) *CreateSimpleOfficeSiteRequest
	GetTenantId() *string
	SetVSwitchId(v []*string) *CreateSimpleOfficeSiteRequest
	GetVSwitchId() []*string
	SetVerifyCode(v string) *CreateSimpleOfficeSiteRequest
	GetVerifyCode() *string
	SetVpcType(v string) *CreateSimpleOfficeSiteRequest
	GetVpcType() *string
}

type CreateSimpleOfficeSiteRequest struct {
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AuthorityHost *string `json:"AuthorityHost,omitempty" xml:"AuthorityHost,omitempty"`
	// The maximum public bandwidth. Value range: 10 to 200. Unit: Mbit/s. This parameter is available if you set `EnableInternetAccess` to `true`.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Cloud Enterprise Network (CEN) instance ID.
	//
	// >  If you want end users to connect to cloud computers from Alibaba Cloud Workspace clients over VPCs, you can attach the office network to a CEN instance. The CEN instance is the one that connects to your on-premises network over VPN Gateway or Express Connect.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account to which the Cloud Enterprise Network (CEN) instance belongs.
	//
	// - If you do not specify the CenId parameter, or the CEN instance that is specified by the CenId parameter belongs to the current Alibaba Cloud account, skip this parameter.
	//
	// - If you specify the CenId parameter and the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, enter the ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 118272523431****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block that you want the office network to use in the virtual private cloud (VPC) of the office network. The system automatically creates a VPC for the office network based on the IPv4 CIDR block. We recommend that you set this parameter to one of the following CIDR blocks and their subnets:
	//
	// 	- `10.0.0.0/12` (subnet mask range: 12 to 14 bits)
	//
	// 	- `172.16.0.0/12` (subnet mask range: 12 to 24 bits)
	//
	// 	- `192.168.0.0/16` (subnet mask range: 16 to 24 bits)
	//
	// example:
	//
	// 172.16.0.0/12
	CidrBlock    *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// Specifies whether to create a CloudBox-based office network.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	CloudBoxOfficeSite *bool `json:"CloudBoxOfficeSite,omitempty" xml:"CloudBoxOfficeSite,omitempty"`
	// The method to connect to cloud computers from Alibaba Cloud Workspace clients.
	//
	// >  The VPC connection depends on Alibaba Cloud PrivateLink. You can use PrivateLink for free. When you set this parameter to VPC or Any, PrivateLink is automatically activated.````
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Eid               *string `json:"Eid,omitempty" xml:"Eid,omitempty"`
	// Specifies whether to grant the local administrator permissions to users that are authorized to use cloud computers in the office network.
	//
	// Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable Internet access.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false (default)
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	EnableInternetAccess *bool `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	// Specifies whether to enable trusted device verification.
	//
	// example:
	//
	// true
	NeedVerifyZeroDevice *bool `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	// The office network name. The name must be 2 to 255 characters in length. It can contain digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestOfficeSite_Simple
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The IDs of the vSwitches that you want to specify in VPCs. This parameter is required only when you create CloudBox-based office networks.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the [SendVerifyCode](https://help.aliyun.com/document_detail/335132.html) operation to obtain the verification code.
	//
	// example:
	//
	// 123456
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// The network type of the office network.
	//
	// Valid values:
	//
	// 	- standard: advanced
	//
	// 	- basic: basic
	//
	// example:
	//
	// standard
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s CreateSimpleOfficeSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimpleOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateSimpleOfficeSiteRequest) GetAuthorityHost() *string {
	return s.AuthorityHost
}

func (s *CreateSimpleOfficeSiteRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateSimpleOfficeSiteRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateSimpleOfficeSiteRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *CreateSimpleOfficeSiteRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateSimpleOfficeSiteRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CreateSimpleOfficeSiteRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateSimpleOfficeSiteRequest) GetCloudBoxOfficeSite() *bool {
	return s.CloudBoxOfficeSite
}

func (s *CreateSimpleOfficeSiteRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *CreateSimpleOfficeSiteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateSimpleOfficeSiteRequest) GetEid() *string {
	return s.Eid
}

func (s *CreateSimpleOfficeSiteRequest) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *CreateSimpleOfficeSiteRequest) GetEnableInternetAccess() *bool {
	return s.EnableInternetAccess
}

func (s *CreateSimpleOfficeSiteRequest) GetNeedVerifyZeroDevice() *bool {
	return s.NeedVerifyZeroDevice
}

func (s *CreateSimpleOfficeSiteRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *CreateSimpleOfficeSiteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSimpleOfficeSiteRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateSimpleOfficeSiteRequest) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *CreateSimpleOfficeSiteRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *CreateSimpleOfficeSiteRequest) GetVpcType() *string {
	return s.VpcType
}

func (s *CreateSimpleOfficeSiteRequest) SetAccountType(v string) *CreateSimpleOfficeSiteRequest {
	s.AccountType = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetAuthorityHost(v string) *CreateSimpleOfficeSiteRequest {
	s.AuthorityHost = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetBandwidth(v int32) *CreateSimpleOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCenId(v string) *CreateSimpleOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCenOwnerId(v int64) *CreateSimpleOfficeSiteRequest {
	s.CenOwnerId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCidrBlock(v string) *CreateSimpleOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetClientId(v string) *CreateSimpleOfficeSiteRequest {
	s.ClientId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetClientSecret(v string) *CreateSimpleOfficeSiteRequest {
	s.ClientSecret = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCloudBoxOfficeSite(v bool) *CreateSimpleOfficeSiteRequest {
	s.CloudBoxOfficeSite = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetDesktopAccessType(v string) *CreateSimpleOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetDomainName(v string) *CreateSimpleOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEid(v string) *CreateSimpleOfficeSiteRequest {
	s.Eid = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetNeedVerifyZeroDevice(v bool) *CreateSimpleOfficeSiteRequest {
	s.NeedVerifyZeroDevice = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetOfficeSiteName(v string) *CreateSimpleOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetRegionId(v string) *CreateSimpleOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetTenantId(v string) *CreateSimpleOfficeSiteRequest {
	s.TenantId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetVSwitchId(v []*string) *CreateSimpleOfficeSiteRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetVerifyCode(v string) *CreateSimpleOfficeSiteRequest {
	s.VerifyCode = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetVpcType(v string) *CreateSimpleOfficeSiteRequest {
	s.VpcType = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) Validate() error {
	return dara.Validate(s)
}
