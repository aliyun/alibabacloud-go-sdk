// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterStopMemberSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterStopMemberSubscriptionRequest
	GetBalanceType() *string
}

type ModelRouterStopMemberSubscriptionRequest struct {
	// The balance type. Valid values: permanent, monthly.
	//
	// example:
	//
	// monthly
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
}

func (s ModelRouterStopMemberSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterStopMemberSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterStopMemberSubscriptionRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterStopMemberSubscriptionRequest) SetBalanceType(v string) *ModelRouterStopMemberSubscriptionRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
