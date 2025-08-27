// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityAddResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntityAddResponseBody
  GetCode() *string 
  SetMessage(v string) *EntityAddResponseBody
  GetMessage() *string 
  SetModule(v *EntityAddResponseBodyModule) *EntityAddResponseBody
  GetModule() *EntityAddResponseBodyModule 
  SetRequestId(v string) *EntityAddResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EntityAddResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *EntityAddResponseBody
  GetTraceId() *string 
}

type EntityAddResponseBody struct {
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *EntityAddResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  // example:
  // 
  // B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // traceId
  // 
  // example:
  // 
  // 21041ce316577904808056433edbb2
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s EntityAddResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntityAddResponseBody) GoString() string {
  return s.String()
}

func (s *EntityAddResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntityAddResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntityAddResponseBody) GetModule() *EntityAddResponseBodyModule  {
  return s.Module
}

func (s *EntityAddResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntityAddResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EntityAddResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *EntityAddResponseBody) SetCode(v string) *EntityAddResponseBody {
  s.Code = &v
  return s
}

func (s *EntityAddResponseBody) SetMessage(v string) *EntityAddResponseBody {
  s.Message = &v
  return s
}

func (s *EntityAddResponseBody) SetModule(v *EntityAddResponseBodyModule) *EntityAddResponseBody {
  s.Module = v
  return s
}

func (s *EntityAddResponseBody) SetRequestId(v string) *EntityAddResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntityAddResponseBody) SetSuccess(v bool) *EntityAddResponseBody {
  s.Success = &v
  return s
}

func (s *EntityAddResponseBody) SetTraceId(v string) *EntityAddResponseBody {
  s.TraceId = &v
  return s
}

func (s *EntityAddResponseBody) Validate() error {
  return dara.Validate(s)
}

type EntityAddResponseBodyModule struct {
  // example:
  // 
  // 1
  AddNum *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
  // example:
  // 
  // 2
  SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s EntityAddResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s EntityAddResponseBodyModule) GoString() string {
  return s.String()
}

func (s *EntityAddResponseBodyModule) GetAddNum() *int32  {
  return s.AddNum
}

func (s *EntityAddResponseBodyModule) GetSelectedUserNum() *int32  {
  return s.SelectedUserNum
}

func (s *EntityAddResponseBodyModule) SetAddNum(v int32) *EntityAddResponseBodyModule {
  s.AddNum = &v
  return s
}

func (s *EntityAddResponseBodyModule) SetSelectedUserNum(v int32) *EntityAddResponseBodyModule {
  s.SelectedUserNum = &v
  return s
}

func (s *EntityAddResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

