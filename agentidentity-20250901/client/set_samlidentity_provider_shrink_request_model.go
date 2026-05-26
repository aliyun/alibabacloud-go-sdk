// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSAMLIdentityProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *SetSAMLIdentityProviderShrinkRequest
	GetEntityId() *string
	SetJITProvisionStatus(v string) *SetSAMLIdentityProviderShrinkRequest
	GetJITProvisionStatus() *string
	SetJITUpdateStatus(v string) *SetSAMLIdentityProviderShrinkRequest
	GetJITUpdateStatus() *string
	SetLoginURL(v string) *SetSAMLIdentityProviderShrinkRequest
	GetLoginURL() *string
	SetSAMLBindingType(v string) *SetSAMLIdentityProviderShrinkRequest
	GetSAMLBindingType() *string
	SetSSOStatus(v string) *SetSAMLIdentityProviderShrinkRequest
	GetSSOStatus() *string
	SetUserPoolName(v string) *SetSAMLIdentityProviderShrinkRequest
	GetUserPoolName() *string
	SetX509CertificatesShrink(v string) *SetSAMLIdentityProviderShrinkRequest
	GetX509CertificatesShrink() *string
}

type SetSAMLIdentityProviderShrinkRequest struct {
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// Enabled
	JITProvisionStatus *string `json:"JITProvisionStatus,omitempty" xml:"JITProvisionStatus,omitempty"`
	// example:
	//
	// Enabled
	JITUpdateStatus *string `json:"JITUpdateStatus,omitempty" xml:"JITUpdateStatus,omitempty"`
	// example:
	//
	// https://idp.example.com/sso/saml
	LoginURL *string `json:"LoginURL,omitempty" xml:"LoginURL,omitempty"`
	// example:
	//
	// HTTP-Redirect
	SAMLBindingType *string `json:"SAMLBindingType,omitempty" xml:"SAMLBindingType,omitempty"`
	// example:
	//
	// Enabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName           *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
	X509CertificatesShrink *string `json:"X509Certificates,omitempty" xml:"X509Certificates,omitempty"`
}

func (s SetSAMLIdentityProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSAMLIdentityProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetJITProvisionStatus() *string {
	return s.JITProvisionStatus
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetJITUpdateStatus() *string {
	return s.JITUpdateStatus
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetLoginURL() *string {
	return s.LoginURL
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetSAMLBindingType() *string {
	return s.SAMLBindingType
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *SetSAMLIdentityProviderShrinkRequest) GetX509CertificatesShrink() *string {
	return s.X509CertificatesShrink
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetEntityId(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetJITProvisionStatus(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.JITProvisionStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetJITUpdateStatus(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.JITUpdateStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetLoginURL(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.LoginURL = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetSAMLBindingType(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.SAMLBindingType = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetSSOStatus(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.SSOStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetUserPoolName(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.UserPoolName = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) SetX509CertificatesShrink(v string) *SetSAMLIdentityProviderShrinkRequest {
	s.X509CertificatesShrink = &v
	return s
}

func (s *SetSAMLIdentityProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
