// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAIStudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeLiveAIStudioResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveAIStudioResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveAIStudioResponseBody
	GetRequestId() *string
	SetStudioConfigs(v *DescribeLiveAIStudioResponseBodyStudioConfigs) *DescribeLiveAIStudioResponseBody
	GetStudioConfigs() *DescribeLiveAIStudioResponseBodyStudioConfigs
	SetTotal(v int32) *DescribeLiveAIStudioResponseBody
	GetTotal() *int32
}

type DescribeLiveAIStudioResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The virtual studio templates.
	StudioConfigs *DescribeLiveAIStudioResponseBodyStudioConfigs `json:"StudioConfigs,omitempty" xml:"StudioConfigs,omitempty" type:"Struct"`
	// The total number of templates.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveAIStudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIStudioResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIStudioResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveAIStudioResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveAIStudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveAIStudioResponseBody) GetStudioConfigs() *DescribeLiveAIStudioResponseBodyStudioConfigs {
	return s.StudioConfigs
}

func (s *DescribeLiveAIStudioResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLiveAIStudioResponseBody) SetPageNumber(v int32) *DescribeLiveAIStudioResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBody) SetPageSize(v int32) *DescribeLiveAIStudioResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBody) SetRequestId(v string) *DescribeLiveAIStudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBody) SetStudioConfigs(v *DescribeLiveAIStudioResponseBodyStudioConfigs) *DescribeLiveAIStudioResponseBody {
	s.StudioConfigs = v
	return s
}

func (s *DescribeLiveAIStudioResponseBody) SetTotal(v int32) *DescribeLiveAIStudioResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAIStudioResponseBodyStudioConfigs struct {
	SubtitleConfig []*DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig `json:"SubtitleConfig,omitempty" xml:"SubtitleConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveAIStudioResponseBodyStudioConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIStudioResponseBodyStudioConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigs) GetSubtitleConfig() []*DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	return s.SubtitleConfig
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigs) SetSubtitleConfig(v []*DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) *DescribeLiveAIStudioResponseBodyStudioConfigs {
	s.SubtitleConfig = v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig struct {
	// The ID of the background material.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	BackgroundResourceId *string `json:"BackgroundResourceId,omitempty" xml:"BackgroundResourceId,omitempty"`
	// The URL of the background material. Make sure that the URL is accessible over the Internet. Either this parameter or the BackgroundResourceId parameter is returned.
	//
	// example:
	//
	// https://testbucket.xx.com/2.jpg
	BackgroundResourceUrl *string `json:"BackgroundResourceUrl,omitempty" xml:"BackgroundResourceUrl,omitempty"`
	// The type of the background material. Valid values:
	//
	// 	- VOD: a video in ApsaraVideo VOD
	//
	// 	- PIC: an image
	//
	// 	- LIVE: a live stream
	//
	// example:
	//
	// VOD
	BackgroundType *string `json:"BackgroundType,omitempty" xml:"BackgroundType,omitempty"`
	// The custom description.
	//
	// example:
	//
	// user defined description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The preview height. Unit: pixels.
	//
	// The following preview specifications (width × height) are supported:
	//
	// 	- Landscape low definition 360p (640×360)
	//
	// 	- Portrait low definition 360p (360×640)
	//
	// 	- Landscape standard definition 480p (854×480)
	//
	// 	- Portrait standard definition 480p (480×854)
	//
	// 	- Landscape high definition 720p (1280×720)
	//
	// 	- Portrait high definition 720p (720×1280)
	//
	// 	- Landscape ultra-high definition 1080p (1920×1080)
	//
	// 	- Portrait ultra-high definition 1080p (1080×1920)
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The layout information of the multimedia material.
	//
	// example:
	//
	// "{\\"positionY\\":\\"0.0\\",\\"positionX\\":\\"0.0\\",\\"heightNormalized\\":\\"0.5\\"}"
	MattingLayout *string `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty"`
	// The type of chroma key that is performed on ingested streams. Valid values:
	//
	// 	- green: green-key chroma key
	//
	// 	- blue: blue-screen chroma key
	//
	// 	- complex: background replacement
	//
	// example:
	//
	// complex
	MattingType *string `json:"MattingType,omitempty" xml:"MattingType,omitempty"`
	// LIVE, live streaming
	//
	// example:
	//
	// "{\\"positionY\\":\\"0.0\\",\\"positionX\\":\\"0.0\\",\\"heightNormalized\\":\\"0.5\\"}"
	MediaLayout *string `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty"`
	// The ID of the multimedia material in ApsaraVideo VOD.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	MediaResourceId *string `json:"MediaResourceId,omitempty" xml:"MediaResourceId,omitempty"`
	// The URL of the multimedia material.
	//
	// example:
	//
	// https://testbucket.xx.com/2.jpg
	MediaResourceUrl *string `json:"MediaResourceUrl,omitempty" xml:"MediaResourceUrl,omitempty"`
	// The type of the multimedia material. Valid values:
	//
	// 	- VOD: a video in ApsaraVideo VOD
	//
	// 	- PIC: an image
	//
	// 	- LIVE: a live stream
	//
	// example:
	//
	// VOD
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The IDs of the bound rules.
	RuleIds *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The ID of the virtual studio template.
	//
	// example:
	//
	// 24654384-f5ac-40ea-823b-74e85a61dd9f
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the virtual studio template. The name is the same as the value of the StudioName parameter that was specified when you called the CreateLiveAiStudio operation to create the virtual studio template.
	//
	// example:
	//
	// studio1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The preview width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetBackgroundResourceId() *string {
	return s.BackgroundResourceId
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetBackgroundResourceUrl() *string {
	return s.BackgroundResourceUrl
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetBackgroundType() *string {
	return s.BackgroundType
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetHeight() *string {
	return s.Height
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetMattingLayout() *string {
	return s.MattingLayout
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetMattingType() *string {
	return s.MattingType
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetMediaLayout() *string {
	return s.MediaLayout
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetMediaResourceId() *string {
	return s.MediaResourceId
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetMediaResourceUrl() *string {
	return s.MediaResourceUrl
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetMediaType() *string {
	return s.MediaType
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetRuleIds() *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds {
	return s.RuleIds
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) GetWidth() *string {
	return s.Width
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetBackgroundResourceId(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.BackgroundResourceId = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetBackgroundResourceUrl(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.BackgroundResourceUrl = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetBackgroundType(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.BackgroundType = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetDescription(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.Description = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetHeight(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.Height = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetMattingLayout(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.MattingLayout = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetMattingType(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.MattingType = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetMediaLayout(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.MediaLayout = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetMediaResourceId(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.MediaResourceId = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetMediaResourceUrl(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.MediaResourceUrl = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetMediaType(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.MediaType = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetRuleIds(v *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.RuleIds = v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetTemplateId(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.TemplateId = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetTemplateName(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.TemplateName = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) SetWidth(v string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig {
	s.Width = &v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds struct {
	RuleId []*string `json:"ruleId,omitempty" xml:"ruleId,omitempty" type:"Repeated"`
}

func (s DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds) SetRuleId(v []*string) *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds) Validate() error {
	return dara.Validate(s)
}
