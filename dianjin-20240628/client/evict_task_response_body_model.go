// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvictTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCost(v int64) *EvictTaskResponseBody
  GetCost() *int64 
  SetData(v string) *EvictTaskResponseBody
  GetData() *string 
  SetDataType(v string) *EvictTaskResponseBody
  GetDataType() *string 
  SetErrCode(v string) *EvictTaskResponseBody
  GetErrCode() *string 
  SetMessage(v string) *EvictTaskResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EvictTaskResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EvictTaskResponseBody
  GetSuccess() *bool 
  SetTime(v string) *EvictTaskResponseBody
  GetTime() *string 
}

type EvictTaskResponseBody struct {
  // example:
  // 
  // null
  Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
  // example:
  // 
  // 17071319
  Data *string `json:"data,omitempty" xml:"data,omitempty"`
  // example:
  // 
  // null
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
  // example:
  // 
  // 0
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // example:
  // 
  // 44BD277A-87F9-5310-8D63-3E6645F1DA85
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // example:
  // 
  // 2024-04-24 11:54:34
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EvictTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvictTaskResponseBody) GoString() string {
  return s.String()
}

func (s *EvictTaskResponseBody) GetCost() *int64  {
  return s.Cost
}

func (s *EvictTaskResponseBody) GetData() *string  {
  return s.Data
}

func (s *EvictTaskResponseBody) GetDataType() *string  {
  return s.DataType
}

func (s *EvictTaskResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *EvictTaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EvictTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvictTaskResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EvictTaskResponseBody) GetTime() *string  {
  return s.Time
}

func (s *EvictTaskResponseBody) SetCost(v int64) *EvictTaskResponseBody {
  s.Cost = &v
  return s
}

func (s *EvictTaskResponseBody) SetData(v string) *EvictTaskResponseBody {
  s.Data = &v
  return s
}

func (s *EvictTaskResponseBody) SetDataType(v string) *EvictTaskResponseBody {
  s.DataType = &v
  return s
}

func (s *EvictTaskResponseBody) SetErrCode(v string) *EvictTaskResponseBody {
  s.ErrCode = &v
  return s
}

func (s *EvictTaskResponseBody) SetMessage(v string) *EvictTaskResponseBody {
  s.Message = &v
  return s
}

func (s *EvictTaskResponseBody) SetRequestId(v string) *EvictTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvictTaskResponseBody) SetSuccess(v bool) *EvictTaskResponseBody {
  s.Success = &v
  return s
}

func (s *EvictTaskResponseBody) SetTime(v string) *EvictTaskResponseBody {
  s.Time = &v
  return s
}

func (s *EvictTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

