// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecificIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderType(v string) *GetSpecificIdentityProviderRequest
	GetIdentityProviderType() *string
	SetUserPoolName(v string) *GetSpecificIdentityProviderRequest
	GetUserPoolName() *string
}

type GetSpecificIdentityProviderRequest struct {
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	UserPoolName         *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetSpecificIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetSpecificIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetSpecificIdentityProviderRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetSpecificIdentityProviderRequest) SetIdentityProviderType(v string) *GetSpecificIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *GetSpecificIdentityProviderRequest) SetUserPoolName(v string) *GetSpecificIdentityProviderRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetSpecificIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
