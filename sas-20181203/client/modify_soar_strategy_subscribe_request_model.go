// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySoarStrategySubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v int64) *ModifySoarStrategySubscribeRequest
	GetStrategyId() *int64
	SetSubscribeStatus(v bool) *ModifySoarStrategySubscribeRequest
	GetSubscribeStatus() *bool
}

type ModifySoarStrategySubscribeRequest struct {
	// The ID of the policy.
	//
	// >  You can call the [DescribeSoarStrategies](~~DescribeSoarStrategies~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8000
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// Specifies whether to create the policy template. Valid values:
	//
	// 	- true: creates the policy template
	//
	// 	- false: deletes the policy template
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SubscribeStatus *bool `json:"SubscribeStatus,omitempty" xml:"SubscribeStatus,omitempty"`
}

func (s ModifySoarStrategySubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySoarStrategySubscribeRequest) GoString() string {
	return s.String()
}

func (s *ModifySoarStrategySubscribeRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *ModifySoarStrategySubscribeRequest) GetSubscribeStatus() *bool {
	return s.SubscribeStatus
}

func (s *ModifySoarStrategySubscribeRequest) SetStrategyId(v int64) *ModifySoarStrategySubscribeRequest {
	s.StrategyId = &v
	return s
}

func (s *ModifySoarStrategySubscribeRequest) SetSubscribeStatus(v bool) *ModifySoarStrategySubscribeRequest {
	s.SubscribeStatus = &v
	return s
}

func (s *ModifySoarStrategySubscribeRequest) Validate() error {
	return dara.Validate(s)
}
