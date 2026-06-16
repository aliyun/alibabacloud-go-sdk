// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceScope(v string) *UpdateAuthorizationRuleRequest
	GetAuthorizationResourceScope() *string
	SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetAuthorizationRuleName(v string) *UpdateAuthorizationRuleRequest
	GetAuthorizationRuleName() *string
	SetClientToken(v string) *UpdateAuthorizationRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateAuthorizationRuleRequest
	GetInstanceId() *string
}

type UpdateAuthorizationRuleRequest struct {
	// The scope of the authorized resources. Valid values:
	//
	// - global: global resources within the project.
	//
	// - custom: specified resources within the project.
	//
	// example:
	//
	// global
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The name of the authorization rule. The name can be up to 64 characters long.
	//
	// example:
	//
	// test_rule
	AuthorizationRuleName *string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	// A client token that ensures the idempotence of the request. Generate a unique value for this parameter from your client. The token can contain only ASCII characters and must be no more than 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
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
}

func (s UpdateAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleRequest) GetAuthorizationResourceScope() *string {
	return s.AuthorizationResourceScope
}

func (s *UpdateAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *UpdateAuthorizationRuleRequest) GetAuthorizationRuleName() *string {
	return s.AuthorizationRuleName
}

func (s *UpdateAuthorizationRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationRuleRequest) SetAuthorizationResourceScope(v string) *UpdateAuthorizationRuleRequest {
	s.AuthorizationResourceScope = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetAuthorizationRuleName(v string) *UpdateAuthorizationRuleRequest {
	s.AuthorizationRuleName = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetClientToken(v string) *UpdateAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetInstanceId(v string) *UpdateAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
