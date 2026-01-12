// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSitesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeOfficeSitesResponseBody
	GetNextToken() *string
	SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody
	GetOfficeSites() []*DescribeOfficeSitesResponseBodyOfficeSites
	SetRequestId(v string) *DescribeOfficeSitesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOfficeSitesResponseBody
	GetTotalCount() *int32
}

type DescribeOfficeSitesResponseBody struct {
	// The token that determines the start point of the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office networks.
	OfficeSites []*DescribeOfficeSitesResponseBodyOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOfficeSitesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOfficeSitesResponseBody) GetOfficeSites() []*DescribeOfficeSitesResponseBodyOfficeSites {
	return s.OfficeSites
}

func (s *DescribeOfficeSitesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOfficeSitesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOfficeSitesResponseBody) SetNextToken(v string) *DescribeOfficeSitesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody {
	s.OfficeSites = v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetRequestId(v string) *DescribeOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetTotalCount(v int32) *DescribeOfficeSitesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) Validate() error {
	if s.OfficeSites != nil {
		for _, item := range s.OfficeSites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOfficeSitesResponseBodyOfficeSites struct {
	// Details of AD connectors.
	ADConnectors []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1astu3yrplkzoo2****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The hostname of the domain controller. The hostname must comply with the hostname naming convention of Windows.
	//
	// example:
	//
	// beijing-ad01
	AdHostname    *string `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	AuthorityHost *string `json:"AuthorityHost,omitempty" xml:"AuthorityHost,omitempty"`
	// The hostname of the secondary domain controller.
	//
	// example:
	//
	// beijing-ad02
	BackupDCHostname *string `json:"BackupDCHostname,omitempty" xml:"BackupDCHostname,omitempty"`
	// The DNS address of the secondary domain controller.
	//
	// example:
	//
	// 172.24.XX.XX
	BackupDns *string `json:"BackupDns,omitempty" xml:"BackupDns,omitempty"`
	// The maximum public bandwidth value. Valid values: 0 to 1000.\\
	//
	// If you leave this parameter empty or set this parameter to 0, Internet access is not enabled.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The CEN instance status.
	//
	// example:
	//
	// attached
	CenAttachStatus *string `json:"CenAttachStatus,omitempty" xml:"CenAttachStatus,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The IPv4 CIDR block of the VPC that the office network uses.
	//
	// example:
	//
	// 172.16.0.0/16
	CidrBlock    *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// Indicates whether the CloudBox-based office network is created.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CloudBoxOfficeSite *bool `json:"CloudBoxOfficeSite,omitempty" xml:"CloudBoxOfficeSite,omitempty"`
	// The time when the office network was created.
	//
	// example:
	//
	// 2021-05-06T05:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The custom endpoint of the access gateway.
	//
	// example:
	//
	// gw-****.com
	CustomAccessPoint *string `json:"CustomAccessPoint,omitempty" xml:"CustomAccessPoint,omitempty"`
	// The custom DNS addresses.
	CustomDnsAddress []*string `json:"CustomDnsAddress,omitempty" xml:"CustomDnsAddress,omitempty" type:"Repeated"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1ce64o4g9mdf5u****
	CustomSecurityGroupId *string `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	// The method that is used to connect cloud computers that reside in the office network from Alibaba Cloud Workspace clients.
	//
	// >  The VPC connection depends on Alibaba Cloud PrivateLink. You can use Alibaba Cloud PrivateLink for free. When you set this parameter to `VPC` or `Any`, PrivateLink is automatically activated.
	//
	// Valid values:
	//
	// 	- INTERNET (default): Cloud computers are connected from Alibaba Cloud Workspace clients over the Internet.
	//
	// 	- VPC: Cloud computers are connected from Alibaba Cloud Workspace clients over the VPC.
	//
	// 	- ANY: Cloud computers are connected from Alibaba Cloud Workspace clients over the Internet or the VPC. When end users connect to cloud computers from Alibaba Cloud Workspace clients, you can choose a connection method based on your business requirements.
	//
	// example:
	//
	// INTERNET
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The number of cloud computers that are created.
	//
	// example:
	//
	// 1
	DesktopCount *int64 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The endpoint that is used to connect to cloud computers in the directory over a VPC.
	//
	// example:
	//
	// http://ep-bp1s2vmbj55r5rzc****.epsrv-bp1pcfhpwvlpny01****.cn-hangzhou.privatelink.aliyuncs.com
	DesktopVpcEndpoint *string `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	// The DNS addresses for the AD domains.
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The username of a Domain Name System (DNS) user.
	//
	// example:
	//
	// testDnsUserName
	DnsUserName *string `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	// The domain name of the enterprise AD.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The password of the domain administrator.
	//
	// example:
	//
	// testPassword
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	// The username of the domain administrator.
	//
	// example:
	//
	// Administrator
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	Eid            *string `json:"Eid,omitempty" xml:"Eid,omitempty"`
	// Indicates whether the local administrator permissions are granted to users that are authorized to use cloud computers in the office network.
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
	// Indicates whether the connection between cloud computers in the office network is enabled. After you enable the connection between cloud computers in the office network, cloud computers in the office network can access each other.
	//
	// example:
	//
	// false
	EnableCrossDesktopAccess *bool `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// example:
	//
	// false
	EnableInternetAccess *bool `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	// Indicates whether route access control is enabled for cloud services.
	//
	// example:
	//
	// false
	EnableServiceRoute *bool   `json:"EnableServiceRoute,omitempty" xml:"EnableServiceRoute,omitempty"`
	EnvType            *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// An array of File Storage NAS (NAS) file system IDs.
	FileSystemIds []*string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Repeated"`
	IsLdap        *bool     `json:"IsLdap,omitempty" xml:"IsLdap,omitempty"`
	LdapUrl       *string   `json:"LdapUrl,omitempty" xml:"LdapUrl,omitempty"`
	// Details about registration logs.
	Logs []*DescribeOfficeSitesResponseBodyOfficeSitesLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Indicates whether multi-factor authentication (MFA) is enabled.
	//
	// example:
	//
	// false
	MfaEnabled *bool `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	// The name of the office network. The name is unique in a region.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether two-factor verification is enabled when an end user logs on to an Alibaba Cloud Workspace client. This parameter is required only for convenience office networks. If two-factor verification is enabled, the system checks whether security risks exist within the logon account when a convenience user logs on to the client. If risks are detected, the system sends a verification code to the email address that is associated with the account. Then, the convenience user can log on to the client only after the user enters the correct verification code.
	//
	// example:
	//
	// false
	NeedVerifyLoginRisk *bool `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	// Indicates whether the trusted device verification is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	NeedVerifyZeroDevice *bool `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	// The premium bandwidth plan ID.
	//
	// example:
	//
	// np-amtp8e8q1o9e4****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The network version. The new version supports App Streaming.
	//
	// Valid values:
	//
	// 	- DEFAULT: the old version.
	//
	// 	- NM: the new version.
	//
	// example:
	//
	// NM
	NmVersion *string `json:"NmVersion,omitempty" xml:"NmVersion,omitempty"`
	// The IDs of the office networks.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The account type of the office network.
	//
	// Valid values:
	//
	// 	- SIMPLE: the convenience account
	//
	// 	- AD_CONNECTOR: the enterprise AD account
	//
	// example:
	//
	// AD_CONNECTOR
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The organizational unit (OU) in the AD domain to which the office network is connected.
	//
	// example:
	//
	// example.com/Domain Controllers
	OuName *string `json:"OuName,omitempty" xml:"OuName,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX
	//
	// 	- ASP
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The IP address of the RDS license.
	//
	// example:
	//
	// 47.100.XX.XX
	RdsLicenseAddress *string `json:"RdsLicenseAddress,omitempty" xml:"RdsLicenseAddress,omitempty"`
	// The domain name of the RDS license.
	//
	// example:
	//
	// test.com
	RdsLicenseDomainName *string `json:"RdsLicenseDomainName,omitempty" xml:"RdsLicenseDomainName,omitempty"`
	// The remote desktop service (RDS) license status.
	//
	// example:
	//
	// 2
	RdsLicenseStatus *string `json:"RdsLicenseStatus,omitempty" xml:"RdsLicenseStatus,omitempty"`
	// The number of resources.
	ResourceAmounts []*DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts `json:"ResourceAmounts,omitempty" xml:"ResourceAmounts,omitempty" type:"Repeated"`
	// The security protection setting of the office network.
	//
	// Valid values:
	//
	// 	- SASE: SASE is configured.
	//
	// 	- OFF: No security protection setting is configured.
	//
	// example:
	//
	// SASE
	SecurityProtection *string `json:"SecurityProtection,omitempty" xml:"SecurityProtection,omitempty"`
	// Indicates whether single sign-on (SSO) is enabled.
	//
	// example:
	//
	// false
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// The SSO type.
	//
	// Valid values:
	//
	// 	- SAML.
	//
	// example:
	//
	// null
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	// The office network status.
	//
	// Valid values:
	//
	// 	- REGISTERING: The office network is being registered.
	//
	// 	- DEREGISTERING: The office network is being deregistered.
	//
	// 	- REGISTERED: The office network is registered.
	//
	// 	- NEEDCONFIGTRUST: A trust relationship is required for the office network.
	//
	// 	- CONFIGTRUSTFAILED: A trust relationship fails to be configured for the office network.
	//
	// 	- DEREGISTERED: The office network is deregistered.
	//
	// 	- ERROR: One or more configurations of the office network are invalid.
	//
	// 	- CONFIGTRUSTING: A trust relationship is being configured for the office network.
	//
	// 	- NEEDCONFIGUSER: Users are required for the office network.
	//
	// example:
	//
	// REGISTERED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DNS addresses for the AD subdomains.
	SubDnsAddress []*string `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" type:"Repeated"`
	// The username of enterprise AD subdomain.
	//
	// example:
	//
	// testSubDnsUserName
	SubDomainName *string `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	// The subnet mode of the office network.
	//
	// Valid values:
	//
	// 	- 0: disabled.
	//
	// 	- 1: enabled.
	//
	// example:
	//
	// 0
	SubnetMode *string `json:"SubnetMode,omitempty" xml:"SubnetMode,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The total number of cloud computers.
	//
	// example:
	//
	// 0
	TotalEdsCount *int64 `json:"TotalEdsCount,omitempty" xml:"TotalEdsCount,omitempty"`
	// The number of cloud computers in the cloud computer share.
	//
	// example:
	//
	// 0
	TotalEdsCountForGroup *int64 `json:"TotalEdsCountForGroup,omitempty" xml:"TotalEdsCountForGroup,omitempty"`
	// The number of network interface controllers (NICs).
	//
	// example:
	//
	// 1
	TotalResourceAmount *int64 `json:"TotalResourceAmount,omitempty" xml:"TotalResourceAmount,omitempty"`
	// >  This parameter is unavailable.
	//
	// example:
	//
	// null
	TrustPassword *string `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	// An array of VSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6tz5k67puge5jn8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC type.
	//
	// Valid values:
	//
	// 	- Basic
	//
	// 	- Customized
	//
	// 	- Standard
	//
	// example:
	//
	// Basic
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetADConnectors() []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	return s.ADConnectors
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetAdHostname() *string {
	return s.AdHostname
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetAuthorityHost() *string {
	return s.AuthorityHost
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetBackupDCHostname() *string {
	return s.BackupDCHostname
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetBackupDns() *string {
	return s.BackupDns
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCenAttachStatus() *string {
	return s.CenAttachStatus
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCenId() *string {
	return s.CenId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCloudBoxOfficeSite() *bool {
	return s.CloudBoxOfficeSite
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCustomAccessPoint() *string {
	return s.CustomAccessPoint
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCustomDnsAddress() []*string {
	return s.CustomDnsAddress
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetCustomSecurityGroupId() *string {
	return s.CustomSecurityGroupId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDesktopCount() *int64 {
	return s.DesktopCount
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDesktopVpcEndpoint() *string {
	return s.DesktopVpcEndpoint
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDnsAddress() []*string {
	return s.DnsAddress
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDnsUserName() *string {
	return s.DnsUserName
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDomainPassword() *string {
	return s.DomainPassword
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDomainUserName() *string {
	return s.DomainUserName
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetEid() *string {
	return s.Eid
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetEnableCrossDesktopAccess() *bool {
	return s.EnableCrossDesktopAccess
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetEnableInternetAccess() *bool {
	return s.EnableInternetAccess
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetEnableServiceRoute() *bool {
	return s.EnableServiceRoute
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetEnvType() *string {
	return s.EnvType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetFileSystemIds() []*string {
	return s.FileSystemIds
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetIsLdap() *bool {
	return s.IsLdap
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetLdapUrl() *string {
	return s.LdapUrl
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetLogs() []*DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	return s.Logs
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetMfaEnabled() *bool {
	return s.MfaEnabled
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetName() *string {
	return s.Name
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetNeedVerifyLoginRisk() *bool {
	return s.NeedVerifyLoginRisk
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetNeedVerifyZeroDevice() *bool {
	return s.NeedVerifyZeroDevice
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetNmVersion() *string {
	return s.NmVersion
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetOuName() *string {
	return s.OuName
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetRdsLicenseAddress() *string {
	return s.RdsLicenseAddress
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetRdsLicenseDomainName() *string {
	return s.RdsLicenseDomainName
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetRdsLicenseStatus() *string {
	return s.RdsLicenseStatus
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetResourceAmounts() []*DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts {
	return s.ResourceAmounts
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSecurityProtection() *string {
	return s.SecurityProtection
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSsoEnabled() *bool {
	return s.SsoEnabled
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSsoType() *string {
	return s.SsoType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetStatus() *string {
	return s.Status
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSubDnsAddress() []*string {
	return s.SubDnsAddress
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSubDomainName() *string {
	return s.SubDomainName
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSubnetMode() *string {
	return s.SubnetMode
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetTotalEdsCount() *int64 {
	return s.TotalEdsCount
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetTotalEdsCountForGroup() *int64 {
	return s.TotalEdsCountForGroup
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetTotalResourceAmount() *int64 {
	return s.TotalResourceAmount
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetTrustPassword() *string {
	return s.TrustPassword
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetADConnectors(v []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ADConnectors = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetAcceleratorId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetAccountType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.AccountType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetAdHostname(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.AdHostname = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetAuthorityHost(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.AuthorityHost = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetBackupDCHostname(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.BackupDCHostname = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetBackupDns(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.BackupDns = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetBandwidth(v int32) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Bandwidth = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCenAttachStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CenAttachStatus = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCenId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CenId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCidrBlock(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CidrBlock = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetClientId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ClientId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetClientSecret(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ClientSecret = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCloudBoxOfficeSite(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CloudBoxOfficeSite = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCreationTime(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CreationTime = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCustomAccessPoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CustomAccessPoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCustomDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CustomDnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCustomSecurityGroupId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopCount(v int64) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopCount = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDnsUserName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DnsUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainPassword(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainUserName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEid(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Eid = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableAdminAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableCrossDesktopAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableInternetAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableServiceRoute(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableServiceRoute = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnvType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnvType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetFileSystemIds(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.FileSystemIds = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetIsLdap(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.IsLdap = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetLdapUrl(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.LdapUrl = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetLogs(v []*DescribeOfficeSitesResponseBodyOfficeSitesLogs) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Logs = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetMfaEnabled(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Name = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNeedVerifyLoginRisk(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NeedVerifyLoginRisk = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNeedVerifyZeroDevice(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NeedVerifyZeroDevice = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNetworkPackageId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNmVersion(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NmVersion = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOuName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OuName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetProtocolType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ProtocolType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetRdsLicenseAddress(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.RdsLicenseAddress = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetRdsLicenseDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.RdsLicenseDomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetRdsLicenseStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.RdsLicenseStatus = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetResourceAmounts(v []*DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ResourceAmounts = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSecurityProtection(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SecurityProtection = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoEnabled(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Status = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubDomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubnetMode(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubnetMode = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTenantId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TenantId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTotalEdsCount(v int64) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TotalEdsCount = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTotalEdsCountForGroup(v int64) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TotalEdsCountForGroup = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTotalResourceAmount(v int64) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TotalResourceAmount = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTrustPassword(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TrustPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVSwitchIds(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VSwitchIds = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVpcId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VpcId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVpcType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VpcType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) Validate() error {
	if s.ADConnectors != nil {
		for _, item := range s.ADConnectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceAmounts != nil {
		for _, item := range s.ResourceAmounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOfficeSitesResponseBodyOfficeSitesADConnectors struct {
	// The connection address of the AD connector.
	//
	// example:
	//
	// 172.24.*.*
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	// The status of the AD connector.
	//
	// Valid values:
	//
	// 	- CONNECT_ERROR
	//
	// 	- RUNNING
	//
	// 	- CONNECTING (You must configure the AD domain in which the AD connector is used.)
	//
	// 	- EXPIRED
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	ConnectorStatus *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	// The ID of an elastic network interface (ENI) to which the AD connector is mounted.
	//
	// example:
	//
	// eni-bp1i4wx78lgosrj6****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The AD connector type.
	//
	// Valid values:
	//
	// 	- 1: General
	//
	// 	- 2: Advanced
	//
	// example:
	//
	// 1
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The trust password that is specified when you configure the AD trust relationship.
	//
	// example:
	//
	// password123***
	TrustKey *string `json:"TrustKey,omitempty" xml:"TrustKey,omitempty"`
	// The ID of the vSwitch that resides in the network of the AD connector.
	//
	// example:
	//
	// vsw-bp19ocz3erfx15uon****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GetADConnectorAddress() *string {
	return s.ADConnectorAddress
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GetTrustKey() *string {
	return s.TrustKey
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetADConnectorAddress(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetConnectorStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetNetworkInterfaceId(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetSpecification(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.Specification = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetTrustKey(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.TrustKey = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetVSwitchId(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) Validate() error {
	return dara.Validate(s)
}

type DescribeOfficeSitesResponseBodyOfficeSitesLogs struct {
	// The log severity.
	//
	// Valid values:
	//
	// 	- ERROR
	//
	// 	- INFO
	//
	// 	- WARN
	//
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Details of the log entry.
	//
	// example:
	//
	// code:success | message:Create Connector complete
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The step in the log entry.
	//
	// example:
	//
	// CREATE_CONNECTOR
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// The time when the log entry was printed.
	//
	// example:
	//
	// 2021-05-12T09:42Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesLogs) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) GetLevel() *string {
	return s.Level
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) GetMessage() *string {
	return s.Message
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) GetStep() *string {
	return s.Step
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetLevel(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Level = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetMessage(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Message = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetStep(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Step = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetTimeStamp(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) Validate() error {
	return dara.Validate(s)
}

type DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts struct {
	// The number of resources.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- desktop: the cloud computer.
	//
	// 	- DesktopGroup: the cloud computer share.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) GetAmount() *int64 {
	return s.Amount
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) SetAmount(v int64) *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts {
	s.Amount = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) SetResourceType(v string) *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts {
	s.ResourceType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesResourceAmounts) Validate() error {
	return dara.Validate(s)
}
