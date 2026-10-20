// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSliderCaptchaPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *GetSliderCaptchaPageResponseBodyConfigs) *GetSliderCaptchaPageResponseBody
	GetConfigs() *GetSliderCaptchaPageResponseBodyConfigs
	SetCustomContent(v string) *GetSliderCaptchaPageResponseBody
	GetCustomContent() *string
	SetForceLanguage(v string) *GetSliderCaptchaPageResponseBody
	GetForceLanguage() *string
	SetId(v int64) *GetSliderCaptchaPageResponseBody
	GetId() *int64
	SetName(v string) *GetSliderCaptchaPageResponseBody
	GetName() *string
	SetRequestId(v string) *GetSliderCaptchaPageResponseBody
	GetRequestId() *string
	SetSiteIds(v []*int64) *GetSliderCaptchaPageResponseBody
	GetSiteIds() []*int64
}

type GetSliderCaptchaPageResponseBody struct {
	// The predefined configurations for the custom slider page.
	Configs *GetSliderCaptchaPageResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
	// The custom slider configurations.
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
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of bound site IDs.
	SiteIds []*int64 `json:"SiteIds,omitempty" xml:"SiteIds,omitempty" type:"Repeated"`
}

func (s GetSliderCaptchaPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSliderCaptchaPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSliderCaptchaPageResponseBody) GetConfigs() *GetSliderCaptchaPageResponseBodyConfigs {
	return s.Configs
}

func (s *GetSliderCaptchaPageResponseBody) GetCustomContent() *string {
	return s.CustomContent
}

func (s *GetSliderCaptchaPageResponseBody) GetForceLanguage() *string {
	return s.ForceLanguage
}

func (s *GetSliderCaptchaPageResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetSliderCaptchaPageResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSliderCaptchaPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSliderCaptchaPageResponseBody) GetSiteIds() []*int64 {
	return s.SiteIds
}

func (s *GetSliderCaptchaPageResponseBody) SetConfigs(v *GetSliderCaptchaPageResponseBodyConfigs) *GetSliderCaptchaPageResponseBody {
	s.Configs = v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) SetCustomContent(v string) *GetSliderCaptchaPageResponseBody {
	s.CustomContent = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) SetForceLanguage(v string) *GetSliderCaptchaPageResponseBody {
	s.ForceLanguage = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) SetId(v int64) *GetSliderCaptchaPageResponseBody {
	s.Id = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) SetName(v string) *GetSliderCaptchaPageResponseBody {
	s.Name = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) SetRequestId(v string) *GetSliderCaptchaPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) SetSiteIds(v []*int64) *GetSliderCaptchaPageResponseBody {
	s.SiteIds = v
	return s
}

func (s *GetSliderCaptchaPageResponseBody) Validate() error {
	if s.Configs != nil {
		if err := s.Configs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSliderCaptchaPageResponseBodyConfigs struct {
	// The predefined English style.
	En *GetSliderCaptchaPageResponseBodyConfigsEn `json:"En,omitempty" xml:"En,omitempty" type:"Struct"`
	// The predefined Chinese style.
	Zh *GetSliderCaptchaPageResponseBodyConfigsZh `json:"Zh,omitempty" xml:"Zh,omitempty" type:"Struct"`
}

func (s GetSliderCaptchaPageResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetSliderCaptchaPageResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *GetSliderCaptchaPageResponseBodyConfigs) GetEn() *GetSliderCaptchaPageResponseBodyConfigsEn {
	return s.En
}

func (s *GetSliderCaptchaPageResponseBodyConfigs) GetZh() *GetSliderCaptchaPageResponseBodyConfigsZh {
	return s.Zh
}

func (s *GetSliderCaptchaPageResponseBodyConfigs) SetEn(v *GetSliderCaptchaPageResponseBodyConfigsEn) *GetSliderCaptchaPageResponseBodyConfigs {
	s.En = v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigs) SetZh(v *GetSliderCaptchaPageResponseBodyConfigsZh) *GetSliderCaptchaPageResponseBodyConfigs {
	s.Zh = v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigs) Validate() error {
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

type GetSliderCaptchaPageResponseBodyConfigsEn struct {
	// Specifies whether to display the footer.
	//
	// example:
	//
	// true
	ShowFooter *bool `json:"ShowFooter,omitempty" xml:"ShowFooter,omitempty"`
	// Specifies whether to display the request ID.
	//
	// example:
	//
	// true
	ShowRequestId *bool `json:"ShowRequestId,omitempty" xml:"ShowRequestId,omitempty"`
	// The slider color.
	//
	// example:
	//
	// #FFFFFF
	SliderColor *string `json:"SliderColor,omitempty" xml:"SliderColor,omitempty"`
	// The title.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSliderCaptchaPageResponseBodyConfigsEn) String() string {
	return dara.Prettify(s)
}

func (s GetSliderCaptchaPageResponseBodyConfigsEn) GoString() string {
	return s.String()
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) GetShowFooter() *bool {
	return s.ShowFooter
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) GetShowRequestId() *bool {
	return s.ShowRequestId
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) GetSliderColor() *string {
	return s.SliderColor
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) GetTitle() *string {
	return s.Title
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) SetShowFooter(v bool) *GetSliderCaptchaPageResponseBodyConfigsEn {
	s.ShowFooter = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) SetShowRequestId(v bool) *GetSliderCaptchaPageResponseBodyConfigsEn {
	s.ShowRequestId = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) SetSliderColor(v string) *GetSliderCaptchaPageResponseBodyConfigsEn {
	s.SliderColor = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) SetTitle(v string) *GetSliderCaptchaPageResponseBodyConfigsEn {
	s.Title = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsEn) Validate() error {
	return dara.Validate(s)
}

type GetSliderCaptchaPageResponseBodyConfigsZh struct {
	// Specifies whether to display the footer.
	//
	// example:
	//
	// false
	ShowFooter *bool `json:"ShowFooter,omitempty" xml:"ShowFooter,omitempty"`
	// Specifies whether to display the request ID.
	//
	// example:
	//
	// false
	ShowRequestId *bool `json:"ShowRequestId,omitempty" xml:"ShowRequestId,omitempty"`
	// The slider color.
	//
	// example:
	//
	// #FFFFFF
	SliderColor *string `json:"SliderColor,omitempty" xml:"SliderColor,omitempty"`
	// The title.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSliderCaptchaPageResponseBodyConfigsZh) String() string {
	return dara.Prettify(s)
}

func (s GetSliderCaptchaPageResponseBodyConfigsZh) GoString() string {
	return s.String()
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) GetShowFooter() *bool {
	return s.ShowFooter
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) GetShowRequestId() *bool {
	return s.ShowRequestId
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) GetSliderColor() *string {
	return s.SliderColor
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) GetTitle() *string {
	return s.Title
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) SetShowFooter(v bool) *GetSliderCaptchaPageResponseBodyConfigsZh {
	s.ShowFooter = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) SetShowRequestId(v bool) *GetSliderCaptchaPageResponseBodyConfigsZh {
	s.ShowRequestId = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) SetSliderColor(v string) *GetSliderCaptchaPageResponseBodyConfigsZh {
	s.SliderColor = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) SetTitle(v string) *GetSliderCaptchaPageResponseBodyConfigsZh {
	s.Title = &v
	return s
}

func (s *GetSliderCaptchaPageResponseBodyConfigsZh) Validate() error {
	return dara.Validate(s)
}
