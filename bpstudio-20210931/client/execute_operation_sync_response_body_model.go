// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationSyncResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteOperationSyncResponseBody
  GetCode() *int32 
  SetData(v *ExecuteOperationSyncResponseBodyData) *ExecuteOperationSyncResponseBody
  GetData() *ExecuteOperationSyncResponseBodyData 
  SetMessage(v string) *ExecuteOperationSyncResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteOperationSyncResponseBody
  GetRequestId() *string 
}

type ExecuteOperationSyncResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // op_xxxxxxxxxxxxxxxxxx_dds_modifyInstanceType_BYSOQGWUV6PME412_ERMEZLXNN3K9N3OL
  Data *ExecuteOperationSyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // Access key ID cannot be null.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // FDC73B05-5331-57AA-BA93-4C9882792FF5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteOperationSyncResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationSyncResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteOperationSyncResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteOperationSyncResponseBody) GetData() *ExecuteOperationSyncResponseBodyData  {
  return s.Data
}

func (s *ExecuteOperationSyncResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteOperationSyncResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteOperationSyncResponseBody) SetCode(v int32) *ExecuteOperationSyncResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteOperationSyncResponseBody) SetData(v *ExecuteOperationSyncResponseBodyData) *ExecuteOperationSyncResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteOperationSyncResponseBody) SetMessage(v string) *ExecuteOperationSyncResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteOperationSyncResponseBody) SetRequestId(v string) *ExecuteOperationSyncResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteOperationSyncResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteOperationSyncResponseBodyData struct {
  Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ExecuteOperationSyncResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationSyncResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteOperationSyncResponseBodyData) GetArguments() *string  {
  return s.Arguments
}

func (s *ExecuteOperationSyncResponseBodyData) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteOperationSyncResponseBodyData) GetOperationId() *string  {
  return s.OperationId
}

func (s *ExecuteOperationSyncResponseBodyData) GetStatus() *string  {
  return s.Status
}

func (s *ExecuteOperationSyncResponseBodyData) SetArguments(v string) *ExecuteOperationSyncResponseBodyData {
  s.Arguments = &v
  return s
}

func (s *ExecuteOperationSyncResponseBodyData) SetMessage(v string) *ExecuteOperationSyncResponseBodyData {
  s.Message = &v
  return s
}

func (s *ExecuteOperationSyncResponseBodyData) SetOperationId(v string) *ExecuteOperationSyncResponseBodyData {
  s.OperationId = &v
  return s
}

func (s *ExecuteOperationSyncResponseBodyData) SetStatus(v string) *ExecuteOperationSyncResponseBodyData {
  s.Status = &v
  return s
}

func (s *ExecuteOperationSyncResponseBodyData) Validate() error {
  return dara.Validate(s)
}

