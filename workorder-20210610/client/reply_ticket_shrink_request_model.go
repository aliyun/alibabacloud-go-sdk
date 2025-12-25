// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ReplyTicketShrinkRequest
	GetContent() *string
	SetEncrypt(v bool) *ReplyTicketShrinkRequest
	GetEncrypt() *bool
	SetFileNameListShrink(v string) *ReplyTicketShrinkRequest
	GetFileNameListShrink() *string
	SetTicketId(v string) *ReplyTicketShrinkRequest
	GetTicketId() *string
	SetUid(v string) *ReplyTicketShrinkRequest
	GetUid() *string
}

type ReplyTicketShrinkRequest struct {
	// Content of the ticket reply
	//
	// This parameter is required.
	//
	// example:
	//
	// Why ECS backup failed?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Encryption status
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// The list of attachment names, GetAttachmentUploadUrl the ObjectKey field returned by the interface.
	FileNameListShrink *string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty"`
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// Alibaba Cloud UID
	//
	// example:
	//
	// 1289427240739141
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ReplyTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplyTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplyTicketShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *ReplyTicketShrinkRequest) GetEncrypt() *bool {
	return s.Encrypt
}

func (s *ReplyTicketShrinkRequest) GetFileNameListShrink() *string {
	return s.FileNameListShrink
}

func (s *ReplyTicketShrinkRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ReplyTicketShrinkRequest) GetUid() *string {
	return s.Uid
}

func (s *ReplyTicketShrinkRequest) SetContent(v string) *ReplyTicketShrinkRequest {
	s.Content = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetEncrypt(v bool) *ReplyTicketShrinkRequest {
	s.Encrypt = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetFileNameListShrink(v string) *ReplyTicketShrinkRequest {
	s.FileNameListShrink = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetTicketId(v string) *ReplyTicketShrinkRequest {
	s.TicketId = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetUid(v string) *ReplyTicketShrinkRequest {
	s.Uid = &v
	return s
}

func (s *ReplyTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
