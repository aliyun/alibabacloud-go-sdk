// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAIProduceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *AddLiveAIProduceRulesRequest
	GetApp() *string
	SetDescription(v string) *AddLiveAIProduceRulesRequest
	GetDescription() *string
	SetDomain(v string) *AddLiveAIProduceRulesRequest
	GetDomain() *string
	SetIsLazy(v bool) *AddLiveAIProduceRulesRequest
	GetIsLazy() *bool
	SetLiveTemplate(v string) *AddLiveAIProduceRulesRequest
	GetLiveTemplate() *string
	SetOwnerId(v int64) *AddLiveAIProduceRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveAIProduceRulesRequest
	GetRegionId() *string
	SetStudioName(v string) *AddLiveAIProduceRulesRequest
	GetStudioName() *string
	SetSubtitleName(v string) *AddLiveAIProduceRulesRequest
	GetSubtitleName() *string
	SetSuffix(v string) *AddLiveAIProduceRulesRequest
	GetSuffix() *string
}

type AddLiveAIProduceRulesRequest struct {
	// The name of the live streaming application. The name can be up to 256 characters long and can contain digits, uppercase and lowercase letters, hyphens (-), and underscores (_). The AppName must match the AppName in the ingest URL for the template to take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The description of the subtitle rule. The description can contain Chinese and English characters, digits, and special characters. It can be up to 128 characters in length.
	//
	// example:
	//
	// live AI subtitle template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether the rule is triggered by stream pulling. Valid values:
	//
	// - true: Subtitles are generated when stream pulling starts. If no stream is pulled for 5 minutes, subtitle generation stops. Subtitle generation resumes when stream pulling starts again.
	//
	// - false: Subtitles are generated when stream ingest starts, regardless of whether a stream is being pulled.
	//
	// example:
	//
	// true
	IsLazy *bool `json:"IsLazy,omitempty" xml:"IsLazy,omitempty"`
	// The specifications of the output subtitles. Valid values:
	//
	// - Landscape low definition 360p (640 × 360): `lp_ld`
	//
	// - Portrait low definition 360p (360 × 640): `lp_ld_v`
	//
	// - Landscape standard definition 480p (854 × 480): `lp_sd`
	//
	// - Portrait standard definition 480p (480 × 854): `lp_sd_v`
	//
	// - Landscape high definition 720p (1280 × 720): `lp_hd`
	//
	// - Portrait high definition 720p (720 × 1280): `lp_hd_v`
	//
	// - Landscape ultra high definition 1080p (1920 × 1080): `lp_ud`
	//
	// - Portrait ultra high definition 1080p (1080 × 1920): `lp_ud_v`
	//
	// This parameter is required.
	//
	// example:
	//
	// lp_ld
	LiveTemplate *string `json:"LiveTemplate,omitempty" xml:"LiveTemplate,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the virtual background template.
	//
	// example:
	//
	// sub02
	StudioName *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
	// The name of the subtitle template.
	//
	// example:
	//
	// sub01
	SubtitleName *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
	// The suffix to match.
	//
	// example:
	//
	// test01
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s AddLiveAIProduceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAIProduceRulesRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAIProduceRulesRequest) GetApp() *string {
	return s.App
}

func (s *AddLiveAIProduceRulesRequest) GetDescription() *string {
	return s.Description
}

func (s *AddLiveAIProduceRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddLiveAIProduceRulesRequest) GetIsLazy() *bool {
	return s.IsLazy
}

func (s *AddLiveAIProduceRulesRequest) GetLiveTemplate() *string {
	return s.LiveTemplate
}

func (s *AddLiveAIProduceRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveAIProduceRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveAIProduceRulesRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *AddLiveAIProduceRulesRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *AddLiveAIProduceRulesRequest) GetSuffix() *string {
	return s.Suffix
}

func (s *AddLiveAIProduceRulesRequest) SetApp(v string) *AddLiveAIProduceRulesRequest {
	s.App = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetDescription(v string) *AddLiveAIProduceRulesRequest {
	s.Description = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetDomain(v string) *AddLiveAIProduceRulesRequest {
	s.Domain = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetIsLazy(v bool) *AddLiveAIProduceRulesRequest {
	s.IsLazy = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetLiveTemplate(v string) *AddLiveAIProduceRulesRequest {
	s.LiveTemplate = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetOwnerId(v int64) *AddLiveAIProduceRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetRegionId(v string) *AddLiveAIProduceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetStudioName(v string) *AddLiveAIProduceRulesRequest {
	s.StudioName = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetSubtitleName(v string) *AddLiveAIProduceRulesRequest {
	s.SubtitleName = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) SetSuffix(v string) *AddLiveAIProduceRulesRequest {
	s.Suffix = &v
	return s
}

func (s *AddLiveAIProduceRulesRequest) Validate() error {
	return dara.Validate(s)
}
