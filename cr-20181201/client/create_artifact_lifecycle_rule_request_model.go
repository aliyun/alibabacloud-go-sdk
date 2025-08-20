// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuto(v bool) *CreateArtifactLifecycleRuleRequest
	GetAuto() *bool
	SetEnableDeleteTag(v bool) *CreateArtifactLifecycleRuleRequest
	GetEnableDeleteTag() *bool
	SetInstanceId(v string) *CreateArtifactLifecycleRuleRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateArtifactLifecycleRuleRequest
	GetNamespaceName() *string
	SetRepoName(v string) *CreateArtifactLifecycleRuleRequest
	GetRepoName() *string
	SetRetentionTagCount(v int64) *CreateArtifactLifecycleRuleRequest
	GetRetentionTagCount() *int64
	SetScheduleTime(v string) *CreateArtifactLifecycleRuleRequest
	GetScheduleTime() *string
	SetScope(v string) *CreateArtifactLifecycleRuleRequest
	GetScope() *string
	SetTagRegexp(v string) *CreateArtifactLifecycleRuleRequest
	GetTagRegexp() *string
}

type CreateArtifactLifecycleRuleRequest struct {
	// Specify whether to automatically execute the lifecycle management rule.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// Specify whether to enable lifecycle management for the artifact.
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
	// cri-gbwfk7qbgrxe****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// dev-backend
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
	// The execution cycle of the lifecycle management rule.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The deletion scope.
	//
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that is used to indicate which image tags are retained.
	//
	// example:
	//
	// release-.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s CreateArtifactLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactLifecycleRuleRequest) GetAuto() *bool {
	return s.Auto
}

func (s *CreateArtifactLifecycleRuleRequest) GetEnableDeleteTag() *bool {
	return s.EnableDeleteTag
}

func (s *CreateArtifactLifecycleRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateArtifactLifecycleRuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateArtifactLifecycleRuleRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateArtifactLifecycleRuleRequest) GetRetentionTagCount() *int64 {
	return s.RetentionTagCount
}

func (s *CreateArtifactLifecycleRuleRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *CreateArtifactLifecycleRuleRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateArtifactLifecycleRuleRequest) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *CreateArtifactLifecycleRuleRequest) SetAuto(v bool) *CreateArtifactLifecycleRuleRequest {
	s.Auto = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *CreateArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetInstanceId(v string) *CreateArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetNamespaceName(v string) *CreateArtifactLifecycleRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetRepoName(v string) *CreateArtifactLifecycleRuleRequest {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetRetentionTagCount(v int64) *CreateArtifactLifecycleRuleRequest {
	s.RetentionTagCount = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetScheduleTime(v string) *CreateArtifactLifecycleRuleRequest {
	s.ScheduleTime = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetScope(v string) *CreateArtifactLifecycleRuleRequest {
	s.Scope = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetTagRegexp(v string) *CreateArtifactLifecycleRuleRequest {
	s.TagRegexp = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
