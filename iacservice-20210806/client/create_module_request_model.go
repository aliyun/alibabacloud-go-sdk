// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateModuleRequest
	GetClientToken() *string
	SetDescription(v string) *CreateModuleRequest
	GetDescription() *string
	SetGroupInfo(v *CreateModuleRequestGroupInfo) *CreateModuleRequest
	GetGroupInfo() *CreateModuleRequestGroupInfo
	SetName(v string) *CreateModuleRequest
	GetName() *string
	SetSource(v string) *CreateModuleRequest
	GetSource() *string
	SetSourcePath(v string) *CreateModuleRequest
	GetSourcePath() *string
	SetStatePath(v string) *CreateModuleRequest
	GetStatePath() *string
	SetTags(v []*CreateModuleRequestTags) *CreateModuleRequest
	GetTags() []*CreateModuleRequestTags
	SetVersionStrategy(v string) *CreateModuleRequest
	GetVersionStrategy() *string
}

type CreateModuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string                       `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *CreateModuleRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// OSS：
	//
	// "oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip"
	//
	// Registry：
	//
	// "alibaba/security-group/alicloud:2.1.0"
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string                    `json:"statePath,omitempty" xml:"statePath,omitempty"`
	Tags      []*CreateModuleRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s CreateModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleRequest) GoString() string {
	return s.String()
}

func (s *CreateModuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModuleRequest) GetGroupInfo() *CreateModuleRequestGroupInfo {
	return s.GroupInfo
}

func (s *CreateModuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateModuleRequest) GetSource() *string {
	return s.Source
}

func (s *CreateModuleRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *CreateModuleRequest) GetStatePath() *string {
	return s.StatePath
}

func (s *CreateModuleRequest) GetTags() []*CreateModuleRequestTags {
	return s.Tags
}

func (s *CreateModuleRequest) GetVersionStrategy() *string {
	return s.VersionStrategy
}

func (s *CreateModuleRequest) SetClientToken(v string) *CreateModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModuleRequest) SetDescription(v string) *CreateModuleRequest {
	s.Description = &v
	return s
}

func (s *CreateModuleRequest) SetGroupInfo(v *CreateModuleRequestGroupInfo) *CreateModuleRequest {
	s.GroupInfo = v
	return s
}

func (s *CreateModuleRequest) SetName(v string) *CreateModuleRequest {
	s.Name = &v
	return s
}

func (s *CreateModuleRequest) SetSource(v string) *CreateModuleRequest {
	s.Source = &v
	return s
}

func (s *CreateModuleRequest) SetSourcePath(v string) *CreateModuleRequest {
	s.SourcePath = &v
	return s
}

func (s *CreateModuleRequest) SetStatePath(v string) *CreateModuleRequest {
	s.StatePath = &v
	return s
}

func (s *CreateModuleRequest) SetTags(v []*CreateModuleRequestTags) *CreateModuleRequest {
	s.Tags = v
	return s
}

func (s *CreateModuleRequest) SetVersionStrategy(v string) *CreateModuleRequest {
	s.VersionStrategy = &v
	return s
}

func (s *CreateModuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateModuleRequestGroupInfo struct {
	// example:
	//
	// g-5fd38c9b92d541a7083a86432e2
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// p-433aead75605713865c386cb9d
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s CreateModuleRequestGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateModuleRequestGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateModuleRequestGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateModuleRequestGroupInfo) SetGroupId(v string) *CreateModuleRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *CreateModuleRequestGroupInfo) SetProjectId(v string) *CreateModuleRequestGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *CreateModuleRequestGroupInfo) Validate() error {
	return dara.Validate(s)
}

type CreateModuleRequestTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s CreateModuleRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleRequestTags) GoString() string {
	return s.String()
}

func (s *CreateModuleRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateModuleRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateModuleRequestTags) SetTagKey(v string) *CreateModuleRequestTags {
	s.TagKey = &v
	return s
}

func (s *CreateModuleRequestTags) SetTagValue(v string) *CreateModuleRequestTags {
	s.TagValue = &v
	return s
}

func (s *CreateModuleRequestTags) Validate() error {
	return dara.Validate(s)
}
