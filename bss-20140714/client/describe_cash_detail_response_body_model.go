// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCashDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAmountOwed(v string) *DescribeCashDetailResponseBody
	GetAmountOwed() *string
	SetAvailableCredit(v string) *DescribeCashDetailResponseBody
	GetAvailableCredit() *string
	SetBalanceAmount(v string) *DescribeCashDetailResponseBody
	GetBalanceAmount() *string
	SetCreditCardAmount(v string) *DescribeCashDetailResponseBody
	GetCreditCardAmount() *string
	SetCreditLimit(v string) *DescribeCashDetailResponseBody
	GetCreditLimit() *string
	SetEnableThresholdAlert(v string) *DescribeCashDetailResponseBody
	GetEnableThresholdAlert() *string
	SetFrozenAmount(v string) *DescribeCashDetailResponseBody
	GetFrozenAmount() *string
	SetMiniAlertThreshold(v int64) *DescribeCashDetailResponseBody
	GetMiniAlertThreshold() *int64
	SetRemmitanceAmount(v string) *DescribeCashDetailResponseBody
	GetRemmitanceAmount() *string
	SetRequestId(v string) *DescribeCashDetailResponseBody
	GetRequestId() *string
}

type DescribeCashDetailResponseBody struct {
	AmountOwed           *string `json:"AmountOwed,omitempty" xml:"AmountOwed,omitempty"`
	AvailableCredit      *string `json:"AvailableCredit,omitempty" xml:"AvailableCredit,omitempty"`
	BalanceAmount        *string `json:"BalanceAmount,omitempty" xml:"BalanceAmount,omitempty"`
	CreditCardAmount     *string `json:"CreditCardAmount,omitempty" xml:"CreditCardAmount,omitempty"`
	CreditLimit          *string `json:"CreditLimit,omitempty" xml:"CreditLimit,omitempty"`
	EnableThresholdAlert *string `json:"EnableThresholdAlert,omitempty" xml:"EnableThresholdAlert,omitempty"`
	FrozenAmount         *string `json:"FrozenAmount,omitempty" xml:"FrozenAmount,omitempty"`
	MiniAlertThreshold   *int64  `json:"MiniAlertThreshold,omitempty" xml:"MiniAlertThreshold,omitempty"`
	RemmitanceAmount     *string `json:"RemmitanceAmount,omitempty" xml:"RemmitanceAmount,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCashDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCashDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCashDetailResponseBody) GetAmountOwed() *string {
	return s.AmountOwed
}

func (s *DescribeCashDetailResponseBody) GetAvailableCredit() *string {
	return s.AvailableCredit
}

func (s *DescribeCashDetailResponseBody) GetBalanceAmount() *string {
	return s.BalanceAmount
}

func (s *DescribeCashDetailResponseBody) GetCreditCardAmount() *string {
	return s.CreditCardAmount
}

func (s *DescribeCashDetailResponseBody) GetCreditLimit() *string {
	return s.CreditLimit
}

func (s *DescribeCashDetailResponseBody) GetEnableThresholdAlert() *string {
	return s.EnableThresholdAlert
}

func (s *DescribeCashDetailResponseBody) GetFrozenAmount() *string {
	return s.FrozenAmount
}

func (s *DescribeCashDetailResponseBody) GetMiniAlertThreshold() *int64 {
	return s.MiniAlertThreshold
}

func (s *DescribeCashDetailResponseBody) GetRemmitanceAmount() *string {
	return s.RemmitanceAmount
}

func (s *DescribeCashDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCashDetailResponseBody) SetAmountOwed(v string) *DescribeCashDetailResponseBody {
	s.AmountOwed = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetAvailableCredit(v string) *DescribeCashDetailResponseBody {
	s.AvailableCredit = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetBalanceAmount(v string) *DescribeCashDetailResponseBody {
	s.BalanceAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetCreditCardAmount(v string) *DescribeCashDetailResponseBody {
	s.CreditCardAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetCreditLimit(v string) *DescribeCashDetailResponseBody {
	s.CreditLimit = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetEnableThresholdAlert(v string) *DescribeCashDetailResponseBody {
	s.EnableThresholdAlert = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetFrozenAmount(v string) *DescribeCashDetailResponseBody {
	s.FrozenAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetMiniAlertThreshold(v int64) *DescribeCashDetailResponseBody {
	s.MiniAlertThreshold = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetRemmitanceAmount(v string) *DescribeCashDetailResponseBody {
	s.RemmitanceAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetRequestId(v string) *DescribeCashDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCashDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
