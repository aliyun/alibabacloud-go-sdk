// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWaitingRoomResponseBody
	GetRequestId() *string
}

type UpdateWaitingRoomResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0195619f-eab3-4a66-ac00-ed53d913e72e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaitingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWaitingRoomResponseBody) SetRequestId(v string) *UpdateWaitingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWaitingRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
