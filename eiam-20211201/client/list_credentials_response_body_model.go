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
	// The list of credentials.
	Credentials []*ListCredentialsResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	// The maximum number of entries per page for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	// The creation time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The credential content.
	CredentialContent *ListCredentialsResponseBodyCredentialsCredentialContent `json:"CredentialContent,omitempty" xml:"CredentialContent,omitempty" type:"Struct"`
	// The creation type of the credential. Valid values:
	//
	// - system_init: Created by the system.
	//
	// - user_custom: Created by the user.
	//
	// example:
	//
	// user_custom
	CredentialCreationType *string `json:"CredentialCreationType,omitempty" xml:"CredentialCreationType,omitempty"`
	CredentialExternalId   *string `json:"CredentialExternalId,omitempty" xml:"CredentialExternalId,omitempty"`
	// The credential ID.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The credential identifier.
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"CredentialIdentifier,omitempty" xml:"CredentialIdentifier,omitempty"`
	// The credential name.
	//
	// example:
	//
	// credential_name
	CredentialName *string `json:"CredentialName,omitempty" xml:"CredentialName,omitempty"`
	// The scenarios label of the credential. Valid values:
	//
	// - llm: large language model.
	//
	// - saas: third-party SaaS service.
	//
	// example:
	//
	// llm
	CredentialScenarioLabel *string `json:"CredentialScenarioLabel,omitempty" xml:"CredentialScenarioLabel,omitempty"`
	CredentialSharingScope  *string `json:"CredentialSharingScope,omitempty" xml:"CredentialSharingScope,omitempty"`
	// The subject ID to which the credential belongs.
	//
	// example:
	//
	// apt_werthgfdsasffxxxxx
	CredentialSubjectId *string `json:"CredentialSubjectId,omitempty" xml:"CredentialSubjectId,omitempty"`
	// The subject type to which the credential belongs. Valid values:
	//
	// - authentication_token_provider: authentication token provider.
	//
	// example:
	//
	// authentication_token_provider
	CredentialSubjectType *string `json:"CredentialSubjectType,omitempty" xml:"CredentialSubjectType,omitempty"`
	// The credential type. Valid values:
	//
	// - api_key: API key authentication credential.
	//
	// - oauth_client: OAuth client authentication credential.
	//
	// example:
	//
	// api_key
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The credential description.
	//
	// example:
	//
	// credential_description
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExclusiveUserId *string `json:"ExclusiveUserId,omitempty" xml:"ExclusiveUserId,omitempty"`
	// The EIAM instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The credential status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - diasbled: Disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time, in UNIX timestamp format. Unit: milliseconds.
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

func (s *ListCredentialsResponseBodyCredentials) GetCredentialExternalId() *string {
	return s.CredentialExternalId
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

func (s *ListCredentialsResponseBodyCredentials) GetCredentialSharingScope() *string {
	return s.CredentialSharingScope
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

func (s *ListCredentialsResponseBodyCredentials) GetExclusiveUserId() *string {
	return s.ExclusiveUserId
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

func (s *ListCredentialsResponseBodyCredentials) SetCredentialExternalId(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialExternalId = &v
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

func (s *ListCredentialsResponseBodyCredentials) SetCredentialSharingScope(v string) *ListCredentialsResponseBodyCredentials {
	s.CredentialSharingScope = &v
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

func (s *ListCredentialsResponseBodyCredentials) SetExclusiveUserId(v string) *ListCredentialsResponseBodyCredentials {
	s.ExclusiveUserId = &v
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
	// The credential content of the OAuth client authentication credential type.
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
	// The client_id of the OAuth protocol.
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
