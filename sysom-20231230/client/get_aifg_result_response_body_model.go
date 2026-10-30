// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAifgResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAifgResultResponseBody
	GetCode() *string
	SetData(v *GetAifgResultResponseBodyData) *GetAifgResultResponseBody
	GetData() *GetAifgResultResponseBodyData
	SetMessage(v string) *GetAifgResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAifgResultResponseBody
	GetRequestId() *string
}

type GetAifgResultResponseBody struct {
	// The status code.
	//
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data details.
	Data *GetAifgResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The message content.
	//
	// example:
	//
	// "Execution succeeded"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAifgResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAifgResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAifgResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAifgResultResponseBody) GetData() *GetAifgResultResponseBodyData {
	return s.Data
}

func (s *GetAifgResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAifgResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAifgResultResponseBody) SetCode(v string) *GetAifgResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAifgResultResponseBody) SetData(v *GetAifgResultResponseBodyData) *GetAifgResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAifgResultResponseBody) SetMessage(v string) *GetAifgResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAifgResultResponseBody) SetRequestId(v string) *GetAifgResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAifgResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAifgResultResponseBodyData struct {
	// The diagnostic results.
	//
	// example:
	//
	// success
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// The service name.
	//
	// example:
	//
	// user-service
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The root data status.
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The data creation time.
	//
	// example:
	//
	// 2023-01-01T00:00:00Z
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The unique identifier of the task.
	//
	// example:
	//
	// 1234567890
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The training task ID.
	//
	// example:
	//
	// task-as2993fsx
	TraceTaskId *string `json:"traceTaskId,omitempty" xml:"traceTaskId,omitempty"`
}

func (s GetAifgResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAifgResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAifgResultResponseBodyData) GetResult() interface{} {
	return s.Result
}

func (s *GetAifgResultResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAifgResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAifgResultResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetAifgResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAifgResultResponseBodyData) GetTraceTaskId() *string {
	return s.TraceTaskId
}

func (s *GetAifgResultResponseBodyData) SetResult(v interface{}) *GetAifgResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetAifgResultResponseBodyData) SetServiceName(v string) *GetAifgResultResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetAifgResultResponseBodyData) SetStatus(v string) *GetAifgResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAifgResultResponseBodyData) SetSubmitTime(v string) *GetAifgResultResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *GetAifgResultResponseBodyData) SetTaskId(v string) *GetAifgResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAifgResultResponseBodyData) SetTraceTaskId(v string) *GetAifgResultResponseBodyData {
	s.TraceTaskId = &v
	return s
}

func (s *GetAifgResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
