// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEduRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEduRoomResponseBody
	GetRequestId() *string
}

type DeleteEduRoomResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEduRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEduRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEduRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEduRoomResponseBody) SetRequestId(v string) *DeleteEduRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEduRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
