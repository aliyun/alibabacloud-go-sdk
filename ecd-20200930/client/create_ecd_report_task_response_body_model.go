// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEcdReportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEcdReportTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateEcdReportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEcdReportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEcdReportTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateEcdReportTaskResponseBody
	GetTaskId() *string
}

type CreateEcdReportTaskResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ret-g67ip******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateEcdReportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEcdReportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEcdReportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEcdReportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEcdReportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEcdReportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEcdReportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateEcdReportTaskResponseBody) SetCode(v string) *CreateEcdReportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEcdReportTaskResponseBody) SetMessage(v string) *CreateEcdReportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEcdReportTaskResponseBody) SetRequestId(v string) *CreateEcdReportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEcdReportTaskResponseBody) SetSuccess(v bool) *CreateEcdReportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEcdReportTaskResponseBody) SetTaskId(v string) *CreateEcdReportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateEcdReportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
