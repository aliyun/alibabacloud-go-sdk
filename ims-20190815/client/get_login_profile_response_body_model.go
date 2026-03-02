// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginProfile(v *GetLoginProfileResponseBodyLoginProfile) *GetLoginProfileResponseBody
	GetLoginProfile() *GetLoginProfileResponseBodyLoginProfile
	SetRequestId(v string) *GetLoginProfileResponseBody
	GetRequestId() *string
}

type GetLoginProfileResponseBody struct {
	// The logon information for the console.
	LoginProfile *GetLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E517F18B-632C-48FC-93F1-CDCBCC6F8444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBody) GetLoginProfile() *GetLoginProfileResponseBodyLoginProfile {
	return s.LoginProfile
}

func (s *GetLoginProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginProfileResponseBody) SetLoginProfile(v *GetLoginProfileResponseBodyLoginProfile) *GetLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *GetLoginProfileResponseBody) SetRequestId(v string) *GetLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginProfileResponseBody) Validate() error {
	if s.LoginProfile != nil {
		if err := s.LoginProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether console logon is automatically disabled if the user is inactive. This feature is enabled by default and cannot be disabled.
	//
	// example:
	//
	// true
	AutoDisableLoginStatus *string `json:"AutoDisableLoginStatus,omitempty" xml:"AutoDisableLoginStatus,omitempty"`
	// The time when the RAM user last logged on to the console. The time is in UTC.
	//
	// example:
	//
	// 2020-10-14T07:25:25Z
	LastLoginTime *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	// Indicates whether multi-factor authentication (MFA) is required for the user. Valid values:
	//
	// - false: MFA is not required.
	//
	// - true: MFA is required.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user must reset the password at the next logon. Valid values:
	//
	// - false: The RAM user is not required to reset the password.
	//
	// - true: The RAM user is required to reset the password.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The status of the initial password. An initial password is the password that is configured when you create a logon profile or re-enable console logon.
	//
	// Valid values
	//
	// - "NotInitial": The password is not an initial password.
	//
	// - "InitialValid": The initial password is valid.
	//
	// - "InitialExpired": The initial password has expired.
	//
	// example:
	//
	// NotInitial
	PasswordStatus *string `json:"PasswordStatus,omitempty" xml:"PasswordStatus,omitempty"`
	// The status of console logon. Valid values:
	//
	// - Active: Console logon is enabled.
	//
	// - Inactive: Console logon is disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the logon profile was last updated. The time is in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2020-10-14T06:56:45Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetLoginProfileResponseBodyLoginProfile) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetAutoDisableLoginStatus() *string {
	return s.AutoDisableLoginStatus
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetPasswordStatus() *string {
	return s.PasswordStatus
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetStatus() *string {
	return s.Status
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetAutoDisableLoginStatus(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.AutoDisableLoginStatus = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetLastLoginTime(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.LastLoginTime = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetPasswordStatus(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.PasswordStatus = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetStatus(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.Status = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetUpdateDate(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetUserPrincipalName(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.UserPrincipalName = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) Validate() error {
	return dara.Validate(s)
}
