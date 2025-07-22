// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMPUTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopMPUTaskResponseBody
	GetRequestId() *string
}

type StopMPUTaskResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMPUTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopMPUTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMPUTaskResponseBody) SetRequestId(v string) *StopMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMPUTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
