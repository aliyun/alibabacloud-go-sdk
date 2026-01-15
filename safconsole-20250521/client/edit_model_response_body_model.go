// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditModelResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int64) *EditModelResponseBody
  GetCode() *int64 
  SetHttpStatusCode(v int64) *EditModelResponseBody
  GetHttpStatusCode() *int64 
  SetRequestId(v string) *EditModelResponseBody
  GetRequestId() *string 
  SetResultObject(v bool) *EditModelResponseBody
  GetResultObject() *bool 
  SetSuccess(v bool) *EditModelResponseBody
  GetSuccess() *bool 
}

type EditModelResponseBody struct {
  // example:
  // 
  // 200
  Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditModelResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditModelResponseBody) GoString() string {
  return s.String()
}

func (s *EditModelResponseBody) GetCode() *int64  {
  return s.Code
}

func (s *EditModelResponseBody) GetHttpStatusCode() *int64  {
  return s.HttpStatusCode
}

func (s *EditModelResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditModelResponseBody) GetResultObject() *bool  {
  return s.ResultObject
}

func (s *EditModelResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditModelResponseBody) SetCode(v int64) *EditModelResponseBody {
  s.Code = &v
  return s
}

func (s *EditModelResponseBody) SetHttpStatusCode(v int64) *EditModelResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EditModelResponseBody) SetRequestId(v string) *EditModelResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditModelResponseBody) SetResultObject(v bool) *EditModelResponseBody {
  s.ResultObject = &v
  return s
}

func (s *EditModelResponseBody) SetSuccess(v bool) *EditModelResponseBody {
  s.Success = &v
  return s
}

func (s *EditModelResponseBody) Validate() error {
  return dara.Validate(s)
}

