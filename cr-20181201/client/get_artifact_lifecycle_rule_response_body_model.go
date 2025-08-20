// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuto(v bool) *GetArtifactLifecycleRuleResponseBody
	GetAuto() *bool
	SetCode(v string) *GetArtifactLifecycleRuleResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetArtifactLifecycleRuleResponseBody
	GetCreateTime() *int64
	SetEnableDeleteTag(v bool) *GetArtifactLifecycleRuleResponseBody
	GetEnableDeleteTag() *bool
	SetInstanceId(v string) *GetArtifactLifecycleRuleResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetArtifactLifecycleRuleResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetArtifactLifecycleRuleResponseBody
	GetModifiedTime() *int64
	SetNamespaceName(v string) *GetArtifactLifecycleRuleResponseBody
	GetNamespaceName() *string
	SetNextTime(v int64) *GetArtifactLifecycleRuleResponseBody
	GetNextTime() *int64
	SetPolicies(v []*GetArtifactLifecycleRuleResponseBodyPolicies) *GetArtifactLifecycleRuleResponseBody
	GetPolicies() []*GetArtifactLifecycleRuleResponseBodyPolicies
	SetRepoName(v string) *GetArtifactLifecycleRuleResponseBody
	GetRepoName() *string
	SetRequestId(v string) *GetArtifactLifecycleRuleResponseBody
	GetRequestId() *string
	SetRetentionTagCount(v int64) *GetArtifactLifecycleRuleResponseBody
	GetRetentionTagCount() *int64
	SetRuleId(v string) *GetArtifactLifecycleRuleResponseBody
	GetRuleId() *string
	SetScheduleTime(v string) *GetArtifactLifecycleRuleResponseBody
	GetScheduleTime() *string
	SetScope(v string) *GetArtifactLifecycleRuleResponseBody
	GetScope() *string
	SetTagRegexp(v string) *GetArtifactLifecycleRuleResponseBody
	GetTagRegexp() *string
}

type GetArtifactLifecycleRuleResponseBody struct {
	// Indicates whether the lifecycle management rule is automatically executed.
	//
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the lifecycle management rule was created.
	//
	// example:
	//
	// 1571926439000
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
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the lifecycle management rule was last modified.
	//
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The time when the lifecycle management rule is next executed.
	//
	// example:
	//
	// 1701878400000
	NextTime *int64                                          `json:"NextTime,omitempty" xml:"NextTime,omitempty"`
	Policies []*GetArtifactLifecycleRuleResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The name of the image repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 724402D0-75CD-4794-BC20-7D37208****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// cralr-a18bkiajy8****
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

func (s GetArtifactLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponseBody) GetAuto() *bool {
	return s.Auto
}

func (s *GetArtifactLifecycleRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArtifactLifecycleRuleResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetArtifactLifecycleRuleResponseBody) GetEnableDeleteTag() *bool {
	return s.EnableDeleteTag
}

func (s *GetArtifactLifecycleRuleResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactLifecycleRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetArtifactLifecycleRuleResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetArtifactLifecycleRuleResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetArtifactLifecycleRuleResponseBody) GetNextTime() *int64 {
	return s.NextTime
}

func (s *GetArtifactLifecycleRuleResponseBody) GetPolicies() []*GetArtifactLifecycleRuleResponseBodyPolicies {
	return s.Policies
}

func (s *GetArtifactLifecycleRuleResponseBody) GetRepoName() *string {
	return s.RepoName
}

func (s *GetArtifactLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactLifecycleRuleResponseBody) GetRetentionTagCount() *int64 {
	return s.RetentionTagCount
}

func (s *GetArtifactLifecycleRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *GetArtifactLifecycleRuleResponseBody) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetArtifactLifecycleRuleResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetArtifactLifecycleRuleResponseBody) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *GetArtifactLifecycleRuleResponseBody) SetAuto(v bool) *GetArtifactLifecycleRuleResponseBody {
	s.Auto = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetCode(v string) *GetArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetCreateTime(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetEnableDeleteTag(v bool) *GetArtifactLifecycleRuleResponseBody {
	s.EnableDeleteTag = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetInstanceId(v string) *GetArtifactLifecycleRuleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *GetArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetModifiedTime(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetNamespaceName(v string) *GetArtifactLifecycleRuleResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetNextTime(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.NextTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetPolicies(v []*GetArtifactLifecycleRuleResponseBodyPolicies) *GetArtifactLifecycleRuleResponseBody {
	s.Policies = v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRepoName(v string) *GetArtifactLifecycleRuleResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRequestId(v string) *GetArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRetentionTagCount(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.RetentionTagCount = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRuleId(v string) *GetArtifactLifecycleRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetScheduleTime(v string) *GetArtifactLifecycleRuleResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetScope(v string) *GetArtifactLifecycleRuleResponseBody {
	s.Scope = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetTagRegexp(v string) *GetArtifactLifecycleRuleResponseBody {
	s.TagRegexp = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetArtifactLifecycleRuleResponseBodyPolicies struct {
	Condition *GetArtifactLifecycleRuleResponseBodyPoliciesCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	Filter    *GetArtifactLifecycleRuleResponseBodyPoliciesFilter    `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	Type      *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetArtifactLifecycleRuleResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) GetCondition() *GetArtifactLifecycleRuleResponseBodyPoliciesCondition {
	return s.Condition
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) GetFilter() *GetArtifactLifecycleRuleResponseBodyPoliciesFilter {
	return s.Filter
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) GetType() *string {
	return s.Type
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) SetCondition(v *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) *GetArtifactLifecycleRuleResponseBodyPolicies {
	s.Condition = v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) SetFilter(v *GetArtifactLifecycleRuleResponseBodyPoliciesFilter) *GetArtifactLifecycleRuleResponseBodyPolicies {
	s.Filter = v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) SetType(v string) *GetArtifactLifecycleRuleResponseBodyPolicies {
	s.Type = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}

type GetArtifactLifecycleRuleResponseBodyPoliciesCondition struct {
	LastPullOlderThanDays *int32 `json:"LastPullOlderThanDays,omitempty" xml:"LastPullOlderThanDays,omitempty"`
	LastPushOlderThanDays *int32 `json:"LastPushOlderThanDays,omitempty" xml:"LastPushOlderThanDays,omitempty"`
	LatestTagCount        *int32 `json:"LatestTagCount,omitempty" xml:"LatestTagCount,omitempty"`
}

func (s GetArtifactLifecycleRuleResponseBodyPoliciesCondition) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponseBodyPoliciesCondition) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) GetLastPullOlderThanDays() *int32 {
	return s.LastPullOlderThanDays
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) GetLastPushOlderThanDays() *int32 {
	return s.LastPushOlderThanDays
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) GetLatestTagCount() *int32 {
	return s.LatestTagCount
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) SetLastPullOlderThanDays(v int32) *GetArtifactLifecycleRuleResponseBodyPoliciesCondition {
	s.LastPullOlderThanDays = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) SetLastPushOlderThanDays(v int32) *GetArtifactLifecycleRuleResponseBodyPoliciesCondition {
	s.LastPushOlderThanDays = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) SetLatestTagCount(v int32) *GetArtifactLifecycleRuleResponseBodyPoliciesCondition {
	s.LatestTagCount = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesCondition) Validate() error {
	return dara.Validate(s)
}

type GetArtifactLifecycleRuleResponseBodyPoliciesFilter struct {
	TagWildcard *string `json:"TagWildcard,omitempty" xml:"TagWildcard,omitempty"`
}

func (s GetArtifactLifecycleRuleResponseBodyPoliciesFilter) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponseBodyPoliciesFilter) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesFilter) GetTagWildcard() *string {
	return s.TagWildcard
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesFilter) SetTagWildcard(v string) *GetArtifactLifecycleRuleResponseBodyPoliciesFilter {
	s.TagWildcard = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBodyPoliciesFilter) Validate() error {
	return dara.Validate(s)
}
