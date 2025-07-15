// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *WithdrawTicketRequest
	GetComment() *string
	SetInstanceId(v string) *WithdrawTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *WithdrawTicketRequest
	GetTicketId() *string
}

type WithdrawTicketRequest struct {
	// This parameter is required.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5491d3b4-14ee-4341-b5f1-db2c78beddeb
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s WithdrawTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawTicketRequest) GoString() string {
	return s.String()
}

func (s *WithdrawTicketRequest) GetComment() *string {
	return s.Comment
}

func (s *WithdrawTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *WithdrawTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *WithdrawTicketRequest) SetComment(v string) *WithdrawTicketRequest {
	s.Comment = &v
	return s
}

func (s *WithdrawTicketRequest) SetInstanceId(v string) *WithdrawTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *WithdrawTicketRequest) SetTicketId(v string) *WithdrawTicketRequest {
	s.TicketId = &v
	return s
}

func (s *WithdrawTicketRequest) Validate() error {
	return dara.Validate(s)
}
