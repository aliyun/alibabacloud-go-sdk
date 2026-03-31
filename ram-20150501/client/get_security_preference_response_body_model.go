// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSecurityPreferenceResponseBody
	GetRequestId() *string
	SetSecurityPreference(v *GetSecurityPreferenceResponseBodySecurityPreference) *GetSecurityPreferenceResponseBody
	GetSecurityPreference() *GetSecurityPreferenceResponseBodySecurityPreference
}

type GetSecurityPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC1213F1-A9D5-4A01-A996-44983689126C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security preferences.
	SecurityPreference *GetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityPreferenceResponseBody) GetSecurityPreference() *GetSecurityPreferenceResponseBodySecurityPreference {
	return s.SecurityPreference
}

func (s *GetSecurityPreferenceResponseBody) SetRequestId(v string) *GetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityPreferenceResponseBody) SetSecurityPreference(v *GetSecurityPreferenceResponseBodySecurityPreference) *GetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBody) Validate() error {
	if s.SecurityPreference != nil {
		if err := s.SecurityPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The multi-factor authentication (MFA) preference.
	MFAPreference *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The public key preference.
	//
	// >  This parameter is valid only for the Japan site.
	PublicKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference `json:"PublicKeyPreference,omitempty" xml:"PublicKeyPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetAccessKeyPreference() *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	return s.AccessKeyPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetLoginProfilePreference() *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	return s.LoginProfilePreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetMFAPreference() *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	return s.MFAPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetPublicKeyPreference() *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference {
	return s.PublicKeyPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetPublicKeyPreference(v *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.PublicKeyPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) Validate() error {
	if s.AccessKeyPreference != nil {
		if err := s.AccessKeyPreference.Validate(); err != nil {
			return err
		}
	}
	if s.LoginProfilePreference != nil {
		if err := s.LoginProfilePreference.Validate(); err != nil {
			return err
		}
	}
	if s.MFAPreference != nil {
		if err := s.MFAPreference.Validate(); err != nil {
			return err
		}
	}
	if s.PublicKeyPreference != nil {
		if err := s.PublicKeyPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether Resource Access Management (RAM) users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can change their passwords. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether RAM users can save security codes for MFA during logon. Each security code is valid for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that indicates the IP addresses from which logon to the Alibaba Cloud Management Console is allowed. This parameter applies to password-based logon and single sign-on (SSO). However, this parameter does not apply to API calls that are authenticated based on AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you want to specify more than one subnet mask, separate the masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference struct {
	// Indicates whether RAM users can manage their public keys. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AllowUserToManagePublicKeys *bool `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) GetAllowUserToManagePublicKeys() *bool {
	return s.AllowUserToManagePublicKeys
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) SetAllowUserToManagePublicKeys(v bool) *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference {
	s.AllowUserToManagePublicKeys = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) Validate() error {
	return dara.Validate(s)
}
