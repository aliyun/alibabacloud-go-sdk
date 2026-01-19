// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AuthorizeResourceServerScopesToUserRequest
	GetApplicationId() *string
	SetClientToken(v string) *AuthorizeResourceServerScopesToUserRequest
	GetClientToken() *string
	SetInstanceId(v string) *AuthorizeResourceServerScopesToUserRequest
	GetInstanceId() *string
	SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToUserRequest
	GetResourceServerScopeIds() []*string
	SetUserId(v string) *AuthorizeResourceServerScopesToUserRequest
	GetUserId() *string
}

type AuthorizeResourceServerScopesToUserRequest struct {
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
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s AuthorizeResourceServerScopesToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToUserRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToUserRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AuthorizeResourceServerScopesToUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeResourceServerScopesToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeResourceServerScopesToUserRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *AuthorizeResourceServerScopesToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AuthorizeResourceServerScopesToUserRequest) SetApplicationId(v string) *AuthorizeResourceServerScopesToUserRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToUserRequest) SetClientToken(v string) *AuthorizeResourceServerScopesToUserRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeResourceServerScopesToUserRequest) SetInstanceId(v string) *AuthorizeResourceServerScopesToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToUserRequest) SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToUserRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *AuthorizeResourceServerScopesToUserRequest) SetUserId(v string) *AuthorizeResourceServerScopesToUserRequest {
	s.UserId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToUserRequest) Validate() error {
	return dara.Validate(s)
}
