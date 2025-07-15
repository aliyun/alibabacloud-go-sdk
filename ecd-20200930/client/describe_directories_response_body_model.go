// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdHostname(v string) *DescribeDirectoriesResponseBody
	GetAdHostname() *string
	SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody
	GetDirectories() []*DescribeDirectoriesResponseBodyDirectories
	SetNextToken(v string) *DescribeDirectoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDirectoriesResponseBody
	GetRequestId() *string
}

type DescribeDirectoriesResponseBody struct {
	// The hostname of the domain controller. The hostname must comply with the hostname naming convention of Windows. This parameter is returned only when the directory type is AD office network.
	//
	// example:
	//
	// cnshsv21hmc****
	AdHostname *string `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	// The directories.
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The token that is used for the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F369A091-002F-49C8-AD55-02A776297C7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) GetAdHostname() *string {
	return s.AdHostname
}

func (s *DescribeDirectoriesResponseBody) GetDirectories() []*DescribeDirectoriesResponseBodyDirectories {
	return s.Directories
}

func (s *DescribeDirectoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDirectoriesResponseBody) SetAdHostname(v string) *DescribeDirectoriesResponseBody {
	s.AdHostname = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetNextToken(v string) *DescribeDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDirectoriesResponseBodyDirectories struct {
	// Details of the AD connector.
	ADConnectors []*DescribeDirectoriesResponseBodyDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	// The hostname of the domain controller.
	//
	// example:
	//
	// dc001
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
	// The time when the directory was created.
	//
	// example:
	//
	// 2020-11-02T01:44Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The security group ID. This parameter is returned only when the directory type is AD office network.
	//
	// example:
	//
	// sg-bp1ce64o4g9mdf5u****
	CustomSecurityGroupId *string `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	// The method in which the cloud computer is connected.
	//
	// Valid values:
	//
	// 	- VPC
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Internet
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Any
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The endpoint that is used to connect to cloud computers in the directory over a VPC.
	//
	// example:
	//
	// http://ep-bp1s2vmbj55r5rzc****.epsrv-bp1pcfhpwvlpny01****.cn-hangzhou.privatelink.aliyuncs.com
	DesktopVpcEndpoint *string `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The directory type.
	//
	// Valid values:
	//
	// 	- AD_CONNECTOR: AD directory
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- RAM: RAM directory
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// RAM
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The DNS address of the directory.
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The username of a DNS user.
	//
	// example:
	//
	// testDnsUserName
	DnsUserName *string `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The password of the domain administrator. This parameter is returned only when the directory type is AD office network.
	//
	// example:
	//
	// testPassword
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	// The username of the domain administrator.
	//
	// example:
	//
	// sAMAccountName
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	// Indicates whether the local administrator permissions are granted to users that use cloud computers in the office network.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Indicates whether cloud computers can communicate with each other in the directory.
	//
	// example:
	//
	// true
	EnableCrossDesktopAccess *bool `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	// Indicates whether access over the Internet is enabled.
	//
	// >  This parameter is unavailable.
	//
	// example:
	//
	// false
	EnableInternetAccess *bool `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	// The IDs of File Storage NAS (NAS) file systems.
	FileSystemIds []*string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Repeated"`
	// The registration logs. This parameter is returned only when the directory type is AD office network.
	Logs []*DescribeDirectoriesResponseBodyDirectoriesLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Indicates whether MFA is enabled.
	//
	// example:
	//
	// false
	MfaEnabled *bool `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	// The directory name.
	//
	// example:
	//
	// testDirectoryName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether two-step verification for logons is enabled. This parameter is returned only for directories of convenience account type.\\
	//
	// If two-factor verification is enabled, the system checks whether security risks exist within the logon account when a convenience user logs on to an Alibaba Cloud Workspace client. If risks are detected, the system sends a verification code to the email address that is associated with the account. Then, the convenience user can log on to the client only after the user enters the correct verification code.
	//
	// example:
	//
	// false
	NeedVerifyLoginRisk *bool `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	// The organization unit that you selected when you added the cloud computer to the domain.
	//
	// example:
	//
	// example.com/Domain Controllers
	OuName *string `json:"OuName,omitempty" xml:"OuName,omitempty"`
	// Indicates whether single sign-on (SSO) is enabled.
	//
	// example:
	//
	// false
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// The status of the AD directory.
	//
	// Valid values:
	//
	// 	- REGISTERING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- REGISTERED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// REGISTERING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DNS address of the enterprise AD subdomain.
	SubDnsAddress []*string `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" type:"Repeated"`
	// The fully qualified domain name (FQDN) of the existing AD subdomain. The value contains both the host name and the domain name.
	//
	// example:
	//
	// child.example.com
	SubDomainName *string `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	// The AD trust password. This parameter is returned only when the directory type is AD office network.
	//
	// example:
	//
	// 82Tg****
	TrustPassword *string `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	// The IDs of the vSwitches specified when the directory was created.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC to which the vSwitch belongs. This parameter is returned only when the directory type is AD office network.
	//
	// example:
	//
	// vpc-uf6tz5k67puge5jn8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetADConnectors() []*DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	return s.ADConnectors
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetAdHostname() *string {
	return s.AdHostname
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetBackupDCHostname() *string {
	return s.BackupDCHostname
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetBackupDns() *string {
	return s.BackupDns
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetCustomSecurityGroupId() *string {
	return s.CustomSecurityGroupId
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDesktopVpcEndpoint() *string {
	return s.DesktopVpcEndpoint
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDnsAddress() []*string {
	return s.DnsAddress
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDnsUserName() *string {
	return s.DnsUserName
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDomainPassword() *string {
	return s.DomainPassword
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDomainUserName() *string {
	return s.DomainUserName
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetEnableCrossDesktopAccess() *bool {
	return s.EnableCrossDesktopAccess
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetEnableInternetAccess() *bool {
	return s.EnableInternetAccess
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetFileSystemIds() []*string {
	return s.FileSystemIds
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetLogs() []*DescribeDirectoriesResponseBodyDirectoriesLogs {
	return s.Logs
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetMfaEnabled() *bool {
	return s.MfaEnabled
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetName() *string {
	return s.Name
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetNeedVerifyLoginRisk() *bool {
	return s.NeedVerifyLoginRisk
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetOuName() *string {
	return s.OuName
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetSsoEnabled() *bool {
	return s.SsoEnabled
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetStatus() *string {
	return s.Status
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetSubDnsAddress() []*string {
	return s.SubDnsAddress
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetSubDomainName() *string {
	return s.SubDomainName
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetTrustPassword() *string {
	return s.TrustPassword
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetADConnectors(v []*DescribeDirectoriesResponseBodyDirectoriesADConnectors) *DescribeDirectoriesResponseBodyDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetAdHostname(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.AdHostname = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetBackupDCHostname(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.BackupDCHostname = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetBackupDns(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.BackupDns = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCreationTime(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CreationTime = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopVpcEndpoint(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableAdminAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableCrossDesktopAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableInternetAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetFileSystemIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.FileSystemIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetLogs(v []*DescribeDirectoriesResponseBodyDirectoriesLogs) *DescribeDirectoriesResponseBodyDirectories {
	s.Logs = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetMfaEnabled(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetNeedVerifyLoginRisk(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.NeedVerifyLoginRisk = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetOuName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.OuName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSsoEnabled(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetStatus(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSubDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSubDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.SubDomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.TrustPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVSwitchIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVpcId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.VpcId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) Validate() error {
	return dara.Validate(s)
}

type DescribeDirectoriesResponseBodyDirectoriesADConnectors struct {
	// The connection address.
	//
	// example:
	//
	// ``172.17.**.**``
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	// Valid values:
	//
	// 	- CONNECT_ERROR
	//
	// 	- RUNNING
	//
	// 	- CONNECTING: You must configure domain trust for your AD system.
	//
	// 	- EXPIRED
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	ConnectorStatus *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	// The ID of the NIC to which the AD connector is mounted.
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
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The trust password of the AD domain controller.
	//
	// example:
	//
	// yfpoAD****
	TrustKey *string `json:"TrustKey,omitempty" xml:"TrustKey,omitempty"`
	// The ID of the vSwitch with which the AD connector is associated.
	//
	// example:
	//
	// vsw-bp19ocz3erfx15uon****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) GetADConnectorAddress() *string {
	return s.ADConnectorAddress
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) GetTrustKey() *string {
	return s.TrustKey
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetConnectorStatus(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetNetworkInterfaceId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetSpecification(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.Specification = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetTrustKey(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.TrustKey = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetVSwitchId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) Validate() error {
	return dara.Validate(s)
}

type DescribeDirectoriesResponseBodyDirectoriesLogs struct {
	// The level of the log entry.
	//
	// Valid values:
	//
	// 	- ERROR
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- INFO
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- WARN
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Details of the log entry.
	//
	// example:
	//
	// code:success | message:Create Connector complete.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The step that corresponds to the log entry.
	//
	// example:
	//
	// DescribeDirectories
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// The time when the log entry was printed.
	//
	// example:
	//
	// 2021-01-22T06:45Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesLogs) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) GetLevel() *string {
	return s.Level
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) GetMessage() *string {
	return s.Message
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) GetStep() *string {
	return s.Step
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetLevel(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Level = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetMessage(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Message = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetStep(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Step = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetTimeStamp(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) Validate() error {
	return dara.Validate(s)
}
