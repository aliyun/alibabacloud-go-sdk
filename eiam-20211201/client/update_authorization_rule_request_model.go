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
	// 授权资源范围，枚举值：global（Project下的所有资源）、custom（自定义资源范围）。
	//
	// example:
	//
	// global
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 授权规则名称，长度限制最大64个字符。
	//
	// example:
	//
	// test_rule
	AuthorizationRuleName *string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
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
