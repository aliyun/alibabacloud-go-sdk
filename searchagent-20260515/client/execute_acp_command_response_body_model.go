// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAcpCommandResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteAcpCommandResponseBody
  GetCode() *string 
  SetData(v *ExecuteAcpCommandResponseBodyData) *ExecuteAcpCommandResponseBody
  GetData() *ExecuteAcpCommandResponseBodyData 
  SetMessage(v string) *ExecuteAcpCommandResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteAcpCommandResponseBody
  GetRequestId() *string 
}

type ExecuteAcpCommandResponseBody struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Data *ExecuteAcpCommandResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteAcpCommandResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAcpCommandResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAcpCommandResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteAcpCommandResponseBody) GetData() *ExecuteAcpCommandResponseBodyData  {
  return s.Data
}

func (s *ExecuteAcpCommandResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteAcpCommandResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAcpCommandResponseBody) SetCode(v string) *ExecuteAcpCommandResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteAcpCommandResponseBody) SetData(v *ExecuteAcpCommandResponseBodyData) *ExecuteAcpCommandResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAcpCommandResponseBody) SetMessage(v string) *ExecuteAcpCommandResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteAcpCommandResponseBody) SetRequestId(v string) *ExecuteAcpCommandResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAcpCommandResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAcpCommandResponseBodyData struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Jsonrpc *string `json:"jsonrpc,omitempty" xml:"jsonrpc,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
  Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ExecuteAcpCommandResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAcpCommandResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAcpCommandResponseBodyData) GetId() *string  {
  return s.Id
}

func (s *ExecuteAcpCommandResponseBodyData) GetJsonrpc() *string  {
  return s.Jsonrpc
}

func (s *ExecuteAcpCommandResponseBodyData) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAcpCommandResponseBodyData) GetResult() map[string]interface{}  {
  return s.Result
}

func (s *ExecuteAcpCommandResponseBodyData) GetTimestamp() *int64  {
  return s.Timestamp
}

func (s *ExecuteAcpCommandResponseBodyData) SetId(v string) *ExecuteAcpCommandResponseBodyData {
  s.Id = &v
  return s
}

func (s *ExecuteAcpCommandResponseBodyData) SetJsonrpc(v string) *ExecuteAcpCommandResponseBodyData {
  s.Jsonrpc = &v
  return s
}

func (s *ExecuteAcpCommandResponseBodyData) SetRequestId(v string) *ExecuteAcpCommandResponseBodyData {
  s.RequestId = &v
  return s
}

func (s *ExecuteAcpCommandResponseBodyData) SetResult(v map[string]interface{}) *ExecuteAcpCommandResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteAcpCommandResponseBodyData) SetTimestamp(v int64) *ExecuteAcpCommandResponseBodyData {
  s.Timestamp = &v
  return s
}

func (s *ExecuteAcpCommandResponseBodyData) Validate() error {
  return dara.Validate(s)
}

