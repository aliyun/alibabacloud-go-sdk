// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetText(v string) *CreateAIStaffConversationRequest
	GetText() *string
}

type CreateAIStaffConversationRequest struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CreateAIStaffConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateAIStaffConversationRequest) GetText() *string {
	return s.Text
}

func (s *CreateAIStaffConversationRequest) SetText(v string) *CreateAIStaffConversationRequest {
	s.Text = &v
	return s
}

func (s *CreateAIStaffConversationRequest) Validate() error {
	return dara.Validate(s)
}
