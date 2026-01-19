// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceServerScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateResourceServerScopeRequest
	GetApplicationId() *string
	SetInstanceId(v string) *UpdateResourceServerScopeRequest
	GetInstanceId() *string
	SetResourceServerScopeId(v string) *UpdateResourceServerScopeRequest
	GetResourceServerScopeId() *string
	SetResourceServerScopeName(v string) *UpdateResourceServerScopeRequest
	GetResourceServerScopeName() *string
}

type UpdateResourceServerScopeRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ResourceServer下Scope唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// ress_neg35flu6byysxwutaxu3dxxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
	// 权限名称
	//
	// This parameter is required.
	//
	// example:
	//
	// Read All User
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
}

func (s UpdateResourceServerScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceServerScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceServerScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateResourceServerScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateResourceServerScopeRequest) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *UpdateResourceServerScopeRequest) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *UpdateResourceServerScopeRequest) SetApplicationId(v string) *UpdateResourceServerScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateResourceServerScopeRequest) SetInstanceId(v string) *UpdateResourceServerScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceServerScopeRequest) SetResourceServerScopeId(v string) *UpdateResourceServerScopeRequest {
	s.ResourceServerScopeId = &v
	return s
}

func (s *UpdateResourceServerScopeRequest) SetResourceServerScopeName(v string) *UpdateResourceServerScopeRequest {
	s.ResourceServerScopeName = &v
	return s
}

func (s *UpdateResourceServerScopeRequest) Validate() error {
	return dara.Validate(s)
}
