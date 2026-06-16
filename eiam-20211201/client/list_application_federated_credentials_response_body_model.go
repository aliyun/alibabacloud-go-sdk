// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsResponseBody
	GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials
	SetMaxResults(v int32) *ListApplicationFederatedCredentialsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationFederatedCredentialsResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationFederatedCredentialsResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListApplicationFederatedCredentialsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationFederatedCredentialsResponseBody
	GetTotalCount() *int64
}

type ListApplicationFederatedCredentialsResponseBody struct {
	// The list of application federated credentials.
	ApplicationFederatedCredentials []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials `json:"ApplicationFederatedCredentials,omitempty" xml:"ApplicationFederatedCredentials,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The query token returned by this call.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	return s.ApplicationFederatedCredentials
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsResponseBody {
	s.ApplicationFederatedCredentials = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetMaxResults(v int32) *ListApplicationFederatedCredentialsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetNextToken(v string) *ListApplicationFederatedCredentialsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetPreviousToken(v string) *ListApplicationFederatedCredentialsResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetRequestId(v string) *ListApplicationFederatedCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetTotalCount(v int64) *ListApplicationFederatedCredentialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) Validate() error {
	if s.ApplicationFederatedCredentials != nil {
		for _, item := range s.ApplicationFederatedCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials struct {
	// The ID of the application federated credential.
	//
	// example:
	//
	// afc_adsa1sdaxxxxx
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// The name of the application federated credential.
	//
	// example:
	//
	// test
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// The type of the application federated credential.
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app_xxxasda1
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the credential was created.
	//
	// example:
	//
	// 1758785994982
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application federated credential.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the federated credential provider.
	//
	// example:
	//
	// fcp_das1asda1xxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the credential was last used.
	//
	// example:
	//
	// 1758785994982
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The status of the application federated credential.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the credential was last updated.
	//
	// example:
	//
	// 1758785994982
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialName(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialType(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetCreateTime(v int64) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetDescription(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.Description = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetFederatedCredentialProviderId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetInstanceId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetLastUsedTime(v int64) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetStatus(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.Status = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetUpdateTime(v int64) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) Validate() error {
	return dara.Validate(s)
}
