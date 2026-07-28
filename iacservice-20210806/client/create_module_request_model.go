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
	// The idempotency parameter. We recommend that you use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the template. The description can be up to 256 characters in length.
	//
	// example:
	//
	// ECS instance module
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The project group information to which the template belongs.
	GroupInfo *CreateModuleRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// The name of the template. The name must meet the following requirements:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). The name cannot start or end with a hyphen, underscore, or period.
	//
	// - The name must be unique among all templates under the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-ecs-module
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source from which the template is created. Valid values:
	//
	// - OSS: imports from a ZIP file stored in OSS.
	//
	// - Registry: creates from a module in the template registry.
	//
	// - ExportTask: references a template exported by a resource export task.
	//
	// - Editor: creates a blank template that supports online editing.
	//
	// - Upload: uploads a local template file to generate the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path of the template source. This parameter takes effect when source is set to Registry, OSS, or ExportTask.
	//
	// - If source is set to Registry, the value is in the format of \\<workspace name>/\\<module name>:\\<module version>. Example: terraform-alicloud-modules/rds:1.0.0.
	//
	// - If source is set to OSS, the value is in the format of oss::<file URL>. The file must be a ZIP file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If source is set to ExportTask, the value is in the format of \\<export task ID>:\\<exported version>. Example: ex-3b6cb9fa4751afff298da723c24ac:v1.
	//
	// - If source is set to Editor or Upload, leave this parameter empty.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path of the State file that corresponds to the template. This parameter is valid only when source is set to OSS.
	//
	// The value is in the format of oss::\\<OSS file path>/terraform.tfstate.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// The list of tags for the template.
	Tags []*CreateModuleRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The version generation strategy. Valid values:
	//
	// - Manual: manually generates a version. This is the default value.
	//
	// - SourcePathUpdated: generates a new version when sourcePath is modified.
	//
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

type CreateModuleRequestGroupInfo struct {
	// The group ID.
	//
	// example:
	//
	// g-5fd38c9b92d541a7083a86432e2
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The project ID.
	//
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
