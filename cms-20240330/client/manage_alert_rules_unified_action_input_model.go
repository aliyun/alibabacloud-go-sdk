// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesUnifiedActionInput interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ManageAlertRulesUnifiedActionInput
	GetAction() *string
	SetActionIntegrationConfig(v *ActionIntegrationConfig) *ManageAlertRulesUnifiedActionInput
	GetActionIntegrationConfig() *ActionIntegrationConfig
	SetAnnotations(v map[string]*string) *ManageAlertRulesUnifiedActionInput
	GetAnnotations() map[string]*string
	SetArmsIntegrationConfig(v *ArmsIntegrationConfig) *ManageAlertRulesUnifiedActionInput
	GetArmsIntegrationConfig() *ArmsIntegrationConfig
	SetBizSource(v string) *ManageAlertRulesUnifiedActionInput
	GetBizSource() *string
	SetConditionConfig(v *ConditionConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetConditionConfig() *ConditionConfigUnified
	SetContentTemplate(v string) *ManageAlertRulesUnifiedActionInput
	GetContentTemplate() *string
	SetDatasourceConfig(v *DatasourceConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetDatasourceConfig() *DatasourceConfigUnified
	SetDisplayName(v string) *ManageAlertRulesUnifiedActionInput
	GetDisplayName() *string
	SetDrillEvalTime(v int64) *ManageAlertRulesUnifiedActionInput
	GetDrillEvalTime() *int64
	SetDrillMode(v string) *ManageAlertRulesUnifiedActionInput
	GetDrillMode() *string
	SetDrillTimeRange(v *ManageAlertRulesUnifiedActionInputDrillTimeRange) *ManageAlertRulesUnifiedActionInput
	GetDrillTimeRange() *ManageAlertRulesUnifiedActionInputDrillTimeRange
	SetEnabled(v bool) *ManageAlertRulesUnifiedActionInput
	GetEnabled() *bool
	SetLabels(v map[string]*string) *ManageAlertRulesUnifiedActionInput
	GetLabels() map[string]*string
	SetNotifyConfig(v *NotifyConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetNotifyConfig() *NotifyConfigUnified
	SetObserveResourceConfig(v *ObserveResourceConfig) *ManageAlertRulesUnifiedActionInput
	GetObserveResourceConfig() *ObserveResourceConfig
	SetObserveResourceInstanceId(v string) *ManageAlertRulesUnifiedActionInput
	GetObserveResourceInstanceId() *string
	SetObserveResourceType(v string) *ManageAlertRulesUnifiedActionInput
	GetObserveResourceType() *string
	SetQueryConfig(v *QueryConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetQueryConfig() *QueryConfigUnified
	SetRcaConfig(v *AlertRuleRcaConfig) *ManageAlertRulesUnifiedActionInput
	GetRcaConfig() *AlertRuleRcaConfig
	SetRegionId(v string) *ManageAlertRulesUnifiedActionInput
	GetRegionId() *string
	SetScheduleConfig(v *ScheduleConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetScheduleConfig() *ScheduleConfigUnified
	SetUuid(v string) *ManageAlertRulesUnifiedActionInput
	GetUuid() *string
	SetUuidList(v []*string) *ManageAlertRulesUnifiedActionInput
	GetUuidList() []*string
	SetWorkspace(v string) *ManageAlertRulesUnifiedActionInput
	GetWorkspace() *string
}

type ManageAlertRulesUnifiedActionInput struct {
	// The operation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The action integration configuration.
	//
	// example:
	//
	// {"enabled":true,"actions":[]}
	ActionIntegrationConfig *ActionIntegrationConfig `json:"actionIntegrationConfig,omitempty" xml:"actionIntegrationConfig,omitempty"`
	// The annotations.
	//
	// example:
	//
	// {"summary":"CPU usage exceeds 90%"}
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// The ARMS integration configuration.
	//
	// example:
	//
	// {"enabled":false}
	ArmsIntegrationConfig *ArmsIntegrationConfig `json:"armsIntegrationConfig,omitempty" xml:"armsIntegrationConfig,omitempty"`
	// The business source (optional). Examples include managed_service_for_prometheus, umodel, application_insights, cloud_monitoring, and sls. Pass as needed for CREATE, UPDATE, and PATCH.
	//
	// example:
	//
	// Sample value
	BizSource *string `json:"bizSource,omitempty" xml:"bizSource,omitempty"`
	// The condition aggregation configuration.
	//
	// example:
	//
	// {"type":"PROMETHEUS_SIMPLE_CONDITION","durationSecs":0,"severity":"CRITICAL"}
	ConditionConfig *ConditionConfigUnified `json:"conditionConfig,omitempty" xml:"conditionConfig,omitempty"`
	// The content template.
	//
	// example:
	//
	// Instance {{instance}} CPU usage exceeds {{threshold}}%
	ContentTemplate *string `json:"contentTemplate,omitempty" xml:"contentTemplate,omitempty"`
	// The data source aggregation configuration.
	//
	// example:
	//
	// {"type":"PROMETHEUS","instanceId":"prom-cn-hangzhou-xxxxx"}
	DatasourceConfig *DatasourceConfigUnified `json:"datasourceConfig,omitempty" xml:"datasourceConfig,omitempty"`
	// The display name.
	//
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The alert drill time, in Unix milliseconds. Required by the backend when drillMode is ONE_SHOT. It cannot be later than the current time or earlier than 90 days ago.
	//
	// example:
	//
	// 1783728000000
	DrillEvalTime *int64 `json:"drillEvalTime,omitempty" xml:"drillEvalTime,omitempty"`
	// The alert drill mode. Required by the backend when action is DRILL_ALERT. ONE_SHOT specifies a single time point to simulate a complete lifecycle (three events: OCCURRED, PERSISTENT, and RECOVERED). TIME_TRAVEL specifies a time period to simulate round by round according to the rule scheduling interval (OCCURRED, PERSISTENT, ..., RECOVERED).
	//
	// example:
	//
	// ONE_SHOT
	DrillMode *string `json:"drillMode,omitempty" xml:"drillMode,omitempty"`
	// The alert drill time range, in Unix milliseconds. Required by the backend when drillMode is TIME_TRAVEL. The endTime must be later than the startTime and cannot be later than the current time. The span between endTime and startTime cannot exceed 1 hour (the backend returns 400 if exceeded). Simulates round by round according to the rule scheduling interval, with a maximum of 500 rounds (the backend returns 400 if exceeded).
	//
	// example:
	//
	// {"startTime":1783728000000,"endTime":1783731600000}
	DrillTimeRange *ManageAlertRulesUnifiedActionInputDrillTimeRange `json:"drillTimeRange,omitempty" xml:"drillTimeRange,omitempty" type:"Struct"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The labels.
	//
	// example:
	//
	// {"env":"production"}
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The notification aggregation configuration.
	//
	// example:
	//
	// {"type":"DIRECT_NOTIFY","channels":[{"type":"GROUP","identifiers":["<contact-group-id>"]}]}
	NotifyConfig *NotifyConfigUnified `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty"`
	// The observable resource configuration.
	//
	// example:
	//
	// {"relationType":"CLOUD_INSTANCE","resources":[]}
	ObserveResourceConfig *ObserveResourceConfig `json:"observeResourceConfig,omitempty" xml:"observeResourceConfig,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The monitored object instance ID. Use observeResourceConfig.resources for new integrations. Retained only for compatibility with legacy SDKs.
	//
	// example:
	//
	// example-id-001
	ObserveResourceInstanceId *string `json:"observeResourceInstanceId,omitempty" xml:"observeResourceInstanceId,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The monitored object resource type. Use observeResourceConfig.entityType for new integrations. Retained only for compatibility with legacy SDKs.
	//
	// example:
	//
	// default
	ObserveResourceType *string `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	// The query aggregation configuration.
	//
	// example:
	//
	// {"type":"PROMETHEUS_SINGLE_QUERY","expr":"up > 0"}
	QueryConfig *QueryConfigUnified `json:"queryConfig,omitempty" xml:"queryConfig,omitempty"`
	// The root cause analysis (RCA) configuration (optional). Pass as needed for CREATE, UPDATE, and PATCH.
	//
	// example:
	//
	// {"enabled":false}
	RcaConfig *AlertRuleRcaConfig `json:"rcaConfig,omitempty" xml:"rcaConfig,omitempty"`
	// The region ID. Aligned with V1 AlertRule.regionId. The gateway callerRegionId is used if not provided.
	//
	// example:
	//
	// example-id-001
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The scheduling aggregation configuration.
	//
	// example:
	//
	// {"type":"FIXED","intervalSecs":60}
	ScheduleConfig *ScheduleConfigUnified `json:"scheduleConfig,omitempty" xml:"scheduleConfig,omitempty"`
	// The rule UUID. Required for UPDATE and PATCH.
	//
	// example:
	//
	// xxx-xxx-xxx
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The list of rule UUIDs to be deleted (BATCH_DELETE).
	//
	// example:
	//
	// ["cms_v2_81d01768-fee9-459f-af49-ab2704f1db2b"]
	UuidList []*string `json:"uuidList,omitempty" xml:"uuidList,omitempty" type:"Repeated"`
	// The workspace (for CREATE, UPDATE, and other actions).
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ManageAlertRulesUnifiedActionInput) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesUnifiedActionInput) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesUnifiedActionInput) GetAction() *string {
	return s.Action
}

func (s *ManageAlertRulesUnifiedActionInput) GetActionIntegrationConfig() *ActionIntegrationConfig {
	return s.ActionIntegrationConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *ManageAlertRulesUnifiedActionInput) GetArmsIntegrationConfig() *ArmsIntegrationConfig {
	return s.ArmsIntegrationConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetBizSource() *string {
	return s.BizSource
}

func (s *ManageAlertRulesUnifiedActionInput) GetConditionConfig() *ConditionConfigUnified {
	return s.ConditionConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetContentTemplate() *string {
	return s.ContentTemplate
}

func (s *ManageAlertRulesUnifiedActionInput) GetDatasourceConfig() *DatasourceConfigUnified {
	return s.DatasourceConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ManageAlertRulesUnifiedActionInput) GetDrillEvalTime() *int64 {
	return s.DrillEvalTime
}

func (s *ManageAlertRulesUnifiedActionInput) GetDrillMode() *string {
	return s.DrillMode
}

func (s *ManageAlertRulesUnifiedActionInput) GetDrillTimeRange() *ManageAlertRulesUnifiedActionInputDrillTimeRange {
	return s.DrillTimeRange
}

func (s *ManageAlertRulesUnifiedActionInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *ManageAlertRulesUnifiedActionInput) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ManageAlertRulesUnifiedActionInput) GetNotifyConfig() *NotifyConfigUnified {
	return s.NotifyConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetObserveResourceConfig() *ObserveResourceConfig {
	return s.ObserveResourceConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetObserveResourceInstanceId() *string {
	return s.ObserveResourceInstanceId
}

func (s *ManageAlertRulesUnifiedActionInput) GetObserveResourceType() *string {
	return s.ObserveResourceType
}

func (s *ManageAlertRulesUnifiedActionInput) GetQueryConfig() *QueryConfigUnified {
	return s.QueryConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetRcaConfig() *AlertRuleRcaConfig {
	return s.RcaConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetRegionId() *string {
	return s.RegionId
}

func (s *ManageAlertRulesUnifiedActionInput) GetScheduleConfig() *ScheduleConfigUnified {
	return s.ScheduleConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetUuid() *string {
	return s.Uuid
}

func (s *ManageAlertRulesUnifiedActionInput) GetUuidList() []*string {
	return s.UuidList
}

func (s *ManageAlertRulesUnifiedActionInput) GetWorkspace() *string {
	return s.Workspace
}

func (s *ManageAlertRulesUnifiedActionInput) SetAction(v string) *ManageAlertRulesUnifiedActionInput {
	s.Action = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetActionIntegrationConfig(v *ActionIntegrationConfig) *ManageAlertRulesUnifiedActionInput {
	s.ActionIntegrationConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetAnnotations(v map[string]*string) *ManageAlertRulesUnifiedActionInput {
	s.Annotations = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetArmsIntegrationConfig(v *ArmsIntegrationConfig) *ManageAlertRulesUnifiedActionInput {
	s.ArmsIntegrationConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetBizSource(v string) *ManageAlertRulesUnifiedActionInput {
	s.BizSource = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetConditionConfig(v *ConditionConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.ConditionConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetContentTemplate(v string) *ManageAlertRulesUnifiedActionInput {
	s.ContentTemplate = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDatasourceConfig(v *DatasourceConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.DatasourceConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDisplayName(v string) *ManageAlertRulesUnifiedActionInput {
	s.DisplayName = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDrillEvalTime(v int64) *ManageAlertRulesUnifiedActionInput {
	s.DrillEvalTime = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDrillMode(v string) *ManageAlertRulesUnifiedActionInput {
	s.DrillMode = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDrillTimeRange(v *ManageAlertRulesUnifiedActionInputDrillTimeRange) *ManageAlertRulesUnifiedActionInput {
	s.DrillTimeRange = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetEnabled(v bool) *ManageAlertRulesUnifiedActionInput {
	s.Enabled = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetLabels(v map[string]*string) *ManageAlertRulesUnifiedActionInput {
	s.Labels = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetNotifyConfig(v *NotifyConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.NotifyConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetObserveResourceConfig(v *ObserveResourceConfig) *ManageAlertRulesUnifiedActionInput {
	s.ObserveResourceConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetObserveResourceInstanceId(v string) *ManageAlertRulesUnifiedActionInput {
	s.ObserveResourceInstanceId = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetObserveResourceType(v string) *ManageAlertRulesUnifiedActionInput {
	s.ObserveResourceType = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetQueryConfig(v *QueryConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.QueryConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetRcaConfig(v *AlertRuleRcaConfig) *ManageAlertRulesUnifiedActionInput {
	s.RcaConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetRegionId(v string) *ManageAlertRulesUnifiedActionInput {
	s.RegionId = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetScheduleConfig(v *ScheduleConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.ScheduleConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetUuid(v string) *ManageAlertRulesUnifiedActionInput {
	s.Uuid = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetUuidList(v []*string) *ManageAlertRulesUnifiedActionInput {
	s.UuidList = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetWorkspace(v string) *ManageAlertRulesUnifiedActionInput {
	s.Workspace = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) Validate() error {
	if s.ActionIntegrationConfig != nil {
		if err := s.ActionIntegrationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ArmsIntegrationConfig != nil {
		if err := s.ArmsIntegrationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConditionConfig != nil {
		if err := s.ConditionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DatasourceConfig != nil {
		if err := s.DatasourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DrillTimeRange != nil {
		if err := s.DrillTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyConfig != nil {
		if err := s.NotifyConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ObserveResourceConfig != nil {
		if err := s.ObserveResourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.QueryConfig != nil {
		if err := s.QueryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RcaConfig != nil {
		if err := s.RcaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ManageAlertRulesUnifiedActionInputDrillTimeRange struct {
	// The drill end time, in Unix milliseconds. Must be later than the start time, and the span from the start time cannot exceed 1 hour.
	//
	// example:
	//
	// 1783731600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The drill start time, in Unix milliseconds.
	//
	// example:
	//
	// 1783728000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ManageAlertRulesUnifiedActionInputDrillTimeRange) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesUnifiedActionInputDrillTimeRange) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesUnifiedActionInputDrillTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ManageAlertRulesUnifiedActionInputDrillTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ManageAlertRulesUnifiedActionInputDrillTimeRange) SetEndTime(v int64) *ManageAlertRulesUnifiedActionInputDrillTimeRange {
	s.EndTime = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInputDrillTimeRange) SetStartTime(v int64) *ManageAlertRulesUnifiedActionInputDrillTimeRange {
	s.StartTime = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInputDrillTimeRange) Validate() error {
	return dara.Validate(s)
}
