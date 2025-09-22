// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserSsoSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthnSignAlgo(v string) *SetUserSsoSettingsRequest
	GetAuthnSignAlgo() *string
	SetAuxiliaryDomain(v string) *SetUserSsoSettingsRequest
	GetAuxiliaryDomain() *string
	SetMetadataDocument(v string) *SetUserSsoSettingsRequest
	GetMetadataDocument() *string
	SetSsoEnabled(v bool) *SetUserSsoSettingsRequest
	GetSsoEnabled() *bool
	SetSsoLoginWithDomain(v bool) *SetUserSsoSettingsRequest
	GetSsoLoginWithDomain() *bool
}

type SetUserSsoSettingsRequest struct {
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The auxiliary domain name.
	//
	// example:
	//
	// example.com
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file, which is Base64-encoded.
	//
	// The file is provided by an identity provider (IdP) that supports Security Assertion Markup Language (SAML) 2.0.
	//
	// example:
	//
	// PD94bWwgdmVy****
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Specifies whether to enable SSO for the RAM user. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Specifies whether the SAML SSO requires a domain name in the `<saml:NameID>` element of the SAML response. If yes, the username specified by the IdP for SSO must have a domain name as the suffix.
	//
	// 	- If the value of the parameter is `true`, the `<saml:NameID>` element **must*	- be in the `username@domain` format. You can set `domain` to the default domain name or the configured domain alias.
	//
	// 	- If the value of the parameter is `false`, the `<saml:NameID>` element **must*	- be in the `username` format and **cannot*	- contain the `domain` suffix.
	//
	// Set the value to the default `true`.
	//
	// example:
	//
	// true
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s SetUserSsoSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserSsoSettingsRequest) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsRequest) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *SetUserSsoSettingsRequest) GetAuxiliaryDomain() *string {
	return s.AuxiliaryDomain
}

func (s *SetUserSsoSettingsRequest) GetMetadataDocument() *string {
	return s.MetadataDocument
}

func (s *SetUserSsoSettingsRequest) GetSsoEnabled() *bool {
	return s.SsoEnabled
}

func (s *SetUserSsoSettingsRequest) GetSsoLoginWithDomain() *bool {
	return s.SsoLoginWithDomain
}

func (s *SetUserSsoSettingsRequest) SetAuthnSignAlgo(v string) *SetUserSsoSettingsRequest {
	s.AuthnSignAlgo = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetAuxiliaryDomain(v string) *SetUserSsoSettingsRequest {
	s.AuxiliaryDomain = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetMetadataDocument(v string) *SetUserSsoSettingsRequest {
	s.MetadataDocument = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetSsoEnabled(v bool) *SetUserSsoSettingsRequest {
	s.SsoEnabled = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetSsoLoginWithDomain(v bool) *SetUserSsoSettingsRequest {
	s.SsoLoginWithDomain = &v
	return s
}

func (s *SetUserSsoSettingsRequest) Validate() error {
	return dara.Validate(s)
}
