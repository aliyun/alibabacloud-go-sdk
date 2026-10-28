// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteSubResellerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountInfoList(v []*InviteSubResellerRequestAccountInfoList) *InviteSubResellerRequest
	GetAccountInfoList() []*InviteSubResellerRequestAccountInfoList
}

type InviteSubResellerRequest struct {
	// The list of account information for the invitation. A maximum of five accounts can be specified at a time.
	AccountInfoList []*InviteSubResellerRequestAccountInfoList `json:"AccountInfoList,omitempty" xml:"AccountInfoList,omitempty" type:"Repeated"`
}

func (s InviteSubResellerRequest) String() string {
	return dara.Prettify(s)
}

func (s InviteSubResellerRequest) GoString() string {
	return s.String()
}

func (s *InviteSubResellerRequest) GetAccountInfoList() []*InviteSubResellerRequestAccountInfoList {
	return s.AccountInfoList
}

func (s *InviteSubResellerRequest) SetAccountInfoList(v []*InviteSubResellerRequestAccountInfoList) *InviteSubResellerRequest {
	s.AccountInfoList = v
	return s
}

func (s *InviteSubResellerRequest) Validate() error {
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

type InviteSubResellerRequestAccountInfoList struct {
	// The name of the distribution customer.
	//
	// - If the distribution customer is an enterprise, this is the company name.
	//
	// - If the distribution customer is a T2 Reseller partner, this is the partner name.
	//
	// example:
	//
	// XXX技术有限公司
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The total budget quota that the partner allocates to the distribution customer.
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The reason for applying for a cross-region association.
	//
	// example:
	//
	// XXX
	CrossScopeRemark *string `json:"CrossScopeRemark,omitempty" xml:"CrossScopeRemark,omitempty"`
	// The customer account manager. The value can be up to 50 characters in length.
	//
	// example:
	//
	// 张三
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// The email address to which the invitation email is sent.
	//
	// example:
	//
	// nejatox206@getasail.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The initial order placement status. Valid values:
	//
	// - ban: Purchase is prohibited. After the customer is registered and associated, the customer cannot place orders to purchase cloud resources immediately. You must set the order control to Normal before the customer can place orders.
	//
	// - normal: Normal. After the customer is registered and associated, the customer can place orders to purchase cloud resources immediately.
	//
	// example:
	//
	// Normal
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	// The country or region of the invited T2 Reseller. The value must comply with the ISO 3166-1 alpha-2 code. You can call the ListCountries operation to query the contract coverage area list of the current UID.
	//
	// The system automatically determines whether the invitation is cross-region based on whether this parameter falls within the T1 contract coverage area.
	//
	// - If the invitation is cross-region, a cross-region invitation approval process is initiated. After the approval is granted, an invitation registration email is sent to the specified email address.
	//
	// - If the invitation is not cross-region, an invitation registration email is sent directly.
	//
	// example:
	//
	// SG
	RegisterNation *string `json:"RegisterNation,omitempty" xml:"RegisterNation,omitempty"`
	// The description of the distribution customer.
	//
	// example:
	//
	// 发展xx为客户的邀请
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// 5
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// The shutdown policy that the partner sets for the distribution customer.
	//
	// - 1: delayStop. The account uses the delayed shutdown quota for deferred instance shutdown. Alibaba Cloud takes over the resources. When the delayed shutdown quota is exhausted, instance shutdown is triggered. If the payment is not made within 15 days, storage resources are released.
	//
	// - 2: noStop. Manual management with no shutdown on overdue payment. The system does not manage the resource lifecycle of the non-trusteeship linked account. The partner must manually manage the instance shutdown status of the customer.
	//
	// - 3: immediatelyStop. Immediate shutdown on overdue payment. When the active quota of the customer account is less than 0, the account enters the overdue payment status and instance shutdown is triggered.
	//
	// example:
	//
	// 1
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
}

func (s InviteSubResellerRequestAccountInfoList) String() string {
	return dara.Prettify(s)
}

func (s InviteSubResellerRequestAccountInfoList) GoString() string {
	return s.String()
}

func (s *InviteSubResellerRequestAccountInfoList) GetAccountNickname() *string {
	return s.AccountNickname
}

func (s *InviteSubResellerRequestAccountInfoList) GetCreditLine() *string {
	return s.CreditLine
}

func (s *InviteSubResellerRequestAccountInfoList) GetCrossScopeRemark() *string {
	return s.CrossScopeRemark
}

func (s *InviteSubResellerRequestAccountInfoList) GetCustomerBd() *string {
	return s.CustomerBd
}

func (s *InviteSubResellerRequestAccountInfoList) GetEmailAddress() *string {
	return s.EmailAddress
}

func (s *InviteSubResellerRequestAccountInfoList) GetNewBuyStatus() *string {
	return s.NewBuyStatus
}

func (s *InviteSubResellerRequestAccountInfoList) GetRegisterNation() *string {
	return s.RegisterNation
}

func (s *InviteSubResellerRequestAccountInfoList) GetRemark() *string {
	return s.Remark
}

func (s *InviteSubResellerRequestAccountInfoList) GetSubAccountType() *string {
	return s.SubAccountType
}

func (s *InviteSubResellerRequestAccountInfoList) GetZeroCreditShutdownPolicy() *string {
	return s.ZeroCreditShutdownPolicy
}

func (s *InviteSubResellerRequestAccountInfoList) SetAccountNickname(v string) *InviteSubResellerRequestAccountInfoList {
	s.AccountNickname = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetCreditLine(v string) *InviteSubResellerRequestAccountInfoList {
	s.CreditLine = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetCrossScopeRemark(v string) *InviteSubResellerRequestAccountInfoList {
	s.CrossScopeRemark = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetCustomerBd(v string) *InviteSubResellerRequestAccountInfoList {
	s.CustomerBd = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetEmailAddress(v string) *InviteSubResellerRequestAccountInfoList {
	s.EmailAddress = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetNewBuyStatus(v string) *InviteSubResellerRequestAccountInfoList {
	s.NewBuyStatus = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetRegisterNation(v string) *InviteSubResellerRequestAccountInfoList {
	s.RegisterNation = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetRemark(v string) *InviteSubResellerRequestAccountInfoList {
	s.Remark = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetSubAccountType(v string) *InviteSubResellerRequestAccountInfoList {
	s.SubAccountType = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) SetZeroCreditShutdownPolicy(v string) *InviteSubResellerRequestAccountInfoList {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

func (s *InviteSubResellerRequestAccountInfoList) Validate() error {
	return dara.Validate(s)
}
