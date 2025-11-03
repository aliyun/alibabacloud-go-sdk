// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListArtifactLifecycleRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListArtifactLifecycleRuleResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListArtifactLifecycleRuleResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListArtifactLifecycleRuleResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListArtifactLifecycleRuleResponseBody
	GetRequestId() *string
	SetRules(v []*ListArtifactLifecycleRuleResponseBodyRules) *ListArtifactLifecycleRuleResponseBody
	GetRules() []*ListArtifactLifecycleRuleResponseBodyRules
	SetTotalCount(v int32) *ListArtifactLifecycleRuleResponseBody
	GetTotalCount() *int32
}

type ListArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F92D82F9-A4C4-5A4A-97B9-E495BF1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// _
	Rules []*ListArtifactLifecycleRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 39
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListArtifactLifecycleRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListArtifactLifecycleRuleResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListArtifactLifecycleRuleResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactLifecycleRuleResponseBody) GetRules() []*ListArtifactLifecycleRuleResponseBodyRules {
	return s.Rules
}

func (s *ListArtifactLifecycleRuleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListArtifactLifecycleRuleResponseBody) SetCode(v string) *ListArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *ListArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetPageNo(v int32) *ListArtifactLifecycleRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetPageSize(v int32) *ListArtifactLifecycleRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetRequestId(v string) *ListArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetRules(v []*ListArtifactLifecycleRuleResponseBodyRules) *ListArtifactLifecycleRuleResponseBody {
	s.Rules = v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetTotalCount(v int32) *ListArtifactLifecycleRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactLifecycleRuleResponseBodyRules struct {
	// Indicates whether the lifecycle management rule is automatically executed.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The time when the lifecycle management rule was created.
	//
	// example:
	//
	// 1638187989000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether lifecycle management is enabled for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-brlg4cbj2yl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the lifecycle management rule was last modified.
	//
	// example:
	//
	// 1678341923385
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The time when the lifecycle management rule is next executed.
	//
	// example:
	//
	// 1638187989000
	NextTime *int64                                                `json:"NextTime,omitempty" xml:"NextTime,omitempty"`
	Policies []*ListArtifactLifecycleRuleResponseBodyRulesPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The name of the image repository.
	//
	// example:
	//
	// test_1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The number of retained images.
	//
	// example:
	//
	// 30
	RetentionTagCount *int64 `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cralr-yqx1q5sir6d****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The execution cycle of the lifecycle management rule.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The deletion scope of artifacts.
	//
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that indicates which image tags are retained.
	//
	// example:
	//
	// .*-alpine
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetAuto() *bool {
	return s.Auto
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetEnableDeleteTag() *bool {
	return s.EnableDeleteTag
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetNextTime() *int64 {
	return s.NextTime
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetPolicies() []*ListArtifactLifecycleRuleResponseBodyRulesPolicies {
	return s.Policies
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetRepoName() *string {
	return s.RepoName
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetRetentionTagCount() *int64 {
	return s.RetentionTagCount
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetScope() *string {
	return s.Scope
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetAuto(v bool) *ListArtifactLifecycleRuleResponseBodyRules {
	s.Auto = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetCreateTime(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetEnableDeleteTag(v bool) *ListArtifactLifecycleRuleResponseBodyRules {
	s.EnableDeleteTag = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetInstanceId(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetModifiedTime(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.ModifiedTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetNamespaceName(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.NamespaceName = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetNextTime(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.NextTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetPolicies(v []*ListArtifactLifecycleRuleResponseBodyRulesPolicies) *ListArtifactLifecycleRuleResponseBodyRules {
	s.Policies = v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetRepoName(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.RepoName = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetRetentionTagCount(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.RetentionTagCount = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetRuleId(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetScheduleTime(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.ScheduleTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetScope(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.Scope = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetTagRegexp(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.TagRegexp = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactLifecycleRuleResponseBodyRulesPolicies struct {
	Condition *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	Filter    *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter    `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	Type      *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBodyRulesPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBodyRulesPolicies) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) GetCondition() *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition {
	return s.Condition
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) GetFilter() *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter {
	return s.Filter
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) GetType() *string {
	return s.Type
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) SetCondition(v *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) *ListArtifactLifecycleRuleResponseBodyRulesPolicies {
	s.Condition = v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) SetFilter(v *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter) *ListArtifactLifecycleRuleResponseBodyRulesPolicies {
	s.Filter = v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) SetType(v string) *ListArtifactLifecycleRuleResponseBodyRulesPolicies {
	s.Type = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPolicies) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition struct {
	LastPullOlderThanDays *int32 `json:"LastPullOlderThanDays,omitempty" xml:"LastPullOlderThanDays,omitempty"`
	LastPushOlderThanDays *int32 `json:"LastPushOlderThanDays,omitempty" xml:"LastPushOlderThanDays,omitempty"`
	LatestTagCount        *int32 `json:"LatestTagCount,omitempty" xml:"LatestTagCount,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) GetLastPullOlderThanDays() *int32 {
	return s.LastPullOlderThanDays
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) GetLastPushOlderThanDays() *int32 {
	return s.LastPushOlderThanDays
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) GetLatestTagCount() *int32 {
	return s.LatestTagCount
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) SetLastPullOlderThanDays(v int32) *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition {
	s.LastPullOlderThanDays = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) SetLastPushOlderThanDays(v int32) *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition {
	s.LastPushOlderThanDays = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) SetLatestTagCount(v int32) *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition {
	s.LatestTagCount = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesCondition) Validate() error {
	return dara.Validate(s)
}

type ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter struct {
	TagWildcard *string `json:"TagWildcard,omitempty" xml:"TagWildcard,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter) GetTagWildcard() *string {
	return s.TagWildcard
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter) SetTagWildcard(v string) *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter {
	s.TagWildcard = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRulesPoliciesFilter) Validate() error {
	return dara.Validate(s)
}
