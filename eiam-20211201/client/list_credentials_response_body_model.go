// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredentials(v []*ListCredentialsResponseBodyCredentials) *ListCredentialsResponseBody
	GetCredentials() []*ListCredentialsResponseBodyCredentials
	SetMaxResults(v int32) *ListCredentialsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCredentialsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCredentialsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCredentialsResponseBody
	GetTotalCount() *int64
}

type ListCredentialsResponseBody struct {
	Credentials []*ListCredentialsResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	// 分页查询时每页行数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBody) GetCredentials() []*ListCredentialsResponseBodyCredentials {
	return s.Credentials
}

func (s *ListCredentialsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCredentialsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCredentialsResponseBody) SetCredentials(v []*ListCredentialsResponseBodyCredentials) *ListCredentialsResponseBody {
	s.Credentials = v
	return s
}

func (s *ListCredentialsResponseBody) SetMaxResults(v int32) *ListCredentialsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialsResponseBody) SetNextToken(v string) *ListCredentialsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCredentialsResponseBody) SetRequestId(v string) *ListCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResponseBody) SetTotalCount(v int64) *ListCredentialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCredentialsResponseBody) Validate() error {
	if s.Credentials != nil {
		for _, item := range s.Credentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCredentialsResponseBodyCredentials struct {
	// 云角色创建时间
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 凭据的内容。
	CredentialContent *ListCredentialsResponseBodyCredentialsCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
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

func (s ListCredentialsResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyCredentials) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialContent() *ListCredentialsResponseBodyCredentialsCredentialContent {
	return s.CredentialContent
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialCreationType() *string {
	return s.CredentialCreationType
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialScenarioLabel() *string {
	return s.CredentialScenarioLabel
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialSubjectId() *string {
	return s.CredentialSubjectId
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialSubjectType() *string {
	return s.CredentialSubjectType
}

func (s *ListCredentialsResponseBodyCredentials) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ListCredentialsResponseBodyCredentials) GetDescription() *string {
	return s.Description
}

func (s *ListCredentialsResponseBodyCredentials) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCredentialsResponseBodyCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListCredentialsResponseBodyCredentials) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCredentialsResponseBodyCredentials) SetCreateTime(v int64) *ListCredentialsResponseBodyCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialContent(v *ListCredentialsResponseBodyCredentialsCredentialContent) *ListCredentialsResponseBodyCredentials {
	s.CredentialContent = v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialCreationType(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialCreationType = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialId(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialId = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialIdentifier(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialIdentifier = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialName(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialName = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialScenarioLabel(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialScenarioLabel = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialSubjectId(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialSubjectId = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialSubjectType(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialSubjectType = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetCredentialType(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialType = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetDescription(v string) *ListCredentialsResponseBodyCredentials {
	s.Description = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetInstanceId(v string) *ListCredentialsResponseBodyCredentials {
	s.InstanceId = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetStatus(v string) *ListCredentialsResponseBodyCredentials {
	s.Status = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) SetUpdateTime(v int64) *ListCredentialsResponseBodyCredentials {
	s.UpdateTime = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentials) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCredentialsResponseBodyCredentialsCredentialContent struct {
	// OAuth客户端认证凭证类型的凭据内容。
	OAuthClientContent *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent `json:"OAuthClientContent,omitempty" xml:"OAuthClientContent,omitempty" type:"Struct"`
}

func (s ListCredentialsResponseBodyCredentialsCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyCredentialsCredentialContent) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyCredentialsCredentialContent) GetOAuthClientContent() *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent {
	return s.OAuthClientContent
}

func (s *ListCredentialsResponseBodyCredentialsCredentialContent) SetOAuthClientContent(v *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent) *ListCredentialsResponseBodyCredentialsCredentialContent {
	s.OAuthClientContent = v
	return s
}

func (s *ListCredentialsResponseBodyCredentialsCredentialContent) Validate() error {
	if s.OAuthClientContent != nil {
		if err := s.OAuthClientContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent struct {
	// OAuth协议的client_id
	//
	// example:
	//
	// dmvncmxersdxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent) GetClientId() *string {
	return s.ClientId
}

func (s *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent) SetClientId(v string) *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent {
	s.ClientId = &v
	return s
}

func (s *ListCredentialsResponseBodyCredentialsCredentialContentOAuthClientContent) Validate() error {
	return dara.Validate(s)
}
