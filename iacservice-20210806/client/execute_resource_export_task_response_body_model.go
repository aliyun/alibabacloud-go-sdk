// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteResourceExportTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetExportTaskId(v string) *ExecuteResourceExportTaskResponseBody
  GetExportTaskId() *string 
  SetExportVersion(v string) *ExecuteResourceExportTaskResponseBody
  GetExportVersion() *string 
  SetRequestId(v string) *ExecuteResourceExportTaskResponseBody
  GetRequestId() *string 
}

type ExecuteResourceExportTaskResponseBody struct {
  // example:
  // 
  // ex-3b6cb9fa4751a6e645ad8365e6
  ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
  // example:
  // 
  // v1
  ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
  // example:
  // 
  // 0B0A7C19-9077-5975-ACBD-DEE718787992
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteResourceExportTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteResourceExportTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteResourceExportTaskResponseBody) GetExportTaskId() *string  {
  return s.ExportTaskId
}

func (s *ExecuteResourceExportTaskResponseBody) GetExportVersion() *string  {
  return s.ExportVersion
}

func (s *ExecuteResourceExportTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteResourceExportTaskResponseBody) SetExportTaskId(v string) *ExecuteResourceExportTaskResponseBody {
  s.ExportTaskId = &v
  return s
}

func (s *ExecuteResourceExportTaskResponseBody) SetExportVersion(v string) *ExecuteResourceExportTaskResponseBody {
  s.ExportVersion = &v
  return s
}

func (s *ExecuteResourceExportTaskResponseBody) SetRequestId(v string) *ExecuteResourceExportTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteResourceExportTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

