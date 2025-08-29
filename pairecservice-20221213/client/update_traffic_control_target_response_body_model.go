// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficControlTargetResponseBody
	GetRequestId() *string
}

type UpdateTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficControlTargetResponseBody) SetRequestId(v string) *UpdateTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficControlTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
