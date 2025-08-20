// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoBuildRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildArgs(v []*string) *CreateRepoBuildRuleRequest
	GetBuildArgs() []*string
	SetDockerfileLocation(v string) *CreateRepoBuildRuleRequest
	GetDockerfileLocation() *string
	SetDockerfileName(v string) *CreateRepoBuildRuleRequest
	GetDockerfileName() *string
	SetImageTag(v string) *CreateRepoBuildRuleRequest
	GetImageTag() *string
	SetInstanceId(v string) *CreateRepoBuildRuleRequest
	GetInstanceId() *string
	SetPlatforms(v []*string) *CreateRepoBuildRuleRequest
	GetPlatforms() []*string
	SetPushName(v string) *CreateRepoBuildRuleRequest
	GetPushName() *string
	SetPushType(v string) *CreateRepoBuildRuleRequest
	GetPushType() *string
	SetRepoId(v string) *CreateRepoBuildRuleRequest
	GetRepoId() *string
}

type CreateRepoBuildRuleRequest struct {
	// Building arguments.
	BuildArgs []*string `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
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
	// This parameter is required.
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
	// cri-xkx6vujuhay0****
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
	// 	- `inux/arm/v6`
	//
	// Default value: `linux/amd64`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	//
	// This parameter is required.
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
	// This parameter is required.
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
	// crr-8dz3aedjqlmk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateRepoBuildRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoBuildRuleRequest) GetBuildArgs() []*string {
	return s.BuildArgs
}

func (s *CreateRepoBuildRuleRequest) GetDockerfileLocation() *string {
	return s.DockerfileLocation
}

func (s *CreateRepoBuildRuleRequest) GetDockerfileName() *string {
	return s.DockerfileName
}

func (s *CreateRepoBuildRuleRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *CreateRepoBuildRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoBuildRuleRequest) GetPlatforms() []*string {
	return s.Platforms
}

func (s *CreateRepoBuildRuleRequest) GetPushName() *string {
	return s.PushName
}

func (s *CreateRepoBuildRuleRequest) GetPushType() *string {
	return s.PushType
}

func (s *CreateRepoBuildRuleRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepoBuildRuleRequest) SetBuildArgs(v []*string) *CreateRepoBuildRuleRequest {
	s.BuildArgs = v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetDockerfileLocation(v string) *CreateRepoBuildRuleRequest {
	s.DockerfileLocation = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetDockerfileName(v string) *CreateRepoBuildRuleRequest {
	s.DockerfileName = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetImageTag(v string) *CreateRepoBuildRuleRequest {
	s.ImageTag = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetInstanceId(v string) *CreateRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetPlatforms(v []*string) *CreateRepoBuildRuleRequest {
	s.Platforms = v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetPushName(v string) *CreateRepoBuildRuleRequest {
	s.PushName = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetPushType(v string) *CreateRepoBuildRuleRequest {
	s.PushType = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetRepoId(v string) *CreateRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) Validate() error {
	return dara.Validate(s)
}
