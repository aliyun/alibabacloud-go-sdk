// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceAdbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopInstanceAdbResponseBody
	GetRequestId() *string
}

type StopInstanceAdbResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceAdbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceAdbResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceAdbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstanceAdbResponseBody) SetRequestId(v string) *StopInstanceAdbResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceAdbResponseBody) Validate() error {
	return dara.Validate(s)
}
