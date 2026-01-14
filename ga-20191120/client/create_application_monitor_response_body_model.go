// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateApplicationMonitorResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateApplicationMonitorResponseBody
	GetTaskId() *string
}

type CreateApplicationMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the origin probing task.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateApplicationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationMonitorResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateApplicationMonitorResponseBody) SetRequestId(v string) *CreateApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationMonitorResponseBody) SetTaskId(v string) *CreateApplicationMonitorResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateApplicationMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
