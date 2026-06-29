// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnnotationsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOssPath(v string) *ExportAnnotationsRequest
  GetOssPath() *string 
  SetRegisterDataset(v string) *ExportAnnotationsRequest
  GetRegisterDataset() *string 
  SetTarget(v string) *ExportAnnotationsRequest
  GetTarget() *string 
}

type ExportAnnotationsRequest struct {
  // OSS path.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // oss://***-hz-oss.oss-cn-hangzhou.aliyuncs.com/output/
  OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
  // Specifies whether to register the result as a PAI dataset. Valid values:
  // 
  // - true: Registers the annotation result as a PAI dataset.
  // 
  // - false: Exports the annotation result directly to an OSS folder without registering it as a dataset.
  // 
  // example:
  // 
  // true
  RegisterDataset *string `json:"RegisterDataset,omitempty" xml:"RegisterDataset,omitempty"`
  // Target.
  // 
  // example:
  // 
  // PAI
  Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ExportAnnotationsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportAnnotationsRequest) GoString() string {
  return s.String()
}

func (s *ExportAnnotationsRequest) GetOssPath() *string  {
  return s.OssPath
}

func (s *ExportAnnotationsRequest) GetRegisterDataset() *string  {
  return s.RegisterDataset
}

func (s *ExportAnnotationsRequest) GetTarget() *string  {
  return s.Target
}

func (s *ExportAnnotationsRequest) SetOssPath(v string) *ExportAnnotationsRequest {
  s.OssPath = &v
  return s
}

func (s *ExportAnnotationsRequest) SetRegisterDataset(v string) *ExportAnnotationsRequest {
  s.RegisterDataset = &v
  return s
}

func (s *ExportAnnotationsRequest) SetTarget(v string) *ExportAnnotationsRequest {
  s.Target = &v
  return s
}

func (s *ExportAnnotationsRequest) Validate() error {
  return dara.Validate(s)
}

