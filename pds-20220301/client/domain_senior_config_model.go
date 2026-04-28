// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainSeniorConfig interface {
	dara.Model
	String() string
	GoString() string
	SetClientDownloadEnable(v bool) *DomainSeniorConfig
	GetClientDownloadEnable() *bool
	SetCspFrameAncestors(v string) *DomainSeniorConfig
	GetCspFrameAncestors() *string
	SetCustomLoginAppid(v string) *DomainSeniorConfig
	GetCustomLoginAppid() *string
	SetCustomLoginUrl(v string) *DomainSeniorConfig
	GetCustomLoginUrl() *string
	SetCustomLogoutUrl(v string) *DomainSeniorConfig
	GetCustomLogoutUrl() *string
	SetCustomSideLinkList(v []*CustomSideLinkConfig) *DomainSeniorConfig
	GetCustomSideLinkList() []*CustomSideLinkConfig
	SetHomePageBgImageUrl(v string) *DomainSeniorConfig
	GetHomePageBgImageUrl() *string
	SetHomePageFooter(v string) *DomainSeniorConfig
	GetHomePageFooter() *string
	SetHomePageFooter2(v string) *DomainSeniorConfig
	GetHomePageFooter2() *string
	SetHomePageSlogan(v string) *DomainSeniorConfig
	GetHomePageSlogan() *string
	SetRefererEnable(v bool) *DomainSeniorConfig
	GetRefererEnable() *bool
	SetWxTxtList(v *WxTrustedDomainConfig) *DomainSeniorConfig
	GetWxTxtList() *WxTrustedDomainConfig
}

type DomainSeniorConfig struct {
	ClientDownloadEnable *bool                   `json:"client_download_enable,omitempty" xml:"client_download_enable,omitempty"`
	CspFrameAncestors    *string                 `json:"csp_frame_ancestors,omitempty" xml:"csp_frame_ancestors,omitempty"`
	CustomLoginAppid     *string                 `json:"custom_login_appid,omitempty" xml:"custom_login_appid,omitempty"`
	CustomLoginUrl       *string                 `json:"custom_login_url,omitempty" xml:"custom_login_url,omitempty"`
	CustomLogoutUrl      *string                 `json:"custom_logout_url,omitempty" xml:"custom_logout_url,omitempty"`
	CustomSideLinkList   []*CustomSideLinkConfig `json:"custom_side_link_list,omitempty" xml:"custom_side_link_list,omitempty" type:"Repeated"`
	HomePageBgImageUrl   *string                 `json:"home_page_bg_image_url,omitempty" xml:"home_page_bg_image_url,omitempty"`
	HomePageFooter       *string                 `json:"home_page_footer,omitempty" xml:"home_page_footer,omitempty"`
	HomePageFooter2      *string                 `json:"home_page_footer2,omitempty" xml:"home_page_footer2,omitempty"`
	HomePageSlogan       *string                 `json:"home_page_slogan,omitempty" xml:"home_page_slogan,omitempty"`
	RefererEnable        *bool                   `json:"referer_enable,omitempty" xml:"referer_enable,omitempty"`
	WxTxtList            *WxTrustedDomainConfig  `json:"wx_txt_list,omitempty" xml:"wx_txt_list,omitempty"`
}

func (s DomainSeniorConfig) String() string {
	return dara.Prettify(s)
}

func (s DomainSeniorConfig) GoString() string {
	return s.String()
}

func (s *DomainSeniorConfig) GetClientDownloadEnable() *bool {
	return s.ClientDownloadEnable
}

func (s *DomainSeniorConfig) GetCspFrameAncestors() *string {
	return s.CspFrameAncestors
}

func (s *DomainSeniorConfig) GetCustomLoginAppid() *string {
	return s.CustomLoginAppid
}

func (s *DomainSeniorConfig) GetCustomLoginUrl() *string {
	return s.CustomLoginUrl
}

func (s *DomainSeniorConfig) GetCustomLogoutUrl() *string {
	return s.CustomLogoutUrl
}

func (s *DomainSeniorConfig) GetCustomSideLinkList() []*CustomSideLinkConfig {
	return s.CustomSideLinkList
}

func (s *DomainSeniorConfig) GetHomePageBgImageUrl() *string {
	return s.HomePageBgImageUrl
}

func (s *DomainSeniorConfig) GetHomePageFooter() *string {
	return s.HomePageFooter
}

func (s *DomainSeniorConfig) GetHomePageFooter2() *string {
	return s.HomePageFooter2
}

func (s *DomainSeniorConfig) GetHomePageSlogan() *string {
	return s.HomePageSlogan
}

func (s *DomainSeniorConfig) GetRefererEnable() *bool {
	return s.RefererEnable
}

func (s *DomainSeniorConfig) GetWxTxtList() *WxTrustedDomainConfig {
	return s.WxTxtList
}

func (s *DomainSeniorConfig) SetClientDownloadEnable(v bool) *DomainSeniorConfig {
	s.ClientDownloadEnable = &v
	return s
}

func (s *DomainSeniorConfig) SetCspFrameAncestors(v string) *DomainSeniorConfig {
	s.CspFrameAncestors = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomLoginAppid(v string) *DomainSeniorConfig {
	s.CustomLoginAppid = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomLoginUrl(v string) *DomainSeniorConfig {
	s.CustomLoginUrl = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomLogoutUrl(v string) *DomainSeniorConfig {
	s.CustomLogoutUrl = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomSideLinkList(v []*CustomSideLinkConfig) *DomainSeniorConfig {
	s.CustomSideLinkList = v
	return s
}

func (s *DomainSeniorConfig) SetHomePageBgImageUrl(v string) *DomainSeniorConfig {
	s.HomePageBgImageUrl = &v
	return s
}

func (s *DomainSeniorConfig) SetHomePageFooter(v string) *DomainSeniorConfig {
	s.HomePageFooter = &v
	return s
}

func (s *DomainSeniorConfig) SetHomePageFooter2(v string) *DomainSeniorConfig {
	s.HomePageFooter2 = &v
	return s
}

func (s *DomainSeniorConfig) SetHomePageSlogan(v string) *DomainSeniorConfig {
	s.HomePageSlogan = &v
	return s
}

func (s *DomainSeniorConfig) SetRefererEnable(v bool) *DomainSeniorConfig {
	s.RefererEnable = &v
	return s
}

func (s *DomainSeniorConfig) SetWxTxtList(v *WxTrustedDomainConfig) *DomainSeniorConfig {
	s.WxTxtList = v
	return s
}

func (s *DomainSeniorConfig) Validate() error {
	if s.CustomSideLinkList != nil {
		for _, item := range s.CustomSideLinkList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WxTxtList != nil {
		if err := s.WxTxtList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
