// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTransactionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *QueryAccountTransactionsRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *QueryAccountTransactionsRequest
	GetCreateTimeStart() *string
	SetPageNum(v int32) *QueryAccountTransactionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryAccountTransactionsRequest
	GetPageSize() *int32
	SetRecordID(v string) *QueryAccountTransactionsRequest
	GetRecordID() *string
	SetTransactionChannel(v string) *QueryAccountTransactionsRequest
	GetTransactionChannel() *string
	SetTransactionChannelSN(v string) *QueryAccountTransactionsRequest
	GetTransactionChannelSN() *string
	SetTransactionFlow(v string) *QueryAccountTransactionsRequest
	GetTransactionFlow() *string
	SetTransactionNumber(v string) *QueryAccountTransactionsRequest
	GetTransactionNumber() *string
	SetTransactionType(v string) *QueryAccountTransactionsRequest
	GetTransactionType() *string
}

type QueryAccountTransactionsRequest struct {
	// The end time of the creation time range. If not specified, the default value is the current time. Example: 2018-01-01T00:00:00Z.
	//
	// example:
	//
	// 2020-03-06T01:55:00Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start time of the creation time range. If not specified, the default value is one month before the current time. Example: 2018-01-01T00:00:00Z.
	//
	// example:
	//
	// 2020-03-05T01:46:09Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The order or bill number.
	//
	// example:
	//
	// 20200302
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// The transaction channel. If you specify a transaction channel, the query returns results of the corresponding type. If no results exist, an empty result is returned. If you do not specify this parameter, all types are returned by default. Valid values:
	//
	// - AccountBalance: Account balance.
	//
	// - BankTransfer: Bank transfer.
	//
	// - Alipay: Alipay.
	//
	// - AntCreditPay: Ant Credit Pay.
	//
	// - OfflineRemittance: Offline remittance.
	//
	// - RegularBankCreditRefund: Credit-based payment refund.
	//
	// - CreditCard: Credit card.
	//
	// - MyBankCredit: MYbank Credit.
	//
	// - HuaxiaBankCInstallment: Huaxia Bank installment.
	//
	// - ApplePay: Apple Pay.
	//
	// example:
	//
	// AccountBalance
	TransactionChannel *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	// The payment reference No.
	//
	// example:
	//
	// 12342134
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	// The income and expense type. If you specify an income and expense type, the query returns results of the corresponding type. If no results exist, an empty result is returned. If you do not specify this parameter, all types are returned by default. Valid values:
	//
	// - Income: Income.
	//
	// - Expense: Expense.
	//
	// example:
	//
	// Income
	TransactionFlow *string `json:"TransactionFlow,omitempty" xml:"TransactionFlow,omitempty"`
	// The transaction number.
	//
	// example:
	//
	// 133314076
	TransactionNumber *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// The transaction type. If you specify a transaction type, the query returns results of the corresponding type. If no results exist, an empty result is returned. If you do not specify this parameter, all types are returned by default. Valid values:
	//
	// - Payment: Top-up.
	//
	// - Withdraw: Withdrawal.
	//
	// - Refund: Refund.
	//
	// - Consumption: Consumption.
	//
	// - Transfer: Transfer.
	//
	// - Adjust: Adjustment.
	//
	// example:
	//
	// Payment
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s QueryAccountTransactionsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionsRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *QueryAccountTransactionsRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *QueryAccountTransactionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAccountTransactionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAccountTransactionsRequest) GetRecordID() *string {
	return s.RecordID
}

func (s *QueryAccountTransactionsRequest) GetTransactionChannel() *string {
	return s.TransactionChannel
}

func (s *QueryAccountTransactionsRequest) GetTransactionChannelSN() *string {
	return s.TransactionChannelSN
}

func (s *QueryAccountTransactionsRequest) GetTransactionFlow() *string {
	return s.TransactionFlow
}

func (s *QueryAccountTransactionsRequest) GetTransactionNumber() *string {
	return s.TransactionNumber
}

func (s *QueryAccountTransactionsRequest) GetTransactionType() *string {
	return s.TransactionType
}

func (s *QueryAccountTransactionsRequest) SetCreateTimeEnd(v string) *QueryAccountTransactionsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetCreateTimeStart(v string) *QueryAccountTransactionsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetPageNum(v int32) *QueryAccountTransactionsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetPageSize(v int32) *QueryAccountTransactionsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetRecordID(v string) *QueryAccountTransactionsRequest {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionChannel(v string) *QueryAccountTransactionsRequest {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionChannelSN(v string) *QueryAccountTransactionsRequest {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionFlow(v string) *QueryAccountTransactionsRequest {
	s.TransactionFlow = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionNumber(v string) *QueryAccountTransactionsRequest {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionType(v string) *QueryAccountTransactionsRequest {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionsRequest) Validate() error {
	return dara.Validate(s)
}
