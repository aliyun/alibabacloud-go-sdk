// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCallSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *UpdateCallSummaryRequest
	GetContext() *string
	SetInstanceId(v string) *UpdateCallSummaryRequest
	GetInstanceId() *string
	SetTicketId(v string) *UpdateCallSummaryRequest
	GetTicketId() *string
}

type UpdateCallSummaryRequest struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
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

func (s UpdateCallSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCallSummaryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCallSummaryRequest) GetContext() *string {
	return s.Context
}

func (s *UpdateCallSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCallSummaryRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *UpdateCallSummaryRequest) SetContext(v string) *UpdateCallSummaryRequest {
	s.Context = &v
	return s
}

func (s *UpdateCallSummaryRequest) SetInstanceId(v string) *UpdateCallSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCallSummaryRequest) SetTicketId(v string) *UpdateCallSummaryRequest {
	s.TicketId = &v
	return s
}

func (s *UpdateCallSummaryRequest) Validate() error {
	return dara.Validate(s)
}
