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
	// The console logon configurations.
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
	return dara.Validate(s)
}

type UpdateLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether console logon is automatically disabled if a RAM user does not log on to the console in the previous specified number of days. The number of days is specified by MaxIdleDaysForUsers. The default value is true, and you cannot change the value.
	//
	// example:
	//
	// true
	AutoDisableLoginStatus *string `json:"AutoDisableLoginStatus,omitempty" xml:"AutoDisableLoginStatus,omitempty"`
	// Indicates whether MFA must be enabled.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user is required to reset the password upon the next logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Indicates whether to enable password-based logons to the console.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The modification time.
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
