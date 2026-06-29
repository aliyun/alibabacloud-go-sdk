// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityWatchResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityWatchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityWatchResponseBody
	GetMessage() *string
	SetQualityWatchInfo(v *GetQualityWatchResponseBodyQualityWatchInfo) *GetQualityWatchResponseBody
	GetQualityWatchInfo() *GetQualityWatchResponseBodyQualityWatchInfo
	SetRequestId(v string) *GetQualityWatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityWatchResponseBody
	GetSuccess() *bool
}

type GetQualityWatchResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the monitored object.
	QualityWatchInfo *GetQualityWatchResponseBodyQualityWatchInfo `json:"QualityWatchInfo,omitempty" xml:"QualityWatchInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityWatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityWatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityWatchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityWatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityWatchResponseBody) GetQualityWatchInfo() *GetQualityWatchResponseBodyQualityWatchInfo {
	return s.QualityWatchInfo
}

func (s *GetQualityWatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityWatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityWatchResponseBody) SetCode(v string) *GetQualityWatchResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityWatchResponseBody) SetHttpStatusCode(v int32) *GetQualityWatchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityWatchResponseBody) SetMessage(v string) *GetQualityWatchResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityWatchResponseBody) SetQualityWatchInfo(v *GetQualityWatchResponseBodyQualityWatchInfo) *GetQualityWatchResponseBody {
	s.QualityWatchInfo = v
	return s
}

func (s *GetQualityWatchResponseBody) SetRequestId(v string) *GetQualityWatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityWatchResponseBody) SetSuccess(v bool) *GetQualityWatchResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityWatchResponseBody) Validate() error {
	if s.QualityWatchInfo != nil {
		if err := s.QualityWatchInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityWatchResponseBodyQualityWatchInfo struct {
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creator name.
	//
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The datasource config details.
	DataSourceInfo *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo `json:"DataSourceInfo,omitempty" xml:"DataSourceInfo,omitempty" type:"Struct"`
	// The number of enabled rules.
	//
	// example:
	//
	// 1
	EnabledRuleCount *int64 `json:"EnabledRuleCount,omitempty" xml:"EnabledRuleCount,omitempty"`
	// The monitor ID.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric details.
	IndexInfo *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo `json:"IndexInfo,omitempty" xml:"IndexInfo,omitempty" type:"Struct"`
	// The ID of the most recent quality watchtask for the monitored object.
	//
	// example:
	//
	// 1
	LatestWatchTaskId *int64 `json:"LatestWatchTaskId,omitempty" xml:"LatestWatchTaskId,omitempty"`
	// The status of the most recent quality watchtask for the monitored object.
	//
	// example:
	//
	// SUCCESS
	LatestWatchTaskStatus *string `json:"LatestWatchTaskStatus,omitempty" xml:"LatestWatchTaskStatus,omitempty"`
	// The user ID of the last modifier.
	//
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The monitor name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the quality owner.
	//
	// example:
	//
	// 30012011
	QualityOwner *string `json:"QualityOwner,omitempty" xml:"QualityOwner,omitempty"`
	// The display name of the quality owner.
	//
	// example:
	//
	// test
	QualityOwnerName *string `json:"QualityOwnerName,omitempty" xml:"QualityOwnerName,omitempty"`
	// The number of rules.
	//
	// example:
	//
	// 11
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The status. Valid values:
	//
	// - ENABLE
	//
	// - DISABLE.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The monitored table object.
	TableInfo *GetQualityWatchResponseBodyQualityWatchInfoTableInfo `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
	// The monitored object type. Valid values:
	//
	// - TABLE: Dataphin table.
	//
	// - DATASOURCE_TABLE: full-domain table.
	//
	// - DATASOURCE: data source.
	//
	// - INDEX: metric.
	//
	// - REALTIME_LOGICAL_TABLE: real-time meta table.
	//
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQualityWatchResponseBodyQualityWatchInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchResponseBodyQualityWatchInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetDataSourceInfo() *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	return s.DataSourceInfo
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetEnabledRuleCount() *int64 {
	return s.EnabledRuleCount
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetIndexInfo() *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	return s.IndexInfo
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetLatestWatchTaskId() *int64 {
	return s.LatestWatchTaskId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetLatestWatchTaskStatus() *string {
	return s.LatestWatchTaskStatus
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetQualityOwner() *string {
	return s.QualityOwner
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetQualityOwnerName() *string {
	return s.QualityOwnerName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetStatus() *string {
	return s.Status
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetTableInfo() *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	return s.TableInfo
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetCreateTime(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetCreator(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetCreatorName(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.CreatorName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetDataSourceInfo(v *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.DataSourceInfo = v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetEnabledRuleCount(v int64) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.EnabledRuleCount = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetId(v int64) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetIndexInfo(v *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.IndexInfo = v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetLatestWatchTaskId(v int64) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.LatestWatchTaskId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetLatestWatchTaskStatus(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.LatestWatchTaskStatus = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetModifier(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetModifyTime(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetName(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetQualityOwner(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.QualityOwner = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetQualityOwnerName(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.QualityOwnerName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetRuleCount(v int64) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.RuleCount = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetStatus(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.Status = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetTableInfo(v *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.TableInfo = v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) SetType(v string) *GetQualityWatchResponseBodyQualityWatchInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfo) Validate() error {
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

type GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo struct {
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creator name.
	//
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The environment identifier. Valid values:
	//
	// - PROD
	//
	// - DEV.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The data source name.
	//
	// example:
	//
	// 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner.
	//
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The data source type. Valid values:
	//
	// - MAX_COMPUTE
	//
	// - HADOOP.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetEnv() *string {
	return s.Env
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetId() *string {
	return s.Id
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetCreateTime(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetCreator(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetCreatorName(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetEnv(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.Env = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetId(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetModifyTime(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetName(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetOwner(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetOwnerName(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) SetType(v string) *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type GetQualityWatchResponseBodyQualityWatchInfoIndexInfo struct {
	// The business unit ID.
	//
	// example:
	//
	// 1121
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The business unit name.
	//
	// example:
	//
	// test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The metric catalog.
	//
	// example:
	//
	// test
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The cell aggregate table name.
	//
	// example:
	//
	// dws_all
	CellSumLogicTableName *string `json:"CellSumLogicTableName,omitempty" xml:"CellSumLogicTableName,omitempty"`
	// The metric computation type. Valid values:
	//
	// - AUTO
	//
	// - CUSTOM
	//
	// - MOUNT
	//
	// - COMBINE.
	//
	// example:
	//
	// AUTO
	ComputeType *string `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	// The metric data type.
	//
	// example:
	//
	// bigint
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metric display name.
	//
	// example:
	//
	// logic
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The statistical granularity name.
	//
	// example:
	//
	// 全站汇总表
	GranularityDisplayName *string `json:"GranularityDisplayName,omitempty" xml:"GranularityDisplayName,omitempty"`
	// The statistical granularity ID.
	//
	// example:
	//
	// 18755764
	GranularityId *int64 `json:"GranularityId,omitempty" xml:"GranularityId,omitempty"`
	// The metric GUID.
	//
	// example:
	//
	// 1121
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The metric ID.
	//
	// example:
	//
	// 11
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric name.
	//
	// example:
	//
	// logic
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner.
	//
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The metric type. Valid values:
	//
	// - INDEX.
	//
	// example:
	//
	// INDEX
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetCellSumLogicTableName() *string {
	return s.CellSumLogicTableName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetComputeType() *string {
	return s.ComputeType
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetDateType() *string {
	return s.DateType
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetGranularityDisplayName() *string {
	return s.GranularityDisplayName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetGranularityId() *int64 {
	return s.GranularityId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetGuid() *string {
	return s.Guid
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetId() *string {
	return s.Id
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetBizUnitId(v int64) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetBizUnitName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetCatalog(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Catalog = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetCellSumLogicTableName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.CellSumLogicTableName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetComputeType(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.ComputeType = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetDateType(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.DateType = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetDescription(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Description = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetDisplayName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.DisplayName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetGranularityDisplayName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.GranularityDisplayName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetGranularityId(v int64) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.GranularityId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetGuid(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Guid = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetId(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetOwner(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetOwnerName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetProjectId(v int64) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.ProjectId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetProjectName(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.ProjectName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) SetType(v string) *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoIndexInfo) Validate() error {
	return dara.Validate(s)
}

type GetQualityWatchResponseBodyQualityWatchInfoTableInfo struct {
	// The business unit ID.
	//
	// example:
	//
	// 1121
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The business unit name.
	//
	// example:
	//
	// test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The table catalog.
	//
	// example:
	//
	// test
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 22
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source type.
	//
	// example:
	//
	// MAX_COMPUTE
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment identifier. Valid values:
	//
	// - DEV
	//
	// - PROD.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The table ID.
	//
	// example:
	//
	// test
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the table is a partitioned table.
	IsPartitionTable *bool `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// The table name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner.
	//
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The type. Valid values:
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
	//
	// example:
	//
	// LOGIC_DIM_TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQualityWatchResponseBodyQualityWatchInfoTableInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetEnv() *string {
	return s.Env
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetId() *string {
	return s.Id
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetBizUnitId(v int64) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetBizUnitName(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetCatalog(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Catalog = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetDataSourceId(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.DataSourceId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetDataSourceType(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.DataSourceType = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetDescription(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Description = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetEnv(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Env = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetId(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetIsPartitionTable(v bool) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.IsPartitionTable = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetName(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetOwner(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetOwnerName(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetProjectId(v int64) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.ProjectId = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetProjectName(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.ProjectName = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) SetType(v string) *GetQualityWatchResponseBodyQualityWatchInfoTableInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchResponseBodyQualityWatchInfoTableInfo) Validate() error {
	return dara.Validate(s)
}
