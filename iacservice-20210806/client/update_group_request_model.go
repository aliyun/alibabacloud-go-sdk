// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDestroy(v bool) *UpdateGroupRequest
	GetAutoDestroy() *bool
	SetAutoTrigger(v bool) *UpdateGroupRequest
	GetAutoTrigger() *bool
	SetClientToken(v string) *UpdateGroupRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateGroupRequest
	GetDescription() *string
	SetForcedSetting(v bool) *UpdateGroupRequest
	GetForcedSetting() *bool
	SetName(v string) *UpdateGroupRequest
	GetName() *string
	SetNotifyConfig(v []*UpdateGroupRequestNotifyConfig) *UpdateGroupRequest
	GetNotifyConfig() []*UpdateGroupRequestNotifyConfig
	SetNotifyOperationTypes(v []*string) *UpdateGroupRequest
	GetNotifyOperationTypes() []*string
	SetRamRole(v string) *UpdateGroupRequest
	GetRamRole() *string
	SetReportExportField(v []*string) *UpdateGroupRequest
	GetReportExportField() []*string
	SetReportExportPath(v string) *UpdateGroupRequest
	GetReportExportPath() *string
	SetTerraformProviderVersion(v string) *UpdateGroupRequest
	GetTerraformProviderVersion() *string
	SetTriggerConfig(v []*UpdateGroupRequestTriggerConfig) *UpdateGroupRequest
	GetTriggerConfig() []*UpdateGroupRequestTriggerConfig
	SetTriggerResourceType(v []*string) *UpdateGroupRequest
	GetTriggerResourceType() []*string
}

type UpdateGroupRequest struct {
	// Specifies whether to automatically delete the group.
	//
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// Specifies whether to enable the automatic trigger policy. Valid values: - **true**: enabled. - **false**: disabled.
	//
	// example:
	//
	// true
	AutoTrigger *bool `json:"autoTrigger,omitempty" xml:"autoTrigger,omitempty"`
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. We recommend that you use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The group description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to forcibly use the group configuration.
	//
	// example:
	//
	// true
	ForcedSetting *bool `json:"forcedSetting,omitempty" xml:"forcedSetting,omitempty"`
	// The group name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The notification configuration.
	NotifyConfig []*UpdateGroupRequestNotifyConfig `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty" type:"Repeated"`
	// The list of notification operation types.
	NotifyOperationTypes []*string `json:"notifyOperationTypes,omitempty" xml:"notifyOperationTypes,omitempty" type:"Repeated"`
	// The RAM role (1-128 characters). The system assumes this role to execute the template when a new job is triggered. This parameter is required when the job trigger method is not manual.
	//
	// example:
	//
	// {}
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The list of export fields for the report.
	ReportExportField []*string `json:"reportExportField,omitempty" xml:"reportExportField,omitempty" type:"Repeated"`
	// The export path for the execution report. OSS paths are supported.
	//
	// example:
	//
	// /
	ReportExportPath *string `json:"reportExportPath,omitempty" xml:"reportExportPath,omitempty"`
	// The Terraform Provider version. Select a Terraform Provider version. Nodes in the group execute plans based on the specified Terraform Provider version. The version configured on a node takes higher priority.
	//
	// example:
	//
	// 1.183.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The trigger policy. This parameter cannot be empty when autoTrigger is set to true.
	TriggerConfig []*UpdateGroupRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Repeated"`
	// The resource type that triggers execution. Valid values:
	//
	// - Task: regular node.
	//
	// - SceneTestingTask: scenario-based testing node.
	TriggerResourceType []*string `json:"triggerResourceType,omitempty" xml:"triggerResourceType,omitempty" type:"Repeated"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *UpdateGroupRequest) GetAutoTrigger() *bool {
	return s.AutoTrigger
}

func (s *UpdateGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupRequest) GetForcedSetting() *bool {
	return s.ForcedSetting
}

func (s *UpdateGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGroupRequest) GetNotifyConfig() []*UpdateGroupRequestNotifyConfig {
	return s.NotifyConfig
}

func (s *UpdateGroupRequest) GetNotifyOperationTypes() []*string {
	return s.NotifyOperationTypes
}

func (s *UpdateGroupRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateGroupRequest) GetReportExportField() []*string {
	return s.ReportExportField
}

func (s *UpdateGroupRequest) GetReportExportPath() *string {
	return s.ReportExportPath
}

func (s *UpdateGroupRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *UpdateGroupRequest) GetTriggerConfig() []*UpdateGroupRequestTriggerConfig {
	return s.TriggerConfig
}

func (s *UpdateGroupRequest) GetTriggerResourceType() []*string {
	return s.TriggerResourceType
}

func (s *UpdateGroupRequest) SetAutoDestroy(v bool) *UpdateGroupRequest {
	s.AutoDestroy = &v
	return s
}

func (s *UpdateGroupRequest) SetAutoTrigger(v bool) *UpdateGroupRequest {
	s.AutoTrigger = &v
	return s
}

func (s *UpdateGroupRequest) SetClientToken(v string) *UpdateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupRequest) SetDescription(v string) *UpdateGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupRequest) SetForcedSetting(v bool) *UpdateGroupRequest {
	s.ForcedSetting = &v
	return s
}

func (s *UpdateGroupRequest) SetName(v string) *UpdateGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupRequest) SetNotifyConfig(v []*UpdateGroupRequestNotifyConfig) *UpdateGroupRequest {
	s.NotifyConfig = v
	return s
}

func (s *UpdateGroupRequest) SetNotifyOperationTypes(v []*string) *UpdateGroupRequest {
	s.NotifyOperationTypes = v
	return s
}

func (s *UpdateGroupRequest) SetRamRole(v string) *UpdateGroupRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateGroupRequest) SetReportExportField(v []*string) *UpdateGroupRequest {
	s.ReportExportField = v
	return s
}

func (s *UpdateGroupRequest) SetReportExportPath(v string) *UpdateGroupRequest {
	s.ReportExportPath = &v
	return s
}

func (s *UpdateGroupRequest) SetTerraformProviderVersion(v string) *UpdateGroupRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *UpdateGroupRequest) SetTriggerConfig(v []*UpdateGroupRequestTriggerConfig) *UpdateGroupRequest {
	s.TriggerConfig = v
	return s
}

func (s *UpdateGroupRequest) SetTriggerResourceType(v []*string) *UpdateGroupRequest {
	s.TriggerResourceType = v
	return s
}

func (s *UpdateGroupRequest) Validate() error {
	if s.NotifyConfig != nil {
		for _, item := range s.NotifyConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TriggerConfig != nil {
		for _, item := range s.TriggerConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGroupRequestNotifyConfig struct {
	// The path configuration for notifications.
	//
	// example:
	//
	// /
	NotifyPath *string `json:"notifyPath,omitempty" xml:"notifyPath,omitempty"`
	// The notification type. Valid values: DingDing.
	//
	// example:
	//
	// DingDing
	NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty"`
}

func (s UpdateGroupRequestNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequestNotifyConfig) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequestNotifyConfig) GetNotifyPath() *string {
	return s.NotifyPath
}

func (s *UpdateGroupRequestNotifyConfig) GetNotifyType() *string {
	return s.NotifyType
}

func (s *UpdateGroupRequestNotifyConfig) SetNotifyPath(v string) *UpdateGroupRequestNotifyConfig {
	s.NotifyPath = &v
	return s
}

func (s *UpdateGroupRequestNotifyConfig) SetNotifyType(v string) *UpdateGroupRequestNotifyConfig {
	s.NotifyType = &v
	return s
}

func (s *UpdateGroupRequestNotifyConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateGroupRequestTriggerConfig struct {
	// The trigger strategy. Valid values:
	//
	// - ProviderNewVersion: triggered when a new Provider version is released.
	//
	// - Cron: triggered on a schedule.
	//
	// example:
	//
	// Cron
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The policy value that must be maintained for scheduled triggers. This value is a cron expression.
	//
	// example:
	//
	// 0 0 	- 	- 	- ？
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s UpdateGroupRequestTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequestTriggerConfig) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *UpdateGroupRequestTriggerConfig) GetTriggerValue() *string {
	return s.TriggerValue
}

func (s *UpdateGroupRequestTriggerConfig) SetTriggerStrategy(v string) *UpdateGroupRequestTriggerConfig {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateGroupRequestTriggerConfig) SetTriggerValue(v string) *UpdateGroupRequestTriggerConfig {
	s.TriggerValue = &v
	return s
}

func (s *UpdateGroupRequestTriggerConfig) Validate() error {
	return dara.Validate(s)
}
