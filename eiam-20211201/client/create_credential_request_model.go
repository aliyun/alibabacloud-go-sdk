// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCredentialRequest
	GetClientToken() *string
	SetCredentialContent(v *CreateCredentialRequestCredentialContent) *CreateCredentialRequest
	GetCredentialContent() *CreateCredentialRequestCredentialContent
	SetCredentialIdentifier(v string) *CreateCredentialRequest
	GetCredentialIdentifier() *string
	SetCredentialName(v string) *CreateCredentialRequest
	GetCredentialName() *string
	SetCredentialScenarioLabel(v string) *CreateCredentialRequest
	GetCredentialScenarioLabel() *string
	SetCredentialSubjectId(v string) *CreateCredentialRequest
	GetCredentialSubjectId() *string
	SetCredentialSubjectType(v string) *CreateCredentialRequest
	GetCredentialSubjectType() *string
	SetCredentialType(v string) *CreateCredentialRequest
	GetCredentialType() *string
	SetDescription(v string) *CreateCredentialRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCredentialRequest
	GetInstanceId() *string
}

type CreateCredentialRequest struct {
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 凭据的内容。
	//
	// This parameter is required.
	CredentialContent *CreateCredentialRequestCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
	// 凭据标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"CredentialIdentifier,omitempty" xml:"CredentialIdentifier,omitempty"`
	// 凭据名称。
	//
	// This parameter is required.
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
	// This parameter is required.
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
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCredentialRequest) GetCredentialContent() *CreateCredentialRequestCredentialContent {
	return s.CredentialContent
}

func (s *CreateCredentialRequest) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *CreateCredentialRequest) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateCredentialRequest) GetCredentialScenarioLabel() *string {
	return s.CredentialScenarioLabel
}

func (s *CreateCredentialRequest) GetCredentialSubjectId() *string {
	return s.CredentialSubjectId
}

func (s *CreateCredentialRequest) GetCredentialSubjectType() *string {
	return s.CredentialSubjectType
}

func (s *CreateCredentialRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateCredentialRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCredentialRequest) SetClientToken(v string) *CreateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCredentialRequest) SetCredentialContent(v *CreateCredentialRequestCredentialContent) *CreateCredentialRequest {
	s.CredentialContent = v
	return s
}

func (s *CreateCredentialRequest) SetCredentialIdentifier(v string) *CreateCredentialRequest {
	s.CredentialIdentifier = &v
	return s
}

func (s *CreateCredentialRequest) SetCredentialName(v string) *CreateCredentialRequest {
	s.CredentialName = &v
	return s
}

func (s *CreateCredentialRequest) SetCredentialScenarioLabel(v string) *CreateCredentialRequest {
	s.CredentialScenarioLabel = &v
	return s
}

func (s *CreateCredentialRequest) SetCredentialSubjectId(v string) *CreateCredentialRequest {
	s.CredentialSubjectId = &v
	return s
}

func (s *CreateCredentialRequest) SetCredentialSubjectType(v string) *CreateCredentialRequest {
	s.CredentialSubjectType = &v
	return s
}

func (s *CreateCredentialRequest) SetCredentialType(v string) *CreateCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *CreateCredentialRequest) SetDescription(v string) *CreateCredentialRequest {
	s.Description = &v
	return s
}

func (s *CreateCredentialRequest) SetInstanceId(v string) *CreateCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCredentialRequest) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialRequestCredentialContent struct {
	// Api Key的内容。
	ApiKeyContent *CreateCredentialRequestCredentialContentApiKeyContent `json:"ApiKeyContent,omitempty" xml:"ApiKeyContent,omitempty" type:"Struct"`
	// OAuth客户端认证凭证类型的凭据内容。
	OAuthClientContent *CreateCredentialRequestCredentialContentOAuthClientContent `json:"OAuthClientContent,omitempty" xml:"OAuthClientContent,omitempty" type:"Struct"`
}

func (s CreateCredentialRequestCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequestCredentialContent) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequestCredentialContent) GetApiKeyContent() *CreateCredentialRequestCredentialContentApiKeyContent {
	return s.ApiKeyContent
}

func (s *CreateCredentialRequestCredentialContent) GetOAuthClientContent() *CreateCredentialRequestCredentialContentOAuthClientContent {
	return s.OAuthClientContent
}

func (s *CreateCredentialRequestCredentialContent) SetApiKeyContent(v *CreateCredentialRequestCredentialContentApiKeyContent) *CreateCredentialRequestCredentialContent {
	s.ApiKeyContent = v
	return s
}

func (s *CreateCredentialRequestCredentialContent) SetOAuthClientContent(v *CreateCredentialRequestCredentialContentOAuthClientContent) *CreateCredentialRequestCredentialContent {
	s.OAuthClientContent = v
	return s
}

func (s *CreateCredentialRequestCredentialContent) Validate() error {
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

type CreateCredentialRequestCredentialContentApiKeyContent struct {
	// API Key 凭证类型的凭据内容。
	//
	// example:
	//
	// nsklnertyt5ddwizncxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
}

func (s CreateCredentialRequestCredentialContentApiKeyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequestCredentialContentApiKeyContent) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequestCredentialContentApiKeyContent) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateCredentialRequestCredentialContentApiKeyContent) SetApiKey(v string) *CreateCredentialRequestCredentialContentApiKeyContent {
	s.ApiKey = &v
	return s
}

func (s *CreateCredentialRequestCredentialContentApiKeyContent) Validate() error {
	return dara.Validate(s)
}

type CreateCredentialRequestCredentialContentOAuthClientContent struct {
	// OAuth协议的client_id。
	//
	// example:
	//
	// dmvncmxersdxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// OAuth协议的client_secret。
	//
	// example:
	//
	// nsklncmwizncxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s CreateCredentialRequestCredentialContentOAuthClientContent) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequestCredentialContentOAuthClientContent) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequestCredentialContentOAuthClientContent) GetClientId() *string {
	return s.ClientId
}

func (s *CreateCredentialRequestCredentialContentOAuthClientContent) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateCredentialRequestCredentialContentOAuthClientContent) SetClientId(v string) *CreateCredentialRequestCredentialContentOAuthClientContent {
	s.ClientId = &v
	return s
}

func (s *CreateCredentialRequestCredentialContentOAuthClientContent) SetClientSecret(v string) *CreateCredentialRequestCredentialContentOAuthClientContent {
	s.ClientSecret = &v
	return s
}

func (s *CreateCredentialRequestCredentialContentOAuthClientContent) Validate() error {
	return dara.Validate(s)
}
