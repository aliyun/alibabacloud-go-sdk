// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRequest interface {
  dara.Model
  String() string
  GoString() string
  SetConsoleSessionId(v string) *ExportRequest
  GetConsoleSessionId() *string 
  SetExportType(v int32) *ExportRequest
  GetExportType() *int32 
  SetInstanceId(v string) *ExportRequest
  GetInstanceId() *string 
  SetVhostName(v string) *ExportRequest
  GetVhostName() *string 
}

type ExportRequest struct {
  ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
  // This parameter is required.
  ExportType *int32 `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // This parameter is required.
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ExportRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRequest) GoString() string {
  return s.String()
}

func (s *ExportRequest) GetConsoleSessionId() *string  {
  return s.ConsoleSessionId
}

func (s *ExportRequest) GetExportType() *int32  {
  return s.ExportType
}

func (s *ExportRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportRequest) GetVhostName() *string  {
  return s.VhostName
}

func (s *ExportRequest) SetConsoleSessionId(v string) *ExportRequest {
  s.ConsoleSessionId = &v
  return s
}

func (s *ExportRequest) SetExportType(v int32) *ExportRequest {
  s.ExportType = &v
  return s
}

func (s *ExportRequest) SetInstanceId(v string) *ExportRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportRequest) SetVhostName(v string) *ExportRequest {
  s.VhostName = &v
  return s
}

func (s *ExportRequest) Validate() error {
  return dara.Validate(s)
}

