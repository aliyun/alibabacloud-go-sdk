// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcomVideoRecreationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EcomVideoRecreationResponseBody
  GetCode() *string 
  SetData(v *EcomVideoRecreationResponseBodyData) *EcomVideoRecreationResponseBody
  GetData() *EcomVideoRecreationResponseBodyData 
  SetMessage(v string) *EcomVideoRecreationResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EcomVideoRecreationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EcomVideoRecreationResponseBody
  GetSuccess() *bool 
}

type EcomVideoRecreationResponseBody struct {
  // The result code. `success` indicates success. An error code is returned upon failure.
  // 
  // example:
  // 
  // success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The asynchronous task submit status.
  Data *EcomVideoRecreationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The response message. An error description is returned upon failure.
  // 
  // example:
  // 
  // Task submitted
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID, used to identify a unique call.
  // 
  // example:
  // 
  // 70CBEFDF-BB17-1EB3-8A21-569F3124738F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the submission is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EcomVideoRecreationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationResponseBody) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EcomVideoRecreationResponseBody) GetData() *EcomVideoRecreationResponseBodyData  {
  return s.Data
}

func (s *EcomVideoRecreationResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EcomVideoRecreationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EcomVideoRecreationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EcomVideoRecreationResponseBody) SetCode(v string) *EcomVideoRecreationResponseBody {
  s.Code = &v
  return s
}

func (s *EcomVideoRecreationResponseBody) SetData(v *EcomVideoRecreationResponseBodyData) *EcomVideoRecreationResponseBody {
  s.Data = v
  return s
}

func (s *EcomVideoRecreationResponseBody) SetMessage(v string) *EcomVideoRecreationResponseBody {
  s.Message = &v
  return s
}

func (s *EcomVideoRecreationResponseBody) SetRequestId(v string) *EcomVideoRecreationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EcomVideoRecreationResponseBody) SetSuccess(v bool) *EcomVideoRecreationResponseBody {
  s.Success = &v
  return s
}

func (s *EcomVideoRecreationResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EcomVideoRecreationResponseBodyData struct {
  // The asynchronous task ID for QueryAsyncTaskResult queries.
  // 
  // example:
  // 
  // task_778fa8bd21804828a5d147050e30edac
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EcomVideoRecreationResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationResponseBodyData) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationResponseBodyData) GetTaskId() *string  {
  return s.TaskId
}

func (s *EcomVideoRecreationResponseBodyData) SetTaskId(v string) *EcomVideoRecreationResponseBodyData {
  s.TaskId = &v
  return s
}

func (s *EcomVideoRecreationResponseBodyData) Validate() error {
  return dara.Validate(s)
}

