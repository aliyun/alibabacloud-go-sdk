// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactSubscriptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListArtifactSubscriptionRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListArtifactSubscriptionRuleResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListArtifactSubscriptionRuleResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListArtifactSubscriptionRuleResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListArtifactSubscriptionRuleResponseBody
	GetRequestId() *string
	SetRules(v []*ListArtifactSubscriptionRuleResponseBodyRules) *ListArtifactSubscriptionRuleResponseBody
	GetRules() []*ListArtifactSubscriptionRuleResponseBodyRules
	SetTotalCount(v int32) *ListArtifactSubscriptionRuleResponseBody
	GetTotalCount() *int32
}

type ListArtifactSubscriptionRuleResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - `true`: The call was successful.
	//
	// - `false`: The call failed.
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
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 07FC5654-C82A-59FA-A9D1-78B4EE443F86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of rules.
	Rules []*ListArtifactSubscriptionRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactSubscriptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetRules() []*ListArtifactSubscriptionRuleResponseBodyRules {
	return s.Rules
}

func (s *ListArtifactSubscriptionRuleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetCode(v string) *ListArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *ListArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetPageNo(v int32) *ListArtifactSubscriptionRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetPageSize(v int32) *ListArtifactSubscriptionRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *ListArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetRules(v []*ListArtifactSubscriptionRuleResponseBodyRules) *ListArtifactSubscriptionRuleResponseBody {
	s.Rules = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetTotalCount(v int32) *ListArtifactSubscriptionRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) Validate() error {
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

type ListArtifactSubscriptionRuleResponseBodyRules struct {
	// Indicates whether the acceleration link is enabled. The subscription acceleration feature is in public preview and can improve image subscription speed based on scheduling policies and network link optimization.
	//
	// example:
	//
	// true
	Accelerate *bool `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1638187989000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-brlg4cbj2yl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The modification time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1678341923385
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The namespace name.
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
	// The operating system and architecture. If the source repository contains multi-architecture images, only the specified operating system and architecture are subscribed to the destination repository of the Enterprise instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The repository name.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The domain name of the source image repository for the subscription rule.
	//
	// example:
	//
	// anolis-registry.cn-zhangjiakou.cr.aliyuncs.com
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
	// The number of subscribed images.
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The list of subscribed tags.
	//
	// example:
	//
	// latest
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The image versions subscribed from the source repository. Regular expressions are supported.
	//
	// example:
	//
	// release.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s ListArtifactSubscriptionRuleResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionRuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetAccelerate() *bool {
	return s.Accelerate
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetOverride() *bool {
	return s.Override
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetPlatform() []*string {
	return s.Platform
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetRepoName() *string {
	return s.RepoName
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetSourceDomain() *string {
	return s.SourceDomain
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetSourceNamespaceName() *string {
	return s.SourceNamespaceName
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetSourceProvider() *string {
	return s.SourceProvider
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetTagCount() *int64 {
	return s.TagCount
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetTagList() []*string {
	return s.TagList
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) GetTagRegexp() *string {
	return s.TagRegexp
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetAccelerate(v bool) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.Accelerate = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetCreateTime(v int64) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetInstanceId(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetModifiedTime(v int64) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.ModifiedTime = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetNamespaceName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.NamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetOverride(v bool) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.Override = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetPlatform(v []*string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.Platform = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetRepoName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.RepoName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetRuleId(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceDomain(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceDomain = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceNamespaceName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceNamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceProvider(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceProvider = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceRepoName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceRepoName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetTagCount(v int64) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.TagCount = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetTagList(v []*string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.TagList = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetTagRegexp(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.TagRegexp = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
