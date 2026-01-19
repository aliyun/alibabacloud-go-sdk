// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceServerScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetResourceServerScopeRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetResourceServerScopeRequest
	GetInstanceId() *string
	SetResourceServerScopeId(v string) *GetResourceServerScopeRequest
	GetResourceServerScopeId() *string
}

type GetResourceServerScopeRequest struct {
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
}

func (s GetResourceServerScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceServerScopeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceServerScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetResourceServerScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetResourceServerScopeRequest) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *GetResourceServerScopeRequest) SetApplicationId(v string) *GetResourceServerScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetResourceServerScopeRequest) SetInstanceId(v string) *GetResourceServerScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetResourceServerScopeRequest) SetResourceServerScopeId(v string) *GetResourceServerScopeRequest {
	s.ResourceServerScopeId = &v
	return s
}

func (s *GetResourceServerScopeRequest) Validate() error {
	return dara.Validate(s)
}
