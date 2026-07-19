// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRbacConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizId(v string) *ExportRbacConfigRequest
  GetBizId() *string 
}

type ExportRbacConfigRequest struct {
  // example:
  // 
  // WS20250814102215000001
  BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s ExportRbacConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRbacConfigRequest) GoString() string {
  return s.String()
}

func (s *ExportRbacConfigRequest) GetBizId() *string  {
  return s.BizId
}

func (s *ExportRbacConfigRequest) SetBizId(v string) *ExportRbacConfigRequest {
  s.BizId = &v
  return s
}

func (s *ExportRbacConfigRequest) Validate() error {
  return dara.Validate(s)
}

