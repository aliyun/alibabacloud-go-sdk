// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateADConnectorOfficeSiteRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The hostname of the domain controller. The hostname must comply with the naming conventions for Windows hosts.
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
	// The maximum public bandwidth of the Internet access package. Valid values: 0 to 200.\\
	//
	// If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
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
	// The Alibaba Cloud account that creates the Cloud Enterprise Network (CEN) instance.
	//
	// 	- If you do not specify the CenId parameter, or the CEN instance that is specified by the CenId parameter belongs to the current Alibaba Cloud account, skip this parameter.
	//
	// 	- If you specify the CenId parameter and the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, enter the ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 102681951715****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block of the virtual private cloud (VPC) that your office network uses. The system creates a VPC for your office network based on the IPv4 CIDR block. We recommend that you set this parameter to one of the following CIDR blocks and their subnets:
	//
	// 	- `10.0.0.0/12` (subnet mask range: 12 to 24 bits)
	//
	// 	- `172.16.0.0/12` (subnet mask range: 12 to 24 bits)
	//
	// 	- `192.168.0.0/16` (subnet mask range: 16 to 24 bits)
	//
	// example:
	//
	// 47.100.XX.XX
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The method to connect to cloud computers from Alibaba Cloud Workspace clients.
	//
	// >  The VPC connection depends on Alibaba Cloud PrivateLink. You can use PrivateLink for free. When you set this parameter to `VPC` or `Any`, PrivateLink is automatically activated.
	//
	// Valid values:
	//
	// - Internet: connects clients to cloud desktops only over the Internet. [Default]
	//
	// - VPC: connects clients to cloud desktops only over a VPC.
	//
	// - Any: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method based on your business requirements when you connect to your cloud desktop from a client.
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The IP address of the DNS server of the enterprise AD system. You can specify only one IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The domain name of the enterprise AD system. You can register each domain name only once.
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
	// > Specify the username by using sAMAccountName instead of userPrincipalName.
	//
	// example:
	//
	// Administrator
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	// Specifies whether to grant the local administrator permissions to users that are authorized to use cloud computers in the office network.
	//
	// Valid values:
	//
	// 	- <!-- -->
	//
	//     true
	//
	//     <!-- -->
	//
	//     (default)
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable Internet access.
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
	// The office network name. The name must be 2 to 255 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.\\
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The protocol type.
	//
	// Valid value:
	//
	// 	- Adaptive Streaming Protocol (ASP)
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The AD connector type.
	//
	// Valid values:
	//
	// 	- 1: General
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- 2: Advanced
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// 1
	Specification *int64 `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The DNS address of the enterprise AD subdomain. If you specify `SubDomainName` but do not specify this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
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
	// The array of the vSwitch IDs.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the [SendVerifyCode](https://help.aliyun.com/document_detail/436847.html) operation to obtain the verification code.
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
