// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMFAAuthenticationAdvanceSettings(v string) *GetMFAAuthenticationSettingsResponseBody
	GetMFAAuthenticationAdvanceSettings() *string
	SetRequestId(v string) *GetMFAAuthenticationSettingsResponseBody
	GetRequestId() *string
}

type GetMFAAuthenticationSettingsResponseBody struct {
	// Indicates whether MFA is enabled for all users. Valid values:
	//
	// 	- Enabled: MFA is enabled for all users.
	//
	// 	- Byuser: User-specific settings are applied.
	//
	// 	- Disabled: MFA is disabled for all users.
	//
	// example:
	//
	// Enabled
	MFAAuthenticationAdvanceSettings *string `json:"MFAAuthenticationAdvanceSettings,omitempty" xml:"MFAAuthenticationAdvanceSettings,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A2BC00C5-76A2-5FFC-A340-927940A98377
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMFAAuthenticationSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingsResponseBody) GetMFAAuthenticationAdvanceSettings() *string {
	return s.MFAAuthenticationAdvanceSettings
}

func (s *GetMFAAuthenticationSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMFAAuthenticationSettingsResponseBody) SetMFAAuthenticationAdvanceSettings(v string) *GetMFAAuthenticationSettingsResponseBody {
	s.MFAAuthenticationAdvanceSettings = &v
	return s
}

func (s *GetMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *GetMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMFAAuthenticationSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
