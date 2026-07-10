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
	// The application ID of the ResourceServer.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
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
	// The organizational unit ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The list of scope permission IDs under the ResourceServer.
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
