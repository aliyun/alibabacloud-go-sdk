// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTablesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListTablesResponseBodyItems) *ListTablesResponseBody
	GetItems() []*ListTablesResponseBodyItems
	SetMarker(v string) *ListTablesResponseBody
	GetMarker() *string
	SetNextMarker(v string) *ListTablesResponseBody
	GetNextMarker() *string
	SetPageSize(v int32) *ListTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTablesResponseBody
	GetTotalCount() *int32
	SetTruncated(v string) *ListTablesResponseBody
	GetTruncated() *string
}

type ListTablesResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// 1001
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
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
	// example:
	//
	// false
	Truncated *string `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTablesResponseBody) GetItems() []*ListTablesResponseBodyItems {
	return s.Items
}

func (s *ListTablesResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListTablesResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTablesResponseBody) GetTruncated() *string {
	return s.Truncated
}

func (s *ListTablesResponseBody) SetCurrentPage(v int32) *ListTablesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTablesResponseBody) SetItems(v []*ListTablesResponseBodyItems) *ListTablesResponseBody {
	s.Items = v
	return s
}

func (s *ListTablesResponseBody) SetMarker(v string) *ListTablesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListTablesResponseBody) SetNextMarker(v string) *ListTablesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListTablesResponseBody) SetPageSize(v int32) *ListTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetTotalCount(v int32) *ListTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTablesResponseBody) SetTruncated(v string) *ListTablesResponseBody {
	s.Truncated = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
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

type ListTablesResponseBodyItems struct {
	// example:
	//
	// 客户联系方式
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1788566400000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// 业务数据库实例
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// example:
	//
	// customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456789012****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// example:
	//
	// S2
	RiskLevelName *string                                `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	RuleList      []*ListTablesResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	Sensitive     *bool                                  `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// example:
	//
	// 20
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// example:
	//
	// 0.2
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// example:
	//
	// 业务租户
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTablesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyItems) GetComment() *string {
	return s.Comment
}

func (s *ListTablesResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListTablesResponseBodyItems) GetDataAssetSourceId() *string {
	return s.DataAssetSourceId
}

func (s *ListTablesResponseBodyItems) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListTablesResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ListTablesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListTablesResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *ListTablesResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListTablesResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListTablesResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *ListTablesResponseBodyItems) GetRuleList() []*ListTablesResponseBodyItemsRuleList {
	return s.RuleList
}

func (s *ListTablesResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *ListTablesResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *ListTablesResponseBodyItems) GetSensitiveRatio() *string {
	return s.SensitiveRatio
}

func (s *ListTablesResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *ListTablesResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTablesResponseBodyItems) SetComment(v string) *ListTablesResponseBodyItems {
	s.Comment = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetCreationTime(v int64) *ListTablesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetDataAssetSourceId(v string) *ListTablesResponseBodyItems {
	s.DataAssetSourceId = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetDataSourceName(v string) *ListTablesResponseBodyItems {
	s.DataSourceName = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetInstanceDescription(v string) *ListTablesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetName(v string) *ListTablesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetOwner(v string) *ListTablesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetProductCode(v string) *ListTablesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetRiskLevelId(v int64) *ListTablesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetRiskLevelName(v string) *ListTablesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetRuleList(v []*ListTablesResponseBodyItemsRuleList) *ListTablesResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *ListTablesResponseBodyItems) SetSensitive(v bool) *ListTablesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetSensitiveCount(v int32) *ListTablesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetSensitiveRatio(v string) *ListTablesResponseBodyItems {
	s.SensitiveRatio = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetTenantName(v string) *ListTablesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *ListTablesResponseBodyItems) SetTotalCount(v int32) *ListTablesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *ListTablesResponseBodyItems) Validate() error {
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

type ListTablesResponseBodyItemsRuleList struct {
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 手机号
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
}

func (s ListTablesResponseBodyItemsRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyItemsRuleList) GetCount() *int64 {
	return s.Count
}

func (s *ListTablesResponseBodyItemsRuleList) GetName() *string {
	return s.Name
}

func (s *ListTablesResponseBodyItemsRuleList) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ListTablesResponseBodyItemsRuleList) SetCount(v int64) *ListTablesResponseBodyItemsRuleList {
	s.Count = &v
	return s
}

func (s *ListTablesResponseBodyItemsRuleList) SetName(v string) *ListTablesResponseBodyItemsRuleList {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyItemsRuleList) SetRiskLevelId(v int64) *ListTablesResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *ListTablesResponseBodyItemsRuleList) Validate() error {
	return dara.Validate(s)
}
