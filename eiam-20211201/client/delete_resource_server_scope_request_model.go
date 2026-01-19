// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceServerScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteResourceServerScopeRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DeleteResourceServerScopeRequest
	GetInstanceId() *string
	SetResourceServerScopeId(v string) *DeleteResourceServerScopeRequest
	GetResourceServerScopeId() *string
}

type DeleteResourceServerScopeRequest struct {
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

func (s DeleteResourceServerScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceServerScopeRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceServerScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteResourceServerScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteResourceServerScopeRequest) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *DeleteResourceServerScopeRequest) SetApplicationId(v string) *DeleteResourceServerScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteResourceServerScopeRequest) SetInstanceId(v string) *DeleteResourceServerScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteResourceServerScopeRequest) SetResourceServerScopeId(v string) *DeleteResourceServerScopeRequest {
	s.ResourceServerScopeId = &v
	return s
}

func (s *DeleteResourceServerScopeRequest) Validate() error {
	return dara.Validate(s)
}
