// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskFromResourceImportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTaskFromResourceImportResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateTaskFromResourceImportResponseBody
	GetTaskId() *string
}

type CreateTaskFromResourceImportResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateTaskFromResourceImportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskFromResourceImportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskFromResourceImportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskFromResourceImportResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskFromResourceImportResponseBody) SetRequestId(v string) *CreateTaskFromResourceImportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskFromResourceImportResponseBody) SetTaskId(v string) *CreateTaskFromResourceImportResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTaskFromResourceImportResponseBody) Validate() error {
	return dara.Validate(s)
}
