// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeTablesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeTablesResponseBodyItems) *DescribeTablesResponseBody
	GetItems() []*DescribeTablesResponseBodyItems
	SetPageSize(v int32) *DescribeTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTablesResponseBody
	GetTotalCount() *int32
}

type DescribeTablesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of data asset tables.
	Items []*DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTablesResponseBody) GetItems() []*DescribeTablesResponseBodyItems {
	return s.Items
}

func (s *DescribeTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTablesResponseBody) SetCurrentPage(v int32) *DescribeTablesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesResponseBody) SetItems(v []*DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotalCount(v int32) *DescribeTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTablesResponseBody) Validate() error {
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

type DescribeTablesResponseBodyItems struct {
	// The time when the data asset table was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The unique ID of the data asset table.
	//
	// example:
	//
	// 222
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// instance description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the instance to which the data asset table belongs.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// rm-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the data asset table.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account that owns the data asset table.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the product to which the data asset table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**. For more information about the supported products, see [Data asset types that support sensitive data detection](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the product to which the data asset table belongs.
	//
	// example:
	//
	// 1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the risk level for the data asset table. Each risk level ID corresponds to a risk level name. Valid values:
	//
	// - **1**: N/A. No sensitive data is detected.
	//
	// - **2**: S1. Level 1 sensitive data.
	//
	// - **3**: S2. Level 2 sensitive data.
	//
	// - **4**: S3. Level 3 sensitive data.
	//
	// - **5**: S4. Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the risk level for the data asset table. Valid values:
	//
	// - **N/A**: No sensitive data is detected.
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
	// The information about the sensitive data detection rules that the data asset table hits.
	RuleList []*DescribeTablesResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// Indicates whether the data asset table contains sensitive fields.
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive fields in the data asset table.
	//
	// example:
	//
	// 32
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The percentage of sensitive fields in the data asset table.
	//
	// example:
	//
	// 21%
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// Tenate001
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of fields in the data asset table.
	//
	// example:
	//
	// 1234
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTablesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeTablesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeTablesResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeTablesResponseBodyItems) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeTablesResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTablesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeTablesResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribeTablesResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeTablesResponseBodyItems) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeTablesResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeTablesResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeTablesResponseBodyItems) GetRuleList() []*DescribeTablesResponseBodyItemsRuleList {
	return s.RuleList
}

func (s *DescribeTablesResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *DescribeTablesResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *DescribeTablesResponseBodyItems) GetSensitiveRatio() *string {
	return s.SensitiveRatio
}

func (s *DescribeTablesResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeTablesResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTablesResponseBodyItems) SetCreationTime(v int64) *DescribeTablesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetId(v int64) *DescribeTablesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceDescription(v string) *DescribeTablesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceId(v int64) *DescribeTablesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceName(v string) *DescribeTablesResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetName(v string) *DescribeTablesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetOwner(v string) *DescribeTablesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetProductCode(v string) *DescribeTablesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetProductId(v string) *DescribeTablesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRiskLevelId(v int64) *DescribeTablesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRiskLevelName(v string) *DescribeTablesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRuleList(v []*DescribeTablesResponseBodyItemsRuleList) *DescribeTablesResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitive(v bool) *DescribeTablesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitiveCount(v int32) *DescribeTablesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitiveRatio(v string) *DescribeTablesResponseBodyItems {
	s.SensitiveRatio = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetTenantName(v string) *DescribeTablesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetTotalCount(v int32) *DescribeTablesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) Validate() error {
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

type DescribeTablesResponseBodyItemsRuleList struct {
	// The total number of rules.
	//
	// example:
	//
	// 12
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the risk level for the sensitive data detection rule. Valid values:
	//
	// - **1**: N/A. No sensitive data is detected.
	//
	// - **2**: S1. Level 1 sensitive data.
	//
	// - **3**: S2. Level 2 sensitive data.
	//
	// - **4**: S3. Level 3 sensitive data.
	//
	// - **5**: S4. Level 4 sensitive data.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
}

func (s DescribeTablesResponseBodyItemsRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsRuleList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeTablesResponseBodyItemsRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeTablesResponseBodyItemsRuleList) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetCount(v int64) *DescribeTablesResponseBodyItemsRuleList {
	s.Count = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetName(v string) *DescribeTablesResponseBodyItemsRuleList {
	s.Name = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeTablesResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsRuleList) Validate() error {
	return dara.Validate(s)
}
