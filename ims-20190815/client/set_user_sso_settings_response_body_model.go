// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserSsoSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetUserSsoSettingsResponseBody
	GetRequestId() *string
	SetUserSsoSettings(v *SetUserSsoSettingsResponseBodyUserSsoSettings) *SetUserSsoSettingsResponseBody
	GetUserSsoSettings() *SetUserSsoSettingsResponseBodyUserSsoSettings
}

type SetUserSsoSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 87F2E3F6-28A0-43F3-A77F-F7760E62F61E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of user-based SSO.
	UserSsoSettings *SetUserSsoSettingsResponseBodyUserSsoSettings `json:"UserSsoSettings,omitempty" xml:"UserSsoSettings,omitempty" type:"Struct"`
}

func (s SetUserSsoSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetUserSsoSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetUserSsoSettingsResponseBody) GetUserSsoSettings() *SetUserSsoSettingsResponseBodyUserSsoSettings {
	return s.UserSsoSettings
}

func (s *SetUserSsoSettingsResponseBody) SetRequestId(v string) *SetUserSsoSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserSsoSettingsResponseBody) SetUserSsoSettings(v *SetUserSsoSettingsResponseBodyUserSsoSettings) *SetUserSsoSettingsResponseBody {
	s.UserSsoSettings = v
	return s
}

func (s *SetUserSsoSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetUserSsoSettingsResponseBodyUserSsoSettings struct {
	// The auxiliary domain name.
	//
	// example:
	//
	// example.com
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file, which is Base64-encoded.
	//
	// example:
	//
	// PD94bWwgdmVy****
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Indicates whether user-based SSO is enabled.
	//
	// example:
	//
	// true
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Indicates whether the SAML SSO requires a domain name in the `<saml:NameID>` element of the SAML response. If yes, the username specified by the IdP for SSO must have a domain name as the suffix.
	//
	// 	- If the value of the parameter is `true`, the `<saml:NameID>` element **must*	- be in the `username@domain` format. You can set `domain` to the default domain name or the configured domain alias.
	//
	// 	- If the value of the parameter is `false`, the `<saml:NameID>` element **must*	- be in the `username` format and **cannot*	- contain the `domain` suffix.
	//
	// The default value is `true`.
	//
	// example:
	//
	// true
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s SetUserSsoSettingsResponseBodyUserSsoSettings) String() string {
	return dara.Prettify(s)
}

func (s SetUserSsoSettingsResponseBodyUserSsoSettings) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) GetAuxiliaryDomain() *string {
	return s.AuxiliaryDomain
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) GetMetadataDocument() *string {
	return s.MetadataDocument
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) GetSsoEnabled() *bool {
	return s.SsoEnabled
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) GetSsoLoginWithDomain() *bool {
	return s.SsoLoginWithDomain
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetAuxiliaryDomain(v string) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.AuxiliaryDomain = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetMetadataDocument(v string) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.MetadataDocument = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoEnabled(v bool) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoEnabled = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoLoginWithDomain(v bool) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoLoginWithDomain = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) Validate() error {
	return dara.Validate(s)
}
