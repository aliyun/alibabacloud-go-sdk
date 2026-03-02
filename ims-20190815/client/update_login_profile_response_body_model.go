// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginProfile(v *UpdateLoginProfileResponseBodyLoginProfile) *UpdateLoginProfileResponseBody
	GetLoginProfile() *UpdateLoginProfileResponseBodyLoginProfile
	SetRequestId(v string) *UpdateLoginProfileResponseBody
	GetRequestId() *string
}

type UpdateLoginProfileResponseBody struct {
	// The console logon information.
	LoginProfile *UpdateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BCDB6A7F-2199-41D9-B577-4FA536A5ADE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoginProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseBody) GetLoginProfile() *UpdateLoginProfileResponseBodyLoginProfile {
	return s.LoginProfile
}

func (s *UpdateLoginProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoginProfileResponseBody) SetLoginProfile(v *UpdateLoginProfileResponseBodyLoginProfile) *UpdateLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *UpdateLoginProfileResponseBody) SetRequestId(v string) *UpdateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoginProfileResponseBody) Validate() error {
	if s.LoginProfile != nil {
		if err := s.LoginProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether to automatically disable console logon for an inactive account. This feature is enabled by default and cannot be disabled.
	//
	// example:
	//
	// true
	AutoDisableLoginStatus *string `json:"AutoDisableLoginStatus,omitempty" xml:"AutoDisableLoginStatus,omitempty"`
	// Indicates whether MFA is enforced for the user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user must reset the password at the next logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The status of the initial password. An initial password is the one set when a logon profile is created or console logon is re-enabled.
	//
	// Valid values:
	//
	// - "NotInitial": Not an initial password.
	//
	// - "InitialValid": The initial password is valid.
	//
	// - "InitialExpired": The initial password has expired.
	//
	// example:
	//
	// NotInitial
	PasswordStatus *string `json:"PasswordStatus,omitempty" xml:"PasswordStatus,omitempty"`
	// Specifies whether password logon to the console is enabled or disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the logon profile was updated.
	//
	// example:
	//
	// 2020-10-14T07:48:41Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example11.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateLoginProfileResponseBodyLoginProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetAutoDisableLoginStatus() *string {
	return s.AutoDisableLoginStatus
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetPasswordStatus() *string {
	return s.PasswordStatus
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetStatus() *string {
	return s.Status
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetAutoDisableLoginStatus(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.AutoDisableLoginStatus = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *UpdateLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *UpdateLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetPasswordStatus(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.PasswordStatus = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetStatus(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.Status = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetUpdateDate(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetUserPrincipalName(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) Validate() error {
	return dara.Validate(s)
}
