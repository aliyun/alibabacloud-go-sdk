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
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCatalogAssetDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 1
	ApiCallMode *string `json:"ApiCallMode,omitempty" xml:"ApiCallMode,omitempty"`
	// example:
	//
	// 默认API分组
	ApiGroupName *string `json:"ApiGroupName,omitempty" xml:"ApiGroupName,omitempty"`
	// example:
	//
	// 10441
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// 1
	ApiRequestMethod *string `json:"ApiRequestMethod,omitempty" xml:"ApiRequestMethod,omitempty"`
	// example:
	//
	// abc
	AssetDescription *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	// example:
	//
	// https://dataphin.poc.lydaas.com/market/catalog/detail/table/...
	AssetDetailUrl *string `json:"AssetDetailUrl,omitempty" xml:"AssetDetailUrl,omitempty"`
	// example:
	//
	// abc表
	AssetDisplayName *string `json:"AssetDisplayName,omitempty" xml:"AssetDisplayName,omitempty"`
	// example:
	//
	// Dataphin-中间层-服饰零售 (LD_Fashion)
	AssetFrom *string `json:"AssetFrom,omitempty" xml:"AssetFrom,omitempty"`
	// example:
	//
	// dwd_all.abc
	AssetFullName *string `json:"AssetFullName,omitempty" xml:"AssetFullName,omitempty"`
	// example:
	//
	// abc
	AssetName *string   `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	AssetTags []*string `json:"AssetTags,omitempty" xml:"AssetTags,omitempty" type:"Repeated"`
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// dataphin演示空间
	BiCatalog *string `json:"BiCatalog,omitempty" xml:"BiCatalog,omitempty"`
	// example:
	//
	// 6865277495315392
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// 服饰零售（LD_Fashion）
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 23
	ChartCount *int64 `json:"ChartCount,omitempty" xml:"ChartCount,omitempty"`
	// example:
	//
	// 0
	CollectionCount *int64                                           `json:"CollectionCount,omitempty" xml:"CollectionCount,omitempty"`
	Columns         []*GetCatalogAssetDetailsResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-10-11 16:10:19
	CreateTime       *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomAttributes []*GetCatalogAssetDetailsResponseBodyDataCustomAttributes `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty" type:"Repeated"`
	// example:
	//
	// 49837403
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// example:
	//
	// 课程域
	DataCellName *string `json:"DataCellName,omitempty" xml:"DataCellName,omitempty"`
	// example:
	//
	// demo_mysql
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// 7305549302863001856
	DatasourceId *int64                                               `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	Directories  []*GetCatalogAssetDetailsResponseBodyDataDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-05-22 10:06:20
	FirstOnShelveTime *string                                                  `json:"FirstOnShelveTime,omitempty" xml:"FirstOnShelveTime,omitempty"`
	FirstOnShelveUser *GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser `json:"FirstOnShelveUser,omitempty" xml:"FirstOnShelveUser,omitempty" type:"Struct"`
	// example:
	//
	// 课程
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// dp_ds_table.300023201.7311626611751680256.load_test.abc
	Guid             *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	Instruction      *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	IsDeleted        *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	IsPartitionTable *bool   `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// example:
	//
	// 2024-10-11 16:10:19
	LastDdlTime *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// example:
	//
	// 2024-10-11 16:10:19
	LastDmlTime *string `json:"LastDmlTime,omitempty" xml:"LastDmlTime,omitempty"`
	// example:
	//
	// 2025-05-22 10:06:20
	LastOnShelveTime   *string                                                 `json:"LastOnShelveTime,omitempty" xml:"LastOnShelveTime,omitempty"`
	LastOnShelveUser   *GetCatalogAssetDetailsResponseBodyDataLastOnShelveUser `json:"LastOnShelveUser,omitempty" xml:"LastOnShelveUser,omitempty" type:"Struct"`
	MaintainUserGroups []*string                                               `json:"MaintainUserGroups,omitempty" xml:"MaintainUserGroups,omitempty" type:"Repeated"`
	MaintainUserIds    []*string                                               `json:"MaintainUserIds,omitempty" xml:"MaintainUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// L3
	MaxSecurityLevel *string `json:"MaxSecurityLevel,omitempty" xml:"MaxSecurityLevel,omitempty"`
	// example:
	//
	// 2024-10-11 16:10:19
	ModifyTime *string                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Owner      *GetCatalogAssetDetailsResponseBodyDataOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// ds
	PartitionKey *string `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	// example:
	//
	// employee_id
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// example:
	//
	// ALL_USERS_CAN_VIEW
	ProfilingReportViewScopeType       *string   `json:"ProfilingReportViewScopeType,omitempty" xml:"ProfilingReportViewScopeType,omitempty"`
	ProfilingReportViewScopeUserGroups []*string `json:"ProfilingReportViewScopeUserGroups,omitempty" xml:"ProfilingReportViewScopeUserGroups,omitempty" type:"Repeated"`
	ProfilingReportViewScopeUserIds    []*string `json:"ProfilingReportViewScopeUserIds,omitempty" xml:"ProfilingReportViewScopeUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 6865331517728384
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// train
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 5
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// example:
	//
	// ALL_USERS_CAN_VIEW
	ShelveViewScopeType       *string                                                  `json:"ShelveViewScopeType,omitempty" xml:"ShelveViewScopeType,omitempty"`
	ShelveViewScopeUserGroups []*string                                                `json:"ShelveViewScopeUserGroups,omitempty" xml:"ShelveViewScopeUserGroups,omitempty" type:"Repeated"`
	ShelveViewScopeUserIds    []*string                                                `json:"ShelveViewScopeUserIds,omitempty" xml:"ShelveViewScopeUserIds,omitempty" type:"Repeated"`
	SimpleNodeInfos           []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos `json:"SimpleNodeInfos,omitempty" xml:"SimpleNodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// DIM_NORMAL
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// odps.300023201.test.ads_gross
	SumTableGuid *string `json:"SumTableGuid,omitempty" xml:"SumTableGuid,omitempty"`
	// example:
	//
	// ads_gross
	SumTableName *string `json:"SumTableName,omitempty" xml:"SumTableName,omitempty"`
	// example:
	//
	// 36000
	TableLifeCycle *string `json:"TableLifeCycle,omitempty" xml:"TableLifeCycle,omitempty"`
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
	AssociatedEntity *GetCatalogAssetDetailsResponseBodyDataColumnsAssociatedEntity `json:"AssociatedEntity,omitempty" xml:"AssociatedEntity,omitempty" type:"Struct"`
	// example:
	//
	// DIMENSION
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// /交易信息/0000001
	ClassifyName *string `json:"ClassifyName,omitempty" xml:"ClassifyName,omitempty"`
	// example:
	//
	// double
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 门店客流转化率
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// dp_table.300023201.ld_fashion.dws_lulu_location.conversion_rate
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// L4
	LevelShortName *string `json:"LevelShortName,omitempty" xml:"LevelShortName,omitempty"`
	// example:
	//
	// conversion_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.0
	QualityScore *float64                                                  `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	Standards    []*GetCatalogAssetDetailsResponseBodyDataColumnsStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
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
	// example:
	//
	// 7137404445633152
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// LD_train
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 上药erp数据源
	DimensionDisplayName *string `json:"DimensionDisplayName,omitempty" xml:"DimensionDisplayName,omitempty"`
	// example:
	//
	// 68014359
	DimensionId *int64 `json:"DimensionId,omitempty" xml:"DimensionId,omitempty"`
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
	// example:
	//
	// hr_person_id
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 120350
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// example:
	//
	// MANAGEMENT
	AttrType *string `json:"AttrType,omitempty" xml:"AttrType,omitempty"`
	// example:
	//
	// gkglbm
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 归口管理部门
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 102260
	DirectoryId *int64 `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 线上电商平台
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// example:
	//
	// 101676
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// example:
	//
	// 全渠道数据专题
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetCatalogAssetDetailsResponseBodyDataDirectories) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponseBodyDataDirectories) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetDirectoryId() *int64 {
	return s.DirectoryId
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) GetTopicName() *string {
	return s.TopicName
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetDirectoryId(v int64) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.DirectoryId = &v
	return s
}

func (s *GetCatalogAssetDetailsResponseBodyDataDirectories) SetDirectoryName(v string) *GetCatalogAssetDetailsResponseBodyDataDirectories {
	s.DirectoryName = &v
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
	return dara.Validate(s)
}

type GetCatalogAssetDetailsResponseBodyDataFirstOnShelveUser struct {
	// example:
	//
	// 张三
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	// example:
	//
	// 张三
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	// example:
	//
	// 张三
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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

type GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfos struct {
	BizUnit *string `json:"BizUnit,omitempty" xml:"BizUnit,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// n_7443633109495119872
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 2345
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// NORMAL
	NodeScheduleType *string                                                        `json:"NodeScheduleType,omitempty" xml:"NodeScheduleType,omitempty"`
	Owners           []*GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	Project          *GetCatalogAssetDetailsResponseBodyDataSimpleNodeInfosProject  `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
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
	// example:
	//
	// 张三
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	// example:
	//
	// 6865331517728384
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
