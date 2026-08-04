// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistBindsByOuterIdResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExistBindsByOuterIdResponseBody
  GetCode() *string 
  SetData(v bool) *ExistBindsByOuterIdResponseBody
  GetData() *bool 
  SetHttpCode(v string) *ExistBindsByOuterIdResponseBody
  GetHttpCode() *string 
  SetMessage(v string) *ExistBindsByOuterIdResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExistBindsByOuterIdResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExistBindsByOuterIdResponseBody
  GetSuccess() *bool 
}

type ExistBindsByOuterIdResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExistBindsByOuterIdResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExistBindsByOuterIdResponseBody) GoString() string {
  return s.String()
}

func (s *ExistBindsByOuterIdResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExistBindsByOuterIdResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExistBindsByOuterIdResponseBody) GetHttpCode() *string  {
  return s.HttpCode
}

func (s *ExistBindsByOuterIdResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExistBindsByOuterIdResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExistBindsByOuterIdResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExistBindsByOuterIdResponseBody) SetCode(v string) *ExistBindsByOuterIdResponseBody {
  s.Code = &v
  return s
}

func (s *ExistBindsByOuterIdResponseBody) SetData(v bool) *ExistBindsByOuterIdResponseBody {
  s.Data = &v
  return s
}

func (s *ExistBindsByOuterIdResponseBody) SetHttpCode(v string) *ExistBindsByOuterIdResponseBody {
  s.HttpCode = &v
  return s
}

func (s *ExistBindsByOuterIdResponseBody) SetMessage(v string) *ExistBindsByOuterIdResponseBody {
  s.Message = &v
  return s
}

func (s *ExistBindsByOuterIdResponseBody) SetRequestId(v string) *ExistBindsByOuterIdResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExistBindsByOuterIdResponseBody) SetSuccess(v bool) *ExistBindsByOuterIdResponseBody {
  s.Success = &v
  return s
}

func (s *ExistBindsByOuterIdResponseBody) Validate() error {
  return dara.Validate(s)
}

