// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoBuildRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildArgs(v []*string) *UpdateRepoBuildRuleRequest
	GetBuildArgs() []*string
	SetBuildRuleId(v string) *UpdateRepoBuildRuleRequest
	GetBuildRuleId() *string
	SetDockerfileLocation(v string) *UpdateRepoBuildRuleRequest
	GetDockerfileLocation() *string
	SetDockerfileName(v string) *UpdateRepoBuildRuleRequest
	GetDockerfileName() *string
	SetImageTag(v string) *UpdateRepoBuildRuleRequest
	GetImageTag() *string
	SetInstanceId(v string) *UpdateRepoBuildRuleRequest
	GetInstanceId() *string
	SetPlatforms(v []*string) *UpdateRepoBuildRuleRequest
	GetPlatforms() []*string
	SetPushName(v string) *UpdateRepoBuildRuleRequest
	GetPushName() *string
	SetPushType(v string) *UpdateRepoBuildRuleRequest
	GetPushType() *string
	SetRepoId(v string) *UpdateRepoBuildRuleRequest
	GetRepoId() *string
}

type UpdateRepoBuildRuleRequest struct {
	// Building arguments.
	BuildArgs []*string `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The ID of the building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The path of the Dockerfile.
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
	// v0.9.5
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Architecture for image building. Valid values:
	//
	// 	- `linux/amd64`
	//
	// 	- `linux/arm64`
	//
	// 	- `linux/386`
	//
	// 	- `linux/arm/v7`
	//
	// 	- `linux/arm/v6`
	//
	// Default value: `linux/amd64`
	//
	// example:
	//
	// linux/amd64
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	//
	// example:
	//
	// master
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the building rule. Valid values:
	//
	// 	- `GIT_TAG`: tag push
	//
	// 	- `GIT_BRANCH`: branch push
	//
	// example:
	//
	// GIT_BRANCH
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s UpdateRepoBuildRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepoBuildRuleRequest) GetBuildArgs() []*string {
	return s.BuildArgs
}

func (s *UpdateRepoBuildRuleRequest) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *UpdateRepoBuildRuleRequest) GetDockerfileLocation() *string {
	return s.DockerfileLocation
}

func (s *UpdateRepoBuildRuleRequest) GetDockerfileName() *string {
	return s.DockerfileName
}

func (s *UpdateRepoBuildRuleRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *UpdateRepoBuildRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRepoBuildRuleRequest) GetPlatforms() []*string {
	return s.Platforms
}

func (s *UpdateRepoBuildRuleRequest) GetPushName() *string {
	return s.PushName
}

func (s *UpdateRepoBuildRuleRequest) GetPushType() *string {
	return s.PushType
}

func (s *UpdateRepoBuildRuleRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *UpdateRepoBuildRuleRequest) SetBuildArgs(v []*string) *UpdateRepoBuildRuleRequest {
	s.BuildArgs = v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetBuildRuleId(v string) *UpdateRepoBuildRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetDockerfileLocation(v string) *UpdateRepoBuildRuleRequest {
	s.DockerfileLocation = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetDockerfileName(v string) *UpdateRepoBuildRuleRequest {
	s.DockerfileName = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetImageTag(v string) *UpdateRepoBuildRuleRequest {
	s.ImageTag = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetInstanceId(v string) *UpdateRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetPlatforms(v []*string) *UpdateRepoBuildRuleRequest {
	s.Platforms = v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetPushName(v string) *UpdateRepoBuildRuleRequest {
	s.PushName = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetPushType(v string) *UpdateRepoBuildRuleRequest {
	s.PushType = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetRepoId(v string) *UpdateRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) Validate() error {
	return dara.Validate(s)
}
