// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFinancialAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFinancialAccountInfoResponseBody
	GetCode() *string
	SetData(v *QueryFinancialAccountInfoResponseBodyData) *QueryFinancialAccountInfoResponseBody
	GetData() *QueryFinancialAccountInfoResponseBodyData
	SetMessage(v string) *QueryFinancialAccountInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryFinancialAccountInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFinancialAccountInfoResponseBody
	GetSuccess() *bool
}

type QueryFinancialAccountInfoResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryFinancialAccountInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 97426BC2-792E-41F5-92B8-73C710DA7779
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFinancialAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFinancialAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFinancialAccountInfoResponseBody) GetData() *QueryFinancialAccountInfoResponseBodyData {
	return s.Data
}

func (s *QueryFinancialAccountInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFinancialAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFinancialAccountInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFinancialAccountInfoResponseBody) SetCode(v string) *QueryFinancialAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetData(v *QueryFinancialAccountInfoResponseBodyData) *QueryFinancialAccountInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetMessage(v string) *QueryFinancialAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetRequestId(v string) *QueryFinancialAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetSuccess(v bool) *QueryFinancialAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryFinancialAccountInfoResponseBodyData struct {
	// The type of the financial account. A value of MASTER indicates that the account is a management account. A value of MEMBER indicates that the account is a member.
	//
	// example:
	//
	// MEMBER
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Indicates whether the account is a financial account. A value of true indicates that the account is a financial account. A value of false indicates that the account is not a financial account.
	//
	// example:
	//
	// true
	IsFinancialAccount *bool `json:"IsFinancialAccount,omitempty" xml:"IsFinancialAccount,omitempty"`
	// The ID of the group to which the member belongs.
	//
	// example:
	//
	// 124
	MemberGroupId *int64 `json:"MemberGroupId,omitempty" xml:"MemberGroupId,omitempty"`
	// The name of the group to which the member belongs.
	//
	// example:
	//
	// Test group
	MemberGroupName *string `json:"MemberGroupName,omitempty" xml:"MemberGroupName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// Test
	MemberNickName *string `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// killortop3
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryFinancialAccountInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFinancialAccountInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoResponseBodyData) GetAccountType() *string {
	return s.AccountType
}

func (s *QueryFinancialAccountInfoResponseBodyData) GetIsFinancialAccount() *bool {
	return s.IsFinancialAccount
}

func (s *QueryFinancialAccountInfoResponseBodyData) GetMemberGroupId() *int64 {
	return s.MemberGroupId
}

func (s *QueryFinancialAccountInfoResponseBodyData) GetMemberGroupName() *string {
	return s.MemberGroupName
}

func (s *QueryFinancialAccountInfoResponseBodyData) GetMemberNickName() *string {
	return s.MemberNickName
}

func (s *QueryFinancialAccountInfoResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetAccountType(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetIsFinancialAccount(v bool) *QueryFinancialAccountInfoResponseBodyData {
	s.IsFinancialAccount = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetMemberGroupId(v int64) *QueryFinancialAccountInfoResponseBodyData {
	s.MemberGroupId = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetMemberGroupName(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.MemberGroupName = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetMemberNickName(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.MemberNickName = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetUserName(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.UserName = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
