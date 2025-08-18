// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWaitingRoomResponseBody
	GetRequestId() *string
}

type DeleteWaitingRoomResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaitingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWaitingRoomResponseBody) SetRequestId(v string) *DeleteWaitingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWaitingRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
