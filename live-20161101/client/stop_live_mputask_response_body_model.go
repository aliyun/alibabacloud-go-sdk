// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveMPUTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLiveMPUTaskResponseBody
	GetRequestId() *string
}

type StopLiveMPUTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0F72851F-5DC1-1979-9B2C-450040316C3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveMPUTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLiveMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveMPUTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLiveMPUTaskResponseBody) SetRequestId(v string) *StopLiveMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLiveMPUTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
