// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConversationLockStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAppConversationLockStatusRequest
	GetConversationId() *string
}

type GetAppConversationLockStatusRequest struct {
	// session ID
	//
	// example:
	//
	// 5b7105a2-2999-430b-ba23-ba09149d5434
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s GetAppConversationLockStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationLockStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAppConversationLockStatusRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppConversationLockStatusRequest) SetConversationId(v string) *GetAppConversationLockStatusRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAppConversationLockStatusRequest) Validate() error {
	return dara.Validate(s)
}
