// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactSubscriptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerate(v string) *UpdateArtifactSubscriptionRuleRequest
	GetAccelerate() *string
	SetInstanceId(v string) *UpdateArtifactSubscriptionRuleRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *UpdateArtifactSubscriptionRuleRequest
	GetNamespaceName() *string
	SetOverride(v string) *UpdateArtifactSubscriptionRuleRequest
	GetOverride() *string
	SetPlatform(v []*string) *UpdateArtifactSubscriptionRuleRequest
	GetPlatform() []*string
	SetRepoName(v string) *UpdateArtifactSubscriptionRuleRequest
	GetRepoName() *string
	SetRuleId(v string) *UpdateArtifactSubscriptionRuleRequest
	GetRuleId() *string
	SetSourceNamespaceName(v string) *UpdateArtifactSubscriptionRuleRequest
	GetSourceNamespaceName() *string
	SetSourceProvider(v string) *UpdateArtifactSubscriptionRuleRequest
	GetSourceProvider() *string
	SetSourceRepoName(v string) *UpdateArtifactSubscriptionRuleRequest
	GetSourceRepoName() *string
	SetTagCount(v int64) *UpdateArtifactSubscriptionRuleRequest
	GetTagCount() *int64
	SetTagList(v []*string) *UpdateArtifactSubscriptionRuleRequest
	GetTagList() []*string
	SetTagRegexp(v string) *UpdateArtifactSubscriptionRuleRequest
	GetTagRegexp() *string
}

type UpdateArtifactSubscriptionRuleRequest struct {
	// Specifies whether to enable the acceleration link. The subscription acceleration feature is in public preview. It improves the image subscription speed based on scheduling policies and network link optimization.
	//
	// example:
	//
	// true
	Accelerate *string `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ACR namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Specifies whether to overwrite images.
	//
	// example:
	//
	// true
	Override *string `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains multi-architecture images, only images with the specified operating system and architecture are subscribed to the destination repository of the Enterprise instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The ACR repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The artifact source.
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The number of images to subscribe to.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// - The image tags to subscribe to when matching by tag list. A maximum of 10 tags are supported.
	//
	// - This parameter is valid only when the rule is created by using tag list matching. It cannot be specified together with TagRegexp or TagCount.
	//
	// - If this parameter is not specified or an empty list is specified, the original tag list is retained and cannot be cleared. The matching method cannot be modified. To modify the matching method, delete the rule and create a new one.
	//
	// example:
	//
	// latest
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The image versions in the source repository to subscribe to. Regular expressions are supported.
	//
	// example:
	//
	// release-v.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s UpdateArtifactSubscriptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetAccelerate() *string {
	return s.Accelerate
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetOverride() *string {
	return s.Override
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetPlatform() []*string {
	return s.Platform
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetSourceNamespaceName() *string {
	return s.SourceNamespaceName
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetSourceProvider() *string {
	return s.SourceProvider
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetTagCount() *int64 {
	return s.TagCount
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetTagList() []*string {
	return s.TagList
}

func (s *UpdateArtifactSubscriptionRuleRequest) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetAccelerate(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.Accelerate = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetInstanceId(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetNamespaceName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetOverride(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.Override = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetPlatform(v []*string) *UpdateArtifactSubscriptionRuleRequest {
	s.Platform = v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetRepoName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetRuleId(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetSourceNamespaceName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.SourceNamespaceName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetSourceProvider(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.SourceProvider = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetSourceRepoName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.SourceRepoName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetTagCount(v int64) *UpdateArtifactSubscriptionRuleRequest {
	s.TagCount = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetTagList(v []*string) *UpdateArtifactSubscriptionRuleRequest {
	s.TagList = v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetTagRegexp(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.TagRegexp = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
