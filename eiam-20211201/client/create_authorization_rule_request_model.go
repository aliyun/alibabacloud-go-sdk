// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceScope(v string) *CreateAuthorizationRuleRequest
	GetAuthorizationResourceScope() *string
	SetAuthorizationRuleName(v string) *CreateAuthorizationRuleRequest
	GetAuthorizationRuleName() *string
	SetClientToken(v string) *CreateAuthorizationRuleRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAuthorizationRuleRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateAuthorizationRuleRequest
	GetInstanceId() *string
	SetProjectId(v string) *CreateAuthorizationRuleRequest
	GetProjectId() *string
}

type CreateAuthorizationRuleRequest struct {
	// The scope of authorized resources. Valid values:
	//
	// - global: all resources within the project.
	//
	// - custom: specified resources within the project.
	//
	// example:
	//
	// global
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// The name of the authorization rule. The name can be up to 64 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_rule
	AuthorizationRuleName *string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	// A unique identifier that you provide to ensure the idempotence of the request. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the authorization rule. The description can be up to 128 characters long.
	//
	// example:
	//
	// this is a test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the project to associate with the authorization rule. If you are unsure which project to use, you can associate the rule with the default project, iprj_system_default.
	//
	// This parameter is required.
	//
	// example:
	//
	// iprj_system_default
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleRequest) GetAuthorizationResourceScope() *string {
	return s.AuthorizationResourceScope
}

func (s *CreateAuthorizationRuleRequest) GetAuthorizationRuleName() *string {
	return s.AuthorizationRuleName
}

func (s *CreateAuthorizationRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAuthorizationRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAuthorizationRuleRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateAuthorizationRuleRequest) SetAuthorizationResourceScope(v string) *CreateAuthorizationRuleRequest {
	s.AuthorizationResourceScope = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetAuthorizationRuleName(v string) *CreateAuthorizationRuleRequest {
	s.AuthorizationRuleName = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetClientToken(v string) *CreateAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetDescription(v string) *CreateAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetInstanceId(v string) *CreateAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetProjectId(v string) *CreateAuthorizationRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
