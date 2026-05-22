// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSiteId(v int64) *ExportRecordsRequest
  GetSiteId() *int64 
}

type ExportRecordsRequest struct {
  // The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
  // 
  // example:
  // 
  // 1234567890123
  SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ExportRecordsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordsRequest) GoString() string {
  return s.String()
}

func (s *ExportRecordsRequest) GetSiteId() *int64  {
  return s.SiteId
}

func (s *ExportRecordsRequest) SetSiteId(v int64) *ExportRecordsRequest {
  s.SiteId = &v
  return s
}

func (s *ExportRecordsRequest) Validate() error {
  return dara.Validate(s)
}

