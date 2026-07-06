// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerate(v bool) *GetArtifactSubscriptionRuleResponseBody
	GetAccelerate() *bool
	SetCode(v string) *GetArtifactSubscriptionRuleResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetArtifactSubscriptionRuleResponseBody
	GetCreateTime() *int64
	SetInstanceId(v string) *GetArtifactSubscriptionRuleResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetArtifactSubscriptionRuleResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetArtifactSubscriptionRuleResponseBody
	GetModifiedTime() *int64
	SetNamespaceName(v string) *GetArtifactSubscriptionRuleResponseBody
	GetNamespaceName() *string
	SetOverride(v bool) *GetArtifactSubscriptionRuleResponseBody
	GetOverride() *bool
	SetPlatform(v []*string) *GetArtifactSubscriptionRuleResponseBody
	GetPlatform() []*string
	SetRepoName(v string) *GetArtifactSubscriptionRuleResponseBody
	GetRepoName() *string
	SetRequestId(v string) *GetArtifactSubscriptionRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *GetArtifactSubscriptionRuleResponseBody
	GetRuleId() *string
	SetSourceDomain(v string) *GetArtifactSubscriptionRuleResponseBody
	GetSourceDomain() *string
	SetSourceNamespaceName(v string) *GetArtifactSubscriptionRuleResponseBody
	GetSourceNamespaceName() *string
	SetSourceProvider(v string) *GetArtifactSubscriptionRuleResponseBody
	GetSourceProvider() *string
	SetSourceRepoName(v string) *GetArtifactSubscriptionRuleResponseBody
	GetSourceRepoName() *string
	SetTagCount(v int64) *GetArtifactSubscriptionRuleResponseBody
	GetTagCount() *int64
	SetTagRegexp(v string) *GetArtifactSubscriptionRuleResponseBody
	GetTagRegexp() *string
}

type GetArtifactSubscriptionRuleResponseBody struct {
	// Indicates whether to enable the accelerated data transfer feature. This feature is in public preview. It optimizes scheduling policies and network paths to improve the speed of artifact subscription.
	//
	// example:
	//
	// true
	Accelerate *bool `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the rule was created.
	//
	// example:
	//
	// 1570759546000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request succeeded.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the rule was last modified.
	//
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The destination ACR namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Indicates whether to overwrite the existing images that have the same tag in the destination repository.
	//
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating systems and architectures. If a source repository contains multi-architecture images, only images that match the specified platforms are synchronized to the destination repository of the Enterprise Edition instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The destination ACR repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D4978DCC-ECBD-40B0-A714-EE6959B22C77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The domain name of the artifact source.
	SourceDomain *string `json:"SourceDomain,omitempty" xml:"SourceDomain,omitempty"`
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
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The regular expression that is used to match the tags of images in the source repository for subscription.
	//
	// example:
	//
	// release-v.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s GetArtifactSubscriptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetAccelerate() *bool {
	return s.Accelerate
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetOverride() *bool {
	return s.Override
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetPlatform() []*string {
	return s.Platform
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetRepoName() *string {
	return s.RepoName
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetSourceDomain() *string {
	return s.SourceDomain
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetSourceNamespaceName() *string {
	return s.SourceNamespaceName
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetSourceProvider() *string {
	return s.SourceProvider
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetTagCount() *int64 {
	return s.TagCount
}

func (s *GetArtifactSubscriptionRuleResponseBody) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetAccelerate(v bool) *GetArtifactSubscriptionRuleResponseBody {
	s.Accelerate = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetCode(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetCreateTime(v int64) *GetArtifactSubscriptionRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetInstanceId(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *GetArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetModifiedTime(v int64) *GetArtifactSubscriptionRuleResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetNamespaceName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetOverride(v bool) *GetArtifactSubscriptionRuleResponseBody {
	s.Override = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetPlatform(v []*string) *GetArtifactSubscriptionRuleResponseBody {
	s.Platform = v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetRepoName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetRuleId(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceDomain(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceDomain = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceNamespaceName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceNamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceProvider(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceProvider = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceRepoName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceRepoName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetTagCount(v int64) *GetArtifactSubscriptionRuleResponseBody {
	s.TagCount = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetTagRegexp(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.TagRegexp = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
