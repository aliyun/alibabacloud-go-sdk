// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConversationalAutomationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandsShrink(v string) *UpdateConversationalAutomationShrinkRequest
	GetCommandsShrink() *string
	SetCustSpaceId(v string) *UpdateConversationalAutomationShrinkRequest
	GetCustSpaceId() *string
	SetEnableWelcomeMessage(v bool) *UpdateConversationalAutomationShrinkRequest
	GetEnableWelcomeMessage() *bool
	SetOwnerId(v int64) *UpdateConversationalAutomationShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *UpdateConversationalAutomationShrinkRequest
	GetPhoneNumber() *string
	SetPromptsShrink(v string) *UpdateConversationalAutomationShrinkRequest
	GetPromptsShrink() *string
	SetResourceOwnerAccount(v string) *UpdateConversationalAutomationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateConversationalAutomationShrinkRequest
	GetResourceOwnerId() *int64
}

type UpdateConversationalAutomationShrinkRequest struct {
	// The commands.
	CommandsShrink *string `json:"Commands,omitempty" xml:"Commands,omitempty"`
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
	PromptsShrink        *string `json:"Prompts,omitempty" xml:"Prompts,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateConversationalAutomationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConversationalAutomationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConversationalAutomationShrinkRequest) GetCommandsShrink() *string {
	return s.CommandsShrink
}

func (s *UpdateConversationalAutomationShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateConversationalAutomationShrinkRequest) GetEnableWelcomeMessage() *bool {
	return s.EnableWelcomeMessage
}

func (s *UpdateConversationalAutomationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateConversationalAutomationShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateConversationalAutomationShrinkRequest) GetPromptsShrink() *string {
	return s.PromptsShrink
}

func (s *UpdateConversationalAutomationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateConversationalAutomationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateConversationalAutomationShrinkRequest) SetCommandsShrink(v string) *UpdateConversationalAutomationShrinkRequest {
	s.CommandsShrink = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetCustSpaceId(v string) *UpdateConversationalAutomationShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetEnableWelcomeMessage(v bool) *UpdateConversationalAutomationShrinkRequest {
	s.EnableWelcomeMessage = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetOwnerId(v int64) *UpdateConversationalAutomationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetPhoneNumber(v string) *UpdateConversationalAutomationShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetPromptsShrink(v string) *UpdateConversationalAutomationShrinkRequest {
	s.PromptsShrink = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetResourceOwnerAccount(v string) *UpdateConversationalAutomationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) SetResourceOwnerId(v int64) *UpdateConversationalAutomationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateConversationalAutomationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
