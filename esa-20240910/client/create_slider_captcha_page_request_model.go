// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSliderCaptchaPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *CreateSliderCaptchaPageRequestConfigs) *CreateSliderCaptchaPageRequest
	GetConfigs() *CreateSliderCaptchaPageRequestConfigs
	SetCustomContent(v string) *CreateSliderCaptchaPageRequest
	GetCustomContent() *string
	SetForceLanguage(v string) *CreateSliderCaptchaPageRequest
	GetForceLanguage() *string
	SetName(v string) *CreateSliderCaptchaPageRequest
	GetName() *string
	SetSiteIds(v []*int64) *CreateSliderCaptchaPageRequest
	GetSiteIds() []*int64
}

type CreateSliderCaptchaPageRequest struct {
	// The predefined configurations for the custom slider page.
	Configs *CreateSliderCaptchaPageRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
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
	SiteIds []*int64 `json:"SiteIds,omitempty" xml:"SiteIds,omitempty" type:"Repeated"`
}

func (s CreateSliderCaptchaPageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageRequest) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageRequest) GetConfigs() *CreateSliderCaptchaPageRequestConfigs {
	return s.Configs
}

func (s *CreateSliderCaptchaPageRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *CreateSliderCaptchaPageRequest) GetForceLanguage() *string {
	return s.ForceLanguage
}

func (s *CreateSliderCaptchaPageRequest) GetName() *string {
	return s.Name
}

func (s *CreateSliderCaptchaPageRequest) GetSiteIds() []*int64 {
	return s.SiteIds
}

func (s *CreateSliderCaptchaPageRequest) SetConfigs(v *CreateSliderCaptchaPageRequestConfigs) *CreateSliderCaptchaPageRequest {
	s.Configs = v
	return s
}

func (s *CreateSliderCaptchaPageRequest) SetCustomContent(v string) *CreateSliderCaptchaPageRequest {
	s.CustomContent = &v
	return s
}

func (s *CreateSliderCaptchaPageRequest) SetForceLanguage(v string) *CreateSliderCaptchaPageRequest {
	s.ForceLanguage = &v
	return s
}

func (s *CreateSliderCaptchaPageRequest) SetName(v string) *CreateSliderCaptchaPageRequest {
	s.Name = &v
	return s
}

func (s *CreateSliderCaptchaPageRequest) SetSiteIds(v []*int64) *CreateSliderCaptchaPageRequest {
	s.SiteIds = v
	return s
}

func (s *CreateSliderCaptchaPageRequest) Validate() error {
	if s.Configs != nil {
		if err := s.Configs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSliderCaptchaPageRequestConfigs struct {
	// The predefined English style.
	En *CreateSliderCaptchaPageRequestConfigsEn `json:"En,omitempty" xml:"En,omitempty" type:"Struct"`
	// The predefined Chinese style.
	Zh *CreateSliderCaptchaPageRequestConfigsZh `json:"Zh,omitempty" xml:"Zh,omitempty" type:"Struct"`
}

func (s CreateSliderCaptchaPageRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageRequestConfigs) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageRequestConfigs) GetEn() *CreateSliderCaptchaPageRequestConfigsEn {
	return s.En
}

func (s *CreateSliderCaptchaPageRequestConfigs) GetZh() *CreateSliderCaptchaPageRequestConfigsZh {
	return s.Zh
}

func (s *CreateSliderCaptchaPageRequestConfigs) SetEn(v *CreateSliderCaptchaPageRequestConfigsEn) *CreateSliderCaptchaPageRequestConfigs {
	s.En = v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigs) SetZh(v *CreateSliderCaptchaPageRequestConfigsZh) *CreateSliderCaptchaPageRequestConfigs {
	s.Zh = v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigs) Validate() error {
	if s.En != nil {
		if err := s.En.Validate(); err != nil {
			return err
		}
	}
	if s.Zh != nil {
		if err := s.Zh.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSliderCaptchaPageRequestConfigsEn struct {
	// Specifies whether to show the footer.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ShowFooter *bool `json:"ShowFooter,omitempty" xml:"ShowFooter,omitempty"`
	// Specifies whether to show the request ID.
	//
	// example:
	//
	// true
	ShowRequestId *bool `json:"ShowRequestId,omitempty" xml:"ShowRequestId,omitempty"`
	// The color of the slider.
	//
	// This parameter is required.
	//
	// example:
	//
	// #FFFFFF
	SliderColor *string `json:"SliderColor,omitempty" xml:"SliderColor,omitempty"`
	// The title.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSliderCaptchaPageRequestConfigsEn) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageRequestConfigsEn) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) GetShowFooter() *bool {
	return s.ShowFooter
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) GetShowRequestId() *bool {
	return s.ShowRequestId
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) GetSliderColor() *string {
	return s.SliderColor
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) GetTitle() *string {
	return s.Title
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) SetShowFooter(v bool) *CreateSliderCaptchaPageRequestConfigsEn {
	s.ShowFooter = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) SetShowRequestId(v bool) *CreateSliderCaptchaPageRequestConfigsEn {
	s.ShowRequestId = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) SetSliderColor(v string) *CreateSliderCaptchaPageRequestConfigsEn {
	s.SliderColor = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) SetTitle(v string) *CreateSliderCaptchaPageRequestConfigsEn {
	s.Title = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsEn) Validate() error {
	return dara.Validate(s)
}

type CreateSliderCaptchaPageRequestConfigsZh struct {
	// Specifies whether to show the footer.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	ShowFooter *bool `json:"ShowFooter,omitempty" xml:"ShowFooter,omitempty"`
	// Specifies whether to show the request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	ShowRequestId *bool `json:"ShowRequestId,omitempty" xml:"ShowRequestId,omitempty"`
	// The color of the slider.
	//
	// This parameter is required.
	//
	// example:
	//
	// #FFFFFF
	SliderColor *string `json:"SliderColor,omitempty" xml:"SliderColor,omitempty"`
	// The title.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSliderCaptchaPageRequestConfigsZh) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageRequestConfigsZh) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) GetShowFooter() *bool {
	return s.ShowFooter
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) GetShowRequestId() *bool {
	return s.ShowRequestId
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) GetSliderColor() *string {
	return s.SliderColor
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) GetTitle() *string {
	return s.Title
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) SetShowFooter(v bool) *CreateSliderCaptchaPageRequestConfigsZh {
	s.ShowFooter = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) SetShowRequestId(v bool) *CreateSliderCaptchaPageRequestConfigsZh {
	s.ShowRequestId = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) SetSliderColor(v string) *CreateSliderCaptchaPageRequestConfigsZh {
	s.SliderColor = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) SetTitle(v string) *CreateSliderCaptchaPageRequestConfigsZh {
	s.Title = &v
	return s
}

func (s *CreateSliderCaptchaPageRequestConfigsZh) Validate() error {
	return dara.Validate(s)
}
