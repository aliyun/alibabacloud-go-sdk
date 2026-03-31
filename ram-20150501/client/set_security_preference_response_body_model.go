// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSecurityPreferenceResponseBody
	GetRequestId() *string
	SetSecurityPreference(v *SetSecurityPreferenceResponseBodySecurityPreference) *SetSecurityPreferenceResponseBody
	GetSecurityPreference() *SetSecurityPreferenceResponseBodySecurityPreference
}

type SetSecurityPreferenceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A978915D-F279-4CA0-A89B-9A71219FFB3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security preferences.
	SecurityPreference *SetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSecurityPreferenceResponseBody) GetSecurityPreference() *SetSecurityPreferenceResponseBodySecurityPreference {
	return s.SecurityPreference
}

func (s *SetSecurityPreferenceResponseBody) SetRequestId(v string) *SetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecurityPreferenceResponseBody) SetSecurityPreference(v *SetSecurityPreferenceResponseBodySecurityPreference) *SetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBody) Validate() error {
	if s.SecurityPreference != nil {
		if err := s.SecurityPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA preference.
	MFAPreference *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The public key preference.
	//
	// >  This parameter is valid only for the Japan site.
	PublicKeyPreference *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference `json:"PublicKeyPreference,omitempty" xml:"PublicKeyPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetAccessKeyPreference() *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	return s.AccessKeyPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetLoginProfilePreference() *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	return s.LoginProfilePreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetMFAPreference() *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	return s.MFAPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetPublicKeyPreference() *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference {
	return s.PublicKeyPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetPublicKeyPreference(v *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.PublicKeyPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) Validate() error {
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

type SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether RAM users can manage their AccessKey pairs.
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can change their passwords.
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether the MFA devices of RAM users are remembered.
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their MFA devices.
	//
	// example:
	//
	// false
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference struct {
	// Indicates whether RAM users can manage their public keys.
	//
	// example:
	//
	// false
	AllowUserToManagePublicKeys *bool `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) GetAllowUserToManagePublicKeys() *bool {
	return s.AllowUserToManagePublicKeys
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) SetAllowUserToManagePublicKeys(v bool) *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference {
	s.AllowUserToManagePublicKeys = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) Validate() error {
	return dara.Validate(s)
}
