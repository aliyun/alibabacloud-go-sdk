// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuto(v bool) *UpdateArtifactLifecycleRuleRequest
	GetAuto() *bool
	SetDryRun(v bool) *UpdateArtifactLifecycleRuleRequest
	GetDryRun() *bool
	SetEnableDeleteTag(v bool) *UpdateArtifactLifecycleRuleRequest
	GetEnableDeleteTag() *bool
	SetEnableDeleteUntaggedManifest(v bool) *UpdateArtifactLifecycleRuleRequest
	GetEnableDeleteUntaggedManifest() *bool
	SetInstanceId(v string) *UpdateArtifactLifecycleRuleRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *UpdateArtifactLifecycleRuleRequest
	GetNamespaceName() *string
	SetRepoName(v string) *UpdateArtifactLifecycleRuleRequest
	GetRepoName() *string
	SetRetentionTagCount(v int64) *UpdateArtifactLifecycleRuleRequest
	GetRetentionTagCount() *int64
	SetRuleId(v string) *UpdateArtifactLifecycleRuleRequest
	GetRuleId() *string
	SetScheduleTime(v string) *UpdateArtifactLifecycleRuleRequest
	GetScheduleTime() *string
	SetScope(v string) *UpdateArtifactLifecycleRuleRequest
	GetScope() *string
	SetTagRegexp(v string) *UpdateArtifactLifecycleRuleRequest
	GetTagRegexp() *string
}

type UpdateArtifactLifecycleRuleRequest struct {
	// Specifies whether to automatically execute the rule.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// Specifies whether to enable DryRun mode. If DryRun mode is enabled, only the lifecycle task scan is performed and no actual data cleanup is performed. DryRun mode is disabled by default.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable lifecycle management.
	//
	// Only one of this parameter and EnableDeleteUntaggedManifest can be set to true.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// Specifies whether to enable artifact cleanup.
	//
	// Only one of this parameter and EnableDeleteTag can be set to true.
	//
	// example:
	//
	// false
	EnableDeleteUntaggedManifest *bool `json:"EnableDeleteUntaggedManifest,omitempty" xml:"EnableDeleteUntaggedManifest,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-r6ym0lerldp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The image repository name.
	//
	// example:
	//
	// test_1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The number of images to retain.
	//
	// example:
	//
	// 30
	RetentionTagCount *int64 `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cralr-luq6qiegzvx****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The execution cycle.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The cleanup scope.
	//
	// example:
	//
	// REPO
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression used to retain image versions.
	//
	// example:
	//
	// .*production_.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s UpdateArtifactLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactLifecycleRuleRequest) GetAuto() *bool {
	return s.Auto
}

func (s *UpdateArtifactLifecycleRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateArtifactLifecycleRuleRequest) GetEnableDeleteTag() *bool {
	return s.EnableDeleteTag
}

func (s *UpdateArtifactLifecycleRuleRequest) GetEnableDeleteUntaggedManifest() *bool {
	return s.EnableDeleteUntaggedManifest
}

func (s *UpdateArtifactLifecycleRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateArtifactLifecycleRuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateArtifactLifecycleRuleRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *UpdateArtifactLifecycleRuleRequest) GetRetentionTagCount() *int64 {
	return s.RetentionTagCount
}

func (s *UpdateArtifactLifecycleRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateArtifactLifecycleRuleRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *UpdateArtifactLifecycleRuleRequest) GetScope() *string {
	return s.Scope
}

func (s *UpdateArtifactLifecycleRuleRequest) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *UpdateArtifactLifecycleRuleRequest) SetAuto(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.Auto = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetDryRun(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetEnableDeleteUntaggedManifest(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.EnableDeleteUntaggedManifest = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetInstanceId(v string) *UpdateArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetNamespaceName(v string) *UpdateArtifactLifecycleRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetRepoName(v string) *UpdateArtifactLifecycleRuleRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetRetentionTagCount(v int64) *UpdateArtifactLifecycleRuleRequest {
	s.RetentionTagCount = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetRuleId(v string) *UpdateArtifactLifecycleRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetScheduleTime(v string) *UpdateArtifactLifecycleRuleRequest {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetScope(v string) *UpdateArtifactLifecycleRuleRequest {
	s.Scope = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetTagRegexp(v string) *UpdateArtifactLifecycleRuleRequest {
	s.TagRegexp = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
