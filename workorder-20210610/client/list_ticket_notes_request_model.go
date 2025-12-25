// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketNotesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicketId(v string) *ListTicketNotesRequest
	GetTicketId() *string
	SetUid(v string) *ListTicketNotesRequest
	GetUid() *string
}

type ListTicketNotesRequest struct {
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
	// 1936753548534516
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListTicketNotesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesRequest) GoString() string {
	return s.String()
}

func (s *ListTicketNotesRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketNotesRequest) GetUid() *string {
	return s.Uid
}

func (s *ListTicketNotesRequest) SetTicketId(v string) *ListTicketNotesRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketNotesRequest) SetUid(v string) *ListTicketNotesRequest {
	s.Uid = &v
	return s
}

func (s *ListTicketNotesRequest) Validate() error {
	return dara.Validate(s)
}
