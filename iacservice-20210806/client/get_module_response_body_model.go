// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *GetModuleResponseBodyModule) *GetModuleResponseBody
	GetModule() *GetModuleResponseBodyModule
	SetRequestId(v string) *GetModuleResponseBody
	GetRequestId() *string
}

type GetModuleResponseBody struct {
	Module *GetModuleResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1E7BA3EB-B0EF-53F5-9999-07CAD6D9F8A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBody) GetModule() *GetModuleResponseBodyModule {
	return s.Module
}

func (s *GetModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModuleResponseBody) SetModule(v *GetModuleResponseBodyModule) *GetModuleResponseBody {
	s.Module = v
	return s
}

func (s *GetModuleResponseBody) SetRequestId(v string) *GetModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetModuleResponseBodyModule struct {
	// example:
	//
	// 2022-09-06T06:11:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test1
	Description *string                               `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *GetModuleResponseBodyModuleGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// example:
	//
	// v1
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// example:
	//
	// mod-4267dcfbf1b6d14625614ddbe15
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// /
	OutputPath *string `json:"outputPath,omitempty" xml:"outputPath,omitempty"`
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
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// example:
	//
	// Created
	Status *string                            `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*GetModuleResponseBodyModuleTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s GetModuleResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetModuleResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBodyModule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetModuleResponseBodyModule) GetDescription() *string {
	return s.Description
}

func (s *GetModuleResponseBodyModule) GetGroupInfo() *GetModuleResponseBodyModuleGroupInfo {
	return s.GroupInfo
}

func (s *GetModuleResponseBodyModule) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetModuleResponseBodyModule) GetModuleId() *string {
	return s.ModuleId
}

func (s *GetModuleResponseBodyModule) GetName() *string {
	return s.Name
}

func (s *GetModuleResponseBodyModule) GetOutputPath() *string {
	return s.OutputPath
}

func (s *GetModuleResponseBodyModule) GetSource() *string {
	return s.Source
}

func (s *GetModuleResponseBodyModule) GetSourcePath() *string {
	return s.SourcePath
}

func (s *GetModuleResponseBodyModule) GetStatePath() *string {
	return s.StatePath
}

func (s *GetModuleResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *GetModuleResponseBodyModule) GetTags() []*GetModuleResponseBodyModuleTags {
	return s.Tags
}

func (s *GetModuleResponseBodyModule) GetVersionStrategy() *string {
	return s.VersionStrategy
}

func (s *GetModuleResponseBodyModule) SetCreateTime(v string) *GetModuleResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetDescription(v string) *GetModuleResponseBodyModule {
	s.Description = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetGroupInfo(v *GetModuleResponseBodyModuleGroupInfo) *GetModuleResponseBodyModule {
	s.GroupInfo = v
	return s
}

func (s *GetModuleResponseBodyModule) SetLatestVersion(v string) *GetModuleResponseBodyModule {
	s.LatestVersion = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetModuleId(v string) *GetModuleResponseBodyModule {
	s.ModuleId = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetName(v string) *GetModuleResponseBodyModule {
	s.Name = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetOutputPath(v string) *GetModuleResponseBodyModule {
	s.OutputPath = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetSource(v string) *GetModuleResponseBodyModule {
	s.Source = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetSourcePath(v string) *GetModuleResponseBodyModule {
	s.SourcePath = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetStatePath(v string) *GetModuleResponseBodyModule {
	s.StatePath = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetStatus(v string) *GetModuleResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetTags(v []*GetModuleResponseBodyModuleTags) *GetModuleResponseBodyModule {
	s.Tags = v
	return s
}

func (s *GetModuleResponseBodyModule) SetVersionStrategy(v string) *GetModuleResponseBodyModule {
	s.VersionStrategy = &v
	return s
}

func (s *GetModuleResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type GetModuleResponseBodyModuleGroupInfo struct {
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName   *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s GetModuleResponseBodyModuleGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s GetModuleResponseBodyModuleGroupInfo) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBodyModuleGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *GetModuleResponseBodyModuleGroupInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *GetModuleResponseBodyModuleGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetModuleResponseBodyModuleGroupInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetModuleResponseBodyModuleGroupInfo) SetGroupId(v string) *GetModuleResponseBodyModuleGroupInfo {
	s.GroupId = &v
	return s
}

func (s *GetModuleResponseBodyModuleGroupInfo) SetGroupName(v string) *GetModuleResponseBodyModuleGroupInfo {
	s.GroupName = &v
	return s
}

func (s *GetModuleResponseBodyModuleGroupInfo) SetProjectId(v string) *GetModuleResponseBodyModuleGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *GetModuleResponseBodyModuleGroupInfo) SetProjectName(v string) *GetModuleResponseBodyModuleGroupInfo {
	s.ProjectName = &v
	return s
}

func (s *GetModuleResponseBodyModuleGroupInfo) Validate() error {
	return dara.Validate(s)
}

type GetModuleResponseBodyModuleTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s GetModuleResponseBodyModuleTags) String() string {
	return dara.Prettify(s)
}

func (s GetModuleResponseBodyModuleTags) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBodyModuleTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetModuleResponseBodyModuleTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetModuleResponseBodyModuleTags) SetTagKey(v string) *GetModuleResponseBodyModuleTags {
	s.TagKey = &v
	return s
}

func (s *GetModuleResponseBodyModuleTags) SetTagValue(v string) *GetModuleResponseBodyModuleTags {
	s.TagValue = &v
	return s
}

func (s *GetModuleResponseBodyModuleTags) Validate() error {
	return dara.Validate(s)
}
