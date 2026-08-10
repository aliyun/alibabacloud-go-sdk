// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDDLResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteDDLResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *ExecuteDDLResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecuteDDLResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteDDLResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteDDLResponseBody
  GetSuccess() *bool 
}

type ExecuteDDLResponseBody struct {
  // The backend response code.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The backend exception details.
  // 
  // example:
  // 
  // internal error
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteDDLResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteDDLResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteDDLResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteDDLResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteDDLResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteDDLResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteDDLResponseBody) SetCode(v string) *ExecuteDDLResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteDDLResponseBody) SetHttpStatusCode(v int32) *ExecuteDDLResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteDDLResponseBody) SetMessage(v string) *ExecuteDDLResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteDDLResponseBody) SetRequestId(v string) *ExecuteDDLResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteDDLResponseBody) SetSuccess(v bool) *ExecuteDDLResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteDDLResponseBody) Validate() error {
  return dara.Validate(s)
}

