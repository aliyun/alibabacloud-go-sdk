// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCatalogAssetsResponseBody
	GetCode() *string
	SetData(v *ListCatalogAssetsResponseBodyData) *ListCatalogAssetsResponseBody
	GetData() *ListCatalogAssetsResponseBodyData
	SetHttpStatusCode(v int32) *ListCatalogAssetsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCatalogAssetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCatalogAssetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCatalogAssetsResponseBody
	GetSuccess() *bool
}

type ListCatalogAssetsResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCatalogAssetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCatalogAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCatalogAssetsResponseBody) GetData() *ListCatalogAssetsResponseBodyData {
	return s.Data
}

func (s *ListCatalogAssetsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCatalogAssetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCatalogAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCatalogAssetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCatalogAssetsResponseBody) SetCode(v string) *ListCatalogAssetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCatalogAssetsResponseBody) SetData(v *ListCatalogAssetsResponseBodyData) *ListCatalogAssetsResponseBody {
	s.Data = v
	return s
}

func (s *ListCatalogAssetsResponseBody) SetHttpStatusCode(v int32) *ListCatalogAssetsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCatalogAssetsResponseBody) SetMessage(v string) *ListCatalogAssetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCatalogAssetsResponseBody) SetRequestId(v string) *ListCatalogAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCatalogAssetsResponseBody) SetSuccess(v bool) *ListCatalogAssetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCatalogAssetsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCatalogAssetsResponseBodyData struct {
	AssetList []*ListCatalogAssetsResponseBodyDataAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCatalogAssetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsResponseBodyData) GetAssetList() []*ListCatalogAssetsResponseBodyDataAssetList {
	return s.AssetList
}

func (s *ListCatalogAssetsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCatalogAssetsResponseBodyData) SetAssetList(v []*ListCatalogAssetsResponseBodyDataAssetList) *ListCatalogAssetsResponseBodyData {
	s.AssetList = v
	return s
}

func (s *ListCatalogAssetsResponseBodyData) SetTotalCount(v int64) *ListCatalogAssetsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyData) Validate() error {
	if s.AssetList != nil {
		for _, item := range s.AssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCatalogAssetsResponseBodyDataAssetList struct {
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
	DatasourceId *int64                                                   `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	Directories  []*ListCatalogAssetsResponseBodyDataAssetListDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// example:
	//
	// 课程
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// dp_ds_table.300023201.7311626611751680256.load_test.abc
	Guid      *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	IsDeleted *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// example:
	//
	// L3
	MaxSecurityLevel *string `json:"MaxSecurityLevel,omitempty" xml:"MaxSecurityLevel,omitempty"`
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
}

func (s ListCatalogAssetsResponseBodyDataAssetList) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsResponseBodyDataAssetList) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetApiCallMode() *string {
	return s.ApiCallMode
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetApiGroupName() *string {
	return s.ApiGroupName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetApiRequestMethod() *string {
	return s.ApiRequestMethod
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetDisplayName() *string {
	return s.AssetDisplayName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetFrom() *string {
	return s.AssetFrom
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetFullName() *string {
	return s.AssetFullName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetName() *string {
	return s.AssetName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetTags() []*string {
	return s.AssetTags
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetAssetType() *string {
	return s.AssetType
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetBiCatalog() *string {
	return s.BiCatalog
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetChartCount() *int64 {
	return s.ChartCount
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetDataCellId() *string {
	return s.DataCellId
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetDataCellName() *string {
	return s.DataCellName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetDirectories() []*ListCatalogAssetsResponseBodyDataAssetListDirectories {
	return s.Directories
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetGranularity() *string {
	return s.Granularity
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetGuid() *string {
	return s.Guid
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetMaxSecurityLevel() *string {
	return s.MaxSecurityLevel
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetSubType() *string {
	return s.SubType
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetSumTableGuid() *string {
	return s.SumTableGuid
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) GetSumTableName() *string {
	return s.SumTableName
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetApiCallMode(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ApiCallMode = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetApiGroupName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ApiGroupName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetApiId(v int64) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ApiId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetApiRequestMethod(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ApiRequestMethod = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetDescription(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetDescription = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetDisplayName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetDisplayName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetFrom(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetFrom = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetFullName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetFullName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetTags(v []*string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetTags = v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetAssetType(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.AssetType = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetBiCatalog(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.BiCatalog = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetBizUnitId(v int64) *ListCatalogAssetsResponseBodyDataAssetList {
	s.BizUnitId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetBizUnitName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.BizUnitName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetChartCount(v int64) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ChartCount = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetDataCellId(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.DataCellId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetDataCellName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.DataCellName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetDataSourceName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.DataSourceName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetDatasourceId(v int64) *ListCatalogAssetsResponseBodyDataAssetList {
	s.DatasourceId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetDirectories(v []*ListCatalogAssetsResponseBodyDataAssetListDirectories) *ListCatalogAssetsResponseBodyDataAssetList {
	s.Directories = v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetGranularity(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.Granularity = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetGuid(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.Guid = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetIsDeleted(v bool) *ListCatalogAssetsResponseBodyDataAssetList {
	s.IsDeleted = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetMaxSecurityLevel(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.MaxSecurityLevel = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetProjectId(v int64) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ProjectId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetProjectName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.ProjectName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetSubType(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.SubType = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetSumTableGuid(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.SumTableGuid = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) SetSumTableName(v string) *ListCatalogAssetsResponseBodyDataAssetList {
	s.SumTableName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetList) Validate() error {
	if s.Directories != nil {
		for _, item := range s.Directories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCatalogAssetsResponseBodyDataAssetListDirectories struct {
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

func (s ListCatalogAssetsResponseBodyDataAssetListDirectories) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsResponseBodyDataAssetListDirectories) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) GetDirectoryId() *int64 {
	return s.DirectoryId
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) GetTopicName() *string {
	return s.TopicName
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) SetDirectoryId(v int64) *ListCatalogAssetsResponseBodyDataAssetListDirectories {
	s.DirectoryId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) SetDirectoryName(v string) *ListCatalogAssetsResponseBodyDataAssetListDirectories {
	s.DirectoryName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) SetTopicId(v int64) *ListCatalogAssetsResponseBodyDataAssetListDirectories {
	s.TopicId = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) SetTopicName(v string) *ListCatalogAssetsResponseBodyDataAssetListDirectories {
	s.TopicName = &v
	return s
}

func (s *ListCatalogAssetsResponseBodyDataAssetListDirectories) Validate() error {
	return dara.Validate(s)
}
