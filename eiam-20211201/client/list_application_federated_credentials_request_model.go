// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialType(v string) *ListApplicationFederatedCredentialsRequest
	GetApplicationFederatedCredentialType() *string
	SetApplicationId(v string) *ListApplicationFederatedCredentialsRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListApplicationFederatedCredentialsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListApplicationFederatedCredentialsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationFederatedCredentialsRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationFederatedCredentialsRequest
	GetPreviousToken() *string
}

type ListApplicationFederatedCredentialsRequest struct {
	// 应用联邦凭证提供者类型
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// 应用ID
	//
	// This parameter is required.
	//
	// example:
	//
	// app_xxxasda1
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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

func (s ListApplicationFederatedCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsRequest) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *ListApplicationFederatedCredentialsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationFederatedCredentialsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationFederatedCredentialsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationFederatedCredentialsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationFederatedCredentialsRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationFederatedCredentialsRequest) SetApplicationFederatedCredentialType(v string) *ListApplicationFederatedCredentialsRequest {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *ListApplicationFederatedCredentialsRequest) SetApplicationId(v string) *ListApplicationFederatedCredentialsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsRequest) SetInstanceId(v string) *ListApplicationFederatedCredentialsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsRequest) SetMaxResults(v int32) *ListApplicationFederatedCredentialsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationFederatedCredentialsRequest) SetNextToken(v string) *ListApplicationFederatedCredentialsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsRequest) SetPreviousToken(v string) *ListApplicationFederatedCredentialsRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
