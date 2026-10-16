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
	// The IdP metadata. For the DingTalk type, the metadata is in JSON format and includes fields such as corpId, appKey, appSecret, encryptKey, verificationToken, and syncCallbackStatus.
	//
	// example:
	//
	// {"corpId":"dingxxxxxxxxxx","appKey":"dingxxxxxxxxxx","appSecret":"your_app_secret","encryptKey":"your_encrypt_key","verificationToken":"your_verification_token","syncCallbackStatus":"Enabled"}
	IDPMetadata *string `json:"IDPMetadata,omitempty" xml:"IDPMetadata,omitempty"`
	// The type of the identity provider. Currently, only DingTalk is supported.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// The SSO enabling status. Valid values:
	//
	// - Enabled: Enabled.
	//
	// - Disabled: Disabled.
	//
	// example:
	//
	// Enabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// The name of the user pool.
	//
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
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
