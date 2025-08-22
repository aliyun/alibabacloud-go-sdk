// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateModuleAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateModuleAttributeRequest
	GetDescription() *string
	SetGroupInfo(v *UpdateModuleAttributeRequestGroupInfo) *UpdateModuleAttributeRequest
	GetGroupInfo() *UpdateModuleAttributeRequestGroupInfo
	SetName(v string) *UpdateModuleAttributeRequest
	GetName() *string
	SetSourcePath(v string) *UpdateModuleAttributeRequest
	GetSourcePath() *string
	SetStatePath(v string) *UpdateModuleAttributeRequest
	GetStatePath() *string
	SetTags(v []*UpdateModuleAttributeRequestTags) *UpdateModuleAttributeRequest
	GetTags() []*UpdateModuleAttributeRequestTags
	SetVersionStrategy(v string) *UpdateModuleAttributeRequest
	GetVersionStrategy() *string
}

type UpdateModuleAttributeRequest struct {
	// This parameter is required.
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string                                `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *UpdateModuleAttributeRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	StatePath *string                             `json:"statePath,omitempty" xml:"statePath,omitempty"`
	Tags      []*UpdateModuleAttributeRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s UpdateModuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModuleAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModuleAttributeRequest) GetGroupInfo() *UpdateModuleAttributeRequestGroupInfo {
	return s.GroupInfo
}

func (s *UpdateModuleAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateModuleAttributeRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *UpdateModuleAttributeRequest) GetStatePath() *string {
	return s.StatePath
}

func (s *UpdateModuleAttributeRequest) GetTags() []*UpdateModuleAttributeRequestTags {
	return s.Tags
}

func (s *UpdateModuleAttributeRequest) GetVersionStrategy() *string {
	return s.VersionStrategy
}

func (s *UpdateModuleAttributeRequest) SetClientToken(v string) *UpdateModuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetDescription(v string) *UpdateModuleAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetGroupInfo(v *UpdateModuleAttributeRequestGroupInfo) *UpdateModuleAttributeRequest {
	s.GroupInfo = v
	return s
}

func (s *UpdateModuleAttributeRequest) SetName(v string) *UpdateModuleAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetSourcePath(v string) *UpdateModuleAttributeRequest {
	s.SourcePath = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetStatePath(v string) *UpdateModuleAttributeRequest {
	s.StatePath = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetTags(v []*UpdateModuleAttributeRequestTags) *UpdateModuleAttributeRequest {
	s.Tags = v
	return s
}

func (s *UpdateModuleAttributeRequest) SetVersionStrategy(v string) *UpdateModuleAttributeRequest {
	s.VersionStrategy = &v
	return s
}

func (s *UpdateModuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateModuleAttributeRequestGroupInfo struct {
	GroupId   *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s UpdateModuleAttributeRequestGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleAttributeRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeRequestGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateModuleAttributeRequestGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateModuleAttributeRequestGroupInfo) SetGroupId(v string) *UpdateModuleAttributeRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *UpdateModuleAttributeRequestGroupInfo) SetProjectId(v string) *UpdateModuleAttributeRequestGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *UpdateModuleAttributeRequestGroupInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateModuleAttributeRequestTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s UpdateModuleAttributeRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleAttributeRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateModuleAttributeRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateModuleAttributeRequestTags) SetTagKey(v string) *UpdateModuleAttributeRequestTags {
	s.TagKey = &v
	return s
}

func (s *UpdateModuleAttributeRequestTags) SetTagValue(v string) *UpdateModuleAttributeRequestTags {
	s.TagValue = &v
	return s
}

func (s *UpdateModuleAttributeRequestTags) Validate() error {
	return dara.Validate(s)
}
