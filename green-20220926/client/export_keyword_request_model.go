// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportKeywordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLibId(v string) *ExportKeywordRequest
  GetLibId() *string 
  SetRegionId(v string) *ExportKeywordRequest
  GetRegionId() *string 
}

type ExportKeywordRequest struct {
  // Keyword library ID.
  // 
  // example:
  // 
  // customxx_xxxx
  LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
  // Region ID.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportKeywordRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportKeywordRequest) GoString() string {
  return s.String()
}

func (s *ExportKeywordRequest) GetLibId() *string  {
  return s.LibId
}

func (s *ExportKeywordRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportKeywordRequest) SetLibId(v string) *ExportKeywordRequest {
  s.LibId = &v
  return s
}

func (s *ExportKeywordRequest) SetRegionId(v string) *ExportKeywordRequest {
  s.RegionId = &v
  return s
}

func (s *ExportKeywordRequest) Validate() error {
  return dara.Validate(s)
}

