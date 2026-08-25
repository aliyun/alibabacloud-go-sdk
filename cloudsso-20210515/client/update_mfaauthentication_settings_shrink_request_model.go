// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMFAAuthenticationSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedVerificationTypesShrink(v string) *UpdateMFAAuthenticationSettingsShrinkRequest
	GetAllowedVerificationTypesShrink() *string
	SetDirectoryId(v string) *UpdateMFAAuthenticationSettingsShrinkRequest
	GetDirectoryId() *string
	SetMFAAuthenticationSettings(v string) *UpdateMFAAuthenticationSettingsShrinkRequest
	GetMFAAuthenticationSettings() *string
	SetOperationForRiskLogin(v string) *UpdateMFAAuthenticationSettingsShrinkRequest
	GetOperationForRiskLogin() *string
}

type UpdateMFAAuthenticationSettingsShrinkRequest struct {
	AllowedVerificationTypesShrink *string `json:"AllowedVerificationTypes,omitempty" xml:"AllowedVerificationTypes,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The global MFA settings. Valid values:
	//
	// - Enabled: MFA verification is enabled for all users.
	//
	// - Byuser: MFA verification depends on the individual MFA settings of each user. For more information about individual user MFA settings, see [UpdateUserMFAAuthenticationSettings](https://help.aliyun.com/document_detail/450135.html).
	//
	// - Disabled: MFA verification is disabled for all users.
	//
	// - OnlyRiskyLogin: MFA verification is required only for unusual logon attempts.
	//
	// example:
	//
	// Enabled
	MFAAuthenticationSettings *string `json:"MFAAuthenticationSettings,omitempty" xml:"MFAAuthenticationSettings,omitempty"`
	// The action to take when the MFA settings option is set to verify only for unusual logon attempts. Valid values:
	//
	// - Autonomous: Users can skip MFA binding during unusual logon, but users who have already bound MFA must complete MFA verification.
	//
	// - EnforceVerify: Users who have not bound MFA are required to bind it, and users who have already bound MFA must complete verification.
	//
	// example:
	//
	// Autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s UpdateMFAAuthenticationSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) GetAllowedVerificationTypesShrink() *string {
	return s.AllowedVerificationTypesShrink
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) GetMFAAuthenticationSettings() *string {
	return s.MFAAuthenticationSettings
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) SetAllowedVerificationTypesShrink(v string) *UpdateMFAAuthenticationSettingsShrinkRequest {
	s.AllowedVerificationTypesShrink = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) SetDirectoryId(v string) *UpdateMFAAuthenticationSettingsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) SetMFAAuthenticationSettings(v string) *UpdateMFAAuthenticationSettingsShrinkRequest {
	s.MFAAuthenticationSettings = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) SetOperationForRiskLogin(v string) *UpdateMFAAuthenticationSettingsShrinkRequest {
	s.OperationForRiskLogin = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
