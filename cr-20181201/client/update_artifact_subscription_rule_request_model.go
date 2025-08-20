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
	SetTagRegexp(v string) *UpdateArtifactSubscriptionRuleRequest
	GetTagRegexp() *string
}

type UpdateArtifactSubscriptionRuleRequest struct {
	// Specifies whether to enable an acceleration link for image subscription. The subscription acceleration feature is in public preview. The feature is optimized based on scheduling policies and network links to accelerate image subscription.
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
	// The name of the Container Registry namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Specifies whether to overwrite the original image.
	//
	// example:
	//
	// true
	Override *string `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains multi-arch images, only the images with the specified operating system and architecture are subscribed to the destination repository of the Enterprise Edition instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The name of the Container Registry repository.
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
	// The name of the source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The source of the artifacts.
	//
	// Valid values:
	//
	// 	- DOCKER_HUB: Docker Hub
	//
	// 	- GCR: GCR
	//
	// 	- QUAY: Quay.io
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
	// The number of subscribed images.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The image tags in the subscription source repository. Regular expressions are supported.
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

func (s *UpdateArtifactSubscriptionRuleRequest) SetTagRegexp(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.TagRegexp = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
