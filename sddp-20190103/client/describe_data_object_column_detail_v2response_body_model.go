// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectColumnDetailV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataObjectColumnDetailV2ResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataObjectColumnDetailV2ResponseBodyItems) *DescribeDataObjectColumnDetailV2ResponseBody
	GetItems() []*DescribeDataObjectColumnDetailV2ResponseBodyItems
	SetPageSize(v int32) *DescribeDataObjectColumnDetailV2ResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataObjectColumnDetailV2ResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataObjectColumnDetailV2ResponseBody
	GetTotalCount() *int32
}

type DescribeDataObjectColumnDetailV2ResponseBody struct {
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of detection results for the columns in the data table.
	Items []*DescribeDataObjectColumnDetailV2ResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 231
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) GetItems() []*DescribeDataObjectColumnDetailV2ResponseBodyItems {
	return s.Items
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetItems(v []*DescribeDataObjectColumnDetailV2ResponseBodyItems) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetPageSize(v int32) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetRequestId(v string) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetTotalCount(v int32) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) Validate() error {
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

type DescribeDataObjectColumnDetailV2ResponseBodyItems struct {
	// The list of industry-specific categories for the sensitive data.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The comments on the column.
	//
	// example:
	//
	// column comment
	ColumnComment *string `json:"ColumnComment,omitempty" xml:"ColumnComment,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// hide14
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The engine type. Valid values:
	//
	// - **MySQL**.
	//
	// - **MariaDB**.
	//
	// - **Oracle**.
	//
	// - **PostgreSQL**.
	//
	// - **SQLServer**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The unique ID of the data object.
	//
	// example:
	//
	// 1392973973691383808
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the instance for the data asset table.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The encryption status of the column. Valid values:
	//
	// - **-1**: Not encrypted
	//
	// - **1**: Encrypted
	//
	// - **2**: Encryption failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// The list of data tags.
	ModelTags []*DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// Indicates whether the column is a primary key. Valid values:
	//
	// - **true**: The column is a primary key.
	//
	// - **false**: The column is not a primary key.
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The ID of the product to which the data object belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: Tablestore
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
	// The region where the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the sensitivity level. Valid values:
	//
	// - **1**: N/A: No sensitive data is detected.
	//
	// - **2**: S1: level-1 sensitive data.
	//
	// - **3**: S2: level-2 sensitive data.
	//
	// - **4**: S3: level-3 sensitive data.
	//
	// - **5**: S4: level-4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level. Valid values:
	//
	// - **N/A**: No sensitive data is detected.
	//
	// - **S1**: level-1 sensitive data.
	//
	// - **S2**: level-2 sensitive data.
	//
	// - **S3**: level-3 sensitive data.
	//
	// - **S4**: level-4 sensitive data.
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the hit detection model.
	//
	// example:
	//
	// 51
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the hit detection model.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetCategories() []*string {
	return s.Categories
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetColumnComment() *string {
	return s.ColumnComment
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetMaskingStatus() *int32 {
	return s.MaskingStatus
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetModelTags() []*DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags {
	return s.ModelTags
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetCategories(v []*string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.Categories = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetColumnComment(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ColumnComment = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetColumnName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ColumnName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetDataType(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetEngineType(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetId(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetInstanceName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetMaskingStatus(v int32) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetModelTags(v []*DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetPrimaryKey(v bool) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetProductId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRegionId(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRiskLevelName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRuleId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRuleName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetTableName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) Validate() error {
	if s.ModelTags != nil {
		for _, item := range s.ModelTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags struct {
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
	// - Personal sensitive information
	//
	// - Personal information
	//
	// - General information
	//
	// example:
	//
	// personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) GetName() *string {
	return s.Name
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) SetId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) SetName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) Validate() error {
	return dara.Validate(s)
}
