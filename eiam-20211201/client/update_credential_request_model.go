// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCredentialRequest
	GetClientToken() *string
	SetCredentialContent(v *UpdateCredentialRequestCredentialContent) *UpdateCredentialRequest
	GetCredentialContent() *UpdateCredentialRequestCredentialContent
	SetCredentialId(v string) *UpdateCredentialRequest
	GetCredentialId() *string
	SetCredentialName(v string) *UpdateCredentialRequest
	GetCredentialName() *string
	SetInstanceId(v string) *UpdateCredentialRequest
	GetInstanceId() *string
}

type UpdateCredentialRequest struct {
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 凭据的内容。
	CredentialContent *UpdateCredentialRequestCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
	// 凭据ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// 凭据名称。
	//
	// example:
	//
	// credential_name
	CredentialName *string `json:"CredentialName,omitempty" xml:"CredentialName,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialRequest) GetCredentialContent() *UpdateCredentialRequestCredentialContent {
	return s.CredentialContent
}

func (s *UpdateCredentialRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateCredentialRequest) GetCredentialName() *string {
	return s.CredentialName
}

func (s *UpdateCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCredentialRequest) SetClientToken(v string) *UpdateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialRequest) SetCredentialContent(v *UpdateCredentialRequestCredentialContent) *UpdateCredentialRequest {
	s.CredentialContent = v
	return s
}

func (s *UpdateCredentialRequest) SetCredentialId(v string) *UpdateCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *UpdateCredentialRequest) SetCredentialName(v string) *UpdateCredentialRequest {
	s.CredentialName = &v
	return s
}

func (s *UpdateCredentialRequest) SetInstanceId(v string) *UpdateCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCredentialRequest) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialRequestCredentialContent struct {
	// Api Key的内容。
	ApiKeyContent *UpdateCredentialRequestCredentialContentApiKeyContent `json:"ApiKeyContent,omitempty" xml:"ApiKeyContent,omitempty" type:"Struct"`
	// OAuth客户端认证凭证类型的凭据内容。
	OAuthClientContent *UpdateCredentialRequestCredentialContentOAuthClientContent `json:"OAuthClientContent,omitempty" xml:"OAuthClientContent,omitempty" type:"Struct"`
}

func (s UpdateCredentialRequestCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequestCredentialContent) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequestCredentialContent) GetApiKeyContent() *UpdateCredentialRequestCredentialContentApiKeyContent {
	return s.ApiKeyContent
}

func (s *UpdateCredentialRequestCredentialContent) GetOAuthClientContent() *UpdateCredentialRequestCredentialContentOAuthClientContent {
	return s.OAuthClientContent
}

func (s *UpdateCredentialRequestCredentialContent) SetApiKeyContent(v *UpdateCredentialRequestCredentialContentApiKeyContent) *UpdateCredentialRequestCredentialContent {
	s.ApiKeyContent = v
	return s
}

func (s *UpdateCredentialRequestCredentialContent) SetOAuthClientContent(v *UpdateCredentialRequestCredentialContentOAuthClientContent) *UpdateCredentialRequestCredentialContent {
	s.OAuthClientContent = v
	return s
}

func (s *UpdateCredentialRequestCredentialContent) Validate() error {
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

type UpdateCredentialRequestCredentialContentApiKeyContent struct {
	// API Key 凭证类型的凭据内容。
	//
	// example:
	//
	// nsklnertyt5ddwizncxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
}

func (s UpdateCredentialRequestCredentialContentApiKeyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequestCredentialContentApiKeyContent) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequestCredentialContentApiKeyContent) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateCredentialRequestCredentialContentApiKeyContent) SetApiKey(v string) *UpdateCredentialRequestCredentialContentApiKeyContent {
	s.ApiKey = &v
	return s
}

func (s *UpdateCredentialRequestCredentialContentApiKeyContent) Validate() error {
	return dara.Validate(s)
}

type UpdateCredentialRequestCredentialContentOAuthClientContent struct {
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

func (s UpdateCredentialRequestCredentialContentOAuthClientContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequestCredentialContentOAuthClientContent) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequestCredentialContentOAuthClientContent) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateCredentialRequestCredentialContentOAuthClientContent) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateCredentialRequestCredentialContentOAuthClientContent) SetClientId(v string) *UpdateCredentialRequestCredentialContentOAuthClientContent {
	s.ClientId = &v
	return s
}

func (s *UpdateCredentialRequestCredentialContentOAuthClientContent) SetClientSecret(v string) *UpdateCredentialRequestCredentialContentOAuthClientContent {
	s.ClientSecret = &v
	return s
}

func (s *UpdateCredentialRequestCredentialContentOAuthClientContent) Validate() error {
	return dara.Validate(s)
}
