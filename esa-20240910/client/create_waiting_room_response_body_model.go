// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWaitingRoomResponseBody
	GetRequestId() *string
	SetWaitingRoomId(v string) *CreateWaitingRoomResponseBody
	GetWaitingRoomId() *string
}

type CreateWaitingRoomResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s CreateWaitingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWaitingRoomResponseBody) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *CreateWaitingRoomResponseBody) SetRequestId(v string) *CreateWaitingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWaitingRoomResponseBody) SetWaitingRoomId(v string) *CreateWaitingRoomResponseBody {
	s.WaitingRoomId = &v
	return s
}

func (s *CreateWaitingRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
