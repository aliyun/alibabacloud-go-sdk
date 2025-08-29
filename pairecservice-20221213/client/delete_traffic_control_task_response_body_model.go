// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficControlTaskResponseBody
	GetRequestId() *string
}

type DeleteTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficControlTaskResponseBody) SetRequestId(v string) *DeleteTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
