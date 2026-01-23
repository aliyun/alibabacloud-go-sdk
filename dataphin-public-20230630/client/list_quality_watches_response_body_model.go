// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityWatchesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListQualityWatchesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListQualityWatchesResponseBody
	GetMessage() *string
	SetPageResult(v *ListQualityWatchesResponseBodyPageResult) *ListQualityWatchesResponseBody
	GetPageResult() *ListQualityWatchesResponseBodyPageResult
	SetRequestId(v string) *ListQualityWatchesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityWatchesResponseBody
	GetSuccess() *bool
}

type ListQualityWatchesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListQualityWatchesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityWatchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityWatchesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityWatchesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityWatchesResponseBody) GetPageResult() *ListQualityWatchesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListQualityWatchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityWatchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityWatchesResponseBody) SetCode(v string) *ListQualityWatchesResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityWatchesResponseBody) SetHttpStatusCode(v int32) *ListQualityWatchesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityWatchesResponseBody) SetMessage(v string) *ListQualityWatchesResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityWatchesResponseBody) SetPageResult(v *ListQualityWatchesResponseBodyPageResult) *ListQualityWatchesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListQualityWatchesResponseBody) SetRequestId(v string) *ListQualityWatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityWatchesResponseBody) SetSuccess(v bool) *ListQualityWatchesResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityWatchesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityWatchesResponseBodyPageResult struct {
	QualityWatchList []*ListQualityWatchesResponseBodyPageResultQualityWatchList `json:"QualityWatchList,omitempty" xml:"QualityWatchList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityWatchesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponseBodyPageResult) GetQualityWatchList() []*ListQualityWatchesResponseBodyPageResultQualityWatchList {
	return s.QualityWatchList
}

func (s *ListQualityWatchesResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityWatchesResponseBodyPageResult) SetQualityWatchList(v []*ListQualityWatchesResponseBodyPageResultQualityWatchList) *ListQualityWatchesResponseBodyPageResult {
	s.QualityWatchList = v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResult) SetTotalCount(v int64) *ListQualityWatchesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResult) Validate() error {
	if s.QualityWatchList != nil {
		for _, item := range s.QualityWatchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityWatchesResponseBodyPageResultQualityWatchList struct {
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName    *string                                                                 `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	DataSourceInfo *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo `json:"DataSourceInfo,omitempty" xml:"DataSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1
	EnabledRuleCount *int64 `json:"EnabledRuleCount,omitempty" xml:"EnabledRuleCount,omitempty"`
	// example:
	//
	// 11
	Id        *int64                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	IndexInfo *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo `json:"IndexInfo,omitempty" xml:"IndexInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1
	LatestWatchTaskId *int64 `json:"LatestWatchTaskId,omitempty" xml:"LatestWatchTaskId,omitempty"`
	// example:
	//
	// SUCCESS
	LatestWatchTaskStatus *string `json:"LatestWatchTaskStatus,omitempty" xml:"LatestWatchTaskStatus,omitempty"`
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012011
	QualityOwner *string `json:"QualityOwner,omitempty" xml:"QualityOwner,omitempty"`
	// example:
	//
	// test
	QualityOwnerName *string `json:"QualityOwnerName,omitempty" xml:"QualityOwnerName,omitempty"`
	// example:
	//
	// 11
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// example:
	//
	// ENABLE
	Status    *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TableInfo *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchList) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetCreator() *string {
	return s.Creator
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetDataSourceInfo() *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	return s.DataSourceInfo
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetEnabledRuleCount() *int64 {
	return s.EnabledRuleCount
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetId() *int64 {
	return s.Id
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetIndexInfo() *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	return s.IndexInfo
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetLatestWatchTaskId() *int64 {
	return s.LatestWatchTaskId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetLatestWatchTaskStatus() *string {
	return s.LatestWatchTaskStatus
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetModifier() *string {
	return s.Modifier
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetName() *string {
	return s.Name
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetQualityOwner() *string {
	return s.QualityOwner
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetQualityOwnerName() *string {
	return s.QualityOwnerName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetStatus() *string {
	return s.Status
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetTableInfo() *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	return s.TableInfo
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) GetType() *string {
	return s.Type
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetCreateTime(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.CreateTime = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetCreator(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.Creator = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetCreatorName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.CreatorName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetDataSourceInfo(v *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.DataSourceInfo = v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetEnabledRuleCount(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.EnabledRuleCount = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.Id = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetIndexInfo(v *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.IndexInfo = v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetLatestWatchTaskId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.LatestWatchTaskId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetLatestWatchTaskStatus(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.LatestWatchTaskStatus = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetModifier(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.Modifier = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetModifyTime(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.ModifyTime = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.Name = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetQualityOwner(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.QualityOwner = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetQualityOwnerName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.QualityOwnerName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetRuleCount(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.RuleCount = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetStatus(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.Status = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetTableInfo(v *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.TableInfo = v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) SetType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchList {
	s.Type = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchList) Validate() error {
	if s.DataSourceInfo != nil {
		if err := s.DataSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.IndexInfo != nil {
		if err := s.IndexInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TableInfo != nil {
		if err := s.TableInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo struct {
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetEnv() *string {
	return s.Env
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetId() *string {
	return s.Id
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetOwner() *string {
	return s.Owner
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetCreateTime(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetCreator(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetCreatorName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetEnv(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.Env = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetId(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.Id = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetModifyTime(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetOwner(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetOwnerName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) SetType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo {
	s.Type = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo struct {
	// example:
	//
	// 1121
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// test
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// dws_all
	CellSumLogicTableName *string `json:"CellSumLogicTableName,omitempty" xml:"CellSumLogicTableName,omitempty"`
	// example:
	//
	// AUTO
	ComputeType *string `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	// example:
	//
	// bigint
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// logic
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 全站汇总表
	GranularityDisplayName *string `json:"GranularityDisplayName,omitempty" xml:"GranularityDisplayName,omitempty"`
	// example:
	//
	// 18755764
	GranularityId *int64 `json:"GranularityId,omitempty" xml:"GranularityId,omitempty"`
	// example:
	//
	// 1121
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// 11
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// logic
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// INDEX
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetCellSumLogicTableName() *string {
	return s.CellSumLogicTableName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetComputeType() *string {
	return s.ComputeType
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetDateType() *string {
	return s.DateType
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetDescription() *string {
	return s.Description
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetGranularityDisplayName() *string {
	return s.GranularityDisplayName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetGranularityId() *int64 {
	return s.GranularityId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetGuid() *string {
	return s.Guid
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetId() *string {
	return s.Id
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetName() *string {
	return s.Name
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetOwner() *string {
	return s.Owner
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) GetType() *string {
	return s.Type
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetBizUnitId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.BizUnitId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetBizUnitName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.BizUnitName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetCatalog(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Catalog = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetCellSumLogicTableName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.CellSumLogicTableName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetComputeType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.ComputeType = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetDateType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.DateType = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetDescription(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Description = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetDisplayName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.DisplayName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetGranularityDisplayName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.GranularityDisplayName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetGranularityId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.GranularityId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetGuid(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Guid = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetId(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Id = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Name = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetOwner(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Owner = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetOwnerName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.OwnerName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetProjectId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.ProjectId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetProjectName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.ProjectName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) SetType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo {
	s.Type = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListIndexInfo) Validate() error {
	return dara.Validate(s)
}

type ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo struct {
	// example:
	//
	// 1121
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// test
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// 22
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// test
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsPartitionTable *bool   `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// LOGIC_DIM_TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetDescription() *string {
	return s.Description
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetEnv() *string {
	return s.Env
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetId() *string {
	return s.Id
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetName() *string {
	return s.Name
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetOwner() *string {
	return s.Owner
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) GetType() *string {
	return s.Type
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetBizUnitId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.BizUnitId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetBizUnitName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.BizUnitName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetCatalog(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Catalog = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetDataSourceId(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.DataSourceId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetDataSourceType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.DataSourceType = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetDescription(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Description = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetEnv(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Env = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetId(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Id = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetIsPartitionTable(v bool) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.IsPartitionTable = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Name = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetOwner(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Owner = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetOwnerName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.OwnerName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetProjectId(v int64) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.ProjectId = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetProjectName(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.ProjectName = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) SetType(v string) *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo {
	s.Type = &v
	return s
}

func (s *ListQualityWatchesResponseBodyPageResultQualityWatchListTableInfo) Validate() error {
	return dara.Validate(s)
}
