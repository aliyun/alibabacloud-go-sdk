// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredential(v *ObtainCredentialResponseBodyCredential) *ObtainCredentialResponseBody
	GetCredential() *ObtainCredentialResponseBodyCredential
	SetRequestId(v string) *ObtainCredentialResponseBody
	GetRequestId() *string
}

type ObtainCredentialResponseBody struct {
	Credential *ObtainCredentialResponseBodyCredential `json:"Credential,omitempty" xml:"Credential,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ObtainCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBody) GetCredential() *ObtainCredentialResponseBodyCredential {
	return s.Credential
}

func (s *ObtainCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ObtainCredentialResponseBody) SetCredential(v *ObtainCredentialResponseBodyCredential) *ObtainCredentialResponseBody {
	s.Credential = v
	return s
}

func (s *ObtainCredentialResponseBody) SetRequestId(v string) *ObtainCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *ObtainCredentialResponseBody) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCredentialResponseBodyCredential struct {
	// 云角色创建时间
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 凭据的内容。
	CredentialContent *ObtainCredentialResponseBodyCredentialCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
	// 凭据的创建类型。
	//
	// example:
	//
	// user_custom
	CredentialCreationType *string `json:"CredentialCreationType,omitempty" xml:"CredentialCreationType,omitempty"`
	// 凭据ID。
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// 凭据标识
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"CredentialIdentifier,omitempty" xml:"CredentialIdentifier,omitempty"`
	// 凭据名称
	//
	// example:
	//
	// credential_name
	CredentialName *string `json:"CredentialName,omitempty" xml:"CredentialName,omitempty"`
	// 凭据的使用场景标签。
	//
	// example:
	//
	// llm
	CredentialScenarioLabel *string `json:"CredentialScenarioLabel,omitempty" xml:"CredentialScenarioLabel,omitempty"`
	// 凭据所属的主体ID。
	//
	// example:
	//
	// apt_werthgfdsasffxxxxx
	CredentialSubjectId *string `json:"CredentialSubjectId,omitempty" xml:"CredentialSubjectId,omitempty"`
	// 凭据所属的主体类型。
	//
	// example:
	//
	// authentication_token_provider
	CredentialSubjectType *string `json:"CredentialSubjectType,omitempty" xml:"CredentialSubjectType,omitempty"`
	// 凭据类型。
	//
	// example:
	//
	// api_key
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// 描述
	//
	// example:
	//
	// credential_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// EIAM实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 凭据状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 云角色更新时间
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ObtainCredentialResponseBodyCredential) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredential) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialContent() *ObtainCredentialResponseBodyCredentialCredentialContent {
	return s.CredentialContent
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialCreationType() *string {
	return s.CredentialCreationType
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialScenarioLabel() *string {
	return s.CredentialScenarioLabel
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialSubjectId() *string {
	return s.CredentialSubjectId
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialSubjectType() *string {
	return s.CredentialSubjectType
}

func (s *ObtainCredentialResponseBodyCredential) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ObtainCredentialResponseBodyCredential) GetDescription() *string {
	return s.Description
}

func (s *ObtainCredentialResponseBodyCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainCredentialResponseBodyCredential) GetStatus() *string {
	return s.Status
}

func (s *ObtainCredentialResponseBodyCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ObtainCredentialResponseBodyCredential) SetCreateTime(v int64) *ObtainCredentialResponseBodyCredential {
	s.CreateTime = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialContent(v *ObtainCredentialResponseBodyCredentialCredentialContent) *ObtainCredentialResponseBodyCredential {
	s.CredentialContent = v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialCreationType(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialCreationType = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialId(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialId = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialIdentifier(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialIdentifier = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialName(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialName = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialScenarioLabel(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialScenarioLabel = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialSubjectId(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialSubjectId = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialSubjectType(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialSubjectType = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetCredentialType(v string) *ObtainCredentialResponseBodyCredential {
	s.CredentialType = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetDescription(v string) *ObtainCredentialResponseBodyCredential {
	s.Description = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetInstanceId(v string) *ObtainCredentialResponseBodyCredential {
	s.InstanceId = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetStatus(v string) *ObtainCredentialResponseBodyCredential {
	s.Status = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) SetUpdateTime(v int64) *ObtainCredentialResponseBodyCredential {
	s.UpdateTime = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredential) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCredentialResponseBodyCredentialCredentialContent struct {
	ApiKeyContent *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent `json:"ApiKeyContent,omitempty" xml:"ApiKeyContent,omitempty" type:"Struct"`
	// OAuth客户端认证凭证类型的凭据内容。
	OAuthClientContent *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent `json:"OAuthClientContent,omitempty" xml:"OAuthClientContent,omitempty" type:"Struct"`
}

func (s ObtainCredentialResponseBodyCredentialCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredentialCredentialContent) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContent) GetApiKeyContent() *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent {
	return s.ApiKeyContent
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContent) GetOAuthClientContent() *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent {
	return s.OAuthClientContent
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContent) SetApiKeyContent(v *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent) *ObtainCredentialResponseBodyCredentialCredentialContent {
	s.ApiKeyContent = v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContent) SetOAuthClientContent(v *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) *ObtainCredentialResponseBodyCredentialCredentialContent {
	s.OAuthClientContent = v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContent) Validate() error {
	if s.ApiKeyContent != nil {
		if err := s.ApiKeyContent.Validate(); err != nil {
			return err
		}
	}
	if s.OAuthClientContent != nil {
		if err := s.OAuthClientContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent struct {
	// example:
	//
	// nsklncmwizncxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
}

func (s ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent) GetApiKey() *string {
	return s.ApiKey
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent) SetApiKey(v string) *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent {
	s.ApiKey = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentApiKeyContent) Validate() error {
	return dara.Validate(s)
}

type ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent struct {
	// OAuth协议的client_id
	//
	// example:
	//
	// dmvncmxersdxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// nsklnertyt5ddwizncxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) GetClientId() *string {
	return s.ClientId
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) SetClientId(v string) *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent {
	s.ClientId = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) SetClientSecret(v string) *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent {
	s.ClientSecret = &v
	return s
}

func (s *ObtainCredentialResponseBodyCredentialCredentialContentOAuthClientContent) Validate() error {
	return dara.Validate(s)
}
