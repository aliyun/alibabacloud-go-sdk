// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicket4AiProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateTicket4AiProRequest
	GetAccountName() *string
	SetAccountType(v int32) *CreateTicket4AiProRequest
	GetAccountType() *int32
	SetAiProId(v string) *CreateTicket4AiProRequest
	GetAiProId() *string
	SetCustomParam(v string) *CreateTicket4AiProRequest
	GetCustomParam() *string
	SetExpireTime(v int32) *CreateTicket4AiProRequest
	GetExpireTime() *int32
	SetQuickBIUserId(v string) *CreateTicket4AiProRequest
	GetQuickBIUserId() *string
	SetSystemUid(v string) *CreateTicket4AiProRequest
	GetSystemUid() *string
	SetTicketNum(v int32) *CreateTicket4AiProRequest
	GetTicketNum() *int32
}

type CreateTicket4AiProRequest struct {
	// The account of the access user to bind. Use this parameter together with AccountType when QuickBIUserId is not specified.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account type of the user. Valid values:
	//
	// - 1: Alibaba Cloud account.
	//
	// - 3: Quick BI custom account.
	//
	// - 4: DingTalk.
	//
	// - 5: RAM user.
	//
	// - 9: WeCom.
	//
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AIPro embedding configuration. The configuration must exist, belong to the current organization, and have a status of 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AiProId *string `json:"AiProId,omitempty" xml:"AiProId,omitempty"`
	// The custom pass-through parameter. If specified, it must be a valid JSON string.
	//
	// example:
	//
	// test
	CustomParam *string `json:"CustomParam,omitempty" xml:"CustomParam,omitempty"`
	// The validity period of the ticket, in minutes. Default value: 240. The value must be greater than 0.
	//
	// example:
	//
	// 22
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The QuickBI user ID of the access user to bind. Specify either this parameter or the combination of AccountName and AccountType.
	//
	// example:
	//
	// test
	QuickBIUserId *string `json:"QuickBIUserId,omitempty" xml:"QuickBIUserId,omitempty"`
	// The custom system identifier. The length must be less than 150 characters.
	//
	// example:
	//
	// test
	SystemUid *string `json:"SystemUid,omitempty" xml:"SystemUid,omitempty"`
	// The number of times the ticket can be used. Default value: 1. Valid values: 1 to 99999.
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
}

func (s CreateTicket4AiProRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicket4AiProRequest) GoString() string {
	return s.String()
}

func (s *CreateTicket4AiProRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateTicket4AiProRequest) GetAccountType() *int32 {
	return s.AccountType
}

func (s *CreateTicket4AiProRequest) GetAiProId() *string {
	return s.AiProId
}

func (s *CreateTicket4AiProRequest) GetCustomParam() *string {
	return s.CustomParam
}

func (s *CreateTicket4AiProRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *CreateTicket4AiProRequest) GetQuickBIUserId() *string {
	return s.QuickBIUserId
}

func (s *CreateTicket4AiProRequest) GetSystemUid() *string {
	return s.SystemUid
}

func (s *CreateTicket4AiProRequest) GetTicketNum() *int32 {
	return s.TicketNum
}

func (s *CreateTicket4AiProRequest) SetAccountName(v string) *CreateTicket4AiProRequest {
	s.AccountName = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetAccountType(v int32) *CreateTicket4AiProRequest {
	s.AccountType = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetAiProId(v string) *CreateTicket4AiProRequest {
	s.AiProId = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetCustomParam(v string) *CreateTicket4AiProRequest {
	s.CustomParam = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetExpireTime(v int32) *CreateTicket4AiProRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetQuickBIUserId(v string) *CreateTicket4AiProRequest {
	s.QuickBIUserId = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetSystemUid(v string) *CreateTicket4AiProRequest {
	s.SystemUid = &v
	return s
}

func (s *CreateTicket4AiProRequest) SetTicketNum(v int32) *CreateTicket4AiProRequest {
	s.TicketNum = &v
	return s
}

func (s *CreateTicket4AiProRequest) Validate() error {
	return dara.Validate(s)
}
