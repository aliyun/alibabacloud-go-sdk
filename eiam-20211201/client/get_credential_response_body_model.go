// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredential(v *GetCredentialResponseBodyCredential) *GetCredentialResponseBody
	GetCredential() *GetCredentialResponseBodyCredential
	SetRequestId(v string) *GetCredentialResponseBody
	GetRequestId() *string
}

type GetCredentialResponseBody struct {
	// The credential details.
	Credential *GetCredentialResponseBodyCredential `json:"Credential,omitempty" xml:"Credential,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBody) GetCredential() *GetCredentialResponseBodyCredential {
	return s.Credential
}

func (s *GetCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialResponseBody) SetCredential(v *GetCredentialResponseBodyCredential) *GetCredentialResponseBody {
	s.Credential = v
	return s
}

func (s *GetCredentialResponseBody) SetRequestId(v string) *GetCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialResponseBody) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialResponseBodyCredential struct {
	// The creation time of the credential, in Unix timestamp format (milliseconds).
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The content of the credential.
	CredentialContent *GetCredentialResponseBodyCredentialCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
	// How the credential was created. Valid values:
	//
	// - `system_init`: The credential was created by the system.
	//
	// - `user_custom`: The credential was created by a user.
	//
	// example:
	//
	// user_custom
	CredentialCreationType *string `json:"CredentialCreationType,omitempty" xml:"CredentialCreationType,omitempty"`
	CredentialExternalId   *string `json:"CredentialExternalId,omitempty" xml:"CredentialExternalId,omitempty"`
	// The ID of the credential.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The identifier of the credential.
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"CredentialIdentifier,omitempty" xml:"CredentialIdentifier,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// credential_name
	CredentialName *string `json:"CredentialName,omitempty" xml:"CredentialName,omitempty"`
	// The use case of the credential. Valid values:
	//
	// - `llm`: a large language model (LLM).
	//
	// - `saas`: a third-party Software as a Service (SaaS) application.
	//
	// example:
	//
	// llm
	CredentialScenarioLabel *string `json:"CredentialScenarioLabel,omitempty" xml:"CredentialScenarioLabel,omitempty"`
	CredentialSharingScope  *string `json:"CredentialSharingScope,omitempty" xml:"CredentialSharingScope,omitempty"`
	// The ID of the subject that owns the credential.
	//
	// example:
	//
	// apt_werthgfdsasffxxxxx
	CredentialSubjectId *string `json:"CredentialSubjectId,omitempty" xml:"CredentialSubjectId,omitempty"`
	// The type of the subject that owns the credential. Valid value:
	//
	// - `authentication_token_provider`: The subject is an authentication token provider.
	//
	// example:
	//
	// authentication_token_provider
	CredentialSubjectType *string `json:"CredentialSubjectType,omitempty" xml:"CredentialSubjectType,omitempty"`
	// The type of the credential. Valid values:
	//
	// - `api_key`: An API key.
	//
	// - `oauth_client`: An OAuth client.
	//
	// example:
	//
	// api_key
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The user-defined description of the credential.
	//
	// example:
	//
	// credential_description
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExclusiveUserId *string `json:"ExclusiveUserId,omitempty" xml:"ExclusiveUserId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the credential. Valid values:
	//
	// - `enabled`: The credential is active.
	//
	// - `disabled`: The credential is inactive.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time the credential was last updated, in Unix timestamp format (milliseconds).
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCredentialResponseBodyCredential) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyCredential) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCredentialResponseBodyCredential) GetCredentialContent() *GetCredentialResponseBodyCredentialCredentialContent {
	return s.CredentialContent
}

func (s *GetCredentialResponseBodyCredential) GetCredentialCreationType() *string {
	return s.CredentialCreationType
}

func (s *GetCredentialResponseBodyCredential) GetCredentialExternalId() *string {
	return s.CredentialExternalId
}

func (s *GetCredentialResponseBodyCredential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetCredentialResponseBodyCredential) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *GetCredentialResponseBodyCredential) GetCredentialName() *string {
	return s.CredentialName
}

func (s *GetCredentialResponseBodyCredential) GetCredentialScenarioLabel() *string {
	return s.CredentialScenarioLabel
}

func (s *GetCredentialResponseBodyCredential) GetCredentialSharingScope() *string {
	return s.CredentialSharingScope
}

func (s *GetCredentialResponseBodyCredential) GetCredentialSubjectId() *string {
	return s.CredentialSubjectId
}

func (s *GetCredentialResponseBodyCredential) GetCredentialSubjectType() *string {
	return s.CredentialSubjectType
}

func (s *GetCredentialResponseBodyCredential) GetCredentialType() *string {
	return s.CredentialType
}

func (s *GetCredentialResponseBodyCredential) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialResponseBodyCredential) GetExclusiveUserId() *string {
	return s.ExclusiveUserId
}

func (s *GetCredentialResponseBodyCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCredentialResponseBodyCredential) GetStatus() *string {
	return s.Status
}

func (s *GetCredentialResponseBodyCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCredentialResponseBodyCredential) SetCreateTime(v int64) *GetCredentialResponseBodyCredential {
	s.CreateTime = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialContent(v *GetCredentialResponseBodyCredentialCredentialContent) *GetCredentialResponseBodyCredential {
	s.CredentialContent = v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialCreationType(v string) *GetCredentialResponseBodyCredential {
	s.CredentialCreationType = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialExternalId(v string) *GetCredentialResponseBodyCredential {
	s.CredentialExternalId = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialId(v string) *GetCredentialResponseBodyCredential {
	s.CredentialId = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialIdentifier(v string) *GetCredentialResponseBodyCredential {
	s.CredentialIdentifier = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialName(v string) *GetCredentialResponseBodyCredential {
	s.CredentialName = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialScenarioLabel(v string) *GetCredentialResponseBodyCredential {
	s.CredentialScenarioLabel = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialSharingScope(v string) *GetCredentialResponseBodyCredential {
	s.CredentialSharingScope = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialSubjectId(v string) *GetCredentialResponseBodyCredential {
	s.CredentialSubjectId = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialSubjectType(v string) *GetCredentialResponseBodyCredential {
	s.CredentialSubjectType = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetCredentialType(v string) *GetCredentialResponseBodyCredential {
	s.CredentialType = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetDescription(v string) *GetCredentialResponseBodyCredential {
	s.Description = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetExclusiveUserId(v string) *GetCredentialResponseBodyCredential {
	s.ExclusiveUserId = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetInstanceId(v string) *GetCredentialResponseBodyCredential {
	s.InstanceId = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetStatus(v string) *GetCredentialResponseBodyCredential {
	s.Status = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) SetUpdateTime(v int64) *GetCredentialResponseBodyCredential {
	s.UpdateTime = &v
	return s
}

func (s *GetCredentialResponseBodyCredential) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialResponseBodyCredentialCredentialContent struct {
	// The credential content for an OAuth client. This parameter is returned only when `CredentialType` is `oauth_client`.
	OAuthClientContent *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent `json:"OAuthClientContent,omitempty" xml:"OAuthClientContent,omitempty" type:"Struct"`
}

func (s GetCredentialResponseBodyCredentialCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyCredentialCredentialContent) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyCredentialCredentialContent) GetOAuthClientContent() *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent {
	return s.OAuthClientContent
}

func (s *GetCredentialResponseBodyCredentialCredentialContent) SetOAuthClientContent(v *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent) *GetCredentialResponseBodyCredentialCredentialContent {
	s.OAuthClientContent = v
	return s
}

func (s *GetCredentialResponseBodyCredentialCredentialContent) Validate() error {
	if s.OAuthClientContent != nil {
		if err := s.OAuthClientContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent struct {
	// The OAuth client ID.
	//
	// example:
	//
	// dmvncmxersdxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent) GetClientId() *string {
	return s.ClientId
}

func (s *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent) SetClientId(v string) *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent {
	s.ClientId = &v
	return s
}

func (s *GetCredentialResponseBodyCredentialCredentialContentOAuthClientContent) Validate() error {
	return dara.Validate(s)
}
