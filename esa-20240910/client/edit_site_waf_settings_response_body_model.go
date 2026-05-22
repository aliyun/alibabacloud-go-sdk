// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditSiteWafSettingsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EditSiteWafSettingsResponseBody
  GetRequestId() *string 
}

type EditSiteWafSettingsResponseBody struct {
  // Request ID.
  // 
  // example:
  // 
  // 36af3fcc-43d0-441c-86b1-428951dc8225
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditSiteWafSettingsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditSiteWafSettingsResponseBody) GoString() string {
  return s.String()
}

func (s *EditSiteWafSettingsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditSiteWafSettingsResponseBody) SetRequestId(v string) *EditSiteWafSettingsResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditSiteWafSettingsResponseBody) Validate() error {
  return dara.Validate(s)
}

