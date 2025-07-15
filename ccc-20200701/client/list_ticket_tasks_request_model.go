// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTicketTasksRequest
	GetInstanceId() *string
	SetTicketId(v string) *ListTicketTasksRequest
	GetTicketId() *string
}

type ListTicketTasksRequest struct {
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

func (s ListTicketTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTicketTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTicketTasksRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketTasksRequest) SetInstanceId(v string) *ListTicketTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTicketTasksRequest) SetTicketId(v string) *ListTicketTasksRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketTasksRequest) Validate() error {
	return dara.Validate(s)
}
