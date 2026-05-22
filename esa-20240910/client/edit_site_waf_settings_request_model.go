// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditSiteWafSettingsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSettings(v *WafSiteSettings) *EditSiteWafSettingsRequest
  GetSettings() *WafSiteSettings 
  SetSiteId(v int64) *EditSiteWafSettingsRequest
  GetSiteId() *int64 
  SetSiteVersion(v int32) *EditSiteWafSettingsRequest
  GetSiteVersion() *int32 
}

type EditSiteWafSettingsRequest struct {
  Settings *WafSiteSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
  // example:
  // 
  // 1
  SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
  SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s EditSiteWafSettingsRequest) String() string {
  return dara.Prettify(s)
}

func (s EditSiteWafSettingsRequest) GoString() string {
  return s.String()
}

func (s *EditSiteWafSettingsRequest) GetSettings() *WafSiteSettings  {
  return s.Settings
}

func (s *EditSiteWafSettingsRequest) GetSiteId() *int64  {
  return s.SiteId
}

func (s *EditSiteWafSettingsRequest) GetSiteVersion() *int32  {
  return s.SiteVersion
}

func (s *EditSiteWafSettingsRequest) SetSettings(v *WafSiteSettings) *EditSiteWafSettingsRequest {
  s.Settings = v
  return s
}

func (s *EditSiteWafSettingsRequest) SetSiteId(v int64) *EditSiteWafSettingsRequest {
  s.SiteId = &v
  return s
}

func (s *EditSiteWafSettingsRequest) SetSiteVersion(v int32) *EditSiteWafSettingsRequest {
  s.SiteVersion = &v
  return s
}

func (s *EditSiteWafSettingsRequest) Validate() error {
  if s.Settings != nil {
    if err := s.Settings.Validate(); err != nil {
      return err
    }
  }
  return nil
}

