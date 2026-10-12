// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateChatSessionRequest
	GetName() *string
	SetProducts(v []*string) *CreateChatSessionRequest
	GetProducts() []*string
	SetTalkType(v string) *CreateChatSessionRequest
	GetTalkType() *string
}

type CreateChatSessionRequest struct {
	// The name of the chat session.
	//
	// example:
	//
	// Product consultation session
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of API products associated with the chat session.
	Products []*string `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// The conversation type of the session. Valid values: MODEL and AGENT.
	//
	// example:
	//
	// MODEL
	TalkType *string `json:"talkType,omitempty" xml:"talkType,omitempty"`
}

func (s CreateChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateChatSessionRequest) GetName() *string {
	return s.Name
}

func (s *CreateChatSessionRequest) GetProducts() []*string {
	return s.Products
}

func (s *CreateChatSessionRequest) GetTalkType() *string {
	return s.TalkType
}

func (s *CreateChatSessionRequest) SetName(v string) *CreateChatSessionRequest {
	s.Name = &v
	return s
}

func (s *CreateChatSessionRequest) SetProducts(v []*string) *CreateChatSessionRequest {
	s.Products = v
	return s
}

func (s *CreateChatSessionRequest) SetTalkType(v string) *CreateChatSessionRequest {
	s.TalkType = &v
	return s
}

func (s *CreateChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
