// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMajorProtectionBlackIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeMajorProtectionBlackIpsRequest
	GetInstanceId() *string
	SetIpLike(v string) *DescribeMajorProtectionBlackIpsRequest
	GetIpLike() *string
	SetOrderBy(v string) *DescribeMajorProtectionBlackIpsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeMajorProtectionBlackIpsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMajorProtectionBlackIpsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMajorProtectionBlackIpsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeMajorProtectionBlackIpsRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DescribeMajorProtectionBlackIpsRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsRequest
	GetTemplateId() *int64
}

type DescribeMajorProtectionBlackIpsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address that you want to query. You can specify this parameter to query an IP address in the IP address blacklist for major event protection by using fuzzy matching.
	//
	// example:
	//
	// 192.0.XX.XX
	IpLike *string `json:"IpLike,omitempty" xml:"IpLike,omitempty"`
	// The method that you want to use to sort the IP addresses **in descending order**. Valid values:
	//
	// 	- **gmtModified:*	- sorts the IP addresses by most recent modification time.
	//
	// 	- **ip:*	- sorts the IP addresses by IP address.
	//
	// 	- **templateId:*	- sorts the IP addresses by template ID.
	//
	// 	- **id:*	- sorts the IP addresses by primary key.
	//
	// example:
	//
	// gmtModified
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	//
	// example:
	//
	// 20013199
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the rule template for major event protection.
	//
	// example:
	//
	// 5673
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetIpLike() *string {
	return s.IpLike
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeMajorProtectionBlackIpsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetInstanceId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetIpLike(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.IpLike = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetOrderBy(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetPageNumber(v int32) *DescribeMajorProtectionBlackIpsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetPageSize(v int32) *DescribeMajorProtectionBlackIpsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetRegionId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetResourceManagerResourceGroupId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetRuleId(v int64) *DescribeMajorProtectionBlackIpsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) Validate() error {
	return dara.Validate(s)
}
