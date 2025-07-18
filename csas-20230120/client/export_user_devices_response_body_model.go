// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportUserDevicesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExportUserDevicesResponseBody
  GetRequestId() *string 
  SetSignedUrl(v string) *ExportUserDevicesResponseBody
  GetSignedUrl() *string 
}

type ExportUserDevicesResponseBody struct {
  // example:
  // 
  // 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // https://sase-export.oss-cn-hangzhou.aliyuncs.com/export%2Fapp-device%2F20240607154831.xlsx?Expires=1717746571&OSSAccessKeyId=********************
  SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s ExportUserDevicesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportUserDevicesResponseBody) GoString() string {
  return s.String()
}

func (s *ExportUserDevicesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportUserDevicesResponseBody) GetSignedUrl() *string  {
  return s.SignedUrl
}

func (s *ExportUserDevicesResponseBody) SetRequestId(v string) *ExportUserDevicesResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportUserDevicesResponseBody) SetSignedUrl(v string) *ExportUserDevicesResponseBody {
  s.SignedUrl = &v
  return s
}

func (s *ExportUserDevicesResponseBody) Validate() error {
  return dara.Validate(s)
}

