// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlExecJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSqlExecJobResponseBodyData) *CreateSqlExecJobResponseBody
	GetData() *CreateSqlExecJobResponseBodyData
	SetErrCode(v string) *CreateSqlExecJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateSqlExecJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateSqlExecJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSqlExecJobResponseBody
	GetSuccess() *bool
}

type CreateSqlExecJobResponseBody struct {
	// The data body returned by the operation. For the field structure, see the child parameters.
	Data *CreateSqlExecJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call succeeded.
	//
	// - false: The call failed. Check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSqlExecJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlExecJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlExecJobResponseBody) GetData() *CreateSqlExecJobResponseBodyData {
	return s.Data
}

func (s *CreateSqlExecJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateSqlExecJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateSqlExecJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSqlExecJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSqlExecJobResponseBody) SetData(v *CreateSqlExecJobResponseBodyData) *CreateSqlExecJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateSqlExecJobResponseBody) SetErrCode(v string) *CreateSqlExecJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSqlExecJobResponseBody) SetErrMessage(v string) *CreateSqlExecJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSqlExecJobResponseBody) SetRequestId(v string) *CreateSqlExecJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlExecJobResponseBody) SetSuccess(v bool) *CreateSqlExecJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSqlExecJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSqlExecJobResponseBodyData struct {
	// The execution job ID.
	//
	// example:
	//
	// 12345
	ExecJobId *int64 `json:"execJobId,omitempty" xml:"execJobId,omitempty"`
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the submission succeeded.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateSqlExecJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlExecJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSqlExecJobResponseBodyData) GetExecJobId() *int64 {
	return s.ExecJobId
}

func (s *CreateSqlExecJobResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateSqlExecJobResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSqlExecJobResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateSqlExecJobResponseBodyData) SetExecJobId(v int64) *CreateSqlExecJobResponseBodyData {
	s.ExecJobId = &v
	return s
}

func (s *CreateSqlExecJobResponseBodyData) SetMessage(v string) *CreateSqlExecJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateSqlExecJobResponseBodyData) SetSuccess(v bool) *CreateSqlExecJobResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreateSqlExecJobResponseBodyData) SetTaskId(v int64) *CreateSqlExecJobResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateSqlExecJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
