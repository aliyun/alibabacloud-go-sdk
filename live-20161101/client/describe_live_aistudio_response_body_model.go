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
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.StudioConfigs != nil {
		if err := s.StudioConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.SubtitleConfig != nil {
		for _, item := range s.SubtitleConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfig struct {
	BackgroundResourceId  *string                                                             `json:"BackgroundResourceId,omitempty" xml:"BackgroundResourceId,omitempty"`
	BackgroundResourceUrl *string                                                             `json:"BackgroundResourceUrl,omitempty" xml:"BackgroundResourceUrl,omitempty"`
	BackgroundType        *string                                                             `json:"BackgroundType,omitempty" xml:"BackgroundType,omitempty"`
	Description           *string                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Height                *string                                                             `json:"Height,omitempty" xml:"Height,omitempty"`
	MattingLayout         *string                                                             `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty"`
	MattingType           *string                                                             `json:"MattingType,omitempty" xml:"MattingType,omitempty"`
	MediaLayout           *string                                                             `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty"`
	MediaResourceId       *string                                                             `json:"MediaResourceId,omitempty" xml:"MediaResourceId,omitempty"`
	MediaResourceUrl      *string                                                             `json:"MediaResourceUrl,omitempty" xml:"MediaResourceUrl,omitempty"`
	MediaType             *string                                                             `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	RuleIds               *DescribeLiveAIStudioResponseBodyStudioConfigsSubtitleConfigRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	TemplateId            *string                                                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName          *string                                                             `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Width                 *string                                                             `json:"Width,omitempty" xml:"Width,omitempty"`
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
	if s.RuleIds != nil {
		if err := s.RuleIds.Validate(); err != nil {
			return err
		}
	}
	return nil
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
