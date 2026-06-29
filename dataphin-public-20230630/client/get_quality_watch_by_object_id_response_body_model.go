// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchByObjectIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityWatchByObjectIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityWatchByObjectIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityWatchByObjectIdResponseBody
	GetMessage() *string
	SetQualityWatchInfo(v *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) *GetQualityWatchByObjectIdResponseBody
	GetQualityWatchInfo() *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo
	SetRequestId(v string) *GetQualityWatchByObjectIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityWatchByObjectIdResponseBody
	GetSuccess() *bool
}

type GetQualityWatchByObjectIdResponseBody struct {
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
	// The error details from the backend.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the monitored object.
	QualityWatchInfo *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo `json:"QualityWatchInfo,omitempty" xml:"QualityWatchInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityWatchByObjectIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityWatchByObjectIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityWatchByObjectIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityWatchByObjectIdResponseBody) GetQualityWatchInfo() *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	return s.QualityWatchInfo
}

func (s *GetQualityWatchByObjectIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityWatchByObjectIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityWatchByObjectIdResponseBody) SetCode(v string) *GetQualityWatchByObjectIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBody) SetHttpStatusCode(v int32) *GetQualityWatchByObjectIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBody) SetMessage(v string) *GetQualityWatchByObjectIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBody) SetQualityWatchInfo(v *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) *GetQualityWatchByObjectIdResponseBody {
	s.QualityWatchInfo = v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBody) SetRequestId(v string) *GetQualityWatchByObjectIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBody) SetSuccess(v bool) *GetQualityWatchByObjectIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBody) Validate() error {
	if s.QualityWatchInfo != nil {
		if err := s.QualityWatchInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityWatchByObjectIdResponseBodyQualityWatchInfo struct {
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
	// The name of the creator.
	//
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The datasource details.
	DataSourceInfo *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo `json:"DataSourceInfo,omitempty" xml:"DataSourceInfo,omitempty" type:"Struct"`
	// The number of enabled rules.
	//
	// example:
	//
	// 1
	EnabledRuleCount *int64 `json:"EnabledRuleCount,omitempty" xml:"EnabledRuleCount,omitempty"`
	// The monitoring ID.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric details.
	IndexInfo *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo `json:"IndexInfo,omitempty" xml:"IndexInfo,omitempty" type:"Struct"`
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
	// The monitoring name.
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
	TableInfo *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
	// The monitored object type. Valid values:
	//
	// - TABLE: Dataphin table.
	//
	// - DATASOURCE_TABLE: global table.
	//
	// - DATASOURCE: datasource.
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

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetDataSourceInfo() *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	return s.DataSourceInfo
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetEnabledRuleCount() *int64 {
	return s.EnabledRuleCount
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetIndexInfo() *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	return s.IndexInfo
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetLatestWatchTaskId() *int64 {
	return s.LatestWatchTaskId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetLatestWatchTaskStatus() *string {
	return s.LatestWatchTaskStatus
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetQualityOwner() *string {
	return s.QualityOwner
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetQualityOwnerName() *string {
	return s.QualityOwnerName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetStatus() *string {
	return s.Status
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetTableInfo() *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	return s.TableInfo
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetCreateTime(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetCreator(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetCreatorName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.CreatorName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetDataSourceInfo(v *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.DataSourceInfo = v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetEnabledRuleCount(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.EnabledRuleCount = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetIndexInfo(v *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.IndexInfo = v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetLatestWatchTaskId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.LatestWatchTaskId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetLatestWatchTaskStatus(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.LatestWatchTaskStatus = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetModifier(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetModifyTime(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetQualityOwner(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.QualityOwner = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetQualityOwnerName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.QualityOwnerName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetRuleCount(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.RuleCount = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetStatus(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.Status = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetTableInfo(v *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.TableInfo = v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) SetType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfo) Validate() error {
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

type GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo struct {
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
	// The name of the creator.
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
	// The datasource ID.
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
	// The datasource name.
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
	// The datasource type. Valid values:
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

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetEnv() *string {
	return s.Env
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetId() *string {
	return s.Id
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetCreateTime(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetCreator(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetCreatorName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetEnv(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.Env = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetId(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetModifyTime(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetOwner(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetOwnerName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) SetType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo struct {
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
	// The name of the cell aggregate table.
	//
	// example:
	//
	// dws_all
	CellSumLogicTableName *string `json:"CellSumLogicTableName,omitempty" xml:"CellSumLogicTableName,omitempty"`
	// The computation type of the metric. Valid values:
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
	// The data type of the metric.
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
	// The display name of the metric.
	//
	// example:
	//
	// logic
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The display name of the statistical granularity.
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

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetCellSumLogicTableName() *string {
	return s.CellSumLogicTableName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetComputeType() *string {
	return s.ComputeType
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetDateType() *string {
	return s.DateType
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetGranularityDisplayName() *string {
	return s.GranularityDisplayName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetGranularityId() *int64 {
	return s.GranularityId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetGuid() *string {
	return s.Guid
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetId() *string {
	return s.Id
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetBizUnitId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetBizUnitName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetCatalog(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Catalog = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetCellSumLogicTableName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.CellSumLogicTableName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetComputeType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.ComputeType = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetDateType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.DateType = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetDescription(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Description = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetDisplayName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.DisplayName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetGranularityDisplayName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.GranularityDisplayName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetGranularityId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.GranularityId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetGuid(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Guid = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetId(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetOwner(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetOwnerName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetProjectId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.ProjectId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetProjectName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.ProjectName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) SetType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoIndexInfo) Validate() error {
	return dara.Validate(s)
}

type GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo struct {
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
	// The datasource ID.
	//
	// example:
	//
	// 22
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The datasource type.
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

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetEnv() *string {
	return s.Env
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetId() *string {
	return s.Id
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetBizUnitId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetBizUnitName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetCatalog(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Catalog = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetDataSourceId(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.DataSourceId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetDataSourceType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.DataSourceType = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetDescription(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Description = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetEnv(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Env = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetId(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetIsPartitionTable(v bool) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.IsPartitionTable = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Name = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetOwner(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetOwnerName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetProjectId(v int64) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.ProjectId = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetProjectName(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.ProjectName = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) SetType(v string) *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo {
	s.Type = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponseBodyQualityWatchInfoTableInfo) Validate() error {
	return dara.Validate(s)
}
