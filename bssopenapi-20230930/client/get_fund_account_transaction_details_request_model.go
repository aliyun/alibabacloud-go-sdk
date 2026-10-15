// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountTransactionDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillNumber(v string) *GetFundAccountTransactionDetailsRequest
	GetBillNumber() *string
	SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsRequest
	GetChannelTransactionNumber() *string
	SetCurrentPage(v int32) *GetFundAccountTransactionDetailsRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetFundAccountTransactionDetailsRequest
	GetEndTime() *int64
	SetFundAccountId(v int64) *GetFundAccountTransactionDetailsRequest
	GetFundAccountId() *int64
	SetPageSize(v int32) *GetFundAccountTransactionDetailsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *GetFundAccountTransactionDetailsRequest
	GetStartTime() *int64
	SetTransactionChannelList(v []*string) *GetFundAccountTransactionDetailsRequest
	GetTransactionChannelList() []*string
	SetTransactionDirection(v string) *GetFundAccountTransactionDetailsRequest
	GetTransactionDirection() *string
	SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsRequest
	GetTransactionNumber() *int64
	SetTransactionType(v string) *GetFundAccountTransactionDetailsRequest
	GetTransactionType() *string
	SetTransactionTypeList(v []*string) *GetFundAccountTransactionDetailsRequest
	GetTransactionTypeList() []*string
}

type GetFundAccountTransactionDetailsRequest struct {
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
	TransactionChannelList []*string `json:"TransactionChannelList,omitempty" xml:"TransactionChannelList,omitempty" type:"Repeated"`
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
	TransactionTypeList []*string `json:"TransactionTypeList,omitempty" xml:"TransactionTypeList,omitempty" type:"Repeated"`
}

func (s GetFundAccountTransactionDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountTransactionDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsRequest) GetBillNumber() *string {
	return s.BillNumber
}

func (s *GetFundAccountTransactionDetailsRequest) GetChannelTransactionNumber() *string {
	return s.ChannelTransactionNumber
}

func (s *GetFundAccountTransactionDetailsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetFundAccountTransactionDetailsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetFundAccountTransactionDetailsRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountTransactionDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFundAccountTransactionDetailsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetFundAccountTransactionDetailsRequest) GetTransactionChannelList() []*string {
	return s.TransactionChannelList
}

func (s *GetFundAccountTransactionDetailsRequest) GetTransactionDirection() *string {
	return s.TransactionDirection
}

func (s *GetFundAccountTransactionDetailsRequest) GetTransactionNumber() *int64 {
	return s.TransactionNumber
}

func (s *GetFundAccountTransactionDetailsRequest) GetTransactionType() *string {
	return s.TransactionType
}

func (s *GetFundAccountTransactionDetailsRequest) GetTransactionTypeList() []*string {
	return s.TransactionTypeList
}

func (s *GetFundAccountTransactionDetailsRequest) SetBillNumber(v string) *GetFundAccountTransactionDetailsRequest {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsRequest {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetEndTime(v int64) *GetFundAccountTransactionDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetPageSize(v int32) *GetFundAccountTransactionDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetStartTime(v int64) *GetFundAccountTransactionDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionChannelList(v []*string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionChannelList = v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsRequest {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionType(v string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionTypeList(v []*string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionTypeList = v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) Validate() error {
	return dara.Validate(s)
}
