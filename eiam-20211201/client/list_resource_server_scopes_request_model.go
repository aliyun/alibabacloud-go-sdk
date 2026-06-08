// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServerScopesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListResourceServerScopesRequest
	GetApplicationId() *string
	SetAuthorizationType(v string) *ListResourceServerScopesRequest
	GetAuthorizationType() *string
	SetInstanceId(v string) *ListResourceServerScopesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListResourceServerScopesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceServerScopesRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListResourceServerScopesRequest
	GetPreviousToken() *string
	SetResourceServerScopeIds(v []*string) *ListResourceServerScopesRequest
	GetResourceServerScopeIds() []*string
	SetResourceServerScopeName(v string) *ListResourceServerScopesRequest
	GetResourceServerScopeName() *string
	SetResourceServerScopeType(v string) *ListResourceServerScopesRequest
	GetResourceServerScopeType() *string
	SetResourceServerScopeValue(v string) *ListResourceServerScopesRequest
	GetResourceServerScopeValue() *string
}

type ListResourceServerScopesRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
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
	// example:
	//
	// PTxxxxxexample
	PreviousToken          *string   `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
	// 权限名称
	//
	// example:
	//
	// 读取全部用户
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
	// example:
	//
	// urn:alibaba:idaas:resourceserver:scope:delegated
	ResourceServerScopeType *string `json:"ResourceServerScopeType,omitempty" xml:"ResourceServerScopeType,omitempty"`
	// 权限值，大小写不敏感，格式(${ResourceType}:${ResourceOption}:${ResourceRestrict})
	//
	// example:
	//
	// User:Write:ALL
	ResourceServerScopeValue *string `json:"ResourceServerScopeValue,omitempty" xml:"ResourceServerScopeValue,omitempty"`
}

func (s ListResourceServerScopesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServerScopesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceServerScopesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListResourceServerScopesRequest) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *ListResourceServerScopesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceServerScopesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceServerScopesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceServerScopesRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListResourceServerScopesRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *ListResourceServerScopesRequest) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *ListResourceServerScopesRequest) GetResourceServerScopeType() *string {
	return s.ResourceServerScopeType
}

func (s *ListResourceServerScopesRequest) GetResourceServerScopeValue() *string {
	return s.ResourceServerScopeValue
}

func (s *ListResourceServerScopesRequest) SetApplicationId(v string) *ListResourceServerScopesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetAuthorizationType(v string) *ListResourceServerScopesRequest {
	s.AuthorizationType = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetInstanceId(v string) *ListResourceServerScopesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetMaxResults(v int32) *ListResourceServerScopesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetNextToken(v string) *ListResourceServerScopesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetPreviousToken(v string) *ListResourceServerScopesRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetResourceServerScopeIds(v []*string) *ListResourceServerScopesRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *ListResourceServerScopesRequest) SetResourceServerScopeName(v string) *ListResourceServerScopesRequest {
	s.ResourceServerScopeName = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetResourceServerScopeType(v string) *ListResourceServerScopesRequest {
	s.ResourceServerScopeType = &v
	return s
}

func (s *ListResourceServerScopesRequest) SetResourceServerScopeValue(v string) *ListResourceServerScopesRequest {
	s.ResourceServerScopeValue = &v
	return s
}

func (s *ListResourceServerScopesRequest) Validate() error {
	return dara.Validate(s)
}
