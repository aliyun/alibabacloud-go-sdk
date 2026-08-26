// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptionLayerContent(v string) *ModifyCasterComponentRequest
	GetCaptionLayerContent() *string
	SetCasterId(v string) *ModifyCasterComponentRequest
	GetCasterId() *string
	SetComponentId(v string) *ModifyCasterComponentRequest
	GetComponentId() *string
	SetComponentLayer(v string) *ModifyCasterComponentRequest
	GetComponentLayer() *string
	SetComponentName(v string) *ModifyCasterComponentRequest
	GetComponentName() *string
	SetComponentType(v string) *ModifyCasterComponentRequest
	GetComponentType() *string
	SetEffect(v string) *ModifyCasterComponentRequest
	GetEffect() *string
	SetImageLayerContent(v string) *ModifyCasterComponentRequest
	GetImageLayerContent() *string
	SetOwnerId(v int64) *ModifyCasterComponentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCasterComponentRequest
	GetRegionId() *string
	SetTextLayerContent(v string) *ModifyCasterComponentRequest
	GetTextLayerContent() *string
}

type ModifyCasterComponentRequest struct {
	// The properties of the caption layer. The value is a JSON string. The following properties are supported:
	//
	// 	Notice:
	//
	// This parameter is required if you set ComponentType to caption.
	//
	//
	//
	// - **SizeNormalized**: The normalized font size. The font size is calculated using the formula: font_size/output_height. The value must be in the range of `[0,1]`. If the calculated font size is greater than 1024, the value 1024 is used.
	//
	// - **BorderWidthNormalized**: The normalized width of the text border. The normalized width is calculated based on the font size using the formula: BorderWidth/FontSize. The value must be in the range of `[0,1]`. If the calculated value is greater than 16, the value 16 is used. Default value: 0.
	//
	// - **FontName**: The font name. For more information about valid values, see **Production studio fonts**. Default value: KaiTi.
	//
	// - **BorderColor**: The color of the text border. Valid values are from 0x000000 to 0xffffff. The default value is an empty string, which indicates that this parameter is not used.
	//
	// - **LocationId**: The channel ID of the translation source.
	//
	// - **SourceLan**: The source language of the audio in the video source. Valid values are en (English), cn (Chinese), es (Spanish), and ru (Russian). Default value: cn.
	//
	// - **TargetLan**: The target language for translation. If you do not set this parameter, only speech recognition is performed. If you set this parameter, translation is also performed. Valid values are en (English), cn (Chinese), es (Spanish), and ru (Russian).
	//
	// - **ShowSourceLan**: Specifies whether to display the source language. Valid values are true (display) and false (do not display). Default value: false.
	//
	// - **Truncation**: Specifies whether to truncate the caption. Valid values are true (truncate) and false (do not truncate). Default value: false.
	//
	// - **SourceLanPerLineWordCount**: The number of words per line for the source language. This parameter takes effect only if Truncation is set to true. Default value: 20.
	//
	// - **TargetLanPerLineWordCount**: The number of words per line for the target language. This parameter takes effect only if Truncation is set to true. Default value: 20.
	//
	// example:
	//
	// {"BorderWidthNormalized":0.01,"SizeNormalized":0.05,"Color":"0x000000","LocationId":"RV01","SourceLan":"cn","FontName":"KaiTi","BorderColor":"0xffffff"}
	CaptionLayerContent *string `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty"`
	// The ID of the production studio.
	//
	// - The ID is returned after you call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation.
	//
	// - If you create a production studio in the LIVE console, go to the **LIVE*	- > **Production Studio*	- > **Cloud Production Studio*	- page to find the ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The component ID. The ID is returned after you call the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The size and layout of the layer. The value is a JSON string. The following properties are supported:
	//
	// - **HeightNormalized**: The normalized height.
	//
	// - **WidthNormalized**: The normalized width.
	//
	// - **PositionNormalized**: The normalized position of the layer.
	//
	// - **PositionRefer**: The reference point for the position of the layer.
	//
	// example:
	//
	// {"HeightNormalized":"1","PositionRefer":"topRight","WidthNormalized":"0","PositionNormalized":["0.1","0.2"]}
	ComponentLayer *string `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty"`
	// The name of the component. The default value is the component ID.
	//
	// example:
	//
	// text01
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The type of the component. Valid values:
	//
	// - **text**: A text component. The TextLayerContent parameter is required only if you set ComponentType to text.
	//
	// - **image**: An image component. The ImageLayerContent parameter is required only if you set ComponentType to image.
	//
	// - **caption**: A translation caption component. The CaptionLayerContent parameter is required only if you set ComponentType to caption.
	//
	// example:
	//
	// text
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The display effect of the component. Valid values:
	//
	// - **none*	- (default): no effect.
	//
	// - **animateH**: horizontal scroll.
	//
	// - **animateV**: vertical scroll.
	//
	// example:
	//
	// animateV
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The properties of the image layer. The value is a JSON string.
	//
	// 	Notice:
	//
	// This parameter is required if you set ComponentType to image.
	//
	//
	//
	// MaterialId is the ID of the material in the media asset library.
	//
	// example:
	//
	// {"MaterialId":"6cf724c6ebfd4a59b5b3cec6f10d5ecf"}
	ImageLayerContent *string `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The properties of the text layer. The value is a JSON string. The following properties are supported:
	//
	// 	Notice:
	//
	// This parameter is required if you set ComponentType to text.
	//
	//
	//
	// - **SizeNormalized**: The normalized font size. The font size is calculated using the formula: font_size/output_height. The value must be in the range of `[0,1]`. If the calculated font size is greater than 1024, the value 1024 is used.
	//
	// - **BorderWidthNormalized**: The normalized width of the text border. The normalized width is calculated based on the font size using the formula: BorderWidth/FontSize. The value must be in the range of `[0,1]`. If the calculated value is greater than 16, the value 16 is used. Default value: 0.
	//
	// - **FontName**: The font name. For more information about valid values, see **Production studio fonts**. Default value: KaiTi.
	//
	// - **BorderColor**: The color of the text border. Valid values are from 0x000000 to 0xffffff. The default value is an empty string, which indicates that this parameter is not used.
	//
	// - **Text**: The text content. The default value is an empty string.
	//
	// - **Color**: The color of the text. Default value: 0xff0000, which is red.
	//
	// example:
	//
	// {"BorderWidthNormalized":"1","SizeNormalized":"0.2","Color":"0x000000","FontName":"KaiTi","BorderColor":"0x000000","Text":"hello world!"}
	TextLayerContent *string `json:"TextLayerContent,omitempty" xml:"TextLayerContent,omitempty"`
}

func (s ModifyCasterComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterComponentRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterComponentRequest) GetCaptionLayerContent() *string {
	return s.CaptionLayerContent
}

func (s *ModifyCasterComponentRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterComponentRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *ModifyCasterComponentRequest) GetComponentLayer() *string {
	return s.ComponentLayer
}

func (s *ModifyCasterComponentRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *ModifyCasterComponentRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *ModifyCasterComponentRequest) GetEffect() *string {
	return s.Effect
}

func (s *ModifyCasterComponentRequest) GetImageLayerContent() *string {
	return s.ImageLayerContent
}

func (s *ModifyCasterComponentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCasterComponentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCasterComponentRequest) GetTextLayerContent() *string {
	return s.TextLayerContent
}

func (s *ModifyCasterComponentRequest) SetCaptionLayerContent(v string) *ModifyCasterComponentRequest {
	s.CaptionLayerContent = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetCasterId(v string) *ModifyCasterComponentRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentId(v string) *ModifyCasterComponentRequest {
	s.ComponentId = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentLayer(v string) *ModifyCasterComponentRequest {
	s.ComponentLayer = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentName(v string) *ModifyCasterComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentType(v string) *ModifyCasterComponentRequest {
	s.ComponentType = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetEffect(v string) *ModifyCasterComponentRequest {
	s.Effect = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetImageLayerContent(v string) *ModifyCasterComponentRequest {
	s.ImageLayerContent = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetOwnerId(v int64) *ModifyCasterComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetRegionId(v string) *ModifyCasterComponentRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetTextLayerContent(v string) *ModifyCasterComponentRequest {
	s.TextLayerContent = &v
	return s
}

func (s *ModifyCasterComponentRequest) Validate() error {
	return dara.Validate(s)
}
