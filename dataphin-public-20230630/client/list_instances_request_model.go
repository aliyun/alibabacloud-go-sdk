// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListInstancesRequest
	GetEnv() *string
	SetListQuery(v *ListInstancesRequestListQuery) *ListInstancesRequest
	GetListQuery() *ListInstancesRequestListQuery
	SetOpTenantId(v int64) *ListInstancesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListInstancesRequest
	GetOpUserId() *string
}

type ListInstancesRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: Development environment.
	//
	// - PROD (default): Production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The query request.
	ListQuery *ListInstancesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
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

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetEnv() *string {
	return s.Env
}

func (s *ListInstancesRequest) GetListQuery() *ListInstancesRequestListQuery {
	return s.ListQuery
}

func (s *ListInstancesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListInstancesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListInstancesRequest) SetEnv(v string) *ListInstancesRequest {
	s.Env = &v
	return s
}

func (s *ListInstancesRequest) SetListQuery(v *ListInstancesRequestListQuery) *ListInstancesRequest {
	s.ListQuery = v
	return s
}

func (s *ListInstancesRequest) SetOpTenantId(v int64) *ListInstancesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListInstancesRequest) SetOpUserId(v string) *ListInstancesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesRequestListQuery struct {
	// The business type. Valid values:
	//
	// - SCRIPT: Script instance.
	//
	// - LOGICAL_TABLE: Logical table.
	//
	// example:
	//
	// SCRIPT
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The business unit ID. Required when querying aggregate logical tables.
	//
	// example:
	//
	// 6232322111
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 1021
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The end business date and time. The time format must match the partition format specified by the business unit.
	//
	// example:
	//
	// 2024-05-31
	MaxBizDate *string `json:"MaxBizDate,omitempty" xml:"MaxBizDate,omitempty"`
	// The maximum instance run time.
	//
	// example:
	//
	// 2024-05-31
	MaxRunDate *string `json:"MaxRunDate,omitempty" xml:"MaxRunDate,omitempty"`
	// The start business date and time. The time format must match the partition format specified by the business unit.
	//
	// example:
	//
	// 2024-05-30
	MinBizDate *string `json:"MinBizDate,omitempty" xml:"MinBizDate,omitempty"`
	// The minimum instance run time.
	//
	// example:
	//
	// 2024-05-30
	MinRunDate *string `json:"MinRunDate,omitempty" xml:"MinRunDate,omitempty"`
	// The node ID.
	//
	// example:
	//
	// n_23131
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node owners.
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority. Valid values:
	//
	// - HIGHEST
	//
	// - HIGH
	//
	// - MIDDLE
	//
	// - LOW
	//
	// - LOWEST
	PriorityList []*string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131311111321
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The run status. Valid values:
	//
	// - INIT: Init.
	//
	// - WAIT_SUBMISSION: Waiting for submission.
	//
	// - WAIT_SCHEDULE: Waiting for schedule time.
	//
	// - DISPATCH_BLOCKED: Throttled.
	//
	// - WAIT_RESOURCE: Waiting for schedule resource.
	//
	// - RUNNING: Running.
	//
	// - SUCCESS: Succeeded.
	//
	// - FAILED: Failed.
	RunStatusList []*string `json:"RunStatusList,omitempty" xml:"RunStatusList,omitempty" type:"Repeated"`
	// Specifies whether scheduling is paused.
	SchedulePaused *bool `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	// The scheduling period. Valid values:
	//
	// - YEARLY
	//
	// - MONTHLY
	//
	// - WEEKLY
	//
	// - DAILY
	//
	// - HOURLY
	//
	// - MINUTELY
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// The instance scheduling type. Valid values:
	//
	// - NORMAL: Periodic instance.
	//
	// - MANUAL: Manual instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// Fuzzy match by node name or exact match by node ID.
	//
	// example:
	//
	// xx
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// The sub-business type. Valid values:
	//
	// - MAX_COMPUTE_SQL
	//
	// - HIVE_SQL
	//
	// - SHELL
	//
	// - PYTHON
	//
	// - ONE_SERVICE_SQL
	//
	// - DATABASE_SQL
	SubBizTypeList []*string `json:"SubBizTypeList,omitempty" xml:"SubBizTypeList,omitempty" type:"Repeated"`
	// The node tag filter list. Each element is a numeric string of a node tag ID (such as "123"). Filters the instance list by node tags. If not specified or empty, no filtering is applied and all instances are returned. Multiple tags use OR logic. Invalid elements (non-numeric or overflow) are ignored.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s ListInstancesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestListQuery) GetBizType() *string {
	return s.BizType
}

func (s *ListInstancesRequestListQuery) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListInstancesRequestListQuery) GetFlowId() *string {
	return s.FlowId
}

func (s *ListInstancesRequestListQuery) GetMaxBizDate() *string {
	return s.MaxBizDate
}

func (s *ListInstancesRequestListQuery) GetMaxRunDate() *string {
	return s.MaxRunDate
}

func (s *ListInstancesRequestListQuery) GetMinBizDate() *string {
	return s.MinBizDate
}

func (s *ListInstancesRequestListQuery) GetMinRunDate() *string {
	return s.MinRunDate
}

func (s *ListInstancesRequestListQuery) GetNodeId() *string {
	return s.NodeId
}

func (s *ListInstancesRequestListQuery) GetOwnerList() []*string {
	return s.OwnerList
}

func (s *ListInstancesRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListInstancesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequestListQuery) GetPriorityList() []*string {
	return s.PriorityList
}

func (s *ListInstancesRequestListQuery) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListInstancesRequestListQuery) GetRunStatusList() []*string {
	return s.RunStatusList
}

func (s *ListInstancesRequestListQuery) GetSchedulePaused() *bool {
	return s.SchedulePaused
}

func (s *ListInstancesRequestListQuery) GetSchedulePeriodList() []*string {
	return s.SchedulePeriodList
}

func (s *ListInstancesRequestListQuery) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListInstancesRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListInstancesRequestListQuery) GetSubBizTypeList() []*string {
	return s.SubBizTypeList
}

func (s *ListInstancesRequestListQuery) GetTagList() []*string {
	return s.TagList
}

func (s *ListInstancesRequestListQuery) SetBizType(v string) *ListInstancesRequestListQuery {
	s.BizType = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetBizUnitId(v int64) *ListInstancesRequestListQuery {
	s.BizUnitId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetFlowId(v string) *ListInstancesRequestListQuery {
	s.FlowId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMaxBizDate(v string) *ListInstancesRequestListQuery {
	s.MaxBizDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMaxRunDate(v string) *ListInstancesRequestListQuery {
	s.MaxRunDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMinBizDate(v string) *ListInstancesRequestListQuery {
	s.MinBizDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMinRunDate(v string) *ListInstancesRequestListQuery {
	s.MinRunDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetNodeId(v string) *ListInstancesRequestListQuery {
	s.NodeId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetOwnerList(v []*string) *ListInstancesRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetPage(v int32) *ListInstancesRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetPageSize(v int32) *ListInstancesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetPriorityList(v []*string) *ListInstancesRequestListQuery {
	s.PriorityList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetProjectId(v int64) *ListInstancesRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetRunStatusList(v []*string) *ListInstancesRequestListQuery {
	s.RunStatusList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetSchedulePaused(v bool) *ListInstancesRequestListQuery {
	s.SchedulePaused = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetSchedulePeriodList(v []*string) *ListInstancesRequestListQuery {
	s.SchedulePeriodList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetScheduleType(v string) *ListInstancesRequestListQuery {
	s.ScheduleType = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetSearchText(v string) *ListInstancesRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetSubBizTypeList(v []*string) *ListInstancesRequestListQuery {
	s.SubBizTypeList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetTagList(v []*string) *ListInstancesRequestListQuery {
	s.TagList = v
	return s
}

func (s *ListInstancesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
