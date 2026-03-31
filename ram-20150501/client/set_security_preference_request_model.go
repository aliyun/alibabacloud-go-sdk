// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToChangePassword() *bool
	SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManageAccessKeys() *bool
	SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManageMFADevices() *bool
	SetAllowUserToManagePublicKeys(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManagePublicKeys() *bool
	SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest
	GetEnableSaveMFATicket() *bool
	SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest
	GetLoginNetworkMasks() *string
	SetLoginSessionDuration(v int32) *SetSecurityPreferenceRequest
	GetLoginSessionDuration() *int32
}

type SetSecurityPreferenceRequest struct {
	// Specifies whether RAM users can change their passwords. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can manage their public keys. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// >  This parameter is valid only for the Japan site.
	//
	// example:
	//
	// false
	AllowUserToManagePublicKeys *bool `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
	// Specifies whether to remember the multi-factor authentication (MFA) devices of Resource Access Management (RAM) users for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// You can specify up to 40 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// Valid values: 1 to 24. Default value: 6. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
}

func (s SetSecurityPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManagePublicKeys() *bool {
	return s.AllowUserToManagePublicKeys
}

func (s *SetSecurityPreferenceRequest) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *SetSecurityPreferenceRequest) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *SetSecurityPreferenceRequest) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManagePublicKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManagePublicKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginSessionDuration(v int32) *SetSecurityPreferenceRequest {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
