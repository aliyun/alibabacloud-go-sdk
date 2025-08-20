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
	SetEnableDeleteTag(v bool) *UpdateArtifactLifecycleRuleRequest
	GetEnableDeleteTag() *bool
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
	// Specifies whether to automatically execute the lifecycle management rule.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// Specifies whether to enable lifecycle management for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-r6ym0lerldp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test_1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The number of images that you want to retain.
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
	// REPO
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that indicates which image tags you want to retain.
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

func (s *UpdateArtifactLifecycleRuleRequest) GetEnableDeleteTag() *bool {
	return s.EnableDeleteTag
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

func (s *UpdateArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
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
