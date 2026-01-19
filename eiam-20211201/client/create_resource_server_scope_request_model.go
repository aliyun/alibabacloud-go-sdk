// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceServerScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateResourceServerScopeRequest
	GetApplicationId() *string
	SetAuthorizationType(v string) *CreateResourceServerScopeRequest
	GetAuthorizationType() *string
	SetInstanceId(v string) *CreateResourceServerScopeRequest
	GetInstanceId() *string
	SetResourceServerScopeName(v string) *CreateResourceServerScopeRequest
	GetResourceServerScopeName() *string
	SetResourceServerScopeType(v string) *CreateResourceServerScopeRequest
	GetResourceServerScopeType() *string
	SetResourceServerScopeValue(v string) *CreateResourceServerScopeRequest
	GetResourceServerScopeValue() *string
}

type CreateResourceServerScopeRequest struct {
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
	// 权限名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 读取全部用户
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
	// 权限类型
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:resourceserver:scope:delegated
	ResourceServerScopeType *string `json:"ResourceServerScopeType,omitempty" xml:"ResourceServerScopeType,omitempty"`
	// 权限值，大小写不敏感，格式(${ResourceType}:${ResourceOption}:${ResourceRestrict})
	//
	// This parameter is required.
	//
	// example:
	//
	// User:Read:ALL
	ResourceServerScopeValue *string `json:"ResourceServerScopeValue,omitempty" xml:"ResourceServerScopeValue,omitempty"`
}

func (s CreateResourceServerScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceServerScopeRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceServerScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateResourceServerScopeRequest) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *CreateResourceServerScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateResourceServerScopeRequest) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *CreateResourceServerScopeRequest) GetResourceServerScopeType() *string {
	return s.ResourceServerScopeType
}

func (s *CreateResourceServerScopeRequest) GetResourceServerScopeValue() *string {
	return s.ResourceServerScopeValue
}

func (s *CreateResourceServerScopeRequest) SetApplicationId(v string) *CreateResourceServerScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateResourceServerScopeRequest) SetAuthorizationType(v string) *CreateResourceServerScopeRequest {
	s.AuthorizationType = &v
	return s
}

func (s *CreateResourceServerScopeRequest) SetInstanceId(v string) *CreateResourceServerScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceServerScopeRequest) SetResourceServerScopeName(v string) *CreateResourceServerScopeRequest {
	s.ResourceServerScopeName = &v
	return s
}

func (s *CreateResourceServerScopeRequest) SetResourceServerScopeType(v string) *CreateResourceServerScopeRequest {
	s.ResourceServerScopeType = &v
	return s
}

func (s *CreateResourceServerScopeRequest) SetResourceServerScopeValue(v string) *CreateResourceServerScopeRequest {
	s.ResourceServerScopeValue = &v
	return s
}

func (s *CreateResourceServerScopeRequest) Validate() error {
	return dara.Validate(s)
}
