// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportKgSchemaRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOpTenantId(v int64) *ExportKgSchemaRequest
  GetOpTenantId() *int64 
  SetOutputFormat(v string) *ExportKgSchemaRequest
  GetOutputFormat() *string 
  SetVersionId(v int32) *ExportKgSchemaRequest
  GetVersionId() *int32 
  SetWorkspaceId(v string) *ExportKgSchemaRequest
  GetWorkspaceId() *string 
}

type ExportKgSchemaRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
  // example:
  // 
  // json
  OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
  // example:
  // 
  // 0
  VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // f1d4559a4db044158305e2d89bccf81f
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportKgSchemaRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportKgSchemaRequest) GoString() string {
  return s.String()
}

func (s *ExportKgSchemaRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExportKgSchemaRequest) GetOutputFormat() *string  {
  return s.OutputFormat
}

func (s *ExportKgSchemaRequest) GetVersionId() *int32  {
  return s.VersionId
}

func (s *ExportKgSchemaRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportKgSchemaRequest) SetOpTenantId(v int64) *ExportKgSchemaRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExportKgSchemaRequest) SetOutputFormat(v string) *ExportKgSchemaRequest {
  s.OutputFormat = &v
  return s
}

func (s *ExportKgSchemaRequest) SetVersionId(v int32) *ExportKgSchemaRequest {
  s.VersionId = &v
  return s
}

func (s *ExportKgSchemaRequest) SetWorkspaceId(v string) *ExportKgSchemaRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportKgSchemaRequest) Validate() error {
  return dara.Validate(s)
}

