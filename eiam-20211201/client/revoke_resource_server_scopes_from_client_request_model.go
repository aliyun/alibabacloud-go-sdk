// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientApplicationId(v string) *RevokeResourceServerScopesFromClientRequest
	GetClientApplicationId() *string
	SetInstanceId(v string) *RevokeResourceServerScopesFromClientRequest
	GetInstanceId() *string
	SetResourceServerApplicationId(v string) *RevokeResourceServerScopesFromClientRequest
	GetResourceServerApplicationId() *string
	SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromClientRequest
	GetResourceServerScopeIds() []*string
}

type RevokeResourceServerScopesFromClientRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientApplicationId *string `json:"ClientApplicationId,omitempty" xml:"ClientApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ResourceServerApplicationId *string `json:"ResourceServerApplicationId,omitempty" xml:"ResourceServerApplicationId,omitempty"`
	// ResourceServer权限ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ress_XXXXX","ress_XXXXX"]
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
}

func (s RevokeResourceServerScopesFromClientRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromClientRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromClientRequest) GetClientApplicationId() *string {
	return s.ClientApplicationId
}

func (s *RevokeResourceServerScopesFromClientRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeResourceServerScopesFromClientRequest) GetResourceServerApplicationId() *string {
	return s.ResourceServerApplicationId
}

func (s *RevokeResourceServerScopesFromClientRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *RevokeResourceServerScopesFromClientRequest) SetClientApplicationId(v string) *RevokeResourceServerScopesFromClientRequest {
	s.ClientApplicationId = &v
	return s
}

func (s *RevokeResourceServerScopesFromClientRequest) SetInstanceId(v string) *RevokeResourceServerScopesFromClientRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeResourceServerScopesFromClientRequest) SetResourceServerApplicationId(v string) *RevokeResourceServerScopesFromClientRequest {
	s.ResourceServerApplicationId = &v
	return s
}

func (s *RevokeResourceServerScopesFromClientRequest) SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromClientRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *RevokeResourceServerScopesFromClientRequest) Validate() error {
	return dara.Validate(s)
}
