// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestSGResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteRequestSGResponseBody
  GetCode() *int32 
  SetData(v map[string]interface{}) *ExecuteRequestSGResponseBody
  GetData() map[string]interface{} 
  SetMessage(v string) *ExecuteRequestSGResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteRequestSGResponseBody
  GetRequestId() *string 
}

type ExecuteRequestSGResponseBody struct {
  // The status code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The data information, including fields such as tags, score, and extend.
  // 
  // example:
  // 
  // {"score": 80.4, "tags": "ra0601,rnk0002,rn0312"}
  Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
  // The result message.
  // 
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // A5F41DC6-2070-5CE7-8B2D-7A23D882932D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteRequestSGResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestSGResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteRequestSGResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteRequestSGResponseBody) GetData() map[string]interface{}  {
  return s.Data
}

func (s *ExecuteRequestSGResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteRequestSGResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteRequestSGResponseBody) SetCode(v int32) *ExecuteRequestSGResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteRequestSGResponseBody) SetData(v map[string]interface{}) *ExecuteRequestSGResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteRequestSGResponseBody) SetMessage(v string) *ExecuteRequestSGResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteRequestSGResponseBody) SetRequestId(v string) *ExecuteRequestSGResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteRequestSGResponseBody) Validate() error {
  return dara.Validate(s)
}

