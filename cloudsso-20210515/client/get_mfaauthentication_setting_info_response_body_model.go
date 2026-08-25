// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationSettingInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMFAAuthenticationSettingInfo(v *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) *GetMFAAuthenticationSettingInfoResponseBody
	GetMFAAuthenticationSettingInfo() *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo
	SetRequestId(v string) *GetMFAAuthenticationSettingInfoResponseBody
	GetRequestId() *string
}

type GetMFAAuthenticationSettingInfoResponseBody struct {
	// The global MFA verification configuration.
	MFAAuthenticationSettingInfo *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo `json:"MFAAuthenticationSettingInfo,omitempty" xml:"MFAAuthenticationSettingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 95D3B107-DA80-5B34-A3D0-9E82F8F0DA0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) GetMFAAuthenticationSettingInfo() *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo {
	return s.MFAAuthenticationSettingInfo
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) SetMFAAuthenticationSettingInfo(v *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) *GetMFAAuthenticationSettingInfoResponseBody {
	s.MFAAuthenticationSettingInfo = v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) SetRequestId(v string) *GetMFAAuthenticationSettingInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) Validate() error {
	if s.MFAAuthenticationSettingInfo != nil {
		if err := s.MFAAuthenticationSettingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo struct {
	AllowedVerificationTypes []*string `json:"AllowedVerificationTypes,omitempty" xml:"AllowedVerificationTypes,omitempty" type:"Repeated"`
	// The global MFA verification policy. Valid values:
	//
	// - Enabled: MFA verification is enabled for all users.
	//
	// - Byuser: MFA verification depends on the independent MFA configuration of each user. For more information about user-specific MFA configuration, see [UpdateUserMFAAuthenticationSettings](https://help.aliyun.com/document_detail/450135.html).
	//
	// - Disabled: MFA verification is disabled for all users.
	//
	// - OnlyRiskyLogin: MFA verification is required only for unusual logon attempts.
	//
	// example:
	//
	// OnlyRiskyLogin
	MfaAuthenticationAdvanceSettings *string `json:"MfaAuthenticationAdvanceSettings,omitempty" xml:"MfaAuthenticationAdvanceSettings,omitempty"`
	// The MFA verification policy for unusual logon attempts. Valid values:
	//
	// - Autonomous: Users can skip MFA binding during unusual logon, but users who have already bound MFA must complete verification.
	//
	// - EnforceVerify: Users are required to bind or verify MFA during unusual logon.
	//
	// > This parameter is displayed only when MfaAuthenticationAdvanceSettings is set to Byuser or OnlyRiskyLogin.
	//
	// example:
	//
	// EnforceVerify
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) GetAllowedVerificationTypes() []*string {
	return s.AllowedVerificationTypes
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) GetMfaAuthenticationAdvanceSettings() *string {
	return s.MfaAuthenticationAdvanceSettings
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) SetAllowedVerificationTypes(v []*string) *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo {
	s.AllowedVerificationTypes = v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) SetMfaAuthenticationAdvanceSettings(v string) *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo {
	s.MfaAuthenticationAdvanceSettings = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) SetOperationForRiskLogin(v string) *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo {
	s.OperationForRiskLogin = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) Validate() error {
	return dara.Validate(s)
}
