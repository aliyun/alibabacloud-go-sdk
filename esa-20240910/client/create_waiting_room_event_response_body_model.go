// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWaitingRoomEventResponseBody
	GetRequestId() *string
	SetWaitingRoomEventId(v int64) *CreateWaitingRoomEventResponseBody
	GetWaitingRoomEventId() *int64
}

type CreateWaitingRoomEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WaitingRoomEventId *int64  `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
}

func (s CreateWaitingRoomEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomEventResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWaitingRoomEventResponseBody) GetWaitingRoomEventId() *int64 {
	return s.WaitingRoomEventId
}

func (s *CreateWaitingRoomEventResponseBody) SetRequestId(v string) *CreateWaitingRoomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWaitingRoomEventResponseBody) SetWaitingRoomEventId(v int64) *CreateWaitingRoomEventResponseBody {
	s.WaitingRoomEventId = &v
	return s
}

func (s *CreateWaitingRoomEventResponseBody) Validate() error {
	return dara.Validate(s)
}
