// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteWafSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSiteWafSettingsResponseBody
	GetRequestId() *string
	SetSettings(v *WafSiteSettings) *GetSiteWafSettingsResponseBody
	GetSettings() *WafSiteSettings
}

type GetSiteWafSettingsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of site WAF configuration.
	Settings *WafSiteSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s GetSiteWafSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteWafSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteWafSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteWafSettingsResponseBody) GetSettings() *WafSiteSettings {
	return s.Settings
}

func (s *GetSiteWafSettingsResponseBody) SetRequestId(v string) *GetSiteWafSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteWafSettingsResponseBody) SetSettings(v *WafSiteSettings) *GetSiteWafSettingsResponseBody {
	s.Settings = v
	return s
}

func (s *GetSiteWafSettingsResponseBody) Validate() error {
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			return err
		}
	}
	return nil
}
