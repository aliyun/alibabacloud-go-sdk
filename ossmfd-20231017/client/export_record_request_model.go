// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExportType(v string) *ExportRecordRequest
  GetExportType() *string 
  SetLang(v string) *ExportRecordRequest
  GetLang() *string 
  SetParams(v string) *ExportRecordRequest
  GetParams() *string 
}

type ExportRecordRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // exportObjectScanEvents
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // example:
  // 
  // {"currentPage":1,"pageSize":10}
  Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ExportRecordRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordRequest) GoString() string {
  return s.String()
}

func (s *ExportRecordRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportRecordRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExportRecordRequest) GetParams() *string  {
  return s.Params
}

func (s *ExportRecordRequest) SetExportType(v string) *ExportRecordRequest {
  s.ExportType = &v
  return s
}

func (s *ExportRecordRequest) SetLang(v string) *ExportRecordRequest {
  s.Lang = &v
  return s
}

func (s *ExportRecordRequest) SetParams(v string) *ExportRecordRequest {
  s.Params = &v
  return s
}

func (s *ExportRecordRequest) Validate() error {
  return dara.Validate(s)
}

