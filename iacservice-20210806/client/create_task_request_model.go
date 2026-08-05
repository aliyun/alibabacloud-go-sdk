// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoApply(v bool) *CreateTaskRequest
	GetAutoApply() *bool
	SetAutoDestroy(v bool) *CreateTaskRequest
	GetAutoDestroy() *bool
	SetClientToken(v string) *CreateTaskRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTaskRequest
	GetDescription() *string
	SetGroupInfo(v *CreateTaskRequestGroupInfo) *CreateTaskRequest
	GetGroupInfo() *CreateTaskRequestGroupInfo
	SetInitModuleState(v bool) *CreateTaskRequest
	GetInitModuleState() *bool
	SetModuleId(v string) *CreateTaskRequest
	GetModuleId() *string
	SetModuleVersion(v string) *CreateTaskRequest
	GetModuleVersion() *string
	SetName(v string) *CreateTaskRequest
	GetName() *string
	SetParameterSetIds(v []*string) *CreateTaskRequest
	GetParameterSetIds() []*string
	SetProtectionStrategy(v []*string) *CreateTaskRequest
	GetProtectionStrategy() []*string
	SetRamRole(v string) *CreateTaskRequest
	GetRamRole() *string
	SetSkipPropertyValidation(v bool) *CreateTaskRequest
	GetSkipPropertyValidation() *bool
	SetSkipRegionValidation(v bool) *CreateTaskRequest
	GetSkipRegionValidation() *bool
	SetTags(v []*CreateTaskRequestTags) *CreateTaskRequest
	GetTags() []*CreateTaskRequestTags
	SetTaskBackend(v *CreateTaskRequestTaskBackend) *CreateTaskRequest
	GetTaskBackend() *CreateTaskRequestTaskBackend
	SetTerraformProviderVersion(v string) *CreateTaskRequest
	GetTerraformProviderVersion() *string
	SetTerraformVersion(v string) *CreateTaskRequest
	GetTerraformVersion() *string
	SetTriggerStrategy(v string) *CreateTaskRequest
	GetTriggerStrategy() *string
}

type CreateTaskRequest struct {
	// Specifies whether to automatically execute the node. Default value: false.
	//
	// - true: After the preview is complete (terraform plan), the execution (terraform apply) is automatically performed without manual confirmation.
	//
	// - false: After the preview is complete (terraform plan), manual confirmation is required before the execution (terraform apply) starts.
	//
	// example:
	//
	// false
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
	// The description of the node.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The project group information.
	GroupInfo *CreateTaskRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// Specifies whether to use a state file. Default value: false. This parameter is applicable when the template originates from resource export. Only one node can use this parameter.
	//
	// example:
	//
	// false
	InitModuleState *bool `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mod-144fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// The template version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The node name. The name must meet the following requirements:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). The name cannot start or end with a hyphen, underscore, or period.
	//
	// - The name must be unique among all node resources within the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The collection of associated parameter set IDs.
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// The list of resource protection strategies.
	ProtectionStrategy []*string `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// The RAM role. The system assumes this role to execute the template when a new job is triggered. This parameter is required when the job trigger method is not manual.
	//
	// example:
	//
	// RoleName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// Specifies whether to skip enumeration value validation. Default value: false.
	//
	// example:
	//
	// true
	SkipPropertyValidation *bool `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	SkipRegionValidation   *bool `json:"skipRegionValidation,omitempty" xml:"skipRegionValidation,omitempty"`
	// The list of tags for the node.
	Tags []*CreateTaskRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The node backend configuration. After this parameter is configured, runtime log information is saved to the specified OSS bucket.
	TaskBackend *CreateTaskRequestTaskBackend `json:"taskBackend,omitempty" xml:"taskBackend,omitempty" type:"Struct"`
	// example:
	//
	// 1.248.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform version. Call the **ListAvailableTerraformVersions*	- operation to obtain the list of supported versions. Default value: 1.5.7.
	//
	// example:
	//
	// 1.5.7
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// The job trigger method. Valid values:
	//
	// - Manual: manual trigger (default).
	//
	// - NewVersion: triggered when a new template version is published.
	//
	// - ParameterSetUpdated: triggered when the parameter set content changes or the parameter set attach relationship changes.
	//
	// - Auto: automatically triggered when the node properties change, such as node creation, execution version change, or job trigger policy change (when changed from another value to Auto).
	//
	// The **ramRole*	- parameter is required when the trigger method is not manual.
	//
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *CreateTaskRequest) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *CreateTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTaskRequest) GetGroupInfo() *CreateTaskRequestGroupInfo {
	return s.GroupInfo
}

func (s *CreateTaskRequest) GetInitModuleState() *bool {
	return s.InitModuleState
}

func (s *CreateTaskRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *CreateTaskRequest) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *CreateTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateTaskRequest) GetParameterSetIds() []*string {
	return s.ParameterSetIds
}

func (s *CreateTaskRequest) GetProtectionStrategy() []*string {
	return s.ProtectionStrategy
}

func (s *CreateTaskRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateTaskRequest) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *CreateTaskRequest) GetSkipRegionValidation() *bool {
	return s.SkipRegionValidation
}

func (s *CreateTaskRequest) GetTags() []*CreateTaskRequestTags {
	return s.Tags
}

func (s *CreateTaskRequest) GetTaskBackend() *CreateTaskRequestTaskBackend {
	return s.TaskBackend
}

func (s *CreateTaskRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *CreateTaskRequest) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *CreateTaskRequest) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *CreateTaskRequest) SetAutoApply(v bool) *CreateTaskRequest {
	s.AutoApply = &v
	return s
}

func (s *CreateTaskRequest) SetAutoDestroy(v bool) *CreateTaskRequest {
	s.AutoDestroy = &v
	return s
}

func (s *CreateTaskRequest) SetClientToken(v string) *CreateTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTaskRequest) SetDescription(v string) *CreateTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTaskRequest) SetGroupInfo(v *CreateTaskRequestGroupInfo) *CreateTaskRequest {
	s.GroupInfo = v
	return s
}

func (s *CreateTaskRequest) SetInitModuleState(v bool) *CreateTaskRequest {
	s.InitModuleState = &v
	return s
}

func (s *CreateTaskRequest) SetModuleId(v string) *CreateTaskRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateTaskRequest) SetModuleVersion(v string) *CreateTaskRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CreateTaskRequest) SetName(v string) *CreateTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTaskRequest) SetParameterSetIds(v []*string) *CreateTaskRequest {
	s.ParameterSetIds = v
	return s
}

func (s *CreateTaskRequest) SetProtectionStrategy(v []*string) *CreateTaskRequest {
	s.ProtectionStrategy = v
	return s
}

func (s *CreateTaskRequest) SetRamRole(v string) *CreateTaskRequest {
	s.RamRole = &v
	return s
}

func (s *CreateTaskRequest) SetSkipPropertyValidation(v bool) *CreateTaskRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *CreateTaskRequest) SetSkipRegionValidation(v bool) *CreateTaskRequest {
	s.SkipRegionValidation = &v
	return s
}

func (s *CreateTaskRequest) SetTags(v []*CreateTaskRequestTags) *CreateTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateTaskRequest) SetTaskBackend(v *CreateTaskRequestTaskBackend) *CreateTaskRequest {
	s.TaskBackend = v
	return s
}

func (s *CreateTaskRequest) SetTerraformProviderVersion(v string) *CreateTaskRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateTaskRequest) SetTerraformVersion(v string) *CreateTaskRequest {
	s.TerraformVersion = &v
	return s
}

func (s *CreateTaskRequest) SetTriggerStrategy(v string) *CreateTaskRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateTaskRequest) Validate() error {
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
	if s.TaskBackend != nil {
		if err := s.TaskBackend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskRequestGroupInfo struct {
	// The group ID.
	//
	// example:
	//
	// g-5fd38c9b83a86432e2
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-433aeade5d9167608
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s CreateTaskRequestGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateTaskRequestGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateTaskRequestGroupInfo) SetGroupId(v string) *CreateTaskRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *CreateTaskRequestGroupInfo) SetProjectId(v string) *CreateTaskRequestGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *CreateTaskRequestGroupInfo) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestTags struct {
	// The tag key of the node.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the node.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s CreateTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateTaskRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateTaskRequestTags) SetTagKey(v string) *CreateTaskRequestTags {
	s.TagKey = &v
	return s
}

func (s *CreateTaskRequestTags) SetTagValue(v string) *CreateTaskRequestTags {
	s.TagValue = &v
	return s
}

func (s *CreateTaskRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestTaskBackend struct {
	// The endpoint information.
	//
	// example:
	//
	// ss-cn-beijing.aliyuncs.com
	BucketEndpoint *string `json:"bucketEndpoint,omitempty" xml:"bucketEndpoint,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// iac-runtime-test
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The object path.
	//
	// example:
	//
	// /log
	ObjectPath *string `json:"objectPath,omitempty" xml:"objectPath,omitempty"`
}

func (s CreateTaskRequestTaskBackend) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestTaskBackend) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestTaskBackend) GetBucketEndpoint() *string {
	return s.BucketEndpoint
}

func (s *CreateTaskRequestTaskBackend) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateTaskRequestTaskBackend) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *CreateTaskRequestTaskBackend) SetBucketEndpoint(v string) *CreateTaskRequestTaskBackend {
	s.BucketEndpoint = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) SetBucketName(v string) *CreateTaskRequestTaskBackend {
	s.BucketName = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) SetObjectPath(v string) *CreateTaskRequestTaskBackend {
	s.ObjectPath = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) Validate() error {
	return dara.Validate(s)
}
