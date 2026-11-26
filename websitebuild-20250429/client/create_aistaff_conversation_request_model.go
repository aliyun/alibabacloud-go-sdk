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
	// The text of the user question. The first 100 characters are truncated and used as the session title.
	//
	// example:
	//
	// Bathroom hardware products, featuring a brushed stainless steel surface with a modern silver-gray color tone, presenting a minimalist and textured style. The products include showerheads, faucets, and towel racks, all of which are rust-proof and corrosion-resistant, making them suitable for humid environments. The connections between components are precisely designed with hidden screw structures to enhance overall aesthetics. The faucets and showerheads support multiple water flow modes, and the handles operate smoothly, conforming to ergonomic design. All hardware components have passed the national water-saving certification, support cold and hot water connections, and are wall-mounted
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
