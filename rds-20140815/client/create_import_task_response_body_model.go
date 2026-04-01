// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateImportTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateImportTaskResponseBody
	GetTaskId() *string
}

type CreateImportTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3384382
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImportTaskResponseBody) SetRequestId(v string) *CreateImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImportTaskResponseBody) SetTaskId(v string) *CreateImportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
