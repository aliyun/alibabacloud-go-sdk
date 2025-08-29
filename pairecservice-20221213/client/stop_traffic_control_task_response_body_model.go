// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTrafficControlTaskResponseBody
	GetRequestId() *string
}

type StopTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTrafficControlTaskResponseBody) SetRequestId(v string) *StopTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
