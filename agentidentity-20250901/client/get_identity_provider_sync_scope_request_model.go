// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderType(v string) *GetIdentityProviderSyncScopeRequest
	GetIdentityProviderType() *string
	SetUserPoolName(v string) *GetIdentityProviderSyncScopeRequest
	GetUserPoolName() *string
}

type GetIdentityProviderSyncScopeRequest struct {
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// example:
	//
	// syncscope-example
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetIdentityProviderSyncScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetIdentityProviderSyncScopeRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetIdentityProviderSyncScopeRequest) SetIdentityProviderType(v string) *GetIdentityProviderSyncScopeRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *GetIdentityProviderSyncScopeRequest) SetUserPoolName(v string) *GetIdentityProviderSyncScopeRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetIdentityProviderSyncScopeRequest) Validate() error {
	return dara.Validate(s)
}
