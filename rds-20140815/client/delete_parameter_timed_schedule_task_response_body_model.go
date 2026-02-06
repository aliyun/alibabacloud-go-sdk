// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterTimedScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParameterTimedScheduleTaskResponseBody
	GetRequestId() *string
}

type DeleteParameterTimedScheduleTaskResponseBody struct {
	// example:
	//
	// 16C62438-491B-5C02-9B49-BA924A1372A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterTimedScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterTimedScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterTimedScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterTimedScheduleTaskResponseBody) SetRequestId(v string) *DeleteParameterTimedScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterTimedScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
