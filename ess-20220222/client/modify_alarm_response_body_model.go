// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmTaskId(v string) *ModifyAlarmResponseBody
	GetAlarmTaskId() *string
	SetRequestId(v string) *ModifyAlarmResponseBody
	GetRequestId() *string
}

type ModifyAlarmResponseBody struct {
	// The ID of the event-triggered task.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_83948190-acdd-483f-98f7-b77f4778****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BACACF83-7070-4953-A8FD-D81F89F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAlarmResponseBody) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *ModifyAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAlarmResponseBody) SetAlarmTaskId(v string) *ModifyAlarmResponseBody {
	s.AlarmTaskId = &v
	return s
}

func (s *ModifyAlarmResponseBody) SetRequestId(v string) *ModifyAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
