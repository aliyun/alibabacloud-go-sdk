// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWaitingRoomEventResponseBody
	GetRequestId() *string
}

type DeleteWaitingRoomEventResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaitingRoomEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWaitingRoomEventResponseBody) SetRequestId(v string) *DeleteWaitingRoomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWaitingRoomEventResponseBody) Validate() error {
	return dara.Validate(s)
}
