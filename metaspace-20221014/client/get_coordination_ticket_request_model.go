// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoordinationTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoId(v string) *GetCoordinationTicketRequest
	GetCoId() *string
	SetTaskId(v string) *GetCoordinationTicketRequest
	GetTaskId() *string
}

type GetCoordinationTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// co-ik1nu2hxg5zbu****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// example:
	//
	// eab51156-7832-4922-9623-ff905****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCoordinationTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinationTicketRequest) GoString() string {
	return s.String()
}

func (s *GetCoordinationTicketRequest) GetCoId() *string {
	return s.CoId
}

func (s *GetCoordinationTicketRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCoordinationTicketRequest) SetCoId(v string) *GetCoordinationTicketRequest {
	s.CoId = &v
	return s
}

func (s *GetCoordinationTicketRequest) SetTaskId(v string) *GetCoordinationTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetCoordinationTicketRequest) Validate() error {
	return dara.Validate(s)
}
