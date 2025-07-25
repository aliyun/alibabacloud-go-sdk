// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateScheduledTaskResponseBody
	GetRequestId() *string
	SetScheduledTaskId(v string) *CreateScheduledTaskResponseBody
	GetScheduledTaskId() *string
}

type CreateScheduledTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The globally unique ID of the scheduled task. The globally unique ID is generated by the system.
	//
	// example:
	//
	// edRtShc57WGXdt8TlPbr****
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledTaskResponseBody) GetScheduledTaskId() *string {
	return s.ScheduledTaskId
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetScheduledTaskId(v string) *CreateScheduledTaskResponseBody {
	s.ScheduledTaskId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
