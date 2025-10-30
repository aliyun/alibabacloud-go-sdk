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
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQuery *ListNodesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
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
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SCRIPT
	NodeBizType *string `json:"NodeBizType,omitempty" xml:"NodeBizType,omitempty"`
	// This parameter is required.
	NodeSubBizTypeList []*string `json:"NodeSubBizTypeList,omitempty" xml:"NodeSubBizTypeList,omitempty" type:"Repeated"`
	OwnerList          []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PriorityList []*string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12111
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// true
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// NORMAL
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
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
