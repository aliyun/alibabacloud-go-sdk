// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailable(v bool) *CheckDeviceResponseBody
	GetAvailable() *bool
	SetRequestId(v string) *CheckDeviceResponseBody
	GetRequestId() *string
}

type CheckDeviceResponseBody struct {
	// example:
	//
	// ture
	Available *bool `json:"Available,omitempty" xml:"Available,omitempty"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDeviceResponseBody) GetAvailable() *bool {
	return s.Available
}

func (s *CheckDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDeviceResponseBody) SetAvailable(v bool) *CheckDeviceResponseBody {
	s.Available = &v
	return s
}

func (s *CheckDeviceResponseBody) SetRequestId(v string) *CheckDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
