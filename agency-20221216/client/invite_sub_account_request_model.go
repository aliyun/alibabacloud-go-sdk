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
	// The list of account information for the invitation. A maximum of five accounts can be invited at a time.
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
	// The name of the distribution customer. Valid values:
	//
	// - If the distribution customer is an enterprise, this is the company name.
	//
	// - If the distribution customer is a T2 reseller partner, this is the partner name.
	//
	// This parameter is required.
	//
	// example:
	//
	// XXX技术有限公司
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The total budget quota that the partner allocates to the distribution customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The account manager of the customer. Maximum length: 50 characters.
	//
	// example:
	//
	// 张三
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// The customer ID returned by the CreateCustomer API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The email address of the distribution customer account. The invitation email is sent to this address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345@163.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The initial order placement status. Valid values:
	//
	// - ban: Purchase is prohibited. After the customer is registered and associated, the customer cannot immediately place orders to purchase cloud resources. You must set the order control to normal before the customer can place orders.
	//
	// - normal: Normal. After the customer is registered and associated, the customer can immediately place orders to purchase cloud resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// ban
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	// The description of the distribution customer.
	//
	// example:
	//
	// 发展xx为客户的邀请
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the sub-account. Valid values:
	//
	// - 1: end user of Agency.
	//
	// - 2: end user of Reseller.
	//
	// - 5: T2 partner of Reseller.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// The shutdown policy that the partner manages for the distribution customer. Valid values:
	//
	// - 1: delayStop. Alibaba Cloud takes over the resources and the customer enjoys a delayed shutdown quota. When the delayed shutdown quota is exhausted, instance shutdown is triggered. If payment is not made within 15 days, storage resources are released.
	//
	// - 2: noStop. Manual management with no shutdown on overdue payment. The system does not manage the resource lifecycle of the non-trusteeship linked account. The partner must manually manage the instance shutdown status of the customer.
	//
	// - 3: immediatelyStop. Instances are immediately shut down on overdue payment. When the available quota of the customer account is less than 0, the account enters the overdue payment status and instance shutdown is triggered.
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
