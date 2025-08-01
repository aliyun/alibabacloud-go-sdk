// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHttp2ResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EnableHttp2ResponseBody
  GetCode() *int32 
  SetData(v bool) *EnableHttp2ResponseBody
  GetData() *bool 
  SetDynamicCode(v string) *EnableHttp2ResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *EnableHttp2ResponseBody
  GetDynamicMessage() *string 
  SetErrorCode(v string) *EnableHttp2ResponseBody
  GetErrorCode() *string 
  SetHttpStatusCode(v int32) *EnableHttp2ResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableHttp2ResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableHttp2ResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableHttp2ResponseBody
  GetSuccess() *bool 
}

type EnableHttp2ResponseBody struct {
  // The status code. A value of 200 is returned if the request is successful.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // Indicates whether HTTP/2 is enabled.
  // 
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The dynamic part in the error message.
  // 
  // example:
  // 
  // code
  DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
  // The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
  // 
  // >  The request parameter **DtsJobId*	- is invalid if **The Value of Input Parameter %s is not valid*	- is returned for **ErrMessage*	- and **DtsJobId*	- is returned for **DynamicMessage**.
  // 
  // example:
  // 
  // The specified parameter is invalid.
  DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
  // The status code.
  // 
  // example:
  // 
  // Success
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The message returned.
  // 
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 52BA6DA6-A702-4362-A32F-DFF79655****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- `true`
  // 
  // 	- `false`
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHttp2ResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHttp2ResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHttp2ResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableHttp2ResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableHttp2ResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *EnableHttp2ResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *EnableHttp2ResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableHttp2ResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableHttp2ResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableHttp2ResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHttp2ResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableHttp2ResponseBody) SetCode(v int32) *EnableHttp2ResponseBody {
  s.Code = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetData(v bool) *EnableHttp2ResponseBody {
  s.Data = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetDynamicCode(v string) *EnableHttp2ResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetDynamicMessage(v string) *EnableHttp2ResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetErrorCode(v string) *EnableHttp2ResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetHttpStatusCode(v int32) *EnableHttp2ResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetMessage(v string) *EnableHttp2ResponseBody {
  s.Message = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetRequestId(v string) *EnableHttp2ResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHttp2ResponseBody) SetSuccess(v bool) *EnableHttp2ResponseBody {
  s.Success = &v
  return s
}

func (s *EnableHttp2ResponseBody) Validate() error {
  return dara.Validate(s)
}

