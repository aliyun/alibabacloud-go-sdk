// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityHost(v string) *ModifyOfficeSiteAttributeRequest
	GetAuthorityHost() *string
	SetClientId(v string) *ModifyOfficeSiteAttributeRequest
	GetClientId() *string
	SetClientSecret(v string) *ModifyOfficeSiteAttributeRequest
	GetClientSecret() *string
	SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest
	GetDesktopAccessType() *string
	SetDomainName(v string) *ModifyOfficeSiteAttributeRequest
	GetDomainName() *string
	SetEnableAdminAccess(v bool) *ModifyOfficeSiteAttributeRequest
	GetEnableAdminAccess() *bool
	SetNeedVerifyLoginRisk(v bool) *ModifyOfficeSiteAttributeRequest
	GetNeedVerifyLoginRisk() *bool
	SetNeedVerifyZeroDevice(v bool) *ModifyOfficeSiteAttributeRequest
	GetNeedVerifyZeroDevice() *bool
	SetOfficeSiteId(v string) *ModifyOfficeSiteAttributeRequest
	GetOfficeSiteId() *string
	SetOfficeSiteName(v string) *ModifyOfficeSiteAttributeRequest
	GetOfficeSiteName() *string
	SetRegionId(v string) *ModifyOfficeSiteAttributeRequest
	GetRegionId() *string
	SetTenantId(v string) *ModifyOfficeSiteAttributeRequest
	GetTenantId() *string
	SetVSwitchId(v []*string) *ModifyOfficeSiteAttributeRequest
	GetVSwitchId() []*string
}

type ModifyOfficeSiteAttributeRequest struct {
	AuthorityHost *string `json:"AuthorityHost,omitempty" xml:"AuthorityHost,omitempty"`
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret  *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The method for connecting to cloud computers.
	//
	// > VPC connections use Alibaba Cloud PrivateLink, a free service. If you set this parameter to VPC or Any, PrivateLink is automatically activated.
	//
	// example:
	//
	// INTERNET
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to grant cloud computer users local administrative permissions.
	//
	// example:
	//
	// false
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable two-factor authentication. This parameter is applicable to only office sites that use convenience accounts. If enabled, the system performs a security check during logon. If the system detects a risk, it sends a verification code to the email address that is associated with the account. The user must enter the correct verification code to log on.
	//
	// example:
	//
	// false
	NeedVerifyLoginRisk *bool `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	// Specifies whether to enable device verification. This feature is available only for office sites that use convenience accounts.
	//
	// example:
	//
	// false
	NeedVerifyZeroDevice *bool `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	// The ID of the office site.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-882398****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the office site. The name must be 2 to 255 characters long. The name must start with a letter or a Chinese character, and cannot start with http\\:// or https\\://. It can contain digits, colons (:), underscores (_), and hyphens (-).<br>
	//
	// This parameter is optional.<br>
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions where Elastic Desktop Service is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId  *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s ModifyOfficeSiteAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeRequest) GetAuthorityHost() *string {
	return s.AuthorityHost
}

func (s *ModifyOfficeSiteAttributeRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ModifyOfficeSiteAttributeRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ModifyOfficeSiteAttributeRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *ModifyOfficeSiteAttributeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyOfficeSiteAttributeRequest) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *ModifyOfficeSiteAttributeRequest) GetNeedVerifyLoginRisk() *bool {
	return s.NeedVerifyLoginRisk
}

func (s *ModifyOfficeSiteAttributeRequest) GetNeedVerifyZeroDevice() *bool {
	return s.NeedVerifyZeroDevice
}

func (s *ModifyOfficeSiteAttributeRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyOfficeSiteAttributeRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *ModifyOfficeSiteAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfficeSiteAttributeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ModifyOfficeSiteAttributeRequest) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *ModifyOfficeSiteAttributeRequest) SetAuthorityHost(v string) *ModifyOfficeSiteAttributeRequest {
	s.AuthorityHost = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetClientId(v string) *ModifyOfficeSiteAttributeRequest {
	s.ClientId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetClientSecret(v string) *ModifyOfficeSiteAttributeRequest {
	s.ClientSecret = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetDomainName(v string) *ModifyOfficeSiteAttributeRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetEnableAdminAccess(v bool) *ModifyOfficeSiteAttributeRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetNeedVerifyLoginRisk(v bool) *ModifyOfficeSiteAttributeRequest {
	s.NeedVerifyLoginRisk = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetNeedVerifyZeroDevice(v bool) *ModifyOfficeSiteAttributeRequest {
	s.NeedVerifyZeroDevice = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteName(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetRegionId(v string) *ModifyOfficeSiteAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetTenantId(v string) *ModifyOfficeSiteAttributeRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetVSwitchId(v []*string) *ModifyOfficeSiteAttributeRequest {
	s.VSwitchId = v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) Validate() error {
	return dara.Validate(s)
}
