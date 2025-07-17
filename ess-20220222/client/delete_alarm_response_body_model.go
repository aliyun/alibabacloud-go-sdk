// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmTaskId(v string) *DeleteAlarmResponseBody
	GetAlarmTaskId() *string
	SetRequestId(v string) *DeleteAlarmResponseBody
	GetRequestId() *string
}

type DeleteAlarmResponseBody struct {
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
	// 6EF9BFEE-FE07-4627-B8FB-14326FB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmResponseBody) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *DeleteAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlarmResponseBody) SetAlarmTaskId(v string) *DeleteAlarmResponseBody {
	s.AlarmTaskId = &v
	return s
}

func (s *DeleteAlarmResponseBody) SetRequestId(v string) *DeleteAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
