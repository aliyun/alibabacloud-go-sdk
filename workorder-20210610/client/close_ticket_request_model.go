// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicketId(v string) *CloseTicketRequest
	GetTicketId() *string
	SetUid(v string) *CloseTicketRequest
	GetUid() *string
}

type CloseTicketRequest struct {
	// Work Order Number
	//
	// This parameter is required.
	//
	// example:
	//
	// G2BKRWG
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// The UID of the Alibaba Cloud account. You can view your UID in the profile picture in the upper-right corner of the DMS console.
	//
	// example:
	//
	// 1139477549527134
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CloseTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseTicketRequest) GoString() string {
	return s.String()
}

func (s *CloseTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *CloseTicketRequest) GetUid() *string {
	return s.Uid
}

func (s *CloseTicketRequest) SetTicketId(v string) *CloseTicketRequest {
	s.TicketId = &v
	return s
}

func (s *CloseTicketRequest) SetUid(v string) *CloseTicketRequest {
	s.Uid = &v
	return s
}

func (s *CloseTicketRequest) Validate() error {
	return dara.Validate(s)
}
