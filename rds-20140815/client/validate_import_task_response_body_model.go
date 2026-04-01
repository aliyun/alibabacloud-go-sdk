// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateImportTaskResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *ValidateImportTaskResponseBody
	GetTaskId() *int64
}

type ValidateImportTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12345
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ValidateImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateImportTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ValidateImportTaskResponseBody) SetRequestId(v string) *ValidateImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateImportTaskResponseBody) SetTaskId(v int64) *ValidateImportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *ValidateImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
