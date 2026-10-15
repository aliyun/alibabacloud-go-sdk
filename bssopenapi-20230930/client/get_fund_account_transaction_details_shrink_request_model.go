// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountTransactionDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetBillNumber() *string
	SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetChannelTransactionNumber() *string
	SetCurrentPage(v int32) *GetFundAccountTransactionDetailsShrinkRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetEndTime() *int64
	SetFundAccountId(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetFundAccountId() *int64
	SetPageSize(v int32) *GetFundAccountTransactionDetailsShrinkRequest
	GetPageSize() *int32
	SetStartTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetStartTime() *int64
	SetTransactionChannelListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionChannelListShrink() *string
	SetTransactionDirection(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionDirection() *string
	SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionNumber() *int64
	SetTransactionType(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionType() *string
	SetTransactionTypeListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionTypeListShrink() *string
}

type GetFundAccountTransactionDetailsShrinkRequest struct {
	// The transaction order number.
	//
	// example:
	//
	// 2023212312321
	BillNumber *string `json:"BillNumber,omitempty" xml:"BillNumber,omitempty"`
	// The payment reference No.
	//
	// example:
	//
	// 20250312334312322
	ChannelTransactionNumber *string `json:"ChannelTransactionNumber,omitempty" xml:"ChannelTransactionNumber,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end timestamp for the query. Unit: milliseconds.
	//
	// example:
	//
	// 1735664561000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the fund account. If you do not specify this parameter, the ID of the account owned by the current user is used by default.
	//
	// example:
	//
	// 123221232
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// The number of entries per page. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start timestamp for the query. Unit: milliseconds.
	//
	// example:
	//
	// 1735664461000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of transaction channels. If you specify this parameter, transactions are queried by channel. If you do not specify this parameter, all channels are queried by default. Valid values:
	//
	// - ACCT_CASH: User balance.
	//
	// - ALIPAY: Alipay.
	//
	// - ALIPAY_ZHIFUTONG: Alipay Zhifutong.
	//
	// - OFFLINE_REMIT: Corporate remittance.
	//
	// - REFUND: Credit limit refund.
	//
	// - UNION_PAY_BANK: Online banking.
	//
	// - CREDIT_CARD: Credit card. Available only on the international site.
	//
	// - AYPAL: PayPal. Available only on the international site.
	TransactionChannelListShrink *string `json:"TransactionChannelList,omitempty" xml:"TransactionChannelList,omitempty"`
	// The direction of the transaction.
	//
	// example:
	//
	// IN
	TransactionDirection *string `json:"TransactionDirection,omitempty" xml:"TransactionDirection,omitempty"`
	// The transaction number.
	//
	// example:
	//
	// 543231231
	TransactionNumber *int64 `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// The transaction type. If you specify a transaction type, the query returns results of the corresponding type. If no results exist, an empty result is returned. If you do not specify this parameter, all types are returned by default. Valid values:
	//
	// - CHARGE: Top-up.
	//
	// - WITHDRAW: Withdrawal.
	//
	// - REFUND: Refund.
	//
	// - PAY: Payment.
	//
	// - TRANSFER: Transfer.
	//
	// - ADJUST: Adjustment.
	//
	// - PAY_FAILED: Refund for failed order.
	//
	// example:
	//
	// CHARGE
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	// The list of transaction types. You can specify multiple parameters for a parameter query.
	TransactionTypeListShrink *string `json:"TransactionTypeList,omitempty" xml:"TransactionTypeList,omitempty"`
}

func (s GetFundAccountTransactionDetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountTransactionDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetBillNumber() *string {
	return s.BillNumber
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetChannelTransactionNumber() *string {
	return s.ChannelTransactionNumber
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionChannelListShrink() *string {
	return s.TransactionChannelListShrink
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionDirection() *string {
	return s.TransactionDirection
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionNumber() *int64 {
	return s.TransactionNumber
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionType() *string {
	return s.TransactionType
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionTypeListShrink() *string {
	return s.TransactionTypeListShrink
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetBillNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetEndTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetPageSize(v int32) *GetFundAccountTransactionDetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetStartTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionChannelListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionChannelListShrink = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionType(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionTypeListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionTypeListShrink = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
