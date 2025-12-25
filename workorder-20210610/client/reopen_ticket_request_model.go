// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ReopenTicketRequest
	GetContent() *string
	SetTicketId(v string) *ReopenTicketRequest
	GetTicketId() *string
	SetUid(v string) *ReopenTicketRequest
	GetUid() *string
}

type ReopenTicketRequest struct {
	// Reply content of rework order
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs backup fails
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Work Order Number
	//
	// This parameter is required.
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// UID
	//
	// example:
	//
	// 1013872004421947
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ReopenTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s ReopenTicketRequest) GoString() string {
	return s.String()
}

func (s *ReopenTicketRequest) GetContent() *string {
	return s.Content
}

func (s *ReopenTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ReopenTicketRequest) GetUid() *string {
	return s.Uid
}

func (s *ReopenTicketRequest) SetContent(v string) *ReopenTicketRequest {
	s.Content = &v
	return s
}

func (s *ReopenTicketRequest) SetTicketId(v string) *ReopenTicketRequest {
	s.TicketId = &v
	return s
}

func (s *ReopenTicketRequest) SetUid(v string) *ReopenTicketRequest {
	s.Uid = &v
	return s
}

func (s *ReopenTicketRequest) Validate() error {
	return dara.Validate(s)
}
