// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *CloseTicketRequest
	GetLanguage() *string
	SetTicketId(v string) *CloseTicketRequest
	GetTicketId() *string
}

type CloseTicketRequest struct {
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3EBYRY7
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s CloseTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseTicketRequest) GoString() string {
	return s.String()
}

func (s *CloseTicketRequest) GetLanguage() *string {
	return s.Language
}

func (s *CloseTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *CloseTicketRequest) SetLanguage(v string) *CloseTicketRequest {
	s.Language = &v
	return s
}

func (s *CloseTicketRequest) SetTicketId(v string) *CloseTicketRequest {
	s.TicketId = &v
	return s
}

func (s *CloseTicketRequest) Validate() error {
	return dara.Validate(s)
}
