// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityWatchesRequestListQuery) *ListQualityWatchesRequest
	GetListQuery() *ListQualityWatchesRequestListQuery
	SetOpTenantId(v int64) *ListQualityWatchesRequest
	GetOpTenantId() *int64
}

type ListQualityWatchesRequest struct {
	// The paged query conditions.
	ListQuery *ListQualityWatchesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityWatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesRequest) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesRequest) GetListQuery() *ListQualityWatchesRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityWatchesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityWatchesRequest) SetListQuery(v *ListQualityWatchesRequestListQuery) *ListQualityWatchesRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityWatchesRequest) SetOpTenantId(v int64) *ListQualityWatchesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityWatchesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityWatchesRequestListQuery struct {
	// The business unit names.
	BizUnitNameList []*string `json:"BizUnitNameList,omitempty" xml:"BizUnitNameList,omitempty" type:"Repeated"`
	// Specifies whether to query only monitored objects owned by the current user.
	CurrentUserOwned *bool `json:"CurrentUserOwned,omitempty" xml:"CurrentUserOwned,omitempty"`
	// The data source IDs.
	DataSourceIdList []*int64 `json:"DataSourceIdList,omitempty" xml:"DataSourceIdList,omitempty" type:"Repeated"`
	// The data source owners.
	DataSourceOwnerList []*string `json:"DataSourceOwnerList,omitempty" xml:"DataSourceOwnerList,omitempty" type:"Repeated"`
	// The data source scope. Valid values:
	//
	// - STREAMING: real-time only
	//
	// - OFFLINE: offline only
	//
	// - ALL: real-time and offline.
	DataSourceScopeList []*string `json:"DataSourceScopeList,omitempty" xml:"DataSourceScopeList,omitempty" type:"Repeated"`
	// The data source type, such as MAX_COMPUTE, HADOOP, or MYSQL.
	DataSourceTypeList []*string `json:"DataSourceTypeList,omitempty" xml:"DataSourceTypeList,omitempty" type:"Repeated"`
	// The metric computation type. Valid values:
	//
	// - AUTO: automated coding
	//
	// - CUSTOM: expert coding
	//
	// - MOUNT: external table registration
	//
	// - COMBINE: derived metric specific.
	IndexComputeTypeList []*string `json:"IndexComputeTypeList,omitempty" xml:"IndexComputeTypeList,omitempty" type:"Repeated"`
	// The metric owners.
	IndexOwnerList []*string `json:"IndexOwnerList,omitempty" xml:"IndexOwnerList,omitempty" type:"Repeated"`
	// The search keyword. This is the name of the monitored table.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The task status. Valid values:
	//
	// - NOT_RUN: not executed
	//
	// - WAITING: waiting
	//
	// - RUNNING: executing
	//
	// - SUCCESS: executed successfully
	//
	// - FAILED: execution failed
	//
	// - CANCEL: canceled
	//
	// - TIMEOUT: timed out
	//
	// - OFFLINE: offline.
	LatestWatchTaskStatusList []*string `json:"LatestWatchTaskStatusList,omitempty" xml:"LatestWatchTaskStatusList,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project names.
	ProjectNameList []*string `json:"ProjectNameList,omitempty" xml:"ProjectNameList,omitempty" type:"Repeated"`
	// The quality owners.
	QualityOwnerList []*string `json:"QualityOwnerList,omitempty" xml:"QualityOwnerList,omitempty" type:"Repeated"`
	// The status of the monitored object. Valid values:
	//
	// - ENABLE: enabled
	//
	// - DISABLE: disabled.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The table owners.
	TableOwnerList []*string `json:"TableOwnerList,omitempty" xml:"TableOwnerList,omitempty" type:"Repeated"`
	// The table type. Valid values:
	//
	// - LOGIC_DIM_TABLE: logical dimension table
	//
	// - LOGIC_FACT_TABLE: logical fact table
	//
	// - LOGIC_SUM_TABLE: logical aggregate table
	//
	// - LOGIC_LABEL_TABLE: logical label table
	//
	// - PHYSICAL_TABLE: physical table
	//
	// - REALTIME_LOGICAL_TABLE: real-time meta table.
	TableTypeList []*string `json:"TableTypeList,omitempty" xml:"TableTypeList,omitempty" type:"Repeated"`
	// The monitored object type. Valid values:
	//
	// - TABLE: Dataphin table
	//
	// - DATASOURCE_TABLE: full-domain table
	//
	// - DATASOURCE: data source
	//
	// - INDEX: metric
	//
	// - REALTIME_LOGICAL_TABLE: real-time meta table.
	WatchTypeList []*string `json:"WatchTypeList,omitempty" xml:"WatchTypeList,omitempty" type:"Repeated"`
}

func (s ListQualityWatchesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesRequestListQuery) GetBizUnitNameList() []*string {
	return s.BizUnitNameList
}

func (s *ListQualityWatchesRequestListQuery) GetCurrentUserOwned() *bool {
	return s.CurrentUserOwned
}

func (s *ListQualityWatchesRequestListQuery) GetDataSourceIdList() []*int64 {
	return s.DataSourceIdList
}

func (s *ListQualityWatchesRequestListQuery) GetDataSourceOwnerList() []*string {
	return s.DataSourceOwnerList
}

func (s *ListQualityWatchesRequestListQuery) GetDataSourceScopeList() []*string {
	return s.DataSourceScopeList
}

func (s *ListQualityWatchesRequestListQuery) GetDataSourceTypeList() []*string {
	return s.DataSourceTypeList
}

func (s *ListQualityWatchesRequestListQuery) GetIndexComputeTypeList() []*string {
	return s.IndexComputeTypeList
}

func (s *ListQualityWatchesRequestListQuery) GetIndexOwnerList() []*string {
	return s.IndexOwnerList
}

func (s *ListQualityWatchesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListQualityWatchesRequestListQuery) GetLatestWatchTaskStatusList() []*string {
	return s.LatestWatchTaskStatusList
}

func (s *ListQualityWatchesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListQualityWatchesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityWatchesRequestListQuery) GetProjectNameList() []*string {
	return s.ProjectNameList
}

func (s *ListQualityWatchesRequestListQuery) GetQualityOwnerList() []*string {
	return s.QualityOwnerList
}

func (s *ListQualityWatchesRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListQualityWatchesRequestListQuery) GetTableOwnerList() []*string {
	return s.TableOwnerList
}

func (s *ListQualityWatchesRequestListQuery) GetTableTypeList() []*string {
	return s.TableTypeList
}

func (s *ListQualityWatchesRequestListQuery) GetWatchTypeList() []*string {
	return s.WatchTypeList
}

func (s *ListQualityWatchesRequestListQuery) SetBizUnitNameList(v []*string) *ListQualityWatchesRequestListQuery {
	s.BizUnitNameList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetCurrentUserOwned(v bool) *ListQualityWatchesRequestListQuery {
	s.CurrentUserOwned = &v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetDataSourceIdList(v []*int64) *ListQualityWatchesRequestListQuery {
	s.DataSourceIdList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetDataSourceOwnerList(v []*string) *ListQualityWatchesRequestListQuery {
	s.DataSourceOwnerList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetDataSourceScopeList(v []*string) *ListQualityWatchesRequestListQuery {
	s.DataSourceScopeList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetDataSourceTypeList(v []*string) *ListQualityWatchesRequestListQuery {
	s.DataSourceTypeList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetIndexComputeTypeList(v []*string) *ListQualityWatchesRequestListQuery {
	s.IndexComputeTypeList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetIndexOwnerList(v []*string) *ListQualityWatchesRequestListQuery {
	s.IndexOwnerList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetKeyword(v string) *ListQualityWatchesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetLatestWatchTaskStatusList(v []*string) *ListQualityWatchesRequestListQuery {
	s.LatestWatchTaskStatusList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetPageNo(v int32) *ListQualityWatchesRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetPageSize(v int32) *ListQualityWatchesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetProjectNameList(v []*string) *ListQualityWatchesRequestListQuery {
	s.ProjectNameList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetQualityOwnerList(v []*string) *ListQualityWatchesRequestListQuery {
	s.QualityOwnerList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetStatusList(v []*string) *ListQualityWatchesRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetTableOwnerList(v []*string) *ListQualityWatchesRequestListQuery {
	s.TableOwnerList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetTableTypeList(v []*string) *ListQualityWatchesRequestListQuery {
	s.TableTypeList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) SetWatchTypeList(v []*string) *ListQualityWatchesRequestListQuery {
	s.WatchTypeList = v
	return s
}

func (s *ListQualityWatchesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
