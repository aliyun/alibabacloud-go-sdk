// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetExternalSAMLIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindingType(v string) *SetExternalSAMLIdentityProviderRequest
	GetBindingType() *string
	SetDirectoryId(v string) *SetExternalSAMLIdentityProviderRequest
	GetDirectoryId() *string
	SetEncodedMetadataDocument(v string) *SetExternalSAMLIdentityProviderRequest
	GetEncodedMetadataDocument() *string
	SetEntityId(v string) *SetExternalSAMLIdentityProviderRequest
	GetEntityId() *string
	SetLoginUrl(v string) *SetExternalSAMLIdentityProviderRequest
	GetLoginUrl() *string
	SetSSOStatus(v string) *SetExternalSAMLIdentityProviderRequest
	GetSSOStatus() *string
	SetWantRequestSigned(v bool) *SetExternalSAMLIdentityProviderRequest
	GetWantRequestSigned() *bool
	SetX509Certificate(v string) *SetExternalSAMLIdentityProviderRequest
	GetX509Certificate() *string
}

type SetExternalSAMLIdentityProviderRequest struct {
	// The binding for sending SAML requests. Valid values:
	//
	// 	- Post: HTTP Post bindings.
	//
	// 	- Redirect: HTTP Redirect bindings.
	//
	// example:
	//
	// Redirect
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the IdP. The value of this parameter is Base64-encoded.
	//
	// The file is provided by the IdP that supports SAML 2.0.
	//
	// example:
	//
	// PD94bWwgdmVyc2lvbj0iMS4****
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the IdP.
	//
	// example:
	//
	// http://www.okta.com/exk3qwgtjhetR2Od****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The logon URL of the IdP.
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The status of SSO logon. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled (default)
	//
	// example:
	//
	// Disabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// Specifies whether CloudSSO needs to sign SAML requests. The requests are sent when users log on to the CloudSSO user portal to initiate SAML-based SSO. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	WantRequestSigned *bool `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
	// The X.509 certificate in the PEM format. If you specify this parameter, all existing certificates are replaced.
	//
	// example:
	//
	// MIIC8DCCAdigAwIBAgIQP9eomUYGeoND****
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s SetExternalSAMLIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderRequest) GetBindingType() *string {
	return s.BindingType
}

func (s *SetExternalSAMLIdentityProviderRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetExternalSAMLIdentityProviderRequest) GetEncodedMetadataDocument() *string {
	return s.EncodedMetadataDocument
}

func (s *SetExternalSAMLIdentityProviderRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SetExternalSAMLIdentityProviderRequest) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *SetExternalSAMLIdentityProviderRequest) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetExternalSAMLIdentityProviderRequest) GetWantRequestSigned() *bool {
	return s.WantRequestSigned
}

func (s *SetExternalSAMLIdentityProviderRequest) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *SetExternalSAMLIdentityProviderRequest) SetBindingType(v string) *SetExternalSAMLIdentityProviderRequest {
	s.BindingType = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetDirectoryId(v string) *SetExternalSAMLIdentityProviderRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetEncodedMetadataDocument(v string) *SetExternalSAMLIdentityProviderRequest {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetEntityId(v string) *SetExternalSAMLIdentityProviderRequest {
	s.EntityId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetLoginUrl(v string) *SetExternalSAMLIdentityProviderRequest {
	s.LoginUrl = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetSSOStatus(v string) *SetExternalSAMLIdentityProviderRequest {
	s.SSOStatus = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetWantRequestSigned(v bool) *SetExternalSAMLIdentityProviderRequest {
	s.WantRequestSigned = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetX509Certificate(v string) *SetExternalSAMLIdentityProviderRequest {
	s.X509Certificate = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
