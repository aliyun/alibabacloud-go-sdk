// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeGroupedInstancesRequest
	GetCurrentPage() *int32
	SetFieldValue(v string) *DescribeGroupedInstancesRequest
	GetFieldValue() *string
	SetGroupField(v string) *DescribeGroupedInstancesRequest
	GetGroupField() *string
	SetLang(v string) *DescribeGroupedInstancesRequest
	GetLang() *string
	SetMachineTypes(v string) *DescribeGroupedInstancesRequest
	GetMachineTypes() *string
	SetNoPage(v bool) *DescribeGroupedInstancesRequest
	GetNoPage() *bool
	SetPageSize(v int32) *DescribeGroupedInstancesRequest
	GetPageSize() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeGroupedInstancesRequest
	GetResourceDirectoryAccountId() *int64
	SetSaleVersionCheckCode(v string) *DescribeGroupedInstancesRequest
	GetSaleVersionCheckCode() *string
	SetVendor(v int32) *DescribeGroupedInstancesRequest
	GetVendor() *int32
	SetVendors(v string) *DescribeGroupedInstancesRequest
	GetVendors() *string
}

type DescribeGroupedInstancesRequest struct {
	// The page number from which query results start to be displayed. Default value: **1**, indicating that query results are displayed starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the asset group to query. Fuzzy search is supported.
	//
	// example:
	//
	// test-01
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The filter condition for querying assets. Valid values:
	//
	// - **groupId**: queries assets by group.
	//
	// - **regionId**: queries assets by region.
	//
	// - **vpcInstanceId**: queries assets by Virtual Private Cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// groupId
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	// The language type for requests and responses. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of assets to query. Fixed value: **ecs**, indicating Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// Specifies whether to enable paginated queries. Default value: **true**. Valid values:
	//
	// - **true**: enables paginated queries.
	//
	// - **false**: disables paginated queries.
	//
	// example:
	//
	// true
	NoPage *bool `json:"NoPage,omitempty" xml:"NoPage,omitempty"`
	// The number of entries per page in a paginated query. Default value: **20**, indicating that 20 entries of asset information are displayed per page.
	//
	// example:
	//
	// 20
	PageSize                   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The edition filter for querying assets. Valid values:
	//
	// - **sas_gte_advanced**: Advanced edition or higher
	//
	// - **sas_gte_enterprise**: Enterprise edition or higher
	//
	// - **sas_gt_basic**: paid editions
	//
	// - **sas_eq_advanced**: Advanced edition only
	//
	// - **sas_gt_anti_virus**: editions higher than Anti-virus edition
	//
	// example:
	//
	// sas_gt_basic
	SaleVersionCheckCode *string `json:"SaleVersionCheckCode,omitempty" xml:"SaleVersionCheckCode,omitempty"`
	// The server vendor. Valid values:
	//
	// - **0**: Alibaba Cloud assets
	//
	// - **1**: non-cloud assets
	//
	// - **2**: IDC assets
	//
	// - **3**, **4**, **5**, **7**: assets from other cloud providers
	//
	// - **8**: lightweight assets
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The server vendors. Separate multiple vendors with commas (,). Valid values:
	//
	// - **0**: Alibaba Cloud assets
	//
	// - **1**: non-cloud assets
	//
	// - **2**: IDC assets
	//
	// - **3**, **4**, **5**, **7**: assets from other cloud providers
	//
	// - **8**: lightweight assets
	//
	// example:
	//
	// 0,8
	Vendors *string `json:"Vendors,omitempty" xml:"Vendors,omitempty"`
}

func (s DescribeGroupedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedInstancesRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeGroupedInstancesRequest) GetGroupField() *string {
	return s.GroupField
}

func (s *DescribeGroupedInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupedInstancesRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeGroupedInstancesRequest) GetNoPage() *bool {
	return s.NoPage
}

func (s *DescribeGroupedInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedInstancesRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeGroupedInstancesRequest) GetSaleVersionCheckCode() *string {
	return s.SaleVersionCheckCode
}

func (s *DescribeGroupedInstancesRequest) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeGroupedInstancesRequest) GetVendors() *string {
	return s.Vendors
}

func (s *DescribeGroupedInstancesRequest) SetCurrentPage(v int32) *DescribeGroupedInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetFieldValue(v string) *DescribeGroupedInstancesRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetGroupField(v string) *DescribeGroupedInstancesRequest {
	s.GroupField = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetLang(v string) *DescribeGroupedInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetMachineTypes(v string) *DescribeGroupedInstancesRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetNoPage(v bool) *DescribeGroupedInstancesRequest {
	s.NoPage = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetPageSize(v int32) *DescribeGroupedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetResourceDirectoryAccountId(v int64) *DescribeGroupedInstancesRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetSaleVersionCheckCode(v string) *DescribeGroupedInstancesRequest {
	s.SaleVersionCheckCode = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetVendor(v int32) *DescribeGroupedInstancesRequest {
	s.Vendor = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) SetVendors(v string) *DescribeGroupedInstancesRequest {
	s.Vendors = &v
	return s
}

func (s *DescribeGroupedInstancesRequest) Validate() error {
	return dara.Validate(s)
}
