// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v *DescribeCasterComponentsResponseBodyComponents) *DescribeCasterComponentsResponseBody
	GetComponents() *DescribeCasterComponentsResponseBodyComponents
	SetRequestId(v string) *DescribeCasterComponentsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCasterComponentsResponseBody
	GetTotal() *int32
}

type DescribeCasterComponentsResponseBody struct {
	// The components.
	Components *DescribeCasterComponentsResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3be7ade8-d907-483c-b24a-0dad45******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCasterComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBody) GetComponents() *DescribeCasterComponentsResponseBodyComponents {
	return s.Components
}

func (s *DescribeCasterComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterComponentsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterComponentsResponseBody) SetComponents(v *DescribeCasterComponentsResponseBodyComponents) *DescribeCasterComponentsResponseBody {
	s.Components = v
	return s
}

func (s *DescribeCasterComponentsResponseBody) SetRequestId(v string) *DescribeCasterComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterComponentsResponseBody) SetTotal(v int32) *DescribeCasterComponentsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterComponentsResponseBody) Validate() error {
	if s.Components != nil {
		if err := s.Components.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterComponentsResponseBodyComponents struct {
	Component []*DescribeCasterComponentsResponseBodyComponentsComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
}

func (s DescribeCasterComponentsResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponents) GetComponent() []*DescribeCasterComponentsResponseBodyComponentsComponent {
	return s.Component
}

func (s *DescribeCasterComponentsResponseBodyComponents) SetComponent(v []*DescribeCasterComponentsResponseBodyComponentsComponent) *DescribeCasterComponentsResponseBodyComponents {
	s.Component = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponents) Validate() error {
	if s.Component != nil {
		for _, item := range s.Component {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterComponentsResponseBodyComponentsComponent struct {
	// The information about the subtitle component.
	CaptionLayerContent *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty" type:"Struct"`
	// The component ID.
	//
	// example:
	//
	// 72d2ec7a-4cd7-4a01-974b-7cd53947****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The information about the component layer, such as the size and layout.
	ComponentLayer *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty" type:"Struct"`
	// The name of the component. By default, the name is the ID of the component.
	//
	// example:
	//
	// component_name
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The type of the component. Valid values:
	//
	// 	- **text**: a text component
	//
	// 	- **image**: an image component
	//
	// 	- **caption**: a caption component
	//
	// example:
	//
	// text
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The display effect for the component. Valid values:
	//
	// 	- **none**
	//
	// 	- **animateH**: horizontal scrolling
	//
	// 	- **animateV**: vertical scrolling
	//
	// example:
	//
	// animateV
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The information about the image component. This parameter is returned only for image components.
	ImageLayerContent *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty" type:"Struct"`
	// The location ID of the component.
	//
	// Each location ID can be assigned to only one component and must be in the RC[Number] format. The values of this parameter are in ascending order, for example, from RC01 to RC12.
	//
	// example:
	//
	// RC01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The information about the text component. This parameter is returned only for text components.
	TextLayerContent *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent `json:"TextLayerContent,omitempty" xml:"TextLayerContent,omitempty" type:"Struct"`
}

func (s DescribeCasterComponentsResponseBodyComponentsComponent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponentsComponent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetCaptionLayerContent() *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	return s.CaptionLayerContent
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetComponentLayer() *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer {
	return s.ComponentLayer
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetEffect() *string {
	return s.Effect
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetImageLayerContent() *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent {
	return s.ImageLayerContent
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetLocationId() *string {
	return s.LocationId
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) GetTextLayerContent() *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	return s.TextLayerContent
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetCaptionLayerContent(v *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.CaptionLayerContent = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetComponentId(v string) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.ComponentId = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetComponentLayer(v *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.ComponentLayer = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetComponentName(v string) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.ComponentName = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetComponentType(v string) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.ComponentType = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetEffect(v string) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.Effect = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetImageLayerContent(v *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.ImageLayerContent = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetLocationId(v string) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.LocationId = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) SetTextLayerContent(v *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) *DescribeCasterComponentsResponseBodyComponentsComponent {
	s.TextLayerContent = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponent) Validate() error {
	if s.CaptionLayerContent != nil {
		if err := s.CaptionLayerContent.Validate(); err != nil {
			return err
		}
	}
	if s.ComponentLayer != nil {
		if err := s.ComponentLayer.Validate(); err != nil {
			return err
		}
	}
	if s.ImageLayerContent != nil {
		if err := s.ImageLayerContent.Validate(); err != nil {
			return err
		}
	}
	if s.TextLayerContent != nil {
		if err := s.TextLayerContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent struct {
	// The color of the text border.
	//
	// Valid values: **0x000000 to 0xffffff**. If the value of this parameter is "", this parameter does not take effect.
	//
	// example:
	//
	// 0x000000
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The normalized value of the width of the text border. The value of this parameter equals the border width divided by the font size.
	//
	// The maximum width of the text border is **16**, even if the border width calculated based on this parameter is greater than **16**.
	//
	// example:
	//
	// 0
	BorderWidthNormalized *float32 `json:"BorderWidthNormalized,omitempty" xml:"BorderWidthNormalized,omitempty"`
	// The color of the text. Valid values: **0x000000 to 0xffffff**.
	//
	// example:
	//
	// 0x000000
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The font of the text, which is specified by the system. Valid values:
	//
	// 	- **KaiTi**
	//
	// 	- **AlibabaPuHuiTi-Regular**
	//
	// 	- **AlibabaPuHuiTi-Bold**
	//
	// 	- **NAlibabaPuHuiTi-Light**
	//
	// 	- **NotoSansHans-Regular**
	//
	// 	- **NotoSansHans-Bold**
	//
	// 	- **NotoSansHans-Light**
	//
	// ****
	//
	// example:
	//
	// KaiTi
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The line spacing, which indicates the interval between every two lines.
	//
	// example:
	//
	// 0
	LineSpaceNormalized *float32 `json:"LineSpaceNormalized,omitempty" xml:"LineSpaceNormalized,omitempty"`
	// The location ID of the component. If the value of the ComponentType parameter is caption, the LocationId parameter indicates the channel ID of the video source that is referenced by the component.
	//
	// example:
	//
	// RV01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The offset between the presentation timestamps (PTS) of the subtitles and the audio.
	//
	// Valid values: **-10000 to 10000**. Default value: **0**.
	//
	// example:
	//
	// 0
	PtsOffset *int32 `json:"PtsOffset,omitempty" xml:"PtsOffset,omitempty"`
	// Indicates whether the source language of the subtitles is displayed. Valid values:
	//
	// 	- **true**: The source language is displayed.
	//
	// 	- **false**: The source language is not displayed.
	//
	// example:
	//
	// false
	ShowSourceLan *bool `json:"ShowSourceLan,omitempty" xml:"ShowSourceLan,omitempty"`
	// The normalized value of the font size of the subtitles. The value of this parameter equals the font size divided by the output height.``
	//
	// The maximum font size of the subtitles is **1,024**, even if the font size calculated based on this parameter is greater than **1,024**. If the value of this parameter is **-1**, this parameter does not take effect.
	//
	// example:
	//
	// 0.1
	SizeNormalized *float32 `json:"SizeNormalized,omitempty" xml:"SizeNormalized,omitempty"`
	// The source language of the audio in the video source. Valid values:
	//
	// 	- **en**: English
	//
	// 	- **cn**: Chinese
	//
	// 	- **es**: Spanish
	//
	// 	- **ru**: Russian
	//
	// example:
	//
	// cn
	SourceLan *string `json:"SourceLan,omitempty" xml:"SourceLan,omitempty"`
	// The target language of the audio in the video source. Valid values:
	//
	// 	- **en**: English
	//
	// 	- **cn**: Chinese
	//
	// 	- **es**: Spanish
	//
	// 	- **ru**: Russian
	//
	// example:
	//
	// cn
	TargetLan *string `json:"TargetLan,omitempty" xml:"TargetLan,omitempty"`
	// The maximum number of words displayed in each line.
	//
	// example:
	//
	// 15
	WordCountPerLine *int32 `json:"WordCountPerLine,omitempty" xml:"WordCountPerLine,omitempty"`
	// The word spacing, which indicates the interval between every two words.
	//
	// example:
	//
	// 0
	WordSpaceNormalized *float32 `json:"WordSpaceNormalized,omitempty" xml:"WordSpaceNormalized,omitempty"`
	// The number of words displayed on the component. The value of this parameter can be specified based on the font size.
	//
	// Valid values: **10 to 50**.
	//
	// example:
	//
	// 35
	WordsCount *int32 `json:"WordsCount,omitempty" xml:"WordsCount,omitempty"`
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetBorderColor() *string {
	return s.BorderColor
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetBorderWidthNormalized() *float32 {
	return s.BorderWidthNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetColor() *string {
	return s.Color
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetFontName() *string {
	return s.FontName
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetLineSpaceNormalized() *float32 {
	return s.LineSpaceNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetLocationId() *string {
	return s.LocationId
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetPtsOffset() *int32 {
	return s.PtsOffset
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetShowSourceLan() *bool {
	return s.ShowSourceLan
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetSizeNormalized() *float32 {
	return s.SizeNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetSourceLan() *string {
	return s.SourceLan
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetTargetLan() *string {
	return s.TargetLan
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetWordCountPerLine() *int32 {
	return s.WordCountPerLine
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetWordSpaceNormalized() *float32 {
	return s.WordSpaceNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) GetWordsCount() *int32 {
	return s.WordsCount
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetBorderColor(v string) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.BorderColor = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetBorderWidthNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.BorderWidthNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetColor(v string) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.Color = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetFontName(v string) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.FontName = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetLineSpaceNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.LineSpaceNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetLocationId(v string) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.LocationId = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetPtsOffset(v int32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.PtsOffset = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetShowSourceLan(v bool) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.ShowSourceLan = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetSizeNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.SizeNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetSourceLan(v string) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.SourceLan = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetTargetLan(v string) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.TargetLan = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetWordCountPerLine(v int32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.WordCountPerLine = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetWordSpaceNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.WordSpaceNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) SetWordsCount(v int32) *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent {
	s.WordsCount = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentCaptionLayerContent) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer struct {
	// The normalized value for heights of the elements in the layer. The widths of the elements are proportionally scaled based on this parameter.
	//
	// If the value of this parameter is **0**, the elements in the layer are not scaled.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The normalized value of the position of the layer, in the format of `[x,y]`. Example: `[0,0]`.
	//
	// >  The values of x and y need to be normalized.
	PositionNormalizeds *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds `json:"PositionNormalizeds,omitempty" xml:"PositionNormalizeds,omitempty" type:"Struct"`
	// The reference coordinates of the layer. Valid values:
	//
	// 	- **topLeft**: the upper-left corner
	//
	// 	- **topRight**: the upper-right corner
	//
	// 	- **bottomLeft**: the lower-left corner
	//
	// 	- **bottomRight**: the lower-right corner
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The transparency of the layer. Valid values: 0 to 255.
	//
	// A value of **0*	- indicates that the layer is completely transparent. A value of **255*	- indicates that the layer is completely opaque.
	//
	// example:
	//
	// 200
	Transparency *int32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
	// The normalized value for widths of the elements in the layer. The heights of the elements are proportionally scaled based on this parameter. If the value of this parameter is **0**, the elements in the layer are not scaled.
	//
	// >  This parameter conflicts with the HeightNormalized parameter. If both of them are specified, only the HeightNormalized parameter takes effect. If only one of them is specified, the latest specified value is used.
	//
	// example:
	//
	// 0.5
	WidthNormalized *float32 `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) GetPositionNormalizeds() *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds {
	return s.PositionNormalizeds
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) GetPositionRefer() *string {
	return s.PositionRefer
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) GetTransparency() *int32 {
	return s.Transparency
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) GetWidthNormalized() *float32 {
	return s.WidthNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) SetHeightNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer {
	s.HeightNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) SetPositionNormalizeds(v *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds) *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer {
	s.PositionNormalizeds = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) SetPositionRefer(v string) *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer {
	s.PositionRefer = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) SetTransparency(v int32) *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer {
	s.Transparency = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) SetWidthNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer {
	s.WidthNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayer) Validate() error {
	if s.PositionNormalizeds != nil {
		if err := s.PositionNormalizeds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds struct {
	Position []*float32 `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds) GetPosition() []*float32 {
	return s.Position
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds) SetPosition(v []*float32) *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds {
	s.Position = v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentComponentLayerPositionNormalizeds) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent struct {
	// The ID of the material from the media library.
	//
	// example:
	//
	// 6cf724c6ebfd4a59b5b3cec6f10d****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent) GetMaterialId() *string {
	return s.MaterialId
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent) SetMaterialId(v string) *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent {
	s.MaterialId = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentImageLayerContent) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent struct {
	// The color of the text border. Valid values: **0x000000 to 0xffffff**. If the value of this parameter is **""**, this parameter does not take effect.
	//
	// example:
	//
	// 0x000000
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The normalized value of the width of the text border. The value of this parameter equals the border width divided by the font size.****
	//
	// The maximum width of the text border is **16**, even if the border width calculated based on this parameter is greater than **16**.
	//
	// example:
	//
	// 0
	BorderWidthNormalized *float32 `json:"BorderWidthNormalized,omitempty" xml:"BorderWidthNormalized,omitempty"`
	// The color of the text. Valid values: **0x000000 to 0xffffff**.
	//
	// example:
	//
	// 0xff0000
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The font of the text, which is specified by the system. Valid values:
	//
	// 	- **KaiTi**
	//
	// 	- **AlibabaPuHuiTi-Regular**
	//
	// 	- **AlibabaPuHuiTi-Bold**
	//
	// 	- **NAlibabaPuHuiTi-Light**
	//
	// 	- **NotoSansHans-Regular**
	//
	// 	- **NotoSansHans-Bold**
	//
	// 	- **NotoSansHans-Light**
	//
	// ****
	//
	// example:
	//
	// KaiTi
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The normalized value of the font size of the text.
	//
	// The value of this parameter equals the font size divided by the output height.***	- The maximum font size of the text is **1,024**, even if the font size calculated based on this parameter is greater than **1,024**. If the value of this parameter is **-1**, this parameter does not take effect.
	//
	// example:
	//
	// 16
	SizeNormalized *float32 `json:"SizeNormalized,omitempty" xml:"SizeNormalized,omitempty"`
	// The content of the text.
	//
	// example:
	//
	// hello world
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GetBorderColor() *string {
	return s.BorderColor
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GetBorderWidthNormalized() *float32 {
	return s.BorderWidthNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GetColor() *string {
	return s.Color
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GetFontName() *string {
	return s.FontName
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GetSizeNormalized() *float32 {
	return s.SizeNormalized
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) GetText() *string {
	return s.Text
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) SetBorderColor(v string) *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	s.BorderColor = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) SetBorderWidthNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	s.BorderWidthNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) SetColor(v string) *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	s.Color = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) SetFontName(v string) *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	s.FontName = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) SetSizeNormalized(v float32) *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	s.SizeNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) SetText(v string) *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent {
	s.Text = &v
	return s
}

func (s *DescribeCasterComponentsResponseBodyComponentsComponentTextLayerContent) Validate() error {
	return dara.Validate(s)
}
