// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *DeleteTicketRequest
	GetTicketId() *string
}

type DeleteTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 3d26b90a-c5d2-4b09-8219-********
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s DeleteTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketRequest) GoString() string {
	return s.String()
}

func (s *DeleteTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *DeleteTicketRequest) SetInstanceId(v string) *DeleteTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTicketRequest) SetTicketId(v string) *DeleteTicketRequest {
	s.TicketId = &v
	return s
}

func (s *DeleteTicketRequest) Validate() error {
	return dara.Validate(s)
}
