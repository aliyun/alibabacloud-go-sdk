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
	AccessAttribute *string `json:"AccessAttribute,omitempty" xml:"AccessAttribute,omitempty"`
	// The domain controller hostname.
	//
	// The hostname must comply with Windows hostname naming conventions.
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
	// The peak public bandwidth, specified in Mbit/s. The value can range from 0 to 200.<br>
	//
	// If you omit this parameter or set it to 0, internet access is disabled.<br>
	//
	// example:
	//
	// 1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the Cloud Enterprise Network (CEN) instance.
	//
	// - If you do not specify `CenId`, or the specified CEN instance belongs to your Alibaba Cloud account, you do not need to specify this parameter.
	//
	// - If the specified CEN instance belongs to another Alibaba Cloud account, you must specify that account\\"s ID.
	//
	// example:
	//
	// 102681951715****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block for the office site\\"s VPC. The system uses this IPv4 CIDR block to automatically create a VPC. We recommend that you use one of the following CIDR blocks or their subnets:
	//
	// - `10.0.0.0/12` (The subnet mask length must be 12 to 24 bits.)
	//
	// - `172.16.0.0/12` (The subnet mask length must be 12 to 24 bits.)
	//
	// - `192.168.0.0/16` (The subnet mask length must be 16 to 24 bits.)
	//
	// example:
	//
	// 47.100.XX.XX
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The method for connecting to cloud desktops.
	//
	// > VPC connections are established using Alibaba Cloud PrivateLink, which is a free service. If you set this parameter to `VPC` or `Any`, PrivateLink is automatically enabled.
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// An array that contains the IP address of the DNS server for the enterprise AD. You can specify only one IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The domain name for the enterprise AD. Each domain name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain administrator\\"s password. The password cannot exceed 64 characters in length.
	//
	// example:
	//
	// testPassword
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	// The domain administrator\\"s username. The username cannot exceed 64 characters in length.
	//
	// > Use the sAMAccountName, not the userPrincipalName.
	//
	// example:
	//
	// Administrator
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	// Specifies whether to grant local administrator permissions to cloud desktop users. Default: true.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// This parameter is deprecated. Use the `Bandwidth` parameter to manage internet access.
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
	// The name of the office site. The name must be 2 to 255 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).<br>
	//
	// This parameter is empty by default.<br>
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The AD Connector type.
	//
	// example:
	//
	// 1
	Specification *int64 `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The DNS address of the enterprise AD child domain. If you specify `SubDomainName` but not this parameter, the DNS address of the child domain is considered the same as that of the parent domain.
	//
	// example:
	//
	// 192.168.XX.XX
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	// The domain name of the enterprise AD child domain.
	//
	// example:
	//
	// child.example.com
	SubDomainName *string `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	// The list of vSwitch IDs.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The verification code. If the `CenId` that you specify belongs to another Alibaba Cloud account, you must first call the [SendVerifyCode](https://help.aliyun.com/document_detail/436847.html) operation to obtain the verification code.
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
