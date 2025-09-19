// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeStrategyTargetRequest
	GetConfig() *string
	SetSourceIp(v string) *DescribeStrategyTargetRequest
	GetSourceIp() *string
	SetType(v string) *DescribeStrategyTargetRequest
	GetType() *string
}

type DescribeStrategyTargetRequest struct {
	// The ID of the baseline check policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"strategyId":8167126}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the policy. Set the value to hc_strategy, which indicates baseline check policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc_strategy
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeStrategyTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyTargetRequest) GoString() string {
	return s.String()
}

func (s *DescribeStrategyTargetRequest) GetConfig() *string {
	return s.Config
}

func (s *DescribeStrategyTargetRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeStrategyTargetRequest) GetType() *string {
	return s.Type
}

func (s *DescribeStrategyTargetRequest) SetConfig(v string) *DescribeStrategyTargetRequest {
	s.Config = &v
	return s
}

func (s *DescribeStrategyTargetRequest) SetSourceIp(v string) *DescribeStrategyTargetRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStrategyTargetRequest) SetType(v string) *DescribeStrategyTargetRequest {
	s.Type = &v
	return s
}

func (s *DescribeStrategyTargetRequest) Validate() error {
	return dara.Validate(s)
}
