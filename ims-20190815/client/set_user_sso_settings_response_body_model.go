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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user-based SSO settings.
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
	if s.UserSsoSettings != nil {
		if err := s.UserSsoSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetUserSsoSettingsResponseBodyUserSsoSettings struct {
	// The signature algorithm that is supported by the Alibaba Cloud SP. Valid values:
	//
	// - rsa-sha256
	//
	// - rsa-sha1 (default)
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The auxiliary domain name.
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file. The file is Base64-encoded.
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Indicates whether user-based SSO is enabled.
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Indicates whether the `<saml:NameID>` element in a SAML response must contain a domain name when a user logs on using SAML-based SSO. This applies if the username that is specified on the IdP for logon matching contains a domain name suffix.
	//
	// - If this parameter is set to `true`, the value of the `<saml:NameID>` element **must*	- be in the `username@domain` format, which includes a domain name suffix. The `domain` can be the default domain name or a domain alias if one is configured.
	//
	// - If this parameter is set to `false`, the value of the `<saml:NameID>` element **must*	- be the `username` only. The value **must not*	- contain the `domain` part.
	//
	// The default value is `true`.
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s SetUserSsoSettingsResponseBodyUserSsoSettings) String() string {
	return dara.Prettify(s)
}

func (s SetUserSsoSettingsResponseBodyUserSsoSettings) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
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

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetAuthnSignAlgo(v string) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.AuthnSignAlgo = &v
	return s
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
