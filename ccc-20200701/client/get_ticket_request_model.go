// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *GetTicketRequest
	GetTicketId() *string
}

type GetTicketRequest struct {
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
	// 3d26b90a-c5d2-4b09-8219-********
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s GetTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *GetTicketRequest) SetInstanceId(v string) *GetTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTicketRequest) SetTicketId(v string) *GetTicketRequest {
	s.TicketId = &v
	return s
}

func (s *GetTicketRequest) Validate() error {
	return dara.Validate(s)
}
