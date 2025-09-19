// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVulResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFileName(v string) *ExportVulResponseBody
  GetFileName() *string 
  SetId(v int64) *ExportVulResponseBody
  GetId() *int64 
  SetRequestId(v string) *ExportVulResponseBody
  GetRequestId() *string 
}

type ExportVulResponseBody struct {
  // The name of the exported file.
  // 
  // example:
  // 
  // app_20211101
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // The ID of the exported file.
  // 
  // example:
  // 
  // 81634
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // E1FAB2B8-DF4D-55DF-BC3D-5C3CA6FD5B13
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportVulResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportVulResponseBody) GoString() string {
  return s.String()
}

func (s *ExportVulResponseBody) GetFileName() *string  {
  return s.FileName
}

func (s *ExportVulResponseBody) GetId() *int64  {
  return s.Id
}

func (s *ExportVulResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportVulResponseBody) SetFileName(v string) *ExportVulResponseBody {
  s.FileName = &v
  return s
}

func (s *ExportVulResponseBody) SetId(v int64) *ExportVulResponseBody {
  s.Id = &v
  return s
}

func (s *ExportVulResponseBody) SetRequestId(v string) *ExportVulResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportVulResponseBody) Validate() error {
  return dara.Validate(s)
}

