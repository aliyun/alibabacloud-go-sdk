// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillOrderEntryDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *BillOrderEntryDTO
	GetAmount() *float64
	SetBalanceAfter(v float64) *BillOrderEntryDTO
	GetBalanceAfter() *float64
	SetBalanceBefore(v float64) *BillOrderEntryDTO
	GetBalanceBefore() *float64
	SetBalanceType(v string) *BillOrderEntryDTO
	GetBalanceType() *string
	SetCreateTime(v string) *BillOrderEntryDTO
	GetCreateTime() *string
	SetDirection(v string) *BillOrderEntryDTO
	GetDirection() *string
	SetModelCode(v string) *BillOrderEntryDTO
	GetModelCode() *string
	SetOperatorId(v string) *BillOrderEntryDTO
	GetOperatorId() *string
	SetOrderId(v string) *BillOrderEntryDTO
	GetOrderId() *string
	SetOrderType(v string) *BillOrderEntryDTO
	GetOrderType() *string
	SetRemark(v string) *BillOrderEntryDTO
	GetRemark() *string
	SetSource(v string) *BillOrderEntryDTO
	GetSource() *string
	SetTotalAfter(v float64) *BillOrderEntryDTO
	GetTotalAfter() *float64
	SetTotalBefore(v float64) *BillOrderEntryDTO
	GetTotalBefore() *float64
}

type BillOrderEntryDTO struct {
	// The change amount.
	//
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The balance after the change.
	//
	// example:
	//
	// 100.00
	BalanceAfter *float64 `json:"balanceAfter,omitempty" xml:"balanceAfter,omitempty"`
	// The balance before the change.
	//
	// example:
	//
	// 0.00
	BalanceBefore *float64 `json:"balanceBefore,omitempty" xml:"balanceBefore,omitempty"`
	// The balance type: permanent/monthly.
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The change direction: in/out.
	//
	// example:
	//
	// in
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The model code (available when gateway billing deduction is applied).
	//
	// example:
	//
	// qwen-max
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 1
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// The change record ID.
	//
	// example:
	//
	// ord_xxxxxxxx
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The change type.
	//
	// example:
	//
	// recharge
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// The remark.
	//
	// example:
	//
	// Top-up
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The source.
	//
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The total balance after the change.
	//
	// example:
	//
	// 100.00
	TotalAfter *float64 `json:"totalAfter,omitempty" xml:"totalAfter,omitempty"`
	// The total balance before the change.
	//
	// example:
	//
	// 0.00
	TotalBefore *float64 `json:"totalBefore,omitempty" xml:"totalBefore,omitempty"`
}

func (s BillOrderEntryDTO) String() string {
	return dara.Prettify(s)
}

func (s BillOrderEntryDTO) GoString() string {
	return s.String()
}

func (s *BillOrderEntryDTO) GetAmount() *float64 {
	return s.Amount
}

func (s *BillOrderEntryDTO) GetBalanceAfter() *float64 {
	return s.BalanceAfter
}

func (s *BillOrderEntryDTO) GetBalanceBefore() *float64 {
	return s.BalanceBefore
}

func (s *BillOrderEntryDTO) GetBalanceType() *string {
	return s.BalanceType
}

func (s *BillOrderEntryDTO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BillOrderEntryDTO) GetDirection() *string {
	return s.Direction
}

func (s *BillOrderEntryDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *BillOrderEntryDTO) GetOperatorId() *string {
	return s.OperatorId
}

func (s *BillOrderEntryDTO) GetOrderId() *string {
	return s.OrderId
}

func (s *BillOrderEntryDTO) GetOrderType() *string {
	return s.OrderType
}

func (s *BillOrderEntryDTO) GetRemark() *string {
	return s.Remark
}

func (s *BillOrderEntryDTO) GetSource() *string {
	return s.Source
}

func (s *BillOrderEntryDTO) GetTotalAfter() *float64 {
	return s.TotalAfter
}

func (s *BillOrderEntryDTO) GetTotalBefore() *float64 {
	return s.TotalBefore
}

func (s *BillOrderEntryDTO) SetAmount(v float64) *BillOrderEntryDTO {
	s.Amount = &v
	return s
}

func (s *BillOrderEntryDTO) SetBalanceAfter(v float64) *BillOrderEntryDTO {
	s.BalanceAfter = &v
	return s
}

func (s *BillOrderEntryDTO) SetBalanceBefore(v float64) *BillOrderEntryDTO {
	s.BalanceBefore = &v
	return s
}

func (s *BillOrderEntryDTO) SetBalanceType(v string) *BillOrderEntryDTO {
	s.BalanceType = &v
	return s
}

func (s *BillOrderEntryDTO) SetCreateTime(v string) *BillOrderEntryDTO {
	s.CreateTime = &v
	return s
}

func (s *BillOrderEntryDTO) SetDirection(v string) *BillOrderEntryDTO {
	s.Direction = &v
	return s
}

func (s *BillOrderEntryDTO) SetModelCode(v string) *BillOrderEntryDTO {
	s.ModelCode = &v
	return s
}

func (s *BillOrderEntryDTO) SetOperatorId(v string) *BillOrderEntryDTO {
	s.OperatorId = &v
	return s
}

func (s *BillOrderEntryDTO) SetOrderId(v string) *BillOrderEntryDTO {
	s.OrderId = &v
	return s
}

func (s *BillOrderEntryDTO) SetOrderType(v string) *BillOrderEntryDTO {
	s.OrderType = &v
	return s
}

func (s *BillOrderEntryDTO) SetRemark(v string) *BillOrderEntryDTO {
	s.Remark = &v
	return s
}

func (s *BillOrderEntryDTO) SetSource(v string) *BillOrderEntryDTO {
	s.Source = &v
	return s
}

func (s *BillOrderEntryDTO) SetTotalAfter(v float64) *BillOrderEntryDTO {
	s.TotalAfter = &v
	return s
}

func (s *BillOrderEntryDTO) SetTotalBefore(v float64) *BillOrderEntryDTO {
	s.TotalBefore = &v
	return s
}

func (s *BillOrderEntryDTO) Validate() error {
	return dara.Validate(s)
}
