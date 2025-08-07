// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelScheduleTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelScheduleTasksResponseBody
	GetSuccess() *bool
}

type CancelScheduleTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7F2007D3-7E74-4ECB-89A8-BF130D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelScheduleTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelScheduleTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelScheduleTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelScheduleTasksResponseBody) SetRequestId(v string) *CancelScheduleTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelScheduleTasksResponseBody) SetSuccess(v bool) *CancelScheduleTasksResponseBody {
	s.Success = &v
	return s
}

func (s *CancelScheduleTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
