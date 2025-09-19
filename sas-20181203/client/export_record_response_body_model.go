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
  // The name of the exported file.
  // 
  // example:
  // 
  // cms_20171101.xlsx
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // The ID of the exported file.
  // 
  // example:
  // 
  // 131231
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // 6673D49C-A9AB-40DD-B4A2-B92306701AE7
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

