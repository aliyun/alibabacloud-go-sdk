// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInvocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopInvocationResponseBody
	GetRequestId() *string
}

type StopInvocationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInvocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *StopInvocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInvocationResponseBody) SetRequestId(v string) *StopInvocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInvocationResponseBody) Validate() error {
	return dara.Validate(s)
}
