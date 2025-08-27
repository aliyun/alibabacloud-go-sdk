// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntitySetResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntitySetResponseBody
  GetCode() *string 
  SetMessage(v string) *EntitySetResponseBody
  GetMessage() *string 
  SetModule(v *EntitySetResponseBodyModule) *EntitySetResponseBody
  GetModule() *EntitySetResponseBodyModule 
  SetMorePage(v bool) *EntitySetResponseBody
  GetMorePage() *bool 
  SetRequestId(v string) *EntitySetResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EntitySetResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *EntitySetResponseBody
  GetTraceId() *string 
}

type EntitySetResponseBody struct {
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *EntitySetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  // example:
  // 
  // true
  MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
  // example:
  // 
  // 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s EntitySetResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntitySetResponseBody) GoString() string {
  return s.String()
}

func (s *EntitySetResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntitySetResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntitySetResponseBody) GetModule() *EntitySetResponseBodyModule  {
  return s.Module
}

func (s *EntitySetResponseBody) GetMorePage() *bool  {
  return s.MorePage
}

func (s *EntitySetResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntitySetResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EntitySetResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *EntitySetResponseBody) SetCode(v string) *EntitySetResponseBody {
  s.Code = &v
  return s
}

func (s *EntitySetResponseBody) SetMessage(v string) *EntitySetResponseBody {
  s.Message = &v
  return s
}

func (s *EntitySetResponseBody) SetModule(v *EntitySetResponseBodyModule) *EntitySetResponseBody {
  s.Module = v
  return s
}

func (s *EntitySetResponseBody) SetMorePage(v bool) *EntitySetResponseBody {
  s.MorePage = &v
  return s
}

func (s *EntitySetResponseBody) SetRequestId(v string) *EntitySetResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntitySetResponseBody) SetSuccess(v bool) *EntitySetResponseBody {
  s.Success = &v
  return s
}

func (s *EntitySetResponseBody) SetTraceId(v string) *EntitySetResponseBody {
  s.TraceId = &v
  return s
}

func (s *EntitySetResponseBody) Validate() error {
  return dara.Validate(s)
}

type EntitySetResponseBodyModule struct {
  // example:
  // 
  // 1
  AddNum *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
  // example:
  // 
  // 1
  RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
  // example:
  // 
  // 1
  SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s EntitySetResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s EntitySetResponseBodyModule) GoString() string {
  return s.String()
}

func (s *EntitySetResponseBodyModule) GetAddNum() *int32  {
  return s.AddNum
}

func (s *EntitySetResponseBodyModule) GetRemoveNum() *int32  {
  return s.RemoveNum
}

func (s *EntitySetResponseBodyModule) GetSelectedUserNum() *int32  {
  return s.SelectedUserNum
}

func (s *EntitySetResponseBodyModule) SetAddNum(v int32) *EntitySetResponseBodyModule {
  s.AddNum = &v
  return s
}

func (s *EntitySetResponseBodyModule) SetRemoveNum(v int32) *EntitySetResponseBodyModule {
  s.RemoveNum = &v
  return s
}

func (s *EntitySetResponseBodyModule) SetSelectedUserNum(v int32) *EntitySetResponseBodyModule {
  s.SelectedUserNum = &v
  return s
}

func (s *EntitySetResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

