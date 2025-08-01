// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableProxyProtocolResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EnableProxyProtocolResponseBody
  GetCode() *int32 
  SetData(v bool) *EnableProxyProtocolResponseBody
  GetData() *bool 
  SetDynamicCode(v string) *EnableProxyProtocolResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *EnableProxyProtocolResponseBody
  GetDynamicMessage() *string 
  SetErrorCode(v string) *EnableProxyProtocolResponseBody
  GetErrorCode() *string 
  SetHttpStatusCode(v int32) *EnableProxyProtocolResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableProxyProtocolResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableProxyProtocolResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableProxyProtocolResponseBody
  GetSuccess() *bool 
}

type EnableProxyProtocolResponseBody struct {
  // The status code. A value of 200 is returned if the request is successful.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // Indicates whether the request was successful. Valid values:
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
  // The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
  // 
  // >  For example, if the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the **DtsJobId*	- parameter in the request is invalid.
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
  // The returned message.
  // 
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 316F5F64-F73D-42DC-8632-01E308B6****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- `true`: The request was successful.
  // 
  // 	- `false`: The request failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableProxyProtocolResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableProxyProtocolResponseBody) GoString() string {
  return s.String()
}

func (s *EnableProxyProtocolResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableProxyProtocolResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableProxyProtocolResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *EnableProxyProtocolResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *EnableProxyProtocolResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableProxyProtocolResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableProxyProtocolResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableProxyProtocolResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableProxyProtocolResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableProxyProtocolResponseBody) SetCode(v int32) *EnableProxyProtocolResponseBody {
  s.Code = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetData(v bool) *EnableProxyProtocolResponseBody {
  s.Data = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetDynamicCode(v string) *EnableProxyProtocolResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetDynamicMessage(v string) *EnableProxyProtocolResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetErrorCode(v string) *EnableProxyProtocolResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetHttpStatusCode(v int32) *EnableProxyProtocolResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetMessage(v string) *EnableProxyProtocolResponseBody {
  s.Message = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetRequestId(v string) *EnableProxyProtocolResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) SetSuccess(v bool) *EnableProxyProtocolResponseBody {
  s.Success = &v
  return s
}

func (s *EnableProxyProtocolResponseBody) Validate() error {
  return dara.Validate(s)
}

