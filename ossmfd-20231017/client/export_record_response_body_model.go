// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFileName(v string) *ExportRecordResponseBody
  GetFileName() *string 
  SetId(v int64) *ExportRecordResponseBody
  GetId() *int64 
  SetRequestId(v string) *ExportRecordResponseBody
  GetRequestId() *string 
}

type ExportRecordResponseBody struct {
  // example:
  // 
  // health_check_export_20******
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // example:
  // 
  // 100******
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // example:
  // 
  // AAC546A5-5EDC-5939-86A3-56DFAF******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportRecordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordResponseBody) GoString() string {
  return s.String()
}

func (s *ExportRecordResponseBody) GetFileName() *string  {
  return s.FileName
}

func (s *ExportRecordResponseBody) GetId() *int64  {
  return s.Id
}

func (s *ExportRecordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportRecordResponseBody) SetFileName(v string) *ExportRecordResponseBody {
  s.FileName = &v
  return s
}

func (s *ExportRecordResponseBody) SetId(v int64) *ExportRecordResponseBody {
  s.Id = &v
  return s
}

func (s *ExportRecordResponseBody) SetRequestId(v string) *ExportRecordResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportRecordResponseBody) Validate() error {
  return dara.Validate(s)
}

