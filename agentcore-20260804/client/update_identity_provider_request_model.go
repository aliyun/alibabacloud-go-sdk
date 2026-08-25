// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateIdentityProviderRequestBody) *UpdateIdentityProviderRequest
	GetBody() *UpdateIdentityProviderRequestBody
}

type UpdateIdentityProviderRequest struct {
	// The request body for updating the external identity provider.
	Body *UpdateIdentityProviderRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequest) GetBody() *UpdateIdentityProviderRequestBody {
	return s.Body
}

func (s *UpdateIdentityProviderRequest) SetBody(v *UpdateIdentityProviderRequestBody) *UpdateIdentityProviderRequest {
	s.Body = v
	return s
}

func (s *UpdateIdentityProviderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderRequestBody struct {
	// Specifies whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The new application configuration of the external identity provider. If not specified, the existing configuration remains unchanged.
	Metadata *UpdateIdentityProviderRequestBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// Specifies whether to enable organization member synchronization. After this feature is enabled, the external identity provider synchronizes organization members as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
}

func (s UpdateIdentityProviderRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestBody) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *UpdateIdentityProviderRequestBody) GetMetadata() *UpdateIdentityProviderRequestBodyMetadata {
	return s.Metadata
}

func (s *UpdateIdentityProviderRequestBody) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderRequestBody) SetLoginEnabled(v bool) *UpdateIdentityProviderRequestBody {
	s.LoginEnabled = &v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetMetadata(v *UpdateIdentityProviderRequestBodyMetadata) *UpdateIdentityProviderRequestBody {
	s.Metadata = v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetSyncEnabled(v bool) *UpdateIdentityProviderRequestBody {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderRequestBody) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderRequestBodyMetadata struct {
	// The App ID of the Lark application. This parameter is required when the binding type is Feishu.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The AppKey of the DingTalk application. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// Required. The secret of the external identity provider application. This parameter is write-only and is not returned by query operations.
	//
	// example:
	//
	// example-app-secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// The CorpId of the DingTalk organization. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The data encryption key for event subscriptions. The value must be consistent with the one configured in the external identity provider application. This parameter is write-only and is not returned by query operations.
	//
	// example:
	//
	// example-encrypt-key
	EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
	// The verification token for event subscriptions. The value must be consistent with the one configured in the external identity provider application. This parameter is write-only and is not returned by query operations.
	//
	// example:
	//
	// example-verification-token
	VerificationToken *string `json:"verificationToken,omitempty" xml:"verificationToken,omitempty"`
}

func (s UpdateIdentityProviderRequestBodyMetadata) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestBodyMetadata) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetAppId() *string {
	return s.AppId
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetAppKey() *string {
	return s.AppKey
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetAppSecret() *string {
	return s.AppSecret
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetAppId(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.AppId = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetAppKey(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.AppKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetAppSecret(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.AppSecret = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetCorpId(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.CorpId = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetEncryptKey(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.EncryptKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetVerificationToken(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.VerificationToken = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) Validate() error {
	return dara.Validate(s)
}
