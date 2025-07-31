// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnswerSampleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLibId(v string) *ExportAnswerSampleRequest
  GetLibId() *string 
  SetRegionId(v string) *ExportAnswerSampleRequest
  GetRegionId() *string 
}

type ExportAnswerSampleRequest struct {
  // example:
  // 
  // alxxx
  LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportAnswerSampleRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportAnswerSampleRequest) GoString() string {
  return s.String()
}

func (s *ExportAnswerSampleRequest) GetLibId() *string  {
  return s.LibId
}

func (s *ExportAnswerSampleRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportAnswerSampleRequest) SetLibId(v string) *ExportAnswerSampleRequest {
  s.LibId = &v
  return s
}

func (s *ExportAnswerSampleRequest) SetRegionId(v string) *ExportAnswerSampleRequest {
  s.RegionId = &v
  return s
}

func (s *ExportAnswerSampleRequest) Validate() error {
  return dara.Validate(s)
}

