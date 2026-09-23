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
	// The creation time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The credential content.
	CredentialContent *ObtainCredentialResponseBodyCredentialContent `json:"credentialContent,omitempty" xml:"credentialContent,omitempty" type:"Struct"`
	// The creation type of the credential. Valid values:
	//
	// - system_init: Created by the system.
	//
	// - user_custom: Created by the user.
	//
	// example:
	//
	// user_custom
	CredentialCreationType *string `json:"credentialCreationType,omitempty" xml:"credentialCreationType,omitempty"`
	// The external unique identifier of the credential.
	//
	// example:
	//
	// 23528e9957304f57b98112c72788b5xxxxx
	CredentialExternalId *string `json:"credentialExternalId,omitempty" xml:"credentialExternalId,omitempty"`
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
	// The Scenarios label of the credential. Valid values:
	//
	// - llm: Large language model.
	//
	// - saas: Third-party SaaS service.
	//
	// example:
	//
	// llm
	CredentialScenarioLabel *string `json:"credentialScenarioLabel,omitempty" xml:"credentialScenarioLabel,omitempty"`
	// The credential sharing scope.
	//
	// example:
	//
	// user_exclusive
	CredentialSharingScope *string `json:"credentialSharingScope,omitempty" xml:"credentialSharingScope,omitempty"`
	// The subject ID that the credential belongs to.
	//
	// example:
	//
	// apt_werthgfdsasffxxxxx
	CredentialSubjectId *string `json:"credentialSubjectId,omitempty" xml:"credentialSubjectId,omitempty"`
	// The subject type that the credential belongs to. Valid values:
	//
	// - authentication_token_provider: Authentication token provider.
	//
	// example:
	//
	// authentication_token_provider
	CredentialSubjectType *string `json:"credentialSubjectType,omitempty" xml:"credentialSubjectType,omitempty"`
	// The credential type. Valid values:
	//
	// - api_key: API Key authentication credential.
	//
	// - oauth_client: OAuth client authentication credential.
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
	// The exclusive account ID of the credential.
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
	// The credential status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time, in UNIX timestamp format. Unit: milliseconds.
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
	// The credential content of the API Key credential type.
	ApiKeyContent *ObtainCredentialResponseBodyCredentialContentApiKeyContent `json:"apiKeyContent,omitempty" xml:"apiKeyContent,omitempty" type:"Struct"`
	// The credential content of the OAuth client authentication credential type.
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
	// The value of the API Key.
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
	// The client_id of the OAuth protocol.
	//
	// example:
	//
	// dmvncmxersdxxxxxx
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The client_secret of the OAuth protocol.
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
