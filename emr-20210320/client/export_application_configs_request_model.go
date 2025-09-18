// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportApplicationConfigsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationConfigFiles(v []*ApplicationConfigFile) *ExportApplicationConfigsRequest
  GetApplicationConfigFiles() []*ApplicationConfigFile 
  SetClusterId(v string) *ExportApplicationConfigsRequest
  GetClusterId() *string 
  SetExportMode(v string) *ExportApplicationConfigsRequest
  GetExportMode() *string 
  SetFileFormat(v string) *ExportApplicationConfigsRequest
  GetFileFormat() *string 
  SetRegionId(v string) *ExportApplicationConfigsRequest
  GetRegionId() *string 
}

type ExportApplicationConfigsRequest struct {
  // 导出应用配置。
  ApplicationConfigFiles []*ApplicationConfigFile `json:"ApplicationConfigFiles,omitempty" xml:"ApplicationConfigFiles,omitempty" type:"Repeated"`
  // 集群ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // c-b933c5aac8fe****
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // example:
  // 
  // MODIFICATION
  ExportMode *string `json:"ExportMode,omitempty" xml:"ExportMode,omitempty"`
  // 导出应用配置的文件格式。
  // 
  // example:
  // 
  // XML
  FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
  // 区域ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportApplicationConfigsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportApplicationConfigsRequest) GoString() string {
  return s.String()
}

func (s *ExportApplicationConfigsRequest) GetApplicationConfigFiles() []*ApplicationConfigFile  {
  return s.ApplicationConfigFiles
}

func (s *ExportApplicationConfigsRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportApplicationConfigsRequest) GetExportMode() *string  {
  return s.ExportMode
}

func (s *ExportApplicationConfigsRequest) GetFileFormat() *string  {
  return s.FileFormat
}

func (s *ExportApplicationConfigsRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportApplicationConfigsRequest) SetApplicationConfigFiles(v []*ApplicationConfigFile) *ExportApplicationConfigsRequest {
  s.ApplicationConfigFiles = v
  return s
}

func (s *ExportApplicationConfigsRequest) SetClusterId(v string) *ExportApplicationConfigsRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportApplicationConfigsRequest) SetExportMode(v string) *ExportApplicationConfigsRequest {
  s.ExportMode = &v
  return s
}

func (s *ExportApplicationConfigsRequest) SetFileFormat(v string) *ExportApplicationConfigsRequest {
  s.FileFormat = &v
  return s
}

func (s *ExportApplicationConfigsRequest) SetRegionId(v string) *ExportApplicationConfigsRequest {
  s.RegionId = &v
  return s
}

func (s *ExportApplicationConfigsRequest) Validate() error {
  return dara.Validate(s)
}

