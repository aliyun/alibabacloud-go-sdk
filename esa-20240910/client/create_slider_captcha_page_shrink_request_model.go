// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSliderCaptchaPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *CreateSliderCaptchaPageShrinkRequest
	GetConfigsShrink() *string
	SetCustomContent(v string) *CreateSliderCaptchaPageShrinkRequest
	GetCustomContent() *string
	SetForceLanguage(v string) *CreateSliderCaptchaPageShrinkRequest
	GetForceLanguage() *string
	SetName(v string) *CreateSliderCaptchaPageShrinkRequest
	GetName() *string
	SetSiteIdsShrink(v string) *CreateSliderCaptchaPageShrinkRequest
	GetSiteIdsShrink() *string
}

type CreateSliderCaptchaPageShrinkRequest struct {
	// The predefined configurations for the custom slider page.
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The custom slider configurations. The value is a JSON string.
	//
	// example:
	//
	// [{"language":"ko","description":"더 나은 경험을 위해 웹 페이지에 접속하기 전에 슬라이드하여 인증을 완료해 주세요.","captcha_color":"#28870d","show_trace_id":true,"show_footer":true,"mandatory_lang":false}]
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The forced language.
	//
	// example:
	//
	// zh
	ForceLanguage *string `json:"ForceLanguage,omitempty" xml:"ForceLanguage,omitempty"`
	// The name of the custom slider page.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of bound site IDs.
	SiteIdsShrink *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
}

func (s CreateSliderCaptchaPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *CreateSliderCaptchaPageShrinkRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *CreateSliderCaptchaPageShrinkRequest) GetForceLanguage() *string {
	return s.ForceLanguage
}

func (s *CreateSliderCaptchaPageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSliderCaptchaPageShrinkRequest) GetSiteIdsShrink() *string {
	return s.SiteIdsShrink
}

func (s *CreateSliderCaptchaPageShrinkRequest) SetConfigsShrink(v string) *CreateSliderCaptchaPageShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *CreateSliderCaptchaPageShrinkRequest) SetCustomContent(v string) *CreateSliderCaptchaPageShrinkRequest {
	s.CustomContent = &v
	return s
}

func (s *CreateSliderCaptchaPageShrinkRequest) SetForceLanguage(v string) *CreateSliderCaptchaPageShrinkRequest {
	s.ForceLanguage = &v
	return s
}

func (s *CreateSliderCaptchaPageShrinkRequest) SetName(v string) *CreateSliderCaptchaPageShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSliderCaptchaPageShrinkRequest) SetSiteIdsShrink(v string) *CreateSliderCaptchaPageShrinkRequest {
	s.SiteIdsShrink = &v
	return s
}

func (s *CreateSliderCaptchaPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
