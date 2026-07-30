// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *ObtainCredentialResponseBody
	GetCreateTime() *int64
	SetCredentialContent(v *ObtainCredentialResponseBodyCredentialContent) *ObtainCredentialResponseBody
	GetCredentialContent() *ObtainCredentialResponseBodyCredentialContent
	SetCredentialCreationType(v string) *ObtainCredentialResponseBody
	GetCredentialCreationType() *string
	SetCredentialExternalId(v string) *ObtainCredentialResponseBody
	GetCredentialExternalId() *string
	SetCredentialId(v string) *ObtainCredentialResponseBody
	GetCredentialId() *string
	SetCredentialIdentifier(v string) *ObtainCredentialResponseBody
	GetCredentialIdentifier() *string
	SetCredentialName(v string) *ObtainCredentialResponseBody
	GetCredentialName() *string
	SetCredentialScenarioLabel(v string) *ObtainCredentialResponseBody
	GetCredentialScenarioLabel() *string
	SetCredentialSharingScope(v string) *ObtainCredentialResponseBody
	GetCredentialSharingScope() *string
	SetCredentialSubjectId(v string) *ObtainCredentialResponseBody
	GetCredentialSubjectId() *string
	SetCredentialSubjectType(v string) *ObtainCredentialResponseBody
	GetCredentialSubjectType() *string
	SetCredentialType(v string) *ObtainCredentialResponseBody
	GetCredentialType() *string
	SetDescription(v string) *ObtainCredentialResponseBody
	GetDescription() *string
	SetExclusiveUserId(v string) *ObtainCredentialResponseBody
	GetExclusiveUserId() *string
	SetInstanceId(v string) *ObtainCredentialResponseBody
	GetInstanceId() *string
	SetStatus(v string) *ObtainCredentialResponseBody
	GetStatus() *string
	SetUpdateTime(v int64) *ObtainCredentialResponseBody
	GetUpdateTime() *int64
}

type ObtainCredentialResponseBody struct {
	// The creation time of the credential, formatted as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The detailed content of the credential. The structure of this object depends on the value of `credentialType`.
	CredentialContent *ObtainCredentialResponseBodyCredentialContent `json:"credentialContent,omitempty" xml:"credentialContent,omitempty" type:"Struct"`
	// Indicates how the credential was created. Valid values:
	//
	// - `system_init`: System-initiated.
	//
	// - `user_custom`: User-created.
	//
	// example:
	//
	// user_custom
	CredentialCreationType *string `json:"credentialCreationType,omitempty" xml:"credentialCreationType,omitempty"`
	CredentialExternalId   *string `json:"credentialExternalId,omitempty" xml:"credentialExternalId,omitempty"`
	// The credential ID.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The credential identifier.
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"credentialIdentifier,omitempty" xml:"credentialIdentifier,omitempty"`
	// The credential name.
	//
	// example:
	//
	// credential_name
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// The usage scenario for the credential. Valid values:
	//
	// - `llm`: For use with a large language model.
	//
	// - `saas`: For use with a third-party SaaS application.
	//
	// example:
	//
	// llm
	CredentialScenarioLabel *string `json:"credentialScenarioLabel,omitempty" xml:"credentialScenarioLabel,omitempty"`
	// The sharing scope of the credential, such as whether it is exclusive to a specific account.
	//
	// example:
	//
	// user_exclusive
	CredentialSharingScope *string `json:"credentialSharingScope,omitempty" xml:"credentialSharingScope,omitempty"`
	// The ID of the credential\\"s subject.
	//
	// example:
	//
	// apt_werthgfdsasffxxxxx
	CredentialSubjectId *string `json:"credentialSubjectId,omitempty" xml:"credentialSubjectId,omitempty"`
	// The credential\\"s subject type. Valid values:
	//
	// - `authentication_token_provider`: An authentication token provider.
	//
	// example:
	//
	// authentication_token_provider
	CredentialSubjectType *string `json:"credentialSubjectType,omitempty" xml:"credentialSubjectType,omitempty"`
	// The credential type. Valid values:
	//
	// - `api_key`: The credential is an API key.
	//
	// - `oauth_client`: The credential represents an OAuth client.
	//
	// example:
	//
	// api_key
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// The credential description.
	//
	// example:
	//
	// credential_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the account that exclusively owns the credential. This field is present only when `credentialSharingScope` is `user_exclusive`.
	//
	// example:
	//
	// user_xxx
	ExclusiveUserId *string `json:"exclusiveUserId,omitempty" xml:"exclusiveUserId,omitempty"`
	// The EIAM instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The status of the credential. Valid values:
	//
	// - `enabled`: The credential can be used.
	//
	// - `disabled`: The credential cannot be used.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last update time of the credential, formatted as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ObtainCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ObtainCredentialResponseBody) GetCredentialContent() *ObtainCredentialResponseBodyCredentialContent {
	return s.CredentialContent
}

func (s *ObtainCredentialResponseBody) GetCredentialCreationType() *string {
	return s.CredentialCreationType
}

func (s *ObtainCredentialResponseBody) GetCredentialExternalId() *string {
	return s.CredentialExternalId
}

func (s *ObtainCredentialResponseBody) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ObtainCredentialResponseBody) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *ObtainCredentialResponseBody) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ObtainCredentialResponseBody) GetCredentialScenarioLabel() *string {
	return s.CredentialScenarioLabel
}

func (s *ObtainCredentialResponseBody) GetCredentialSharingScope() *string {
	return s.CredentialSharingScope
}

func (s *ObtainCredentialResponseBody) GetCredentialSubjectId() *string {
	return s.CredentialSubjectId
}

func (s *ObtainCredentialResponseBody) GetCredentialSubjectType() *string {
	return s.CredentialSubjectType
}

func (s *ObtainCredentialResponseBody) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ObtainCredentialResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ObtainCredentialResponseBody) GetExclusiveUserId() *string {
	return s.ExclusiveUserId
}

func (s *ObtainCredentialResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainCredentialResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ObtainCredentialResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ObtainCredentialResponseBody) SetCreateTime(v int64) *ObtainCredentialResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialContent(v *ObtainCredentialResponseBodyCredentialContent) *ObtainCredentialResponseBody {
	s.CredentialContent = v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialCreationType(v string) *ObtainCredentialResponseBody {
	s.CredentialCreationType = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialExternalId(v string) *ObtainCredentialResponseBody {
	s.CredentialExternalId = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialId(v string) *ObtainCredentialResponseBody {
	s.CredentialId = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialIdentifier(v string) *ObtainCredentialResponseBody {
	s.CredentialIdentifier = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialName(v string) *ObtainCredentialResponseBody {
	s.CredentialName = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialScenarioLabel(v string) *ObtainCredentialResponseBody {
	s.CredentialScenarioLabel = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialSharingScope(v string) *ObtainCredentialResponseBody {
	s.CredentialSharingScope = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialSubjectId(v string) *ObtainCredentialResponseBody {
	s.CredentialSubjectId = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialSubjectType(v string) *ObtainCredentialResponseBody {
	s.CredentialSubjectType = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetCredentialType(v string) *ObtainCredentialResponseBody {
	s.CredentialType = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetDescription(v string) *ObtainCredentialResponseBody {
	s.Description = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetExclusiveUserId(v string) *ObtainCredentialResponseBody {
	s.ExclusiveUserId = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetInstanceId(v string) *ObtainCredentialResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetStatus(v string) *ObtainCredentialResponseBody {
	s.Status = &v
	return s
}

func (s *ObtainCredentialResponseBody) SetUpdateTime(v int64) *ObtainCredentialResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *ObtainCredentialResponseBody) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCredentialResponseBodyCredentialContent struct {
	// Contains details for an API key credential. Returned only when `credentialType` is `api_key`.
	ApiKeyContent *ObtainCredentialResponseBodyCredentialContentApiKeyContent `json:"apiKeyContent,omitempty" xml:"apiKeyContent,omitempty" type:"Struct"`
	// Contains details for an OAuth client credential. Returned only when `credentialType` is `oauth_client`.
	OauthClientContent *ObtainCredentialResponseBodyCredentialContentOauthClientContent `json:"oauthClientContent,omitempty" xml:"oauthClientContent,omitempty" type:"Struct"`
}

func (s ObtainCredentialResponseBodyCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredentialContent) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredentialContent) GetApiKeyContent() *ObtainCredentialResponseBodyCredentialContentApiKeyContent {
	return s.ApiKeyContent
}

func (s *ObtainCredentialResponseBodyCredentialContent) GetOauthClientContent() *ObtainCredentialResponseBodyCredentialContentOauthClientContent {
	return s.OauthClientContent
}

func (s *ObtainCredentialResponseBodyCredentialContent) SetApiKeyContent(v *ObtainCredentialResponseBodyCredentialContentApiKeyContent) *ObtainCredentialResponseBodyCredentialContent {
	s.ApiKeyContent = v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialContent) SetOauthClientContent(v *ObtainCredentialResponseBodyCredentialContentOauthClientContent) *ObtainCredentialResponseBodyCredentialContent {
	s.OauthClientContent = v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialContent) Validate() error {
	if s.ApiKeyContent != nil {
		if err := s.ApiKeyContent.Validate(); err != nil {
			return err
		}
	}
	if s.OauthClientContent != nil {
		if err := s.OauthClientContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCredentialResponseBodyCredentialContentApiKeyContent struct {
	// The API key value.
	//
	// example:
	//
	// sk-nsklncmwizncxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
}

func (s ObtainCredentialResponseBodyCredentialContentApiKeyContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredentialContentApiKeyContent) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredentialContentApiKeyContent) GetApiKey() *string {
	return s.ApiKey
}

func (s *ObtainCredentialResponseBodyCredentialContentApiKeyContent) SetApiKey(v string) *ObtainCredentialResponseBodyCredentialContentApiKeyContent {
	s.ApiKey = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialContentApiKeyContent) Validate() error {
	return dara.Validate(s)
}

type ObtainCredentialResponseBodyCredentialContentOauthClientContent struct {
	// The `client_id` for OAuth 2.0.
	//
	// example:
	//
	// dmvncmxersdxxxxxx
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The `client_secret` for OAuth 2.0.
	//
	// example:
	//
	// nsklnertyt5ddwizncxxxx
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
}

func (s ObtainCredentialResponseBodyCredentialContentOauthClientContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredentialContentOauthClientContent) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredentialContentOauthClientContent) GetClientId() *string {
	return s.ClientId
}

func (s *ObtainCredentialResponseBodyCredentialContentOauthClientContent) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ObtainCredentialResponseBodyCredentialContentOauthClientContent) SetClientId(v string) *ObtainCredentialResponseBodyCredentialContentOauthClientContent {
	s.ClientId = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialContentOauthClientContent) SetClientSecret(v string) *ObtainCredentialResponseBodyCredentialContentOauthClientContent {
	s.ClientSecret = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialContentOauthClientContent) Validate() error {
	return dara.Validate(s)
}
