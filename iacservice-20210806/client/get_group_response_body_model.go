// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody
	GetGroup() *GetGroupResponseBodyGroup
	SetRequestId(v string) *GetGroupResponseBody
	GetRequestId() *string
}

type GetGroupResponseBody struct {
	// The group.
	Group *GetGroupResponseBodyGroup `json:"group,omitempty" xml:"group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) GetGroup() *GetGroupResponseBodyGroup {
	return s.Group
}

func (s *GetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGroupResponseBodyGroup struct {
	// Indicates whether automatic deletion is enabled.
	//
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// Indicates whether automatic triggering is enabled.
	//
	// example:
	//
	// true
	AutoTrigger *bool `json:"autoTrigger,omitempty" xml:"autoTrigger,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-08-21T10:57:11Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The group description.
	//
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the group configuration is forcibly used.
	//
	// example:
	//
	// true
	ForcedSetting *bool `json:"forcedSetting,omitempty" xml:"forcedSetting,omitempty"`
	// The group ID.
	//
	// example:
	//
	// g-14e80de4866bf7ffed0bab6154d738
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The notification configuration.
	NotifyConfig []*GetGroupResponseBodyGroupNotifyConfig `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty" type:"Repeated"`
	// The list of notification operation types.
	NotifyOperationTypes []*string `json:"notifyOperationTypes,omitempty" xml:"notifyOperationTypes,omitempty" type:"Repeated"`
	// The project ID.
	//
	// example:
	//
	// p-4267dcfbf1b6d126edcadf0e949
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The RAM role (1 to 128 characters). The system assumes this role to execute the template when a new job is triggered. This parameter is required when the job trigger mode is not manual.
	//
	// example:
	//
	// ramRoleName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The list of report export field options.
	ReportExportField []*string `json:"reportExportField,omitempty" xml:"reportExportField,omitempty" type:"Repeated"`
	// The export address for the execution report. OSS addresses are supported. Format: https://<OSS bucket address>/<path>.
	//
	// example:
	//
	// /
	ReportExportPath *string `json:"reportExportPath,omitempty" xml:"reportExportPath,omitempty"`
	// The number of tasks.
	//
	// example:
	//
	// 3
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
	// The Terraform provider version. Select a Terraform provider version. Tasks in the group are executed based on the specified Terraform provider version. The version configured on a task takes higher priority. This version may conflict with the Terraform provider version specified in the module.
	//
	// example:
	//
	// 1.191.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The trigger policy. This parameter cannot be empty when autoTrigger is set to true.
	TriggerConfig []*GetGroupResponseBodyGroupTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Repeated"`
	// The resource type that triggers execution. Valid values:
	//
	// - Task: regular task
	//
	// - SceneTestingTask: scenario-based testing task.
	TriggerResourceType []*string `json:"triggerResourceType,omitempty" xml:"triggerResourceType,omitempty" type:"Repeated"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *GetGroupResponseBodyGroup) GetAutoTrigger() *bool {
	return s.AutoTrigger
}

func (s *GetGroupResponseBodyGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetGroupResponseBodyGroup) GetDescription() *string {
	return s.Description
}

func (s *GetGroupResponseBodyGroup) GetForcedSetting() *bool {
	return s.ForcedSetting
}

func (s *GetGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupResponseBodyGroup) GetName() *string {
	return s.Name
}

func (s *GetGroupResponseBodyGroup) GetNotifyConfig() []*GetGroupResponseBodyGroupNotifyConfig {
	return s.NotifyConfig
}

func (s *GetGroupResponseBodyGroup) GetNotifyOperationTypes() []*string {
	return s.NotifyOperationTypes
}

func (s *GetGroupResponseBodyGroup) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetGroupResponseBodyGroup) GetRamRole() *string {
	return s.RamRole
}

func (s *GetGroupResponseBodyGroup) GetReportExportField() []*string {
	return s.ReportExportField
}

func (s *GetGroupResponseBodyGroup) GetReportExportPath() *string {
	return s.ReportExportPath
}

func (s *GetGroupResponseBodyGroup) GetTaskCnt() *int64 {
	return s.TaskCnt
}

func (s *GetGroupResponseBodyGroup) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *GetGroupResponseBodyGroup) GetTriggerConfig() []*GetGroupResponseBodyGroupTriggerConfig {
	return s.TriggerConfig
}

func (s *GetGroupResponseBodyGroup) GetTriggerResourceType() []*string {
	return s.TriggerResourceType
}

func (s *GetGroupResponseBodyGroup) SetAutoDestroy(v bool) *GetGroupResponseBodyGroup {
	s.AutoDestroy = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetAutoTrigger(v bool) *GetGroupResponseBodyGroup {
	s.AutoTrigger = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetCreateTime(v string) *GetGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetDescription(v string) *GetGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetForcedSetting(v bool) *GetGroupResponseBodyGroup {
	s.ForcedSetting = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetName(v string) *GetGroupResponseBodyGroup {
	s.Name = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetNotifyConfig(v []*GetGroupResponseBodyGroupNotifyConfig) *GetGroupResponseBodyGroup {
	s.NotifyConfig = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetNotifyOperationTypes(v []*string) *GetGroupResponseBodyGroup {
	s.NotifyOperationTypes = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetProjectId(v string) *GetGroupResponseBodyGroup {
	s.ProjectId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetRamRole(v string) *GetGroupResponseBodyGroup {
	s.RamRole = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetReportExportField(v []*string) *GetGroupResponseBodyGroup {
	s.ReportExportField = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetReportExportPath(v string) *GetGroupResponseBodyGroup {
	s.ReportExportPath = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTaskCnt(v int64) *GetGroupResponseBodyGroup {
	s.TaskCnt = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTerraformProviderVersion(v string) *GetGroupResponseBodyGroup {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTriggerConfig(v []*GetGroupResponseBodyGroupTriggerConfig) *GetGroupResponseBodyGroup {
	s.TriggerConfig = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTriggerResourceType(v []*string) *GetGroupResponseBodyGroup {
	s.TriggerResourceType = v
	return s
}

func (s *GetGroupResponseBodyGroup) Validate() error {
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

type GetGroupResponseBodyGroupNotifyConfig struct {
	// The path configuration for notifications.
	//
	// example:
	//
	// /
	NotifyPath *string `json:"notifyPath,omitempty" xml:"notifyPath,omitempty"`
	// The notification type. Valid values:
	//
	// DingDing.
	//
	// example:
	//
	// DingDing
	NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty"`
}

func (s GetGroupResponseBodyGroupNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBodyGroupNotifyConfig) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroupNotifyConfig) GetNotifyPath() *string {
	return s.NotifyPath
}

func (s *GetGroupResponseBodyGroupNotifyConfig) GetNotifyType() *string {
	return s.NotifyType
}

func (s *GetGroupResponseBodyGroupNotifyConfig) SetNotifyPath(v string) *GetGroupResponseBodyGroupNotifyConfig {
	s.NotifyPath = &v
	return s
}

func (s *GetGroupResponseBodyGroupNotifyConfig) SetNotifyType(v string) *GetGroupResponseBodyGroupNotifyConfig {
	s.NotifyType = &v
	return s
}

func (s *GetGroupResponseBodyGroupNotifyConfig) Validate() error {
	return dara.Validate(s)
}

type GetGroupResponseBodyGroupTriggerConfig struct {
	// The trigger strategy. Valid values:
	//
	// - ProviderNewVersion: triggered when a new provider version is released
	//
	// - Cron: triggered on a schedule.
	//
	// example:
	//
	// Cron
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The policy value that must be maintained for scheduled triggering. This value is a cron expression.
	//
	// example:
	//
	// 0 0 8 	- 	- ?
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s GetGroupResponseBodyGroupTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBodyGroupTriggerConfig) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroupTriggerConfig) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *GetGroupResponseBodyGroupTriggerConfig) GetTriggerValue() *string {
	return s.TriggerValue
}

func (s *GetGroupResponseBodyGroupTriggerConfig) SetTriggerStrategy(v string) *GetGroupResponseBodyGroupTriggerConfig {
	s.TriggerStrategy = &v
	return s
}

func (s *GetGroupResponseBodyGroupTriggerConfig) SetTriggerValue(v string) *GetGroupResponseBodyGroupTriggerConfig {
	s.TriggerValue = &v
	return s
}

func (s *GetGroupResponseBodyGroupTriggerConfig) Validate() error {
	return dara.Validate(s)
}
