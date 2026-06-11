// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicket4CopilotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateTicket4CopilotRequest
	GetAccountName() *string
	SetAccountType(v int32) *CreateTicket4CopilotRequest
	GetAccountType() *int32
	SetCopilotId(v string) *CreateTicket4CopilotRequest
	GetCopilotId() *string
	SetExpireTime(v int32) *CreateTicket4CopilotRequest
	GetExpireTime() *int32
	SetTicketNum(v int32) *CreateTicket4CopilotRequest
	GetTicketNum() *int32
	SetUserId(v string) *CreateTicket4CopilotRequest
	GetUserId() *string
}

type CreateTicket4CopilotRequest struct {
	// The name of the user account.
	//
	// 	Notice: Note: Specify either UserId or AccountName. If you leave both parameters empty, the ticket is bound to the API caller by default. Access is then granted based on the caller\\"s identity.
	//
	// example:
	//
	// 测试用户
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the user account:
	//
	// - 1: Alibaba Cloud account
	//
	// - 3: Quick BI user
	//
	// - 4: DingTalk
	//
	// - 5: RAM user
	//
	// - 6: Third-party account (an account integrated using protocols such as SAML or OAuth)
	//
	// - 9: WeCom
	//
	// - 10: Lark
	//
	// 	Notice:
	//
	// Note: This parameter is required if you specify AccountName.
	//
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the embedded Copilot module.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccd3428c-dd23-460c-a608-26bae29dffee
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// The expiration time of the ticket.
	//
	// - Unit: minutes. The maximum validity period is 240 minutes (4 hours).
	//
	// - Default: 240 minutes.
	//
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The number of times the ticket can be used. The value can range from 1 to 99,999.
	//
	// - Default: 1.
	//
	// - Recommended: 1.
	//
	// - Maximum: 99,999.
	//
	// Each access decrements the ticket\\"s usage count by one.
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	// The ID of the Quick BI user. This is not your Alibaba Cloud account ID. Call the QueryUserInfoByAccount operation to obtain the user ID. Example: `fe67f61a35a94b7da1a34ba174a7****`.
	//
	// 	Notice:
	//
	// Note: Specify either UserId or AccountName. If you leave both parameters empty, the ticket is bound to the API caller by default. Access is then granted based on the caller\\"s identity.
	//
	// example:
	//
	// 9c-asdawf-casxcasd-asdasd
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateTicket4CopilotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicket4CopilotRequest) GoString() string {
	return s.String()
}

func (s *CreateTicket4CopilotRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateTicket4CopilotRequest) GetAccountType() *int32 {
	return s.AccountType
}

func (s *CreateTicket4CopilotRequest) GetCopilotId() *string {
	return s.CopilotId
}

func (s *CreateTicket4CopilotRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *CreateTicket4CopilotRequest) GetTicketNum() *int32 {
	return s.TicketNum
}

func (s *CreateTicket4CopilotRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateTicket4CopilotRequest) SetAccountName(v string) *CreateTicket4CopilotRequest {
	s.AccountName = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetAccountType(v int32) *CreateTicket4CopilotRequest {
	s.AccountType = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetCopilotId(v string) *CreateTicket4CopilotRequest {
	s.CopilotId = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetExpireTime(v int32) *CreateTicket4CopilotRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetTicketNum(v int32) *CreateTicket4CopilotRequest {
	s.TicketNum = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetUserId(v string) *CreateTicket4CopilotRequest {
	s.UserId = &v
	return s
}

func (s *CreateTicket4CopilotRequest) Validate() error {
	return dara.Validate(s)
}
