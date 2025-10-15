// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsForProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials
	SetMaxResults(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetTotalCount() *int32
}

type ListApplicationFederatedCredentialsForProviderResponseBody struct {
	ApplicationFederatedCredentials []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials `json:"ApplicationFederatedCredentials,omitempty" xml:"ApplicationFederatedCredentials,omitempty" type:"Repeated"`
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
	// 本次调用返回的查询凭证（Token）值，用于上一次翻页查询。
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	return s.ApplicationFederatedCredentials
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.ApplicationFederatedCredentials = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetMaxResults(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetNextToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetPreviousToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetRequestId(v string) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetTotalCount(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) Validate() error {
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

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials struct {
	// 应用联邦凭证ID
	//
	// example:
	//
	// afc_dads12sadxxxxx
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// 应用联邦凭证名称
	//
	// example:
	//
	// test
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// 应用联邦凭证类型
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// 应用ID
	//
	// example:
	//
	// app_asda1dsadxxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 1758785994982
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用联邦凭证描述
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用联邦凭证提供者ID
	//
	// example:
	//
	// fcp_adasd12dxxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// EAIM 实例ID
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 最近使用时间
	//
	// example:
	//
	// 1758785994982
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// 应用联邦凭证状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 1758785994982
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialName(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialType(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetCreateTime(v int64) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetDescription(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.Description = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetFederatedCredentialProviderId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetInstanceId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetLastUsedTime(v int64) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetStatus(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.Status = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetUpdateTime(v int64) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) Validate() error {
	return dara.Validate(s)
}
