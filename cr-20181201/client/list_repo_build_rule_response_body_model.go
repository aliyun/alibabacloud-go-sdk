// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRules(v []*ListRepoBuildRuleResponseBodyBuildRules) *ListRepoBuildRuleResponseBody
	GetBuildRules() []*ListRepoBuildRuleResponseBodyBuildRules
	SetCode(v string) *ListRepoBuildRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoBuildRuleResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoBuildRuleResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoBuildRuleResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoBuildRuleResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRepoBuildRuleResponseBody
	GetTotalCount() *string
}

type ListRepoBuildRuleResponseBody struct {
	// The list of image building rules.
	BuildRules []*ListRepoBuildRuleResponseBodyBuildRules `json:"BuildRules,omitempty" xml:"BuildRules,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
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
	// The ID of the request.
	//
	// example:
	//
	// 42D782C8-E8F6-4A32-BEA0-6A6AC854C22A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoBuildRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleResponseBody) GetBuildRules() []*ListRepoBuildRuleResponseBodyBuildRules {
	return s.BuildRules
}

func (s *ListRepoBuildRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoBuildRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoBuildRuleResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoBuildRuleResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoBuildRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoBuildRuleResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRepoBuildRuleResponseBody) SetBuildRules(v []*ListRepoBuildRuleResponseBodyBuildRules) *ListRepoBuildRuleResponseBody {
	s.BuildRules = v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetCode(v string) *ListRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetIsSuccess(v bool) *ListRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetPageNo(v int32) *ListRepoBuildRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetPageSize(v int32) *ListRepoBuildRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetRequestId(v string) *ListRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetTotalCount(v string) *ListRepoBuildRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRepoBuildRuleResponseBodyBuildRules struct {
	BuildArgs []*string `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The ID of the image building rule.
	//
	// example:
	//
	// crbr-khys0nd3asbe****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The directory of the Dockerfile.
	//
	// example:
	//
	// /
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	//
	// example:
	//
	// Dockerfile
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// v0.1
	ImageTag  *string   `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	//
	// example:
	//
	// v0.1
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the image building rule. Valid values:
	//
	// 	- GIT_BRANCH: branch push
	//
	// 	- GIT_TAG: tag push
	//
	// example:
	//
	// GIT_BRANCH
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
}

func (s ListRepoBuildRuleResponseBodyBuildRules) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRuleResponseBodyBuildRules) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetBuildArgs() []*string {
	return s.BuildArgs
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetDockerfileLocation() *string {
	return s.DockerfileLocation
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetDockerfileName() *string {
	return s.DockerfileName
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetImageTag() *string {
	return s.ImageTag
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetPlatforms() []*string {
	return s.Platforms
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetPushName() *string {
	return s.PushName
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) GetPushType() *string {
	return s.PushType
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetBuildArgs(v []*string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.BuildArgs = v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetBuildRuleId(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.BuildRuleId = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetDockerfileLocation(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.DockerfileLocation = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetDockerfileName(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.DockerfileName = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetImageTag(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.ImageTag = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetPlatforms(v []*string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.Platforms = v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetPushName(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.PushName = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetPushType(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.PushType = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) Validate() error {
	return dara.Validate(s)
}
