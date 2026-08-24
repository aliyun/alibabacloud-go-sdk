// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusFileStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVirusFileStatusResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateVirusFileStatusResponseBody
	GetTaskId() *string
}

type UpdateVirusFileStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The disposal task ID. This parameter is returned when Operation is set to AdminQuarantine. An empty string is returned when Operation is set to AdminTrust. You can check the execution result on the user terminal device by using the TaskExecutionInfo field of ListVirusFileStatuses.
	//
	// example:
	//
	// v1:1024773
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateVirusFileStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirusFileStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVirusFileStatusResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateVirusFileStatusResponseBody) SetRequestId(v string) *UpdateVirusFileStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVirusFileStatusResponseBody) SetTaskId(v string) *UpdateVirusFileStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateVirusFileStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
