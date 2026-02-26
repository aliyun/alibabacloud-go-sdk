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
	Credential *GetCredentialResponseBodyCredential `json:"Credential,omitempty" xml:"Credential,omitempty" type:"Struct"`
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
	// 云角色创建时间
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 凭据的内容。
	CredentialContent *GetCredentialResponseBodyCredentialCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
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
	// OAuth客户端认证凭证类型的凭据内容。
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
	// OAuth协议的client_id
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
