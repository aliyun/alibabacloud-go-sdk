// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADConnectorOfficeSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdHostname(v string) *ModifyADConnectorOfficeSiteRequest
	GetAdHostname() *string
	SetBackupDCHostname(v string) *ModifyADConnectorOfficeSiteRequest
	GetBackupDCHostname() *string
	SetBackupDns(v string) *ModifyADConnectorOfficeSiteRequest
	GetBackupDns() *string
	SetDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest
	GetDnsAddress() []*string
	SetDomainName(v string) *ModifyADConnectorOfficeSiteRequest
	GetDomainName() *string
	SetDomainPassword(v string) *ModifyADConnectorOfficeSiteRequest
	GetDomainPassword() *string
	SetDomainUserName(v string) *ModifyADConnectorOfficeSiteRequest
	GetDomainUserName() *string
	SetMfaEnabled(v bool) *ModifyADConnectorOfficeSiteRequest
	GetMfaEnabled() *bool
	SetOUName(v string) *ModifyADConnectorOfficeSiteRequest
	GetOUName() *string
	SetOfficeSiteId(v string) *ModifyADConnectorOfficeSiteRequest
	GetOfficeSiteId() *string
	SetOfficeSiteName(v string) *ModifyADConnectorOfficeSiteRequest
	GetOfficeSiteName() *string
	SetRegionId(v string) *ModifyADConnectorOfficeSiteRequest
	GetRegionId() *string
	SetSubDomainDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest
	GetSubDomainDnsAddress() []*string
	SetSubDomainName(v string) *ModifyADConnectorOfficeSiteRequest
	GetSubDomainName() *string
}

type ModifyADConnectorOfficeSiteRequest struct {
	// The hostname of the domain controller. The hostname must comply with the naming conventions for hostnames in Windows.
	//
	// example:
	//
	// beijing-ad01
	AdHostname *string `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	// The hostname of the secondary domain controller.
	//
	// example:
	//
	// dc002
	BackupDCHostname *string `json:"BackupDCHostname,omitempty" xml:"BackupDCHostname,omitempty"`
	// The IP address of the DNS server corresponding to the secondary domain controller.
	//
	// example:
	//
	// 192.168.2.100
	BackupDns *string `json:"BackupDns,omitempty" xml:"BackupDns,omitempty"`
	// The IP addresses of the DNS servers corresponding to the enterprise ADs. You can specify only one DNS IP address.
	//
	// example:
	//
	// 127.0.*.*
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The domain name of the enterprise AD system. You can register each domain name only once.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The password of the domain administrator. The username can be up to 64 characters in length.
	//
	// example:
	//
	// testPassword
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	// The username of the domain administrator. The username can be up to 64 characters in length.
	//
	// > Specify the value of the sAMAccountName parameter instead of the value of the userPrincipalName parameter as the username.
	//
	// example:
	//
	// Administrator
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	// Specifies whether to enable multi-factor authentication (MFA).
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
	MfaEnabled *bool `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	// The name of the organizational unit (OU) in the AD domain. You can call the [ListUserAdOrganizationUnits](https://help.aliyun.com/document_detail/311259.html) operation to obtain OUs.
	//
	// example:
	//
	// oldad.com/Domain Controllers
	OUName *string `json:"OUName,omitempty" xml:"OUName,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name. The name must be 2 to 255 characters in length. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP addresses of the DNS servers corresponding to the enterprise AD subdomains. You can specify only one DNS IP address. If you specify `SubDomainName` and leave this parameter empty, the value is the same as that of the enterprise AD domain.
	//
	// example:
	//
	// 127.0.*.*
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	// The name of the subdomain in the enterprise AD domain.
	//
	// example:
	//
	// childexample.com
	SubDomainName *string `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
}

func (s ModifyADConnectorOfficeSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteRequest) GetAdHostname() *string {
	return s.AdHostname
}

func (s *ModifyADConnectorOfficeSiteRequest) GetBackupDCHostname() *string {
	return s.BackupDCHostname
}

func (s *ModifyADConnectorOfficeSiteRequest) GetBackupDns() *string {
	return s.BackupDns
}

func (s *ModifyADConnectorOfficeSiteRequest) GetDnsAddress() []*string {
	return s.DnsAddress
}

func (s *ModifyADConnectorOfficeSiteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyADConnectorOfficeSiteRequest) GetDomainPassword() *string {
	return s.DomainPassword
}

func (s *ModifyADConnectorOfficeSiteRequest) GetDomainUserName() *string {
	return s.DomainUserName
}

func (s *ModifyADConnectorOfficeSiteRequest) GetMfaEnabled() *bool {
	return s.MfaEnabled
}

func (s *ModifyADConnectorOfficeSiteRequest) GetOUName() *string {
	return s.OUName
}

func (s *ModifyADConnectorOfficeSiteRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyADConnectorOfficeSiteRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *ModifyADConnectorOfficeSiteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyADConnectorOfficeSiteRequest) GetSubDomainDnsAddress() []*string {
	return s.SubDomainDnsAddress
}

func (s *ModifyADConnectorOfficeSiteRequest) GetSubDomainName() *string {
	return s.SubDomainName
}

func (s *ModifyADConnectorOfficeSiteRequest) SetAdHostname(v string) *ModifyADConnectorOfficeSiteRequest {
	s.AdHostname = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetBackupDCHostname(v string) *ModifyADConnectorOfficeSiteRequest {
	s.BackupDCHostname = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetBackupDns(v string) *ModifyADConnectorOfficeSiteRequest {
	s.BackupDns = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainPassword(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainUserName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *ModifyADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOUName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OUName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetRegionId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) Validate() error {
	return dara.Validate(s)
}
