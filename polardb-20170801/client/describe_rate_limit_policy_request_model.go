// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRateLimitPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DescribeRateLimitPolicyRequest
	GetGwClusterId() *string
	SetPageNumber(v int32) *DescribeRateLimitPolicyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRateLimitPolicyRequest
	GetPageSize() *int32
	SetPolicyId(v string) *DescribeRateLimitPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *DescribeRateLimitPolicyRequest
	GetRegionId() *string
	SetScopeRefId(v string) *DescribeRateLimitPolicyRequest
	GetScopeRefId() *string
	SetScopeType(v string) *DescribeRateLimitPolicyRequest
	GetScopeType() *string
}

type DescribeRateLimitPolicyRequest struct {
	// The ID of the gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The page number to return. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**. The default value is **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the rate limit policy.
	//
	// example:
	//
	// 02eccf7c61cf4d05a543075ee907f3**
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the target resource, which can be a consumer group or a consumer, depending on the `ScopeType` value.
	//
	// example:
	//
	// cg-xxxxxxxx
	ScopeRefId *string `json:"ScopeRefId,omitempty" xml:"ScopeRefId,omitempty"`
	// The scope of the rate limit policy. Valid values:
	//
	// - **ConsumerGroup**: The policy applies to a consumer group.
	//
	// - **Consumer**: The policy applies to a specific consumer.
	//
	// example:
	//
	// ConsumerGroup
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s DescribeRateLimitPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRateLimitPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeRateLimitPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeRateLimitPolicyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRateLimitPolicyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRateLimitPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeRateLimitPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRateLimitPolicyRequest) GetScopeRefId() *string {
	return s.ScopeRefId
}

func (s *DescribeRateLimitPolicyRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *DescribeRateLimitPolicyRequest) SetGwClusterId(v string) *DescribeRateLimitPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) SetPageNumber(v int32) *DescribeRateLimitPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) SetPageSize(v int32) *DescribeRateLimitPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) SetPolicyId(v string) *DescribeRateLimitPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) SetRegionId(v string) *DescribeRateLimitPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) SetScopeRefId(v string) *DescribeRateLimitPolicyRequest {
	s.ScopeRefId = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) SetScopeType(v string) *DescribeRateLimitPolicyRequest {
	s.ScopeType = &v
	return s
}

func (s *DescribeRateLimitPolicyRequest) Validate() error {
	return dara.Validate(s)
}
