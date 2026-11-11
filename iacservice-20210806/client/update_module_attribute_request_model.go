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
	// The idempotence token. Format: [0-9a-zA-Z-]{1,64}. Use a UUID.
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
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The project group information.
	GroupInfo *UpdateModuleAttributeRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// The name of the template. The name must meet the following requirements:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). It cannot start or end with a hyphen, underscore, or period.
	//
	// - The name must be unique among the template resources within the current account.
	//
	// example:
	//
	// ModuleName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The path value of the template source.
	//
	// - If the source is Registry, the value is in the <Workspace name>/<Module name>:<Module version> format. Example: terraform-alicloud-modules/rds:1.0.0.
	//
	// - If the source is OSS, the value is in the oss::<File URL> format. The file must be a ZIP file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If the source is ExportTask, the value is in the <Export task ID>:<Exported version> format. Example: ex-3b6cb9fa4751afff298da723c24ac:v1.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path value of the state file corresponding to the template. Currently, only OSS paths are supported. The value is in the oss::<OSS path of the file>/terraform.tfstate format.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// The list of tags for the template.
	Tags []*UpdateModuleAttributeRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The version generation strategy. Valid values:
	//
	// - Manual: manually generated version (default).
	//
	// - SourcePathUpdated: a new version is generated when sourcePath is modified.
	//
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

type UpdateModuleAttributeRequestGroupInfo struct {
	// The group ID.
	//
	// example:
	//
	// g-433aead7560571e66e31274ffd3
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-433aead75605713865c386cb9d
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
