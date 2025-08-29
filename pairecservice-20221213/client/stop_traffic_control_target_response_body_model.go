// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTrafficControlTargetResponseBody
	GetRequestId() *string
}

type StopTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTrafficControlTargetResponseBody) SetRequestId(v string) *StopTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTrafficControlTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
