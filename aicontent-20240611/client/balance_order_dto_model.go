// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBalanceOrderDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *BalanceOrderDTO
	GetAmount() *float64
	SetBalanceAfter(v float64) *BalanceOrderDTO
	GetBalanceAfter() *float64
	SetBalanceBefore(v float64) *BalanceOrderDTO
	GetBalanceBefore() *float64
	SetBalanceType(v string) *BalanceOrderDTO
	GetBalanceType() *string
	SetCreateTime(v string) *BalanceOrderDTO
	GetCreateTime() *string
	SetDirection(v string) *BalanceOrderDTO
	GetDirection() *string
	SetModelCode(v string) *BalanceOrderDTO
	GetModelCode() *string
	SetOperatorId(v string) *BalanceOrderDTO
	GetOperatorId() *string
	SetOrderId(v string) *BalanceOrderDTO
	GetOrderId() *string
	SetOrderType(v string) *BalanceOrderDTO
	GetOrderType() *string
	SetRemark(v string) *BalanceOrderDTO
	GetRemark() *string
	SetSource(v string) *BalanceOrderDTO
	GetSource() *string
	SetTotalAfter(v float64) *BalanceOrderDTO
	GetTotalAfter() *float64
	SetTotalBefore(v float64) *BalanceOrderDTO
	GetTotalBefore() *float64
}

type BalanceOrderDTO struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 600.00
	BalanceAfter *float64 `json:"balanceAfter,omitempty" xml:"balanceAfter,omitempty"`
	// example:
	//
	// 500.00
	BalanceBefore *float64 `json:"balanceBefore,omitempty" xml:"balanceBefore,omitempty"`
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// example:
	//
	// 2024-07-15T10:30:00+08:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// example:
	//
	// qwen-max
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// 12345
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// ord_20240715_abc123
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// recharge
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// example:
	//
	// Administrator manual recharge
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 1600.00
	TotalAfter *float64 `json:"totalAfter,omitempty" xml:"totalAfter,omitempty"`
	// example:
	//
	// 1500.00
	TotalBefore *float64 `json:"totalBefore,omitempty" xml:"totalBefore,omitempty"`
}

func (s BalanceOrderDTO) String() string {
	return dara.Prettify(s)
}

func (s BalanceOrderDTO) GoString() string {
	return s.String()
}

func (s *BalanceOrderDTO) GetAmount() *float64 {
	return s.Amount
}

func (s *BalanceOrderDTO) GetBalanceAfter() *float64 {
	return s.BalanceAfter
}

func (s *BalanceOrderDTO) GetBalanceBefore() *float64 {
	return s.BalanceBefore
}

func (s *BalanceOrderDTO) GetBalanceType() *string {
	return s.BalanceType
}

func (s *BalanceOrderDTO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BalanceOrderDTO) GetDirection() *string {
	return s.Direction
}

func (s *BalanceOrderDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *BalanceOrderDTO) GetOperatorId() *string {
	return s.OperatorId
}

func (s *BalanceOrderDTO) GetOrderId() *string {
	return s.OrderId
}

func (s *BalanceOrderDTO) GetOrderType() *string {
	return s.OrderType
}

func (s *BalanceOrderDTO) GetRemark() *string {
	return s.Remark
}

func (s *BalanceOrderDTO) GetSource() *string {
	return s.Source
}

func (s *BalanceOrderDTO) GetTotalAfter() *float64 {
	return s.TotalAfter
}

func (s *BalanceOrderDTO) GetTotalBefore() *float64 {
	return s.TotalBefore
}

func (s *BalanceOrderDTO) SetAmount(v float64) *BalanceOrderDTO {
	s.Amount = &v
	return s
}

func (s *BalanceOrderDTO) SetBalanceAfter(v float64) *BalanceOrderDTO {
	s.BalanceAfter = &v
	return s
}

func (s *BalanceOrderDTO) SetBalanceBefore(v float64) *BalanceOrderDTO {
	s.BalanceBefore = &v
	return s
}

func (s *BalanceOrderDTO) SetBalanceType(v string) *BalanceOrderDTO {
	s.BalanceType = &v
	return s
}

func (s *BalanceOrderDTO) SetCreateTime(v string) *BalanceOrderDTO {
	s.CreateTime = &v
	return s
}

func (s *BalanceOrderDTO) SetDirection(v string) *BalanceOrderDTO {
	s.Direction = &v
	return s
}

func (s *BalanceOrderDTO) SetModelCode(v string) *BalanceOrderDTO {
	s.ModelCode = &v
	return s
}

func (s *BalanceOrderDTO) SetOperatorId(v string) *BalanceOrderDTO {
	s.OperatorId = &v
	return s
}

func (s *BalanceOrderDTO) SetOrderId(v string) *BalanceOrderDTO {
	s.OrderId = &v
	return s
}

func (s *BalanceOrderDTO) SetOrderType(v string) *BalanceOrderDTO {
	s.OrderType = &v
	return s
}

func (s *BalanceOrderDTO) SetRemark(v string) *BalanceOrderDTO {
	s.Remark = &v
	return s
}

func (s *BalanceOrderDTO) SetSource(v string) *BalanceOrderDTO {
	s.Source = &v
	return s
}

func (s *BalanceOrderDTO) SetTotalAfter(v float64) *BalanceOrderDTO {
	s.TotalAfter = &v
	return s
}

func (s *BalanceOrderDTO) SetTotalBefore(v float64) *BalanceOrderDTO {
	s.TotalBefore = &v
	return s
}

func (s *BalanceOrderDTO) Validate() error {
	return dara.Validate(s)
}
