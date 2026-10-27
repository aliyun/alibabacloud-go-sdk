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
	// The request content for binding an external identity provider.
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
	// The type of the external identity provider. Valid values: DingTalk and Feishu.
	//
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// Specifies whether to allow workspace users to log on through the external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. When you bind DingTalk, you must provide appKey, appSecret, and corpId. When you bind Feishu, you must provide appId and appSecret.
	Metadata *CreateIdentityProviderRequestBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// Specifies whether to enable organization member synchronization. After this feature is enabled, the external identity provider synchronizes organization members as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// Controls the first synchronization when DingTalk or Feishu is bound. If this parameter is omitted or set to null, the default value is true. If this parameter is set to false, only the configuration is saved, and synchronization is explicitly triggered later. CustomSSO does not use the vendor pull link.
	SyncOnBind *bool `json:"syncOnBind,omitempty" xml:"syncOnBind,omitempty"`
	// Retains the department-based synchronization scope definition. If this parameter is omitted during creation, the default value is AllAuthorized, and the first vendor synchronization is determined by syncOnBind. If this parameter is omitted or set to null during an update, the original value is retained. Updating and saving the scope does not automatically perform full synchronization.
	SyncScope *CreateIdentityProviderRequestBodySyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
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

func (s *CreateIdentityProviderRequestBody) GetSyncOnBind() *bool {
	return s.SyncOnBind
}

func (s *CreateIdentityProviderRequestBody) GetSyncScope() *CreateIdentityProviderRequestBodySyncScope {
	return s.SyncScope
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

func (s *CreateIdentityProviderRequestBody) SetSyncOnBind(v bool) *CreateIdentityProviderRequestBody {
	s.SyncOnBind = &v
	return s
}

func (s *CreateIdentityProviderRequestBody) SetSyncScope(v *CreateIdentityProviderRequestBodySyncScope) *CreateIdentityProviderRequestBody {
	s.SyncScope = v
	return s
}

func (s *CreateIdentityProviderRequestBody) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.SyncScope != nil {
		if err := s.SyncScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderRequestBodyMetadata struct {
	// The app ID of the Feishu application. This parameter is required when the binding type is Feishu.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The appKey of the DingTalk application. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// Required. The secret of the external identity provider application. This parameter is used only for writing and is not returned by the query API.
	//
	// example:
	//
	// example-app-secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// The corpId of the DingTalk enterprise. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The data encryption key for event subscription. The value must be the same as the one configured in the external identity provider application. This parameter is used only for writing and is not returned by the query API.
	//
	// example:
	//
	// example-encrypt-key
	EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
	// The verification token for event subscription. The value must be the same as the one configured in the external identity provider application. This parameter is used only for writing and is not returned by the query API.
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

type CreateIdentityProviderRequestBodySyncScope struct {
	// When the mode is SelectedDepartments, at least one non-empty department ID must be included, which should be obtained from ListIdentityProviderDepartments. When the mode is AllAuthorized, pass an empty array or omit this parameter.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// Valid values:
	//
	// - AllAuthorized: The full authorization scope, which is limited by the application address book authorization.
	//
	// - SelectedDepartments: The selected departments and all their sub-departments.
	//
	// When you explicitly submit syncScope, you must provide a valid mode, which is verified by the backend.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s CreateIdentityProviderRequestBodySyncScope) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestBodySyncScope) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestBodySyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *CreateIdentityProviderRequestBodySyncScope) GetMode() *string {
	return s.Mode
}

func (s *CreateIdentityProviderRequestBodySyncScope) SetDepartmentIds(v []*string) *CreateIdentityProviderRequestBodySyncScope {
	s.DepartmentIds = v
	return s
}

func (s *CreateIdentityProviderRequestBodySyncScope) SetMode(v string) *CreateIdentityProviderRequestBodySyncScope {
	s.Mode = &v
	return s
}

func (s *CreateIdentityProviderRequestBodySyncScope) Validate() error {
	return dara.Validate(s)
}
