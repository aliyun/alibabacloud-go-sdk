// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v int64) *DescribeSoarStrategyParamRequest
	GetStrategyId() *int64
}

type DescribeSoarStrategyParamRequest struct {
	// The ID of the policy.
	//
	// >  You can call the [DescribeSoarStrategies](~~DescribeSoarStrategies~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15553
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeSoarStrategyParamRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyParamRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyParamRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeSoarStrategyParamRequest) SetStrategyId(v int64) *DescribeSoarStrategyParamRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeSoarStrategyParamRequest) Validate() error {
	return dara.Validate(s)
}
