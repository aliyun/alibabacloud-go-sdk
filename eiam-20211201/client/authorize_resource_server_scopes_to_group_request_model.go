// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AuthorizeResourceServerScopesToGroupRequest
	GetApplicationId() *string
	SetClientToken(v string) *AuthorizeResourceServerScopesToGroupRequest
	GetClientToken() *string
	SetGroupId(v string) *AuthorizeResourceServerScopesToGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *AuthorizeResourceServerScopesToGroupRequest
	GetInstanceId() *string
	SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToGroupRequest
	GetResourceServerScopeIds() []*string
}

type AuthorizeResourceServerScopesToGroupRequest struct {
	// The application ID of the resource server.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_mkv7rgt4d7i4u7zqtzev2mxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of scope IDs under the resource server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ress_XXXXX","ress_XXXXX"]
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
}

func (s AuthorizeResourceServerScopesToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToGroupRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AuthorizeResourceServerScopesToGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeResourceServerScopesToGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AuthorizeResourceServerScopesToGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeResourceServerScopesToGroupRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *AuthorizeResourceServerScopesToGroupRequest) SetApplicationId(v string) *AuthorizeResourceServerScopesToGroupRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupRequest) SetClientToken(v string) *AuthorizeResourceServerScopesToGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupRequest) SetGroupId(v string) *AuthorizeResourceServerScopesToGroupRequest {
	s.GroupId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupRequest) SetInstanceId(v string) *AuthorizeResourceServerScopesToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupRequest) SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToGroupRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupRequest) Validate() error {
	return dara.Validate(s)
}
