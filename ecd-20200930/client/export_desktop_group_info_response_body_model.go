// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDesktopGroupInfoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExportDesktopGroupInfoResponseBody
  GetRequestId() *string 
  SetUrl(v string) *ExportDesktopGroupInfoResponseBody
  GetUrl() *string 
}

type ExportDesktopGroupInfoResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 24E05D3E-08F4-551E-B1F0-F6D84EE0BCCC
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The download URL of the exported shared cloud desktop list file. The exported file contains the following details of shared cloud desktops:
  // 
  // - Shared cloud desktop ID/name
  // 
  // - Office network ID/name
  // 
  // - Shared cloud desktop template
  // 
  // - CPU/memory
  // 
  // - System cloud disk/data cloud disk
  // 
  // - Security policy name
  // 
  // - Number of currently authorized users
  // 
  // - Billing method
  // 
  // - Creation time
  // 
  // - Expiration time
  // 
  // example:
  // 
  // https://cn-hangzhou-servicemanager.oss-cn-hangzhou.aliyuncs.com/A0_DESKTOP/EDS_CloudDesktopGroups_202203********_xBjqdCT***.xlsx?*********
  Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportDesktopGroupInfoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopGroupInfoResponseBody) GoString() string {
  return s.String()
}

func (s *ExportDesktopGroupInfoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportDesktopGroupInfoResponseBody) GetUrl() *string  {
  return s.Url
}

func (s *ExportDesktopGroupInfoResponseBody) SetRequestId(v string) *ExportDesktopGroupInfoResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportDesktopGroupInfoResponseBody) SetUrl(v string) *ExportDesktopGroupInfoResponseBody {
  s.Url = &v
  return s
}

func (s *ExportDesktopGroupInfoResponseBody) Validate() error {
  return dara.Validate(s)
}

