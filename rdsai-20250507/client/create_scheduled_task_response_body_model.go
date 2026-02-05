// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScheduledTaskResponseBody
	GetRequestId() *string
	SetScheduledId(v string) *CreateScheduledTaskResponseBody
	GetScheduledId() *string
	SetSuccess(v bool) *CreateScheduledTaskResponseBody
	GetSuccess() *bool
}

type CreateScheduledTaskResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D984FD38-6C2D-55DF-B0D7-8BCAC2E1F8C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 847268a4-196f-416b-aa12-bfe0c115****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledTaskResponseBody) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *CreateScheduledTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScheduledTaskResponseBody) SetMessage(v string) *CreateScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetScheduledId(v string) *CreateScheduledTaskResponseBody {
	s.ScheduledId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetSuccess(v bool) *CreateScheduledTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
