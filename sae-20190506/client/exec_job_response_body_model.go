// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecJobResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecJobResponseBody
  GetCode() *string 
  SetData(v *ExecJobResponseBodyData) *ExecJobResponseBody
  GetData() *ExecJobResponseBodyData 
  SetErrorCode(v string) *ExecJobResponseBody
  GetErrorCode() *string 
  SetMessage(v string) *ExecJobResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecJobResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecJobResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ExecJobResponseBody
  GetTraceId() *string 
}

type ExecJobResponseBody struct {
  // The HTTP status code or a POP error code.
  // 
  // - **2xx**: Success.
  // 
  // - **3xx**: Redirection.
  // 
  // - **4xx**: Request error.
  // 
  // - **5xx**: Server error.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned data.
  Data *ExecJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The error code.
  // 
  // - This parameter is omitted for successful requests.
  // 
  // - This parameter is included for failed requests. For more information, see the **Error codes*	- section of this topic.
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The returned message.
  // 
  // - If the request is successful, **success*	- is returned.
  // 
  // - If the request fails, an error code is returned.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 67DD9A98-9CCC-5BE8-8C9E-B45E72F4****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // - **true**: The request was successful.
  // 
  // - **false**: The request failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // The trace ID for retrieving detailed information about the call.
  // 
  // example:
  // 
  // 0b87b7e716575071334387401e****
  TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ExecJobResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecJobResponseBody) GoString() string {
  return s.String()
}

func (s *ExecJobResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecJobResponseBody) GetData() *ExecJobResponseBodyData  {
  return s.Data
}

func (s *ExecJobResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecJobResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecJobResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecJobResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecJobResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ExecJobResponseBody) SetCode(v string) *ExecJobResponseBody {
  s.Code = &v
  return s
}

func (s *ExecJobResponseBody) SetData(v *ExecJobResponseBodyData) *ExecJobResponseBody {
  s.Data = v
  return s
}

func (s *ExecJobResponseBody) SetErrorCode(v string) *ExecJobResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecJobResponseBody) SetMessage(v string) *ExecJobResponseBody {
  s.Message = &v
  return s
}

func (s *ExecJobResponseBody) SetRequestId(v string) *ExecJobResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecJobResponseBody) SetSuccess(v bool) *ExecJobResponseBody {
  s.Success = &v
  return s
}

func (s *ExecJobResponseBody) SetTraceId(v string) *ExecJobResponseBody {
  s.TraceId = &v
  return s
}

func (s *ExecJobResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecJobResponseBodyData struct {
  // The HTTP status code or a POP error code.
  // 
  // - **2xx**: Success.
  // 
  // - **3xx**: Redirection.
  // 
  // - **4xx**: Request error.
  // 
  // - **5xx**: Server error.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The job ID.
  // 
  // example:
  // 
  // manual-3db7a8fa-5d40-4edc-92e4-49d50eab****
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The returned message.
  // 
  // - If the request is successful, **success*	- is returned.
  // 
  // - If the request fails, an error code is returned.
  // 
  // example:
  // 
  // success
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // Whether the job was successfully executed.
  // 
  // - **true**: The execution was successful.
  // 
  // - **false**: The execution failed.
  // 
  // example:
  // 
  // true
  Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecJobResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecJobResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecJobResponseBodyData) GetCode() *string  {
  return s.Code
}

func (s *ExecJobResponseBodyData) GetData() *string  {
  return s.Data
}

func (s *ExecJobResponseBodyData) GetMsg() *string  {
  return s.Msg
}

func (s *ExecJobResponseBodyData) GetSuccess() *string  {
  return s.Success
}

func (s *ExecJobResponseBodyData) SetCode(v string) *ExecJobResponseBodyData {
  s.Code = &v
  return s
}

func (s *ExecJobResponseBodyData) SetData(v string) *ExecJobResponseBodyData {
  s.Data = &v
  return s
}

func (s *ExecJobResponseBodyData) SetMsg(v string) *ExecJobResponseBodyData {
  s.Msg = &v
  return s
}

func (s *ExecJobResponseBodyData) SetSuccess(v string) *ExecJobResponseBodyData {
  s.Success = &v
  return s
}

func (s *ExecJobResponseBodyData) Validate() error {
  return dara.Validate(s)
}

