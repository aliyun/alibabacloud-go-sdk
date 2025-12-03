// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRulesRequest
	GetResourceOwnerId() *int64
	SetRuleIds(v string) *DeleteRulesRequest
	GetRuleIds() *string
}

type DeleteRulesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of forwarding rules that you want to delete.
	//
	// >  The RuleIds parameter is required. You can specify up to 10 forwarding rules in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["rule-bp1z9ce******","rule-bp1tuc******4"]
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s DeleteRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRulesRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *DeleteRulesRequest) SetOwnerAccount(v string) *DeleteRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRulesRequest) SetOwnerId(v int64) *DeleteRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRulesRequest) SetRegionId(v string) *DeleteRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRulesRequest) SetResourceOwnerAccount(v string) *DeleteRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRulesRequest) SetResourceOwnerId(v int64) *DeleteRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRulesRequest) SetRuleIds(v string) *DeleteRulesRequest {
	s.RuleIds = &v
	return s
}

func (s *DeleteRulesRequest) Validate() error {
	return dara.Validate(s)
}
