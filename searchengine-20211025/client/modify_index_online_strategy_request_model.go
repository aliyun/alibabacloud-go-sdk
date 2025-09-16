// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexOnlineStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeRate(v int32) *ModifyIndexOnlineStrategyRequest
	GetChangeRate() *int32
}

type ModifyIndexOnlineStrategyRequest struct {
	// The index change rate.
	//
	// example:
	//
	// 20
	ChangeRate *int32 `json:"changeRate,omitempty" xml:"changeRate,omitempty"`
}

func (s ModifyIndexOnlineStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexOnlineStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexOnlineStrategyRequest) GetChangeRate() *int32 {
	return s.ChangeRate
}

func (s *ModifyIndexOnlineStrategyRequest) SetChangeRate(v int32) *ModifyIndexOnlineStrategyRequest {
	s.ChangeRate = &v
	return s
}

func (s *ModifyIndexOnlineStrategyRequest) Validate() error {
	return dara.Validate(s)
}
