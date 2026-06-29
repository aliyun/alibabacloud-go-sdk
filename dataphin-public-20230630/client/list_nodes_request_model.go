// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListNodesRequest
	GetEnv() *string
	SetListQuery(v *ListNodesRequestListQuery) *ListNodesRequest
	GetListQuery() *ListNodesRequestListQuery
	SetOpTenantId(v int64) *ListNodesRequest
	GetOpTenantId() *int64
}

type ListNodesRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: development environment
	//
	// - PROD (default): production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The query conditions.
	//
	// This parameter is required.
	ListQuery *ListNodesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetEnv() *string {
	return s.Env
}

func (s *ListNodesRequest) GetListQuery() *ListNodesRequestListQuery {
	return s.ListQuery
}

func (s *ListNodesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListNodesRequest) SetEnv(v string) *ListNodesRequest {
	s.Env = &v
	return s
}

func (s *ListNodesRequest) SetListQuery(v *ListNodesRequestListQuery) *ListNodesRequest {
	s.ListQuery = v
	return s
}

func (s *ListNodesRequest) SetOpTenantId(v int64) *ListNodesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListNodesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesRequestListQuery struct {
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The node business type. Valid values:
	//
	// - SCRIPT: script
	//
	// - LOGICAL_TABLE: logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCRIPT
	NodeBizType *string `json:"NodeBizType,omitempty" xml:"NodeBizType,omitempty"`
	// The sub-business types. Valid values:
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
	// - DATABASE_SQL.
	//
	// This parameter is required.
	NodeSubBizTypeList []*string `json:"NodeSubBizTypeList,omitempty" xml:"NodeSubBizTypeList,omitempty" type:"Repeated"`
	// The user IDs of the owners.
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The node priorities. Valid values:
	//
	// - HIGHEST
	//
	// - HIGH
	//
	// - MIDDLE
	//
	// - LOW
	//
	// - LOWEST.
	PriorityList []*string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12111
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether scheduling is paused.
	//
	// example:
	//
	// true
	SchedulePaused *bool `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	// The scheduling periods. Valid values:
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
	// - MINUTELY.
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// The node scheduling type. Valid values:
	//
	// - NORMAL: periodic scheduling
	//
	// - SUPPLEMENT: data backfill
	//
	// - MANUAL: manual scheduling.
	//
	// example:
	//
	// NORMAL
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The search keyword. Fuzzy search by node name and exact search by node ID are supported.
	//
	// example:
	//
	// xx
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListNodesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListNodesRequestListQuery) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListNodesRequestListQuery) GetNodeBizType() *string {
	return s.NodeBizType
}

func (s *ListNodesRequestListQuery) GetNodeSubBizTypeList() []*string {
	return s.NodeSubBizTypeList
}

func (s *ListNodesRequestListQuery) GetOwnerList() []*string {
	return s.OwnerList
}

func (s *ListNodesRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListNodesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesRequestListQuery) GetPriorityList() []*string {
	return s.PriorityList
}

func (s *ListNodesRequestListQuery) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesRequestListQuery) GetSchedulePaused() *bool {
	return s.SchedulePaused
}

func (s *ListNodesRequestListQuery) GetSchedulePeriodList() []*string {
	return s.SchedulePeriodList
}

func (s *ListNodesRequestListQuery) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListNodesRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListNodesRequestListQuery) SetDryRun(v bool) *ListNodesRequestListQuery {
	s.DryRun = &v
	return s
}

func (s *ListNodesRequestListQuery) SetNodeBizType(v string) *ListNodesRequestListQuery {
	s.NodeBizType = &v
	return s
}

func (s *ListNodesRequestListQuery) SetNodeSubBizTypeList(v []*string) *ListNodesRequestListQuery {
	s.NodeSubBizTypeList = v
	return s
}

func (s *ListNodesRequestListQuery) SetOwnerList(v []*string) *ListNodesRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListNodesRequestListQuery) SetPage(v int32) *ListNodesRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListNodesRequestListQuery) SetPageSize(v int32) *ListNodesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequestListQuery) SetPriorityList(v []*string) *ListNodesRequestListQuery {
	s.PriorityList = v
	return s
}

func (s *ListNodesRequestListQuery) SetProjectId(v int64) *ListNodesRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequestListQuery) SetSchedulePaused(v bool) *ListNodesRequestListQuery {
	s.SchedulePaused = &v
	return s
}

func (s *ListNodesRequestListQuery) SetSchedulePeriodList(v []*string) *ListNodesRequestListQuery {
	s.SchedulePeriodList = v
	return s
}

func (s *ListNodesRequestListQuery) SetScheduleType(v string) *ListNodesRequestListQuery {
	s.ScheduleType = &v
	return s
}

func (s *ListNodesRequestListQuery) SetSearchText(v string) *ListNodesRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListNodesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
