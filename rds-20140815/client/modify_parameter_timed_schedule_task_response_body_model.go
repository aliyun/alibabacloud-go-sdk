// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterTimedScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyParameterTimedScheduleTaskResponseBody
	GetRequestId() *string
}

type ModifyParameterTimedScheduleTaskResponseBody struct {
	// example:
	//
	// 6EF82B07-28D2-48D1-B5D6-7E78FED277C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterTimedScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterTimedScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterTimedScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParameterTimedScheduleTaskResponseBody) SetRequestId(v string) *ModifyParameterTimedScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParameterTimedScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
