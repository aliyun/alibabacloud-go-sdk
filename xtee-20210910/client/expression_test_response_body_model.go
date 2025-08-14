// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpressionTestResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExpressionTestResponseBody
  GetRequestId() *string 
  SetResultObject(v bool) *ExpressionTestResponseBody
  GetResultObject() *bool 
}

type ExpressionTestResponseBody struct {
  // Request ID.
  // 
  // example:
  // 
  // A32FE941-35F2-5378-B37C-4B8FDB16F094
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Return object
  // 
  // example:
  // 
  // true
  ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s ExpressionTestResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExpressionTestResponseBody) GoString() string {
  return s.String()
}

func (s *ExpressionTestResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExpressionTestResponseBody) GetResultObject() *bool  {
  return s.ResultObject
}

func (s *ExpressionTestResponseBody) SetRequestId(v string) *ExpressionTestResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExpressionTestResponseBody) SetResultObject(v bool) *ExpressionTestResponseBody {
  s.ResultObject = &v
  return s
}

func (s *ExpressionTestResponseBody) Validate() error {
  return dara.Validate(s)
}

