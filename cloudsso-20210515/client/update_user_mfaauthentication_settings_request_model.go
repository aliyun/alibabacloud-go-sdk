// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserMFAAuthenticationSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateUserMFAAuthenticationSettingsRequest
	GetDirectoryId() *string
	SetUserId(v string) *UpdateUserMFAAuthenticationSettingsRequest
	GetUserId() *string
	SetUserMFAAuthenticationSettings(v string) *UpdateUserMFAAuthenticationSettingsRequest
	GetUserMFAAuthenticationSettings() *string
}

type UpdateUserMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Specifies whether to enable MFA for the user. Valid values:
	//
	// - Enabled: enables MFA for the user.
	//
	// - Disabled: disables MFA for the user.
	//
	// example:
	//
	// Enabled
	UserMFAAuthenticationSettings *string `json:"UserMFAAuthenticationSettings,omitempty" xml:"UserMFAAuthenticationSettings,omitempty"`
}

func (s UpdateUserMFAAuthenticationSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) GetUserMFAAuthenticationSettings() *string {
	return s.UserMFAAuthenticationSettings
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *UpdateUserMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) SetUserId(v string) *UpdateUserMFAAuthenticationSettingsRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) SetUserMFAAuthenticationSettings(v string) *UpdateUserMFAAuthenticationSettingsRequest {
	s.UserMFAAuthenticationSettings = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) Validate() error {
	return dara.Validate(s)
}
