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
	// The result code. Valid values:
	//
	// - 200: OK.
	//
	// - 1109: System error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetCreditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message information.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Alibaba Cloud uses this ID for error tracking.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current API call is successful. Valid values:
	//
	// - True: Successful.
	//
	// - False: Failed.
	//
	// This does not indicate whether the subsequent business operation is successful.
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
	// The credit control status of the sub-account. Valid values:
	//
	// 1. normal: Normal.
	//
	// 2. arrearsNotShutdown: The account has an overdue payment but is not shut down.
	//
	// 3. shutdown: Shut down.
	//
	// example:
	//
	// normal
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The percentage value. When the available credit drops below this percentage of the credit limit, an email notification is sent to the main account.
	//
	// example:
	//
	// 20
	AlarmThreshold *string `json:"AlarmThreshold,omitempty" xml:"AlarmThreshold,omitempty"`
	// The available credit.
	//
	// example:
	//
	// 800
	AvailableCredit *string `json:"AvailableCredit,omitempty" xml:"AvailableCredit,omitempty"`
	// The total outstanding amount of virtual bills before simulated deduction.
	//
	// example:
	//
	// 0.000000
	ConsumedUndeductedValue *string `json:"ConsumedUndeductedValue,omitempty" xml:"ConsumedUndeductedValue,omitempty"`
	// The credit limit of the sub-account.
	//
	// example:
	//
	// 1000
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The credit amount consumed by the sub-account that has not yet been paid.
	//
	// example:
	//
	// 200
	OutstandingBalance *string `json:"OutstandingBalance,omitempty" xml:"OutstandingBalance,omitempty"`
	// The freeze status of pay-as-you-go. Valid values:
	//
	// - freeze: Frozen.
	//
	// - normal: Normal.
	//
	// example:
	//
	// normal
	PAYGFreezeStatus *string `json:"PAYGFreezeStatus,omitempty" xml:"PAYGFreezeStatus,omitempty"`
	// The policy that the system uses to manage resources when the active credit is 0 or less. Valid values:
	//
	// 1: delayStop. Alibaba Cloud takes over the resources. Instances are shut down after the credit is exhausted with a delay epoch. If payment is not made within 15 days, storage resources are released.
	//
	// 2: noStop. The system does not manage the resource lifecycle of the non-trusteeship linked account. The partner must manage it manually.
	//
	// 3: immediatelyStop. Instances are shut down immediately when the account enters the overdue payment state, which triggers instance shutdown.
	//
	// example:
	//
	// delayStop
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
	// The status that manages the purchase actions of the customer. Valid values:
	//
	// - ban: New purchases are prohibited.
	//
	// - cancelBan: Normal status.
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

func (s *GetCreditInfoResponseBodyData) GetPAYGFreezeStatus() *string {
	return s.PAYGFreezeStatus
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

func (s *GetCreditInfoResponseBodyData) SetPAYGFreezeStatus(v string) *GetCreditInfoResponseBodyData {
	s.PAYGFreezeStatus = &v
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
