// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResubmitTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ResubmitTicketRequest
	GetComment() *string
	SetInstanceId(v string) *ResubmitTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *ResubmitTicketRequest
	GetTicketId() *string
}

type ResubmitTicketRequest struct {
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
	// f2c6722b-cd13-442d-bf10-22a07c70d6d5
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ResubmitTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s ResubmitTicketRequest) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequest) GetComment() *string {
	return s.Comment
}

func (s *ResubmitTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResubmitTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ResubmitTicketRequest) SetComment(v string) *ResubmitTicketRequest {
	s.Comment = &v
	return s
}

func (s *ResubmitTicketRequest) SetInstanceId(v string) *ResubmitTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *ResubmitTicketRequest) SetTicketId(v string) *ResubmitTicketRequest {
	s.TicketId = &v
	return s
}

func (s *ResubmitTicketRequest) Validate() error {
	return dara.Validate(s)
}
