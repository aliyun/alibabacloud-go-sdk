// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDestroy(v bool) *CreateGroupRequest
	GetAutoDestroy() *bool
	SetAutoTrigger(v bool) *CreateGroupRequest
	GetAutoTrigger() *bool
	SetClientToken(v string) *CreateGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateGroupRequest
	GetDescription() *string
	SetForcedSetting(v bool) *CreateGroupRequest
	GetForcedSetting() *bool
	SetName(v string) *CreateGroupRequest
	GetName() *string
	SetNotifyConfig(v []*CreateGroupRequestNotifyConfig) *CreateGroupRequest
	GetNotifyConfig() []*CreateGroupRequestNotifyConfig
	SetNotifyOperationTypes(v []*string) *CreateGroupRequest
	GetNotifyOperationTypes() []*string
	SetProjectId(v string) *CreateGroupRequest
	GetProjectId() *string
	SetRamRole(v string) *CreateGroupRequest
	GetRamRole() *string
	SetReportExportField(v []*string) *CreateGroupRequest
	GetReportExportField() []*string
	SetReportExportPath(v string) *CreateGroupRequest
	GetReportExportPath() *string
	SetTerraformProviderVersion(v string) *CreateGroupRequest
	GetTerraformProviderVersion() *string
	SetTriggerConfig(v []*CreateGroupRequestTriggerConfig) *CreateGroupRequest
	GetTriggerConfig() []*CreateGroupRequestTriggerConfig
	SetTriggerResourceType(v []*string) *CreateGroupRequest
	GetTriggerResourceType() []*string
}

type CreateGroupRequest struct {
	// Specifies whether to delete the group after creation.
	//
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// Specifies whether to enable the automatic trigger policy. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	AutoTrigger *bool `json:"autoTrigger,omitempty" xml:"autoTrigger,omitempty"`
	// The idempotence token. Format: [0-9a-zA-Z-]{1,64}. Use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the group.
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
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The notification configuration.
	NotifyConfig []*CreateGroupRequestNotifyConfig `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty" type:"Repeated"`
	// The list of notification operation types.
	NotifyOperationTypes []*string `json:"notifyOperationTypes,omitempty" xml:"notifyOperationTypes,omitempty" type:"Repeated"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-433aead7560571a87349d054b4
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The RAM role (1-128 characters).
	//
	// The system assumes this role to execute the template when a new job is triggered.
	//
	// This parameter is required when the job trigger method is not manual.
	//
	// example:
	//
	// ramName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The list of export fields for the report.
	ReportExportField []*string `json:"reportExportField,omitempty" xml:"reportExportField,omitempty" type:"Repeated"`
	// The export address for the execution report. OSS addresses are supported.
	//
	// https://<OSS bucket address>/<path>.
	//
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/test/test
	ReportExportPath *string `json:"reportExportPath,omitempty" xml:"reportExportPath,omitempty"`
	// The Terraform Provider version.
	//
	// Select a Terraform Provider version. The version configured on the task takes higher priority.
	//
	// example:
	//
	// 1.189.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The trigger policy.
	//
	// This parameter cannot be empty when autoTrigger is set to true.
	TriggerConfig []*CreateGroupRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Repeated"`
	// The resource type for triggered execution. Valid values:
	//
	// - Task: regular task.
	//
	// - SceneTestingTask: scenario-based testing task.
	TriggerResourceType []*string `json:"triggerResourceType,omitempty" xml:"triggerResourceType,omitempty" type:"Repeated"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *CreateGroupRequest) GetAutoTrigger() *bool {
	return s.AutoTrigger
}

func (s *CreateGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupRequest) GetForcedSetting() *bool {
	return s.ForcedSetting
}

func (s *CreateGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupRequest) GetNotifyConfig() []*CreateGroupRequestNotifyConfig {
	return s.NotifyConfig
}

func (s *CreateGroupRequest) GetNotifyOperationTypes() []*string {
	return s.NotifyOperationTypes
}

func (s *CreateGroupRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateGroupRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateGroupRequest) GetReportExportField() []*string {
	return s.ReportExportField
}

func (s *CreateGroupRequest) GetReportExportPath() *string {
	return s.ReportExportPath
}

func (s *CreateGroupRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *CreateGroupRequest) GetTriggerConfig() []*CreateGroupRequestTriggerConfig {
	return s.TriggerConfig
}

func (s *CreateGroupRequest) GetTriggerResourceType() []*string {
	return s.TriggerResourceType
}

func (s *CreateGroupRequest) SetAutoDestroy(v bool) *CreateGroupRequest {
	s.AutoDestroy = &v
	return s
}

func (s *CreateGroupRequest) SetAutoTrigger(v bool) *CreateGroupRequest {
	s.AutoTrigger = &v
	return s
}

func (s *CreateGroupRequest) SetClientToken(v string) *CreateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetForcedSetting(v bool) *CreateGroupRequest {
	s.ForcedSetting = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetNotifyConfig(v []*CreateGroupRequestNotifyConfig) *CreateGroupRequest {
	s.NotifyConfig = v
	return s
}

func (s *CreateGroupRequest) SetNotifyOperationTypes(v []*string) *CreateGroupRequest {
	s.NotifyOperationTypes = v
	return s
}

func (s *CreateGroupRequest) SetProjectId(v string) *CreateGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateGroupRequest) SetRamRole(v string) *CreateGroupRequest {
	s.RamRole = &v
	return s
}

func (s *CreateGroupRequest) SetReportExportField(v []*string) *CreateGroupRequest {
	s.ReportExportField = v
	return s
}

func (s *CreateGroupRequest) SetReportExportPath(v string) *CreateGroupRequest {
	s.ReportExportPath = &v
	return s
}

func (s *CreateGroupRequest) SetTerraformProviderVersion(v string) *CreateGroupRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateGroupRequest) SetTriggerConfig(v []*CreateGroupRequestTriggerConfig) *CreateGroupRequest {
	s.TriggerConfig = v
	return s
}

func (s *CreateGroupRequest) SetTriggerResourceType(v []*string) *CreateGroupRequest {
	s.TriggerResourceType = v
	return s
}

func (s *CreateGroupRequest) Validate() error {
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

type CreateGroupRequestNotifyConfig struct {
	// The path configuration for notifications.
	//
	// example:
	//
	// /
	NotifyPath *string `json:"notifyPath,omitempty" xml:"notifyPath,omitempty"`
	// The notification type:
	//
	// DingDing.
	//
	// example:
	//
	// DingDing
	NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty"`
}

func (s CreateGroupRequestNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequestNotifyConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestNotifyConfig) GetNotifyPath() *string {
	return s.NotifyPath
}

func (s *CreateGroupRequestNotifyConfig) GetNotifyType() *string {
	return s.NotifyType
}

func (s *CreateGroupRequestNotifyConfig) SetNotifyPath(v string) *CreateGroupRequestNotifyConfig {
	s.NotifyPath = &v
	return s
}

func (s *CreateGroupRequestNotifyConfig) SetNotifyType(v string) *CreateGroupRequestNotifyConfig {
	s.NotifyType = &v
	return s
}

func (s *CreateGroupRequestNotifyConfig) Validate() error {
	return dara.Validate(s)
}

type CreateGroupRequestTriggerConfig struct {
	// The trigger policy. Valid values:
	//
	// - ProviderNewVersion: triggered when a new Provider version is released.
	//
	// - Cron: triggered on a schedule.
	//
	// example:
	//
	// Cron
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The policy value to maintain for scheduled triggering. This is a cron expression.
	//
	// example:
	//
	// 0 0 19 	- 	- ？
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s CreateGroupRequestTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestTriggerConfig) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *CreateGroupRequestTriggerConfig) GetTriggerValue() *string {
	return s.TriggerValue
}

func (s *CreateGroupRequestTriggerConfig) SetTriggerStrategy(v string) *CreateGroupRequestTriggerConfig {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateGroupRequestTriggerConfig) SetTriggerValue(v string) *CreateGroupRequestTriggerConfig {
	s.TriggerValue = &v
	return s
}

func (s *CreateGroupRequestTriggerConfig) Validate() error {
	return dara.Validate(s)
}
