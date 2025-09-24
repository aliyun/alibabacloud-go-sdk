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
	// The end of the creation time range to query. By default, the transactions in the last month are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2018-01-01T00:00:00Z.
	//
	// example:
	//
	// 2020-03-06T01:55:00Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The beginning of the creation time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2018-01-01T00:00:00Z.
	//
	// example:
	//
	// 2020-03-05T01:46:09Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The number of the page to return. Default value is 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value is 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the order or bill.
	//
	// example:
	//
	// 20200302
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// The transaction channel. If you specify one of the following transaction channels for this parameter, the results for the specified transaction channel are returned. If the transaction channel that you specify does not belong to the following transaction channels, no result is returned. If you leave this parameter empty, the results for all the following transaction channels are returned by default. Valid values:
	//
	// 	- AccountBalance
	//
	// 	- BankTransfer
	//
	// 	- Alipay
	//
	// 	- AntCreditPay
	//
	// 	- OfflineRemittance
	//
	// 	- RegularBankCreditRefund
	//
	// 	- CreditCard
	//
	// 	- MyBankCredit
	//
	// 	- HuaxiaBankCInstallment
	//
	// 	- ApplePay
	//
	// example:
	//
	// AccountBalance
	TransactionChannel *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	// The serial number of the transaction channel.
	//
	// example:
	//
	// 12342134
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	// The type of the transaction flow. If you specify one of the following types for this parameter, the results for the specified type are returned. If the type that you specify does not belong to the following types, no result is returned. If you leave this parameter empty, the results for the following two types are returned by default. Valid values:
	//
	// 	- Income
	//
	// 	- Expense
	//
	// example:
	//
	// Income
	TransactionFlow *string `json:"TransactionFlow,omitempty" xml:"TransactionFlow,omitempty"`
	// The number of the transaction.
	//
	// example:
	//
	// 133314076
	TransactionNumber *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// The type of the transaction. If you specify one of the following transaction types for this parameter, the results for the specified transaction type are returned. If the transaction type that you specify does not belong to the following types, no result is returned. If you leave this parameter empty, the results for all the following transaction types are returned by default. Valid values:
	//
	// 	- Payment
	//
	// 	- Withdraw
	//
	// 	- Refund
	//
	// 	- Consumption
	//
	// 	- Transfer
	//
	// 	- Adjust
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
