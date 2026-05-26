// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSAMLIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *SetSAMLIdentityProviderRequest
	GetEntityId() *string
	SetJITProvisionStatus(v string) *SetSAMLIdentityProviderRequest
	GetJITProvisionStatus() *string
	SetJITUpdateStatus(v string) *SetSAMLIdentityProviderRequest
	GetJITUpdateStatus() *string
	SetLoginURL(v string) *SetSAMLIdentityProviderRequest
	GetLoginURL() *string
	SetSAMLBindingType(v string) *SetSAMLIdentityProviderRequest
	GetSAMLBindingType() *string
	SetSSOStatus(v string) *SetSAMLIdentityProviderRequest
	GetSSOStatus() *string
	SetUserPoolName(v string) *SetSAMLIdentityProviderRequest
	GetUserPoolName() *string
	SetX509Certificates(v []*string) *SetSAMLIdentityProviderRequest
	GetX509Certificates() []*string
}

type SetSAMLIdentityProviderRequest struct {
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
	UserPoolName     *string   `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
	X509Certificates []*string `json:"X509Certificates,omitempty" xml:"X509Certificates,omitempty" type:"Repeated"`
}

func (s SetSAMLIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *SetSAMLIdentityProviderRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SetSAMLIdentityProviderRequest) GetJITProvisionStatus() *string {
	return s.JITProvisionStatus
}

func (s *SetSAMLIdentityProviderRequest) GetJITUpdateStatus() *string {
	return s.JITUpdateStatus
}

func (s *SetSAMLIdentityProviderRequest) GetLoginURL() *string {
	return s.LoginURL
}

func (s *SetSAMLIdentityProviderRequest) GetSAMLBindingType() *string {
	return s.SAMLBindingType
}

func (s *SetSAMLIdentityProviderRequest) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetSAMLIdentityProviderRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *SetSAMLIdentityProviderRequest) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *SetSAMLIdentityProviderRequest) SetEntityId(v string) *SetSAMLIdentityProviderRequest {
	s.EntityId = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetJITProvisionStatus(v string) *SetSAMLIdentityProviderRequest {
	s.JITProvisionStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetJITUpdateStatus(v string) *SetSAMLIdentityProviderRequest {
	s.JITUpdateStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetLoginURL(v string) *SetSAMLIdentityProviderRequest {
	s.LoginURL = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetSAMLBindingType(v string) *SetSAMLIdentityProviderRequest {
	s.SAMLBindingType = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetSSOStatus(v string) *SetSAMLIdentityProviderRequest {
	s.SSOStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetUserPoolName(v string) *SetSAMLIdentityProviderRequest {
	s.UserPoolName = &v
	return s
}

func (s *SetSAMLIdentityProviderRequest) SetX509Certificates(v []*string) *SetSAMLIdentityProviderRequest {
	s.X509Certificates = v
	return s
}

func (s *SetSAMLIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
