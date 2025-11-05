// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteSubAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountInfoList(v []*InviteSubAccountRequestAccountInfoList) *InviteSubAccountRequest
	GetAccountInfoList() []*InviteSubAccountRequestAccountInfoList
}

type InviteSubAccountRequest struct {
	// List of invited account information,  less than 5 accounts at a time.</br>
	//
	// `Sub-levels <= 5`
	//
	// This parameter is required.
	AccountInfoList []*InviteSubAccountRequestAccountInfoList `json:"AccountInfoList,omitempty" xml:"AccountInfoList,omitempty" type:"Repeated"`
}

func (s InviteSubAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountRequest) GoString() string {
	return s.String()
}

func (s *InviteSubAccountRequest) GetAccountInfoList() []*InviteSubAccountRequestAccountInfoList {
	return s.AccountInfoList
}

func (s *InviteSubAccountRequest) SetAccountInfoList(v []*InviteSubAccountRequestAccountInfoList) *InviteSubAccountRequest {
	s.AccountInfoList = v
	return s
}

func (s *InviteSubAccountRequest) Validate() error {
	if s.AccountInfoList != nil {
		for _, item := range s.AccountInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InviteSubAccountRequestAccountInfoList struct {
	// The name of Sub Account:</br>
	//
	// 1. Use the official name of Company, if Sub Account is an enterprise.</br>
	//
	// 2. Use the official name of Partner, if Sub Account is a T2 reseller.</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// XXX Technology LTD.
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The total budget Credit of Sub Account that distributed by Partner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// Customer ID, Returning ID from CreateCustomer API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The email address of End User,  which will receive the invitation email.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345@163.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// Initial Order Status</br>
	//
	// 1. ban：Ban the new purchase action--After End User finish registration and authorization, they can\\"t issue Cloud Resource order immediately. Partner should manually update the "Order Control" settings as "Normal" to enable new order.</br>
	//
	// 2. normal：Normal--After End User finished registration and authorization, they can issue Cloud Resource order immediately.</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// ban
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	// Description of Sub Account.
	//
	// example:
	//
	// The invitation to develop XX as a Sub Account
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of Sub Account</br>
	//
	// 1 Agency\\"s End User</br>
	//
	// 2 Reseller\\"s End user</br>
	//
	// 5 Reseller\\"s T2 Partner</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// Partner\\"s Shutdown Policy Management for Sub Account.</br>
	//
	// 1: delayStop. The account have Shutdown-delay Privilege,  After Shutdown-delay Credit is ran out, Alibaba Cloud will take over resources and keep the instance for 15 days. In addition, the instance will be released if Sub Account failed to pay the bill within these 15 days.</br>
	//
	// 2: noStop. Partner will manually manage Shutdown Status for Sub Account. Meanwhile, System would not manage the resource\\"s life-circle of Sub Account.</br>
	//
	// 3: immediatelyStop. Once valid quota of Sub Account falls below 0 and be identified as defaulting account, it will trigger the instance shutdown immediately.</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
}

func (s InviteSubAccountRequestAccountInfoList) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountRequestAccountInfoList) GoString() string {
	return s.String()
}

func (s *InviteSubAccountRequestAccountInfoList) GetAccountNickname() *string {
	return s.AccountNickname
}

func (s *InviteSubAccountRequestAccountInfoList) GetCreditLine() *string {
	return s.CreditLine
}

func (s *InviteSubAccountRequestAccountInfoList) GetCustomerBd() *string {
	return s.CustomerBd
}

func (s *InviteSubAccountRequestAccountInfoList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *InviteSubAccountRequestAccountInfoList) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *InviteSubAccountRequestAccountInfoList) GetNewBuyStatus() *string {
	return s.NewBuyStatus
}

func (s *InviteSubAccountRequestAccountInfoList) GetRemark() *string {
	return s.Remark
}

func (s *InviteSubAccountRequestAccountInfoList) GetSubAccountType() *string {
	return s.SubAccountType
}

func (s *InviteSubAccountRequestAccountInfoList) GetZeroCreditShutdownPolicy() *string {
	return s.ZeroCreditShutdownPolicy
}

func (s *InviteSubAccountRequestAccountInfoList) SetAccountNickname(v string) *InviteSubAccountRequestAccountInfoList {
	s.AccountNickname = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCreditLine(v string) *InviteSubAccountRequestAccountInfoList {
	s.CreditLine = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCustomerBd(v string) *InviteSubAccountRequestAccountInfoList {
	s.CustomerBd = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCustomerId(v string) *InviteSubAccountRequestAccountInfoList {
	s.CustomerId = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetEmailAddress(v string) *InviteSubAccountRequestAccountInfoList {
	s.EmailAddress = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetNewBuyStatus(v string) *InviteSubAccountRequestAccountInfoList {
	s.NewBuyStatus = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetRemark(v string) *InviteSubAccountRequestAccountInfoList {
	s.Remark = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetSubAccountType(v string) *InviteSubAccountRequestAccountInfoList {
	s.SubAccountType = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetZeroCreditShutdownPolicy(v string) *InviteSubAccountRequestAccountInfoList {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) Validate() error {
	return dara.Validate(s)
}
