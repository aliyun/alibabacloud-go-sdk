// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserSsoSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserSsoSettingsResponseBody
	GetRequestId() *string
	SetUserSsoSettings(v *GetUserSsoSettingsResponseBodyUserSsoSettings) *GetUserSsoSettingsResponseBody
	GetUserSsoSettings() *GetUserSsoSettingsResponseBodyUserSsoSettings
}

type GetUserSsoSettingsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user-based SSO settings.
	UserSsoSettings *GetUserSsoSettingsResponseBodyUserSsoSettings `json:"UserSsoSettings,omitempty" xml:"UserSsoSettings,omitempty" type:"Struct"`
}

func (s GetUserSsoSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserSsoSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserSsoSettingsResponseBody) GetUserSsoSettings() *GetUserSsoSettingsResponseBodyUserSsoSettings {
	return s.UserSsoSettings
}

func (s *GetUserSsoSettingsResponseBody) SetRequestId(v string) *GetUserSsoSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserSsoSettingsResponseBody) SetUserSsoSettings(v *GetUserSsoSettingsResponseBodyUserSsoSettings) *GetUserSsoSettingsResponseBody {
	s.UserSsoSettings = v
	return s
}

func (s *GetUserSsoSettingsResponseBody) Validate() error {
	if s.UserSsoSettings != nil {
		if err := s.UserSsoSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserSsoSettingsResponseBodyUserSsoSettings struct {
	// The signature algorithm supported by the Alibaba Cloud service provider (SP). Valid values:
	//
	// - rsa-sha256
	//
	// - rsa-sha1
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The auxiliary domain name.
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file. The value is Base64-encoded.
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Indicates whether user-based SSO is enabled.
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Specifies whether the `<saml:NameID>` element in a SAML response must contain a domain name when a user logs on using SAML SSO. The username for SSO logon matching is specified on the identity provider (IdP) side.
	//
	// - If this parameter is set to `true`, the value of the `<saml:NameID>` element **must*	- be in the `username@domain` format. The `domain` can be the default domain name or a domain alias, if a domain alias is configured.
	//
	// - If this parameter is set to `false`, the value of the `<saml:NameID>` element \\*\\*must\\*\\	- contain only the \\`username\\` and \\*\\*must not\\*\\	- contain the \\`domain\\` part.
	//
	// The default value is `true`.
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s GetUserSsoSettingsResponseBodyUserSsoSettings) String() string {
	return dara.Prettify(s)
}

func (s GetUserSsoSettingsResponseBodyUserSsoSettings) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) GetAuxiliaryDomain() *string {
	return s.AuxiliaryDomain
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) GetMetadataDocument() *string {
	return s.MetadataDocument
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) GetSsoEnabled() *bool {
	return s.SsoEnabled
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) GetSsoLoginWithDomain() *bool {
	return s.SsoLoginWithDomain
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetAuthnSignAlgo(v string) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.AuthnSignAlgo = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetAuxiliaryDomain(v string) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.AuxiliaryDomain = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetMetadataDocument(v string) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.MetadataDocument = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoEnabled(v bool) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoEnabled = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoLoginWithDomain(v bool) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoLoginWithDomain = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) Validate() error {
	return dara.Validate(s)
}
