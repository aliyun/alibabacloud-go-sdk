// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptionLayerContent(v string) *AddCasterComponentRequest
	GetCaptionLayerContent() *string
	SetCasterId(v string) *AddCasterComponentRequest
	GetCasterId() *string
	SetComponentLayer(v string) *AddCasterComponentRequest
	GetComponentLayer() *string
	SetComponentName(v string) *AddCasterComponentRequest
	GetComponentName() *string
	SetComponentType(v string) *AddCasterComponentRequest
	GetComponentType() *string
	SetEffect(v string) *AddCasterComponentRequest
	GetEffect() *string
	SetHtmlLayerContent(v string) *AddCasterComponentRequest
	GetHtmlLayerContent() *string
	SetImageLayerContent(v string) *AddCasterComponentRequest
	GetImageLayerContent() *string
	SetLayerOrder(v string) *AddCasterComponentRequest
	GetLayerOrder() *string
	SetLocationId(v string) *AddCasterComponentRequest
	GetLocationId() *string
	SetOwnerId(v int64) *AddCasterComponentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCasterComponentRequest
	GetRegionId() *string
	SetTextLayerContent(v string) *AddCasterComponentRequest
	GetTextLayerContent() *string
}

type AddCasterComponentRequest struct {
	// The properties of the layer element. The properties are described as follows:
	//
	// 	Notice: This parameter is required when ComponentType is set to caption.
	//
	// - **SizeNormalized**: The normalized font size. This value is calculated as font size / output height. The value must be in the `[0,1]` range and accurate to two decimal places. If the font size calculated from the normalized value is greater than **1024**, the font size is set to **1024**.
	//
	// - **BorderWidthNormalized**: The normalized width of the text border. This value is calculated based on the font size: BorderWidth / FontSize. The value must be in the `[0,1]` range and accurate to two decimal places. If the width calculated from the normalized value is greater than **16**, the width is set to **16**. The default value is **0**.
	//
	// - **FontName**: The font name. For valid values, see **Production studio fonts**. The default font is KaiTi.
	//
	// - **BorderColor**: The color of the text border. The value must be a hexadecimal color code that ranges from 0x000000 to 0xffffff. The default value is an empty string (""), which indicates that no border color is set.
	//
	// - **LocationId**: The channel ID of the translation source.
	//
	// - **SourceLan**: The original audio language of the video source. Valid values: en (English), cn (Chinese), es (Spanish), and ru (Russian). The default value is cn.
	//
	// - **TargetLan**: The target audio language for the video source. If you do not set this parameter, only speech recognition is performed. If you set this parameter, the audio is translated. Valid values: en (English), cn (Chinese), es (Spanish), and ru (Russian).
	//
	// - **ShowSourceLan**: Specifies whether to display the source language. Valid values: true and false. The default value is false.
	//
	// - **Truncation**: Specifies whether captions can be truncated. Valid values: true and false. The default value is false.
	//
	// - **SourceLanPerLineWordCount**: The maximum number of words per line for the source language captions. The default value is 20.
	//
	// - **TargetLanPerLineWordCount**: The maximum number of words per line for the target language captions. The default value is 20.
	//
	// - **SourceLanReservePages**: The number of lines to reserve for the source language captions. This parameter takes effect only when Truncation is set to true. The default value is 2.
	//
	// - **TargetLanReservePages**: The number of lines to reserve for the target language captions. This parameter takes effect only when Truncation is set to true. The default value is 2.
	//
	// The value must be a JSON-formatted string. Parameter names must be in upper-camel case.
	//
	// example:
	//
	// {"BorderWidthNormalized":0.01,"SizeNormalized":0.05,"Color":"0x000000","LocationId":"RV01","SourceLan":"cn","FontName":"KaiTi","BorderColor":"0xffffff"}
	CaptionLayerContent *string `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty"`
	// The ID of the production studio.
	//
	// - If you create the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, find the ID in the CasterId parameter of the response.
	//
	// - If you create the production studio in the LIVE console, go to the **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The size, layout, and other information about the component layer. The elements are described as follows:
	//
	// - **HeightNormalized**: The normalized height.
	//
	// - **WidthNormalized**: The normalized width.
	//
	// - **PositionNormalized**: The normalized position of the layer element.
	//
	// - **PositionRefer**: The reference coordinates for the element\\"s position.
	//
	// The value is a JSON-formatted string. Parameter names must be in upper-camel case.
	//
	// This parameter is required.
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
	// The type of component. Valid values:
	//
	// - **text**: A text component. If you set this parameter to text, you must also set the TextLayerContent parameter.
	//
	// - **image**: An image component. If you set this parameter to image, you must also set the ImageLayerContent parameter.
	//
	// - **caption**: A caption component. If you set this parameter to caption, you must also set the CaptionLayerContent parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The display effect of the component. Valid values:
	//
	// - **none*	- (default): No effect.
	//
	// - **animateH**: Scrolls horizontally.
	//
	// - **animateV**: Scrolls vertically.
	//
	// example:
	//
	// animateH
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The configuration of the H5 component.
	//
	// example:
	//
	// {"htmlUrl":http://caster.example.com}
	HtmlLayerContent *string `json:"HtmlLayerContent,omitempty" xml:"HtmlLayerContent,omitempty"`
	// The properties of the layer element. The properties are described as follows:
	//
	// 	Notice:
	//
	// This parameter is required when ComponentType is set to image.
	//
	//
	//
	// MaterialId: The ID of the media asset. The name that you specify when you upload a media asset is used as the ID of the media asset.
	//
	// The value must be a JSON-formatted string. Parameter names must be in upper-camel case.
	//
	// example:
	//
	// {"MaterialId":"6cf724c6ebfd4a59b5b3cec6f10d****"}
	ImageLayerContent *string `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty"`
	// The layer order of the component.
	//
	// - cover: The component is in the foreground.
	//
	// - background: The component is in the background.
	//
	// example:
	//
	// cover
	LayerOrder *string `json:"LayerOrder,omitempty" xml:"LayerOrder,omitempty"`
	// Specifies the position of the component. Each position can hold only one component. The format must be RC01 to RC99.
	//
	// > If the component type is caption, this parameter specifies the location of the referenced video source.
	//
	// This parameter is required.
	//
	// example:
	//
	// RC01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The properties of the layer element. The properties are described as follows:
	//
	// 	Notice: This parameter is required only when ComponentType is set to text.
	//
	// - **SizeNormalized**: The normalized font size. This value is calculated as font size / output height. The value must be in the `[0,1]` range. If the font size calculated from the normalized value is greater than 1024, the font size is set to 1024.
	//
	// - **BorderWidthNormalized**: The normalized width of the text border. This value is calculated based on the font size: BorderWidth / FontSize. The value must be in the `[0,1]` range. If the width calculated from the normalized value is greater than 16, the width is set to 16. The default value is 0.
	//
	// - **FontName**: The font name. For valid values, see **Production studio fonts**. The default font is KaiTi.
	//
	// - **BorderColor**: The color of the text border. The value must be a hexadecimal color code that ranges from 0x000000 to 0xffffff. The default value is an empty string (""), which indicates that no border color is set.
	//
	// - **Text**: The text content. The default value is an empty string ("").
	//
	// - **Color**: The text color. The default value is 0xff0000, which represents red.
	//
	// The value must be a JSON-formatted string. Parameter names must be in upper-camel case.
	//
	// example:
	//
	// {"BorderWidthNormalized":"1","SizeNormalized":"0.2","Color":"0x000000","FontName":"KaiTi","BorderColor":"0x000000","Text":"hello world!"}
	TextLayerContent *string `json:"TextLayerContent,omitempty" xml:"TextLayerContent,omitempty"`
}

func (s AddCasterComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasterComponentRequest) GoString() string {
	return s.String()
}

func (s *AddCasterComponentRequest) GetCaptionLayerContent() *string {
	return s.CaptionLayerContent
}

func (s *AddCasterComponentRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddCasterComponentRequest) GetComponentLayer() *string {
	return s.ComponentLayer
}

func (s *AddCasterComponentRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *AddCasterComponentRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *AddCasterComponentRequest) GetEffect() *string {
	return s.Effect
}

func (s *AddCasterComponentRequest) GetHtmlLayerContent() *string {
	return s.HtmlLayerContent
}

func (s *AddCasterComponentRequest) GetImageLayerContent() *string {
	return s.ImageLayerContent
}

func (s *AddCasterComponentRequest) GetLayerOrder() *string {
	return s.LayerOrder
}

func (s *AddCasterComponentRequest) GetLocationId() *string {
	return s.LocationId
}

func (s *AddCasterComponentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCasterComponentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCasterComponentRequest) GetTextLayerContent() *string {
	return s.TextLayerContent
}

func (s *AddCasterComponentRequest) SetCaptionLayerContent(v string) *AddCasterComponentRequest {
	s.CaptionLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) SetCasterId(v string) *AddCasterComponentRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterComponentRequest) SetComponentLayer(v string) *AddCasterComponentRequest {
	s.ComponentLayer = &v
	return s
}

func (s *AddCasterComponentRequest) SetComponentName(v string) *AddCasterComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *AddCasterComponentRequest) SetComponentType(v string) *AddCasterComponentRequest {
	s.ComponentType = &v
	return s
}

func (s *AddCasterComponentRequest) SetEffect(v string) *AddCasterComponentRequest {
	s.Effect = &v
	return s
}

func (s *AddCasterComponentRequest) SetHtmlLayerContent(v string) *AddCasterComponentRequest {
	s.HtmlLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) SetImageLayerContent(v string) *AddCasterComponentRequest {
	s.ImageLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) SetLayerOrder(v string) *AddCasterComponentRequest {
	s.LayerOrder = &v
	return s
}

func (s *AddCasterComponentRequest) SetLocationId(v string) *AddCasterComponentRequest {
	s.LocationId = &v
	return s
}

func (s *AddCasterComponentRequest) SetOwnerId(v int64) *AddCasterComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCasterComponentRequest) SetRegionId(v string) *AddCasterComponentRequest {
	s.RegionId = &v
	return s
}

func (s *AddCasterComponentRequest) SetTextLayerContent(v string) *AddCasterComponentRequest {
	s.TextLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) Validate() error {
	return dara.Validate(s)
}
