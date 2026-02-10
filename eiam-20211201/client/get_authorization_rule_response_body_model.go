// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRule(v *GetAuthorizationRuleResponseBodyAuthorizationRule) *GetAuthorizationRuleResponseBody
	GetAuthorizationRule() *GetAuthorizationRuleResponseBodyAuthorizationRule
	SetRequestId(v string) *GetAuthorizationRuleResponseBody
	GetRequestId() *string
}

type GetAuthorizationRuleResponseBody struct {
	AuthorizationRule *GetAuthorizationRuleResponseBodyAuthorizationRule `json:"AuthorizationRule,omitempty" xml:"AuthorizationRule,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationRuleResponseBody) GetAuthorizationRule() *GetAuthorizationRuleResponseBodyAuthorizationRule {
	return s.AuthorizationRule
}

func (s *GetAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorizationRuleResponseBody) SetAuthorizationRule(v *GetAuthorizationRuleResponseBodyAuthorizationRule) *GetAuthorizationRuleResponseBody {
	s.AuthorizationRule = v
	return s
}

func (s *GetAuthorizationRuleResponseBody) SetRequestId(v string) *GetAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationRuleResponseBody) Validate() error {
	if s.AuthorizationRule != nil {
		if err := s.AuthorizationRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorizationRuleResponseBodyAuthorizationRule struct {
	// 授权资源范围，枚举值：global（项目下所有资源）、custom（指定资源）。
	//
	// example:
	//
	// global
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// 授权规则的创建类型，枚举类型：user_created（用户创建)，approval_created（审批创建)。
	//
	// example:
	//
	// user_custom
	AuthorizationRuleCreationType *string `json:"AuthorizationRuleCreationType,omitempty" xml:"AuthorizationRuleCreationType,omitempty"`
	// 授权规则标识。
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 授权规则名称。
	//
	// example:
	//
	// test-name
	AuthorizationRuleName *string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	// 授权规则主体ID，主体类型对应的主体ID。
	AuthorizationRuleSubjectId *string `json:"AuthorizationRuleSubjectId,omitempty" xml:"AuthorizationRuleSubjectId,omitempty"`
	// 授权规则主体范围，枚举类型：shared（共享型，即支持所有主体，包括账户、应用），exclusive（专属类型）
	AuthorizationRuleSubjectScope *string `json:"AuthorizationRuleSubjectScope,omitempty" xml:"AuthorizationRuleSubjectScope,omitempty"`
	// 授权规则主体类型，枚举类型：application（应用)，user（账户)。
	AuthorizationRuleSubjectType *string `json:"AuthorizationRuleSubjectType,omitempty" xml:"AuthorizationRuleSubjectType,omitempty"`
	// 创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 授权规则描述，长度限制为128字符。
	//
	// example:
	//
	// this is a test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 授权规则关联的项目标识。
	//
	// example:
	//
	// iprj_system_default
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 授权规则状态，枚举值：enabled（启用）、disabled（禁用）。
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 最近一次更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAuthorizationRuleResponseBodyAuthorizationRule) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationRuleResponseBodyAuthorizationRule) GoString() string {
	return s.String()
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationResourceScope() *string {
	return s.AuthorizationResourceScope
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleCreationType() *string {
	return s.AuthorizationRuleCreationType
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleName() *string {
	return s.AuthorizationRuleName
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleSubjectId() *string {
	return s.AuthorizationRuleSubjectId
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleSubjectScope() *string {
	return s.AuthorizationRuleSubjectScope
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleSubjectType() *string {
	return s.AuthorizationRuleSubjectType
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetDescription() *string {
	return s.Description
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetStatus() *string {
	return s.Status
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationResourceScope(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationResourceScope = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleCreationType(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleCreationType = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleId(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleId = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleName(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleName = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleSubjectId(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleSubjectId = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleSubjectScope(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleSubjectScope = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleSubjectType(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleSubjectType = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetCreateTime(v int64) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.CreateTime = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetDescription(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.Description = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetInstanceId(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetProjectId(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.ProjectId = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetStatus(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.Status = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetUpdateTime(v int64) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.UpdateTime = &v
	return s
}

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) Validate() error {
	return dara.Validate(s)
}
