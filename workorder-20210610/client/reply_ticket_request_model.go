// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ReplyTicketRequest
	GetContent() *string
	SetEncrypt(v bool) *ReplyTicketRequest
	GetEncrypt() *bool
	SetFileNameList(v []*string) *ReplyTicketRequest
	GetFileNameList() []*string
	SetTicketId(v string) *ReplyTicketRequest
	GetTicketId() *string
	SetUid(v string) *ReplyTicketRequest
	GetUid() *string
}

type ReplyTicketRequest struct {
	// The content of the reply to the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// OK, the renewal was successful on my end.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to encrypt the content.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// A list of attachment names. This is the ObjectKey field returned by the GetAttachmentUploadUrl operation.
	FileNameList []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
	// The ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// The Alibaba Cloud account ID (UID).
	//
	// example:
	//
	// 1289427240739141
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ReplyTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplyTicketRequest) GoString() string {
	return s.String()
}

func (s *ReplyTicketRequest) GetContent() *string {
	return s.Content
}

func (s *ReplyTicketRequest) GetEncrypt() *bool {
	return s.Encrypt
}

func (s *ReplyTicketRequest) GetFileNameList() []*string {
	return s.FileNameList
}

func (s *ReplyTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ReplyTicketRequest) GetUid() *string {
	return s.Uid
}

func (s *ReplyTicketRequest) SetContent(v string) *ReplyTicketRequest {
	s.Content = &v
	return s
}

func (s *ReplyTicketRequest) SetEncrypt(v bool) *ReplyTicketRequest {
	s.Encrypt = &v
	return s
}

func (s *ReplyTicketRequest) SetFileNameList(v []*string) *ReplyTicketRequest {
	s.FileNameList = v
	return s
}

func (s *ReplyTicketRequest) SetTicketId(v string) *ReplyTicketRequest {
	s.TicketId = &v
	return s
}

func (s *ReplyTicketRequest) SetUid(v string) *ReplyTicketRequest {
	s.Uid = &v
	return s
}

func (s *ReplyTicketRequest) Validate() error {
	return dara.Validate(s)
}
