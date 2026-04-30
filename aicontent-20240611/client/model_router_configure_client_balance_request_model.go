// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterConfigureClientBalanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterConfigureClientBalanceRequest
	GetBalanceType() *string
	SetEnableBalance(v bool) *ModelRouterConfigureClientBalanceRequest
	GetEnableBalance() *bool
	SetInitialBalance(v float64) *ModelRouterConfigureClientBalanceRequest
	GetInitialBalance() *float64
}

type ModelRouterConfigureClientBalanceRequest struct {
	// example:
	//
	// amount
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// example:
	//
	// true
	EnableBalance *bool `json:"enableBalance,omitempty" xml:"enableBalance,omitempty"`
	// example:
	//
	// 100.00
	InitialBalance *float64 `json:"initialBalance,omitempty" xml:"initialBalance,omitempty"`
}

func (s ModelRouterConfigureClientBalanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterConfigureClientBalanceRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterConfigureClientBalanceRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterConfigureClientBalanceRequest) GetEnableBalance() *bool {
	return s.EnableBalance
}

func (s *ModelRouterConfigureClientBalanceRequest) GetInitialBalance() *float64 {
	return s.InitialBalance
}

func (s *ModelRouterConfigureClientBalanceRequest) SetBalanceType(v string) *ModelRouterConfigureClientBalanceRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceRequest) SetEnableBalance(v bool) *ModelRouterConfigureClientBalanceRequest {
	s.EnableBalance = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceRequest) SetInitialBalance(v float64) *ModelRouterConfigureClientBalanceRequest {
	s.InitialBalance = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceRequest) Validate() error {
	return dara.Validate(s)
}
