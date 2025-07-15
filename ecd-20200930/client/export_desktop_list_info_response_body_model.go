// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDesktopListInfoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExportDesktopListInfoResponseBody
  GetRequestId() *string 
  SetUrl(v string) *ExportDesktopListInfoResponseBody
  GetUrl() *string 
}

type ExportDesktopListInfoResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The URL of the exported file of the cloud computer list.
  // 
  // example:
  // 
  // https://cn-hangzhou-servicemanager.oss-cn-hangzhou.aliyuncs.com/A0_DESKTOP/EDS_CloudDesktops%20List_2022**********_w3fq******.csv?**********
  Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportDesktopListInfoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopListInfoResponseBody) GoString() string {
  return s.String()
}

func (s *ExportDesktopListInfoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportDesktopListInfoResponseBody) GetUrl() *string  {
  return s.Url
}

func (s *ExportDesktopListInfoResponseBody) SetRequestId(v string) *ExportDesktopListInfoResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportDesktopListInfoResponseBody) SetUrl(v string) *ExportDesktopListInfoResponseBody {
  s.Url = &v
  return s
}

func (s *ExportDesktopListInfoResponseBody) Validate() error {
  return dara.Validate(s)
}

