// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDnsGtmAccessStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *UpdateDnsGtmAccessStrategyResponseBody
	GetStrategyId() *string
}

type UpdateDnsGtmAccessStrategyResponseBody struct {
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

func (s UpdateDnsGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *UpdateDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) SetStrategyId(v string) *UpdateDnsGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
