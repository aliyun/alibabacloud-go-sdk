// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordCnamesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSiteId(v int64) *ExportRecordCnamesRequest
  GetSiteId() *int64 
}

type ExportRecordCnamesRequest struct {
  // The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234567890
  SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ExportRecordCnamesRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordCnamesRequest) GoString() string {
  return s.String()
}

func (s *ExportRecordCnamesRequest) GetSiteId() *int64  {
  return s.SiteId
}

func (s *ExportRecordCnamesRequest) SetSiteId(v int64) *ExportRecordCnamesRequest {
  s.SiteId = &v
  return s
}

func (s *ExportRecordCnamesRequest) Validate() error {
  return dara.Validate(s)
}

