// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIVersion(v int32) *DescribeDataObjectsRequest
	GetAPIVersion() *int32
	SetBucket(v string) *DescribeDataObjectsRequest
	GetBucket() *string
	SetCurrentPage(v int32) *DescribeDataObjectsRequest
	GetCurrentPage() *int32
	SetDbName(v string) *DescribeDataObjectsRequest
	GetDbName() *string
	SetDomainId(v int64) *DescribeDataObjectsRequest
	GetDomainId() *int64
	SetFeatureType(v int32) *DescribeDataObjectsRequest
	GetFeatureType() *int32
	SetFileCategoryCode(v int64) *DescribeDataObjectsRequest
	GetFileCategoryCode() *int64
	SetFileType(v int64) *DescribeDataObjectsRequest
	GetFileType() *int64
	SetInstanceId(v string) *DescribeDataObjectsRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDataObjectsRequest
	GetLang() *string
	SetLogStore(v string) *DescribeDataObjectsRequest
	GetLogStore() *string
	SetLogStoreFlag(v int32) *DescribeDataObjectsRequest
	GetLogStoreFlag() *int32
	SetMemberAccount(v int64) *DescribeDataObjectsRequest
	GetMemberAccount() *int64
	SetModelIds(v string) *DescribeDataObjectsRequest
	GetModelIds() *string
	SetModelTagIds(v string) *DescribeDataObjectsRequest
	GetModelTagIds() *string
	SetPageSize(v int32) *DescribeDataObjectsRequest
	GetPageSize() *int32
	SetParentCategoryIds(v string) *DescribeDataObjectsRequest
	GetParentCategoryIds() *string
	SetPath(v string) *DescribeDataObjectsRequest
	GetPath() *string
	SetProductId(v int32) *DescribeDataObjectsRequest
	GetProductId() *int32
	SetProductIds(v string) *DescribeDataObjectsRequest
	GetProductIds() *string
	SetProject(v string) *DescribeDataObjectsRequest
	GetProject() *string
	SetQueryName(v string) *DescribeDataObjectsRequest
	GetQueryName() *string
	SetRegionId(v string) *DescribeDataObjectsRequest
	GetRegionId() *string
	SetRiskLevelIdList(v string) *DescribeDataObjectsRequest
	GetRiskLevelIdList() *string
	SetRiskLevels(v string) *DescribeDataObjectsRequest
	GetRiskLevels() *string
	SetRuleIds(v string) *DescribeDataObjectsRequest
	GetRuleIds() *string
	SetServiceRegionId(v string) *DescribeDataObjectsRequest
	GetServiceRegionId() *string
	SetTableName(v string) *DescribeDataObjectsRequest
	GetTableName() *string
	SetTaskId(v int64) *DescribeDataObjectsRequest
	GetTaskId() *int64
	SetTemplateId(v int64) *DescribeDataObjectsRequest
	GetTemplateId() *int64
}

type DescribeDataObjectsRequest struct {
	// example:
	//
	// 1
	APIVersion *int32 `json:"APIVersion,omitempty" xml:"APIVersion,omitempty"`
	// example:
	//
	// bucketName
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Page number for the paginated query. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// dataBaseName
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// ID of the data domain to which the data asset belongs.
	//
	// example:
	//
	// 2
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// File category code.
	//
	// example:
	//
	// 1
	FileCategoryCode *int64 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// OSS file types that are supported for recognition.
	//
	// > You can obtain the supported OSS file types by calling [DescribeDocTypes](https://help.aliyun.com/document_detail/2536492.html), using the Code field value from the response. This parameter is only valid for querying OSS-type assets.
	//
	// example:
	//
	// 100001
	FileType *int64 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Keyword for the asset instance ID.
	//
	// example:
	//
	// 8vb54hn2g9j191ddz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language type for request and response messages, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// example:
	//
	// 1
	LogStoreFlag *int32 `json:"LogStoreFlag,omitempty" xml:"LogStoreFlag,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// Model IDs of the industry template, separated by commas.
	//
	// > You can obtain the industry template model identifier ID by calling [DescribeTemplateAllRules](https://help.aliyun.com/document_detail/2536491.html).
	//
	// example:
	//
	// 101
	ModelIds *string `json:"ModelIds,omitempty" xml:"ModelIds,omitempty"`
	// Data labels to be queried, separated by commas. Values:
	//
	// - **101**: Personal Sensitive Information.
	//
	// - **102**: Personal Information.
	//
	// - **107**: General Information.
	//
	// example:
	//
	// 101,102
	ModelTagIds *string `json:"ModelTagIds,omitempty" xml:"ModelTagIds,omitempty"`
	// When performing a paginated query, set the maximum number of data asset instances to display per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of parent category IDs for the template to be queried, separated by commas.
	//
	// example:
	//
	// 234,236,238
	ParentCategoryIds *string `json:"ParentCategoryIds,omitempty" xml:"ParentCategoryIds,omitempty"`
	// example:
	//
	// road
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 5
	ProductId *int32 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// It is recommended to fill in the list of product IDs to be queried, separated by commas. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// > OSS is mutually exclusive with other products, meaning if OSS is included in the query, no other products can be listed; by default, non-OSS products are queried.
	//
	// example:
	//
	// 1,5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// example:
	//
	// project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// Keyword for the data object to be queried.
	//
	// example:
	//
	// t_sddp_selfmysql_pers0
	QueryName *string `json:"QueryName,omitempty" xml:"QueryName,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1,2,3
	RiskLevelIdList *string `json:"RiskLevelIdList,omitempty" xml:"RiskLevelIdList,omitempty"`
	// Specify the risk levels of the data assets to be queried, separated by commas if multiple.
	//
	// - **2**: S1, low risk level.
	//
	// - **3**: S2, medium risk level.
	//
	// - **4**: S3, high risk level.
	//
	// - **5**: S4, highest risk level.
	//
	// example:
	//
	// 2
	RiskLevels *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
	// example:
	//
	// 1,2,3
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// Region where the asset is located. Values:
	//
	// - **cn-beijing**: North China 2 (Beijing).
	//
	// - **cn-zhangjiakou**: North China 3 (Zhangjiakou).
	//
	// - **cn-huhehaote**: North China 5 (Hohhot).
	//
	// - **cn-hangzhou**: East China 1 (Hangzhou).
	//
	// - **cn-shanghai**: East China 2 (Shanghai).
	//
	// - **cn-shenzhen**: South China 1 (Shenzhen).
	//
	// - **cn-hongkong**: Hong Kong, China.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// example:
	//
	// TableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Industry template ID.
	//
	// > You can obtain the industry template identifier ID by calling [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsRequest) GetAPIVersion() *int32 {
	return s.APIVersion
}

func (s *DescribeDataObjectsRequest) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeDataObjectsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataObjectsRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDataObjectsRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *DescribeDataObjectsRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeDataObjectsRequest) GetFileCategoryCode() *int64 {
	return s.FileCategoryCode
}

func (s *DescribeDataObjectsRequest) GetFileType() *int64 {
	return s.FileType
}

func (s *DescribeDataObjectsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDataObjectsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataObjectsRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeDataObjectsRequest) GetLogStoreFlag() *int32 {
	return s.LogStoreFlag
}

func (s *DescribeDataObjectsRequest) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *DescribeDataObjectsRequest) GetModelIds() *string {
	return s.ModelIds
}

func (s *DescribeDataObjectsRequest) GetModelTagIds() *string {
	return s.ModelTagIds
}

func (s *DescribeDataObjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataObjectsRequest) GetParentCategoryIds() *string {
	return s.ParentCategoryIds
}

func (s *DescribeDataObjectsRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeDataObjectsRequest) GetProductId() *int32 {
	return s.ProductId
}

func (s *DescribeDataObjectsRequest) GetProductIds() *string {
	return s.ProductIds
}

func (s *DescribeDataObjectsRequest) GetProject() *string {
	return s.Project
}

func (s *DescribeDataObjectsRequest) GetQueryName() *string {
	return s.QueryName
}

func (s *DescribeDataObjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataObjectsRequest) GetRiskLevelIdList() *string {
	return s.RiskLevelIdList
}

func (s *DescribeDataObjectsRequest) GetRiskLevels() *string {
	return s.RiskLevels
}

func (s *DescribeDataObjectsRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *DescribeDataObjectsRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeDataObjectsRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDataObjectsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeDataObjectsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDataObjectsRequest) SetAPIVersion(v int32) *DescribeDataObjectsRequest {
	s.APIVersion = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetBucket(v string) *DescribeDataObjectsRequest {
	s.Bucket = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetCurrentPage(v int32) *DescribeDataObjectsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetDbName(v string) *DescribeDataObjectsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetDomainId(v int64) *DescribeDataObjectsRequest {
	s.DomainId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetFeatureType(v int32) *DescribeDataObjectsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetFileCategoryCode(v int64) *DescribeDataObjectsRequest {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetFileType(v int64) *DescribeDataObjectsRequest {
	s.FileType = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetInstanceId(v string) *DescribeDataObjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetLang(v string) *DescribeDataObjectsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetLogStore(v string) *DescribeDataObjectsRequest {
	s.LogStore = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetLogStoreFlag(v int32) *DescribeDataObjectsRequest {
	s.LogStoreFlag = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetMemberAccount(v int64) *DescribeDataObjectsRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetModelIds(v string) *DescribeDataObjectsRequest {
	s.ModelIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetModelTagIds(v string) *DescribeDataObjectsRequest {
	s.ModelTagIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetPageSize(v int32) *DescribeDataObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetParentCategoryIds(v string) *DescribeDataObjectsRequest {
	s.ParentCategoryIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetPath(v string) *DescribeDataObjectsRequest {
	s.Path = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetProductId(v int32) *DescribeDataObjectsRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetProductIds(v string) *DescribeDataObjectsRequest {
	s.ProductIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetProject(v string) *DescribeDataObjectsRequest {
	s.Project = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetQueryName(v string) *DescribeDataObjectsRequest {
	s.QueryName = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetRegionId(v string) *DescribeDataObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetRiskLevelIdList(v string) *DescribeDataObjectsRequest {
	s.RiskLevelIdList = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetRiskLevels(v string) *DescribeDataObjectsRequest {
	s.RiskLevels = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetRuleIds(v string) *DescribeDataObjectsRequest {
	s.RuleIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetServiceRegionId(v string) *DescribeDataObjectsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetTableName(v string) *DescribeDataObjectsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetTaskId(v int64) *DescribeDataObjectsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetTemplateId(v int64) *DescribeDataObjectsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataObjectsRequest) Validate() error {
	return dara.Validate(s)
}
