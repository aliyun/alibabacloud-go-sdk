// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TerminateInstanceResponseBody
	GetCode() *string
	SetInstanceId(v string) *TerminateInstanceResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *TerminateInstanceResponseBody
	GetInstanceStatus() *string
	SetMessage(v string) *TerminateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *TerminateInstanceResponseBody
	GetRequestId() *string
	SetStateVersion(v int64) *TerminateInstanceResponseBody
	GetStateVersion() *int64
	SetTerminated(v bool) *TerminateInstanceResponseBody
	GetTerminated() *bool
}

type TerminateInstanceResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The physical instance ID.
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The status after termination.
	//
	// example:
	//
	// Terminated
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The state version.
	//
	// example:
	//
	// 2
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
	// Indicates whether the termination was successful. A value of true indicates success. If the termination fails, the error code and message are returned.
	//
	// example:
	//
	// true
	Terminated *bool `json:"terminated,omitempty" xml:"terminated,omitempty"`
}

func (s TerminateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *TerminateInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateInstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *TerminateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TerminateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateInstanceResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *TerminateInstanceResponseBody) GetTerminated() *bool {
	return s.Terminated
}

func (s *TerminateInstanceResponseBody) SetCode(v string) *TerminateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetInstanceId(v string) *TerminateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetInstanceStatus(v string) *TerminateInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetMessage(v string) *TerminateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetRequestId(v string) *TerminateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetStateVersion(v int64) *TerminateInstanceResponseBody {
	s.StateVersion = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetTerminated(v bool) *TerminateInstanceResponseBody {
	s.Terminated = &v
	return s
}

func (s *TerminateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
