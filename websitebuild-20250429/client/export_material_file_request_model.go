// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMaterialFileRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizId(v string) *ExportMaterialFileRequest
  GetBizId() *string 
  SetFileIds(v []*string) *ExportMaterialFileRequest
  GetFileIds() []*string 
}

type ExportMaterialFileRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // WS20250801154628000001
  BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
  // This parameter is required.
  FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
}

func (s ExportMaterialFileRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportMaterialFileRequest) GoString() string {
  return s.String()
}

func (s *ExportMaterialFileRequest) GetBizId() *string  {
  return s.BizId
}

func (s *ExportMaterialFileRequest) GetFileIds() []*string  {
  return s.FileIds
}

func (s *ExportMaterialFileRequest) SetBizId(v string) *ExportMaterialFileRequest {
  s.BizId = &v
  return s
}

func (s *ExportMaterialFileRequest) SetFileIds(v []*string) *ExportMaterialFileRequest {
  s.FileIds = v
  return s
}

func (s *ExportMaterialFileRequest) Validate() error {
  return dara.Validate(s)
}

