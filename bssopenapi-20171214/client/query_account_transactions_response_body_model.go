// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTransactionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountTransactionsResponseBody
	GetCode() *string
	SetData(v *QueryAccountTransactionsResponseBodyData) *QueryAccountTransactionsResponseBody
	GetData() *QueryAccountTransactionsResponseBodyData
	SetMessage(v string) *QueryAccountTransactionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAccountTransactionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountTransactionsResponseBody
	GetSuccess() *bool
}

type QueryAccountTransactionsResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryAccountTransactionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8634E02D-0942-4B1D-8295-5352FE9A1F39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountTransactionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountTransactionsResponseBody) GetData() *QueryAccountTransactionsResponseBodyData {
	return s.Data
}

func (s *QueryAccountTransactionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountTransactionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountTransactionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountTransactionsResponseBody) SetCode(v string) *QueryAccountTransactionsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetData(v *QueryAccountTransactionsResponseBodyData) *QueryAccountTransactionsResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetMessage(v string) *QueryAccountTransactionsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetRequestId(v string) *QueryAccountTransactionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetSuccess(v bool) *QueryAccountTransactionsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountTransactionsResponseBodyData struct {
	// The name of your Alibaba Cloud account.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The information about transactions.
	AccountTransactionsList *QueryAccountTransactionsResponseBodyDataAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAccountTransactionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryAccountTransactionsResponseBodyData) GetAccountTransactionsList() *QueryAccountTransactionsResponseBodyDataAccountTransactionsList {
	return s.AccountTransactionsList
}

func (s *QueryAccountTransactionsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAccountTransactionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAccountTransactionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryAccountTransactionsResponseBodyData) SetAccountName(v string) *QueryAccountTransactionsResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetAccountTransactionsList(v *QueryAccountTransactionsResponseBodyDataAccountTransactionsList) *QueryAccountTransactionsResponseBodyData {
	s.AccountTransactionsList = v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetPageNum(v int32) *QueryAccountTransactionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetPageSize(v int32) *QueryAccountTransactionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetTotalCount(v int32) *QueryAccountTransactionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) Validate() error {
	if s.AccountTransactionsList != nil {
		if err := s.AccountTransactionsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountTransactionsResponseBodyDataAccountTransactionsList struct {
	AccountTransactionsList []*QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Repeated"`
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsList) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsList) GetAccountTransactionsList() []*QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	return s.AccountTransactionsList
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsList) SetAccountTransactionsList(v []*QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) *QueryAccountTransactionsResponseBodyDataAccountTransactionsList {
	s.AccountTransactionsList = v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsList) Validate() error {
	if s.AccountTransactionsList != nil {
		for _, item := range s.AccountTransactionsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList struct {
	// The amount.
	//
	// example:
	//
	// 0
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The balance of the account.
	//
	// example:
	//
	// 0
	Balance *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The type of transaction payment. Valid values:
	//
	// 	- Cash: pay for the transaction in cash.
	//
	// 	- Deposit: pay for the transaction with deposit.
	//
	// 	- RegularBankCreditRefund: pay for the transaction with credit refund controlled by a bank.
	//
	// 	- DirectPay: directly pay for the transaction.
	//
	// example:
	//
	// Cash
	FundType *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	// The number of the order or bill.
	//
	// example:
	//
	// 2020030242
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// The remarks on the transaction.
	//
	// example:
	//
	// NAT_GW
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The transaction account. For example, the account is a recharge account in Alipay or a transfer account.
	//
	// example:
	//
	// 213562146
	TransactionAccount *string `json:"TransactionAccount,omitempty" xml:"TransactionAccount,omitempty"`
	// The transaction channel.
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
	// 1234354325
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	// The type of the transaction flow.
	//
	// 	- Income
	//
	// 	- Expense
	//
	// example:
	//
	// Expense
	TransactionFlow *string `json:"TransactionFlow,omitempty" xml:"TransactionFlow,omitempty"`
	// The number of the transaction.
	//
	// example:
	//
	// 43342334
	TransactionNumber *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// The time when the transaction was made.
	//
	// example:
	//
	// 2020-03-10T02:03:20Z
	TransactionTime *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	// The type of the transaction.
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
	// Consumption
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetAmount() *string {
	return s.Amount
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetBalance() *string {
	return s.Balance
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetFundType() *string {
	return s.FundType
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetRecordID() *string {
	return s.RecordID
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetRemarks() *string {
	return s.Remarks
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionAccount() *string {
	return s.TransactionAccount
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionChannel() *string {
	return s.TransactionChannel
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionChannelSN() *string {
	return s.TransactionChannelSN
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionFlow() *string {
	return s.TransactionFlow
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionNumber() *string {
	return s.TransactionNumber
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionTime() *string {
	return s.TransactionTime
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionType() *string {
	return s.TransactionType
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetAmount(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Amount = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBalance(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Balance = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBillingCycle(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetFundType(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.FundType = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRecordID(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRemarks(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Remarks = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionAccount(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionAccount = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannel(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannelSN(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionFlow(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionFlow = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionNumber(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionTime(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionTime = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionType(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) Validate() error {
	return dara.Validate(s)
}
