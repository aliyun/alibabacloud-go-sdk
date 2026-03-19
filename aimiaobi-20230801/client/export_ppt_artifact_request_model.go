// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportPptArtifactRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEdit(v bool) *ExportPptArtifactRequest
  GetEdit() *bool 
  SetExportFileType(v string) *ExportPptArtifactRequest
  GetExportFileType() *string 
  SetPptArtifactId(v int64) *ExportPptArtifactRequest
  GetPptArtifactId() *int64 
  SetWorkspaceId(v string) *ExportPptArtifactRequest
  GetWorkspaceId() *string 
  SetZip(v bool) *ExportPptArtifactRequest
  GetZip() *bool 
}

type ExportPptArtifactRequest struct {
  // example:
  // 
  // true
  Edit *bool `json:"Edit,omitempty" xml:"Edit,omitempty"`
  // example:
  // 
  // ppt
  ExportFileType *string `json:"ExportFileType,omitempty" xml:"ExportFileType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 123123
  PptArtifactId *int64 `json:"PptArtifactId,omitempty" xml:"PptArtifactId,omitempty"`
  // example:
  // 
  // llm-xxxx
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
  // example:
  // 
  // true
  Zip *bool `json:"Zip,omitempty" xml:"Zip,omitempty"`
}

func (s ExportPptArtifactRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportPptArtifactRequest) GoString() string {
  return s.String()
}

func (s *ExportPptArtifactRequest) GetEdit() *bool  {
  return s.Edit
}

func (s *ExportPptArtifactRequest) GetExportFileType() *string  {
  return s.ExportFileType
}

func (s *ExportPptArtifactRequest) GetPptArtifactId() *int64  {
  return s.PptArtifactId
}

func (s *ExportPptArtifactRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportPptArtifactRequest) GetZip() *bool  {
  return s.Zip
}

func (s *ExportPptArtifactRequest) SetEdit(v bool) *ExportPptArtifactRequest {
  s.Edit = &v
  return s
}

func (s *ExportPptArtifactRequest) SetExportFileType(v string) *ExportPptArtifactRequest {
  s.ExportFileType = &v
  return s
}

func (s *ExportPptArtifactRequest) SetPptArtifactId(v int64) *ExportPptArtifactRequest {
  s.PptArtifactId = &v
  return s
}

func (s *ExportPptArtifactRequest) SetWorkspaceId(v string) *ExportPptArtifactRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportPptArtifactRequest) SetZip(v bool) *ExportPptArtifactRequest {
  s.Zip = &v
  return s
}

func (s *ExportPptArtifactRequest) Validate() error {
  return dara.Validate(s)
}

