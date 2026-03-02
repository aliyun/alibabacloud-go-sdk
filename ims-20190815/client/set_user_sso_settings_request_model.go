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
	// The signature algorithm that is supported by the Alibaba Cloud service provider (SP). Valid values:
	//
	// - rsa-sha256
	//
	// - rsa-sha1 (default)
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The auxiliary domain name.
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file. The file must be Base64-encoded.
	//
	// The file is provided by an identity provider (IdP) that supports the Security Assertion Markup Language (SAML) 2.0 protocol.
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Specifies whether to enable user-based SSO for Resource Access Management (RAM) users. Valid values:
	//
	// - true: Enables user-based SSO.
	//
	// - false (default): Disables user-based SSO.
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Specifies whether the `<saml:NameID>` element in a SAML response must contain a domain name when a user logs on using SAML-based SSO. This applies if the username that is specified on the IdP for logon matching contains a domain name suffix.
	//
	// - If this parameter is set to `true`, the value of the `<saml:NameID>` element **must*	- be in the `username@domain` format, which includes a domain name suffix. The `domain` can be the default domain name or a domain alias if one is configured.
	//
	// - If this parameter is set to `false`, the value of the `<saml:NameID>` element **must*	- be the `username` only. The value **must not*	- contain the `domain` part.
	//
	// The default value is `true`.
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
