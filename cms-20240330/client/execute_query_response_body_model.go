// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v []map[string]*string) *ExecuteQueryResponseBody
  GetData() []map[string]*string 
  SetMeta(v *ExecuteQueryResponseBodyMeta) *ExecuteQueryResponseBody
  GetMeta() *ExecuteQueryResponseBodyMeta 
  SetRequestId(v string) *ExecuteQueryResponseBody
  GetRequestId() *string 
}

type ExecuteQueryResponseBody struct {
  Data []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
  Meta *ExecuteQueryResponseBodyMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
  // example:
  // 
  // 3B311FD9-A60B-55E0-A896-A0C73*********
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteQueryResponseBody) GetData() []map[string]*string  {
  return s.Data
}

func (s *ExecuteQueryResponseBody) GetMeta() *ExecuteQueryResponseBodyMeta  {
  return s.Meta
}

func (s *ExecuteQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteQueryResponseBody) SetData(v []map[string]*string) *ExecuteQueryResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteQueryResponseBody) SetMeta(v *ExecuteQueryResponseBodyMeta) *ExecuteQueryResponseBody {
  s.Meta = v
  return s
}

func (s *ExecuteQueryResponseBody) SetRequestId(v string) *ExecuteQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteQueryResponseBody) Validate() error {
  if s.Meta != nil {
    if err := s.Meta.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteQueryResponseBodyMeta struct {
  // example:
  // 
  // 1
  AffectedRows *int32 `json:"affectedRows,omitempty" xml:"affectedRows,omitempty"`
  // example:
  // 
  // 1
  Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
  // example:
  // 
  // 1231243
  ElapsedMillisecond *int64 `json:"elapsedMillisecond,omitempty" xml:"elapsedMillisecond,omitempty"`
  // example:
  // 
  // Complete
  Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
}

func (s ExecuteQueryResponseBodyMeta) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryResponseBodyMeta) GoString() string {
  return s.String()
}

func (s *ExecuteQueryResponseBodyMeta) GetAffectedRows() *int32  {
  return s.AffectedRows
}

func (s *ExecuteQueryResponseBodyMeta) GetCount() *int32  {
  return s.Count
}

func (s *ExecuteQueryResponseBodyMeta) GetElapsedMillisecond() *int64  {
  return s.ElapsedMillisecond
}

func (s *ExecuteQueryResponseBodyMeta) GetProgress() *string  {
  return s.Progress
}

func (s *ExecuteQueryResponseBodyMeta) SetAffectedRows(v int32) *ExecuteQueryResponseBodyMeta {
  s.AffectedRows = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetCount(v int32) *ExecuteQueryResponseBodyMeta {
  s.Count = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetElapsedMillisecond(v int64) *ExecuteQueryResponseBodyMeta {
  s.ElapsedMillisecond = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) SetProgress(v string) *ExecuteQueryResponseBodyMeta {
  s.Progress = &v
  return s
}

func (s *ExecuteQueryResponseBodyMeta) Validate() error {
  return dara.Validate(s)
}

