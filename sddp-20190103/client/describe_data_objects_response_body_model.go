// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataObjectsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataObjectsResponseBodyItems) *DescribeDataObjectsResponseBody
	GetItems() []*DescribeDataObjectsResponseBodyItems
	SetPageSize(v int32) *DescribeDataObjectsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataObjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataObjectsResponseBody
	GetTotalCount() *int32
}

type DescribeDataObjectsResponseBody struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of data objects.
	Items []*DescribeDataObjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of data asset instances to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID of the request. Alibaba Cloud generates this ID to help you troubleshoot issues.
	//
	// example:
	//
	// E6F6460E-4330-549A-BD89-C183FB17571E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataObjectsResponseBody) GetItems() []*DescribeDataObjectsResponseBodyItems {
	return s.Items
}

func (s *DescribeDataObjectsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataObjectsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataObjectsResponseBody) SetCurrentPage(v int32) *DescribeDataObjectsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetItems(v []*DescribeDataObjectsResponseBodyItems) *DescribeDataObjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetPageSize(v int32) *DescribeDataObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetRequestId(v string) *DescribeDataObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetTotalCount(v int32) *DescribeDataObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataObjectsResponseBodyItems struct {
	// An array of industry categories to which the sensitive data belongs.
	Categories  []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	ClusterType *string   `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The comment on the column.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The data type of the database column.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// DataBaseName
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The code for the file category.
	//
	// example:
	//
	// 1
	FileCategoryCode *int32 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// The unique ID of the data object.
	//
	// example:
	//
	// 20000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the data asset instance.
	//
	// example:
	//
	// instance description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the data asset instance.
	//
	// example:
	//
	// rm-12*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The revision status.
	//
	// example:
	//
	// -1
	IsRevision *int32 `json:"IsRevision,omitempty" xml:"IsRevision,omitempty"`
	// The last modification time of the file, in milliseconds.
	//
	// example:
	//
	// 1687676649830
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The timestamp of the last scan, in milliseconds.
	//
	// example:
	//
	// 1687676649830
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The name of the Logstore in SLS.
	//
	// example:
	//
	// logStore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The column encryption status.
	//
	// example:
	//
	// -1
	MaskStatus *int32 `json:"MaskStatus,omitempty" xml:"MaskStatus,omitempty"`
	// The ID of the member account.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// A list of data tags.
	ModelTags []*DescribeDataObjectsResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the data object.
	//
	// example:
	//
	// t_sddp_selfmysql_pers0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the file category.
	//
	// example:
	//
	// text file
	ObjectFileCategory *string `json:"ObjectFileCategory,omitempty" xml:"ObjectFileCategory,omitempty"`
	// The type of the data object.
	//
	// example:
	//
	// text type
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The path of the data object.
	//
	// example:
	//
	// rm-12**.db_***
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The name of the product to which the data object belongs. Valid values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADB-MYSQL**
	//
	// - **Table Store**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// - **PolarDB-X**
	//
	// - **PolarDB**
	//
	// - **ADB-PG**
	//
	// - **OceanBase**
	//
	// - **MongoDB**
	//
	// - **Redis**
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the product to which the data object belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: Table Store
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
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the Simple Log Service (SLS) project.
	//
	// example:
	//
	// project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the data object is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The risk level.
	//
	// example:
	//
	// 1
	RiskLevelId *int32 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The number of matched rules.
	//
	// example:
	//
	// 10
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// A list of matched detection models.
	RuleList []*DescribeDataObjectsResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The number of sensitive data fields.
	//
	// example:
	//
	// 1
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The size of the file in bytes.
	//
	// example:
	//
	// 1000
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// A comma-separated string that specifies the count of matched rules for each risk level. The string follows the format `S1,S2...S10`, where the value at each position represents the count for the corresponding risk level.
	//
	// example:
	//
	// 1,2,3,0,0,0,0,5,0,0
	Sx *string `json:"Sx,omitempty" xml:"Sx,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task number.
	//
	// example:
	//
	// 1000
	TaskNumber *int64 `json:"TaskNumber,omitempty" xml:"TaskNumber,omitempty"`
	// The ID of the industry template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeDataObjectsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBodyItems) GetCategories() []*string {
	return s.Categories
}

func (s *DescribeDataObjectsResponseBodyItems) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeDataObjectsResponseBodyItems) GetComment() *string {
	return s.Comment
}

func (s *DescribeDataObjectsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataObjectsResponseBodyItems) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDataObjectsResponseBodyItems) GetFileCategoryCode() *int32 {
	return s.FileCategoryCode
}

func (s *DescribeDataObjectsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeDataObjectsResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeDataObjectsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDataObjectsResponseBodyItems) GetIsRevision() *int32 {
	return s.IsRevision
}

func (s *DescribeDataObjectsResponseBodyItems) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *DescribeDataObjectsResponseBodyItems) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeDataObjectsResponseBodyItems) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeDataObjectsResponseBodyItems) GetMaskStatus() *int32 {
	return s.MaskStatus
}

func (s *DescribeDataObjectsResponseBodyItems) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *DescribeDataObjectsResponseBodyItems) GetModelTags() []*DescribeDataObjectsResponseBodyItemsModelTags {
	return s.ModelTags
}

func (s *DescribeDataObjectsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeDataObjectsResponseBodyItems) GetObjectFileCategory() *string {
	return s.ObjectFileCategory
}

func (s *DescribeDataObjectsResponseBodyItems) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeDataObjectsResponseBodyItems) GetPath() *string {
	return s.Path
}

func (s *DescribeDataObjectsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeDataObjectsResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeDataObjectsResponseBodyItems) GetProject() *string {
	return s.Project
}

func (s *DescribeDataObjectsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataObjectsResponseBodyItems) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeDataObjectsResponseBodyItems) GetRiskLevelId() *int32 {
	return s.RiskLevelId
}

func (s *DescribeDataObjectsResponseBodyItems) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeDataObjectsResponseBodyItems) GetRuleList() []*DescribeDataObjectsResponseBodyItemsRuleList {
	return s.RuleList
}

func (s *DescribeDataObjectsResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *DescribeDataObjectsResponseBodyItems) GetSize() *int64 {
	return s.Size
}

func (s *DescribeDataObjectsResponseBodyItems) GetSx() *string {
	return s.Sx
}

func (s *DescribeDataObjectsResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDataObjectsResponseBodyItems) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeDataObjectsResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeDataObjectsResponseBodyItems) GetTaskNumber() *int64 {
	return s.TaskNumber
}

func (s *DescribeDataObjectsResponseBodyItems) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDataObjectsResponseBodyItems) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDataObjectsResponseBodyItems) SetCategories(v []*string) *DescribeDataObjectsResponseBodyItems {
	s.Categories = v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetClusterType(v string) *DescribeDataObjectsResponseBodyItems {
	s.ClusterType = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetComment(v string) *DescribeDataObjectsResponseBodyItems {
	s.Comment = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetDataType(v string) *DescribeDataObjectsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetDbName(v string) *DescribeDataObjectsResponseBodyItems {
	s.DbName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetFileCategoryCode(v int32) *DescribeDataObjectsResponseBodyItems {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetId(v string) *DescribeDataObjectsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetInstanceDescription(v string) *DescribeDataObjectsResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetInstanceId(v string) *DescribeDataObjectsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetIsRevision(v int32) *DescribeDataObjectsResponseBodyItems {
	s.IsRevision = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetLastModifiedTime(v int64) *DescribeDataObjectsResponseBodyItems {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetLastScanTime(v int64) *DescribeDataObjectsResponseBodyItems {
	s.LastScanTime = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetLogStore(v string) *DescribeDataObjectsResponseBodyItems {
	s.LogStore = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetMaskStatus(v int32) *DescribeDataObjectsResponseBodyItems {
	s.MaskStatus = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetMemberAccount(v int64) *DescribeDataObjectsResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetModelTags(v []*DescribeDataObjectsResponseBodyItemsModelTags) *DescribeDataObjectsResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetName(v string) *DescribeDataObjectsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetObjectFileCategory(v string) *DescribeDataObjectsResponseBodyItems {
	s.ObjectFileCategory = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetObjectType(v string) *DescribeDataObjectsResponseBodyItems {
	s.ObjectType = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetPath(v string) *DescribeDataObjectsResponseBodyItems {
	s.Path = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetProductCode(v string) *DescribeDataObjectsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetProductId(v int64) *DescribeDataObjectsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetProject(v string) *DescribeDataObjectsResponseBodyItems {
	s.Project = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRegionId(v string) *DescribeDataObjectsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRegionName(v string) *DescribeDataObjectsResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRiskLevelId(v int32) *DescribeDataObjectsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRuleCount(v int32) *DescribeDataObjectsResponseBodyItems {
	s.RuleCount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRuleList(v []*DescribeDataObjectsResponseBodyItemsRuleList) *DescribeDataObjectsResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetSensitiveCount(v int32) *DescribeDataObjectsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetSize(v int64) *DescribeDataObjectsResponseBodyItems {
	s.Size = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetSx(v string) *DescribeDataObjectsResponseBodyItems {
	s.Sx = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTableName(v string) *DescribeDataObjectsResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTaskId(v int64) *DescribeDataObjectsResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTaskName(v string) *DescribeDataObjectsResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTaskNumber(v int64) *DescribeDataObjectsResponseBodyItems {
	s.TaskNumber = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTemplateId(v int64) *DescribeDataObjectsResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTemplateName(v string) *DescribeDataObjectsResponseBodyItems {
	s.TemplateName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) Validate() error {
	if s.ModelTags != nil {
		for _, item := range s.ModelTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataObjectsResponseBodyItemsModelTags struct {
	// The ID of the data tag. Valid values:
	//
	// - **101**: Personal sensitive information
	//
	// - **102**: Personal information
	//
	// - **107**: General information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data tag. Valid values:
	//
	// - **Personal sensitive information**
	//
	// - **Personal information**
	//
	// - **General information**
	//
	// example:
	//
	// Personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataObjectsResponseBodyItemsModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectsResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) GetName() *string {
	return s.Name
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) SetId(v int64) *DescribeDataObjectsResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) SetName(v string) *DescribeDataObjectsResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) Validate() error {
	return dara.Validate(s)
}

type DescribeDataObjectsResponseBodyItemsRuleList struct {
	// The ID of the risk level. Valid values:
	//
	// - **1**: N/A - No sensitive data is detected
	//
	// - **2**: S1 - Level-1 sensitive data
	//
	// - **3**: S2 - Level-2 sensitive data
	//
	// - **4**: S3 - Level-3 sensitive data
	//
	// - **5**: S4 - Level-4 sensitive data
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the risk level. Valid values:
	//
	// - **N/A**: No sensitive data is detected
	//
	// - **S1**: Level-1 sensitive data
	//
	// - **S2**: Level-2 sensitive data
	//
	// - **S3**: Level-3 sensitive data
	//
	// - **S4**: Level-4 sensitive data
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The hierarchical category of the rule, from the top-level to the leaf-level category in the template.
	//
	// example:
	//
	// Personal sensitive information-ID card
	RuleCategoryNameList *string `json:"RuleCategoryNameList,omitempty" xml:"RuleCategoryNameList,omitempty"`
	// The number of matched detection models.
	//
	// example:
	//
	// 590
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The ID of the detection model.
	//
	// example:
	//
	// 1080
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the detection model.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The sample data.
	//
	// example:
	//
	// ["Lucy"，"Tom"]
	SampleList *string `json:"SampleList,omitempty" xml:"SampleList,omitempty"`
}

func (s DescribeDataObjectsResponseBodyItemsRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectsResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetRuleCategoryNameList() *string {
	return s.RuleCategoryNameList
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) GetSampleList() *string {
	return s.SampleList
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRiskLevelName(v string) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleCategoryNameList(v string) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleCategoryNameList = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleCount(v int32) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleCount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleId(v int64) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleName(v string) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetSampleList(v string) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.SampleList = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) Validate() error {
	return dara.Validate(s)
}
