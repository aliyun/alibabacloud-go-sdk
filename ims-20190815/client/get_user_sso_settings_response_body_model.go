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
	//
	// example:
	//
	// 69FC3E5E-D3D9-434B-90CA-BBA8E0551A47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of user-based SSO.
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
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
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
	// false
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Indicates whether the Security Assertion Markup Language (SAML) SSO requires a domain name in the `<saml:NameID>` element of the SAML response. If yes, the username specified by the identity provider (IdP) for SSO must have a domain name as the suffix.
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
