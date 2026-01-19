// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientApplicationId(v string) *AuthorizeResourceServerScopesToClientRequest
	GetClientApplicationId() *string
	SetInstanceId(v string) *AuthorizeResourceServerScopesToClientRequest
	GetInstanceId() *string
	SetResourceServerApplicationId(v string) *AuthorizeResourceServerScopesToClientRequest
	GetResourceServerApplicationId() *string
	SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToClientRequest
	GetResourceServerScopeIds() []*string
}

type AuthorizeResourceServerScopesToClientRequest struct {
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

func (s AuthorizeResourceServerScopesToClientRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToClientRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToClientRequest) GetClientApplicationId() *string {
	return s.ClientApplicationId
}

func (s *AuthorizeResourceServerScopesToClientRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeResourceServerScopesToClientRequest) GetResourceServerApplicationId() *string {
	return s.ResourceServerApplicationId
}

func (s *AuthorizeResourceServerScopesToClientRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *AuthorizeResourceServerScopesToClientRequest) SetClientApplicationId(v string) *AuthorizeResourceServerScopesToClientRequest {
	s.ClientApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToClientRequest) SetInstanceId(v string) *AuthorizeResourceServerScopesToClientRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToClientRequest) SetResourceServerApplicationId(v string) *AuthorizeResourceServerScopesToClientRequest {
	s.ResourceServerApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToClientRequest) SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToClientRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *AuthorizeResourceServerScopesToClientRequest) Validate() error {
	return dara.Validate(s)
}
