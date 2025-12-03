// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRuleAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRuleAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRuleAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRuleAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRuleAttributeRequest
	GetResourceOwnerId() *int64
	SetRuleId(v string) *DescribeRuleAttributeRequest
	GetRuleId() *string
}

type DescribeRuleAttributeRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-bp1efemp9****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRuleAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRuleAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRuleAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRuleAttributeRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleAttributeRequest) SetOwnerAccount(v string) *DescribeRuleAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetOwnerId(v int64) *DescribeRuleAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetRegionId(v string) *DescribeRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetResourceOwnerAccount(v string) *DescribeRuleAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetResourceOwnerId(v int64) *DescribeRuleAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetRuleId(v string) *DescribeRuleAttributeRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
