// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoordinationTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinateTicketModel(v *GetCoordinationTicketResponseBodyCoordinateTicketModel) *GetCoordinationTicketResponseBody
	GetCoordinateTicketModel() *GetCoordinationTicketResponseBodyCoordinateTicketModel
	SetRequestId(v string) *GetCoordinationTicketResponseBody
	GetRequestId() *string
}

type GetCoordinationTicketResponseBody struct {
	CoordinateTicketModel *GetCoordinationTicketResponseBodyCoordinateTicketModel `json:"CoordinateTicketModel,omitempty" xml:"CoordinateTicketModel,omitempty" type:"Struct"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCoordinationTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinationTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetCoordinationTicketResponseBody) GetCoordinateTicketModel() *GetCoordinationTicketResponseBodyCoordinateTicketModel {
	return s.CoordinateTicketModel
}

func (s *GetCoordinationTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCoordinationTicketResponseBody) SetCoordinateTicketModel(v *GetCoordinationTicketResponseBodyCoordinateTicketModel) *GetCoordinationTicketResponseBody {
	s.CoordinateTicketModel = v
	return s
}

func (s *GetCoordinationTicketResponseBody) SetRequestId(v string) *GetCoordinationTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCoordinationTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCoordinationTicketResponseBodyCoordinateTicketModel struct {
	// example:
	//
	// co-dk5xrhqlizm42****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// example:
	//
	// W0Rlc2t0b3BdDQpHV1Rva2VuPTAwT1A1bHp1RUlEdU00T0IzemdiZ****
	CoordinateTicket *string `json:"CoordinateTicket,omitempty" xml:"CoordinateTicket,omitempty"`
	// example:
	//
	// eab51156-7832-4922-9623-ff905****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetCoordinationTicketResponseBodyCoordinateTicketModel) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinationTicketResponseBodyCoordinateTicketModel) GoString() string {
	return s.String()
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) GetCoId() *string {
	return s.CoId
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) GetCoordinateTicket() *string {
	return s.CoordinateTicket
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) SetCoId(v string) *GetCoordinationTicketResponseBodyCoordinateTicketModel {
	s.CoId = &v
	return s
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) SetCoordinateTicket(v string) *GetCoordinationTicketResponseBodyCoordinateTicketModel {
	s.CoordinateTicket = &v
	return s
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) SetTaskId(v string) *GetCoordinationTicketResponseBodyCoordinateTicketModel {
	s.TaskId = &v
	return s
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) SetTaskStatus(v string) *GetCoordinationTicketResponseBodyCoordinateTicketModel {
	s.TaskStatus = &v
	return s
}

func (s *GetCoordinationTicketResponseBodyCoordinateTicketModel) Validate() error {
	return dara.Validate(s)
}
