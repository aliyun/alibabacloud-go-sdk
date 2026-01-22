// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountAvailableAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetAvailableAmount() *string
	SetAvailableCreditAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetAvailableCreditAmount() *string
	SetBankAcceptanceAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetBankAcceptanceAmount() *string
	SetCashAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetCashAmount() *string
	SetCreditAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetCreditAmount() *string
	SetCreditRefundAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetCreditRefundAmount() *string
	SetCreditUser(v bool) *GetFundAccountAvailableAmountResponseBody
	GetCreditUser() *bool
	SetCurrency(v string) *GetFundAccountAvailableAmountResponseBody
	GetCurrency() *string
	SetCurrentMonthUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetCurrentMonthUnclearedAmount() *string
	SetExtendLedgerList(v []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList) *GetFundAccountAvailableAmountResponseBody
	GetExtendLedgerList() []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList
	SetFundAccountId(v string) *GetFundAccountAvailableAmountResponseBody
	GetFundAccountId() *string
	SetFundAccountOwnerAccountId(v string) *GetFundAccountAvailableAmountResponseBody
	GetFundAccountOwnerAccountId() *string
	SetFundAccountStatus(v string) *GetFundAccountAvailableAmountResponseBody
	GetFundAccountStatus() *string
	SetFundAccountType(v string) *GetFundAccountAvailableAmountResponseBody
	GetFundAccountType() *string
	SetHistoryMonthUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetHistoryMonthUnclearedAmount() *string
	SetMetadata(v interface{}) *GetFundAccountAvailableAmountResponseBody
	GetMetadata() interface{}
	SetNegativeBillAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetNegativeBillAmount() *string
	SetOriginalCashAmountList(v []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) *GetFundAccountAvailableAmountResponseBody
	GetOriginalCashAmountList() []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList
	SetQuotaAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetQuotaAmount() *string
	SetQuotaConsumedAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetQuotaConsumedAmount() *string
	SetRequestId(v string) *GetFundAccountAvailableAmountResponseBody
	GetRequestId() *string
	SetUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody
	GetUnclearedAmount() *string
}

type GetFundAccountAvailableAmountResponseBody struct {
	// example:
	//
	// 100
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// example:
	//
	// 50
	AvailableCreditAmount *string `json:"AvailableCreditAmount,omitempty" xml:"AvailableCreditAmount,omitempty"`
	// example:
	//
	// 0
	BankAcceptanceAmount *string `json:"BankAcceptanceAmount,omitempty" xml:"BankAcceptanceAmount,omitempty"`
	// example:
	//
	// 50
	CashAmount *string `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// example:
	//
	// 100
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// example:
	//
	// 0
	CreditRefundAmount *string `json:"CreditRefundAmount,omitempty" xml:"CreditRefundAmount,omitempty"`
	CreditUser         *bool   `json:"CreditUser,omitempty" xml:"CreditUser,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 20
	CurrentMonthUnclearedAmount *string                                                      `json:"CurrentMonthUnclearedAmount,omitempty" xml:"CurrentMonthUnclearedAmount,omitempty"`
	ExtendLedgerList            []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList `json:"ExtendLedgerList,omitempty" xml:"ExtendLedgerList,omitempty" type:"Repeated"`
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 1344312434
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// valid
	FundAccountStatus *string `json:"FundAccountStatus,omitempty" xml:"FundAccountStatus,omitempty"`
	// example:
	//
	// REDIRECT_USER
	FundAccountType *string `json:"FundAccountType,omitempty" xml:"FundAccountType,omitempty"`
	// example:
	//
	// 30
	HistoryMonthUnclearedAmount *string `json:"HistoryMonthUnclearedAmount,omitempty" xml:"HistoryMonthUnclearedAmount,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 0
	NegativeBillAmount     *string                                                            `json:"NegativeBillAmount,omitempty" xml:"NegativeBillAmount,omitempty"`
	OriginalCashAmountList []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList `json:"OriginalCashAmountList,omitempty" xml:"OriginalCashAmountList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	QuotaAmount *string `json:"QuotaAmount,omitempty" xml:"QuotaAmount,omitempty"`
	// example:
	//
	// 10
	QuotaConsumedAmount *string `json:"QuotaConsumedAmount,omitempty" xml:"QuotaConsumedAmount,omitempty"`
	// example:
	//
	// F96A2D13-7509-5DF9-A60E-E7E3A3CB68E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	UnclearedAmount *string `json:"UnclearedAmount,omitempty" xml:"UnclearedAmount,omitempty"`
}

func (s GetFundAccountAvailableAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponseBody) GetAvailableAmount() *string {
	return s.AvailableAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetAvailableCreditAmount() *string {
	return s.AvailableCreditAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetBankAcceptanceAmount() *string {
	return s.BankAcceptanceAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetCashAmount() *string {
	return s.CashAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetCreditAmount() *string {
	return s.CreditAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetCreditRefundAmount() *string {
	return s.CreditRefundAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetCreditUser() *bool {
	return s.CreditUser
}

func (s *GetFundAccountAvailableAmountResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountAvailableAmountResponseBody) GetCurrentMonthUnclearedAmount() *string {
	return s.CurrentMonthUnclearedAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetExtendLedgerList() []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	return s.ExtendLedgerList
}

func (s *GetFundAccountAvailableAmountResponseBody) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *GetFundAccountAvailableAmountResponseBody) GetFundAccountOwnerAccountId() *string {
	return s.FundAccountOwnerAccountId
}

func (s *GetFundAccountAvailableAmountResponseBody) GetFundAccountStatus() *string {
	return s.FundAccountStatus
}

func (s *GetFundAccountAvailableAmountResponseBody) GetFundAccountType() *string {
	return s.FundAccountType
}

func (s *GetFundAccountAvailableAmountResponseBody) GetHistoryMonthUnclearedAmount() *string {
	return s.HistoryMonthUnclearedAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountAvailableAmountResponseBody) GetNegativeBillAmount() *string {
	return s.NegativeBillAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetOriginalCashAmountList() []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList {
	return s.OriginalCashAmountList
}

func (s *GetFundAccountAvailableAmountResponseBody) GetQuotaAmount() *string {
	return s.QuotaAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetQuotaConsumedAmount() *string {
	return s.QuotaConsumedAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountAvailableAmountResponseBody) GetUnclearedAmount() *string {
	return s.UnclearedAmount
}

func (s *GetFundAccountAvailableAmountResponseBody) SetAvailableAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.AvailableAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetAvailableCreditAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.AvailableCreditAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetBankAcceptanceAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.BankAcceptanceAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCashAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CashAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCreditAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CreditAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCreditRefundAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CreditRefundAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCreditUser(v bool) *GetFundAccountAvailableAmountResponseBody {
	s.CreditUser = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCurrency(v string) *GetFundAccountAvailableAmountResponseBody {
	s.Currency = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCurrentMonthUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CurrentMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetExtendLedgerList(v []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList) *GetFundAccountAvailableAmountResponseBody {
	s.ExtendLedgerList = v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountId(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountOwnerAccountId(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountStatus(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountStatus = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountType(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountType = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetHistoryMonthUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.HistoryMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountAvailableAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetNegativeBillAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.NegativeBillAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetOriginalCashAmountList(v []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) *GetFundAccountAvailableAmountResponseBody {
	s.OriginalCashAmountList = v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetQuotaAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.QuotaAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetQuotaConsumedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.QuotaConsumedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetRequestId(v string) *GetFundAccountAvailableAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.UnclearedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) Validate() error {
	if s.ExtendLedgerList != nil {
		for _, item := range s.ExtendLedgerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OriginalCashAmountList != nil {
		for _, item := range s.OriginalCashAmountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFundAccountAvailableAmountResponseBodyExtendLedgerList struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 应付对冲账本
	LedgerName *string `json:"LedgerName,omitempty" xml:"LedgerName,omitempty"`
	// example:
	//
	// 50
	OriginalAmount *string `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
}

func (s GetFundAccountAvailableAmountResponseBodyExtendLedgerList) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponseBodyExtendLedgerList) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) GetLedgerName() *string {
	return s.LedgerName
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) GetOriginalAmount() *string {
	return s.OriginalAmount
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) SetCurrency(v string) *GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	s.Currency = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) SetLedgerName(v string) *GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	s.LedgerName = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) SetOriginalAmount(v string) *GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	s.OriginalAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) Validate() error {
	return dara.Validate(s)
}

type GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList struct {
	// example:
	//
	// 10
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
}

func (s GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) GetAmount() *string {
	return s.Amount
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) SetAmount(v string) *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList {
	s.Amount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) SetCurrency(v string) *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList {
	s.Currency = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) Validate() error {
	return dara.Validate(s)
}
