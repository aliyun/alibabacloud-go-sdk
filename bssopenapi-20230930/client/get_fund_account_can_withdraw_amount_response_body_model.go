// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanWithdrawAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanOriginalWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetCanOriginalWithdrawAmount() *string
	SetCanWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetCanWithdrawAmount() *string
	SetCannotOriginalWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetCannotOriginalWithdrawAmount() *string
	SetCashAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetCashAmount() *string
	SetCreditMemoAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetCreditMemoAmount() *string
	SetCurrentMonthUnclearedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetCurrentMonthUnclearedAmount() *string
	SetHistoryMonthUnclearedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetHistoryMonthUnclearedAmount() *string
	SetMetadata(v interface{}) *GetFundAccountCanWithdrawAmountResponseBody
	GetMetadata() interface{}
	SetPayAsYouGoReversedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetPayAsYouGoReversedAmount() *string
	SetRequestId(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetRequestId() *string
	SetTransferAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody
	GetTransferAmount() *string
}

type GetFundAccountCanWithdrawAmountResponseBody struct {
	// Amount withdrawable via original payment method
	//
	// example:
	//
	// 400
	CanOriginalWithdrawAmount *string `json:"CanOriginalWithdrawAmount,omitempty" xml:"CanOriginalWithdrawAmount,omitempty"`
	// Withdrawable amount = Cash balance - Transfer ledger amount - (Outstanding invoice amount + Historical months unsettled amount + Current month unsettled amount - Transfer ledger amount) - Pay-as-you-go reserved amount
	//
	// example:
	//
	// 500
	CanWithdrawAmount *string `json:"CanWithdrawAmount,omitempty" xml:"CanWithdrawAmount,omitempty"`
	// Amount not withdrawable via original payment method
	//
	// example:
	//
	// 100
	CannotOriginalWithdrawAmount *string `json:"CannotOriginalWithdrawAmount,omitempty" xml:"CannotOriginalWithdrawAmount,omitempty"`
	// Cash balance
	//
	// example:
	//
	// 1000
	CashAmount *string `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// Outstanding invoice amount
	//
	// example:
	//
	// 100
	CreditMemoAmount *string `json:"CreditMemoAmount,omitempty" xml:"CreditMemoAmount,omitempty"`
	// Current month unsettled amount
	//
	// example:
	//
	// 200
	CurrentMonthUnclearedAmount *string `json:"CurrentMonthUnclearedAmount,omitempty" xml:"CurrentMonthUnclearedAmount,omitempty"`
	// Historical months unsettled amount
	//
	// example:
	//
	// 100
	HistoryMonthUnclearedAmount *string `json:"HistoryMonthUnclearedAmount,omitempty" xml:"HistoryMonthUnclearedAmount,omitempty"`
	// Response structure metadata
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Pay-as-you-go reserved amount
	//
	// example:
	//
	// 100
	PayAsYouGoReversedAmount *string `json:"PayAsYouGoReversedAmount,omitempty" xml:"PayAsYouGoReversedAmount,omitempty"`
	// Request ID
	//
	// example:
	//
	// DF58589C-A06C-4224-8615-7797E6474FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Transfer amount
	//
	// example:
	//
	// 0
	TransferAmount *string `json:"TransferAmount,omitempty" xml:"TransferAmount,omitempty"`
}

func (s GetFundAccountCanWithdrawAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanWithdrawAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetCanOriginalWithdrawAmount() *string {
	return s.CanOriginalWithdrawAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetCanWithdrawAmount() *string {
	return s.CanWithdrawAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetCannotOriginalWithdrawAmount() *string {
	return s.CannotOriginalWithdrawAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetCashAmount() *string {
	return s.CashAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetCreditMemoAmount() *string {
	return s.CreditMemoAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetCurrentMonthUnclearedAmount() *string {
	return s.CurrentMonthUnclearedAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetHistoryMonthUnclearedAmount() *string {
	return s.HistoryMonthUnclearedAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetPayAsYouGoReversedAmount() *string {
	return s.PayAsYouGoReversedAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) GetTransferAmount() *string {
	return s.TransferAmount
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCanOriginalWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CanOriginalWithdrawAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCanWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CanWithdrawAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCannotOriginalWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CannotOriginalWithdrawAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCashAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CashAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCreditMemoAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CreditMemoAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCurrentMonthUnclearedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CurrentMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetHistoryMonthUnclearedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.HistoryMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanWithdrawAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetPayAsYouGoReversedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.PayAsYouGoReversedAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetRequestId(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetTransferAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.TransferAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) Validate() error {
	return dara.Validate(s)
}
