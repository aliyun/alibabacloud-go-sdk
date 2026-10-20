// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSliderCaptchaPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *UpdateSliderCaptchaPageShrinkRequest
	GetConfigsShrink() *string
	SetCustomContent(v string) *UpdateSliderCaptchaPageShrinkRequest
	GetCustomContent() *string
	SetForceLanguage(v string) *UpdateSliderCaptchaPageShrinkRequest
	GetForceLanguage() *string
	SetId(v int64) *UpdateSliderCaptchaPageShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateSliderCaptchaPageShrinkRequest
	GetName() *string
	SetSiteIdsShrink(v string) *UpdateSliderCaptchaPageShrinkRequest
	GetSiteIdsShrink() *string
}

type UpdateSliderCaptchaPageShrinkRequest struct {
	// The predefined configurations of the custom slider page.
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The custom slider configuration.
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
	// The ID of the custom slider page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom slider page.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of bound site IDs.
	SiteIdsShrink *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
}

func (s UpdateSliderCaptchaPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *UpdateSliderCaptchaPageShrinkRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *UpdateSliderCaptchaPageShrinkRequest) GetForceLanguage() *string {
	return s.ForceLanguage
}

func (s *UpdateSliderCaptchaPageShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateSliderCaptchaPageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSliderCaptchaPageShrinkRequest) GetSiteIdsShrink() *string {
	return s.SiteIdsShrink
}

func (s *UpdateSliderCaptchaPageShrinkRequest) SetConfigsShrink(v string) *UpdateSliderCaptchaPageShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *UpdateSliderCaptchaPageShrinkRequest) SetCustomContent(v string) *UpdateSliderCaptchaPageShrinkRequest {
	s.CustomContent = &v
	return s
}

func (s *UpdateSliderCaptchaPageShrinkRequest) SetForceLanguage(v string) *UpdateSliderCaptchaPageShrinkRequest {
	s.ForceLanguage = &v
	return s
}

func (s *UpdateSliderCaptchaPageShrinkRequest) SetId(v int64) *UpdateSliderCaptchaPageShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateSliderCaptchaPageShrinkRequest) SetName(v string) *UpdateSliderCaptchaPageShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSliderCaptchaPageShrinkRequest) SetSiteIdsShrink(v string) *UpdateSliderCaptchaPageShrinkRequest {
	s.SiteIdsShrink = &v
	return s
}

func (s *UpdateSliderCaptchaPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
