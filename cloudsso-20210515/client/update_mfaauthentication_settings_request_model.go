// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMFAAuthenticationSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedVerificationTypes(v []*string) *UpdateMFAAuthenticationSettingsRequest
	GetAllowedVerificationTypes() []*string
	SetDirectoryId(v string) *UpdateMFAAuthenticationSettingsRequest
	GetDirectoryId() *string
	SetMFAAuthenticationSettings(v string) *UpdateMFAAuthenticationSettingsRequest
	GetMFAAuthenticationSettings() *string
	SetOperationForRiskLogin(v string) *UpdateMFAAuthenticationSettingsRequest
	GetOperationForRiskLogin() *string
}

type UpdateMFAAuthenticationSettingsRequest struct {
	AllowedVerificationTypes []*string `json:"AllowedVerificationTypes,omitempty" xml:"AllowedVerificationTypes,omitempty" type:"Repeated"`
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

func (s UpdateMFAAuthenticationSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsRequest) GetAllowedVerificationTypes() []*string {
	return s.AllowedVerificationTypes
}

func (s *UpdateMFAAuthenticationSettingsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateMFAAuthenticationSettingsRequest) GetMFAAuthenticationSettings() *string {
	return s.MFAAuthenticationSettings
}

func (s *UpdateMFAAuthenticationSettingsRequest) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetAllowedVerificationTypes(v []*string) *UpdateMFAAuthenticationSettingsRequest {
	s.AllowedVerificationTypes = v
	return s
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *UpdateMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetMFAAuthenticationSettings(v string) *UpdateMFAAuthenticationSettingsRequest {
	s.MFAAuthenticationSettings = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetOperationForRiskLogin(v string) *UpdateMFAAuthenticationSettingsRequest {
	s.OperationForRiskLogin = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsRequest) Validate() error {
	return dara.Validate(s)
}
