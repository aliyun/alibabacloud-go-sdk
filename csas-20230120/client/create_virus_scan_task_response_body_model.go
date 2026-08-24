// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirusScanTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateVirusScanTaskResponseBody
	GetTaskId() *string
}

type CreateVirusScanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created virus scan task.
	//
	// example:
	//
	// v1:1024772
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVirusScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVirusScanTaskResponseBody) SetRequestId(v string) *CreateVirusScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanTaskResponseBody) SetTaskId(v string) *CreateVirusScanTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVirusScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
