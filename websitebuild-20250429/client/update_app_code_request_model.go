// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateAppCodeRequest
	GetContent() *string
	SetConversationId(v string) *UpdateAppCodeRequest
	GetConversationId() *string
}

type UpdateAppCodeRequest struct {
	// example:
	//
	// verify_6554d8cc0de584306d16506dd119cbfc
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s UpdateAppCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppCodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppCodeRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateAppCodeRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateAppCodeRequest) SetContent(v string) *UpdateAppCodeRequest {
	s.Content = &v
	return s
}

func (s *UpdateAppCodeRequest) SetConversationId(v string) *UpdateAppCodeRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateAppCodeRequest) Validate() error {
	return dara.Validate(s)
}
