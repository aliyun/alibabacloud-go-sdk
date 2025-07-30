// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDesktopsResponseBody
	GetRequestId() *string
}

type StopDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDesktopsResponseBody) SetRequestId(v string) *StopDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
