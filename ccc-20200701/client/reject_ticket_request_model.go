// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RejectTicketRequest
	GetComment() *string
	SetInstanceId(v string) *RejectTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *RejectTicketRequest
	GetTicketId() *string
}

type RejectTicketRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// f2c6722b-cd13-442d-bf10-22a07c70d6d5
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s RejectTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectTicketRequest) GoString() string {
	return s.String()
}

func (s *RejectTicketRequest) GetComment() *string {
	return s.Comment
}

func (s *RejectTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RejectTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *RejectTicketRequest) SetComment(v string) *RejectTicketRequest {
	s.Comment = &v
	return s
}

func (s *RejectTicketRequest) SetInstanceId(v string) *RejectTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectTicketRequest) SetTicketId(v string) *RejectTicketRequest {
	s.TicketId = &v
	return s
}

func (s *RejectTicketRequest) Validate() error {
	return dara.Validate(s)
}
