// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScheduleMonitorResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScheduleMonitorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScheduleMonitorResponseBody
	GetMessage() *string
	SetPageResult(v *ListScheduleMonitorResponseBodyPageResult) *ListScheduleMonitorResponseBody
	GetPageResult() *ListScheduleMonitorResponseBodyPageResult
	SetRequestId(v string) *ListScheduleMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScheduleMonitorResponseBody
	GetSuccess() *bool
}

type ListScheduleMonitorResponseBody struct {
	// The error code of the request. A value of OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message of the request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paginated result of the scheduling monitoring alert query.
	PageResult *ListScheduleMonitorResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScheduleMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScheduleMonitorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScheduleMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScheduleMonitorResponseBody) GetPageResult() *ListScheduleMonitorResponseBodyPageResult {
	return s.PageResult
}

func (s *ListScheduleMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduleMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduleMonitorResponseBody) SetCode(v string) *ListScheduleMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduleMonitorResponseBody) SetHttpStatusCode(v int32) *ListScheduleMonitorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScheduleMonitorResponseBody) SetMessage(v string) *ListScheduleMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduleMonitorResponseBody) SetPageResult(v *ListScheduleMonitorResponseBodyPageResult) *ListScheduleMonitorResponseBody {
	s.PageResult = v
	return s
}

func (s *ListScheduleMonitorResponseBody) SetRequestId(v string) *ListScheduleMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleMonitorResponseBody) SetSuccess(v bool) *ListScheduleMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduleMonitorResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleMonitorResponseBodyPageResult struct {
	// The list of monitored objects on the current page.
	MonitorItemList []*ListScheduleMonitorResponseBodyPageResultMonitorItemList `json:"MonitorItemList,omitempty" xml:"MonitorItemList,omitempty" type:"Repeated"`
	// The total number of monitored objects that meet the conditions.
	//
	// example:
	//
	// 128
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResult) GetMonitorItemList() []*ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	return s.MonitorItemList
}

func (s *ListScheduleMonitorResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScheduleMonitorResponseBodyPageResult) SetMonitorItemList(v []*ListScheduleMonitorResponseBodyPageResultMonitorItemList) *ListScheduleMonitorResponseBodyPageResult {
	s.MonitorItemList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResult) SetTotalCount(v int32) *ListScheduleMonitorResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResult) Validate() error {
	if s.MonitorItemList != nil {
		for _, item := range s.MonitorItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemList struct {
	// The business segment information.
	BizUnit *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit `json:"BizUnit,omitempty" xml:"BizUnit,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-10 15:30:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The logical table information.
	LogicTable *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable `json:"LogicTable,omitempty" xml:"LogicTable,omitempty" type:"Struct"`
	// The last modification time.
	//
	// example:
	//
	// 2026-09-10 15:30:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The scope of the monitoring task.
	MonitorNodeRange *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange `json:"MonitorNodeRange,omitempty" xml:"MonitorNodeRange,omitempty" type:"Struct"`
	// The ID of the monitored object.
	//
	// example:
	//
	// 6912345678901234
	MonitoredItemId *int64 `json:"MonitoredItemId,omitempty" xml:"MonitoredItemId,omitempty"`
	// The list of monitoring rules.
	MonitoredRuleList []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList `json:"MonitoredRuleList,omitempty" xml:"MonitoredRuleList,omitempty" type:"Repeated"`
	// The monitored task information.
	MonitoredTask *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask `json:"MonitoredTask,omitempty" xml:"MonitoredTask,omitempty" type:"Struct"`
	// The project ID.
	//
	// example:
	//
	// 7086194564164288
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project information.
	ProjectInfo *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// The business subtype.
	//
	// example:
	//
	// MAX_COMPUTE_SQL
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 30001011
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The monitored object type.
	//
	// example:
	//
	// SCHEDULE_PHYSICAL_NODE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemList) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemList) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetBizUnit() *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit {
	return s.BizUnit
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetLogicTable() *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	return s.LogicTable
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetMonitorNodeRange() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange {
	return s.MonitorNodeRange
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetMonitoredItemId() *int64 {
	return s.MonitoredItemId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetMonitoredRuleList() []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	return s.MonitoredRuleList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetMonitoredTask() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask {
	return s.MonitoredTask
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetProjectInfo() *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo {
	return s.ProjectInfo
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetSubType() *string {
	return s.SubType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) GetType() *string {
	return s.Type
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetBizUnit(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.BizUnit = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetCreateTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.CreateTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetLogicTable(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.LogicTable = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetModifiedTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.ModifiedTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetMonitorNodeRange(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.MonitorNodeRange = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetMonitoredItemId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.MonitoredItemId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetMonitoredRuleList(v []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.MonitoredRuleList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetMonitoredTask(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.MonitoredTask = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetProjectId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.ProjectId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetProjectInfo(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.ProjectInfo = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetSubType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.SubType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetTenantId(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.TenantId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) SetType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemList {
	s.Type = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemList) Validate() error {
	if s.BizUnit != nil {
		if err := s.BizUnit.Validate(); err != nil {
			return err
		}
	}
	if s.LogicTable != nil {
		if err := s.LogicTable.Validate(); err != nil {
			return err
		}
	}
	if s.MonitorNodeRange != nil {
		if err := s.MonitorNodeRange.Validate(); err != nil {
			return err
		}
	}
	if s.MonitoredRuleList != nil {
		for _, item := range s.MonitoredRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MonitoredTask != nil {
		if err := s.MonitoredTask.Validate(); err != nil {
			return err
		}
	}
	if s.ProjectInfo != nil {
		if err := s.ProjectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit struct {
	// The display name of the business segment.
	//
	// example:
	//
	// Test segment
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the business segment.
	//
	// example:
	//
	// 6612345678901234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the business segment.
	//
	// example:
	//
	// bd_demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) GetId() *int64 {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) GetName() *string {
	return s.Name
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) SetDisplayName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit {
	s.DisplayName = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) SetId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) SetName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit {
	s.Name = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListBizUnit) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable struct {
	// The ID of the business segment.
	//
	// example:
	//
	// 6612345678901234
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the monitoring task.
	//
	// example:
	//
	// 6812345678901234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitoring task.
	//
	// example:
	//
	// dim_user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subtype of the node.
	//
	// example:
	//
	// MAX_COMPUTE_SQL
	NodeSubType *string `json:"NodeSubType,omitempty" xml:"NodeSubType,omitempty"`
	// The list of O&M owner IDs.
	OwnerIdList []*string `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Repeated"`
	// The name of the project to which the task belongs.
	//
	// example:
	//
	// ds_demo
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The type of the scheduling period.
	//
	// example:
	//
	// DAILY
	SchedulePeriodType *string `json:"SchedulePeriodType,omitempty" xml:"SchedulePeriodType,omitempty"`
	// Specifies whether data delay is supported.
	SupportDataDelay *bool `json:"SupportDataDelay,omitempty" xml:"SupportDataDelay,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetId() *string {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetName() *string {
	return s.Name
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetNodeSubType() *string {
	return s.NodeSubType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetOwnerIdList() []*string {
	return s.OwnerIdList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetSchedulePeriodType() *string {
	return s.SchedulePeriodType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) GetSupportDataDelay() *bool {
	return s.SupportDataDelay
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetBizUnitId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.BizUnitId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetId(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.Name = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetNodeSubType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.NodeSubType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetOwnerIdList(v []*string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.OwnerIdList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetProjectName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.ProjectName = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetSchedulePeriodType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.SchedulePeriodType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) SetSupportDataDelay(v bool) *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable {
	s.SupportDataDelay = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListLogicTable) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange struct {
	// The list of filter rules.
	FilterRuleList []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList `json:"FilterRuleList,omitempty" xml:"FilterRuleList,omitempty" type:"Repeated"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange) GetFilterRuleList() []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList {
	return s.FilterRuleList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange) SetFilterRuleList(v []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange {
	s.FilterRuleList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRange) Validate() error {
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

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList struct {
	// The filter key.
	//
	// example:
	//
	// NODE_NAME
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The list of filter values.
	FilterValueList []*string `json:"FilterValueList,omitempty" xml:"FilterValueList,omitempty" type:"Repeated"`
	// Specifies whether to invert the selection.
	IsExclude *bool `json:"IsExclude,omitempty" xml:"IsExclude,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) GetFilterValueList() []*string {
	return s.FilterValueList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) GetIsExclude() *bool {
	return s.IsExclude
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) SetFilterKey(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList {
	s.FilterKey = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) SetFilterValueList(v []*string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList {
	s.FilterValueList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) SetIsExclude(v bool) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList {
	s.IsExclude = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitorNodeRangeFilterRuleList) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList struct {
	// The reason for the alert.
	AlertReason *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason `json:"AlertReason,omitempty" xml:"AlertReason,omitempty" type:"Struct"`
	// The alert strategy.
	AlertStrategy *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy `json:"AlertStrategy,omitempty" xml:"AlertStrategy,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-10 15:30:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	Creator *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// The monitoring rule ID.
	//
	// example:
	//
	// 6912345678909876
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2026-09-10 15:30:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The last modifier.
	Modifier *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	// The ID of the associated monitoring item.
	//
	// example:
	//
	// 6912345678901234
	MonitoredItemId *int64 `json:"MonitoredItemId,omitempty" xml:"MonitoredItemId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 7086194564164288
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The enablement status of the rule.
	SwitchStatus *bool `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 30001011
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetAlertReason() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason {
	return s.AlertReason
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetAlertStrategy() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy {
	return s.AlertStrategy
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetCreator() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator {
	return s.Creator
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetId() *int64 {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetModifier() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier {
	return s.Modifier
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetMonitoredItemId() *int64 {
	return s.MonitoredItemId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetSwitchStatus() *bool {
	return s.SwitchStatus
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetAlertReason(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.AlertReason = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetAlertStrategy(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.AlertStrategy = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetCreateTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.CreateTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetCreator(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.Creator = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetModifiedTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.ModifiedTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetModifier(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.Modifier = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetMonitoredItemId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.MonitoredItemId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetProjectId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.ProjectId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetSwitchStatus(v bool) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.SwitchStatus = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) SetTenantId(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList {
	s.TenantId = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleList) Validate() error {
	if s.AlertReason != nil {
		if err := s.AlertReason.Validate(); err != nil {
			return err
		}
	}
	if s.AlertStrategy != nil {
		if err := s.AlertStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason struct {
	// The type of the alert reason. Valid values: FAILED, TIME_OUT, UNFINISHED, SUCCESS, and DATA_DELAY.
	//
	// example:
	//
	// TIME_OUT
	AlertReasonType *string `json:"AlertReasonType,omitempty" xml:"AlertReasonType,omitempty"`
	// The timeout period in minutes. This parameter is returned only when the alert reason type is TIME_OUT.
	//
	// example:
	//
	// 60
	FixedElapsedTime *int64 `json:"FixedElapsedTime,omitempty" xml:"FixedElapsedTime,omitempty"`
	// The daily check time for unfinished tasks in HH:mm format. This parameter is returned only when the alert reason type is UNFINISHED.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
	// The detailed configuration of the unfinished rule. This parameter is returned only when the alert reason type is UNFINISHED.
	UnfinishedRuleDetailConfig *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig `json:"UnfinishedRuleDetailConfig,omitempty" xml:"UnfinishedRuleDetailConfig,omitempty" type:"Struct"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) GetAlertReasonType() *string {
	return s.AlertReasonType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) GetFixedElapsedTime() *int64 {
	return s.FixedElapsedTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) GetUnfinishedRuleDetailConfig() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig {
	return s.UnfinishedRuleDetailConfig
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) SetAlertReasonType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason {
	s.AlertReasonType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) SetFixedElapsedTime(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason {
	s.FixedElapsedTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) SetFixedUnfinishedTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) SetUnfinishedRuleDetailConfig(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason {
	s.UnfinishedRuleDetailConfig = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReason) Validate() error {
	if s.UnfinishedRuleDetailConfig != nil {
		if err := s.UnfinishedRuleDetailConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig struct {
	// The daily configuration for unfinished tasks.
	DailyUnfinishedConfig *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig `json:"DailyUnfinishedConfig,omitempty" xml:"DailyUnfinishedConfig,omitempty" type:"Struct"`
	// The multi-period unfinished configuration.
	MultiPeriodUnfinishedConfig *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig `json:"MultiPeriodUnfinishedConfig,omitempty" xml:"MultiPeriodUnfinishedConfig,omitempty" type:"Struct"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) GetDailyUnfinishedConfig() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig {
	return s.DailyUnfinishedConfig
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) GetMultiPeriodUnfinishedConfig() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	return s.MultiPeriodUnfinishedConfig
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) SetDailyUnfinishedConfig(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig {
	s.DailyUnfinishedConfig = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) SetMultiPeriodUnfinishedConfig(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig {
	s.MultiPeriodUnfinishedConfig = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfig) Validate() error {
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

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig struct {
	// The time of day for the unfinished check, in HH:mm format.
	//
	// example:
	//
	// 07:00
	FixedUnfinishedTime *string `json:"FixedUnfinishedTime,omitempty" xml:"FixedUnfinishedTime,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig) SetFixedUnfinishedTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigDailyUnfinishedConfig) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig struct {
	// The list of custom monitoring periods.
	CustomPeriodList []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList `json:"CustomPeriodList,omitempty" xml:"CustomPeriodList,omitempty" type:"Repeated"`
	// The interval value.
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
	// The trigger rule type.
	//
	// example:
	//
	// INTERVAL_AFTER_DUE_TIME
	TriggerRuleType *string `json:"TriggerRuleType,omitempty" xml:"TriggerRuleType,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetCustomPeriodList() []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	return s.CustomPeriodList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetIntervalAfterDueTime() *int32 {
	return s.IntervalAfterDueTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetMonitorPeriodType() *string {
	return s.MonitorPeriodType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) GetTriggerRuleType() *string {
	return s.TriggerRuleType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetCustomPeriodList(v []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.CustomPeriodList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetIntervalAfterDueTime(v int32) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.IntervalAfterDueTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetIntervalUnit(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.IntervalUnit = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetMonitorPeriodType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.MonitorPeriodType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) SetTriggerRuleType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig {
	s.TriggerRuleType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfig) Validate() error {
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

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList struct {
	// The time for the unfinished check, in HH:mm format.
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

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GetFixedUnfinishedTime() *string {
	return s.FixedUnfinishedTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) GetPeriodNum() *int32 {
	return s.PeriodNum
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) SetFixedUnfinishedTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	s.FixedUnfinishedTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) SetPeriodNum(v int32) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList {
	s.PeriodNum = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertReasonUnfinishedRuleDetailConfigMultiPeriodUnfinishedConfigCustomPeriodList) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy struct {
	// The alert trigger frequency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	AlertFrequency *int32 `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	// The monitored time range.
	MonitoredTimeRange *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange `json:"MonitoredTimeRange,omitempty" xml:"MonitoredTimeRange,omitempty" type:"Struct"`
	// The list of recipient configurations.
	//
	// This parameter is required.
	ReceiverConfigList []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList `json:"ReceiverConfigList,omitempty" xml:"ReceiverConfigList,omitempty" type:"Repeated"`
	// The maximum number of alert notifications to send.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SendTimes *int32 `json:"SendTimes,omitempty" xml:"SendTimes,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) GetAlertFrequency() *int32 {
	return s.AlertFrequency
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) GetMonitoredTimeRange() *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange {
	return s.MonitoredTimeRange
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) GetReceiverConfigList() []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList {
	return s.ReceiverConfigList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) GetSendTimes() *int32 {
	return s.SendTimes
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) SetAlertFrequency(v int32) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy {
	s.AlertFrequency = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) SetMonitoredTimeRange(v *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy {
	s.MonitoredTimeRange = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) SetReceiverConfigList(v []*ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy {
	s.ReceiverConfigList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) SetSendTimes(v int32) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy {
	s.SendTimes = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategy) Validate() error {
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

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange struct {
	// The end time in HH:mm format. A value of 0 indicates no limit.
	//
	// example:
	//
	// 11:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time in HH:mm format. A value of 0 indicates no limit.
	//
	// example:
	//
	// 07:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) SetEndTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange {
	s.EndTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) SetStartTime(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange {
	s.StartTime = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyMonitoredTimeRange) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList struct {
	// The list of custom alert channel IDs.
	CustomAlertChannelIdList []*string `json:"CustomAlertChannelIdList,omitempty" xml:"CustomAlertChannelIdList,omitempty" type:"Repeated"`
	// The list of on-duty table IDs. This parameter is required when ReceiverType is set to ON_DUTY_TABLE.
	OnDutyTableIdList []*string `json:"OnDutyTableIdList,omitempty" xml:"OnDutyTableIdList,omitempty" type:"Repeated"`
	// The list of notification methods. Valid values: PHONE, MESSAGE, EMAIL, DING_TALK, FEISHU, WECHAT, and CUSTOM.
	//
	// This parameter is required.
	ReceivedWayList []*string `json:"ReceivedWayList,omitempty" xml:"ReceivedWayList,omitempty" type:"Repeated"`
	// The recipient type. Valid values: OWNER (owner), USER_CUSTOM (custom user), and ON_DUTY_TABLE (on-duty table).
	//
	// This parameter is required.
	//
	// example:
	//
	// USER_CUSTOM
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The list of recipient user IDs. This parameter is required when ReceiverType is set to USER_CUSTOM.
	ReceiverUserIdList []*string `json:"ReceiverUserIdList,omitempty" xml:"ReceiverUserIdList,omitempty" type:"Repeated"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) GetOnDutyTableIdList() []*string {
	return s.OnDutyTableIdList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) GetReceivedWayList() []*string {
	return s.ReceivedWayList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) GetReceiverUserIdList() []*string {
	return s.ReceiverUserIdList
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) SetCustomAlertChannelIdList(v []*string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) SetOnDutyTableIdList(v []*string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList {
	s.OnDutyTableIdList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) SetReceivedWayList(v []*string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList {
	s.ReceivedWayList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) SetReceiverType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList {
	s.ReceiverType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) SetReceiverUserIdList(v []*string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList {
	s.ReceiverUserIdList = v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListAlertStrategyReceiverConfigList) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator struct {
	// The user ID.
	//
	// example:
	//
	// 200012345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// Zhang San
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) GetId() *string {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) GetName() *string {
	return s.Name
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) SetId(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) SetName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator {
	s.Name = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListCreator) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier struct {
	// The user ID.
	//
	// example:
	//
	// 200012345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// Li Si
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) GetId() *string {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) GetName() *string {
	return s.Name
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) SetId(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) SetName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier {
	s.Name = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredRuleListModifier) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask struct {
	// The ID of the monitoring task.
	//
	// example:
	//
	// 6712345678901234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitoring task.
	//
	// example:
	//
	// dwd_order_detail
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the project to which the task belongs.
	//
	// example:
	//
	// ds_demo
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The schedule period type. Valid values: MINUTELY, HOURLY, DAILY, WEEKLY, MONTHLY, and YEARLY.
	//
	// example:
	//
	// DAILY
	SchedulePeriodType *string `json:"SchedulePeriodType,omitempty" xml:"SchedulePeriodType,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) GetId() *string {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) GetName() *string {
	return s.Name
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) GetSchedulePeriodType() *string {
	return s.SchedulePeriodType
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) SetId(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) SetName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask {
	s.Name = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) SetProjectName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask {
	s.ProjectName = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) SetSchedulePeriodType(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask {
	s.SchedulePeriodType = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListMonitoredTask) Validate() error {
	return dara.Validate(s)
}

type ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo struct {
	// The project ID.
	//
	// example:
	//
	// 7086194564164288
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project name.
	//
	// example:
	//
	// ds_demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) GetId() *int64 {
	return s.Id
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) GetName() *string {
	return s.Name
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) SetId(v int64) *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo {
	s.Id = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) SetName(v string) *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo {
	s.Name = &v
	return s
}

func (s *ListScheduleMonitorResponseBodyPageResultMonitorItemListProjectInfo) Validate() error {
	return dara.Validate(s)
}
