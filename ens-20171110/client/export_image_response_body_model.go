// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportImageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetExportedImageURL(v string) *ExportImageResponseBody
  GetExportedImageURL() *string 
  SetRequestId(v string) *ExportImageResponseBody
  GetRequestId() *string 
}

type ExportImageResponseBody struct {
  // The URL that points to the exported image.
  // 
  // example:
  // 
  // http://oss.url
  ExportedImageURL *string `json:"ExportedImageURL,omitempty" xml:"ExportedImageURL,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 43A426AD-3F2E-5DD9-9C08-D4DBDCA48D85
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportImageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportImageResponseBody) GoString() string {
  return s.String()
}

func (s *ExportImageResponseBody) GetExportedImageURL() *string  {
  return s.ExportedImageURL
}

func (s *ExportImageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportImageResponseBody) SetExportedImageURL(v string) *ExportImageResponseBody {
  s.ExportedImageURL = &v
  return s
}

func (s *ExportImageResponseBody) SetRequestId(v string) *ExportImageResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportImageResponseBody) Validate() error {
  return dara.Validate(s)
}

