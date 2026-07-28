// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoApply(v bool) *UpdateTaskAttributeRequest
	GetAutoApply() *bool
	SetAutoDestroy(v bool) *UpdateTaskAttributeRequest
	GetAutoDestroy() *bool
	SetClientToken(v string) *UpdateTaskAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTaskAttributeRequest
	GetDescription() *string
	SetGroupInfo(v *UpdateTaskAttributeRequestGroupInfo) *UpdateTaskAttributeRequest
	GetGroupInfo() *UpdateTaskAttributeRequestGroupInfo
	SetInitModuleState(v bool) *UpdateTaskAttributeRequest
	GetInitModuleState() *bool
	SetModuleVersion(v string) *UpdateTaskAttributeRequest
	GetModuleVersion() *string
	SetName(v string) *UpdateTaskAttributeRequest
	GetName() *string
	SetProtectionStrategy(v []*string) *UpdateTaskAttributeRequest
	GetProtectionStrategy() []*string
	SetRamRole(v string) *UpdateTaskAttributeRequest
	GetRamRole() *string
	SetSkipPropertyValidation(v bool) *UpdateTaskAttributeRequest
	GetSkipPropertyValidation() *bool
	SetSkipRegionValidation(v bool) *UpdateTaskAttributeRequest
	GetSkipRegionValidation() *bool
	SetTags(v []*UpdateTaskAttributeRequestTags) *UpdateTaskAttributeRequest
	GetTags() []*UpdateTaskAttributeRequestTags
	SetTerraformProviderVersion(v string) *UpdateTaskAttributeRequest
	GetTerraformProviderVersion() *string
	SetTerraformVersion(v string) *UpdateTaskAttributeRequest
	GetTerraformVersion() *string
	SetTriggerStrategy(v string) *UpdateTaskAttributeRequest
	GetTriggerStrategy() *string
}

type UpdateTaskAttributeRequest struct {
	// Specifies whether to automatically execute the task. Default value: false.
	//
	// - true: After the preview is complete (terraform plan), the execution (terraform apply) is automatically performed without manual confirmation.
	//
	// - false: After the preview is complete (terraform plan), manual confirmation is required before the execution (terraform apply) starts.
	//
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// Specifies whether to automatically destroy resources after creation. Default value: false.
	//
	// - true: After the execution is complete (terraform apply), the destroy operation (terraform destroy) is automatically performed without manual confirmation.
	//
	// - false: After the execution is complete (terraform apply), no further action is taken.
	//
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. We recommend that you use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The project group information.
	GroupInfo *UpdateTaskAttributeRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// Specifies whether to use a state file. Default value: false. This parameter is applicable when the template originates from resource export. Only one task can use this parameter.
	//
	// example:
	//
	// false
	InitModuleState *bool `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// The template version.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The task name. The name must meet the following requirements:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). It cannot start or end with a hyphen, underscore, or period.
	//
	// - The name must be unique among all tasks under the current account.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of resource protection strategies.
	ProtectionStrategy []*string `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// The RAM role. The system assumes this role to execute the template when a new job is triggered. This parameter is required when the job trigger method is not manual.
	//
	// example:
	//
	// role
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// Specifies whether to skip enum value validation. Default value: false.
	//
	// example:
	//
	// false
	SkipPropertyValidation *bool `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	SkipRegionValidation   *bool `json:"skipRegionValidation,omitempty" xml:"skipRegionValidation,omitempty"`
	// The list of tags for the task.
	Tags                     []*UpdateTaskAttributeRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TerraformProviderVersion *string                           `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform version. Call the **ListAvailableTerraformVersions*	- operation to obtain the list of supported versions. Default value: 1.5.7.
	//
	// example:
	//
	// 1.5.7
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// The job trigger method. Valid values:
	//
	// - Manual: manually triggered (default).
	//
	// - NewVersion: triggered when a new template version is published.
	//
	// - ParameterSetUpdated: triggered when the parameter set content changes or the parameter set attach relationship changes.
	//
	// - Auto: automatically triggered when the task\\"s own properties change, such as task creation, execution version change, or job trigger policy change (when changed from another value to Auto).
	//
	// The **ramRole*	- parameter is required when the trigger method is not manual.
	//
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s UpdateTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequest) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *UpdateTaskAttributeRequest) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *UpdateTaskAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTaskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskAttributeRequest) GetGroupInfo() *UpdateTaskAttributeRequestGroupInfo {
	return s.GroupInfo
}

func (s *UpdateTaskAttributeRequest) GetInitModuleState() *bool {
	return s.InitModuleState
}

func (s *UpdateTaskAttributeRequest) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *UpdateTaskAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAttributeRequest) GetProtectionStrategy() []*string {
	return s.ProtectionStrategy
}

func (s *UpdateTaskAttributeRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateTaskAttributeRequest) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *UpdateTaskAttributeRequest) GetSkipRegionValidation() *bool {
	return s.SkipRegionValidation
}

func (s *UpdateTaskAttributeRequest) GetTags() []*UpdateTaskAttributeRequestTags {
	return s.Tags
}

func (s *UpdateTaskAttributeRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *UpdateTaskAttributeRequest) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *UpdateTaskAttributeRequest) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *UpdateTaskAttributeRequest) SetAutoApply(v bool) *UpdateTaskAttributeRequest {
	s.AutoApply = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetAutoDestroy(v bool) *UpdateTaskAttributeRequest {
	s.AutoDestroy = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetClientToken(v string) *UpdateTaskAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetDescription(v string) *UpdateTaskAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetGroupInfo(v *UpdateTaskAttributeRequestGroupInfo) *UpdateTaskAttributeRequest {
	s.GroupInfo = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetInitModuleState(v bool) *UpdateTaskAttributeRequest {
	s.InitModuleState = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetModuleVersion(v string) *UpdateTaskAttributeRequest {
	s.ModuleVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetName(v string) *UpdateTaskAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetProtectionStrategy(v []*string) *UpdateTaskAttributeRequest {
	s.ProtectionStrategy = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetRamRole(v string) *UpdateTaskAttributeRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetSkipPropertyValidation(v bool) *UpdateTaskAttributeRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetSkipRegionValidation(v bool) *UpdateTaskAttributeRequest {
	s.SkipRegionValidation = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTags(v []*UpdateTaskAttributeRequestTags) *UpdateTaskAttributeRequest {
	s.Tags = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTerraformProviderVersion(v string) *UpdateTaskAttributeRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTerraformVersion(v string) *UpdateTaskAttributeRequest {
	s.TerraformVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTriggerStrategy(v string) *UpdateTaskAttributeRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateTaskAttributeRequest) Validate() error {
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

type UpdateTaskAttributeRequestGroupInfo struct {
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

func (s UpdateTaskAttributeRequestGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequestGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateTaskAttributeRequestGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateTaskAttributeRequestGroupInfo) SetGroupId(v string) *UpdateTaskAttributeRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *UpdateTaskAttributeRequestGroupInfo) SetProjectId(v string) *UpdateTaskAttributeRequestGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *UpdateTaskAttributeRequestGroupInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAttributeRequestTags struct {
	// The tag key of the task.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the task.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s UpdateTaskAttributeRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateTaskAttributeRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateTaskAttributeRequestTags) SetTagKey(v string) *UpdateTaskAttributeRequestTags {
	s.TagKey = &v
	return s
}

func (s *UpdateTaskAttributeRequestTags) SetTagValue(v string) *UpdateTaskAttributeRequestTags {
	s.TagValue = &v
	return s
}

func (s *UpdateTaskAttributeRequestTags) Validate() error {
	return dara.Validate(s)
}
