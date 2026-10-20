// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSliderCaptchaPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *UpdateSliderCaptchaPageRequestConfigs) *UpdateSliderCaptchaPageRequest
	GetConfigs() *UpdateSliderCaptchaPageRequestConfigs
	SetCustomContent(v string) *UpdateSliderCaptchaPageRequest
	GetCustomContent() *string
	SetForceLanguage(v string) *UpdateSliderCaptchaPageRequest
	GetForceLanguage() *string
	SetId(v int64) *UpdateSliderCaptchaPageRequest
	GetId() *int64
	SetName(v string) *UpdateSliderCaptchaPageRequest
	GetName() *string
	SetSiteIds(v []*int64) *UpdateSliderCaptchaPageRequest
	GetSiteIds() []*int64
}

type UpdateSliderCaptchaPageRequest struct {
	// The predefined configurations of the custom slider page.
	Configs *UpdateSliderCaptchaPageRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
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
	SiteIds []*int64 `json:"SiteIds,omitempty" xml:"SiteIds,omitempty" type:"Repeated"`
}

func (s UpdateSliderCaptchaPageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageRequest) GetConfigs() *UpdateSliderCaptchaPageRequestConfigs {
	return s.Configs
}

func (s *UpdateSliderCaptchaPageRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *UpdateSliderCaptchaPageRequest) GetForceLanguage() *string {
	return s.ForceLanguage
}

func (s *UpdateSliderCaptchaPageRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateSliderCaptchaPageRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSliderCaptchaPageRequest) GetSiteIds() []*int64 {
	return s.SiteIds
}

func (s *UpdateSliderCaptchaPageRequest) SetConfigs(v *UpdateSliderCaptchaPageRequestConfigs) *UpdateSliderCaptchaPageRequest {
	s.Configs = v
	return s
}

func (s *UpdateSliderCaptchaPageRequest) SetCustomContent(v string) *UpdateSliderCaptchaPageRequest {
	s.CustomContent = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequest) SetForceLanguage(v string) *UpdateSliderCaptchaPageRequest {
	s.ForceLanguage = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequest) SetId(v int64) *UpdateSliderCaptchaPageRequest {
	s.Id = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequest) SetName(v string) *UpdateSliderCaptchaPageRequest {
	s.Name = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequest) SetSiteIds(v []*int64) *UpdateSliderCaptchaPageRequest {
	s.SiteIds = v
	return s
}

func (s *UpdateSliderCaptchaPageRequest) Validate() error {
	if s.Configs != nil {
		if err := s.Configs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSliderCaptchaPageRequestConfigs struct {
	// The predefined English style.
	En *UpdateSliderCaptchaPageRequestConfigsEn `json:"En,omitempty" xml:"En,omitempty" type:"Struct"`
	// The predefined Chinese style.
	Zh *UpdateSliderCaptchaPageRequestConfigsZh `json:"Zh,omitempty" xml:"Zh,omitempty" type:"Struct"`
}

func (s UpdateSliderCaptchaPageRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageRequestConfigs) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageRequestConfigs) GetEn() *UpdateSliderCaptchaPageRequestConfigsEn {
	return s.En
}

func (s *UpdateSliderCaptchaPageRequestConfigs) GetZh() *UpdateSliderCaptchaPageRequestConfigsZh {
	return s.Zh
}

func (s *UpdateSliderCaptchaPageRequestConfigs) SetEn(v *UpdateSliderCaptchaPageRequestConfigsEn) *UpdateSliderCaptchaPageRequestConfigs {
	s.En = v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigs) SetZh(v *UpdateSliderCaptchaPageRequestConfigsZh) *UpdateSliderCaptchaPageRequestConfigs {
	s.Zh = v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigs) Validate() error {
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

type UpdateSliderCaptchaPageRequestConfigsEn struct {
	// Specifies whether to display the footer.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ShowFooter *bool `json:"ShowFooter,omitempty" xml:"ShowFooter,omitempty"`
	// Specifies whether to display the request ID.
	//
	// This parameter is required.
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

func (s UpdateSliderCaptchaPageRequestConfigsEn) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageRequestConfigsEn) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) GetShowFooter() *bool {
	return s.ShowFooter
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) GetShowRequestId() *bool {
	return s.ShowRequestId
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) GetSliderColor() *string {
	return s.SliderColor
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) GetTitle() *string {
	return s.Title
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) SetShowFooter(v bool) *UpdateSliderCaptchaPageRequestConfigsEn {
	s.ShowFooter = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) SetShowRequestId(v bool) *UpdateSliderCaptchaPageRequestConfigsEn {
	s.ShowRequestId = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) SetSliderColor(v string) *UpdateSliderCaptchaPageRequestConfigsEn {
	s.SliderColor = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) SetTitle(v string) *UpdateSliderCaptchaPageRequestConfigsEn {
	s.Title = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsEn) Validate() error {
	return dara.Validate(s)
}

type UpdateSliderCaptchaPageRequestConfigsZh struct {
	// Specifies whether to display the footer.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	ShowFooter *bool `json:"ShowFooter,omitempty" xml:"ShowFooter,omitempty"`
	// Specifies whether to display the request ID.
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

func (s UpdateSliderCaptchaPageRequestConfigsZh) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageRequestConfigsZh) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) GetShowFooter() *bool {
	return s.ShowFooter
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) GetShowRequestId() *bool {
	return s.ShowRequestId
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) GetSliderColor() *string {
	return s.SliderColor
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) GetTitle() *string {
	return s.Title
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) SetShowFooter(v bool) *UpdateSliderCaptchaPageRequestConfigsZh {
	s.ShowFooter = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) SetShowRequestId(v bool) *UpdateSliderCaptchaPageRequestConfigsZh {
	s.ShowRequestId = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) SetSliderColor(v string) *UpdateSliderCaptchaPageRequestConfigsZh {
	s.SliderColor = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) SetTitle(v string) *UpdateSliderCaptchaPageRequestConfigsZh {
	s.Title = &v
	return s
}

func (s *UpdateSliderCaptchaPageRequestConfigsZh) Validate() error {
	return dara.Validate(s)
}
