// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RevokeResourceServerScopesFromOrganizationalUnitRequest
	GetApplicationId() *string
	SetInstanceId(v string) *RevokeResourceServerScopesFromOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *RevokeResourceServerScopesFromOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
	SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromOrganizationalUnitRequest
	GetResourceServerScopeIds() []*string
}

type RevokeResourceServerScopesFromOrganizationalUnitRequest struct {
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
	// 组织ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// ResourceServer权限ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ress_XXXXX","ress_XXXXX"]
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
}

func (s RevokeResourceServerScopesFromOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) SetApplicationId(v string) *RevokeResourceServerScopesFromOrganizationalUnitRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) SetInstanceId(v string) *RevokeResourceServerScopesFromOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *RevokeResourceServerScopesFromOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromOrganizationalUnitRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
