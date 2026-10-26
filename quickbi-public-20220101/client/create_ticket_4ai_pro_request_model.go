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
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	AiProId *string `json:"AiProId,omitempty" xml:"AiProId,omitempty"`
	// example:
	//
	// test
	CustomParam *string `json:"CustomParam,omitempty" xml:"CustomParam,omitempty"`
	// example:
	//
	// 22
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// test
	QuickBIUserId *string `json:"QuickBIUserId,omitempty" xml:"QuickBIUserId,omitempty"`
	// example:
	//
	// test
	SystemUid *string `json:"SystemUid,omitempty" xml:"SystemUid,omitempty"`
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
