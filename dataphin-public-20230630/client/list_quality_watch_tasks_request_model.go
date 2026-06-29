// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityWatchTasksRequestListQuery) *ListQualityWatchTasksRequest
	GetListQuery() *ListQualityWatchTasksRequestListQuery
	SetOpTenantId(v int64) *ListQualityWatchTasksRequest
	GetOpTenantId() *int64
}

type ListQualityWatchTasksRequest struct {
	// The paged query conditions.
	ListQuery *ListQualityWatchTasksRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityWatchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksRequest) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksRequest) GetListQuery() *ListQualityWatchTasksRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityWatchTasksRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityWatchTasksRequest) SetListQuery(v *ListQualityWatchTasksRequestListQuery) *ListQualityWatchTasksRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityWatchTasksRequest) SetOpTenantId(v int64) *ListQualityWatchTasksRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityWatchTasksRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityWatchTasksRequestListQuery struct {
	// The business date filter.
	//
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The business unit names.
	BizUnitNameList []*string `json:"BizUnitNameList,omitempty" xml:"BizUnitNameList,omitempty" type:"Repeated"`
	// Specifies whether to query only the quality monitoring node objects owned by the current user.
	CurrentUserOwned *bool `json:"CurrentUserOwned,omitempty" xml:"CurrentUserOwned,omitempty"`
	// The data source IDs.
	DataSourceIdList []*string `json:"DataSourceIdList,omitempty" xml:"DataSourceIdList,omitempty" type:"Repeated"`
	// The data source owners.
	DataSourceOwnerList []*string `json:"DataSourceOwnerList,omitempty" xml:"DataSourceOwnerList,omitempty" type:"Repeated"`
	// The data source scopes. Valid values:
	//
	// - STREAMING: real-time only.
	//
	// - OFFLINE: offline only.
	//
	// - ALL: real-time and offline.
	DataSourceScopeList []*string `json:"DataSourceScopeList,omitempty" xml:"DataSourceScopeList,omitempty" type:"Repeated"`
	// The data source types, such as MAX_COMPUTE, HADOOP, and MYSQL.
	DataSourceTypeList []*string `json:"DataSourceTypeList,omitempty" xml:"DataSourceTypeList,omitempty" type:"Repeated"`
	// The rule exception types. Valid values:
	//
	// - STRONG: strong.
	//
	// - WEAK: weak.
	ErrorRuleStrengthList []*string `json:"ErrorRuleStrengthList,omitempty" xml:"ErrorRuleStrengthList,omitempty" type:"Repeated"`
	// The search keyword, which is the name of the monitored table.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project names.
	ProjectNameList []*string `json:"ProjectNameList,omitempty" xml:"ProjectNameList,omitempty" type:"Repeated"`
	// The quality owners.
	QualityOwnerList []*string `json:"QualityOwnerList,omitempty" xml:"QualityOwnerList,omitempty" type:"Repeated"`
	// The task statuses. Valid values:
	//
	// - NOT_RUN: not executed.
	//
	// - WAITING: waiting.
	//
	// - RUNNING: running.
	//
	// - SUCCESS: succeeded.
	//
	// - FAILED: failed.
	//
	// - CANCEL: canceled.
	//
	// - TIMEOUT: timed out.
	//
	// - OFFLINE: offline.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The table owners.
	TableOwnerList []*string `json:"TableOwnerList,omitempty" xml:"TableOwnerList,omitempty" type:"Repeated"`
	// The table types. Valid values:
	//
	// - LOGIC_DIM_TABLE: logical dimension table.
	//
	// - LOGIC_FACT_TABLE: logical fact table.
	//
	// - LOGIC_SUM_TABLE: logical aggregate table.
	//
	// - LOGIC_LABEL_TABLE: logical label table.
	//
	// - PHYSICAL_TABLE: physical table.
	//
	// - REALTIME_LOGICAL_TABLE: real-time meta table.
	TableTypeList []*string `json:"TableTypeList,omitempty" xml:"TableTypeList,omitempty" type:"Repeated"`
	// The monitored object types. Valid values:
	//
	// - TABLE: Dataphin table.
	//
	// - DATASOURCE_TABLE: global table.
	//
	// - DATASOURCE: data source.
	//
	// - INDEX: metric.
	//
	// - REALTIME_LOGICAL_TABLE: real-time meta table.
	WatchTypeList []*string `json:"WatchTypeList,omitempty" xml:"WatchTypeList,omitempty" type:"Repeated"`
}

func (s ListQualityWatchTasksRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksRequestListQuery) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityWatchTasksRequestListQuery) GetBizUnitNameList() []*string {
	return s.BizUnitNameList
}

func (s *ListQualityWatchTasksRequestListQuery) GetCurrentUserOwned() *bool {
	return s.CurrentUserOwned
}

func (s *ListQualityWatchTasksRequestListQuery) GetDataSourceIdList() []*string {
	return s.DataSourceIdList
}

func (s *ListQualityWatchTasksRequestListQuery) GetDataSourceOwnerList() []*string {
	return s.DataSourceOwnerList
}

func (s *ListQualityWatchTasksRequestListQuery) GetDataSourceScopeList() []*string {
	return s.DataSourceScopeList
}

func (s *ListQualityWatchTasksRequestListQuery) GetDataSourceTypeList() []*string {
	return s.DataSourceTypeList
}

func (s *ListQualityWatchTasksRequestListQuery) GetErrorRuleStrengthList() []*string {
	return s.ErrorRuleStrengthList
}

func (s *ListQualityWatchTasksRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListQualityWatchTasksRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListQualityWatchTasksRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityWatchTasksRequestListQuery) GetProjectNameList() []*string {
	return s.ProjectNameList
}

func (s *ListQualityWatchTasksRequestListQuery) GetQualityOwnerList() []*string {
	return s.QualityOwnerList
}

func (s *ListQualityWatchTasksRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListQualityWatchTasksRequestListQuery) GetTableOwnerList() []*string {
	return s.TableOwnerList
}

func (s *ListQualityWatchTasksRequestListQuery) GetTableTypeList() []*string {
	return s.TableTypeList
}

func (s *ListQualityWatchTasksRequestListQuery) GetWatchTypeList() []*string {
	return s.WatchTypeList
}

func (s *ListQualityWatchTasksRequestListQuery) SetBizDate(v string) *ListQualityWatchTasksRequestListQuery {
	s.BizDate = &v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetBizUnitNameList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.BizUnitNameList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetCurrentUserOwned(v bool) *ListQualityWatchTasksRequestListQuery {
	s.CurrentUserOwned = &v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetDataSourceIdList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.DataSourceIdList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetDataSourceOwnerList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.DataSourceOwnerList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetDataSourceScopeList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.DataSourceScopeList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetDataSourceTypeList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.DataSourceTypeList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetErrorRuleStrengthList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.ErrorRuleStrengthList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetKeyword(v string) *ListQualityWatchTasksRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetPageNo(v int32) *ListQualityWatchTasksRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetPageSize(v int32) *ListQualityWatchTasksRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetProjectNameList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.ProjectNameList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetQualityOwnerList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.QualityOwnerList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetStatusList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetTableOwnerList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.TableOwnerList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetTableTypeList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.TableTypeList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) SetWatchTypeList(v []*string) *ListQualityWatchTasksRequestListQuery {
	s.WatchTypeList = v
	return s
}

func (s *ListQualityWatchTasksRequestListQuery) Validate() error {
	return dara.Validate(s)
}
