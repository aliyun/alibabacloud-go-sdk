// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskPopResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EditTaskPopResponseBodyData) *EditTaskPopResponseBody
  GetData() *EditTaskPopResponseBodyData 
  SetErrCode(v string) *EditTaskPopResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *EditTaskPopResponseBody
  GetErrMessage() *string 
  SetRequestId(v string) *EditTaskPopResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditTaskPopResponseBody
  GetSuccess() *bool 
}

type EditTaskPopResponseBody struct {
  // The data body returned by the operation. For the field structure, see the child parameter descriptions.
  Data *EditTaskPopResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
  // - true: Successful.
  // 
  // - false: Failed. Check errCode and errMessage for troubleshooting.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EditTaskPopResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditTaskPopResponseBody) GoString() string {
  return s.String()
}

func (s *EditTaskPopResponseBody) GetData() *EditTaskPopResponseBodyData  {
  return s.Data
}

func (s *EditTaskPopResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *EditTaskPopResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *EditTaskPopResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditTaskPopResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditTaskPopResponseBody) SetData(v *EditTaskPopResponseBodyData) *EditTaskPopResponseBody {
  s.Data = v
  return s
}

func (s *EditTaskPopResponseBody) SetErrCode(v string) *EditTaskPopResponseBody {
  s.ErrCode = &v
  return s
}

func (s *EditTaskPopResponseBody) SetErrMessage(v string) *EditTaskPopResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *EditTaskPopResponseBody) SetRequestId(v string) *EditTaskPopResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditTaskPopResponseBody) SetSuccess(v bool) *EditTaskPopResponseBody {
  s.Success = &v
  return s
}

func (s *EditTaskPopResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EditTaskPopResponseBodyData struct {
  // The file upload and parsing ID.
  // 
  // example:
  // 
  // 1001
  FileUploadParseId *int64 `json:"fileUploadParseId,omitempty" xml:"fileUploadParseId,omitempty"`
  // The primary key ID that uniquely identifies a record.
  // 
  // example:
  // 
  // 10001
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // The message content. In error scenarios, this field contains the error message. In log scenarios, this field contains the log content. In instance progress scenarios, this field may return a status or progress value.
  // 
  // example:
  // 
  // success
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // Indicates whether the call is successful. Valid values:
  // 
  // - true: Successful.
  // 
  // - false: Failed. Check errCode and errMessage for troubleshooting.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // The task ID that uniquely identifies a task.
  // 
  // example:
  // 
  // 10001
  TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s EditTaskPopResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EditTaskPopResponseBodyData) GoString() string {
  return s.String()
}

func (s *EditTaskPopResponseBodyData) GetFileUploadParseId() *int64  {
  return s.FileUploadParseId
}

func (s *EditTaskPopResponseBodyData) GetId() *int64  {
  return s.Id
}

func (s *EditTaskPopResponseBodyData) GetMessage() *string  {
  return s.Message
}

func (s *EditTaskPopResponseBodyData) GetSuccess() *bool  {
  return s.Success
}

func (s *EditTaskPopResponseBodyData) GetTaskId() *int64  {
  return s.TaskId
}

func (s *EditTaskPopResponseBodyData) SetFileUploadParseId(v int64) *EditTaskPopResponseBodyData {
  s.FileUploadParseId = &v
  return s
}

func (s *EditTaskPopResponseBodyData) SetId(v int64) *EditTaskPopResponseBodyData {
  s.Id = &v
  return s
}

func (s *EditTaskPopResponseBodyData) SetMessage(v string) *EditTaskPopResponseBodyData {
  s.Message = &v
  return s
}

func (s *EditTaskPopResponseBodyData) SetSuccess(v bool) *EditTaskPopResponseBodyData {
  s.Success = &v
  return s
}

func (s *EditTaskPopResponseBodyData) SetTaskId(v int64) *EditTaskPopResponseBodyData {
  s.TaskId = &v
  return s
}

func (s *EditTaskPopResponseBodyData) Validate() error {
  return dara.Validate(s)
}

