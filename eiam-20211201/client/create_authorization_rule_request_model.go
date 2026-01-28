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
	// 授权资源范围，枚举值：global（Project下的所有资源）、custom（自定义资源范围）。
	//
	// example:
	//
	// global
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// 授权规则名称，长度限制最大64个字符。
	//
	// This parameter is required.
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
	// 授权规则备注描述，长度限制最大128个字符。
	//
	// example:
	//
	// this is a test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 授权规则关联的项目标识。
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
