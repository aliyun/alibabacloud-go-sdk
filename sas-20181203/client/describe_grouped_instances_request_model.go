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
	SetSaleVersionCheckCode(v string) *DescribeGroupedInstancesRequest
	GetSaleVersionCheckCode() *string
	SetVendor(v int32) *DescribeGroupedInstancesRequest
	GetVendor() *int32
	SetVendors(v string) *DescribeGroupedInstancesRequest
	GetVendors() *string
}

type DescribeGroupedInstancesRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the group to which the assets belong. Fuzzy search is supported.
	//
	// example:
	//
	// test-01
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The filter condition that you want to use to query the assets. Valid values:
	//
	// 	- **groupId**: the group to which the assets belong
	//
	// 	- **regionId**: the region in which the assets reside
	//
	// 	- **vpcInstanceId**: the virtual private cloud (VPC) in which the assets reside
	//
	// This parameter is required.
	//
	// example:
	//
	// groupId
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the assets that you want to query. Set the value to **ecs**, which indicates Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// Specifies whether to enable paged query. Default value: **true**. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	NoPage *bool `json:"NoPage,omitempty" xml:"NoPage,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The edition of Security Center that protects the asset. Valid values:
	//
	// 	- **sas_gte_advanced**: the Advanced edition or higher
	//
	// 	- **sas_gte_enterprise**: the Enterprise edition or higher
	//
	// 	- **sas_gt_basic:*	- a paid edition
	//
	// 	- **sas_eq_advanced:*	- the Advanced edition
	//
	// 	- **sas_gt_anti_virus:*	- an edition higher than the Anti-virus edition
	//
	// example:
	//
	// sas_gt_basic
	SaleVersionCheckCode *string `json:"SaleVersionCheckCode,omitempty" xml:"SaleVersionCheckCode,omitempty"`
	// The source of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// 	- **1**: a third-party cloud server
	//
	// 	- **2**: a server in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight asset
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The source of the server. Separate multiple sources with commas (,).Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// 	- **1**: a third-party cloud server
	//
	// 	- **2**: a server in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight asset
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
