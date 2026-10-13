// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingSkillJobPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentIds(v []*int64) *ListMetaProfilingSkillJobPageRequest
	GetComponentIds() []*int64
	SetDsId(v int64) *ListMetaProfilingSkillJobPageRequest
	GetDsId() *int64
	SetDsName(v string) *ListMetaProfilingSkillJobPageRequest
	GetDsName() *string
	SetDsType(v string) *ListMetaProfilingSkillJobPageRequest
	GetDsType() *string
	SetGroupBy(v string) *ListMetaProfilingSkillJobPageRequest
	GetGroupBy() *string
	SetIsDisplaySystemTask(v bool) *ListMetaProfilingSkillJobPageRequest
	GetIsDisplaySystemTask() *bool
	SetJobId(v int64) *ListMetaProfilingSkillJobPageRequest
	GetJobId() *int64
	SetJobName(v string) *ListMetaProfilingSkillJobPageRequest
	GetJobName() *string
	SetNeedTotalCount(v bool) *ListMetaProfilingSkillJobPageRequest
	GetNeedTotalCount() *bool
	SetOrderBy(v string) *ListMetaProfilingSkillJobPageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMetaProfilingSkillJobPageRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ListMetaProfilingSkillJobPageRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingSkillJobPageRequest
	GetPageSize() *int32
}

type ListMetaProfilingSkillJobPageRequest struct {
	// The list of component IDs used to filter profiling tasks by multiple data source components.
	ComponentIds []*int64 `json:"componentIds,omitempty" xml:"componentIds,omitempty" type:"Repeated"`
	// The data source ID.
	//
	// example:
	//
	// 3001
	DsId *int64 `json:"dsId,omitempty" xml:"dsId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// mysql_source_01
	DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
	// The data source type identifier.
	//
	// example:
	//
	// MYSQL
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// The name of the grouping field used for aggregation.
	//
	// example:
	//
	// dbName
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// Specifies whether to display system built-in tasks.
	IsDisplaySystemTask *bool `json:"isDisplaySystemTask,omitempty" xml:"isDisplaySystemTask,omitempty"`
	// The profiling task ID.
	//
	// example:
	//
	// 100256
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The profiling task name.
	//
	// example:
	//
	// profiling_daily
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// Specifies whether to return the total number of records.
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The name of the sort field.
	//
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort direction, such as ascending (ASC) or descending (DESC).
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of records to return per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMetaProfilingSkillJobPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingSkillJobPageRequest) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingSkillJobPageRequest) GetComponentIds() []*int64 {
	return s.ComponentIds
}

func (s *ListMetaProfilingSkillJobPageRequest) GetDsId() *int64 {
	return s.DsId
}

func (s *ListMetaProfilingSkillJobPageRequest) GetDsName() *string {
	return s.DsName
}

func (s *ListMetaProfilingSkillJobPageRequest) GetDsType() *string {
	return s.DsType
}

func (s *ListMetaProfilingSkillJobPageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListMetaProfilingSkillJobPageRequest) GetIsDisplaySystemTask() *bool {
	return s.IsDisplaySystemTask
}

func (s *ListMetaProfilingSkillJobPageRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListMetaProfilingSkillJobPageRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListMetaProfilingSkillJobPageRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ListMetaProfilingSkillJobPageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetaProfilingSkillJobPageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMetaProfilingSkillJobPageRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingSkillJobPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingSkillJobPageRequest) SetComponentIds(v []*int64) *ListMetaProfilingSkillJobPageRequest {
	s.ComponentIds = v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetDsId(v int64) *ListMetaProfilingSkillJobPageRequest {
	s.DsId = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetDsName(v string) *ListMetaProfilingSkillJobPageRequest {
	s.DsName = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetDsType(v string) *ListMetaProfilingSkillJobPageRequest {
	s.DsType = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetGroupBy(v string) *ListMetaProfilingSkillJobPageRequest {
	s.GroupBy = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetIsDisplaySystemTask(v bool) *ListMetaProfilingSkillJobPageRequest {
	s.IsDisplaySystemTask = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetJobId(v int64) *ListMetaProfilingSkillJobPageRequest {
	s.JobId = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetJobName(v string) *ListMetaProfilingSkillJobPageRequest {
	s.JobName = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetNeedTotalCount(v bool) *ListMetaProfilingSkillJobPageRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetOrderBy(v string) *ListMetaProfilingSkillJobPageRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetOrderDirection(v string) *ListMetaProfilingSkillJobPageRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetPageIndex(v int32) *ListMetaProfilingSkillJobPageRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) SetPageSize(v int32) *ListMetaProfilingSkillJobPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageRequest) Validate() error {
	return dara.Validate(s)
}
