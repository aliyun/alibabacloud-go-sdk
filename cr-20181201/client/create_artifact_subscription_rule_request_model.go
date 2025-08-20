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
	SetTagRegexp(v string) *CreateArtifactSubscriptionRuleRequest
	GetTagRegexp() *string
}

type CreateArtifactSubscriptionRuleRequest struct {
	// Indicates whether an acceleration link is enabled for image subscription. The subscription acceleration feature is in public preview. The feature is optimized based on scheduling policies and network links to accelerate image subscription.
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
	// The name of the Container Registry namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Indicates whether the original image is overwritten.
	//
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains a multi-arch image, only the specified operating system and architecture are subscribed to the destination repository of the Enterprise Edition instance. Subscribe to the destination repository of an Enterprise Edition instance
	//
	// This parameter is required.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The name of the Container Registry repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The source of the artifact.
	//
	// Valid values:
	//
	// 	- DOCKER_HUB: Docker Hub
	//
	// 	- GCR: GCR
	//
	// 	- QUAY: Quay.io
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
	// The number of subscribed images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The image tag in the subscription source repository. Regular expressions are supported.
	//
	// This parameter is required.
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

func (s *CreateArtifactSubscriptionRuleRequest) SetTagRegexp(v string) *CreateArtifactSubscriptionRuleRequest {
	s.TagRegexp = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
