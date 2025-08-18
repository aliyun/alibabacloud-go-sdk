// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOtaTaskByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOtaTaskByTaskIdResponseBody
	GetCode() *string
	SetMessage(v string) *GetOtaTaskByTaskIdResponseBody
	GetMessage() *string
	SetOtaVersion(v string) *GetOtaTaskByTaskIdResponseBody
	GetOtaVersion() *string
	SetReleaseNote(v string) *GetOtaTaskByTaskIdResponseBody
	GetReleaseNote() *string
	SetRequestId(v string) *GetOtaTaskByTaskIdResponseBody
	GetRequestId() *string
	SetTaskStartTime(v string) *GetOtaTaskByTaskIdResponseBody
	GetTaskStartTime() *string
}

type GetOtaTaskByTaskIdResponseBody struct {
	// The error code.
	//
	// example:
	//
	// OtaTask.Running
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The task is running and cannot be sumitted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The OTA version.
	//
	// example:
	//
	// 0.0.1-R-20220708.110604
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// The version description.
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution time of the OTA update task. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s GetOtaTaskByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOtaTaskByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOtaTaskByTaskIdResponseBody) GetOtaVersion() *string {
	return s.OtaVersion
}

func (s *GetOtaTaskByTaskIdResponseBody) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *GetOtaTaskByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOtaTaskByTaskIdResponseBody) GetTaskStartTime() *string {
	return s.TaskStartTime
}

func (s *GetOtaTaskByTaskIdResponseBody) SetCode(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetMessage(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetOtaVersion(v string) *GetOtaTaskByTaskIdResponseBody {
	s.OtaVersion = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetReleaseNote(v string) *GetOtaTaskByTaskIdResponseBody {
	s.ReleaseNote = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetRequestId(v string) *GetOtaTaskByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetTaskStartTime(v string) *GetOtaTaskByTaskIdResponseBody {
	s.TaskStartTime = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) Validate() error {
	return dara.Validate(s)
}
