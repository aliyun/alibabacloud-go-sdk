// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseType(v string) *DescribeDefenseRulesRequest
	GetDefenseType() *string
	SetInstanceId(v string) *DescribeDefenseRulesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseRulesRequest
	GetPageSize() *int32
	SetQuery(v string) *DescribeDefenseRulesRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeDefenseRulesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseRulesRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleType(v string) *DescribeDefenseRulesRequest
	GetRuleType() *string
}

type DescribeDefenseRulesRequest struct {
	// The type of the protection rule. Valid values:
	//
	// - **template*	- (default): template protection rules.
	//
	// - **resource**: rules for protected objects.
	//
	// - **global**: global rules.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The query conditions. Specify this parameter as a JSON string.
	//
	// > The query results for protection rules vary based on the query conditions. For more information, see **Query parameter details**.
	//
	// example:
	//
	// {\\"name\\":\\"IPblock_20220822_10\\",\\"scene\\":\\"custom_acl\\",\\"templateId\\":5327}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
	// The type of the protection rule. Valid values:
	//
	// - **whitelist**: a whitelist rule
	//
	// - **defense*	- (default): a protection rule
	//
	// > This parameter is required only when **DefenseType*	- is set to **template**.
	//
	// example:
	//
	// whitelist
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeDefenseRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *DescribeDefenseRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseRulesRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeDefenseRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseRulesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeDefenseRulesRequest) SetDefenseType(v string) *DescribeDefenseRulesRequest {
	s.DefenseType = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetInstanceId(v string) *DescribeDefenseRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetPageNumber(v int32) *DescribeDefenseRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetPageSize(v int32) *DescribeDefenseRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetQuery(v string) *DescribeDefenseRulesRequest {
	s.Query = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetRegionId(v string) *DescribeDefenseRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseRulesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetRuleType(v string) *DescribeDefenseRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeDefenseRulesRequest) Validate() error {
	return dara.Validate(s)
}
