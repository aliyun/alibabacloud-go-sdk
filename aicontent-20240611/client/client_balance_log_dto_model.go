// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientBalanceLogDTO interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceAfter(v float64) *ClientBalanceLogDTO
	GetBalanceAfter() *float64
	SetBalanceBefore(v float64) *ClientBalanceLogDTO
	GetBalanceBefore() *float64
	SetChangeAmount(v float64) *ClientBalanceLogDTO
	GetChangeAmount() *float64
	SetChangeType(v string) *ClientBalanceLogDTO
	GetChangeType() *string
	SetClientId(v int64) *ClientBalanceLogDTO
	GetClientId() *int64
	SetGmtCreate(v string) *ClientBalanceLogDTO
	GetGmtCreate() *string
	SetId(v int64) *ClientBalanceLogDTO
	GetId() *int64
	SetRemark(v string) *ClientBalanceLogDTO
	GetRemark() *string
}

type ClientBalanceLogDTO struct {
	// The new balance.
	//
	// example:
	//
	// 100.00
	BalanceAfter *float64 `json:"balanceAfter,omitempty" xml:"balanceAfter,omitempty"`
	// The previous balance.
	//
	// example:
	//
	// 90.00
	BalanceBefore *float64 `json:"balanceBefore,omitempty" xml:"balanceBefore,omitempty"`
	// The change amount. A positive value represents a recharge, and a negative value represents a deduction.
	//
	// example:
	//
	// 10.00
	ChangeAmount *float64 `json:"changeAmount,omitempty" xml:"changeAmount,omitempty"`
	// The type of change. Possible values are `auto_deduct`, `deduct`, or `recharge`.
	//
	// example:
	//
	// recharge
	ChangeType *string `json:"changeType,omitempty" xml:"changeType,omitempty"`
	// The client ID.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The creation time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The unique record ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Notes about the balance change.
	//
	// example:
	//
	// 充值
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ClientBalanceLogDTO) String() string {
	return dara.Prettify(s)
}

func (s ClientBalanceLogDTO) GoString() string {
	return s.String()
}

func (s *ClientBalanceLogDTO) GetBalanceAfter() *float64 {
	return s.BalanceAfter
}

func (s *ClientBalanceLogDTO) GetBalanceBefore() *float64 {
	return s.BalanceBefore
}

func (s *ClientBalanceLogDTO) GetChangeAmount() *float64 {
	return s.ChangeAmount
}

func (s *ClientBalanceLogDTO) GetChangeType() *string {
	return s.ChangeType
}

func (s *ClientBalanceLogDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *ClientBalanceLogDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ClientBalanceLogDTO) GetId() *int64 {
	return s.Id
}

func (s *ClientBalanceLogDTO) GetRemark() *string {
	return s.Remark
}

func (s *ClientBalanceLogDTO) SetBalanceAfter(v float64) *ClientBalanceLogDTO {
	s.BalanceAfter = &v
	return s
}

func (s *ClientBalanceLogDTO) SetBalanceBefore(v float64) *ClientBalanceLogDTO {
	s.BalanceBefore = &v
	return s
}

func (s *ClientBalanceLogDTO) SetChangeAmount(v float64) *ClientBalanceLogDTO {
	s.ChangeAmount = &v
	return s
}

func (s *ClientBalanceLogDTO) SetChangeType(v string) *ClientBalanceLogDTO {
	s.ChangeType = &v
	return s
}

func (s *ClientBalanceLogDTO) SetClientId(v int64) *ClientBalanceLogDTO {
	s.ClientId = &v
	return s
}

func (s *ClientBalanceLogDTO) SetGmtCreate(v string) *ClientBalanceLogDTO {
	s.GmtCreate = &v
	return s
}

func (s *ClientBalanceLogDTO) SetId(v int64) *ClientBalanceLogDTO {
	s.Id = &v
	return s
}

func (s *ClientBalanceLogDTO) SetRemark(v string) *ClientBalanceLogDTO {
	s.Remark = &v
	return s
}

func (s *ClientBalanceLogDTO) Validate() error {
	return dara.Validate(s)
}
