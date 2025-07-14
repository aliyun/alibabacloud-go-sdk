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
	// User\\"s account name.
	//
	// <notice	Note: Only one of userId and accountName needs to be filled in. If neither is provided, it will default to the report owner, and the report will be accessed with that user\\"s identity.</notice>
	//
	// example:
	//
	// Test user
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// User\\"s account type:
	//
	// - 1: Alibaba Cloud Primary Account
	//
	// - 3: Quick BI Self-built Account
	//
	// - 4: DingTalk
	//
	// - 5: Alibaba Cloud RAM Account
	//
	// - 6: Third-party Account (SAML, OAuth, etc.)
	//
	// - 9: WeCom
	//
	// - 10: Feishu
	//
	// <notice	Note: If accountName is not empty, then accountType must also be provided.</notice>
	//
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// ID of the Smart Q module to be embedded.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccd3*********ae29dffee
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// Expiration time.
	//
	// - Unit: minutes, maximum 240 (4 hours).
	//
	// - Default: 240.
	//
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Range of ticket quantity:
	//
	// - Default value is 1.
	//
	// - Recommended value is 1.
	//
	// - Maximum value is 99999.
	//
	// Each time a ticket is used, the ticket count decreases by 1.
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	// Quick BI\\"s UserId.
	//
	// - You can obtain this by calling [3.1.7 Get User Details Based on Third-Party Account] or other relevant APIs.
	//
	// <notice	Note: Only one of userId and accountName needs to be filled in. If neither is provided, it will default to the report owner, and the report will be accessed with that user\\"s identity.</notice>
	//
	// example:
	//
	// 9c-asd*****asd-asdasd
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
