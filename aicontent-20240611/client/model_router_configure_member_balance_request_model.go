// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterConfigureMemberBalanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterConfigureMemberBalanceRequest
	GetBalanceType() *string
	SetEnableBalance(v bool) *ModelRouterConfigureMemberBalanceRequest
	GetEnableBalance() *bool
	SetInitialBalance(v float64) *ModelRouterConfigureMemberBalanceRequest
	GetInitialBalance() *float64
}

type ModelRouterConfigureMemberBalanceRequest struct {
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
	// 0
	InitialBalance *float64 `json:"initialBalance,omitempty" xml:"initialBalance,omitempty"`
}

func (s ModelRouterConfigureMemberBalanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterConfigureMemberBalanceRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterConfigureMemberBalanceRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterConfigureMemberBalanceRequest) GetEnableBalance() *bool {
	return s.EnableBalance
}

func (s *ModelRouterConfigureMemberBalanceRequest) GetInitialBalance() *float64 {
	return s.InitialBalance
}

func (s *ModelRouterConfigureMemberBalanceRequest) SetBalanceType(v string) *ModelRouterConfigureMemberBalanceRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceRequest) SetEnableBalance(v bool) *ModelRouterConfigureMemberBalanceRequest {
	s.EnableBalance = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceRequest) SetInitialBalance(v float64) *ModelRouterConfigureMemberBalanceRequest {
	s.InitialBalance = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceRequest) Validate() error {
	return dara.Validate(s)
}
