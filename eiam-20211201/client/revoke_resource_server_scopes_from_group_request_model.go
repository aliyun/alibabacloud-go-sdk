// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RevokeResourceServerScopesFromGroupRequest
	GetApplicationId() *string
	SetGroupId(v string) *RevokeResourceServerScopesFromGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *RevokeResourceServerScopesFromGroupRequest
	GetInstanceId() *string
	SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromGroupRequest
	GetResourceServerScopeIds() []*string
}

type RevokeResourceServerScopesFromGroupRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// group_mkv7rgt4d7i4u7zqtzev2mxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ResourceServer权限ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ress_XXXXX","ress_XXXXX"]
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
}

func (s RevokeResourceServerScopesFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromGroupRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RevokeResourceServerScopesFromGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RevokeResourceServerScopesFromGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeResourceServerScopesFromGroupRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *RevokeResourceServerScopesFromGroupRequest) SetApplicationId(v string) *RevokeResourceServerScopesFromGroupRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeResourceServerScopesFromGroupRequest) SetGroupId(v string) *RevokeResourceServerScopesFromGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RevokeResourceServerScopesFromGroupRequest) SetInstanceId(v string) *RevokeResourceServerScopesFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeResourceServerScopesFromGroupRequest) SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromGroupRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *RevokeResourceServerScopesFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
