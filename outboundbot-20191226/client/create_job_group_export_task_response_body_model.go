// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobGroupExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateJobGroupExportTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateJobGroupExportTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateJobGroupExportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateJobGroupExportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateJobGroupExportTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateJobGroupExportTaskResponseBody
	GetTaskId() *string
}

type CreateJobGroupExportTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 744b27f3-437f-4a8c-a181-f668e492fd24
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateJobGroupExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobGroupExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateJobGroupExportTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateJobGroupExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateJobGroupExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobGroupExportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateJobGroupExportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateJobGroupExportTaskResponseBody) SetCode(v string) *CreateJobGroupExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobGroupExportTaskResponseBody) SetHttpStatusCode(v int32) *CreateJobGroupExportTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateJobGroupExportTaskResponseBody) SetMessage(v string) *CreateJobGroupExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobGroupExportTaskResponseBody) SetRequestId(v string) *CreateJobGroupExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobGroupExportTaskResponseBody) SetSuccess(v bool) *CreateJobGroupExportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobGroupExportTaskResponseBody) SetTaskId(v string) *CreateJobGroupExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateJobGroupExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
