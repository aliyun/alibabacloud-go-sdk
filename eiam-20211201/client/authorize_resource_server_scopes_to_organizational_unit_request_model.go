// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest
	GetApplicationId() *string
	SetClientToken(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest
	GetClientToken() *string
	SetInstanceId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
	SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest
	GetResourceServerScopeIds() []*string
}

type AuthorizeResourceServerScopesToOrganizationalUnitRequest struct {
	// The ID of the resource server application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// A client token to ensure the idempotence of the request. Generate a unique value from your client. This token can contain only ASCII characters and must be no more than 64 characters long. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// A list of scope permission IDs for the resource server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ress_XXXXX","ress_XXXXX"]
	ResourceServerScopeIds []*string `json:"ResourceServerScopeIds,omitempty" xml:"ResourceServerScopeIds,omitempty" type:"Repeated"`
}

func (s AuthorizeResourceServerScopesToOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) GetResourceServerScopeIds() []*string {
	return s.ResourceServerScopeIds
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) SetApplicationId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) SetClientToken(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) SetInstanceId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) SetResourceServerScopeIds(v []*string) *AuthorizeResourceServerScopesToOrganizationalUnitRequest {
	s.ResourceServerScopeIds = v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
