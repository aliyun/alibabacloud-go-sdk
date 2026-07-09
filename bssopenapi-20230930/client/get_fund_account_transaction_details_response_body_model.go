// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountTransactionDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetFundAccountTransactionDetailsResponseBody
	GetCurrentPage() *int32
	SetData(v []*GetFundAccountTransactionDetailsResponseBodyData) *GetFundAccountTransactionDetailsResponseBody
	GetData() []*GetFundAccountTransactionDetailsResponseBodyData
	SetMetadata(v interface{}) *GetFundAccountTransactionDetailsResponseBody
	GetMetadata() interface{}
	SetPageSize(v int32) *GetFundAccountTransactionDetailsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetFundAccountTransactionDetailsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetFundAccountTransactionDetailsResponseBody
	GetTotalCount() *int32
}

type GetFundAccountTransactionDetailsResponseBody struct {
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Data list
	Data []*GetFundAccountTransactionDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Response structure metadata
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetFundAccountTransactionDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountTransactionDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetFundAccountTransactionDetailsResponseBody) GetData() []*GetFundAccountTransactionDetailsResponseBodyData {
	return s.Data
}

func (s *GetFundAccountTransactionDetailsResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountTransactionDetailsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFundAccountTransactionDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountTransactionDetailsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetData(v []*GetFundAccountTransactionDetailsResponseBodyData) *GetFundAccountTransactionDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetMetadata(v interface{}) *GetFundAccountTransactionDetailsResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetPageSize(v int32) *GetFundAccountTransactionDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetRequestId(v string) *GetFundAccountTransactionDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetTotalCount(v int32) *GetFundAccountTransactionDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFundAccountTransactionDetailsResponseBodyData struct {
	// Balance after the operation
	//
	// example:
	//
	// 5
	Balance *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// Order number or bill number
	//
	// example:
	//
	// 2323203243
	BillNumber *string `json:"BillNumber,omitempty" xml:"BillNumber,omitempty"`
	// External transaction serial number
	//
	// example:
	//
	// 20244389232
	ChannelTransactionNumber *string `json:"ChannelTransactionNumber,omitempty" xml:"ChannelTransactionNumber,omitempty"`
	// Transaction amount currency
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// Enterprise entity ID
	//
	// example:
	//
	// 23473943
	FundAccountEcid *string `json:"FundAccountEcid,omitempty" xml:"FundAccountEcid,omitempty"`
	// Account ID
	//
	// example:
	//
	// 1232121
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// Account name
	//
	// example:
	//
	// 云某的名称
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// Alibaba Cloud account ID of the account owner
	//
	// example:
	//
	// 32343231
	FundAccountOwnerAccountId *int64 `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// Fund type
	//
	// example:
	//
	// ACCT_BOOK
	FundType *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	// Primary marketplace
	//
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// Remarks
	//
	// example:
	//
	// 订单备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Site
	//
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// Corresponding transaction account, such as the Alipay top-up account or the counterparty account for transfers.
	//
	// example:
	//
	// 186****3975
	TransactionAccount *string `json:"TransactionAccount,omitempty" xml:"TransactionAccount,omitempty"`
	// Transaction amount
	//
	// example:
	//
	// 10
	TransactionAmount *string `json:"TransactionAmount,omitempty" xml:"TransactionAmount,omitempty"`
	// Transaction channel. If specified, the query filters by transaction channel. If not specified, all channels are queried by default.
	//
	// User balance: ACCT_CASH.
	//
	// Alipay: ALIPAY.
	//
	// Alipay Zhifutong: ALIPAY_ZHIFUTONG.
	//
	// Offline remittance: OFFLINE_REMIT.
	//
	// Credit control quota refund: REFUND.
	//
	// Online banking: UNION_PAY_BANK.
	//
	// Credit card: CREDIT_CARD. (International site only)
	//
	// PayPal: PAYPAL. (International site only)
	//
	// example:
	//
	// ALIPAY
	TransactionChannel *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	// Transaction direction: in/out (income/expenditure)
	//
	// example:
	//
	// IN
	TransactionDirection *string `json:"TransactionDirection,omitempty" xml:"TransactionDirection,omitempty"`
	// Transaction serial number
	//
	// example:
	//
	// 5423121
	TransactionNumber *int64 `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// Formatted transaction time string
	//
	// example:
	//
	// 2024-12-01 12:00:00
	TransactionTime *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	// Transaction type. If a transaction type is specified, only results of that type are returned. If the specified type does not exist, the result is empty. If not specified, all types are returned by default.
	//
	// Top-up: CHARGE.
	//
	// Withdrawal: WITHDRAW.
	//
	// Refund: REFUND.
	//
	// Payment: PAY.
	//
	// Transfer: TRANSFER.
	//
	// Adjustment: ADJUST.
	//
	// Order expiration refund: PAY_FAILED.
	//
	// example:
	//
	// CHARGE
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s GetFundAccountTransactionDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountTransactionDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetBalance() *string {
	return s.Balance
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetBillNumber() *string {
	return s.BillNumber
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetChannelTransactionNumber() *string {
	return s.ChannelTransactionNumber
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetFundAccountEcid() *string {
	return s.FundAccountEcid
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetFundAccountName() *string {
	return s.FundAccountName
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetFundAccountOwnerAccountId() *int64 {
	return s.FundAccountOwnerAccountId
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetFundType() *string {
	return s.FundType
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetNbid() *string {
	return s.Nbid
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetSite() *string {
	return s.Site
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionAccount() *string {
	return s.TransactionAccount
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionAmount() *string {
	return s.TransactionAmount
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionChannel() *string {
	return s.TransactionChannel
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionDirection() *string {
	return s.TransactionDirection
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionNumber() *int64 {
	return s.TransactionNumber
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionTime() *string {
	return s.TransactionTime
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) GetTransactionType() *string {
	return s.TransactionType
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetBalance(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Balance = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetBillNumber(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetCurrency(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountEcid(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountEcid = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountName(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountOwnerAccountId(v int64) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundType(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetNbid(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Nbid = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetRemark(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetSite(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Site = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionAccount(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionAccount = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionAmount(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionAmount = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionChannel(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionChannel = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionTime(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionType(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
