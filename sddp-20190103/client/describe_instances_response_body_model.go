// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstancesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeInstancesResponseBodyItems) *DescribeInstancesResponseBody
	GetItems() []*DescribeInstancesResponseBodyItems
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesResponseBody struct {
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The details of the data asset instances returned.
	Items []*DescribeInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of data asset instances on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71064826-726F-4ADA-B879-05D8055476FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data asset instances returned.
	//
	// example:
	//
	// 231
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstancesResponseBody) GetItems() []*DescribeInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetCurrentPage(v int32) *DescribeInstancesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetItems(v []*DescribeInstancesResponseBodyItems) *DescribeInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
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

type DescribeInstancesResponseBodyItems struct {
	// The time when the data asset instance was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637226782000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The name of the department to which the data asset instance belongs.
	//
	// example:
	//
	// ***DemoCenter
	DepartName *string `json:"DepartName,omitempty" xml:"DepartName,omitempty"`
	// The unique ID of the data asset instance recorded in Data Security Center.
	//
	// example:
	//
	// 11111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the data asset instance.
	//
	// example:
	//
	// instance dscription
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The security status of the data asset instance. Valid values:
	//
	// - **true**: Secure.
	//
	// - **false**: Not secure.
	//
	// example:
	//
	// true
	Labelsec *bool `json:"Labelsec,omitempty" xml:"Labelsec,omitempty"`
	// The time when the most recent scan of the data asset instance was completed. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637622793000
	LastFinishTime *int64 `json:"LastFinishTime,omitempty" xml:"LastFinishTime,omitempty"`
	// If the management account has enabled multi-account management and the asset belongs to another member account, this field displays the UID of the member account.
	//
	// example:
	//
	// 12567890126
	MemberAliUid *string `json:"MemberAliUid,omitempty" xml:"MemberAliUid,omitempty"`
	// The list of data tags.
	ModelTags []*DescribeInstancesResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the data asset instance.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The Alibaba Cloud account that owns the data asset instance.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the product to which the data asset instance belongs, such as MaxCompute, OSS, or RDS. For supported product names, see [Data types from which sensitive data can be detected](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the product to which the data asset instance belongs.
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The protection status of the data asset instance. Valid values:
	//
	// - **true**: Protected.
	//
	// - **false**: Not protected.
	//
	// example:
	//
	// false
	Protection *bool `json:"Protection,omitempty" xml:"Protection,omitempty"`
	// The risk level ID of the data asset instance. A higher risk level ID indicates more sensitive data is detected.
	//
	// - **1**: No sensitive data is detected. No risk.
	//
	// - **2**: Sensitive data risk at level 1.
	//
	// - **3**: Sensitive data risk at level 2.
	//
	// - **4**: Sensitive data risk at level 3.
	//
	// - **5**: Sensitive data risk at level 4.
	//
	// - **6**: Sensitive data risk at level 5.
	//
	// - **7**: Sensitive data risk at level 6.
	//
	// - **8**: Sensitive data risk at level 7.
	//
	// - **9**: Sensitive data risk at level 8.
	//
	// - **10**: Sensitive data risk at level 9.
	//
	// - **11**: Sensitive data risk at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The risk level name of the data asset instance.
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data detection rule that the data asset instance hits.
	//
	// example:
	//
	// **	- rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the data asset instance contains sensitive data. Valid values:
	//
	// - **true**: Contains sensitive data.
	//
	// - **false**: Does not contain sensitive data.
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive data items in the data asset instance. For example, if the data asset is ApsaraDB RDS, this value indicates the total number of sensitive tables in the databases of the instance.
	//
	// example:
	//
	// 123
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// tenant
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of data items in the data asset instance. For example, if the data asset is ApsaraDB RDS, this value indicates the total number of tables in the databases of the instance.
	//
	// example:
	//
	// 231
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeInstancesResponseBodyItems) GetDepartName() *string {
	return s.DepartName
}

func (s *DescribeInstancesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeInstancesResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeInstancesResponseBodyItems) GetLabelsec() *bool {
	return s.Labelsec
}

func (s *DescribeInstancesResponseBodyItems) GetLastFinishTime() *int64 {
	return s.LastFinishTime
}

func (s *DescribeInstancesResponseBodyItems) GetMemberAliUid() *string {
	return s.MemberAliUid
}

func (s *DescribeInstancesResponseBodyItems) GetModelTags() []*DescribeInstancesResponseBodyItemsModelTags {
	return s.ModelTags
}

func (s *DescribeInstancesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesResponseBodyItems) GetOdpsRiskLevelName() *string {
	return s.OdpsRiskLevelName
}

func (s *DescribeInstancesResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribeInstancesResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstancesResponseBodyItems) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeInstancesResponseBodyItems) GetProtection() *bool {
	return s.Protection
}

func (s *DescribeInstancesResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeInstancesResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeInstancesResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeInstancesResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *DescribeInstancesResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *DescribeInstancesResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeInstancesResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBodyItems) SetCreationTime(v int64) *DescribeInstancesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetDepartName(v string) *DescribeInstancesResponseBodyItems {
	s.DepartName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetId(v int64) *DescribeInstancesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetInstanceDescription(v string) *DescribeInstancesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetLabelsec(v bool) *DescribeInstancesResponseBodyItems {
	s.Labelsec = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetLastFinishTime(v int64) *DescribeInstancesResponseBodyItems {
	s.LastFinishTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetMemberAliUid(v string) *DescribeInstancesResponseBodyItems {
	s.MemberAliUid = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetModelTags(v []*DescribeInstancesResponseBodyItemsModelTags) *DescribeInstancesResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetName(v string) *DescribeInstancesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeInstancesResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetOwner(v string) *DescribeInstancesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProductCode(v string) *DescribeInstancesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProductId(v string) *DescribeInstancesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProtection(v bool) *DescribeInstancesResponseBodyItems {
	s.Protection = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRiskLevelId(v int64) *DescribeInstancesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRiskLevelName(v string) *DescribeInstancesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRuleName(v string) *DescribeInstancesResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetSensitive(v bool) *DescribeInstancesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetSensitiveCount(v int32) *DescribeInstancesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetTenantName(v string) *DescribeInstancesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetTotalCount(v int32) *DescribeInstancesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) Validate() error {
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

type DescribeInstancesResponseBodyItemsModelTags struct {
	// The data tag ID. Valid values:
	//
	// - **101**: personal sensitive information
	//
	// - **102**: personal information
	//
	// - **107**: general information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data tag name. Valid values:
	//
	// - 个人敏感信息
	//
	// - 个人信息
	//
	// - 通用信息
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstancesResponseBodyItemsModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyItemsModelTags) GetId() *int64 {
	return s.Id
}

func (s *DescribeInstancesResponseBodyItemsModelTags) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesResponseBodyItemsModelTags) SetId(v int64) *DescribeInstancesResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeInstancesResponseBodyItemsModelTags) SetName(v string) *DescribeInstancesResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyItemsModelTags) Validate() error {
	return dara.Validate(s)
}
