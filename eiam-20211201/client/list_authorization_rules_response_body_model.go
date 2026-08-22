// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRules(v []*ListAuthorizationRulesResponseBodyAuthorizationRules) *ListAuthorizationRulesResponseBody
	GetAuthorizationRules() []*ListAuthorizationRulesResponseBodyAuthorizationRules
	SetMaxResults(v int32) *ListAuthorizationRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuthorizationRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAuthorizationRulesResponseBody
	GetTotalCount() *int64
}

type ListAuthorizationRulesResponseBody struct {
	// The list of authorization rules.
	AuthorizationRules []*ListAuthorizationRulesResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
	// The number of entries per page in the paging query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned for the next page query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponseBody) GetAuthorizationRules() []*ListAuthorizationRulesResponseBodyAuthorizationRules {
	return s.AuthorizationRules
}

func (s *ListAuthorizationRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizationRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizationRulesResponseBody) SetAuthorizationRules(v []*ListAuthorizationRulesResponseBodyAuthorizationRules) *ListAuthorizationRulesResponseBody {
	s.AuthorizationRules = v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetMaxResults(v int32) *ListAuthorizationRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetNextToken(v string) *ListAuthorizationRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetRequestId(v string) *ListAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetTotalCount(v int64) *ListAuthorizationRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) Validate() error {
	if s.AuthorizationRules != nil {
		for _, item := range s.AuthorizationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizationRulesResponseBodyAuthorizationRules struct {
	// The authorization resource scope. Valid values:
	//
	// - global: all resources under the project
	//
	// - custom: specified resources under the project
	//
	// example:
	//
	// global
	AuthorizationResourceScope *string `json:"AuthorizationResourceScope,omitempty" xml:"AuthorizationResourceScope,omitempty"`
	// The creation type of the authorization rule. Valid values:
	//
	// - system_init: created by the system
	//
	// - user_custom: created by the user
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
	// The scenario label of the authorization rule.
	AuthorizationRuleScenarioLabel *string `json:"AuthorizationRuleScenarioLabel,omitempty" xml:"AuthorizationRuleScenarioLabel,omitempty"`
	// The subject ID associated with the authorization rule.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	AuthorizationRuleSubjectId *string `json:"AuthorizationRuleSubjectId,omitempty" xml:"AuthorizationRuleSubjectId,omitempty"`
	// The subject scope of the authorization rule. Valid values:
	//
	// - shared: supports all subjects, including accounts and applications
	//
	// - exclusive: exclusive type
	//
	// example:
	//
	// shared
	AuthorizationRuleSubjectScope *string `json:"AuthorizationRuleSubjectScope,omitempty" xml:"AuthorizationRuleSubjectScope,omitempty"`
	// The subject type associated with the authorization rule. This parameter takes effect only when the subject scope is exclusive. Valid values:
	//
	// - application: application
	//
	// - user: account
	//
	// example:
	//
	// user
	AuthorizationRuleSubjectType *string `json:"AuthorizationRuleSubjectType,omitempty" xml:"AuthorizationRuleSubjectType,omitempty"`
	// The creation time, in UNIX timestamp format, measured in milliseconds.
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
	// The project ID associated with the authorization rule.
	//
	// example:
	//
	// iprj_system_default
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The authorization rule status. Valid values:
	//
	// - enabled: enabled
	//
	// - disabled: disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time, in UNIX timestamp format, measured in milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAuthorizationRulesResponseBodyAuthorizationRules) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesResponseBodyAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationResourceScope() *string {
	return s.AuthorizationResourceScope
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleCreationType() *string {
	return s.AuthorizationRuleCreationType
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleName() *string {
	return s.AuthorizationRuleName
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleScenarioLabel() *string {
	return s.AuthorizationRuleScenarioLabel
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleSubjectId() *string {
	return s.AuthorizationRuleSubjectId
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleSubjectScope() *string {
	return s.AuthorizationRuleSubjectScope
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetAuthorizationRuleSubjectType() *string {
	return s.AuthorizationRuleSubjectType
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetDescription() *string {
	return s.Description
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetStatus() *string {
	return s.Status
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationResourceScope(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationResourceScope = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleCreationType(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleCreationType = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleName(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleName = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleScenarioLabel(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleScenarioLabel = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleSubjectId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleSubjectId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleSubjectScope(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleSubjectScope = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleSubjectType(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleSubjectType = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetCreateTime(v int64) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.CreateTime = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDescription(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Description = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetInstanceId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetProjectId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.ProjectId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetStatus(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Status = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetUpdateTime(v int64) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.UpdateTime = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) Validate() error {
	return dara.Validate(s)
}
