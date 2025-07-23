// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationASyncResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteOperationASyncResponseBody
  GetCode() *int32 
  SetData(v string) *ExecuteOperationASyncResponseBody
  GetData() *string 
  SetMessage(v string) *ExecuteOperationASyncResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteOperationASyncResponseBody
  GetRequestId() *string 
}

type ExecuteOperationASyncResponseBody struct {
  // Result code, 200 for success; Other representatives fail.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The operation ID. You can call the GetExecuteOperationResult operation to asynchronously query the result of an operation. The ID expires after one hour.
  // 
  // example:
  // 
  // op_xxxxxxxxxxxxxxxxxx_ecs_modifyInstanceType_BYSOQGWUV6PME412_ERMEZLXNN3K9N3OL
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // Error message
  // 
  // example:
  // 
  // " "
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Request ID
  // 
  // example:
  // 
  // SD-WEF-DSW-32ED-323DDSD-2332D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteOperationASyncResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationASyncResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteOperationASyncResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteOperationASyncResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExecuteOperationASyncResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteOperationASyncResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteOperationASyncResponseBody) SetCode(v int32) *ExecuteOperationASyncResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteOperationASyncResponseBody) SetData(v string) *ExecuteOperationASyncResponseBody {
  s.Data = &v
  return s
}

func (s *ExecuteOperationASyncResponseBody) SetMessage(v string) *ExecuteOperationASyncResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteOperationASyncResponseBody) SetRequestId(v string) *ExecuteOperationASyncResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteOperationASyncResponseBody) Validate() error {
  return dara.Validate(s)
}

