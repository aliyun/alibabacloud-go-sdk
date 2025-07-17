// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeDispatchRuleRequest
	GetId() *string
	SetRegionId(v string) *DescribeDispatchRuleRequest
	GetRegionId() *string
}

type DescribeDispatchRuleRequest struct {
	// The ID of the dispatch policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDispatchRuleRequest) SetId(v string) *DescribeDispatchRuleRequest {
	s.Id = &v
	return s
}

func (s *DescribeDispatchRuleRequest) SetRegionId(v string) *DescribeDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}
