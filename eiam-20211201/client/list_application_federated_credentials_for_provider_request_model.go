// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsForProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderId(v string) *ListApplicationFederatedCredentialsForProviderRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *ListApplicationFederatedCredentialsForProviderRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListApplicationFederatedCredentialsForProviderRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationFederatedCredentialsForProviderRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationFederatedCredentialsForProviderRequest
	GetPreviousToken() *string
}

type ListApplicationFederatedCredentialsForProviderRequest struct {
	// 联邦凭证提供方ID
	//
	// example:
	//
	// fcp_adasd12dxxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询上一页凭证（Token），取值为上一次API调用返回的previousToken参数值。
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) SetFederatedCredentialProviderId(v string) *ListApplicationFederatedCredentialsForProviderRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) SetInstanceId(v string) *ListApplicationFederatedCredentialsForProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) SetMaxResults(v int32) *ListApplicationFederatedCredentialsForProviderRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) SetNextToken(v string) *ListApplicationFederatedCredentialsForProviderRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) SetPreviousToken(v string) *ListApplicationFederatedCredentialsForProviderRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderRequest) Validate() error {
	return dara.Validate(s)
}
