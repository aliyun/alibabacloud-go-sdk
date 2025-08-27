// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityDeleteResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntityDeleteResponseBody
  GetCode() *string 
  SetMessage(v string) *EntityDeleteResponseBody
  GetMessage() *string 
  SetModule(v *EntityDeleteResponseBodyModule) *EntityDeleteResponseBody
  GetModule() *EntityDeleteResponseBodyModule 
  SetMorePage(v bool) *EntityDeleteResponseBody
  GetMorePage() *bool 
  SetRequestId(v string) *EntityDeleteResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EntityDeleteResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *EntityDeleteResponseBody
  GetTraceId() *string 
}

type EntityDeleteResponseBody struct {
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *EntityDeleteResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s EntityDeleteResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteResponseBody) GoString() string {
  return s.String()
}

func (s *EntityDeleteResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntityDeleteResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntityDeleteResponseBody) GetModule() *EntityDeleteResponseBodyModule  {
  return s.Module
}

func (s *EntityDeleteResponseBody) GetMorePage() *bool  {
  return s.MorePage
}

func (s *EntityDeleteResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntityDeleteResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EntityDeleteResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *EntityDeleteResponseBody) SetCode(v string) *EntityDeleteResponseBody {
  s.Code = &v
  return s
}

func (s *EntityDeleteResponseBody) SetMessage(v string) *EntityDeleteResponseBody {
  s.Message = &v
  return s
}

func (s *EntityDeleteResponseBody) SetModule(v *EntityDeleteResponseBodyModule) *EntityDeleteResponseBody {
  s.Module = v
  return s
}

func (s *EntityDeleteResponseBody) SetMorePage(v bool) *EntityDeleteResponseBody {
  s.MorePage = &v
  return s
}

func (s *EntityDeleteResponseBody) SetRequestId(v string) *EntityDeleteResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntityDeleteResponseBody) SetSuccess(v bool) *EntityDeleteResponseBody {
  s.Success = &v
  return s
}

func (s *EntityDeleteResponseBody) SetTraceId(v string) *EntityDeleteResponseBody {
  s.TraceId = &v
  return s
}

func (s *EntityDeleteResponseBody) Validate() error {
  return dara.Validate(s)
}

type EntityDeleteResponseBodyModule struct {
  // example:
  // 
  // 1
  RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
  // example:
  // 
  // 1
  SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s EntityDeleteResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteResponseBodyModule) GoString() string {
  return s.String()
}

func (s *EntityDeleteResponseBodyModule) GetRemoveNum() *int32  {
  return s.RemoveNum
}

func (s *EntityDeleteResponseBodyModule) GetSelectedUserNum() *int32  {
  return s.SelectedUserNum
}

func (s *EntityDeleteResponseBodyModule) SetRemoveNum(v int32) *EntityDeleteResponseBodyModule {
  s.RemoveNum = &v
  return s
}

func (s *EntityDeleteResponseBodyModule) SetSelectedUserNum(v int32) *EntityDeleteResponseBodyModule {
  s.SelectedUserNum = &v
  return s
}

func (s *EntityDeleteResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

