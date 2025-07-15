// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *TerminateTicketRequest
	GetComment() *string
	SetInstanceId(v string) *TerminateTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *TerminateTicketRequest
	GetTicketId() *string
}

type TerminateTicketRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5491d3b4-14ee-4341-b5f1-db2c78beddeb
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s TerminateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateTicketRequest) GoString() string {
	return s.String()
}

func (s *TerminateTicketRequest) GetComment() *string {
	return s.Comment
}

func (s *TerminateTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *TerminateTicketRequest) SetComment(v string) *TerminateTicketRequest {
	s.Comment = &v
	return s
}

func (s *TerminateTicketRequest) SetInstanceId(v string) *TerminateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *TerminateTicketRequest) SetTicketId(v string) *TerminateTicketRequest {
	s.TicketId = &v
	return s
}

func (s *TerminateTicketRequest) Validate() error {
	return dara.Validate(s)
}
