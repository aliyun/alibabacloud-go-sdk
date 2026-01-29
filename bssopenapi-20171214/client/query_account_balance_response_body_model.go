// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountBalanceResponseBody
	GetCode() *string
	SetData(v *QueryAccountBalanceResponseBodyData) *QueryAccountBalanceResponseBody
	GetData() *QueryAccountBalanceResponseBodyData
	SetMessage(v string) *QueryAccountBalanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAccountBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountBalanceResponseBody
	GetSuccess() *bool
}

type QueryAccountBalanceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryAccountBalanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16176743-6DC7-4CB3-BB25-A13982D8DFAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountBalanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountBalanceResponseBody) GetData() *QueryAccountBalanceResponseBodyData {
	return s.Data
}

func (s *QueryAccountBalanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountBalanceResponseBody) SetCode(v string) *QueryAccountBalanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetData(v *QueryAccountBalanceResponseBodyData) *QueryAccountBalanceResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetMessage(v string) *QueryAccountBalanceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetRequestId(v string) *QueryAccountBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetSuccess(v bool) *QueryAccountBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountBalanceResponseBodyData struct {
	// The available balance of the account.
	//
	// example:
	//
	// 10000.00
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// The available balance in cash.
	//
	// example:
	//
	// 10000.00
	AvailableCashAmount *string `json:"AvailableCashAmount,omitempty" xml:"AvailableCashAmount,omitempty"`
	// The credit balance of the account.
	//
	// example:
	//
	// 0.00
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// The type of the currency. Valid values:
	//
	// 	- CNY: Chinese Yuan
	//
	// 	- USD: US dollar
	//
	// 	- JPY: Japanese Yen
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The credit line controlled by MYbank.
	//
	// example:
	//
	// 0.00
	MybankCreditAmount *string `json:"MybankCreditAmount,omitempty" xml:"MybankCreditAmount,omitempty"`
	// The quota limit for eco customers.
	//
	// example:
	//
	// 10000.00
	QuotaLimit *string `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
}

func (s QueryAccountBalanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBalanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountBalanceResponseBodyData) GetAvailableAmount() *string {
	return s.AvailableAmount
}

func (s *QueryAccountBalanceResponseBodyData) GetAvailableCashAmount() *string {
	return s.AvailableCashAmount
}

func (s *QueryAccountBalanceResponseBodyData) GetCreditAmount() *string {
	return s.CreditAmount
}

func (s *QueryAccountBalanceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryAccountBalanceResponseBodyData) GetMybankCreditAmount() *string {
	return s.MybankCreditAmount
}

func (s *QueryAccountBalanceResponseBodyData) GetQuotaLimit() *string {
	return s.QuotaLimit
}

func (s *QueryAccountBalanceResponseBodyData) SetAvailableAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.AvailableAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetAvailableCashAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.AvailableCashAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetCreditAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.CreditAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetCurrency(v string) *QueryAccountBalanceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetMybankCreditAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.MybankCreditAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetQuotaLimit(v string) *QueryAccountBalanceResponseBodyData {
	s.QuotaLimit = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
