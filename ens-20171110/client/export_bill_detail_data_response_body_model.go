// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportBillDetailDataResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFilePath(v string) *ExportBillDetailDataResponseBody
  GetFilePath() *string 
  SetRequestId(v string) *ExportBillDetailDataResponseBody
  GetRequestId() *string 
}

type ExportBillDetailDataResponseBody struct {
  // The download path of the exported file.
  // 
  // example:
  // 
  // http://test-oss.com/image_01.jpeg
  FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
  // The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
  // 
  // example:
  // 
  // BEA05990-B90D-594F-8C8E-650AEEB94C5D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportBillDetailDataResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportBillDetailDataResponseBody) GoString() string {
  return s.String()
}

func (s *ExportBillDetailDataResponseBody) GetFilePath() *string  {
  return s.FilePath
}

func (s *ExportBillDetailDataResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportBillDetailDataResponseBody) SetFilePath(v string) *ExportBillDetailDataResponseBody {
  s.FilePath = &v
  return s
}

func (s *ExportBillDetailDataResponseBody) SetRequestId(v string) *ExportBillDetailDataResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportBillDetailDataResponseBody) Validate() error {
  return dara.Validate(s)
}

