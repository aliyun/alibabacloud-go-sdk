// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiExportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeApiExportsRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeApiExportsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeApiExportsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApiExportsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApiExportsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApiExportsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApiExportsRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Language type. Valid values:
	//
	// - **cn**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Web Application Firewall (WAF) instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeApiExportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiExportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiExportsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApiExportsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApiExportsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApiExportsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApiExportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiExportsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApiExportsRequest) SetInstanceId(v string) *DescribeApiExportsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiExportsRequest) SetLang(v string) *DescribeApiExportsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApiExportsRequest) SetPageNumber(v int64) *DescribeApiExportsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiExportsRequest) SetPageSize(v int64) *DescribeApiExportsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiExportsRequest) SetRegionId(v string) *DescribeApiExportsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApiExportsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApiExportsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApiExportsRequest) Validate() error {
	return dara.Validate(s)
}
