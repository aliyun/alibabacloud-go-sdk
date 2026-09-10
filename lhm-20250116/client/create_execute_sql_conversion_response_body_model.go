// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecuteSqlConversionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateExecuteSqlConversionResponseBodyData) *CreateExecuteSqlConversionResponseBody
	GetData() *CreateExecuteSqlConversionResponseBodyData
	SetErrCode(v string) *CreateExecuteSqlConversionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateExecuteSqlConversionResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateExecuteSqlConversionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExecuteSqlConversionResponseBody
	GetSuccess() *bool
}

type CreateExecuteSqlConversionResponseBody struct {
	// The data body returned by the operation. For the field structure, see the child field descriptions.
	Data *CreateExecuteSqlConversionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Indicates whether the call is successful. Valid values: true (the call is successful) and false (the call failed). If the call failed, use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateExecuteSqlConversionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExecuteSqlConversionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExecuteSqlConversionResponseBody) GetData() *CreateExecuteSqlConversionResponseBodyData {
	return s.Data
}

func (s *CreateExecuteSqlConversionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateExecuteSqlConversionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateExecuteSqlConversionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExecuteSqlConversionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExecuteSqlConversionResponseBody) SetData(v *CreateExecuteSqlConversionResponseBodyData) *CreateExecuteSqlConversionResponseBody {
	s.Data = v
	return s
}

func (s *CreateExecuteSqlConversionResponseBody) SetErrCode(v string) *CreateExecuteSqlConversionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateExecuteSqlConversionResponseBody) SetErrMessage(v string) *CreateExecuteSqlConversionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateExecuteSqlConversionResponseBody) SetRequestId(v string) *CreateExecuteSqlConversionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExecuteSqlConversionResponseBody) SetSuccess(v bool) *CreateExecuteSqlConversionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExecuteSqlConversionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExecuteSqlConversionResponseBodyData struct {
	// The task ID that uniquely identifies a task.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateExecuteSqlConversionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateExecuteSqlConversionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateExecuteSqlConversionResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateExecuteSqlConversionResponseBodyData) SetTaskId(v int64) *CreateExecuteSqlConversionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateExecuteSqlConversionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
