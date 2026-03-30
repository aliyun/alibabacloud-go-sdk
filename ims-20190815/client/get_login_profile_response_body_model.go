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
	// The console logon configurations.
	LoginProfile *GetLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
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
	// Indicates whether console logon is automatically disabled if a RAM user does not log on to the console in the previous specified number of days. The number of days is specified by MaxIdleDaysForUsers. The default value is true, and you cannot change the value.
	//
	// example:
	//
	// true
	AutoDisableLoginStatus *string `json:"AutoDisableLoginStatus,omitempty" xml:"AutoDisableLoginStatus,omitempty"`
	// The time of the most recent logon. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-14T07:25:25Z
	LastLoginTime *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	// Indicates whether multi-factor authentication (MFA) must be enabled. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user is required to reset the password upon the next logon. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	PasswordStatus        *string `json:"PasswordStatus,omitempty" xml:"PasswordStatus,omitempty"`
	// Indicates whether console logon is enabled. Valid values:
	//
	// 	- Active: enabled.
	//
	// 	- Inactive: disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The modification time. The time is displayed in UTC.
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
