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
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address to query. You can set this parameter to perform a fuzzy query on the added IP address blacklist.
	//
	// example:
	//
	// 192.0.XX.XX
	IpLike *string `json:"IpLike,omitempty" xml:"IpLike,omitempty"`
	// The property by which to sort the results in **descending order**. Valid values:
	//
	// - **gmtModified**: sorts by modification time.
	//
	// - **ip**: sorts by IP address.
	//
	// - **templateId**: sorts by template ID.
	//
	// - **id**: sorts by primary key.
	//
	// example:
	//
	// gmtModified
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**, which indicates 10 entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: regions outside the Chinese mainland.
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
	// The ID of the IP address blacklist rule for critical event protection.
	//
	// example:
	//
	// 20013199
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the critical event protection template.
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
