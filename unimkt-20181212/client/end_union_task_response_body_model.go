// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndUnionTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EndUnionTaskResponseBody
  GetCode() *int32 
  SetData(v bool) *EndUnionTaskResponseBody
  GetData() *bool 
  SetErrorMsg(v string) *EndUnionTaskResponseBody
  GetErrorMsg() *string 
  SetRequestId(v string) *EndUnionTaskResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EndUnionTaskResponseBody
  GetSuccess() *bool 
}

type EndUnionTaskResponseBody struct {
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EndUnionTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndUnionTaskResponseBody) GoString() string {
  return s.String()
}

func (s *EndUnionTaskResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EndUnionTaskResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EndUnionTaskResponseBody) GetErrorMsg() *string  {
  return s.ErrorMsg
}

func (s *EndUnionTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndUnionTaskResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EndUnionTaskResponseBody) SetCode(v int32) *EndUnionTaskResponseBody {
  s.Code = &v
  return s
}

func (s *EndUnionTaskResponseBody) SetData(v bool) *EndUnionTaskResponseBody {
  s.Data = &v
  return s
}

func (s *EndUnionTaskResponseBody) SetErrorMsg(v string) *EndUnionTaskResponseBody {
  s.ErrorMsg = &v
  return s
}

func (s *EndUnionTaskResponseBody) SetRequestId(v string) *EndUnionTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndUnionTaskResponseBody) SetSuccess(v bool) *EndUnionTaskResponseBody {
  s.Success = &v
  return s
}

func (s *EndUnionTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

