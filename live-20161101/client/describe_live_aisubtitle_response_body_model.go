// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAISubtitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveAISubtitleResponseBody
	GetRequestId() *string
	SetSubtitleConfigs(v *DescribeLiveAISubtitleResponseBodySubtitleConfigs) *DescribeLiveAISubtitleResponseBody
	GetSubtitleConfigs() *DescribeLiveAISubtitleResponseBodySubtitleConfigs
}

type DescribeLiveAISubtitleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b96****
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubtitleConfigs *DescribeLiveAISubtitleResponseBodySubtitleConfigs `json:"SubtitleConfigs,omitempty" xml:"SubtitleConfigs,omitempty" type:"Struct"`
}

func (s DescribeLiveAISubtitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveAISubtitleResponseBody) GetSubtitleConfigs() *DescribeLiveAISubtitleResponseBodySubtitleConfigs {
	return s.SubtitleConfigs
}

func (s *DescribeLiveAISubtitleResponseBody) SetRequestId(v string) *DescribeLiveAISubtitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBody) SetSubtitleConfigs(v *DescribeLiveAISubtitleResponseBodySubtitleConfigs) *DescribeLiveAISubtitleResponseBody {
	s.SubtitleConfigs = v
	return s
}

func (s *DescribeLiveAISubtitleResponseBody) Validate() error {
	if s.SubtitleConfigs != nil {
		if err := s.SubtitleConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveAISubtitleResponseBodySubtitleConfigs struct {
	SubtitleConfig []*DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig `json:"SubtitleConfig,omitempty" xml:"SubtitleConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigs) GetSubtitleConfig() []*DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	return s.SubtitleConfig
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigs) SetSubtitleConfig(v []*DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) *DescribeLiveAISubtitleResponseBodySubtitleConfigs {
	s.SubtitleConfig = v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigs) Validate() error {
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

type DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig struct {
	BgColor               *string                                                                            `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	BgWidthNormalized     *float32                                                                           `json:"BgWidthNormalized,omitempty" xml:"BgWidthNormalized,omitempty"`
	BorderWidthNormalized *float32                                                                           `json:"BorderWidthNormalized,omitempty" xml:"BorderWidthNormalized,omitempty"`
	Description           *string                                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	DstLanguage           *string                                                                            `json:"DstLanguage,omitempty" xml:"DstLanguage,omitempty"`
	FontColor             *string                                                                            `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontName              *string                                                                            `json:"FontName,omitempty" xml:"FontName,omitempty"`
	FontSizeNormalized    *string                                                                            `json:"FontSizeNormalized,omitempty" xml:"FontSizeNormalized,omitempty"`
	Height                *string                                                                            `json:"Height,omitempty" xml:"Height,omitempty"`
	MaxLines              *int32                                                                             `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	PositionNormalized    *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Struct"`
	RulesRefer            *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer         `json:"RulesRefer,omitempty" xml:"RulesRefer,omitempty" type:"Struct"`
	ShowSourceLan         *int32                                                                             `json:"ShowSourceLan,omitempty" xml:"ShowSourceLan,omitempty"`
	SrcLanguage           *string                                                                            `json:"SrcLanguage,omitempty" xml:"SrcLanguage,omitempty"`
	SubtitleId            *string                                                                            `json:"SubtitleId,omitempty" xml:"SubtitleId,omitempty"`
	SubtitleName          *string                                                                            `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
	Width                 *string                                                                            `json:"Width,omitempty" xml:"Width,omitempty"`
	WordPerline           *int32                                                                             `json:"WordPerline,omitempty" xml:"WordPerline,omitempty"`
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetBgColor() *string {
	return s.BgColor
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetBgWidthNormalized() *float32 {
	return s.BgWidthNormalized
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetBorderWidthNormalized() *float32 {
	return s.BorderWidthNormalized
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetDstLanguage() *string {
	return s.DstLanguage
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetFontColor() *string {
	return s.FontColor
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetFontName() *string {
	return s.FontName
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetFontSizeNormalized() *string {
	return s.FontSizeNormalized
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetHeight() *string {
	return s.Height
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetMaxLines() *int32 {
	return s.MaxLines
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetPositionNormalized() *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized {
	return s.PositionNormalized
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetRulesRefer() *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer {
	return s.RulesRefer
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetShowSourceLan() *int32 {
	return s.ShowSourceLan
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetSrcLanguage() *string {
	return s.SrcLanguage
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetWidth() *string {
	return s.Width
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) GetWordPerline() *int32 {
	return s.WordPerline
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetBgColor(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.BgColor = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetBgWidthNormalized(v float32) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.BgWidthNormalized = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetBorderWidthNormalized(v float32) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.BorderWidthNormalized = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetDescription(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.Description = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetDstLanguage(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.DstLanguage = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetFontColor(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.FontColor = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetFontName(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.FontName = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetFontSizeNormalized(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.FontSizeNormalized = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetHeight(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.Height = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetMaxLines(v int32) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.MaxLines = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetPositionNormalized(v *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.PositionNormalized = v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetRulesRefer(v *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.RulesRefer = v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetShowSourceLan(v int32) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.ShowSourceLan = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetSrcLanguage(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.SrcLanguage = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetSubtitleId(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.SubtitleId = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetSubtitleName(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.SubtitleName = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetWidth(v string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.Width = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) SetWordPerline(v int32) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig {
	s.WordPerline = &v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfig) Validate() error {
	if s.PositionNormalized != nil {
		if err := s.PositionNormalized.Validate(); err != nil {
			return err
		}
	}
	if s.RulesRefer != nil {
		if err := s.RulesRefer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized struct {
	Position []*float32 `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized) GetPosition() []*float32 {
	return s.Position
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized) SetPosition(v []*float32) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized {
	s.Position = v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigPositionNormalized) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer struct {
	RulesId []*string `json:"RulesId,omitempty" xml:"RulesId,omitempty" type:"Repeated"`
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer) GetRulesId() []*string {
	return s.RulesId
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer) SetRulesId(v []*string) *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer {
	s.RulesId = v
	return s
}

func (s *DescribeLiveAISubtitleResponseBodySubtitleConfigsSubtitleConfigRulesRefer) Validate() error {
	return dara.Validate(s)
}
