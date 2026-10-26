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
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of records returned by the query.
	Items []*ListTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The pagination cursor used for the current query.
	//
	// example:
	//
	// 0
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The cursor used to query the next page.
	//
	// example:
	//
	// 1001
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The number of entries returned per page for the current query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request. You can use the ID to locate the API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the current statistical or query scope.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Indicates whether more pages exist. A value of true indicates that more data exists. A value of false indicates that no more data exists.
	//
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
	// The description of the database table.
	//
	// example:
	//
	// Customer contact information
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the data object was created. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1788566400000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance identity of the cloud instance.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	DataAssetSourceId *string `json:"DataAssetSourceId,omitempty" xml:"DataAssetSourceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// business_db
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The description of the database instance.
	//
	// example:
	//
	// Business database instance
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the database table. Some engines return the name in the schema.table format.
	//
	// example:
	//
	// customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account that owns the database table.
	//
	// example:
	//
	// 123456789012****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the sensitivity level.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The list of detection models or rules.
	RuleList []*ListTablesResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// Indicates whether the data is detected as sensitive data.
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The number of sensitive data entries within the current statistical scope.
	//
	// example:
	//
	// 20
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The proportion of sensitive data.
	//
	// example:
	//
	// 0.2
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// The name of the tenant to which the data asset belongs.
	//
	// example:
	//
	// Business tenant
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of entries that meet the current statistical or query scope.
	//
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
	// The number of hits for the detection rule.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the detection model or rule.
	//
	// example:
	//
	// Phone number
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the sensitivity level.
	//
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
