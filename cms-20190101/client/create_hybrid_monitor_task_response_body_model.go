// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHybridMonitorTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHybridMonitorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHybridMonitorTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateHybridMonitorTaskResponseBody
	GetSuccess() *string
	SetTaskId(v int64) *CreateHybridMonitorTaskResponseBody
	GetTaskId() *int64
}

type CreateHybridMonitorTaskResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// yamlConfigFail
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 28CEA2E0-3E90-51B2-A7E8-B1ED75534E49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the monitoring task.
	//
	// example:
	//
	// 36****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateHybridMonitorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHybridMonitorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHybridMonitorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridMonitorTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateHybridMonitorTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateHybridMonitorTaskResponseBody) SetCode(v string) *CreateHybridMonitorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHybridMonitorTaskResponseBody) SetMessage(v string) *CreateHybridMonitorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHybridMonitorTaskResponseBody) SetRequestId(v string) *CreateHybridMonitorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridMonitorTaskResponseBody) SetSuccess(v string) *CreateHybridMonitorTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHybridMonitorTaskResponseBody) SetTaskId(v int64) *CreateHybridMonitorTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateHybridMonitorTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
