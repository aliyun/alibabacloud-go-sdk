// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditSiteWafSettingsShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSettingsShrink(v string) *EditSiteWafSettingsShrinkRequest
  GetSettingsShrink() *string 
  SetSiteId(v int64) *EditSiteWafSettingsShrinkRequest
  GetSiteId() *int64 
  SetSiteVersion(v int32) *EditSiteWafSettingsShrinkRequest
  GetSiteVersion() *int32 
}

type EditSiteWafSettingsShrinkRequest struct {
  // WAF configuration information for the site, passed in JSON format.
  SettingsShrink *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
  // Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
  // 
  // example:
  // 
  // 1
  SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
  // Site version.
  // 
  // example:
  // 
  // 0
  SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s EditSiteWafSettingsShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EditSiteWafSettingsShrinkRequest) GoString() string {
  return s.String()
}

func (s *EditSiteWafSettingsShrinkRequest) GetSettingsShrink() *string  {
  return s.SettingsShrink
}

func (s *EditSiteWafSettingsShrinkRequest) GetSiteId() *int64  {
  return s.SiteId
}

func (s *EditSiteWafSettingsShrinkRequest) GetSiteVersion() *int32  {
  return s.SiteVersion
}

func (s *EditSiteWafSettingsShrinkRequest) SetSettingsShrink(v string) *EditSiteWafSettingsShrinkRequest {
  s.SettingsShrink = &v
  return s
}

func (s *EditSiteWafSettingsShrinkRequest) SetSiteId(v int64) *EditSiteWafSettingsShrinkRequest {
  s.SiteId = &v
  return s
}

func (s *EditSiteWafSettingsShrinkRequest) SetSiteVersion(v int32) *EditSiteWafSettingsShrinkRequest {
  s.SiteVersion = &v
  return s
}

func (s *EditSiteWafSettingsShrinkRequest) Validate() error {
  return dara.Validate(s)
}

