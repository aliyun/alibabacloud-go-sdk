// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *UpdateTicketRequest
	GetContext() *string
	SetCustomerId(v string) *UpdateTicketRequest
	GetCustomerId() *string
	SetInstanceId(v string) *UpdateTicketRequest
	GetInstanceId() *string
	SetTicketId(v string) *UpdateTicketRequest
	GetTicketId() *string
	SetTitle(v string) *UpdateTicketRequest
	GetTitle() *string
}

type UpdateTicketRequest struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 51e155ce-***-****-b402-13c69597b920
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
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
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) GetContext() *string {
	return s.Context
}

func (s *UpdateTicketRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *UpdateTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *UpdateTicketRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateTicketRequest) SetContext(v string) *UpdateTicketRequest {
	s.Context = &v
	return s
}

func (s *UpdateTicketRequest) SetCustomerId(v string) *UpdateTicketRequest {
	s.CustomerId = &v
	return s
}

func (s *UpdateTicketRequest) SetInstanceId(v string) *UpdateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTicketRequest) SetTicketId(v string) *UpdateTicketRequest {
	s.TicketId = &v
	return s
}

func (s *UpdateTicketRequest) SetTitle(v string) *UpdateTicketRequest {
	s.Title = &v
	return s
}

func (s *UpdateTicketRequest) Validate() error {
	return dara.Validate(s)
}
