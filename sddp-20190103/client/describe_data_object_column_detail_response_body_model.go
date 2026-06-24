// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectColumnDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataObjectColumnDetailResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataObjectColumnDetailResponseBodyItems) *DescribeDataObjectColumnDetailResponseBody
	GetItems() []*DescribeDataObjectColumnDetailResponseBodyItems
	SetPageSize(v int32) *DescribeDataObjectColumnDetailResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataObjectColumnDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataObjectColumnDetailResponseBody
	GetTotalCount() *int32
}

type DescribeDataObjectColumnDetailResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The details of the columns.
	Items []*DescribeDataObjectColumnDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8C8036CC-961D-514E-88E8-3088B5A50CA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 61
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataObjectColumnDetailResponseBody) GetItems() []*DescribeDataObjectColumnDetailResponseBodyItems {
	return s.Items
}

func (s *DescribeDataObjectColumnDetailResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataObjectColumnDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataObjectColumnDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetItems(v []*DescribeDataObjectColumnDetailResponseBodyItems) *DescribeDataObjectColumnDetailResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetPageSize(v int32) *DescribeDataObjectColumnDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetRequestId(v string) *DescribeDataObjectColumnDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetTotalCount(v int32) *DescribeDataObjectColumnDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) Validate() error {
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

type DescribeDataObjectColumnDetailResponseBodyItems struct {
	// The industry-specific data classifications.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The comment on the column.
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
	// The type of the database engine. Valid values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The ID of the column.
	//
	// example:
	//
	// 1509415150052786176
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the instance where the table is located.
	//
	// example:
	//
	// rm-1234
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The data masking status of the column. Valid values:
	//
	// - **-1**: Not masked
	//
	// - **1**: Masked
	//
	// - **2**: Masking failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// The data labels.
	ModelTags []*DescribeDataObjectColumnDetailResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
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
	// The ID of the service to which the data object belongs. Valid values:
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
	// The ID of the region where the data asset is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the sensitivity level. Valid values:
	//
	// - **1**: N/A
	//
	// - **2**: S1
	//
	// - **3**: S2
	//
	// - **4**: S3
	//
	// - **5**: S4
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level. Valid values:
	//
	// - **N/A**
	//
	// - **S1**
	//
	// - **S2**
	//
	// - **S3**
	//
	// - **S4**
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the sensitive data detection rule that was matched.
	//
	// example:
	//
	// 1004
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule that was matched.
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

func (s DescribeDataObjectColumnDetailResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetCategories() []*string {
	return s.Categories
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetColumnComment() *string {
	return s.ColumnComment
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetMaskingStatus() *int32 {
	return s.MaskingStatus
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetModelTags() []*DescribeDataObjectColumnDetailResponseBodyItemsModelTags {
	return s.ModelTags
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetCategories(v []*string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.Categories = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetColumnComment(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ColumnComment = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetColumnName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ColumnName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetDataType(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetEngineType(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetId(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetInstanceName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetMaskingStatus(v int32) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetModelTags(v []*DescribeDataObjectColumnDetailResponseBodyItemsModelTags) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetPrimaryKey(v bool) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetProductId(v int64) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRegionId(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRiskLevelName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRuleId(v int64) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRuleName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetTableName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) Validate() error {
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

type DescribeDataObjectColumnDetailResponseBodyItemsModelTags struct {
	// The ID of the data label. Valid values:
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
	// The name of the data label. Valid values:
	//
	// - **Personal sensitive information**
	//
	// - **Personal information**
	//
	// - **General information**
	//
	// example:
	//
	// personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponseBodyItemsModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) GetName() *string {
	return s.Name
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) SetId(v int64) *DescribeDataObjectColumnDetailResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) SetName(v string) *DescribeDataObjectColumnDetailResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) Validate() error {
	return dara.Validate(s)
}
