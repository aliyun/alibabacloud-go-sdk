// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMaterialFileShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizId(v string) *ExportMaterialFileShrinkRequest
  GetBizId() *string 
  SetFileIdsShrink(v string) *ExportMaterialFileShrinkRequest
  GetFileIdsShrink() *string 
}

type ExportMaterialFileShrinkRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // WS20250801154628000001
  BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
  // This parameter is required.
  FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
}

func (s ExportMaterialFileShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportMaterialFileShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportMaterialFileShrinkRequest) GetBizId() *string  {
  return s.BizId
}

func (s *ExportMaterialFileShrinkRequest) GetFileIdsShrink() *string  {
  return s.FileIdsShrink
}

func (s *ExportMaterialFileShrinkRequest) SetBizId(v string) *ExportMaterialFileShrinkRequest {
  s.BizId = &v
  return s
}

func (s *ExportMaterialFileShrinkRequest) SetFileIdsShrink(v string) *ExportMaterialFileShrinkRequest {
  s.FileIdsShrink = &v
  return s
}

func (s *ExportMaterialFileShrinkRequest) Validate() error {
  return dara.Validate(s)
}

