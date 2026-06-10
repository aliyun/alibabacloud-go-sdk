// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *RenewAppSandboxRequest
	GetConversationId() *string
}

type RenewAppSandboxRequest struct {
	// Session ID
	//
	// example:
	//
	// 5b7105a2-2999-430b-ba23-ba09149d5434
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s RenewAppSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAppSandboxRequest) GoString() string {
	return s.String()
}

func (s *RenewAppSandboxRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *RenewAppSandboxRequest) SetConversationId(v string) *RenewAppSandboxRequest {
	s.ConversationId = &v
	return s
}

func (s *RenewAppSandboxRequest) Validate() error {
	return dara.Validate(s)
}
