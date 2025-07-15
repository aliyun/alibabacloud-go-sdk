// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest
	GetDesktopAccessType() *string
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
}

type ModifyOfficeSiteAttributeRequest struct {
	// The method to connect to cloud computers from Alibaba Cloud Workspace clients.
	//
	// >  VPC connection relies on the Alibaba Cloud PrivateLink service. You can use PrivateLink for free. When you set this parameter to `VPC` or `Any`, the system automatically activates PrivateLink.
	//
	// Valid values:
	//
	// 	- INTERNET (default): allows end users to connect to cloud computers over the Internet.
	//
	// 	- VPC: allows end users to connect to cloud computers over VPCs.
	//
	// 	- ANY: allows end users to connect to cloud computers over the Internet and VPCs. When end users connect to cloud computers from Elastic Desktop Service, you can choose a connection method based on your business requirements.
	//
	// example:
	//
	// INTERNET
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// Specifies whether to grant the local administrator permissions to users that are authorized to use cloud computers in the office network.
	//
	// Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable two-factor verification when an end user logs on to an Alibaba Cloud Workspace client. This parameter is required only for convenience office networks. If two-factor verification is enabled, the system checks whether security risks exist within the logon account when the end user uses a convenience user to log on to the client. If risks are detected, the system sends a verification code to the email address that is associated with the account of the convenience user. Then, the end user can log on to the client only when the verification code is correct.
	//
	// example:
	//
	// false
	NeedVerifyLoginRisk *bool `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	// Specifies whether to enable device verification. This parameter is required only for convenience office networks. This parameter is left empty for enterprise Active Directory (AD) office networks.
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
	// The office network name. The name must be 2 to 255 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.\\
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
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

func (s *ModifyOfficeSiteAttributeRequest) SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest {
	s.DesktopAccessType = &v
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

func (s *ModifyOfficeSiteAttributeRequest) Validate() error {
	return dara.Validate(s)
}
