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
	// The maximum number of entries per page for paging in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned in the current response for querying the next page.
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
	// The total number of entries.
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
	// The scope of authorized resources. Valid values:
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
	// - user_custom: Created by a user.
	//
	// example:
	//
	// user_custom
	AuthorizationRuleCreationType *string `json:"AuthorizationRuleCreationType,omitempty" xml:"AuthorizationRuleCreationType,omitempty"`
	// The ID of the authorization rule.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The name of the authorization rule.
	//
	// example:
	//
	// test-name
	AuthorizationRuleName *string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	// The scenario label of the authorization rule. Valid values:
	//
	// - privileged_cloud_account: Privileged cloud account scenario.
	//
	// - asset: Asset scenario.
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
	// The subject type applicable to the authorization rule. For an exclusive rule, this field indicates the type of the specified subject. For a shared rule with a specified subject type, this field indicates the subject types that the rule can be associated with. Valid values:
	//
	// - application: Application.
	//
	// - user: Account. For a shared privileged cloud account rule, this value indicates the human user type and can be associated with accounts, user groups, and organizational units.
	//
	// > If the subject type is not specified for the rule, this field may be empty. The actual subjects associated with a shared rule are subject to the query results of subject associations.
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
	// The status of the authorization rule. Valid values:
	//
	// - enabled: The rule is enabled.
	//
	// - disabled: The rule is disabled.
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
