// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyTaskId(v int64) *DescribeSoarStrategyTaskParamsRequest
	GetStrategyTaskId() *int64
}

type DescribeSoarStrategyTaskParamsRequest struct {
	// Strategy task ID.
	//
	// > You can obtain this parameter by calling the [DescribeSoarStrategyTasks](~~DescribeSoarStrategyTasks~~) interface.
	//
	// example:
	//
	// 100
	StrategyTaskId *int64 `json:"StrategyTaskId,omitempty" xml:"StrategyTaskId,omitempty"`
}

func (s DescribeSoarStrategyTaskParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskParamsRequest) GetStrategyTaskId() *int64 {
	return s.StrategyTaskId
}

func (s *DescribeSoarStrategyTaskParamsRequest) SetStrategyTaskId(v int64) *DescribeSoarStrategyTaskParamsRequest {
	s.StrategyTaskId = &v
	return s
}

func (s *DescribeSoarStrategyTaskParamsRequest) Validate() error {
	return dara.Validate(s)
}
