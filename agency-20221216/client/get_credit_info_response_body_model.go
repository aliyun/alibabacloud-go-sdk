// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreditInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCreditInfoResponseBody
	GetCode() *string
	SetData(v *GetCreditInfoResponseBodyData) *GetCreditInfoResponseBody
	GetData() *GetCreditInfoResponseBodyData
	SetMessage(v string) *GetCreditInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCreditInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCreditInfoResponseBody
	GetSuccess() *bool
}

type GetCreditInfoResponseBody struct {
	// Result Code:
	//
	// - 200 OK
	//
	// - 1109 System Error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCreditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Message Information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCreditInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCreditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCreditInfoResponseBody) GetData() *GetCreditInfoResponseBodyData {
	return s.Data
}

func (s *GetCreditInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCreditInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCreditInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCreditInfoResponseBody) SetCode(v string) *GetCreditInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetData(v *GetCreditInfoResponseBodyData) *GetCreditInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCreditInfoResponseBody) SetMessage(v string) *GetCreditInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetRequestId(v string) *GetCreditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetSuccess(v bool) *GetCreditInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetCreditInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCreditInfoResponseBodyData struct {
	// The Credit Control status, Value Range:</br>
	//
	// 1. normal - Sub Account status is running as usual.
	//
	// 2. arrearsNotShutdown - Sub Account status is running as usual, but have outstanding bill(s).
	//
	// 3. shutdown -  Sub Account status is down.
	//
	// example:
	//
	// normal
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// Percentage value, when the available credit limit is lower than this credit limit percentage, a notification E-mail will be sent to the main account.
	//
	// example:
	//
	// 20
	AlarmThreshold *string `json:"AlarmThreshold,omitempty" xml:"AlarmThreshold,omitempty"`
	// The Credit available to consume.
	//
	// example:
	//
	// 800
	AvailableCredit *string `json:"AvailableCredit,omitempty" xml:"AvailableCredit,omitempty"`
	// Obtain total unpaid amount on demo bill before simulated deduction.
	//
	// example:
	//
	// 0.000000
	ConsumedUndeductedValue *string `json:"ConsumedUndeductedValue,omitempty" xml:"ConsumedUndeductedValue,omitempty"`
	// The Credit Line of Sub Account
	//
	// example:
	//
	// 1000
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The Credit have been consumed by Sub Account, and haven\\"t be paid.
	//
	// example:
	//
	// 200
	OutstandingBalance *string `json:"OutstandingBalance,omitempty" xml:"OutstandingBalance,omitempty"`
	// The systematic controlling policy for resource management, specifically when the available Credit of Sub Account falls to 0 or less.</br>
	//
	// - 1: delayStop. The account have Shutdown-delay Privilege,  After Shutdown-delay Credit is ran out, Alibaba Cloud will take over resources and keep the instance for 15 days. In addition, the instance will be released if Sub Account failed to pay the bill within these 15 days.</br>
	//
	// - 2: noStop. Partner will manually manage Shutdown Status for Sub Account. Meanwhile, System would not manage the resource\\"s life-circle of Sub Account.</br>
	//
	// - 3: immediatelyStop. Once valid quota of Sub Account falls below 0 and be identified as defaulting account, it will trigger the instance shutdown immediately.</br>
	//
	// example:
	//
	// delayStop
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
	// Manage order operation.
	//
	// - ban：Ban the new purchase action.
	//
	// - normal：The account could raise new purchase order as usual.
	//
	// example:
	//
	// ban
	NewBuyStatus *string `json:"newBuyStatus,omitempty" xml:"newBuyStatus,omitempty"`
}

func (s GetCreditInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCreditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponseBodyData) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *GetCreditInfoResponseBodyData) GetAlarmThreshold() *string {
	return s.AlarmThreshold
}

func (s *GetCreditInfoResponseBodyData) GetAvailableCredit() *string {
	return s.AvailableCredit
}

func (s *GetCreditInfoResponseBodyData) GetConsumedUndeductedValue() *string {
	return s.ConsumedUndeductedValue
}

func (s *GetCreditInfoResponseBodyData) GetCreditLine() *string {
	return s.CreditLine
}

func (s *GetCreditInfoResponseBodyData) GetOutstandingBalance() *string {
	return s.OutstandingBalance
}

func (s *GetCreditInfoResponseBodyData) GetZeroCreditShutdownPolicy() *string {
	return s.ZeroCreditShutdownPolicy
}

func (s *GetCreditInfoResponseBodyData) GetNewBuyStatus() *string {
	return s.NewBuyStatus
}

func (s *GetCreditInfoResponseBodyData) SetAccountStatus(v string) *GetCreditInfoResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetAlarmThreshold(v string) *GetCreditInfoResponseBodyData {
	s.AlarmThreshold = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetAvailableCredit(v string) *GetCreditInfoResponseBodyData {
	s.AvailableCredit = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetConsumedUndeductedValue(v string) *GetCreditInfoResponseBodyData {
	s.ConsumedUndeductedValue = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetCreditLine(v string) *GetCreditInfoResponseBodyData {
	s.CreditLine = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetOutstandingBalance(v string) *GetCreditInfoResponseBodyData {
	s.OutstandingBalance = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetZeroCreditShutdownPolicy(v string) *GetCreditInfoResponseBodyData {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetNewBuyStatus(v string) *GetCreditInfoResponseBodyData {
	s.NewBuyStatus = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
