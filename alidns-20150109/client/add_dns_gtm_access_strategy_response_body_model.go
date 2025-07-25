// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDnsGtmAccessStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *AddDnsGtmAccessStrategyResponseBody
	GetStrategyId() *string
}

type AddDnsGtmAccessStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the access policy.
	//
	// example:
	//
	// testStrategyId1
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s AddDnsGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDnsGtmAccessStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *AddDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *AddDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsGtmAccessStrategyResponseBody) SetStrategyId(v string) *AddDnsGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *AddDnsGtmAccessStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
