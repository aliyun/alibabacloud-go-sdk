// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicketId(v string) *GetTicketRequest
	GetTicketId() *string
	SetUid(v string) *GetTicketRequest
	GetUid() *string
}

type GetTicketRequest struct {
	// Work Order Number
	//
	// This parameter is required.
	//
	// example:
	//
	// 001571BTXC
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// UID
	//
	// example:
	//
	// 1604499804113750
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *GetTicketRequest) GetUid() *string {
	return s.Uid
}

func (s *GetTicketRequest) SetTicketId(v string) *GetTicketRequest {
	s.TicketId = &v
	return s
}

func (s *GetTicketRequest) SetUid(v string) *GetTicketRequest {
	s.Uid = &v
	return s
}

func (s *GetTicketRequest) Validate() error {
	return dara.Validate(s)
}
