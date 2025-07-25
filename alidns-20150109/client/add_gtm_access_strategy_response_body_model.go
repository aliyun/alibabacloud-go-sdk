// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddGtmAccessStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *AddGtmAccessStrategyResponseBody
	GetStrategyId() *string
}

type AddGtmAccessStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the access policy created.
	//
	// example:
	//
	// strategyid
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s AddGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGtmAccessStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *AddGtmAccessStrategyResponseBody) SetRequestId(v string) *AddGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmAccessStrategyResponseBody) SetStrategyId(v string) *AddGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *AddGtmAccessStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
