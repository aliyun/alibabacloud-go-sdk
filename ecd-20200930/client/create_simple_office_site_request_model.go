// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimpleOfficeSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessAttribute(v string) *CreateSimpleOfficeSiteRequest
	GetAccessAttribute() *string
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
	// The access attribute of the office network (workspace).
	//
	// example:
	//
	// Private
	AccessAttribute *string `json:"AccessAttribute,omitempty" xml:"AccessAttribute,omitempty"`
	// The account type.
	//
	// example:
	//
	// SIMPLE
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The authority URL of the identity authentication service.
	//
	// example:
	//
	// https://login.microsoftonline.com
	AuthorityHost *string `json:"AuthorityHost,omitempty" xml:"AuthorityHost,omitempty"`
	// The peak Internet bandwidth. Valid values: 10 to 200. Unit: Mbit/s.
	//
	// You can specify this parameter when `EnableInternetAccess` is set to `true`.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The instance ID of the Cloud Enterprise Network (CEN) instance.
	//
	// > To connect to cloud desktops over a VPC connection, add the office network to a CEN instance. The CEN instance is the one that the on-premises network connects to by using a VPN or Express Connect circuit.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The Alibaba Cloud account ID to which the CEN instance belongs.
	//
	// - If CenId is not specified or the specified CEN instance belongs to the current Alibaba Cloud account, you do not need to specify this parameter.
	//
	// - If the specified CEN instance belongs to another Alibaba Cloud account, specify the Alibaba Cloud account ID of that account.
	//
	// example:
	//
	// 118272523431****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block of the VPC for the office network. This parameter is required for advanced office networks. The system uses automatic creation of a VPC based on the specified IPv4 CIDR block. Use one of the following CIDR blocks or their subnets:
	//
	// - `10.0.0.0/12` (valid mask range: 12 to 24 bits)
	//
	// - `172.16.0.0/12` (valid mask range: 12 to 24 bits)
	//
	// - `192.168.0.0/16` (valid mask range: 16 to 24 bits)
	//
	// example:
	//
	// 172.16.0.0/12
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The client ID registered with the identity provider application.
	//
	// example:
	//
	// a2c8f7e4-1b3d-4c5e-9f0a-6d7b8c9e****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret registered with the identity provider application.
	//
	// example:
	//
	// sct-9f3e2d1c****
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// Specifies whether the office network is a CloudBox office network.
	//
	// example:
	//
	// false
	CloudBoxOfficeSite *bool `json:"CloudBoxOfficeSite,omitempty" xml:"CloudBoxOfficeSite,omitempty"`
	// The access method allowed when connecting to cloud desktops.
	//
	// > The VPC connection method depends on the Alibaba Cloud PrivateLink service, which is free of charge. If this parameter is set to `VPC` or `Any`, the system automatically activates the PrivateLink service.
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The domain name of the enterprise AD.
	//
	// example:
	//
	// domain.local
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The enterprise ID (EID).
	//
	// example:
	//
	// e-1234abcd****
	Eid *string `json:"Eid,omitempty" xml:"Eid,omitempty"`
	// Specifies whether to grant local administrator permissions to users who use cloud desktops.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable public network access.
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
	// The name of the office network. The name must be 2 to 255 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestOfficeSite_Simple
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID of the identity provider.
	//
	// example:
	//
	// 72f988bf-86f1-41af-91ab-2d7cd011****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the vSwitch in the VPC. This parameter is required when you create a CloudBox office network.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The verification code. If the specified CEN instance belongs to another Alibaba Cloud account, call [SendVerifyCode](https://help.aliyun.com/document_detail/335132.html) to obtain the verification code first.
	//
	// example:
	//
	// 123456
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// The type of the office network.
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

func (s *CreateSimpleOfficeSiteRequest) GetAccessAttribute() *string {
	return s.AccessAttribute
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

func (s *CreateSimpleOfficeSiteRequest) SetAccessAttribute(v string) *CreateSimpleOfficeSiteRequest {
	s.AccessAttribute = &v
	return s
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
