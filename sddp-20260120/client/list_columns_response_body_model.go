// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListColumnsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListColumnsResponseBodyItems) *ListColumnsResponseBody
	GetItems() []*ListColumnsResponseBodyItems
	SetPageSize(v int32) *ListColumnsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListColumnsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListColumnsResponseBody
	GetTotalCount() *int32
}

type ListColumnsResponseBody struct {
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of records returned by this query.
	Items []*ListColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of records per page for this query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which can be used to locate this API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the current query scope.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListColumnsResponseBody) GetItems() []*ListColumnsResponseBodyItems {
	return s.Items
}

func (s *ListColumnsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListColumnsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListColumnsResponseBody) SetCurrentPage(v int32) *ListColumnsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListColumnsResponseBody) SetItems(v []*ListColumnsResponseBodyItems) *ListColumnsResponseBody {
	s.Items = v
	return s
}

func (s *ListColumnsResponseBody) SetPageSize(v int32) *ListColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListColumnsResponseBody) SetRequestId(v string) *ListColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListColumnsResponseBody) SetTotalCount(v int32) *ListColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListColumnsResponseBody) Validate() error {
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

type ListColumnsResponseBodyItems struct {
	// The creation time of the data object. This value is a UNIX timestamp in milliseconds since 1970-01-01 00:00:00 UTC.
	//
	// example:
	//
	// 1788566400000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cloud instance identity.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// The database name.
	//
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The data type of the database column.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The cloud instance identity to which the column belongs. This is a compatible field.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The compatible field for column encryption status. To query column encryption results, call ListDataMaskingColumns.
	//
	// example:
	//
	// 0
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// The name of the data column.
	//
	// example:
	//
	// phone
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cloud product code.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the detection result revision record.
	//
	// example:
	//
	// 1001
	RevisionId *int64 `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	// The revision status of the detection result. Valid values:
	//
	// - 0: not revised.
	//
	// - 1: revised.
	//
	// example:
	//
	// 0
	RevisionStatus *int64 `json:"RevisionStatus,omitempty" xml:"RevisionStatus,omitempty"`
	// The sensitivity level ID.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The sensitivity level name.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the matched detection rule.
	//
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the matched detection rule.
	//
	// example:
	//
	// Phone number
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The database schema name.
	//
	// example:
	//
	// public
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// Indicates whether the data is detected as sensitive data.
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The data table name.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListColumnsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListColumnsResponseBodyItems) GetDataAssetSourceId() *string {
	return s.DataAssetSourceId
}

func (s *ListColumnsResponseBodyItems) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListColumnsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *ListColumnsResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *ListColumnsResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListColumnsResponseBodyItems) GetMaskingStatus() *int32 {
	return s.MaskingStatus
}

func (s *ListColumnsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListColumnsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListColumnsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListColumnsResponseBodyItems) GetRevisionId() *int64 {
	return s.RevisionId
}

func (s *ListColumnsResponseBodyItems) GetRevisionStatus() *int64 {
	return s.RevisionStatus
}

func (s *ListColumnsResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListColumnsResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *ListColumnsResponseBodyItems) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListColumnsResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *ListColumnsResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListColumnsResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *ListColumnsResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *ListColumnsResponseBodyItems) SetCreationTime(v int64) *ListColumnsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetDataAssetSourceId(v string) *ListColumnsResponseBodyItems {
	s.DataAssetSourceId = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetDataSourceName(v string) *ListColumnsResponseBodyItems {
	s.DataSourceName = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetDataType(v string) *ListColumnsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetEngineType(v string) *ListColumnsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetInstanceName(v string) *ListColumnsResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetMaskingStatus(v int32) *ListColumnsResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetName(v string) *ListColumnsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetProductCode(v string) *ListColumnsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRegionId(v string) *ListColumnsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRevisionId(v int64) *ListColumnsResponseBodyItems {
	s.RevisionId = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRevisionStatus(v int64) *ListColumnsResponseBodyItems {
	s.RevisionStatus = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRiskLevelId(v int64) *ListColumnsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRiskLevelName(v string) *ListColumnsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRuleId(v int64) *ListColumnsResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetRuleName(v string) *ListColumnsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetSchemaName(v string) *ListColumnsResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetSensitive(v bool) *ListColumnsResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *ListColumnsResponseBodyItems) SetTableName(v string) *ListColumnsResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *ListColumnsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
