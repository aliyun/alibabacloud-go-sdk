// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketNotesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListTicketNotesRequest
	GetLanguage() *string
	SetTicketId(v string) *ListTicketNotesRequest
	GetTicketId() *string
}

type ListTicketNotesRequest struct {
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// G2BKRWG
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ListTicketNotesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesRequest) GoString() string {
	return s.String()
}

func (s *ListTicketNotesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListTicketNotesRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketNotesRequest) SetLanguage(v string) *ListTicketNotesRequest {
	s.Language = &v
	return s
}

func (s *ListTicketNotesRequest) SetTicketId(v string) *ListTicketNotesRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketNotesRequest) Validate() error {
	return dara.Validate(s)
}
