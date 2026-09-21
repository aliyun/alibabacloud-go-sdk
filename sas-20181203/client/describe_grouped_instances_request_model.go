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
	// The page number of the first page to return. Default value: **1**, which indicates that results are returned starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the asset group to query. Fuzzy match is supported.
	//
	// example:
	//
	// test-01
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The filter condition for querying assets. Valid values:
	//
	// - **groupId**: queries assets by asset group.
	//
	// - **regionId**: queries assets by region.
	//
	// - **vpcInstanceId**: queries assets by virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// groupId
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of assets to query. Set the value to **ecs**, which indicates Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// Settings for whether to enable paged query. Default value: **true**. Valid values:
	//
	// - **true**: Paging is enabled.
	//
	// - **false**: Paging is disabled.
	//
	// example:
	//
	// true
	NoPage *bool `json:"NoPage,omitempty" xml:"NoPage,omitempty"`
	// The number of entries per page in a paged query. Default value: **20**, which indicates that 20 entries of asset information are displayed per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with member accounts in a resource folder.
	//
	// >Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The edition-based filter condition for querying assets. Valid values:
	//
	// - **sas_gte_advanced**: Advanced Edition or higher
	//
	// - **sas_gte_enterprise**: Enterprise Edition or higher
	//
	// - **sas_gt_basic**: paid edition
	//
	// - **sas_eq_advanced**: Advanced Edition
	//
	// - **sas_gt_anti_virus**: higher than Anti-virus Edition
	//
	// example:
	//
	// sas_gt_basic
	SaleVersionCheckCode *string `json:"SaleVersionCheckCode,omitempty" xml:"SaleVersionCheckCode,omitempty"`
	// The server vendor. Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: non-cloud asset
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**: third-party cloud asset
	//
	// - **8**: lightweight asset
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The server vendors. Separate multiple vendors with commas (,). Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: non-cloud asset
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**: third-party cloud asset
	//
	// - **8**: lightweight asset
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
