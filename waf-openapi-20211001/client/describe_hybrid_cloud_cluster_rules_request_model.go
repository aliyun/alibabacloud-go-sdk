// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *DescribeHybridCloudClusterRulesRequest
	GetClusterId() *int64
	SetInstanceId(v string) *DescribeHybridCloudClusterRulesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeHybridCloudClusterRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridCloudClusterRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridCloudClusterRulesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClusterRulesRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleContent(v string) *DescribeHybridCloudClusterRulesRequest
	GetRuleContent() *string
	SetRuleMatchType(v string) *DescribeHybridCloudClusterRulesRequest
	GetRuleMatchType() *string
	SetRuleType(v string) *DescribeHybridCloudClusterRulesRequest
	GetRuleType() *string
}

type DescribeHybridCloudClusterRulesRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is available only for hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to obtain information about hybrid cloud clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 428
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
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
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
	// The prefix of the rule value to match.
	//
	// example:
	//
	// 1.1.1.1
	RuleContent *string `json:"RuleContent,omitempty" xml:"RuleContent,omitempty"`
	// The matching logic. Valid values:
	//
	// - exact: exact match
	//
	// - regex: regular expression match
	//
	// example:
	//
	// exact
	RuleMatchType *string `json:"RuleMatchType,omitempty" xml:"RuleMatchType,omitempty"`
	// The type of the rule. Valid value:
	//
	// - **pullin**: traffic redirection
	//
	// This parameter is required.
	//
	// example:
	//
	// pullin
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeHybridCloudClusterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRulesRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *DescribeHybridCloudClusterRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudClusterRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridCloudClusterRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridCloudClusterRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudClusterRulesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudClusterRulesRequest) GetRuleContent() *string {
	return s.RuleContent
}

func (s *DescribeHybridCloudClusterRulesRequest) GetRuleMatchType() *string {
	return s.RuleMatchType
}

func (s *DescribeHybridCloudClusterRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHybridCloudClusterRulesRequest) SetClusterId(v int64) *DescribeHybridCloudClusterRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetInstanceId(v string) *DescribeHybridCloudClusterRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetPageNumber(v int32) *DescribeHybridCloudClusterRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetPageSize(v int32) *DescribeHybridCloudClusterRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetRegionId(v string) *DescribeHybridCloudClusterRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClusterRulesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetRuleContent(v string) *DescribeHybridCloudClusterRulesRequest {
	s.RuleContent = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetRuleMatchType(v string) *DescribeHybridCloudClusterRulesRequest {
	s.RuleMatchType = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) SetRuleType(v string) *DescribeHybridCloudClusterRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesRequest) Validate() error {
	return dara.Validate(s)
}
