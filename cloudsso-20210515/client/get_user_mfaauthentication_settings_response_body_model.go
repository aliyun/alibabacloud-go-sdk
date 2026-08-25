// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAAuthenticationSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserMFAAuthenticationSettingsResponseBody
	GetRequestId() *string
	SetUserMFAAuthenticationSettings(v string) *GetUserMFAAuthenticationSettingsResponseBody
	GetUserMFAAuthenticationSettings() *string
}

type GetUserMFAAuthenticationSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B598B62-85E6-5792-9DF1-246D251B07DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether MFA is enabled for the user. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	UserMFAAuthenticationSettings *string `json:"UserMFAAuthenticationSettings,omitempty" xml:"UserMFAAuthenticationSettings,omitempty"`
}

func (s GetUserMFAAuthenticationSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) GetUserMFAAuthenticationSettings() *string {
	return s.UserMFAAuthenticationSettings
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *GetUserMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) SetUserMFAAuthenticationSettings(v string) *GetUserMFAAuthenticationSettingsResponseBody {
	s.UserMFAAuthenticationSettings = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
