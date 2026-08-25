// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateIdentityProviderRequestBody) *CreateIdentityProviderRequest
	GetBody() *CreateIdentityProviderRequestBody
}

type CreateIdentityProviderRequest struct {
	// The request body for binding an external identity provider.
	Body *CreateIdentityProviderRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequest) GetBody() *CreateIdentityProviderRequestBody {
	return s.Body
}

func (s *CreateIdentityProviderRequest) SetBody(v *CreateIdentityProviderRequestBody) *CreateIdentityProviderRequest {
	s.Body = v
	return s
}

func (s *CreateIdentityProviderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderRequestBody struct {
	// The type of the external identity provider. Valid values: DingTalk, Feishu.
	//
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// Specifies whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. When binding DingTalk, you must provide appKey, appSecret, and corpId. When binding Lark, you must provide appId and appSecret.
	Metadata *CreateIdentityProviderRequestBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// Specifies whether to enable organization member synchronization. After this feature is enabled, the external identity provider synchronizes organization members as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
}

func (s CreateIdentityProviderRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestBody) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *CreateIdentityProviderRequestBody) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *CreateIdentityProviderRequestBody) GetMetadata() *CreateIdentityProviderRequestBodyMetadata {
	return s.Metadata
}

func (s *CreateIdentityProviderRequestBody) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *CreateIdentityProviderRequestBody) SetIdentityProviderType(v string) *CreateIdentityProviderRequestBody {
	s.IdentityProviderType = &v
	return s
}

func (s *CreateIdentityProviderRequestBody) SetLoginEnabled(v bool) *CreateIdentityProviderRequestBody {
	s.LoginEnabled = &v
	return s
}

func (s *CreateIdentityProviderRequestBody) SetMetadata(v *CreateIdentityProviderRequestBodyMetadata) *CreateIdentityProviderRequestBody {
	s.Metadata = v
	return s
}

func (s *CreateIdentityProviderRequestBody) SetSyncEnabled(v bool) *CreateIdentityProviderRequestBody {
	s.SyncEnabled = &v
	return s
}

func (s *CreateIdentityProviderRequestBody) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderRequestBodyMetadata struct {
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
	// Required. The secret of the external identity provider application. This parameter is used only for write operations. The query API does not return this field.
	//
	// example:
	//
	// example-app-secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// The CorpId of the DingTalk enterprise. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The data encryption key for event subscription. The value must be the same as the one configured in the external identity provider application. This parameter is used only for write operations. The query API does not return this field.
	//
	// example:
	//
	// example-encrypt-key
	EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
	// The verification token for event subscription. The value must be the same as the one configured in the external identity provider application. This parameter is used only for write operations. The query API does not return this field.
	//
	// example:
	//
	// example-verification-token
	VerificationToken *string `json:"verificationToken,omitempty" xml:"verificationToken,omitempty"`
}

func (s CreateIdentityProviderRequestBodyMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestBodyMetadata) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestBodyMetadata) GetAppId() *string {
	return s.AppId
}

func (s *CreateIdentityProviderRequestBodyMetadata) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateIdentityProviderRequestBodyMetadata) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateIdentityProviderRequestBodyMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *CreateIdentityProviderRequestBodyMetadata) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *CreateIdentityProviderRequestBodyMetadata) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *CreateIdentityProviderRequestBodyMetadata) SetAppId(v string) *CreateIdentityProviderRequestBodyMetadata {
	s.AppId = &v
	return s
}

func (s *CreateIdentityProviderRequestBodyMetadata) SetAppKey(v string) *CreateIdentityProviderRequestBodyMetadata {
	s.AppKey = &v
	return s
}

func (s *CreateIdentityProviderRequestBodyMetadata) SetAppSecret(v string) *CreateIdentityProviderRequestBodyMetadata {
	s.AppSecret = &v
	return s
}

func (s *CreateIdentityProviderRequestBodyMetadata) SetCorpId(v string) *CreateIdentityProviderRequestBodyMetadata {
	s.CorpId = &v
	return s
}

func (s *CreateIdentityProviderRequestBodyMetadata) SetEncryptKey(v string) *CreateIdentityProviderRequestBodyMetadata {
	s.EncryptKey = &v
	return s
}

func (s *CreateIdentityProviderRequestBodyMetadata) SetVerificationToken(v string) *CreateIdentityProviderRequestBodyMetadata {
	s.VerificationToken = &v
	return s
}

func (s *CreateIdentityProviderRequestBodyMetadata) Validate() error {
	return dara.Validate(s)
}
