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
	CurrentPage *int32                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreationTime      *int64  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	DataSourceName    *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataType          *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EngineType        *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MaskingStatus     *int32  `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RevisionId        *int64  `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	RevisionStatus    *int64  `json:"RevisionStatus,omitempty" xml:"RevisionStatus,omitempty"`
	RiskLevelId       *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RiskLevelName     *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	RuleId            *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SchemaName        *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Sensitive         *bool   `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
