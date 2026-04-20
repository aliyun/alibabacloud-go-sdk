// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndSessionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EndSessionResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *EndSessionResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EndSessionResponseBody
  GetMessage() *string 
  SetParams(v []*string) *EndSessionResponseBody
  GetParams() []*string 
  SetRequestId(v string) *EndSessionResponseBody
  GetRequestId() *string 
}

type EndSessionResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // SUCCESS
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndSessionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndSessionResponseBody) GoString() string {
  return s.String()
}

func (s *EndSessionResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EndSessionResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EndSessionResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EndSessionResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *EndSessionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndSessionResponseBody) SetCode(v string) *EndSessionResponseBody {
  s.Code = &v
  return s
}

func (s *EndSessionResponseBody) SetHttpStatusCode(v int32) *EndSessionResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EndSessionResponseBody) SetMessage(v string) *EndSessionResponseBody {
  s.Message = &v
  return s
}

func (s *EndSessionResponseBody) SetParams(v []*string) *EndSessionResponseBody {
  s.Params = v
  return s
}

func (s *EndSessionResponseBody) SetRequestId(v string) *EndSessionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndSessionResponseBody) Validate() error {
  return dara.Validate(s)
}

