// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmTaskId(v string) *CreateAlarmResponseBody
	GetAlarmTaskId() *string
	SetRequestId(v string) *CreateAlarmResponseBody
	GetRequestId() *string
}

type CreateAlarmResponseBody struct {
	// The ID of the event-triggered task.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponseBody) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *CreateAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlarmResponseBody) SetAlarmTaskId(v string) *CreateAlarmResponseBody {
	s.AlarmTaskId = &v
	return s
}

func (s *CreateAlarmResponseBody) SetRequestId(v string) *CreateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
