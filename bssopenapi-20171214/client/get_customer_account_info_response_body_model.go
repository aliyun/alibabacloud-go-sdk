// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerAccountInfoResponseBody
	GetCode() *string
	SetData(v *GetCustomerAccountInfoResponseBodyData) *GetCustomerAccountInfoResponseBody
	GetData() *GetCustomerAccountInfoResponseBodyData
	SetMessage(v string) *GetCustomerAccountInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerAccountInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerAccountInfoResponseBody
	GetSuccess() *bool
}

type GetCustomerAccountInfoResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCustomerAccountInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 94017C56-1725-5DD9-AB41-B3BAE791D600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerAccountInfoResponseBody) GetData() *GetCustomerAccountInfoResponseBodyData {
	return s.Data
}

func (s *GetCustomerAccountInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerAccountInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerAccountInfoResponseBody) SetCode(v string) *GetCustomerAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetData(v *GetCustomerAccountInfoResponseBodyData) *GetCustomerAccountInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetMessage(v string) *GetCustomerAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetRequestId(v string) *GetCustomerAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetSuccess(v bool) *GetCustomerAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerAccountInfoResponseBodyData struct {
	// The type of the account. A value of 1 indicates an enterprise account. A value of 0 indicates an individual account.
	//
	// example:
	//
	// 1
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The business status of the customer. Valid values:
	//
	// Freeze: The business is frozen.
	//
	// Thaw: The business is unfrozen.
	//
	// Trusteeship: The business is hosted.
	//
	// TrusteeshipCancel: The business is not hosted.
	//
	// example:
	//
	// Freeze
	CreditLimitStatus *string `json:"CreditLimitStatus,omitempty" xml:"CreditLimitStatus,omitempty"`
	// The hosting status of the credit information and instances of the customer. If the credit information and instances of the customer are managed on Alibaba Cloud, Alibaba Cloud suspends a customer service upon overdue payment. Valid values:
	//
	// FREEZE: The business of the customer is frozen.
	//
	// TRUSTEESHIP: The business of the customer is hosted.
	//
	// example:
	//
	// FREEZE
	HostingStatus *string `json:"HostingStatus,omitempty" xml:"HostingStatus,omitempty"`
	// Indicates whether the account passes the real-name verification.
	//
	// example:
	//
	// true
	IsCertified *bool `json:"IsCertified,omitempty" xml:"IsCertified,omitempty"`
	// The email address of the customer.
	//
	// example:
	//
	// xxxx@aliyun.com
	LoginEmail *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// The ID of the management account.
	//
	// example:
	//
	// 1051360339779133
	Mpk *int64 `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
}

func (s GetCustomerAccountInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerAccountInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoResponseBodyData) GetAccountType() *string {
	return s.AccountType
}

func (s *GetCustomerAccountInfoResponseBodyData) GetCreditLimitStatus() *string {
	return s.CreditLimitStatus
}

func (s *GetCustomerAccountInfoResponseBodyData) GetHostingStatus() *string {
	return s.HostingStatus
}

func (s *GetCustomerAccountInfoResponseBodyData) GetIsCertified() *bool {
	return s.IsCertified
}

func (s *GetCustomerAccountInfoResponseBodyData) GetLoginEmail() *string {
	return s.LoginEmail
}

func (s *GetCustomerAccountInfoResponseBodyData) GetMpk() *int64 {
	return s.Mpk
}

func (s *GetCustomerAccountInfoResponseBodyData) SetAccountType(v string) *GetCustomerAccountInfoResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetCreditLimitStatus(v string) *GetCustomerAccountInfoResponseBodyData {
	s.CreditLimitStatus = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetHostingStatus(v string) *GetCustomerAccountInfoResponseBodyData {
	s.HostingStatus = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetIsCertified(v bool) *GetCustomerAccountInfoResponseBodyData {
	s.IsCertified = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetLoginEmail(v string) *GetCustomerAccountInfoResponseBodyData {
	s.LoginEmail = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetMpk(v int64) *GetCustomerAccountInfoResponseBodyData {
	s.Mpk = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
