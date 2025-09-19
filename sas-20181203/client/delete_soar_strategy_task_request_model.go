// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSoarStrategyTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyTaskId(v int64) *DeleteSoarStrategyTaskRequest
	GetStrategyTaskId() *int64
}

type DeleteSoarStrategyTaskRequest struct {
	// The ID of the policy task that is in the waiting state.
	//
	// >  You can call the [DescribeSoarStrategyTasks](~~DescribeSoarStrategyTasks~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11082
	StrategyTaskId *int64 `json:"StrategyTaskId,omitempty" xml:"StrategyTaskId,omitempty"`
}

func (s DeleteSoarStrategyTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSoarStrategyTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteSoarStrategyTaskRequest) GetStrategyTaskId() *int64 {
	return s.StrategyTaskId
}

func (s *DeleteSoarStrategyTaskRequest) SetStrategyTaskId(v int64) *DeleteSoarStrategyTaskRequest {
	s.StrategyTaskId = &v
	return s
}

func (s *DeleteSoarStrategyTaskRequest) Validate() error {
	return dara.Validate(s)
}
