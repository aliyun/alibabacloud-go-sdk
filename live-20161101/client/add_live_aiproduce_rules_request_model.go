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
	// The name of the application to which the live stream belongs. The name can be up to 256 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name must be the same as the application name in the ingest URL. Otherwise, the rule does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The description of the subtitle rule. The description can be up to 128 characters in length and can contain letters, digits, and special characters.
	//
	// example:
	//
	// live AI subtitle template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to trigger the subtitle rule when stream pulling starts. Valid values:
	//
	// 	- true: generates live subtitles when stream pulling starts and stops generating live subtitles when no stream is pulled for 5 minutes. When stream pulling restarts, live subtitles are generated again.
	//
	// 	- false: generates live subtitles when stream ingest starts, regardless of whether stream pulling starts.
	//
	// example:
	//
	// true
	IsLazy *bool `json:"IsLazy,omitempty" xml:"IsLazy,omitempty"`
	// The specification of the output subtitles. Valid values:
	//
	// 	- `lp_ld`: landscape low definition 360p (640×360)
	//
	// 	- `lp_ld_v`: portrait low definition 360p (360×640)
	//
	// 	- `lp_sd`: landscape standard definition 480p (854×480)
	//
	// 	- `lp_sd_v`: portrait standard definition 480p (480×854)
	//
	// 	- `lp_hd`: landscape high definition 720p (1280×720)
	//
	// 	- `lp_hd_v`: portrait high definition 720p (720×1280)
	//
	// 	- `lp_ud`: landscape ultra-high definition 1080p (1920×1080)
	//
	// 	- `lp_ud_v`: portrait ultra-high definition 1080p (1080×1920)
	//
	// This parameter is required.
	//
	// example:
	//
	// lp_ld
	LiveTemplate *string `json:"LiveTemplate,omitempty" xml:"LiveTemplate,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
