// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeColumnsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody
	GetItems() []*DescribeColumnsResponseBodyItems
	SetPageSize(v int32) *DescribeColumnsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeColumnsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeColumnsResponseBody
	GetTotalCount() *int32
}

type DescribeColumnsResponseBody struct {
	// The page number of the current page displayed in the results.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Column data in the data asset table.
	Items []*DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of data entries displayed per page in the results.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID of the result.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data entries in the results.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeColumnsResponseBody) GetItems() []*DescribeColumnsResponseBodyItems {
	return s.Items
}

func (s *DescribeColumnsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColumnsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeColumnsResponseBody) SetCurrentPage(v int32) *DescribeColumnsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetItems(v []*DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsResponseBody) SetPageSize(v int32) *DescribeColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetTotalCount(v int32) *DescribeColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeColumnsResponseBody) Validate() error {
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

type DescribeColumnsResponseBodyItems struct {
	// The creation time of the column data in the data asset table, in milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data type of the column data in the data asset table.
	//
	// example:
	//
	// String
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Database engine type.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The unique ID of the column in the data asset table.
	//
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Column encryption status. Valid values:
	//
	// - **-1**: Not encrypted
	//
	// - **1**: Encryption successful
	//
	// - **2**: Encryption failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// The list of data tags hit by the detection model.
	ModelTags []*DescribeColumnsResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the column in the data asset table.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The risk level name of the asset. Valid values:
	//
	// - **N/A**: No sensitive data detected.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
	//
	// example:
	//
	// S3
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The risk level code of the asset. Valid values:
	//
	// - **1**: N/A.
	//
	// - **2**: S1.
	//
	// - **3**: S2.
	//
	// - **4**: S3.
	//
	// - **5**: S4.
	//
	// example:
	//
	// 3
	OdpsRiskLevelValue *int32 `json:"OdpsRiskLevelValue,omitempty" xml:"OdpsRiskLevelValue,omitempty"`
	// The product name to which the column data in the data asset table belongs. Valid values: **MaxCompute, OSS, ADS, OTS, RDS**, and others.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID corresponding to the product name to which the data object belongs. Valid values:
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
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region where the asset is located.
	//
	// example:
	//
	// cn-***
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Correction record ID.
	//
	// example:
	//
	// 12
	RevisionId *int64 `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	// Correction status. Valid values:
	//
	// - 1: Corrected.
	//
	// - 0: Not corrected.
	//
	// example:
	//
	// 1
	RevisionStatus *int64 `json:"RevisionStatus,omitempty" xml:"RevisionStatus,omitempty"`
	// The risk level ID of the column data in the data asset table. Valid values:
	//
	// - **1**: N/A.
	//
	// - **2**: S1.
	//
	// - **3**: S2.
	//
	// - **4**: S3.
	//
	// - **5**: S4.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The risk level name of the column data in the data asset table. Valid values:
	//
	// - **N/A**: No sensitive data detected.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the sensitive data detection rule hit by the column data in the data asset table.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule hit by the column data in the data asset table.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Sensitivity level name. Valid values:
	//
	// - **N/A**: No sensitive data detected.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// Indicates whether the column data in the data asset table contains sensitive data. Valid values:
	//
	// - true: The column data in the data asset table contains sensitive data.
	//
	// - false: The column data in the data asset table does not contain sensitive data.
	//
	// example:
	//
	// false
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The ID of the asset table to which the column data in the data asset table belongs.
	//
	// example:
	//
	// 123
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table to which the target column for correction belongs.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeColumnsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *DescribeColumnsResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeColumnsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeColumnsResponseBodyItems) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeColumnsResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeColumnsResponseBodyItems) GetMaskingStatus() *int32 {
	return s.MaskingStatus
}

func (s *DescribeColumnsResponseBodyItems) GetModelTags() []*DescribeColumnsResponseBodyItemsModelTags {
	return s.ModelTags
}

func (s *DescribeColumnsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeColumnsResponseBodyItems) GetOdpsRiskLevelName() *string {
	return s.OdpsRiskLevelName
}

func (s *DescribeColumnsResponseBodyItems) GetOdpsRiskLevelValue() *int32 {
	return s.OdpsRiskLevelValue
}

func (s *DescribeColumnsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeColumnsResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeColumnsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColumnsResponseBodyItems) GetRevisionId() *int64 {
	return s.RevisionId
}

func (s *DescribeColumnsResponseBodyItems) GetRevisionStatus() *int64 {
	return s.RevisionStatus
}

func (s *DescribeColumnsResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeColumnsResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeColumnsResponseBodyItems) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeColumnsResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeColumnsResponseBodyItems) GetSensLevelName() *string {
	return s.SensLevelName
}

func (s *DescribeColumnsResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *DescribeColumnsResponseBodyItems) GetTableId() *int64 {
	return s.TableId
}

func (s *DescribeColumnsResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColumnsResponseBodyItems) SetCreationTime(v int64) *DescribeColumnsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetDataType(v string) *DescribeColumnsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetEngineType(v string) *DescribeColumnsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetId(v string) *DescribeColumnsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetInstanceId(v int64) *DescribeColumnsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetInstanceName(v string) *DescribeColumnsResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetMaskingStatus(v int32) *DescribeColumnsResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetModelTags(v []*DescribeColumnsResponseBodyItemsModelTags) *DescribeColumnsResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetName(v string) *DescribeColumnsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetOdpsRiskLevelValue(v int32) *DescribeColumnsResponseBodyItems {
	s.OdpsRiskLevelValue = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetProductCode(v string) *DescribeColumnsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetProductId(v int64) *DescribeColumnsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRegionId(v string) *DescribeColumnsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRevisionId(v int64) *DescribeColumnsResponseBodyItems {
	s.RevisionId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRevisionStatus(v int64) *DescribeColumnsResponseBodyItems {
	s.RevisionStatus = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRiskLevelId(v int64) *DescribeColumnsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRiskLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRuleId(v int64) *DescribeColumnsResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRuleName(v string) *DescribeColumnsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetSensLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetSensitive(v bool) *DescribeColumnsResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetTableId(v int64) *DescribeColumnsResponseBodyItems {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetTableName(v string) *DescribeColumnsResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) Validate() error {
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

type DescribeColumnsResponseBodyItemsModelTags struct {
	// The data tag ID of the detection model.
	//
	// - **101**: Personal sensitive information.
	//
	// - **102**: Personal information.
	//
	// - **103**: Important data.
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data tag name of the detection model.
	//
	// - Personal sensitive information.
	//
	// - Personal information.
	//
	// - Important data.
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsModelTags) GetId() *int64 {
	return s.Id
}

func (s *DescribeColumnsResponseBodyItemsModelTags) GetName() *string {
	return s.Name
}

func (s *DescribeColumnsResponseBodyItemsModelTags) SetId(v int64) *DescribeColumnsResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsModelTags) SetName(v string) *DescribeColumnsResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsModelTags) Validate() error {
	return dara.Validate(s)
}
