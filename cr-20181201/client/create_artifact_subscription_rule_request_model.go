// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactSubscriptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerate(v bool) *CreateArtifactSubscriptionRuleRequest
	GetAccelerate() *bool
	SetInstanceId(v string) *CreateArtifactSubscriptionRuleRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateArtifactSubscriptionRuleRequest
	GetNamespaceName() *string
	SetOverride(v bool) *CreateArtifactSubscriptionRuleRequest
	GetOverride() *bool
	SetPlatform(v []*string) *CreateArtifactSubscriptionRuleRequest
	GetPlatform() []*string
	SetRepoName(v string) *CreateArtifactSubscriptionRuleRequest
	GetRepoName() *string
	SetSourceNamespaceName(v string) *CreateArtifactSubscriptionRuleRequest
	GetSourceNamespaceName() *string
	SetSourceProvider(v string) *CreateArtifactSubscriptionRuleRequest
	GetSourceProvider() *string
	SetSourceRepoName(v string) *CreateArtifactSubscriptionRuleRequest
	GetSourceRepoName() *string
	SetTagCount(v int64) *CreateArtifactSubscriptionRuleRequest
	GetTagCount() *int64
	SetTagList(v []*string) *CreateArtifactSubscriptionRuleRequest
	GetTagList() []*string
	SetTagRegexp(v string) *CreateArtifactSubscriptionRuleRequest
	GetTagRegexp() *string
}

type CreateArtifactSubscriptionRuleRequest struct {
	// Specifies whether to enable the acceleration link. The subscription acceleration feature is in public preview. Based on scheduling policies and network link optimization, it can improve the image subscription speed.
	//
	// example:
	//
	// true
	Accelerate *bool `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
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
	// This parameter is required.
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
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system or architecture. If the source repository contains multi-architecture images, only images with the specified operating system or architecture are subscribed to the destination repository of the Enterprise instance.
	//
	// This parameter is required.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The ACR repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The source namespace. If `SourceProvider` is set to `DOCKER_HUB`, the default value is `library`. This parameter is required in other scenarios.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The artifact source.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The source repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The number of images to subscribe to.
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// - The image tags to subscribe to when matching by tag list. A maximum of 10 tags are supported.
	//
	// - If TagList is specified, these tags are queried one by one. TagRegexp and TagCount cannot be specified at the same time. If TagList is not specified, TagRegexp and TagCount are used for matching, and both must be specified.
	//
	// - The matching method cannot be changed after the rule is created. To change the matching method, delete the rule and create a new one.
	//
	// example:
	//
	// latest
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The image versions to subscribe to from the source repository. Regular expressions are supported.
	//
	// example:
	//
	// release-v.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s CreateArtifactSubscriptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionRuleRequest) GetAccelerate() *bool {
	return s.Accelerate
}

func (s *CreateArtifactSubscriptionRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateArtifactSubscriptionRuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateArtifactSubscriptionRuleRequest) GetOverride() *bool {
	return s.Override
}

func (s *CreateArtifactSubscriptionRuleRequest) GetPlatform() []*string {
	return s.Platform
}

func (s *CreateArtifactSubscriptionRuleRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateArtifactSubscriptionRuleRequest) GetSourceNamespaceName() *string {
	return s.SourceNamespaceName
}

func (s *CreateArtifactSubscriptionRuleRequest) GetSourceProvider() *string {
	return s.SourceProvider
}

func (s *CreateArtifactSubscriptionRuleRequest) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *CreateArtifactSubscriptionRuleRequest) GetTagCount() *int64 {
	return s.TagCount
}

func (s *CreateArtifactSubscriptionRuleRequest) GetTagList() []*string {
	return s.TagList
}

func (s *CreateArtifactSubscriptionRuleRequest) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *CreateArtifactSubscriptionRuleRequest) SetAccelerate(v bool) *CreateArtifactSubscriptionRuleRequest {
	s.Accelerate = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetInstanceId(v string) *CreateArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetNamespaceName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetOverride(v bool) *CreateArtifactSubscriptionRuleRequest {
	s.Override = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetPlatform(v []*string) *CreateArtifactSubscriptionRuleRequest {
	s.Platform = v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetRepoName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetSourceNamespaceName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.SourceNamespaceName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetSourceProvider(v string) *CreateArtifactSubscriptionRuleRequest {
	s.SourceProvider = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetSourceRepoName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.SourceRepoName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetTagCount(v int64) *CreateArtifactSubscriptionRuleRequest {
	s.TagCount = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetTagList(v []*string) *CreateArtifactSubscriptionRuleRequest {
	s.TagList = v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetTagRegexp(v string) *CreateArtifactSubscriptionRuleRequest {
	s.TagRegexp = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
