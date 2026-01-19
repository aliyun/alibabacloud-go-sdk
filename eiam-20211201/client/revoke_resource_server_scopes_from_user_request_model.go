// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RevokeResourceServerScopesFromUserRequest
	GetApplicationId() *string
	SetInstanceId(v string) *RevokeResourceServerScopesFromUserRequest
	GetInstanceId() *string
	SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromUserRequest
	GetResourceServerScopeIds() []*string
	SetUserId(v string) *RevokeResourceServerScopesFromUserRequest
	GetUserId() *string
}

type RevokeResourceServerScopesFromUserRequest struct {
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
	// ResourceServer权限ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ress_XXXXX","ress_XXXXX"]
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
	// 用户ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// user_wovwffm62xifdziem7an7xxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeResourceServerScopesFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromUserRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromUserRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RevokeResourceServerScopesFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeResourceServerScopesFromUserRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *RevokeResourceServerScopesFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *RevokeResourceServerScopesFromUserRequest) SetApplicationId(v string) *RevokeResourceServerScopesFromUserRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeResourceServerScopesFromUserRequest) SetInstanceId(v string) *RevokeResourceServerScopesFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeResourceServerScopesFromUserRequest) SetResourceServerScopeIds(v []*string) *RevokeResourceServerScopesFromUserRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *RevokeResourceServerScopesFromUserRequest) SetUserId(v string) *RevokeResourceServerScopesFromUserRequest {
	s.UserId = &v
	return s
}

func (s *RevokeResourceServerScopesFromUserRequest) Validate() error {
	return dara.Validate(s)
}
