// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFederatedCredentialProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderName(v string) *ListFederatedCredentialProvidersRequest
	GetFederatedCredentialProviderName() *string
	SetFederatedCredentialProviderType(v string) *ListFederatedCredentialProvidersRequest
	GetFederatedCredentialProviderType() *string
	SetInstanceId(v string) *ListFederatedCredentialProvidersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListFederatedCredentialProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFederatedCredentialProvidersRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListFederatedCredentialProvidersRequest
	GetPreviousToken() *string
}

type ListFederatedCredentialProvidersRequest struct {
	// 联邦凭证提供方名称
	//
	// example:
	//
	// test
	FederatedCredentialProviderName *string `json:"FederatedCredentialProviderName,omitempty" xml:"FederatedCredentialProviderName,omitempty"`
	// 联邦凭证提供方类型
	//
	// example:
	//
	// pkcs7
	FederatedCredentialProviderType *string `json:"FederatedCredentialProviderType,omitempty" xml:"FederatedCredentialProviderType,omitempty"`
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

func (s ListFederatedCredentialProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersRequest) GetFederatedCredentialProviderName() *string {
	return s.FederatedCredentialProviderName
}

func (s *ListFederatedCredentialProvidersRequest) GetFederatedCredentialProviderType() *string {
	return s.FederatedCredentialProviderType
}

func (s *ListFederatedCredentialProvidersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFederatedCredentialProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFederatedCredentialProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFederatedCredentialProvidersRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListFederatedCredentialProvidersRequest) SetFederatedCredentialProviderName(v string) *ListFederatedCredentialProvidersRequest {
	s.FederatedCredentialProviderName = &v
	return s
}

func (s *ListFederatedCredentialProvidersRequest) SetFederatedCredentialProviderType(v string) *ListFederatedCredentialProvidersRequest {
	s.FederatedCredentialProviderType = &v
	return s
}

func (s *ListFederatedCredentialProvidersRequest) SetInstanceId(v string) *ListFederatedCredentialProvidersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFederatedCredentialProvidersRequest) SetMaxResults(v int32) *ListFederatedCredentialProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFederatedCredentialProvidersRequest) SetNextToken(v string) *ListFederatedCredentialProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListFederatedCredentialProvidersRequest) SetPreviousToken(v string) *ListFederatedCredentialProvidersRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListFederatedCredentialProvidersRequest) Validate() error {
	return dara.Validate(s)
}
