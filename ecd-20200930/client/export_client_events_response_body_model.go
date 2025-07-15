// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportClientEventsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExportClientEventsResponseBody
  GetRequestId() *string 
  SetUrl(v string) *ExportClientEventsResponseBody
  GetUrl() *string 
}

type ExportClientEventsResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The download URL of the exported files.
  // 
  // example:
  // 
  // https://cn-shanghai-servicemanager.oss-cn-shanghai.aliyuncs.com/A0_CLIENT_EVENT/EDS_Events%20List_20220519234611_w5HuD83KGs.csv?Expires=1652975773&OSSAccessKeyId=****&Signature=4erMG*********k%3D
  Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportClientEventsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportClientEventsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportClientEventsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportClientEventsResponseBody) GetUrl() *string  {
  return s.Url
}

func (s *ExportClientEventsResponseBody) SetRequestId(v string) *ExportClientEventsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportClientEventsResponseBody) SetUrl(v string) *ExportClientEventsResponseBody {
  s.Url = &v
  return s
}

func (s *ExportClientEventsResponseBody) Validate() error {
  return dara.Validate(s)
}

