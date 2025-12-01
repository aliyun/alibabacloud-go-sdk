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
	// An array that consists of tables.
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
	// The point in time when the table was created. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 222
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the data asset.
	//
	// example:
	//
	// Description 1
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the data asset to which the table belongs.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the data asset to which the table belongs.
	//
	// example:
	//
	// Data Asset 1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account to which the table belongs.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the service to which the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**. For more information about the types of data assets from which DSC can scan for sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the table belongs.
	//
	// example:
	//
	// 1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the table. Each sensitivity level ID corresponds to a sensitivity level name. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the table. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The information about the sensitive data detection rules that are hit.
	RuleList []*DescribeTablesResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// Indicates whether the table contains sensitive fields. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive fields in the table.
	//
	// example:
	//
	// 32
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The percentage of sensitive fields in the table.
	//
	// example:
	//
	// 21%
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// Tenant 1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of fields in the table.
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
	// Rule name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
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
