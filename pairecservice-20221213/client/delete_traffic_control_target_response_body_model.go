// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficControlTargetResponseBody
	GetRequestId() *string
}

type DeleteTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficControlTargetResponseBody) SetRequestId(v string) *DeleteTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficControlTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
