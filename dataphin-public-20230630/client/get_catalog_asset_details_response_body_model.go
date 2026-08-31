// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogAssetDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCatalogAssetDetailsResponseBody
	GetCode() *string
	SetData(v *GetCatalogAssetDetailsResponseBodyData) *GetCatalogAssetDetailsResponseBody
	GetData() *GetCatalogAssetDetailsResponseBodyData
	SetHttpStatusCode(v int32) *GetCatalogAssetDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCatalogAssetDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCatalogAssetDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCatalogAssetDetailsResponseBody
	GetSuccess() *bool
}

type GetCatalogAssetDetailsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data catalog asset details.
	Data *GetCatalogAssetDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend response exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCatalogAssetDetailsResponseBody) GetData() *GetCatalogAssetDetailsResponseBodyData {
	return s.Data
}

func (s *GetCatalogAssetDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCatalogAssetDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCatalogAssetDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCatalogAssetDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCatalogAssetDetailsResponseBody) SetCode(v string) *GetCatalogAssetDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBody) SetData(v *GetCatalogAssetDetailsResponseBodyData) *GetCatalogAssetDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBody) SetHttpStatusCode(v int32) *GetCatalogAssetDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBody) SetMessage(v string) *GetCatalogAssetDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBody) SetRequestId(v string) *GetCatalogAssetDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBody) SetSuccess(v bool) *GetCatalogAssetDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCatalogAssetDetailsResponseBodyData struct {
	// The API call mode. Returned when the asset type is API. Valid values: 1=Synchronous call, 2=Asynchronous call.
	//
	// example:
	//
	// 1
	ApiCallMode *string `json:"ApiCallMode,omitempty" xml:"ApiCallMode,omitempty"`
	// The API group name. Returned when the asset type is API.
	//
	// example:
	//
	// Default API group
	ApiGroupName *string `json:"ApiGroupName,omitempty" xml:"ApiGroupName,omitempty"`
	// The API ID. Returned when the asset type is API.
	//
	// example:
	//
	// 10441
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API operation type. Returned when the asset type is API. Valid values: 1=Get, 2=List, 3=Create, 4=Update, 5=Delete.
	//
	// example:
	//
	// 1
	ApiRequestMethod *string `json:"ApiRequestMethod,omitempty" xml:"ApiRequestMethod,omitempty"`
	// The description of the asset.
	//
	// example:
	//
	// abc
	AssetDescription *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	// The URL of the asset catalog detail page.
	//
	// example:
	//
	// https://dataphin.poc.lydaas.com/market/catalog/detail/table/...
	AssetDetailUrl *string `json:"AssetDetailUrl,omitempty" xml:"AssetDetailUrl,omitempty"`
	// The display name of the asset. This parameter is returned when the asset type is TABLE, INDEX, or BIZ_INDEX.
	//
	// example:
	//
	// abc表
	AssetDisplayName *string `json:"AssetDisplayName,omitempty" xml:"AssetDisplayName,omitempty"`
	// The source of the asset. TABLE (physical table) returns "Dataphin-workspace type-project Chinese name (project English name)". TABLE (logical table) returns "Dataphin-workspace type-data domain Chinese name (data domain English name)". TABLE (data source table) returns "source system name-data source name-database/schema name". INDEX (standard modeling metric) returns the asset source of the associated aggregate logical table. INDEX (custom metric) returns the asset source of the source table. API returns "data service project name". PAGE returns "application system name".
	//
	// example:
	//
	// Dataphin-中间层-服饰零售 (LD_Fashion)
	AssetFrom *string `json:"AssetFrom,omitempty" xml:"AssetFrom,omitempty"`
	// The full name of the asset. This parameter is returned when the asset type is TABLE or INDEX.
	//
	// example:
	//
	// dwd_all.abc
	AssetFullName *string `json:"AssetFullName,omitempty" xml:"AssetFullName,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// abc
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The tags of the asset.
	AssetTags []*string `json:"AssetTags,omitempty" xml:"AssetTags,omitempty" type:"Repeated"`
	// The asset type. Valid values:
	//
	// - TABLE: table, including views and materialized views.
	//
	// - INDEX: technical metric.
	//
	// - BIZ_INDEX: business metric.
	//
	// - API: API.
	//
	// - PAGE: dashboard.
	//
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The BI workspace or folder to which the asset belongs. Returned when the asset type is PAGE (dashboard).
	//
	// example:
	//
	// dataphin演示空间
	BiCatalog *string `json:"BiCatalog,omitempty" xml:"BiCatalog,omitempty"`
	// The ID of the data domain to which the asset belongs. This parameter is returned when the asset type is TABLE (logical tables only) or INDEX (technical metrics whose source table is a logical table only).
	//
	// example:
	//
	// 6865277495315392
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The name of the data domain to which the asset belongs. This parameter is returned when the asset type is TABLE (logical tables only) or INDEX (technical metrics whose source table is a logical table only).
	//
	// example:
	//
	// 服饰零售（LD_Fashion）
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The total number of charts. Returned when the asset type is PAGE (dashboard).
	//
	// example:
	//
	// 23
	ChartCount *int64 `json:"ChartCount,omitempty" xml:"ChartCount,omitempty"`
	// The collection count.
	//
	// example:
	//
	// 0
	CollectionCount *int64 `json:"CollectionCount,omitempty" xml:"CollectionCount,omitempty"`
	// The list of columns. This parameter is returned when the asset type is TABLE.
	Columns []*GetCatalogAssetDetailsResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-11 16:10:19
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom attributes. Returned when includeDetailedAttributes is set to true.
	CustomAttributes []*GetCatalogAssetDetailsResponseBodyDataCustomAttributes `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty" type:"Repeated"`
	// The ID of the data domain. Returned when the asset type is TABLE (logical tables only) or INDEX (technical metrics whose source table is a logical table only).
	//
	// example:
	//
	// 49837403
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// The name of the data domain. Returned when the asset type is TABLE (logical tables only) or INDEX (technical metrics whose source table is a logical table only).
	//
	// example:
	//
	// Course domain
	DataCellName *string `json:"DataCellName,omitempty" xml:"DataCellName,omitempty"`
	// The name of the data source to which the asset belongs. This parameter is returned when the asset type is TABLE (data source tables only) or INDEX (technical metrics whose source table is a data source table only).
	//
	// example:
	//
	// demo_mysql
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The ID of the data source to which the asset belongs. This parameter is returned when the asset type is TABLE (data source tables only) or INDEX (technical metrics whose source table is a data source table only).
	//
	// example:
	//
	// 7305549302863001856
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The directories to which the asset belongs, including topic ID, topic name, directory ID, and directory name.
	Directories []*GetCatalogAssetDetailsResponseBodyDataDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The time of the first listing.
	//
	// example:
	//
	// 2025-05-22 10:06:20
	FirstOnShelveTime *string `json:"FirstOnShelveTime,omitempty" xml:"FirstOnShelveTime,omitempty"`
	// The user who performed the first listing.
	FirstOnShelveUser *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser `json:"FirstOnShelveUser,omitempty" xml:"FirstOnShelveUser,omitempty" type:"Struct"`
	// The statistical granularity name of the technical metric. Returned when the asset type is INDEX.
	//
	// example:
	//
	// Course
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The GUID of the asset, which serves as the unique identifier of the asset.
	//
	// example:
	//
	// dp_ds_table.300023201.7311626611751680256.load_test.abc
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The usage instructions.
	//
	// example:
	//
	// test
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// Indicates whether the asset is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// Indicates whether the table is a partitioned table. Returned when the asset type is TABLE. Valid values:
	//
	// - true: The table is a partitioned table.
	//
	// - false: The table is not a partitioned table.
	IsPartitionTable *bool `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// The time of the last DDL change.
	//
	// example:
	//
	// 2024-10-11 16:10:19
	LastDdlTime *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// The time of the last DML update.
	//
	// example:
	//
	// 2024-10-11 16:10:19
	LastDmlTime *string `json:"LastDmlTime,omitempty" xml:"LastDmlTime,omitempty"`
	// The time of the last listing.
	//
	// example:
	//
	// 2025-05-22 10:06:20
	LastOnShelveTime *string `json:"LastOnShelveTime,omitempty" xml:"LastOnShelveTime,omitempty"`
	// The user who performed the last listing.
	LastOnShelveUser *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser `json:"LastOnShelveUser,omitempty" xml:"LastOnShelveUser,omitempty" type:"Struct"`
	// The listing maintenance user groups.
	MaintainUserGroups []*string `json:"MaintainUserGroups,omitempty" xml:"MaintainUserGroups,omitempty" type:"Repeated"`
	// The IDs of the listing maintenance users.
	MaintainUserIds []*string `json:"MaintainUserIds,omitempty" xml:"MaintainUserIds,omitempty" type:"Repeated"`
	// The maximum sensitivity level. This parameter is returned when the asset type is TABLE.
	//
	// example:
	//
	// L3
	MaxSecurityLevel *string `json:"MaxSecurityLevel,omitempty" xml:"MaxSecurityLevel,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-10-11 16:10:19
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The owner.
	Owner *GetCatalogAssetDetailsResponseBodyDataOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// The partition key. Returned when the asset type is TABLE.
	//
	// example:
	//
	// ds
	PartitionKey *string `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	// The primary key. Returned when the asset type is TABLE.
	//
	// example:
	//
	// employee_id
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The visibility scope type of the profiling report. This parameter is returned only when the asset type is TABLE or INDEX. Valid values:
	//
	// - ALL_USERS_CAN_VIEW: Visible to all users.
	//
	// - PART_USERS_CAN_VIEW: Visible to some users.
	//
	// - ALL_USERS_CAN_NOT_VIEW: Not visible to any users.
	//
	// example:
	//
	// ALL_USERS_CAN_VIEW
	ProfilingReportViewScopeType *string `json:"ProfilingReportViewScopeType,omitempty" xml:"ProfilingReportViewScopeType,omitempty"`
	// The user groups within the profiling report visibility scope.
	ProfilingReportViewScopeUserGroups []*string `json:"ProfilingReportViewScopeUserGroups,omitempty" xml:"ProfilingReportViewScopeUserGroups,omitempty" type:"Repeated"`
	// The users within the profiling report visibility scope.
	ProfilingReportViewScopeUserIds []*string `json:"ProfilingReportViewScopeUserIds,omitempty" xml:"ProfilingReportViewScopeUserIds,omitempty" type:"Repeated"`
	// The ID of the project to which the asset belongs. This parameter is returned when the asset type is TABLE (physical tables only) or INDEX (technical metrics whose source table is a physical table only).
	//
	// example:
	//
	// 6865331517728384
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the project to which the asset belongs. This parameter is returned when the asset type is TABLE (physical tables only) or INDEX (technical metrics whose source table is a physical table only).
	//
	// example:
	//
	// train
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The quality score radar chart information. This parameter is returned only when includeDetailedAttributes is set to true. It contains the total score, the number of passed/validated rules, and the score details for each dimension.
	QualityScoreRadar *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar `json:"QualityScoreRadar,omitempty" xml:"QualityScoreRadar,omitempty" type:"Struct"`
	// The view count.
	//
	// example:
	//
	// 5
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The visibility scope type. Valid values:
	//
	// - ALL_USERS_CAN_VIEW: Visible to all users.
	//
	// - PART_USERS_CAN_VIEW: Visible to some users.
	//
	// - PART_USERS_CAN_NOT_VIEW: Not visible to some users.
	//
	// example:
	//
	// ALL_USERS_CAN_VIEW
	ShelveViewScopeType *string `json:"ShelveViewScopeType,omitempty" xml:"ShelveViewScopeType,omitempty"`
	// The user groups within the visibility scope.
	ShelveViewScopeUserGroups []*string `json:"ShelveViewScopeUserGroups,omitempty" xml:"ShelveViewScopeUserGroups,omitempty" type:"Repeated"`
	// The users within the visibility scope.
	ShelveViewScopeUserIds []*string `json:"ShelveViewScopeUserIds,omitempty" xml:"ShelveViewScopeUserIds,omitempty" type:"Repeated"`
	// The output nodes. Returned when the asset type is TABLE.
	SimpleNodeInfos []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos `json:"SimpleNodeInfos,omitempty" xml:"SimpleNodeInfos,omitempty" type:"Repeated"`
	// The subtype. Valid values:
	//
	// - DIM_NORMAL: common logical dimension table.
	//
	// - DIM_ENUM: enumeration logical dimension table.
	//
	// - DIM_VIRTUAL: virtual logical dimension table.
	//
	// - SUM_BIZ_UNIT: aggregate logical table.
	//
	// - FACT_EVENT: event fact logical table.
	//
	// - FACT_SNAPSHOT: snapshot fact logical table.
	//
	// - DATASOURCE_TABLE: data source table.
	//
	// - PHYSICAL_TABLE: physical table.
	//
	// - DATASOURCE_VIEW: view (data source view).
	//
	// - PHYSICAL_VIEW: physical view.
	//
	// - MATERIALIZED_VIEW: materialized view.
	//
	// - BIZ_INDEX: business metric.
	//
	// - INDEX: technical metric (standard modeling metric).
	//
	// - CUSTOM_INDEX: technical metric (custom metric).
	//
	// example:
	//
	// DIM_NORMAL
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The GUID of the aggregate table to which the asset belongs. Returned when the asset type is INDEX.
	//
	// example:
	//
	// odps.300023201.test.ads_gross
	SumTableGuid *string `json:"SumTableGuid,omitempty" xml:"SumTableGuid,omitempty"`
	// The name of the aggregate table to which the asset belongs. Returned when the asset type is INDEX.
	//
	// example:
	//
	// ads_gross
	SumTableName *string `json:"SumTableName,omitempty" xml:"SumTableName,omitempty"`
	// The lifecycle. Returned when the asset type is TABLE.
	//
	// example:
	//
	// 36000
	TableLifeCycle *string `json:"TableLifeCycle,omitempty" xml:"TableLifeCycle,omitempty"`
	// The storage size. This parameter is returned only when the asset type is TABLE.
	//
	// example:
	//
	// 0
	TableSizeInBytes *int64 `json:"TableSizeInBytes,omitempty" xml:"TableSizeInBytes,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetApiCallMode() *string {
	return s.ApiCallMode
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetApiGroupName() *string {
	return s.ApiGroupName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetApiRequestMethod() *string {
	return s.ApiRequestMethod
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetDetailUrl() *string {
	return s.AssetDetailUrl
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetDisplayName() *string {
	return s.AssetDisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetFrom() *string {
	return s.AssetFrom
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetFullName() *string {
	return s.AssetFullName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetName() *string {
	return s.AssetName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetTags() []*string {
	return s.AssetTags
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetAssetType() *string {
	return s.AssetType
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetBiCatalog() *string {
	return s.BiCatalog
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetChartCount() *int64 {
	return s.ChartCount
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetCollectionCount() *int64 {
	return s.CollectionCount
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetColumns() []*GetCatalogAssetDetailsResponseBodyDataColumns {
	return s.Columns
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetCustomAttributes() []*GetCatalogAssetDetailsResponseBodyDataCustomAttributes {
	return s.CustomAttributes
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetDataCellId() *string {
	return s.DataCellId
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetDataCellName() *string {
	return s.DataCellName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetDirectories() []*GetCatalogAssetDetailsResponseBodyDataDirectories {
	return s.Directories
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetFirstOnShelveTime() *string {
	return s.FirstOnShelveTime
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetFirstOnShelveUser() *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser {
	return s.FirstOnShelveUser
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetGranularity() *string {
	return s.Granularity
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetGuid() *string {
	return s.Guid
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetLastDmlTime() *string {
	return s.LastDmlTime
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetLastOnShelveTime() *string {
	return s.LastOnShelveTime
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetLastOnShelveUser() *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser {
	return s.LastOnShelveUser
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetMaintainUserGroups() []*string {
	return s.MaintainUserGroups
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetMaintainUserIds() []*string {
	return s.MaintainUserIds
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetMaxSecurityLevel() *string {
	return s.MaxSecurityLevel
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetOwner() *GetCatalogAssetDetailsResponseBodyDataOwner {
	return s.Owner
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetPartitionKey() *string {
	return s.PartitionKey
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetProfilingReportViewScopeType() *string {
	return s.ProfilingReportViewScopeType
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetProfilingReportViewScopeUserGroups() []*string {
	return s.ProfilingReportViewScopeUserGroups
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetProfilingReportViewScopeUserIds() []*string {
	return s.ProfilingReportViewScopeUserIds
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetQualityScoreRadar() *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar {
	return s.QualityScoreRadar
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetReadCount() *int64 {
	return s.ReadCount
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetShelveViewScopeType() *string {
	return s.ShelveViewScopeType
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetShelveViewScopeUserGroups() []*string {
	return s.ShelveViewScopeUserGroups
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetShelveViewScopeUserIds() []*string {
	return s.ShelveViewScopeUserIds
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetSimpleNodeInfos() []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	return s.SimpleNodeInfos
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetSubType() *string {
	return s.SubType
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetSumTableGuid() *string {
	return s.SumTableGuid
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetSumTableName() *string {
	return s.SumTableName
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetTableLifeCycle() *string {
	return s.TableLifeCycle
}

func (s *GetCatalogAssetDetailsResponseBodyData) GetTableSizeInBytes() *int64 {
	return s.TableSizeInBytes
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetApiCallMode(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ApiCallMode = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetApiGroupName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ApiGroupName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetApiId(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetApiRequestMethod(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ApiRequestMethod = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetDescription(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetDescription = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetDetailUrl(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetDetailUrl = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetDisplayName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetDisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetFrom(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetFrom = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetFullName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetFullName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetTags(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetTags = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetAssetType(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.AssetType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetBiCatalog(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.BiCatalog = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetBizUnitId(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.BizUnitId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetBizUnitName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.BizUnitName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetChartCount(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.ChartCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetCollectionCount(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.CollectionCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetColumns(v []*GetCatalogAssetDetailsResponseBodyDataColumns) *GetCatalogAssetDetailsResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetCreateTime(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetCustomAttributes(v []*GetCatalogAssetDetailsResponseBodyDataCustomAttributes) *GetCatalogAssetDetailsResponseBodyData {
	s.CustomAttributes = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetDataCellId(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.DataCellId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetDataCellName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.DataCellName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetDataSourceName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetDatasourceId(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.DatasourceId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetDirectories(v []*GetCatalogAssetDetailsResponseBodyDataDirectories) *GetCatalogAssetDetailsResponseBodyData {
	s.Directories = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetFirstOnShelveTime(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.FirstOnShelveTime = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetFirstOnShelveUser(v *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) *GetCatalogAssetDetailsResponseBodyData {
	s.FirstOnShelveUser = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetGranularity(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.Granularity = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetGuid(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetInstruction(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetIsDeleted(v bool) *GetCatalogAssetDetailsResponseBodyData {
	s.IsDeleted = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetIsPartitionTable(v bool) *GetCatalogAssetDetailsResponseBodyData {
	s.IsPartitionTable = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetLastDdlTime(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetLastDmlTime(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.LastDmlTime = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetLastOnShelveTime(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.LastOnShelveTime = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetLastOnShelveUser(v *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) *GetCatalogAssetDetailsResponseBodyData {
	s.LastOnShelveUser = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetMaintainUserGroups(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.MaintainUserGroups = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetMaintainUserIds(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.MaintainUserIds = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetMaxSecurityLevel(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.MaxSecurityLevel = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetModifyTime(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetOwner(v *GetCatalogAssetDetailsResponseBodyDataOwner) *GetCatalogAssetDetailsResponseBodyData {
	s.Owner = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetPartitionKey(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.PartitionKey = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetPrimaryKey(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.PrimaryKey = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetProfilingReportViewScopeType(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ProfilingReportViewScopeType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetProfilingReportViewScopeUserGroups(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.ProfilingReportViewScopeUserGroups = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetProfilingReportViewScopeUserIds(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.ProfilingReportViewScopeUserIds = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetProjectId(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetProjectName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetQualityScoreRadar(v *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) *GetCatalogAssetDetailsResponseBodyData {
	s.QualityScoreRadar = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetReadCount(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.ReadCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetShelveViewScopeType(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.ShelveViewScopeType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetShelveViewScopeUserGroups(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.ShelveViewScopeUserGroups = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetShelveViewScopeUserIds(v []*string) *GetCatalogAssetDetailsResponseBodyData {
	s.ShelveViewScopeUserIds = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetSimpleNodeInfos(v []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) *GetCatalogAssetDetailsResponseBodyData {
	s.SimpleNodeInfos = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetSubType(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.SubType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetSumTableGuid(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.SumTableGuid = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetSumTableName(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.SumTableName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetTableLifeCycle(v string) *GetCatalogAssetDetailsResponseBodyData {
	s.TableLifeCycle = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) SetTableSizeInBytes(v int64) *GetCatalogAssetDetailsResponseBodyData {
	s.TableSizeInBytes = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyData) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomAttributes != nil {
		for _, item := range s.CustomAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Directories != nil {
		for _, item := range s.Directories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FirstOnShelveUser != nil {
		if err := s.FirstOnShelveUser.Validate(); err != nil {
			return err
		}
	}
	if s.LastOnShelveUser != nil {
		if err := s.LastOnShelveUser.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	if s.QualityScoreRadar != nil {
		if err := s.QualityScoreRadar.Validate(); err != nil {
			return err
		}
	}
	if s.SimpleNodeInfos != nil {
		for _, item := range s.SimpleNodeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCatalogAssetDetailsResponseBodyDataColumns struct {
	// The associated entity. This parameter is returned when the business type is DIMENSION.
	AssociatedEntity *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity `json:"AssociatedEntity,omitempty" xml:"AssociatedEntity,omitempty" type:"Struct"`
	// The business type. Valid values:
	//
	// - INDEX: metric.
	//
	// - STAT_PERIOD: statistical period.
	//
	// - DIMENSION: dimension.
	//
	// example:
	//
	// DIMENSION
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The data classification.
	//
	// example:
	//
	// /交易信息/0000001
	ClassifyName *string `json:"ClassifyName,omitempty" xml:"ClassifyName,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// double
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The description of the column.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the column.
	//
	// example:
	//
	// Store traffic conversion rate
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The GUID of the column.
	//
	// example:
	//
	// dp_table.300023201.ld_fashion.dws_lulu_location.conversion_rate
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The data classification level.
	//
	// example:
	//
	// L4
	LevelShortName *string `json:"LevelShortName,omitempty" xml:"LevelShortName,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// conversion_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The quality score.
	//
	// example:
	//
	// 0.0
	QualityScore *float64 `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	// The associated standards.
	Standards []*GetCatalogAssetDetailsResponseBodyDataColumnsStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
}

func (s GetCatalogAssetDetailsResponseBodyDataColumns) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetAssociatedEntity() *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity {
	return s.AssociatedEntity
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetBizType() *string {
	return s.BizType
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetClassifyName() *string {
	return s.ClassifyName
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetDataType() *string {
	return s.DataType
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetDescription() *string {
	return s.Description
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetGuid() *string {
	return s.Guid
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetLevelShortName() *string {
	return s.LevelShortName
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetName() *string {
	return s.Name
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetQualityScore() *float64 {
	return s.QualityScore
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) GetStandards() []*GetCatalogAssetDetailsResponseBodyDataColumnsStandards {
	return s.Standards
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetAssociatedEntity(v *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.AssociatedEntity = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetBizType(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.BizType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetClassifyName(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.ClassifyName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetDataType(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.DataType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetDescription(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.Description = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetDisplayName(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.DisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetGuid(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.Guid = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetLevelShortName(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.LevelShortName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetName(v string) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.Name = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetQualityScore(v float64) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.QualityScore = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) SetStandards(v []*GetCatalogAssetDetailsResponseBodyDataColumnsStandards) *GetCatalogAssetDetailsResponseBodyDataColumns {
	s.Standards = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumns) Validate() error {
	if s.AssociatedEntity != nil {
		if err := s.AssociatedEntity.Validate(); err != nil {
			return err
		}
	}
	if s.Standards != nil {
		for _, item := range s.Standards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity struct {
	// The ID of the business unit.
	//
	// example:
	//
	// 7137404445633152
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The name of the business unit.
	//
	// example:
	//
	// LD_train
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The display name of the dimension.
	//
	// example:
	//
	// 上药erp数据源
	DimensionDisplayName *string `json:"DimensionDisplayName,omitempty" xml:"DimensionDisplayName,omitempty"`
	// The ID of the dimension.
	//
	// example:
	//
	// 68014359
	DimensionId *int64 `json:"DimensionId,omitempty" xml:"DimensionId,omitempty"`
	// The name of the dimension.
	//
	// example:
	//
	// etl_source
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) GetDimensionDisplayName() *string {
	return s.DimensionDisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) GetDimensionId() *int64 {
	return s.DimensionId
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) GetDimensionName() *string {
	return s.DimensionName
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) SetBizUnitId(v int64) *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity {
	s.BizUnitId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) SetBizUnitName(v string) *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity {
	s.BizUnitName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) SetDimensionDisplayName(v string) *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity {
	s.DimensionDisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) SetDimensionId(v int64) *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity {
	s.DimensionId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) SetDimensionName(v string) *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity {
	s.DimensionName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataColumnsStandards struct {
	// The code of the standard.
	//
	// example:
	//
	// hr_person_id
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the standard.
	//
	// example:
	//
	// 120350
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the standard.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataColumnsStandards) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataColumnsStandards) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) GetCode() *string {
	return s.Code
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) GetId() *int64 {
	return s.Id
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) GetName() *string {
	return s.Name
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) SetCode(v string) *GetCatalogAssetDetailsResponseBodyDataColumnsStandards {
	s.Code = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) SetId(v int64) *GetCatalogAssetDetailsResponseBodyDataColumnsStandards {
	s.Id = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) SetName(v string) *GetCatalogAssetDetailsResponseBodyDataColumnsStandards {
	s.Name = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataColumnsStandards) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataCustomAttributes struct {
	// The attribute type. Valid values: BUSINESS (business attribute), MANAGEMENT (management attribute), TECHNOLOGY (technical attribute).
	//
	// example:
	//
	// MANAGEMENT
	AttrType *string `json:"AttrType,omitempty" xml:"AttrType,omitempty"`
	// The attribute code.
	//
	// example:
	//
	// gkglbm
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The attribute name.
	//
	// example:
	//
	// Supervising department
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataCustomAttributes) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataCustomAttributes) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) GetAttrType() *string {
	return s.AttrType
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) GetCode() *string {
	return s.Code
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) GetName() *string {
	return s.Name
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) GetValue() *string {
	return s.Value
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) SetAttrType(v string) *GetCatalogAssetDetailsResponseBodyDataCustomAttributes {
	s.AttrType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) SetCode(v string) *GetCatalogAssetDetailsResponseBodyDataCustomAttributes {
	s.Code = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) SetName(v string) *GetCatalogAssetDetailsResponseBodyDataCustomAttributes {
	s.Name = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) SetValue(v string) *GetCatalogAssetDetailsResponseBodyDataCustomAttributes {
	s.Value = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataCustomAttributes) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataDirectories struct {
	// The complete directory hierarchy chain from the top-level directory to the current directory, including the current directory.
	DirectoryChain []*GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain `json:"DirectoryChain,omitempty" xml:"DirectoryChain,omitempty" type:"Repeated"`
	// The directory description.
	//
	// example:
	//
	// Data tables related to order details
	DirectoryDescription *string `json:"DirectoryDescription,omitempty" xml:"DirectoryDescription,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// 102260
	DirectoryId *int64 `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The directory name.
	//
	// example:
	//
	// Online e-commerce platform
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The topic description.
	//
	// example:
	//
	// Core data asset catalog for the entire company, covering all online and offline channels
	TopicDescription *string `json:"TopicDescription,omitempty" xml:"TopicDescription,omitempty"`
	// The topic ID.
	//
	// example:
	//
	// 101676
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Omni-channel data topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataDirectories) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataDirectories) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetDirectoryChain() []*GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain {
	return s.DirectoryChain
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetDirectoryDescription() *string {
	return s.DirectoryDescription
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetDirectoryId() *int64 {
	return s.DirectoryId
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetTopicDescription() *string {
	return s.TopicDescription
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetTopicName() *string {
	return s.TopicName
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetDirectoryChain(v []*GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.DirectoryChain = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetDirectoryDescription(v string) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.DirectoryDescription = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetDirectoryId(v int64) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.DirectoryId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetDirectoryName(v string) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.DirectoryName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetTopicDescription(v string) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.TopicDescription = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetTopicId(v int64) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.TopicId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetTopicName(v string) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.TopicName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) Validate() error {
	if s.DirectoryChain != nil {
		for _, item := range s.DirectoryChain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain struct {
	// The directory description.
	//
	// example:
	//
	// Data assets related to transactions
	DirectoryDescription *string `json:"DirectoryDescription,omitempty" xml:"DirectoryDescription,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// 2001
	DirectoryId *int64 `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The directory name.
	//
	// example:
	//
	// Transaction domain
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The folder level.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) GetDirectoryDescription() *string {
	return s.DirectoryDescription
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) GetDirectoryId() *int64 {
	return s.DirectoryId
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) GetLevel() *int32 {
	return s.Level
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) SetDirectoryDescription(v string) *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain {
	s.DirectoryDescription = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) SetDirectoryId(v int64) *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain {
	s.DirectoryId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) SetDirectoryName(v string) *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain {
	s.DirectoryName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) SetLevel(v int32) *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain {
	s.Level = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectoriesDirectoryChain) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser struct {
	// The username.
	//
	// example:
	//
	// John
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) GetUserId() *string {
	return s.UserId
}

func (s *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) SetDisplayName(v string) *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser {
	s.DisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) SetUserId(v string) *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser {
	s.UserId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser struct {
	// The username.
	//
	// example:
	//
	// John
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) GetUserId() *string {
	return s.UserId
}

func (s *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) SetDisplayName(v string) *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser {
	s.DisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) SetUserId(v string) *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser {
	s.UserId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataOwner struct {
	// The username.
	//
	// example:
	//
	// John
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataOwner) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataOwner) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataOwner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyDataOwner) GetUserId() *string {
	return s.UserId
}

func (s *GetCatalogAssetDetailsResponseBodyDataOwner) SetDisplayName(v string) *GetCatalogAssetDetailsResponseBodyDataOwner {
	s.DisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataOwner) SetUserId(v string) *GetCatalogAssetDetailsResponseBodyDataOwner {
	s.UserId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataOwner) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar struct {
	// The list of dimension scores.
	CatalogScores []*GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores `json:"CatalogScores,omitempty" xml:"CatalogScores,omitempty" type:"Repeated"`
	// The number of passed rules.
	//
	// example:
	//
	// 10
	PassRuleCount *int32 `json:"PassRuleCount,omitempty" xml:"PassRuleCount,omitempty"`
	// The total quality score.
	//
	// example:
	//
	// 85.5
	TotalScore *float64 `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	// The number of validated rules.
	//
	// example:
	//
	// 12
	ValidateRuleCount *int32 `json:"ValidateRuleCount,omitempty" xml:"ValidateRuleCount,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) GetCatalogScores() []*GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	return s.CatalogScores
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) GetPassRuleCount() *int32 {
	return s.PassRuleCount
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) GetTotalScore() *float64 {
	return s.TotalScore
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) GetValidateRuleCount() *int32 {
	return s.ValidateRuleCount
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) SetCatalogScores(v []*GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar {
	s.CatalogScores = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) SetPassRuleCount(v int32) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar {
	s.PassRuleCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) SetTotalScore(v float64) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar {
	s.TotalScore = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) SetValidateRuleCount(v int32) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar {
	s.ValidateRuleCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadar) Validate() error {
	if s.CatalogScores != nil {
		for _, item := range s.CatalogScores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores struct {
	// The dimension name.
	//
	// example:
	//
	// Completeness
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The number of field-level rules.
	//
	// example:
	//
	// 7
	FieldRuleCount *int32 `json:"FieldRuleCount,omitempty" xml:"FieldRuleCount,omitempty"`
	// The pass rate.
	//
	// example:
	//
	// 0.83
	PassRate *float64 `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
	// The number of passed rules.
	//
	// example:
	//
	// 10
	PassRuleCount *int32 `json:"PassRuleCount,omitempty" xml:"PassRuleCount,omitempty"`
	// The dimension score.
	//
	// example:
	//
	// 85.5
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The number of table-level rules.
	//
	// example:
	//
	// 5
	TableRuleCount *int32 `json:"TableRuleCount,omitempty" xml:"TableRuleCount,omitempty"`
	// The number of validated rules.
	//
	// example:
	//
	// 12
	ValidateRuleCount *int32 `json:"ValidateRuleCount,omitempty" xml:"ValidateRuleCount,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetCatalog() *string {
	return s.Catalog
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetFieldRuleCount() *int32 {
	return s.FieldRuleCount
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetPassRate() *float64 {
	return s.PassRate
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetPassRuleCount() *int32 {
	return s.PassRuleCount
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetScore() *float64 {
	return s.Score
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetTableRuleCount() *int32 {
	return s.TableRuleCount
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) GetValidateRuleCount() *int32 {
	return s.ValidateRuleCount
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetCatalog(v string) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.Catalog = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetFieldRuleCount(v int32) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.FieldRuleCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetPassRate(v float64) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.PassRate = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetPassRuleCount(v int32) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.PassRuleCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetScore(v float64) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.Score = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetTableRuleCount(v int32) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.TableRuleCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) SetValidateRuleCount(v int32) *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores {
	s.ValidateRuleCount = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataQualityScoreRadarCatalogScores) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos struct {
	// The business unit to which the node belongs.
	//
	// example:
	//
	// test
	BizUnit *string `json:"BizUnit,omitempty" xml:"BizUnit,omitempty"`
	// The environment to which the asset belongs.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The node ID.
	//
	// example:
	//
	// n_7443633109495119872
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name.
	//
	// example:
	//
	// 2345
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The scheduling type. Valid values: NORMAL (timed scheduling), MANUAL (manual scheduling).
	//
	// example:
	//
	// NORMAL
	NodeScheduleType *string `json:"NodeScheduleType,omitempty" xml:"NodeScheduleType,omitempty"`
	// The list of O&M owners.
	Owners []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	// The project to which the node belongs.
	Project *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The node type. Example valid values: DLINK (offline integration), PYTHON37 (Python compute node).
	//
	// example:
	//
	// DLINK
	SubBizType *string `json:"SubBizType,omitempty" xml:"SubBizType,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetBizUnit() *string {
	return s.BizUnit
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetEnv() *string {
	return s.Env
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetNodeId() *string {
	return s.NodeId
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetNodeScheduleType() *string {
	return s.NodeScheduleType
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetOwners() []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners {
	return s.Owners
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetProject() *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject {
	return s.Project
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) GetSubBizType() *string {
	return s.SubBizType
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetBizUnit(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.BizUnit = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetEnv(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.Env = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetNodeId(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.NodeId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetNodeName(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.NodeName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetNodeScheduleType(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.NodeScheduleType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetOwners(v []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.Owners = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetProject(v *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.Project = v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) SetSubBizType(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos {
	s.SubBizType = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos) Validate() error {
	if s.Owners != nil {
		for _, item := range s.Owners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners struct {
	// The username.
	//
	// example:
	//
	// John
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) GetUserId() *string {
	return s.UserId
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) SetDisplayName(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners {
	s.DisplayName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) SetUserId(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners {
	s.UserId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners) Validate() error {
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject struct {
	// The project ID.
	//
	// example:
	//
	// 6865331517728384
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// fashion_cdm
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) SetProjectId(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject {
	s.ProjectId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) SetProjectName(v string) *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject {
	s.ProjectName = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject) Validate() error {
	return dara.Validate(s)
}
