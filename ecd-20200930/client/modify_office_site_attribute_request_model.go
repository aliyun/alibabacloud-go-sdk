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
	// The Authority URL of the identity authentication service.
	//
	// example:
	//
	// https://login.microsoftonline.com
	AuthorityHost *string `json:"AuthorityHost,omitempty" xml:"AuthorityHost,omitempty"`
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
	// The access method allowed when connecting to cloud computers.
	//
	// > The VPC connection method depends on the Alibaba Cloud PrivateLink service, which is free of charge. If this parameter is set to `VPC` or `Any`, the system automatically activates the PrivateLink service for you.
	//
	// example:
	//
	// INTERNET
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The domain name of the enterprise AD.
	//
	// example:
	//
	// domain.local
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to grant local administrator permissions to cloud computer users.
	//
	// example:
	//
	// false
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// This parameter applies only to convenience account-based office networks. Specifies whether secondary authentication is required during logon. If logon secondary authentication is enabled, the system checks whether the logon account has security risks when a convenience user logs on to the client. If a risk is detected, the system sends a verification code to the email address associated with the account. The convenience user can log on to the client only after passing the verification code check.
	//
	// example:
	//
	// false
	NeedVerifyLoginRisk *bool `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	// This parameter applies only to convenience account-based office networks. Specifies whether to enable device verification. For AD-based office networks, this parameter is empty.
	//
	// example:
	//
	// false
	NeedVerifyZeroDevice *bool `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-882398****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name. The name must be 2 to 255 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// Default value: empty.
	//
	// example:
	//
	// R&D_Office_Network
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
	// The vSwitch ID. Only one vSwitch is supported.
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
