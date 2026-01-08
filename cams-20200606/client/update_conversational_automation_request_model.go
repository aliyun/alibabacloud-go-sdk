// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConversationalAutomationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommands(v []*UpdateConversationalAutomationRequestCommands) *UpdateConversationalAutomationRequest
	GetCommands() []*UpdateConversationalAutomationRequestCommands
	SetCustSpaceId(v string) *UpdateConversationalAutomationRequest
	GetCustSpaceId() *string
	SetEnableWelcomeMessage(v bool) *UpdateConversationalAutomationRequest
	GetEnableWelcomeMessage() *bool
	SetOwnerId(v int64) *UpdateConversationalAutomationRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *UpdateConversationalAutomationRequest
	GetPhoneNumber() *string
	SetPrompts(v []*string) *UpdateConversationalAutomationRequest
	GetPrompts() []*string
	SetResourceOwnerAccount(v string) *UpdateConversationalAutomationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateConversationalAutomationRequest
	GetResourceOwnerId() *int64
}

type UpdateConversationalAutomationRequest struct {
	// The commands.
	Commands []*UpdateConversationalAutomationRequestCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The space ID of the RAM user within the independent software vendor (ISV) account or the instance ID of the customer of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2993****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Specifies whether to enable the welcoming message.
	//
	// example:
	//
	// true
	EnableWelcomeMessage *bool  `json:"EnableWelcomeMessage,omitempty" xml:"EnableWelcomeMessage,omitempty"`
	OwnerId              *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number of the enterprise.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86130000***
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The opening remarks.
	Prompts              []*string `json:"Prompts,omitempty" xml:"Prompts,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateConversationalAutomationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConversationalAutomationRequest) GoString() string {
	return s.String()
}

func (s *UpdateConversationalAutomationRequest) GetCommands() []*UpdateConversationalAutomationRequestCommands {
	return s.Commands
}

func (s *UpdateConversationalAutomationRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateConversationalAutomationRequest) GetEnableWelcomeMessage() *bool {
	return s.EnableWelcomeMessage
}

func (s *UpdateConversationalAutomationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateConversationalAutomationRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateConversationalAutomationRequest) GetPrompts() []*string {
	return s.Prompts
}

func (s *UpdateConversationalAutomationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateConversationalAutomationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateConversationalAutomationRequest) SetCommands(v []*UpdateConversationalAutomationRequestCommands) *UpdateConversationalAutomationRequest {
	s.Commands = v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetCustSpaceId(v string) *UpdateConversationalAutomationRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetEnableWelcomeMessage(v bool) *UpdateConversationalAutomationRequest {
	s.EnableWelcomeMessage = &v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetOwnerId(v int64) *UpdateConversationalAutomationRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetPhoneNumber(v string) *UpdateConversationalAutomationRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetPrompts(v []*string) *UpdateConversationalAutomationRequest {
	s.Prompts = v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetResourceOwnerAccount(v string) *UpdateConversationalAutomationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateConversationalAutomationRequest) SetResourceOwnerId(v int64) *UpdateConversationalAutomationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateConversationalAutomationRequest) Validate() error {
	if s.Commands != nil {
		for _, item := range s.Commands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateConversationalAutomationRequestCommands struct {
	// The description of the command.
	//
	// example:
	//
	// Command 1.
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The command name.
	//
	// example:
	//
	// test
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
}

func (s UpdateConversationalAutomationRequestCommands) String() string {
	return dara.Prettify(s)
}

func (s UpdateConversationalAutomationRequestCommands) GoString() string {
	return s.String()
}

func (s *UpdateConversationalAutomationRequestCommands) GetCommandDescription() *string {
	return s.CommandDescription
}

func (s *UpdateConversationalAutomationRequestCommands) GetCommandName() *string {
	return s.CommandName
}

func (s *UpdateConversationalAutomationRequestCommands) SetCommandDescription(v string) *UpdateConversationalAutomationRequestCommands {
	s.CommandDescription = &v
	return s
}

func (s *UpdateConversationalAutomationRequestCommands) SetCommandName(v string) *UpdateConversationalAutomationRequestCommands {
	s.CommandName = &v
	return s
}

func (s *UpdateConversationalAutomationRequestCommands) Validate() error {
	return dara.Validate(s)
}
