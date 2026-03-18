// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderIds(v []*string) *ListCredentialProvidersRequest
	GetCredentialProviderIds() []*string
	SetCredentialProviderTypes(v []*string) *ListCredentialProvidersRequest
	GetCredentialProviderTypes() []*string
	SetFilter(v []*ListCredentialProvidersRequestFilter) *ListCredentialProvidersRequest
	GetFilter() []*ListCredentialProvidersRequestFilter
	SetInstanceId(v string) *ListCredentialProvidersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListCredentialProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCredentialProvidersRequest
	GetNextToken() *string
	SetStatuses(v []*string) *ListCredentialProvidersRequest
	GetStatuses() []*string
}

type ListCredentialProvidersRequest struct {
	// 认证令牌提供商ID列表。
	CredentialProviderIds []*string `json:"CredentialProviderIds,omitempty" xml:"CredentialProviderIds,omitempty" type:"Repeated"`
	// 认证令牌提供商类型列表。
	CredentialProviderTypes []*string                               `json:"CredentialProviderTypes,omitempty" xml:"CredentialProviderTypes,omitempty" type:"Repeated"`
	Filter                  []*ListCredentialProvidersRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	// 认证令牌提供商状态列表。
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListCredentialProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersRequest) GetCredentialProviderIds() []*string {
	return s.CredentialProviderIds
}

func (s *ListCredentialProvidersRequest) GetCredentialProviderTypes() []*string {
	return s.CredentialProviderTypes
}

func (s *ListCredentialProvidersRequest) GetFilter() []*ListCredentialProvidersRequestFilter {
	return s.Filter
}

func (s *ListCredentialProvidersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCredentialProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialProvidersRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListCredentialProvidersRequest) SetCredentialProviderIds(v []*string) *ListCredentialProvidersRequest {
	s.CredentialProviderIds = v
	return s
}

func (s *ListCredentialProvidersRequest) SetCredentialProviderTypes(v []*string) *ListCredentialProvidersRequest {
	s.CredentialProviderTypes = v
	return s
}

func (s *ListCredentialProvidersRequest) SetFilter(v []*ListCredentialProvidersRequestFilter) *ListCredentialProvidersRequest {
	s.Filter = v
	return s
}

func (s *ListCredentialProvidersRequest) SetInstanceId(v string) *ListCredentialProvidersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCredentialProvidersRequest) SetMaxResults(v int32) *ListCredentialProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialProvidersRequest) SetNextToken(v string) *ListCredentialProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListCredentialProvidersRequest) SetStatuses(v []*string) *ListCredentialProvidersRequest {
	s.Statuses = v
	return s
}

func (s *ListCredentialProvidersRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCredentialProvidersRequestFilter struct {
	// example:
	//
	// CredentialProviderName
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListCredentialProvidersRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersRequestFilter) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListCredentialProvidersRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListCredentialProvidersRequestFilter) SetName(v string) *ListCredentialProvidersRequestFilter {
	s.Name = &v
	return s
}

func (s *ListCredentialProvidersRequestFilter) SetValue(v []*string) *ListCredentialProvidersRequestFilter {
	s.Value = v
	return s
}

func (s *ListCredentialProvidersRequestFilter) Validate() error {
	return dara.Validate(s)
}
