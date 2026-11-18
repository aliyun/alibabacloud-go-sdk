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
	// The authorization rule object.
	AuthorizationRule *GetAuthorizationRuleResponseBodyAuthorizationRule `json:"AuthorizationRule,omitempty" xml:"AuthorizationRule,omitempty" type:"Struct"`
	// The request ID.
	//
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
	// The authorization resource scope. Valid values:
	//
	// - custom: Resources within the specified project scope.
	//
	// example:
	//
	// custom
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// The creation type of the authorization rule. Valid values:
	//
	// - system_init: Created by the system.
	//
	// - user_custom: Created by the user.
	//
	// example:
	//
	// user_custom
	AuthorizationRuleCreationType *string `json:"AuthorizationRuleCreationType,omitempty" xml:"AuthorizationRuleCreationType,omitempty"`
	// The authorization rule ID.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The authorization rule name.
	//
	// example:
	//
	// test-name
	AuthorizationRuleName *string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	// The scenario label of the authorization rule. Valid values:
	//
	// - privileged_cloud_account (Privileged cloud account scenario)
	//
	// - asset (Asset scenario)
	//
	// example:
	//
	// privileged_cloud_account
	AuthorizationRuleScenarioLabel *string `json:"AuthorizationRuleScenarioLabel,omitempty" xml:"AuthorizationRuleScenarioLabel,omitempty"`
	// The ID of the subject associated with the authorization rule.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	AuthorizationRuleSubjectId *string `json:"AuthorizationRuleSubjectId,omitempty" xml:"AuthorizationRuleSubjectId,omitempty"`
	// The subject scope of the authorization rule. Valid values:
	//
	// - shared: Shared. A single rule can be used by multiple subjects that meet the rule constraints.
	//
	// - exclusive: Exclusive. A single rule can only be used by the specified subject.
	//
	// example:
	//
	// shared
	AuthorizationRuleSubjectScope *string `json:"AuthorizationRuleSubjectScope,omitempty" xml:"AuthorizationRuleSubjectScope,omitempty"`
	// The subject category to which the authorization rule applies. For an exclusive rule, this field indicates the type of the specified subject. For a shared rule with a configured subject category, this field indicates the subject category that the rule allows to be associated with. Valid values:
	//
	// - application: Application.
	//
	// - user: Account. For a shared privileged cloud account rule, this value indicates the human user category, which can be associated with accounts, user groups, and organizations.
	//
	// If no subject category is configured for the rule, this field may be empty. The actual subjects associated with a shared rule are determined by the subject association query results.
	//
	// example:
	//
	// user
	AuthorizationRuleSubjectType *string `json:"AuthorizationRuleSubjectType,omitempty" xml:"AuthorizationRuleSubjectType,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the authorization rule.
	//
	// example:
	//
	// this is a test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the project associated with the authorization rule.
	//
	// example:
	//
	// iprj_system_default
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The authorization rule status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time. The value is a UNIX timestamp in milliseconds.
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

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) GetAuthorizationRuleScenarioLabel() *string {
	return s.AuthorizationRuleScenarioLabel
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

func (s *GetAuthorizationRuleResponseBodyAuthorizationRule) SetAuthorizationRuleScenarioLabel(v string) *GetAuthorizationRuleResponseBodyAuthorizationRule {
	s.AuthorizationRuleScenarioLabel = &v
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
