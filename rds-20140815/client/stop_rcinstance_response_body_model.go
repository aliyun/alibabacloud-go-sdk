// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRCInstanceResponseBody
	GetRequestId() *string
}

type StopRCInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3E36DB6E-AE3B-53B6-A703-85F883FD1B2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRCInstanceResponseBody) SetRequestId(v string) *StopRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
