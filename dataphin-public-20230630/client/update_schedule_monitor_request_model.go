// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateScheduleMonitorRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateScheduleMonitorRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateScheduleMonitorRequestUpdateCommand) *UpdateScheduleMonitorRequest
	GetUpdateCommand() *UpdateScheduleMonitorRequestUpdateCommand
}

type UpdateScheduleMonitorRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The request body for updating scheduling monitoring.
	//
	// This parameter is required.
	UpdateCommand *UpdateScheduleMonitorRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateScheduleMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateScheduleMonitorRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateScheduleMonitorRequest) GetUpdateCommand() *UpdateScheduleMonitorRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateScheduleMonitorRequest) SetOpTenantId(v int64) *UpdateScheduleMonitorRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateScheduleMonitorRequest) SetOpUserId(v string) *UpdateScheduleMonitorRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateScheduleMonitorRequest) SetUpdateCommand(v *UpdateScheduleMonitorRequestUpdateCommand) *UpdateScheduleMonitorRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateScheduleMonitorRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommand struct {
	// The monitored object configuration.
	//
	// This parameter is required.
	MonitoredItemConfig *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig `json:"MonitoredItemConfig,omitempty" xml:"MonitoredItemConfig,omitempty" type:"Struct"`
	// The monitored object ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6912345678901234
	MonitoredItemId *int64 `json:"MonitoredItemId,omitempty" xml:"MonitoredItemId,omitempty"`
	// The list of monitored rule configurations. You can configure multiple rules with different alert causes.
	//
	// This parameter is required.
	MonitoredRuleConfigList []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList `json:"MonitoredRuleConfigList,omitempty" xml:"MonitoredRuleConfigList,omitempty" type:"Repeated"`
}

func (s UpdateScheduleMonitorRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) GetMonitoredItemConfig() *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	return s.MonitoredItemConfig
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) GetMonitoredItemId() *int64 {
	return s.MonitoredItemId
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) GetMonitoredRuleConfigList() []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList {
	return s.MonitoredRuleConfigList
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) SetMonitoredItemConfig(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) *UpdateScheduleMonitorRequestUpdateCommand {
	s.MonitoredItemConfig = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) SetMonitoredItemId(v int64) *UpdateScheduleMonitorRequestUpdateCommand {
	s.MonitoredItemId = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) SetMonitoredRuleConfigList(v []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) *UpdateScheduleMonitorRequestUpdateCommand {
	s.MonitoredRuleConfigList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommand) Validate() error {
	if s.MonitoredItemConfig != nil {
		if err := s.MonitoredItemConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MonitoredRuleConfigList != nil {
		for _, item := range s.MonitoredRuleConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig struct {
	// The business unit ID. Required when the summary logical table does not have a project ID.
	//
	// example:
	//
	// 100001
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The logical table ID. Required for the LOGICAL_FIELD and LOGICAL_TABLE_NODE types.
	//
	// example:
	//
	// 6812345678901234
	LogicalTableId *int64 `json:"LogicalTableId,omitempty" xml:"LogicalTableId,omitempty"`
	// The name of the logical table.
	//
	// example:
	//
	// dwd_order_detail
	LogicalTableName *string `json:"LogicalTableName,omitempty" xml:"LogicalTableName,omitempty"`
	// The monitored object type. Valid values: LOGICAL_FIELD (logical field), OND_ID_NODE (logical table node), SCHEDULE_PHYSICAL_NODE (periodically scheduled physical task), MANUAL_PHYSICAL_NODE (manually scheduled physical task), ODM_NODE (source path task), LOGICAL_TABLE_NODE (logical table task), and PROJECT (project).
	//
	// This parameter is required.
	//
	// example:
	//
	// SCHEDULE_PHYSICAL_NODE
	MonitoredItemType *string `json:"MonitoredItemType,omitempty" xml:"MonitoredItemType,omitempty"`
	// The monitored task scope, which specifies project-level monitoring filter rules. Required for the PROJECT type.
	MonitoredNodeRange *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange `json:"MonitoredNodeRange,omitempty" xml:"MonitoredNodeRange,omitempty" type:"Struct"`
	// The monitored task ID, which is the node ID. Required for types other than PROJECT.
	//
	// example:
	//
	// 12345678
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The monitored task name, which is the node name. Required for types other than PROJECT.
	//
	// example:
	//
	// dwd_order_detail_daily
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node subtype. Required for types other than PROJECT and LOGICAL_FIELD.
	//
	// example:
	//
	// MAX_COMPUTE_SQL
	NodeSubType *string `json:"NodeSubType,omitempty" xml:"NodeSubType,omitempty"`
	// The list of node O&M owner IDs. This serves as the source for OWNER recipients.
	OpsOwnerIdList []*string `json:"OpsOwnerIdList,omitempty" xml:"OpsOwnerIdList,omitempty" type:"Repeated"`
	// The project ID. Required for the PROJECT type. For other types, it specifies the project to which the monitored object belongs.
	//
	// example:
	//
	// 7086194564164288
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetLogicalTableId() *int64 {
	return s.LogicalTableId
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetLogicalTableName() *string {
	return s.LogicalTableName
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetMonitoredItemType() *string {
	return s.MonitoredItemType
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetMonitoredNodeRange() *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange {
	return s.MonitoredNodeRange
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetNodeName() *string {
	return s.NodeName
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetNodeSubType() *string {
	return s.NodeSubType
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetOpsOwnerIdList() []*string {
	return s.OpsOwnerIdList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetBizUnitId(v int64) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.BizUnitId = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetLogicalTableId(v int64) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.LogicalTableId = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetLogicalTableName(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.LogicalTableName = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetMonitoredItemType(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.MonitoredItemType = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetMonitoredNodeRange(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.MonitoredNodeRange = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetNodeId(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.NodeId = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetNodeName(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.NodeName = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetNodeSubType(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.NodeSubType = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetOpsOwnerIdList(v []*string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.OpsOwnerIdList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) SetProjectId(v int64) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig {
	s.ProjectId = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfig) Validate() error {
	if s.MonitoredNodeRange != nil {
		if err := s.MonitoredNodeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange struct {
	// The list of filter rules. If left empty, all rules are applied by default.
	FilterRuleList []*UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList `json:"FilterRuleList,omitempty" xml:"FilterRuleList,omitempty" type:"Repeated"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange) GetFilterRuleList() []*UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	return s.FilterRuleList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange) SetFilterRuleList(v []*UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange {
	s.FilterRuleList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRange) Validate() error {
	if s.FilterRuleList != nil {
		for _, item := range s.FilterRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList struct {
	// The filter key.
	//
	// example:
	//
	// NODE_NAME
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The list of filter values. Multiple values are supported.
	FilterValueList []*string `json:"FilterValueList,omitempty" xml:"FilterValueList,omitempty" type:"Repeated"`
	// Specifies whether to exclude the selection.
	IsExclude *bool `json:"IsExclude,omitempty" xml:"IsExclude,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GetFilterKey() *string {
	return s.FilterKey
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GetFilterValueList() []*string {
	return s.FilterValueList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GetIsExclude() *bool {
	return s.IsExclude
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) SetFilterKey(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	s.FilterKey = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) SetFilterValueList(v []*string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	s.FilterValueList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) SetIsExclude(v bool) *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	s.IsExclude = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList struct {
	// The alert cause type. Valid values: FAILED: The node fails to run. TIME_OUT: The node execution times out. UNFINISHED: The node does not generate data. SUCCESS: The node runs successfully. DATA_DELAY: Data is delayed.
	//
	// This parameter is required.
	//
	// example:
	//
	// TIME_OUT
	AlertReasonType *string `json:"AlertReasonType,omitempty" xml:"AlertReasonType,omitempty"`
	// The alert strategy.
	//
	// This parameter is required.
	AlertStrategy *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy `json:"AlertStrategy,omitempty" xml:"AlertStrategy,omitempty" type:"Struct"`
	// The timeout period in minutes. Specify this parameter only when AlertReasonType is set to TIME_OUT.
	//
	// example:
	//
	// 60
	FixedElapsedTime *int64 `json:"FixedElapsedTime,omitempty" xml:"FixedElapsedTime,omitempty"`
	// The daily check time for unfinished nodes in HH:mm format. Specify this parameter only when AlertReasonType is set to UNFINISHED.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
	// The detailed configuration for the unfinished rule. Specify this parameter only when AlertReasonType is set to UNFINISHED.
	UnfinishedRuleDetailConfig *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig `json:"UnfinishedRuleDetailConfig,omitempty" xml:"UnfinishedRuleDetailConfig,omitempty" type:"Struct"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) GetAlertReasonType() *string {
	return s.AlertReasonType
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) GetAlertStrategy() *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy {
	return s.AlertStrategy
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) GetFixedElapsedTime() *int64 {
	return s.FixedElapsedTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) GetUnfinishedRuleDetailConfig() *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig {
	return s.UnfinishedRuleDetailConfig
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) SetAlertReasonType(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList {
	s.AlertReasonType = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) SetAlertStrategy(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList {
	s.AlertStrategy = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) SetFixedElapsedTime(v int64) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList {
	s.FixedElapsedTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) SetFixedUnfinishedTime(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) SetUnfinishedRuleDetailConfig(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList {
	s.UnfinishedRuleDetailConfig = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigList) Validate() error {
	if s.AlertStrategy != nil {
		if err := s.AlertStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.UnfinishedRuleDetailConfig != nil {
		if err := s.UnfinishedRuleDetailConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy struct {
	// The alert trigger frequency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AlertFrequency *int32 `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	// The monitored time range.
	MonitoredTimeRange *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange `json:"MonitoredTimeRange,omitempty" xml:"MonitoredTimeRange,omitempty" type:"Struct"`
	// The list of recipient configurations.
	//
	// This parameter is required.
	ReceiverConfigList []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList `json:"ReceiverConfigList,omitempty" xml:"ReceiverConfigList,omitempty" type:"Repeated"`
	// The maximum number of times an alert can be sent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SendTimes *int32 `json:"SendTimes,omitempty" xml:"SendTimes,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) GetAlertFrequency() *int32 {
	return s.AlertFrequency
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) GetMonitoredTimeRange() *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange {
	return s.MonitoredTimeRange
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) GetReceiverConfigList() []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	return s.ReceiverConfigList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) GetSendTimes() *int32 {
	return s.SendTimes
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) SetAlertFrequency(v int32) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy {
	s.AlertFrequency = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) SetMonitoredTimeRange(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy {
	s.MonitoredTimeRange = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) SetReceiverConfigList(v []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy {
	s.ReceiverConfigList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) SetSendTimes(v int32) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy {
	s.SendTimes = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategy) Validate() error {
	if s.MonitoredTimeRange != nil {
		if err := s.MonitoredTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.ReceiverConfigList != nil {
		for _, item := range s.ReceiverConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange struct {
	// The end time in HH:mm format. Set this parameter to 0 to indicate no limit.
	//
	// example:
	//
	// 11:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time in HH:mm format. Set this parameter to 0 to indicate no limit.
	//
	// example:
	//
	// 07:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) SetEndTime(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) SetStartTime(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList struct {
	// The list of custom alert channel IDs.
	CustomAlertChannelIdList []*string `json:"CustomAlertChannelIdList,omitempty" xml:"CustomAlertChannelIdList,omitempty" type:"Repeated"`
	// The list of on-duty table IDs. This parameter is required when ReceiverType is set to ON_DUTY_TABLE.
	OnDutyTableIdList []*string `json:"OnDutyTableIdList,omitempty" xml:"OnDutyTableIdList,omitempty" type:"Repeated"`
	// The list of receiving methods. Valid values: PHONE: Phone call. MESSAGE: Text message. EMAIL: Email. DING_TALK: DingTalk. FEISHU: Lark. WECHAT: WeChat. CUSTOM: Custom channel.
	//
	// This parameter is required.
	ReceivedWayList []*string `json:"ReceivedWayList,omitempty" xml:"ReceivedWayList,omitempty" type:"Repeated"`
	// The recipient type. Valid values: OWNER: Owner. USER_CUSTOM: Custom user. ON_DUTY_TABLE: On-duty table.
	//
	// This parameter is required.
	//
	// example:
	//
	// OWNER
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The list of recipient user IDs. This parameter is required when ReceiverType is set to USER_CUSTOM.
	ReceiverUserIdList []*string `json:"ReceiverUserIdList,omitempty" xml:"ReceiverUserIdList,omitempty" type:"Repeated"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetOnDutyTableIdList() []*string {
	return s.OnDutyTableIdList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetReceivedWayList() []*string {
	return s.ReceivedWayList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetReceiverUserIdList() []*string {
	return s.ReceiverUserIdList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetCustomAlertChannelIdList(v []*string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetOnDutyTableIdList(v []*string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.OnDutyTableIdList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetReceivedWayList(v []*string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.ReceivedWayList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetReceiverType(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.ReceiverType = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetReceiverUserIdList(v []*string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.ReceiverUserIdList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig struct {
	// The daily unfinished configuration.
	DailyUnfinishedConfig *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig `json:"DailyUnfinishedConfig,omitempty" xml:"DailyUnfinishedConfig,omitempty" type:"Struct"`
	// The multi-period unfinished configuration.
	MultiPeriodUnfinishedConfig *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig `json:"MultiPeriodUnfinishedConfig,omitempty" xml:"MultiPeriodUnfinishedConfig,omitempty" type:"Struct"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) GetDailyUnfinishedConfig() *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig {
	return s.DailyUnfinishedConfig
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) GetMultiPeriodUnfinishedConfig() *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	return s.MultiPeriodUnfinishedConfig
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) SetDailyUnfinishedConfig(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig {
	s.DailyUnfinishedConfig = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) SetMultiPeriodUnfinishedConfig(v *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig {
	s.MultiPeriodUnfinishedConfig = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) Validate() error {
	if s.DailyUnfinishedConfig != nil {
		if err := s.DailyUnfinishedConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MultiPeriodUnfinishedConfig != nil {
		if err := s.MultiPeriodUnfinishedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig struct {
	// The daily check time for unfinished nodes in HH:mm format.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) SetFixedUnfinishedTime(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig struct {
	// The list of custom monitoring periods. Specify this parameter when MonitorPeriodType is set to CUSTOM.
	CustomPeriodList []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList `json:"CustomPeriodList,omitempty" xml:"CustomPeriodList,omitempty" type:"Repeated"`
	// The interval time value. This parameter is used together with IntervalUnit.
	//
	// example:
	//
	// 30
	IntervalAfterDueTime *int32 `json:"IntervalAfterDueTime,omitempty" xml:"IntervalAfterDueTime,omitempty"`
	// The interval unit. Valid values: SECONDS, MINUTES, and HOURS.
	//
	// example:
	//
	// MINUTES
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// The monitoring period type. Valid values: ALL, FIRST, LAST, and CUSTOM.
	//
	// example:
	//
	// ALL
	MonitorPeriodType *string `json:"MonitorPeriodType,omitempty" xml:"MonitorPeriodType,omitempty"`
	// The trigger rule type. Valid values: INTERVAL_AFTER_DUE_TIME: Triggered at an interval after the due time. FIX_TIME: Triggered at a fixed time.
	//
	// example:
	//
	// INTERVAL_AFTER_DUE_TIME
	TriggerRuleType *string `json:"TriggerRuleType,omitempty" xml:"TriggerRuleType,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetCustomPeriodList() []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	return s.CustomPeriodList
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetIntervalAfterDueTime() *int32 {
	return s.IntervalAfterDueTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetMonitorPeriodType() *string {
	return s.MonitorPeriodType
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetTriggerRuleType() *string {
	return s.TriggerRuleType
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetCustomPeriodList(v []*UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.CustomPeriodList = v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetIntervalAfterDueTime(v int32) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.IntervalAfterDueTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetIntervalUnit(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.IntervalUnit = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetMonitorPeriodType(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.MonitorPeriodType = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetTriggerRuleType(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.TriggerRuleType = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) Validate() error {
	if s.CustomPeriodList != nil {
		for _, item := range s.CustomPeriodList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList struct {
	// The check time for unfinished nodes in HH:mm format.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
	// The number of periods.
	//
	// example:
	//
	// 1
	PeriodNum *int32 `json:"PeriodNum,omitempty" xml:"PeriodNum,omitempty"`
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GetPeriodNum() *int32 {
	return s.PeriodNum
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) SetFixedUnfinishedTime(v string) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) SetPeriodNum(v int32) *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	s.PeriodNum = &v
	return s
}

func (s *UpdateScheduleMonitorRequestUpdateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) Validate() error {
	return dara.Validate(s)
}
