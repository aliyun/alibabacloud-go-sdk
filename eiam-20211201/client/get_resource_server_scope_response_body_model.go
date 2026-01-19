// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceServerScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceServerScopeResponseBody
	GetRequestId() *string
	SetResourceServerScope(v *GetResourceServerScopeResponseBodyResourceServerScope) *GetResourceServerScopeResponseBody
	GetResourceServerScope() *GetResourceServerScopeResponseBodyResourceServerScope
}

type GetResourceServerScopeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceServerScope *GetResourceServerScopeResponseBodyResourceServerScope `json:"ResourceServerScope,omitempty" xml:"ResourceServerScope,omitempty" type:"Struct"`
}

func (s GetResourceServerScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceServerScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceServerScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceServerScopeResponseBody) GetResourceServerScope() *GetResourceServerScopeResponseBodyResourceServerScope {
	return s.ResourceServerScope
}

func (s *GetResourceServerScopeResponseBody) SetRequestId(v string) *GetResourceServerScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceServerScopeResponseBody) SetResourceServerScope(v *GetResourceServerScopeResponseBodyResourceServerScope) *GetResourceServerScopeResponseBody {
	s.ResourceServerScope = v
	return s
}

func (s *GetResourceServerScopeResponseBody) Validate() error {
	if s.ResourceServerScope != nil {
		if err := s.ResourceServerScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceServerScopeResponseBodyResourceServerScope struct {
	// IDaaS EIAM 应用Id
	//
	// example:
	//
	// app_xxxxxxxxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_xxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IDaaS EIAM ResourceServer下权限Id
	//
	// example:
	//
	// rss_xxxxxxxxxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
	// IDaaS EIAM ResourceServer下权限名称
	//
	// example:
	//
	// 读取全部用户
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
	// IDaaS EIAM ResourceServer下权限类型
	//
	// example:
	//
	// urn:alibaba:idaas:resourceserver:scope:delegated
	ResourceServerScopeType *string `json:"ResourceServerScopeType,omitempty" xml:"ResourceServerScopeType,omitempty"`
	// IDaaS EIAM ResourceServer下权限值
	//
	// example:
	//
	// User:Read:ALL
	ResourceServerScopeValue *string `json:"ResourceServerScopeValue,omitempty" xml:"ResourceServerScopeValue,omitempty"`
}

func (s GetResourceServerScopeResponseBodyResourceServerScope) String() string {
	return dara.Prettify(s)
}

func (s GetResourceServerScopeResponseBodyResourceServerScope) GoString() string {
	return s.String()
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetResourceServerScopeType() *string {
	return s.ResourceServerScopeType
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) GetResourceServerScopeValue() *string {
	return s.ResourceServerScopeValue
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetApplicationId(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.ApplicationId = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetAuthorizationType(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.AuthorizationType = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetInstanceId(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.InstanceId = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetResourceServerScopeId(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.ResourceServerScopeId = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetResourceServerScopeName(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.ResourceServerScopeName = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetResourceServerScopeType(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.ResourceServerScopeType = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) SetResourceServerScopeValue(v string) *GetResourceServerScopeResponseBodyResourceServerScope {
	s.ResourceServerScopeValue = &v
	return s
}

func (s *GetResourceServerScopeResponseBodyResourceServerScope) Validate() error {
	return dara.Validate(s)
}
