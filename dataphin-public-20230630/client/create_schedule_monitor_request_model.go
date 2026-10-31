// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateScheduleMonitorRequestCreateCommand) *CreateScheduleMonitorRequest
	GetCreateCommand() *CreateScheduleMonitorRequestCreateCommand
	SetOpTenantId(v int64) *CreateScheduleMonitorRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateScheduleMonitorRequest
	GetOpUserId() *string
}

type CreateScheduleMonitorRequest struct {
	// The request body for creating a scheduling monitoring configuration.
	//
	// This parameter is required.
	CreateCommand *CreateScheduleMonitorRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateScheduleMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequest) GetCreateCommand() *CreateScheduleMonitorRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateScheduleMonitorRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateScheduleMonitorRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateScheduleMonitorRequest) SetCreateCommand(v *CreateScheduleMonitorRequestCreateCommand) *CreateScheduleMonitorRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateScheduleMonitorRequest) SetOpTenantId(v int64) *CreateScheduleMonitorRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateScheduleMonitorRequest) SetOpUserId(v string) *CreateScheduleMonitorRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateScheduleMonitorRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScheduleMonitorRequestCreateCommand struct {
	// The monitored object configuration.
	//
	// This parameter is required.
	MonitoredItemConfig *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig `json:"MonitoredItemConfig,omitempty" xml:"MonitoredItemConfig,omitempty" type:"Struct"`
	// The list of monitoring rule configurations. You can configure multiple rules for different alert causes.
	//
	// This parameter is required.
	MonitoredRuleConfigList []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList `json:"MonitoredRuleConfigList,omitempty" xml:"MonitoredRuleConfigList,omitempty" type:"Repeated"`
}

func (s CreateScheduleMonitorRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommand) GetMonitoredItemConfig() *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	return s.MonitoredItemConfig
}

func (s *CreateScheduleMonitorRequestCreateCommand) GetMonitoredRuleConfigList() []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList {
	return s.MonitoredRuleConfigList
}

func (s *CreateScheduleMonitorRequestCreateCommand) SetMonitoredItemConfig(v *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) *CreateScheduleMonitorRequestCreateCommand {
	s.MonitoredItemConfig = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommand) SetMonitoredRuleConfigList(v []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) *CreateScheduleMonitorRequestCreateCommand {
	s.MonitoredRuleConfigList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommand) Validate() error {
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

type CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig struct {
	// The business unit ID. Specify this parameter when the summary logical table does not have a project ID.
	//
	// example:
	//
	// 100001
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The logical table ID. Specify this parameter when the monitored object type is LOGICAL_FIELD or LOGICAL_TABLE_NODE.
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
	// The monitored object type. Valid values:
	//
	// - LOGICAL_FIELD: logical field.
	//
	// - OND_ID_NODE: logical table node.
	//
	// - SCHEDULE_PHYSICAL_NODE: physical task with epoch scheduling.
	//
	// - MANUAL_PHYSICAL_NODE: physical task with manual scheduling.
	//
	// - ODM_NODE: source data path task.
	//
	// - LOGICAL_TABLE_NODE: logical table task.
	//
	// - PROJECT: project.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCHEDULE_PHYSICAL_NODE
	MonitoredItemType *string `json:"MonitoredItemType,omitempty" xml:"MonitoredItemType,omitempty"`
	// The scope of monitored tasks, which serves as the filtering rule for project-level monitoring. Specify this parameter when the monitored object type is PROJECT.
	MonitoredNodeRange *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange `json:"MonitoredNodeRange,omitempty" xml:"MonitoredNodeRange,omitempty" type:"Struct"`
	// The ID of the monitored task, which is the node ID. This parameter is required when the monitored object type is not PROJECT.
	//
	// example:
	//
	// 12345678
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the monitored task, which is the node name. This parameter is required when the monitored object type is not PROJECT.
	//
	// example:
	//
	// dwd_order_detail_daily
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The subtype of the node. This parameter is required when the monitored object type is not PROJECT or LOGICAL_FIELD.
	//
	// example:
	//
	// MAX_COMPUTE_SQL
	NodeSubType *string `json:"NodeSubType,omitempty" xml:"NodeSubType,omitempty"`
	// The list of O&M owner IDs for the node. This serves as the source for the OWNER recipient type.
	OpsOwnerIdList []*string `json:"OpsOwnerIdList,omitempty" xml:"OpsOwnerIdList,omitempty" type:"Repeated"`
	// The project ID. This parameter is required when the monitored object type is PROJECT. For other types, it specifies the project to which the monitored object belongs.
	//
	// example:
	//
	// 7086194564164288
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetLogicalTableId() *int64 {
	return s.LogicalTableId
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetLogicalTableName() *string {
	return s.LogicalTableName
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetMonitoredItemType() *string {
	return s.MonitoredItemType
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetMonitoredNodeRange() *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange {
	return s.MonitoredNodeRange
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetNodeSubType() *string {
	return s.NodeSubType
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetOpsOwnerIdList() []*string {
	return s.OpsOwnerIdList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetBizUnitId(v int64) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.BizUnitId = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetLogicalTableId(v int64) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.LogicalTableId = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetLogicalTableName(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.LogicalTableName = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetMonitoredItemType(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.MonitoredItemType = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetMonitoredNodeRange(v *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.MonitoredNodeRange = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetNodeId(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.NodeId = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetNodeName(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.NodeName = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetNodeSubType(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.NodeSubType = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetOpsOwnerIdList(v []*string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.OpsOwnerIdList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) SetProjectId(v int64) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig {
	s.ProjectId = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfig) Validate() error {
	if s.MonitoredNodeRange != nil {
		if err := s.MonitoredNodeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange struct {
	// The list of filter rules. If this parameter is left empty, all tasks are monitored by default.
	FilterRuleList []*CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList `json:"FilterRuleList,omitempty" xml:"FilterRuleList,omitempty" type:"Repeated"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange) GetFilterRuleList() []*CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	return s.FilterRuleList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange) SetFilterRuleList(v []*CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange {
	s.FilterRuleList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRange) Validate() error {
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

type CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList struct {
	// The filter key.
	//
	// example:
	//
	// NODE_NAME
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The list of filter values. Multiple values are supported.
	FilterValueList []*string `json:"FilterValueList,omitempty" xml:"FilterValueList,omitempty" type:"Repeated"`
	// Specifies whether to invert the selection.
	IsExclude *bool `json:"IsExclude,omitempty" xml:"IsExclude,omitempty"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GetFilterKey() *string {
	return s.FilterKey
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GetFilterValueList() []*string {
	return s.FilterValueList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) GetIsExclude() *bool {
	return s.IsExclude
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) SetFilterKey(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	s.FilterKey = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) SetFilterValueList(v []*string) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	s.FilterValueList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) SetIsExclude(v bool) *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList {
	s.IsExclude = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredItemConfigMonitoredNodeRangeFilterRuleList) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList struct {
	// The alert cause type. Valid values:
	//
	// - FAILED: The task execution failed.
	//
	// - TIME_OUT: The task execution timed out.
	//
	// - UNFINISHED: The task is not completed.
	//
	// - SUCCESS: The task execution succeeded.
	//
	// - DATA_DELAY: The data is delayed.
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
	AlertStrategy *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy `json:"AlertStrategy,omitempty" xml:"AlertStrategy,omitempty" type:"Struct"`
	// The timeout period in minutes. Specify this parameter only when AlertReasonType is set to TIME_OUT.
	//
	// example:
	//
	// 60
	FixedElapsedTime *int64 `json:"FixedElapsedTime,omitempty" xml:"FixedElapsedTime,omitempty"`
	// The daily time to check for unfinished tasks in HH:mm format. Specify this parameter only when AlertReasonType is set to UNFINISHED.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
	// The detailed configuration for the unfinished rule. Specify this parameter only when AlertReasonType is set to UNFINISHED.
	UnfinishedRuleDetailConfig *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig `json:"UnfinishedRuleDetailConfig,omitempty" xml:"UnfinishedRuleDetailConfig,omitempty" type:"Struct"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) GetAlertReasonType() *string {
	return s.AlertReasonType
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) GetAlertStrategy() *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy {
	return s.AlertStrategy
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) GetFixedElapsedTime() *int64 {
	return s.FixedElapsedTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) GetUnfinishedRuleDetailConfig() *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig {
	return s.UnfinishedRuleDetailConfig
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) SetAlertReasonType(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList {
	s.AlertReasonType = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) SetAlertStrategy(v *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList {
	s.AlertStrategy = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) SetFixedElapsedTime(v int64) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList {
	s.FixedElapsedTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) SetFixedUnfinishedTime(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) SetUnfinishedRuleDetailConfig(v *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList {
	s.UnfinishedRuleDetailConfig = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigList) Validate() error {
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

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy struct {
	// The alert trigger frequency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AlertFrequency *int32 `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	// The monitored time range.
	MonitoredTimeRange *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange `json:"MonitoredTimeRange,omitempty" xml:"MonitoredTimeRange,omitempty" type:"Struct"`
	// The list of recipient configurations.
	//
	// This parameter is required.
	ReceiverConfigList []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList `json:"ReceiverConfigList,omitempty" xml:"ReceiverConfigList,omitempty" type:"Repeated"`
	// The maximum number of alert notifications to send.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SendTimes *int32 `json:"SendTimes,omitempty" xml:"SendTimes,omitempty"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) GetAlertFrequency() *int32 {
	return s.AlertFrequency
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) GetMonitoredTimeRange() *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange {
	return s.MonitoredTimeRange
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) GetReceiverConfigList() []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	return s.ReceiverConfigList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) GetSendTimes() *int32 {
	return s.SendTimes
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) SetAlertFrequency(v int32) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy {
	s.AlertFrequency = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) SetMonitoredTimeRange(v *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy {
	s.MonitoredTimeRange = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) SetReceiverConfigList(v []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy {
	s.ReceiverConfigList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) SetSendTimes(v int32) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy {
	s.SendTimes = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategy) Validate() error {
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

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange struct {
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

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) SetEndTime(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange {
	s.EndTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) SetStartTime(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange {
	s.StartTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyMonitoredTimeRange) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList struct {
	// The list of custom alert channel IDs.
	CustomAlertChannelIdList []*string `json:"CustomAlertChannelIdList,omitempty" xml:"CustomAlertChannelIdList,omitempty" type:"Repeated"`
	// The list of on-duty table IDs. This parameter is required when ReceiverType is set to ON_DUTY_TABLE.
	OnDutyTableIdList []*string `json:"OnDutyTableIdList,omitempty" xml:"OnDutyTableIdList,omitempty" type:"Repeated"`
	// The list of notification methods. Valid values:
	//
	// - PHONE: Phone call.
	//
	// - MESSAGE: Text message.
	//
	// - EMAIL: Email.
	//
	// - DING_TALK: DingTalk.
	//
	// - FEISHU: Lark.
	//
	// - WECHAT: WeChat.
	//
	// - CUSTOM: Custom channel.
	//
	// This parameter is required.
	ReceivedWayList []*string `json:"ReceivedWayList,omitempty" xml:"ReceivedWayList,omitempty" type:"Repeated"`
	// The recipient type. Valid values:
	//
	// - OWNER: The owner.
	//
	// - USER_CUSTOM: A custom user.
	//
	// - ON_DUTY_TABLE: An on-duty table.
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

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetOnDutyTableIdList() []*string {
	return s.OnDutyTableIdList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetReceivedWayList() []*string {
	return s.ReceivedWayList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) GetReceiverUserIdList() []*string {
	return s.ReceiverUserIdList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetCustomAlertChannelIdList(v []*string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetOnDutyTableIdList(v []*string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.OnDutyTableIdList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetReceivedWayList(v []*string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.ReceivedWayList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetReceiverType(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.ReceiverType = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) SetReceiverUserIdList(v []*string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList {
	s.ReceiverUserIdList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListAlertStrategyReceiverConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig struct {
	// The daily unfinished configuration.
	DailyUnfinishedConfig *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig `json:"DailyUnfinishedConfig,omitempty" xml:"DailyUnfinishedConfig,omitempty" type:"Struct"`
	// The multi-period unfinished configuration.
	MultiPeriodUnfinishedConfig *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig `json:"MultiPeriodUnfinishedConfig,omitempty" xml:"MultiPeriodUnfinishedConfig,omitempty" type:"Struct"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) GetDailyUnfinishedConfig() *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig {
	return s.DailyUnfinishedConfig
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) GetMultiPeriodUnfinishedConfig() *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	return s.MultiPeriodUnfinishedConfig
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) SetDailyUnfinishedConfig(v *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig {
	s.DailyUnfinishedConfig = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) SetMultiPeriodUnfinishedConfig(v *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig {
	s.MultiPeriodUnfinishedConfig = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfig) Validate() error {
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

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig struct {
	// The daily time to check for unfinished tasks in HH:mm format.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) SetFixedUnfinishedTime(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigDailyUnfinishedConfig) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig struct {
	// The list of custom monitoring periods. Specify this parameter when MonitorPeriodType is set to CUSTOM.
	CustomPeriodList []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList `json:"CustomPeriodList,omitempty" xml:"CustomPeriodList,omitempty" type:"Repeated"`
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
	// The trigger rule type. Valid values:
	//
	// - INTERVAL_AFTER_DUE_TIME: Triggered at an interval after the due time.
	//
	// - FIX_TIME: Triggered at a fixed time.
	//
	// example:
	//
	// INTERVAL_AFTER_DUE_TIME
	TriggerRuleType *string `json:"TriggerRuleType,omitempty" xml:"TriggerRuleType,omitempty"`
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetCustomPeriodList() []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	return s.CustomPeriodList
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetIntervalAfterDueTime() *int32 {
	return s.IntervalAfterDueTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetMonitorPeriodType() *string {
	return s.MonitorPeriodType
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetTriggerRuleType() *string {
	return s.TriggerRuleType
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetCustomPeriodList(v []*CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.CustomPeriodList = v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetIntervalAfterDueTime(v int32) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.IntervalAfterDueTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetIntervalUnit(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.IntervalUnit = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetMonitorPeriodType(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.MonitorPeriodType = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetTriggerRuleType(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.TriggerRuleType = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) Validate() error {
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

type CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList struct {
	// The time to check for unfinished tasks in HH:mm format.
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

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GetPeriodNum() *int32 {
	return s.PeriodNum
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) SetFixedUnfinishedTime(v string) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) SetPeriodNum(v int32) *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	s.PeriodNum = &v
	return s
}

func (s *CreateScheduleMonitorRequestCreateCommandMonitoredRuleConfigListUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) Validate() error {
	return dara.Validate(s)
}
