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
	// The template information.
	Module *GetModuleResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// The request ID.
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModuleResponseBodyModule struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-09-06T06:11:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The template description.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The group information.
	GroupInfo *GetModuleResponseBodyModuleGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// The latest version number.
	//
	// example:
	//
	// v1
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// The template ID.
	//
	// example:
	//
	// mod-4267dcfbf1b6d14625614ddbe15
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// The template name.
	//
	// example:
	//
	// ModuleName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The storage path of the template.
	//
	// example:
	//
	// /
	OutputPath *string `json:"outputPath,omitempty" xml:"outputPath,omitempty"`
	// The template source. Valid values:
	//
	// - OSS: Imported from OSS.
	//
	// - Registry: Created from a template in the template center.
	//
	// - ExportTask: Exported from a resource export task.
	//
	// - Upload: Uploaded as a file.
	//
	// - Shared: Cloned from a shared template.
	//
	// - Editor: Created by using the online editor.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The source path of the template.
	//
	// - If the source is Registry, the value is in the format of <workspace name>/<module name>:<module version>, such as terraform-alicloud-modules/rds:1.0.0.
	//
	// - If the source is OSS, the value is in the format of oss::<file link>, such as oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If the source is ExportTask, the value is in the format of <export task ID>:<exported version>, such as ex-3b6cb9fa4751afff298da723c24ac:v1.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path of the state file that corresponds to the template. Currently, only OSS paths are supported. The value is in the format of oss::<file OSS path>/terraform.tfstate.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// The template status. Valid values:
	//
	// - Creating: The template is being created.
	//
	// - Created: The template is created.
	//
	// After the template is created, you can publish a version.
	//
	// example:
	//
	// Created
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags of the template.
	Tags []*GetModuleResponseBodyModuleTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The version generation strategy. Valid values:
	//
	// - Manual: Versions are generated manually. This is the default value.
	//
	// - SourcePathUpdated: A new version is generated when the sourcePath is modified.
	//
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
	if s.GroupInfo != nil {
		if err := s.GroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModuleResponseBodyModuleGroupInfo struct {
	// The group ID.
	//
	// example:
	//
	// g-fu1a1ol8cob1oni01ekcloi
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// groupName
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-al1d11jlvlsbvr11lf3pqo
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// projectName
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
	// The tag key of the template.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the template.
	//
	// example:
	//
	// TestValue
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
