// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyScheduleTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyScheduleTaskResponseBody
	GetSuccess() *bool
}

type ModifyScheduleTaskResponseBody struct {
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScheduleTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScheduleTaskResponseBody) SetRequestId(v string) *ModifyScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduleTaskResponseBody) SetSuccess(v bool) *ModifyScheduleTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
