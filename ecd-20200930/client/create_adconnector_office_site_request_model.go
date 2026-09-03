// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateADConnectorOfficeSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessAttribute(v string) *CreateADConnectorOfficeSiteRequest
	GetAccessAttribute() *string
	SetAdHostname(v string) *CreateADConnectorOfficeSiteRequest
	GetAdHostname() *string
	SetBackupDCHostname(v string) *CreateADConnectorOfficeSiteRequest
	GetBackupDCHostname() *string
	SetBackupDns(v string) *CreateADConnectorOfficeSiteRequest
	GetBackupDns() *string
	SetBandwidth(v int32) *CreateADConnectorOfficeSiteRequest
	GetBandwidth() *int32
	SetCenId(v string) *CreateADConnectorOfficeSiteRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *CreateADConnectorOfficeSiteRequest
	GetCenOwnerId() *int64
	SetCidrBlock(v string) *CreateADConnectorOfficeSiteRequest
	GetCidrBlock() *string
	SetDesktopAccessType(v string) *CreateADConnectorOfficeSiteRequest
	GetDesktopAccessType() *string
	SetDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest
	GetDnsAddress() []*string
	SetDomainName(v string) *CreateADConnectorOfficeSiteRequest
	GetDomainName() *string
	SetDomainPassword(v string) *CreateADConnectorOfficeSiteRequest
	GetDomainPassword() *string
	SetDomainUserName(v string) *CreateADConnectorOfficeSiteRequest
	GetDomainUserName() *string
	SetEnableAdminAccess(v bool) *CreateADConnectorOfficeSiteRequest
	GetEnableAdminAccess() *bool
	SetEnableInternetAccess(v bool) *CreateADConnectorOfficeSiteRequest
	GetEnableInternetAccess() *bool
	SetMfaEnabled(v bool) *CreateADConnectorOfficeSiteRequest
	GetMfaEnabled() *bool
	SetOfficeSiteName(v string) *CreateADConnectorOfficeSiteRequest
	GetOfficeSiteName() *string
	SetProtocolType(v string) *CreateADConnectorOfficeSiteRequest
	GetProtocolType() *string
	SetRegionId(v string) *CreateADConnectorOfficeSiteRequest
	GetRegionId() *string
	SetSpecification(v int64) *CreateADConnectorOfficeSiteRequest
	GetSpecification() *int64
	SetSubDomainDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest
	GetSubDomainDnsAddress() []*string
	SetSubDomainName(v string) *CreateADConnectorOfficeSiteRequest
	GetSubDomainName() *string
	SetVSwitchId(v []*string) *CreateADConnectorOfficeSiteRequest
	GetVSwitchId() []*string
	SetVerifyCode(v string) *CreateADConnectorOfficeSiteRequest
	GetVerifyCode() *string
}

type CreateADConnectorOfficeSiteRequest struct {
	// The access attribute of the office network (workspace).
	//
	// example:
	//
	// Private
	AccessAttribute *string `json:"AccessAttribute,omitempty" xml:"AccessAttribute,omitempty"`
	// The hostname of the domain controller. The hostname must comply with Windows hostname naming conventions.
	//
	// example:
	//
	// beijing-ad01
	AdHostname *string `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	// The hostname of the backup domain controller.
	//
	// example:
	//
	// dc002
	BackupDCHostname *string `json:"BackupDCHostname,omitempty" xml:"BackupDCHostname,omitempty"`
	// The DNS address of the backup domain controller.
	//
	// example:
	//
	// 192.168.2.100
	BackupDns *string `json:"BackupDns,omitempty" xml:"BackupDns,omitempty"`
	// The peak Internet bandwidth, in Mbit/s. Valid values: 0 to 200.
	//
	// If you do not set this parameter or set it to 0, the Internet access feature is not enabled. Settings take effect immediately.
	//
	// example:
	//
	// 1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The instance ID of the Cloud Enterprise Network (CEN).
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The Alibaba Cloud account ID of the Cloud Enterprise Network (CEN) instance owner.
	//
	// - If CenId is not specified, or the specified CenId belongs to the current Alibaba Cloud account, you do not need to specify this parameter.
	//
	// - If the specified CenId belongs to another Alibaba Cloud account, specify the Alibaba Cloud account ID of that account.
	//
	// example:
	//
	// 102681951715****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block of the office network VPC. The system uses automatic creation to provision a VPC based on the specified IPv4 CIDR block. Use one of the following CIDR blocks or their subnets as the IPv4 CIDR block:
	//
	// - `10.0.0.0/12` (valid mask range: 12 to 24 bits)
	//
	// - `172.16.0.0/12` (valid mask range: 12 to 24 bits)
	//
	// - `192.168.0.0/16` (valid mask range: 16 to 24 bits)
	//
	// example:
	//
	// 47.100.XX.XX
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The access method allowed when connecting to cloud computers.
	//
	// > The VPC connection method depends on the Alibaba Cloud PrivateLink service, which is free of charge. If this parameter is set to `VPC` or `Any`, the system automatically activates the PrivateLink service for you.
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The IP address of the DNS server corresponding to the enterprise AD. Currently, only one IP address is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The domain name of the enterprise AD. The same domain name can be registered only once.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The password of the domain administrator. The password can be up to 64 characters in length.
	//
	// example:
	//
	// testPassword
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	// The username of the domain administrator. The username can be up to 64 characters in length.
	//
	// > Use the sAMAccountName format for the username. Do not use the userPrincipalName format.
	//
	// example:
	//
	// Administrator
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	// Specifies whether to grant local administrator permissions to users who use cloud computers.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether public network access is enabled. This parameter indicates whether the feature is active.
	//
	// example:
	//
	// true
	EnableInternetAccess *bool `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	// Specifies whether to enable multi-factor authentication (MFA).
	//
	// example:
	//
	// false
	MfaEnabled *bool `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	// The name of the office network. The name must be 2 to 255 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// Default value: null.
	//
	// example:
	//
	// RD_Office_Network
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The AD Connector specification.
	//
	// example:
	//
	// 1
	Specification *int64 `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The DNS address of the enterprise AD subdomain. If `SubDomainName` is specified but this parameter is not, the subdomain DNS is considered the same as the parent domain DNS.
	//
	// example:
	//
	// 192.168.XX.XX
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	// The domain name of the enterprise AD subdomain.
	//
	// example:
	//
	// child.example.com
	SubDomainName *string `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	// The list of vSwitch IDs.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The verification code. If the specified CenId belongs to another Alibaba Cloud account, you must first call [SendVerifyCode](https://help.aliyun.com/document_detail/436847.html) to obtain the verification code.
	//
	// example:
	//
	// 12****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CreateADConnectorOfficeSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteRequest) GetAccessAttribute() *string {
	return s.AccessAttribute
}

func (s *CreateADConnectorOfficeSiteRequest) GetAdHostname() *string {
	return s.AdHostname
}

func (s *CreateADConnectorOfficeSiteRequest) GetBackupDCHostname() *string {
	return s.BackupDCHostname
}

func (s *CreateADConnectorOfficeSiteRequest) GetBackupDns() *string {
	return s.BackupDns
}

func (s *CreateADConnectorOfficeSiteRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateADConnectorOfficeSiteRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateADConnectorOfficeSiteRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *CreateADConnectorOfficeSiteRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateADConnectorOfficeSiteRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *CreateADConnectorOfficeSiteRequest) GetDnsAddress() []*string {
	return s.DnsAddress
}

func (s *CreateADConnectorOfficeSiteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateADConnectorOfficeSiteRequest) GetDomainPassword() *string {
	return s.DomainPassword
}

func (s *CreateADConnectorOfficeSiteRequest) GetDomainUserName() *string {
	return s.DomainUserName
}

func (s *CreateADConnectorOfficeSiteRequest) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *CreateADConnectorOfficeSiteRequest) GetEnableInternetAccess() *bool {
	return s.EnableInternetAccess
}

func (s *CreateADConnectorOfficeSiteRequest) GetMfaEnabled() *bool {
	return s.MfaEnabled
}

func (s *CreateADConnectorOfficeSiteRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *CreateADConnectorOfficeSiteRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateADConnectorOfficeSiteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateADConnectorOfficeSiteRequest) GetSpecification() *int64 {
	return s.Specification
}

func (s *CreateADConnectorOfficeSiteRequest) GetSubDomainDnsAddress() []*string {
	return s.SubDomainDnsAddress
}

func (s *CreateADConnectorOfficeSiteRequest) GetSubDomainName() *string {
	return s.SubDomainName
}

func (s *CreateADConnectorOfficeSiteRequest) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *CreateADConnectorOfficeSiteRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *CreateADConnectorOfficeSiteRequest) SetAccessAttribute(v string) *CreateADConnectorOfficeSiteRequest {
	s.AccessAttribute = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetAdHostname(v string) *CreateADConnectorOfficeSiteRequest {
	s.AdHostname = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetBackupDCHostname(v string) *CreateADConnectorOfficeSiteRequest {
	s.BackupDCHostname = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetBackupDns(v string) *CreateADConnectorOfficeSiteRequest {
	s.BackupDns = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetBandwidth(v int32) *CreateADConnectorOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCenId(v string) *CreateADConnectorOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCenOwnerId(v int64) *CreateADConnectorOfficeSiteRequest {
	s.CenOwnerId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCidrBlock(v string) *CreateADConnectorOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDesktopAccessType(v string) *CreateADConnectorOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainPassword(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainUserName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *CreateADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *CreateADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetProtocolType(v string) *CreateADConnectorOfficeSiteRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetRegionId(v string) *CreateADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSpecification(v int64) *CreateADConnectorOfficeSiteRequest {
	s.Specification = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetVSwitchId(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetVerifyCode(v string) *CreateADConnectorOfficeSiteRequest {
	s.VerifyCode = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) Validate() error {
	return dara.Validate(s)
}
