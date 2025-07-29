// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsg(v string) *StopAlertResponseBody
	GetMsg() *string
	SetStatus(v bool) *StopAlertResponseBody
	GetStatus() *bool
}

type StopAlertResponseBody struct {
	// The error message returned if the call fails.
	//
	// example:
	//
	// Success
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The operation result. Valid values:
	//
	// 	- True: The operation is successful.
	//
	// 	- False: The operation failed.
	//
	// example:
	//
	// True
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StopAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StopAlertResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *StopAlertResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *StopAlertResponseBody) SetMsg(v string) *StopAlertResponseBody {
	s.Msg = &v
	return s
}

func (s *StopAlertResponseBody) SetStatus(v bool) *StopAlertResponseBody {
	s.Status = &v
	return s
}

func (s *StopAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
