// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestMLResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteRequestMLResponseBody
  GetCode() *int32 
  SetData(v map[string]interface{}) *ExecuteRequestMLResponseBody
  GetData() map[string]interface{} 
  SetMessage(v string) *ExecuteRequestMLResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteRequestMLResponseBody
  GetRequestId() *string 
}

type ExecuteRequestMLResponseBody struct {
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

func (s ExecuteRequestMLResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestMLResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteRequestMLResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteRequestMLResponseBody) GetData() map[string]interface{}  {
  return s.Data
}

func (s *ExecuteRequestMLResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteRequestMLResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteRequestMLResponseBody) SetCode(v int32) *ExecuteRequestMLResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteRequestMLResponseBody) SetData(v map[string]interface{}) *ExecuteRequestMLResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteRequestMLResponseBody) SetMessage(v string) *ExecuteRequestMLResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteRequestMLResponseBody) SetRequestId(v string) *ExecuteRequestMLResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteRequestMLResponseBody) Validate() error {
  return dara.Validate(s)
}

