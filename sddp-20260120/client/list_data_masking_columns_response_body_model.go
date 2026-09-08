// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDataMaskingColumnsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListDataMaskingColumnsResponseBodyItems) *ListDataMaskingColumnsResponseBody
	GetItems() []*ListDataMaskingColumnsResponseBodyItems
	SetPageSize(v int32) *ListDataMaskingColumnsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataMaskingColumnsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataMaskingColumnsResponseBody
	GetTotalCount() *int32
}

type ListDataMaskingColumnsResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                     `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListDataMaskingColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataMaskingColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataMaskingColumnsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataMaskingColumnsResponseBody) GetItems() []*ListDataMaskingColumnsResponseBodyItems {
	return s.Items
}

func (s *ListDataMaskingColumnsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataMaskingColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataMaskingColumnsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataMaskingColumnsResponseBody) SetCurrentPage(v int32) *ListDataMaskingColumnsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBody) SetItems(v []*ListDataMaskingColumnsResponseBodyItems) *ListDataMaskingColumnsResponseBody {
	s.Items = v
	return s
}

func (s *ListDataMaskingColumnsResponseBody) SetPageSize(v int32) *ListDataMaskingColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBody) SetRequestId(v string) *ListDataMaskingColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBody) SetTotalCount(v int32) *ListDataMaskingColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBody) Validate() error {
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

type ListDataMaskingColumnsResponseBodyItems struct {
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// data_masking_not_running
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 实例未处于运行状态
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// column-example-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Success
	MaskingStatus *string                                             `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	ModelTags     []*ListDataMaskingColumnsResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// example:
	//
	// phone
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// example:
	//
	// 1001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 手机号
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDataMaskingColumnsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetDbName() *string {
	return s.DbName
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetMaskingStatus() *string {
	return s.MaskingStatus
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetModelTags() []*ListDataMaskingColumnsResponseBodyItemsModelTags {
	return s.ModelTags
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *ListDataMaskingColumnsResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetDataType(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetDbName(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.DbName = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetEngineType(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetErrorCode(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetErrorMessage(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetId(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetInstanceId(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetMaskingStatus(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetModelTags(v []*ListDataMaskingColumnsResponseBodyItemsModelTags) *ListDataMaskingColumnsResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetName(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetProductCode(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetProductId(v int64) *ListDataMaskingColumnsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetRegionId(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetRiskLevelId(v int64) *ListDataMaskingColumnsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetRiskLevelName(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetRuleId(v int64) *ListDataMaskingColumnsResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetRuleName(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) SetTableName(v string) *ListDataMaskingColumnsResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItems) Validate() error {
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

type ListDataMaskingColumnsResponseBodyItemsModelTags struct {
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 个人信息
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDataMaskingColumnsResponseBodyItemsModelTags) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingColumnsResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *ListDataMaskingColumnsResponseBodyItemsModelTags) GetId() *int64 {
	return s.Id
}

func (s *ListDataMaskingColumnsResponseBodyItemsModelTags) GetName() *string {
	return s.Name
}

func (s *ListDataMaskingColumnsResponseBodyItemsModelTags) SetId(v int64) *ListDataMaskingColumnsResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItemsModelTags) SetName(v string) *ListDataMaskingColumnsResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

func (s *ListDataMaskingColumnsResponseBodyItemsModelTags) Validate() error {
	return dara.Validate(s)
}
