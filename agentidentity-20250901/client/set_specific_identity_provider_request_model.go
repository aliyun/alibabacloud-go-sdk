// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSpecificIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIDPMetadata(v string) *SetSpecificIdentityProviderRequest
	GetIDPMetadata() *string
	SetIdentityProviderType(v string) *SetSpecificIdentityProviderRequest
	GetIdentityProviderType() *string
	SetSSOStatus(v string) *SetSpecificIdentityProviderRequest
	GetSSOStatus() *string
	SetUserPoolName(v string) *SetSpecificIdentityProviderRequest
	GetUserPoolName() *string
}

type SetSpecificIdentityProviderRequest struct {
	IDPMetadata          *string `json:"IDPMetadata,omitempty" xml:"IDPMetadata,omitempty"`
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	SSOStatus            *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	UserPoolName         *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s SetSpecificIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSpecificIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *SetSpecificIdentityProviderRequest) GetIDPMetadata() *string {
	return s.IDPMetadata
}

func (s *SetSpecificIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *SetSpecificIdentityProviderRequest) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetSpecificIdentityProviderRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *SetSpecificIdentityProviderRequest) SetIDPMetadata(v string) *SetSpecificIdentityProviderRequest {
	s.IDPMetadata = &v
	return s
}

func (s *SetSpecificIdentityProviderRequest) SetIdentityProviderType(v string) *SetSpecificIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *SetSpecificIdentityProviderRequest) SetSSOStatus(v string) *SetSpecificIdentityProviderRequest {
	s.SSOStatus = &v
	return s
}

func (s *SetSpecificIdentityProviderRequest) SetUserPoolName(v string) *SetSpecificIdentityProviderRequest {
	s.UserPoolName = &v
	return s
}

func (s *SetSpecificIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
