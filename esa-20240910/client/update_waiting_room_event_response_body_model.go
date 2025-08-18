// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWaitingRoomEventResponseBody
	GetRequestId() *string
}

type UpdateWaitingRoomEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0195619f-eab3-4a66-ac00-ed53d913e72e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaitingRoomEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomEventResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWaitingRoomEventResponseBody) SetRequestId(v string) *UpdateWaitingRoomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWaitingRoomEventResponseBody) Validate() error {
	return dara.Validate(s)
}
