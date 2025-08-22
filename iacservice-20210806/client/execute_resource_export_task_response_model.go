// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteResourceExportTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteResourceExportTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteResourceExportTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteResourceExportTaskResponseBody) *ExecuteResourceExportTaskResponse
  GetBody() *ExecuteResourceExportTaskResponseBody 
}

type ExecuteResourceExportTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteResourceExportTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteResourceExportTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecuteResourceExportTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteResourceExportTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteResourceExportTaskResponse) GetBody() *ExecuteResourceExportTaskResponseBody  {
  return s.Body
}

func (s *ExecuteResourceExportTaskResponse) SetHeaders(v map[string]*string) *ExecuteResourceExportTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecuteResourceExportTaskResponse) SetStatusCode(v int32) *ExecuteResourceExportTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteResourceExportTaskResponse) SetBody(v *ExecuteResourceExportTaskResponseBody) *ExecuteResourceExportTaskResponse {
  s.Body = v
  return s
}

func (s *ExecuteResourceExportTaskResponse) Validate() error {
  return dara.Validate(s)
}

