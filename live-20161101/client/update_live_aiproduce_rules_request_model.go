// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAIProduceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *UpdateLiveAIProduceRulesRequest
	GetApp() *string
	SetDescription(v string) *UpdateLiveAIProduceRulesRequest
	GetDescription() *string
	SetDomain(v string) *UpdateLiveAIProduceRulesRequest
	GetDomain() *string
	SetIsLazy(v bool) *UpdateLiveAIProduceRulesRequest
	GetIsLazy() *bool
	SetLiveTemplate(v string) *UpdateLiveAIProduceRulesRequest
	GetLiveTemplate() *string
	SetOwnerId(v int64) *UpdateLiveAIProduceRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveAIProduceRulesRequest
	GetRegionId() *string
	SetRulesId(v string) *UpdateLiveAIProduceRulesRequest
	GetRulesId() *string
	SetStudioName(v string) *UpdateLiveAIProduceRulesRequest
	GetStudioName() *string
	SetSubtitleId(v string) *UpdateLiveAIProduceRulesRequest
	GetSubtitleId() *string
	SetSubtitleName(v string) *UpdateLiveAIProduceRulesRequest
	GetSubtitleName() *string
	SetSuffix(v string) *UpdateLiveAIProduceRulesRequest
	GetSuffix() *string
}

type UpdateLiveAIProduceRulesRequest struct {
	// The name of the live stream application.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The description of the subtitle rule. The description can contain letters, digits, Chinese characters, and special characters, and can be up to 128 characters in length.
	//
	// example:
	//
	// live AI subtitle template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The primary streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether subtitles are triggered by stream pulling. Valid values:
	//
	// - true: Subtitles start when a stream is pulled. If no stream is pulled within 5 minutes, the subtitles stop. Subtitles restart when a stream is pulled again.
	//
	// - false: Subtitles start as long as stream ingest is active, regardless of whether a stream is being pulled.
	//
	// example:
	//
	// true
	IsLazy *bool `json:"IsLazy,omitempty" xml:"IsLazy,omitempty"`
	// The output specification of the subtitle. Valid values:
	//
	// - Landscape low definition 360P 640 × 360: `lp_ld`
	//
	// - Portrait low definition 360P 360 × 640: `lp_ld_v`
	//
	// - Landscape standard definition 480P 854 × 480: `lp_sd`
	//
	// - Portrait standard definition 480P 480 × 854: `lp_sd_v`
	//
	// - Landscape high definition 720P 1280 × 720: `lp_hd`
	//
	// - Portrait high definition 720P 720 × 1280: `lp_hd_v`
	//
	// - Landscape ultra-high definition 1080P 1920 × 1080: `lp_ud`
	//
	// - Portrait ultra-high definition 1080P 1080 × 1920: `lp_ud_v`
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
	// The ID of the subtitle rule.
	//
	// example:
	//
	// 445409ec-7eaa-461d -8f29-4bec2eb9****
	RulesId *string `json:"RulesId,omitempty" xml:"RulesId,omitempty"`
	// The name of the virtual background template. You must specify at least one of SubtitleName and StudioName. Otherwise, a MissingParameter error is returned.
	//
	// example:
	//
	// sub02
	StudioName *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
	// The ID of the subtitle template.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	SubtitleId *string `json:"SubtitleId,omitempty" xml:"SubtitleId,omitempty"`
	// The name of the subtitle template. You must specify at least one of SubtitleName and StudioName. Otherwise, a MissingParameter error is returned.
	//
	// example:
	//
	// sub01
	SubtitleName *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
	// The suffix match.
	//
	// example:
	//
	// test01
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s UpdateLiveAIProduceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAIProduceRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAIProduceRulesRequest) GetApp() *string {
	return s.App
}

func (s *UpdateLiveAIProduceRulesRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLiveAIProduceRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateLiveAIProduceRulesRequest) GetIsLazy() *bool {
	return s.IsLazy
}

func (s *UpdateLiveAIProduceRulesRequest) GetLiveTemplate() *string {
	return s.LiveTemplate
}

func (s *UpdateLiveAIProduceRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAIProduceRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveAIProduceRulesRequest) GetRulesId() *string {
	return s.RulesId
}

func (s *UpdateLiveAIProduceRulesRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *UpdateLiveAIProduceRulesRequest) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *UpdateLiveAIProduceRulesRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *UpdateLiveAIProduceRulesRequest) GetSuffix() *string {
	return s.Suffix
}

func (s *UpdateLiveAIProduceRulesRequest) SetApp(v string) *UpdateLiveAIProduceRulesRequest {
	s.App = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetDescription(v string) *UpdateLiveAIProduceRulesRequest {
	s.Description = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetDomain(v string) *UpdateLiveAIProduceRulesRequest {
	s.Domain = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetIsLazy(v bool) *UpdateLiveAIProduceRulesRequest {
	s.IsLazy = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetLiveTemplate(v string) *UpdateLiveAIProduceRulesRequest {
	s.LiveTemplate = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetOwnerId(v int64) *UpdateLiveAIProduceRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetRegionId(v string) *UpdateLiveAIProduceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetRulesId(v string) *UpdateLiveAIProduceRulesRequest {
	s.RulesId = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetStudioName(v string) *UpdateLiveAIProduceRulesRequest {
	s.StudioName = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetSubtitleId(v string) *UpdateLiveAIProduceRulesRequest {
	s.SubtitleId = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetSubtitleName(v string) *UpdateLiveAIProduceRulesRequest {
	s.SubtitleName = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) SetSuffix(v string) *UpdateLiveAIProduceRulesRequest {
	s.Suffix = &v
	return s
}

func (s *UpdateLiveAIProduceRulesRequest) Validate() error {
	return dara.Validate(s)
}
