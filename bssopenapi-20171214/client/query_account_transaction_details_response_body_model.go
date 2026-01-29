// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTransactionDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountTransactionDetailsResponseBody
	GetCode() *string
	SetData(v *QueryAccountTransactionDetailsResponseBodyData) *QueryAccountTransactionDetailsResponseBody
	GetData() *QueryAccountTransactionDetailsResponseBodyData
	SetMessage(v string) *QueryAccountTransactionDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAccountTransactionDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountTransactionDetailsResponseBody
	GetSuccess() *bool
}

type QueryAccountTransactionDetailsResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryAccountTransactionDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// asadadad-edafafafaasd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountTransactionDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountTransactionDetailsResponseBody) GetData() *QueryAccountTransactionDetailsResponseBodyData {
	return s.Data
}

func (s *QueryAccountTransactionDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountTransactionDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountTransactionDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountTransactionDetailsResponseBody) SetCode(v string) *QueryAccountTransactionDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetData(v *QueryAccountTransactionDetailsResponseBodyData) *QueryAccountTransactionDetailsResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetMessage(v string) *QueryAccountTransactionDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetRequestId(v string) *QueryAccountTransactionDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetSuccess(v bool) *QueryAccountTransactionDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountTransactionDetailsResponseBodyData struct {
	// The name of the account.
	//
	// example:
	//
	// yidi
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The details of the transactions within the account.
	AccountTransactionsList *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Struct"`
	// This parameter is invalid.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for paging.
	//
	// example:
	//
	// ASHDADS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAccountTransactionDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryAccountTransactionDetailsResponseBodyData) GetAccountTransactionsList() *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList {
	return s.AccountTransactionsList
}

func (s *QueryAccountTransactionDetailsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryAccountTransactionDetailsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryAccountTransactionDetailsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetAccountName(v string) *QueryAccountTransactionDetailsResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetAccountTransactionsList(v *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) *QueryAccountTransactionDetailsResponseBodyData {
	s.AccountTransactionsList = v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetMaxResults(v int32) *QueryAccountTransactionDetailsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetNextToken(v string) *QueryAccountTransactionDetailsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetTotalCount(v int32) *QueryAccountTransactionDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) Validate() error {
	if s.AccountTransactionsList != nil {
		if err := s.AccountTransactionsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList struct {
	AccountTransactionsList []*QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Repeated"`
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) GetAccountTransactionsList() []*QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	return s.AccountTransactionsList
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) SetAccountTransactionsList(v []*QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList {
	s.AccountTransactionsList = v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) Validate() error {
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

type QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList struct {
	// The amount of the transaction.
	//
	// example:
	//
	// 1.00
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The balance of the account.
	//
	// example:
	//
	// 0
	Balance *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// The billing cycle.
	//
	// example:
	//
	// 2022-10
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The type of transaction payment. Valid values:
	//
	// Cash: pay for the transaction in cash. Deposit: pay for the transaction with deposit. RegularBankCreditRefund: pay for the transaction with credit refund controlled by a bank. DirectPay: directly pay for the transaction.
	//
	// example:
	//
	// Cash
	FundType *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	// The ID of the order or bill.
	//
	// example:
	//
	// 2022120336190912
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Test
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The transaction account.
	//
	// example:
	//
	// fortune_test@xxx.com
	TransactionAccount *string `json:"TransactionAccount,omitempty" xml:"TransactionAccount,omitempty"`
	// The transaction channel.
	//
	// example:
	//
	// ALIPAY
	TransactionChannel *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	// The serial number of the transaction channel.
	//
	// example:
	//
	// 123232434343532
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	// Indicates whether the transaction is of the income type or the expenditure type. If one of the following types is specified, results for the specific type are returned. If the type that you specified for the parameter does not belong to the following types, no result is returned. If the parameter is left empty, results for transactions of the income and expenditure types are all returned. Valid values:
	//
	// Income and Expense.
	//
	// example:
	//
	// Income
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
	// 2022-10-01
	TransactionTime *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	// The type of the transaction. If one of the following transaction types is specified, results for the specified transaction type are returned. If the transaction type that you specified does not belong to the following transaction types, no result is returned. If the parameter is left empty, results for all transaction types are returned. Valid values:
	//
	// Payment, Withdraw, Refund, Consumption, Transfer, and Adjust.
	//
	// example:
	//
	// Consumption
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetAmount() *string {
	return s.Amount
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetBalance() *string {
	return s.Balance
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetFundType() *string {
	return s.FundType
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetRecordID() *string {
	return s.RecordID
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetRemarks() *string {
	return s.Remarks
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionAccount() *string {
	return s.TransactionAccount
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionChannel() *string {
	return s.TransactionChannel
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionChannelSN() *string {
	return s.TransactionChannelSN
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionFlow() *string {
	return s.TransactionFlow
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionNumber() *string {
	return s.TransactionNumber
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionTime() *string {
	return s.TransactionTime
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GetTransactionType() *string {
	return s.TransactionType
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetAmount(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Amount = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBalance(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Balance = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBillingCycle(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetFundType(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.FundType = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRecordID(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRemarks(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Remarks = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionAccount(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionAccount = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannel(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannelSN(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionFlow(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionFlow = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionNumber(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionTime(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionTime = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionType(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) Validate() error {
	return dara.Validate(s)
}
