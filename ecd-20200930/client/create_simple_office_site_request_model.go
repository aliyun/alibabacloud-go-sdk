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
	AccessAttribute *string `json:"AccessAttribute,omitempty" xml:"AccessAttribute,omitempty"`
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AuthorityHost   *string `json:"AuthorityHost,omitempty" xml:"AuthorityHost,omitempty"`
	// The peak public bandwidth. Valid values: 10 to 200. Unit: Mbps.
	//
	// This parameter is valid only when `EnableInternetAccess` is set to `true`.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// > If you want to connect to cloud desktops over a VPC, attach the office site to the same CEN instance that is connected to your on-premises network by a VPN or an Express Connect circuit.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the CEN instance.
	//
	// - If you do not specify CenId, or if the CEN instance belongs to your Alibaba Cloud account, this parameter is not required.
	//
	// - If the CEN instance is owned by another Alibaba Cloud account, specify the ID of that account.
	//
	// example:
	//
	// 118272523431****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block for the office site\\"s Virtual Private Cloud (VPC). This parameter is required for standard office sites. The system automatically creates a VPC based on the specified IPv4 CIDR block. Use one of the following CIDR blocks or their subnets:
	//
	// - `10.0.0.0/12` (The valid mask range is 12 to 24 bits.)
	//
	// - `172.16.0.0/12` (The valid mask range is 12 to 24 bits.)
	//
	// - `192.168.0.0/16` (The valid mask range is 16 to 24 bits.)
	//
	// example:
	//
	// 172.16.0.0/12
	CidrBlock    *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// Specifies whether to create a Cloud Box office site.
	//
	// example:
	//
	// false
	CloudBoxOfficeSite *bool `json:"CloudBoxOfficeSite,omitempty" xml:"CloudBoxOfficeSite,omitempty"`
	// Specifies how clients can connect to cloud desktops.
	//
	// > VPC connections rely on the Alibaba Cloud PrivateLink service, which is free of charge. If you set this parameter to `VPC` or `Any`, the system automatically enables the PrivateLink service.
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Eid               *string `json:"Eid,omitempty" xml:"Eid,omitempty"`
	// Specifies whether to grant users local administrator privileges on their cloud desktops.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable internet access.
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
	// The name of the office site. The name must be 2 to 255 characters in length. It must start with a letter or a Chinese character, and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// TestOfficeSite_Simple
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to get a list of regions that support Elastic Desktop Service (ECD).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The vSwitch ID. This parameter is required when you create a Cloud Box office site.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The verification code. If the CEN instance is owned by another Alibaba Cloud account, you must first call [SendVerifyCode](https://help.aliyun.com/document_detail/335132.html) to obtain a verification code.
	//
	// example:
	//
	// 123456
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// The type of the office site.
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
