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
	// List of invited account information, up to 5 at a time.
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
	// Name of the distribution customer:
	//
	// - If the distribution customer is a company, this is the company name.
	//
	// - If the distribution customer is a T2 reseller partner, this is the partner name.
	//
	// example:
	//
	// XXX技术有限公司
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// Total budget quota allocated by the partner to the distribution customer (quota).
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// Reason for applying for cross-regional association.
	//
	// example:
	//
	// XXX
	CrossScopeRemark *string `json:"CrossScopeRemark,omitempty" xml:"CrossScopeRemark,omitempty"`
	// Customer Business Manager (limited to 50 characters).
	//
	// example:
	//
	// 张三
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// The email address to which the invitation email will be sent.
	//
	// example:
	//
	// nejatox206@getasail.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// Initial order status:
	//
	// - ban: Purchase Ban - The customer cannot place orders immediately after successful registration and association. The \\"Order Control\\" must be set to \\"Normal\\" before placing orders.
	//
	// - normal: Normal - The customer can place orders immediately after successful registration and association.
	//
	// example:
	//
	// Normal
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	// Country or region of the invited T2 Reseller. The parameter should comply with the ISO 3166-1 two-letter code (Alpha-2). You can use the ListCountries API to get the current UID contract coverage area list.
	//
	// The system will automatically determine if the invitation is cross-regional based on whether the `registerNation` parameter is within the T1 contract coverage area:
	//
	// - If it\\"s a cross-regional invitation, a cross-regional approval process will be initiated. After approval by Alibaba Cloud, an invitation registration email will be sent to the invited email address.
	//
	// - If it\\"s not a cross-regional invitation, an invitation registration email will be sent directly.
	//
	// example:
	//
	// SG
	RegisterNation *string `json:"RegisterNation,omitempty" xml:"RegisterNation,omitempty"`
	// Description of the distribution customer.
	//
	// example:
	//
	// 发展xx为客户的邀请
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Do not fill in, deprecated parameter.
	//
	// example:
	//
	// 5
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// Management of the shutdown policy for the distribution customer by the partner.
	//
	// - 1: delayStop - Enjoy delayed shutdown with a grace period: Alibaba Cloud takes over the resources, and when the grace period quota is exhausted, the instance will be shut down. If payment is not made within 15 days, the storage resources will be released.
	//
	// - 2: noStop - Manual management of overdue without stopping: The system does not manage the lifecycle of the sub-account\\"s resources. The partner needs to manually manage the shutdown status of the customer\\"s instances.
	//
	// - 3: immediatelyStop - Immediate shutdown upon overdue: When the available quota of the customer\\"s account is less than 0, the account enters an overdue state, triggering the shutdown of the instance.
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
