// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListScheduleMonitorRequestListQuery) *ListScheduleMonitorRequest
	GetListQuery() *ListScheduleMonitorRequestListQuery
	SetOpTenantId(v int64) *ListScheduleMonitorRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListScheduleMonitorRequest
	GetOpUserId() *string
}

type ListScheduleMonitorRequest struct {
	// The query conditions for scheduling monitoring and alerting.
	//
	// This parameter is required.
	ListQuery *ListScheduleMonitorRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The user ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListScheduleMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorRequest) GetListQuery() *ListScheduleMonitorRequestListQuery {
	return s.ListQuery
}

func (s *ListScheduleMonitorRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListScheduleMonitorRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListScheduleMonitorRequest) SetListQuery(v *ListScheduleMonitorRequestListQuery) *ListScheduleMonitorRequest {
	s.ListQuery = v
	return s
}

func (s *ListScheduleMonitorRequest) SetOpTenantId(v int64) *ListScheduleMonitorRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListScheduleMonitorRequest) SetOpUserId(v string) *ListScheduleMonitorRequest {
	s.OpUserId = &v
	return s
}

func (s *ListScheduleMonitorRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleMonitorRequestListQuery struct {
	// The list of alert cause types. Valid values: FAILED (the task execution failed), TIME_OUT (the task execution timed out), UNFINISHED (the task did not generate output), SUCCESS (the task execution succeeded), and DATA_DELAY (the data is delayed).
	AlertReasonList []*string `json:"AlertReasonList,omitempty" xml:"AlertReasonList,omitempty" type:"Repeated"`
	// The keyword for search. Supports fuzzy matching for monitored task names. If this parameter is empty or not specified, no filtering is applied.
	//
	// example:
	//
	// dwd_order_detail
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The list of logical field IDs.
	LogicFieldIdList []*int64 `json:"LogicFieldIdList,omitempty" xml:"LogicFieldIdList,omitempty" type:"Repeated"`
	// The list of logical table IDs.
	LogicalTableIdList []*int64 `json:"LogicalTableIdList,omitempty" xml:"LogicalTableIdList,omitempty" type:"Repeated"`
	// The list of on-call schedule IDs.
	OnCallTableList []*string `json:"OnCallTableList,omitempty" xml:"OnCallTableList,omitempty" type:"Repeated"`
	// The list of user IDs of the creators of monitored objects.
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The current page number. Pages start from 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of project IDs. If this parameter is empty, all projects that the current user has permissions on are automatically queried.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of notification method types. Valid values: PHONE (phone call), MESSAGE (text message), EMAIL (email), DING_TALK (DingTalk), FEISHU (Lark), WECHAT (WeChat), and CUSTOM (custom channel).
	ReceivedWayTypeList []*string `json:"ReceivedWayTypeList,omitempty" xml:"ReceivedWayTypeList,omitempty" type:"Repeated"`
	// The list of user IDs of the recipients.
	ReceiverList []*string `json:"ReceiverList,omitempty" xml:"ReceiverList,omitempty" type:"Repeated"`
	// The list of monitored object types. Valid values: LOGICAL_FIELD (logical field), OND_ID_NODE (logical table node), SCHEDULE_PHYSICAL_NODE (periodically scheduled physical task), MANUAL_PHYSICAL_NODE (manually scheduled physical task), ODM_NODE (source path task), LOGICAL_TABLE_NODE (logical table task), and PROJECT (project).
	//
	// This parameter is required.
	TypeList []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s ListScheduleMonitorRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorRequestListQuery) GetAlertReasonList() []*string {
	return s.AlertReasonList
}

func (s *ListScheduleMonitorRequestListQuery) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListScheduleMonitorRequestListQuery) GetLogicFieldIdList() []*int64 {
	return s.LogicFieldIdList
}

func (s *ListScheduleMonitorRequestListQuery) GetLogicalTableIdList() []*int64 {
	return s.LogicalTableIdList
}

func (s *ListScheduleMonitorRequestListQuery) GetOnCallTableList() []*string {
	return s.OnCallTableList
}

func (s *ListScheduleMonitorRequestListQuery) GetOwnerList() []*string {
	return s.OwnerList
}

func (s *ListScheduleMonitorRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScheduleMonitorRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduleMonitorRequestListQuery) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListScheduleMonitorRequestListQuery) GetReceivedWayTypeList() []*string {
	return s.ReceivedWayTypeList
}

func (s *ListScheduleMonitorRequestListQuery) GetReceiverList() []*string {
	return s.ReceiverList
}

func (s *ListScheduleMonitorRequestListQuery) GetTypeList() []*string {
	return s.TypeList
}

func (s *ListScheduleMonitorRequestListQuery) SetAlertReasonList(v []*string) *ListScheduleMonitorRequestListQuery {
	s.AlertReasonList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetKeyWord(v string) *ListScheduleMonitorRequestListQuery {
	s.KeyWord = &v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetLogicFieldIdList(v []*int64) *ListScheduleMonitorRequestListQuery {
	s.LogicFieldIdList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetLogicalTableIdList(v []*int64) *ListScheduleMonitorRequestListQuery {
	s.LogicalTableIdList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetOnCallTableList(v []*string) *ListScheduleMonitorRequestListQuery {
	s.OnCallTableList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetOwnerList(v []*string) *ListScheduleMonitorRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetPageNo(v int32) *ListScheduleMonitorRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetPageSize(v int32) *ListScheduleMonitorRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetProjectIds(v []*int64) *ListScheduleMonitorRequestListQuery {
	s.ProjectIds = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetReceivedWayTypeList(v []*string) *ListScheduleMonitorRequestListQuery {
	s.ReceivedWayTypeList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetReceiverList(v []*string) *ListScheduleMonitorRequestListQuery {
	s.ReceiverList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) SetTypeList(v []*string) *ListScheduleMonitorRequestListQuery {
	s.TypeList = v
	return s
}

func (s *ListScheduleMonitorRequestListQuery) Validate() error {
	return dara.Validate(s)
}
