// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterStopSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterStopSubscriptionRequest
	GetBalanceType() *string
}

type ModelRouterStopSubscriptionRequest struct {
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
}

func (s ModelRouterStopSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterStopSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterStopSubscriptionRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterStopSubscriptionRequest) SetBalanceType(v string) *ModelRouterStopSubscriptionRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterStopSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
