// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoDealAccountTodoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseTodoDealAccountTodoResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseTodoDealAccountTodoResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseTodoDealAccountTodoResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseTodoDealAccountTodoResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseTodoDealAccountTodoResponseBody
  GetSuccess() *bool 
}

type EnterpriseTodoDealAccountTodoResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseTodoDealAccountTodoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoDealAccountTodoResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetCode(v string) *EnterpriseTodoDealAccountTodoResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetData(v bool) *EnterpriseTodoDealAccountTodoResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetMessage(v string) *EnterpriseTodoDealAccountTodoResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetRequestId(v string) *EnterpriseTodoDealAccountTodoResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetSuccess(v bool) *EnterpriseTodoDealAccountTodoResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) Validate() error {
  return dara.Validate(s)
}

