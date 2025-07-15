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
	// The information about the caption layer. This parameter contains the following fields:
	//
	// >  This parameter is required when the ComponentType parameter is set to caption.
	//
	// 	- **SizeNormalized**: the normalized font size. The font size is set to font_size/output_height. The value range is `[0,1]` and accurate to two decimal places. If the font size calculated by the system based on the normalization method is greater than **1024**, **1024*	- is used.
	//
	// 	- **BorderWidthNormalized**: the normalized value of the text border width, which is calculated based on the size of the text, namely "BorderWidth/FontSize". The value range is `[0,1]` and accurate to two decimal places. If the value calculated according to the normalization method exceeds **16**, **16*	- is used. The default value is **0**.
	//
	// 	- **FontName**: the font name. For more information about the value, see **Font description**. The default font name is KaiTi.
	//
	// 	- **BorderColor**: the color of the text border. Valid values: 0x000000 to 0xffffff. By default, this parameter is not set. In this case, the color of the text border is transparent.
	//
	// 	- **LocationId**: the channel ID of the source subtitle.
	//
	// 	- **SourceLan**: the source language of the audio in the source video. Valid values: en, cn, es, and ru, which indicate English, Chinese, Spanish, and Russian respectively. Default value: cn.
	//
	// 	- **TargetLan**: the target audio language in the source video. If you do not specify this field, speech recognition is used. If you specify this field, translation is used. Valid values: en, cn, es, and ru, which indicate English, Chinese, Spanish, and Russian respectively. Default value: cn.
	//
	// 	- **ShowSourceLan**: specifies whether to display the source language. Valid values: true: displays the source language. false: hides the source language. Default value: false.
	//
	// 	- **Truncation**: specifies whether to allow caption truncation. Valid values: true: specifies that the caption can be truncated. false: specifies that the caption cannot be truncated. Default value: false.
	//
	// 	- **SourceLanPerLineWordCount**: the number of words displayed in each line when the subtitle is in the source language. Default value: 20.
	//
	// 	- **TargetLanPerLineWordCount**: the number of words displayed in each line when the subtitle is in the destination language. Default value: 20.
	//
	// 	- **SourceLanReservePages**: the number of lines reserved when the subtitle is in the source language. This field takes effect only when the Truncation field is set to true. Default value: 2.
	//
	// 	- **TargetLanReservePages**: the number of lines reserved when the subtitle is in the destination language. This field takes effect only when the Truncation field is set to true. Default value: 2.
	//
	// The value is a JSON string. Use upper camel case for field names.
	//
	// example:
	//
	// {"BorderWidthNormalized":0.01,"SizeNormalized":0.05,"Color":"0x000000","LocationId":"RV01","SourceLan":"cn","FontName":"KaiTi","BorderColor":"0xffffff"}
	CaptionLayerContent *string `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty"`
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The information about the component layer, such as the size and layout. This parameter contains the following fields:
	//
	// 	- **HeightNormalized**: the normalized value of heights for the elements in the component layer
	//
	// 	- **WidthNormalized**: the normalized value of widths for the elements in the component layer
	//
	// 	- **PositionNormalized**: the normalized value of the coordinates of the component layer
	//
	// 	- **PositionRefer**: the reference coordinates of the component layer.
	//
	// The value is a JSON string. Use upper camel case for field names.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"HeightNormalized":"1","PositionRefer":"topRight","WidthNormalized":"0","PositionNormalized":["0.1","0.2"]}
	ComponentLayer *string `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty"`
	// The component name. By default, the component name is the component ID.
	//
	// example:
	//
	// text01
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The component type. Valid values:
	//
	// 	- **text**: a text component. If you set ComponentType to text, you must also specify TextLayerContent.
	//
	// 	- **image**: an image component. If you set ComponentType to image, you must also specify ImageLayerContent.
	//
	// 	- **caption**: a caption component. If you set ComponentType to caption, you must also specify CaptionLayerContent.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The effect of the component. Valid values:
	//
	// 	- **none (default)**
	//
	// 	- **animateH**: horizontal scrolling
	//
	// 	- **animateV**: vertical scrolling
	//
	// example:
	//
	// animateH
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The information about the HTML5 layer.
	//
	// example:
	//
	// {"htmlUrl":http://caster.example.com}
	HtmlLayerContent *string `json:"HtmlLayerContent,omitempty" xml:"HtmlLayerContent,omitempty"`
	// The information about the image layer. This parameter contains the following fields:
	//
	// >  This parameter is required when the ComponentType parameter is set to image.
	//
	// MaterialId: the ID of the asset from the media asset library. The name that you set when you upload an asset is the ID of the asset.
	//
	// The value is a JSON string. Use upper camel case for field names.
	//
	// example:
	//
	// {"MaterialId":"6cf724c6ebfd4a59b5b3cec6f10d****"}
	ImageLayerContent *string `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty"`
	// The layer stacking order of the component. Valid values:
	//
	// 	- cover
	//
	// 	- background
	//
	// example:
	//
	// cover
	LayerOrder *string `json:"LayerOrder,omitempty" xml:"LayerOrder,omitempty"`
	// The location ID of the component. Each location ID can be assigned to only one component and must be in the RC[Number] format. The values specified by this parameter must be in ascending order, such as RC01 to RC99.
	//
	// >  If the ComponentType parameter is set to caption, the LocationId parameter specifies the location ID of the video source referenced by the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// RC01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the text layer. This parameter contains the following fields:
	//
	// >  This parameter is available and required only when the ComponentType parameter is set to text.
	//
	// 	- **SizeNormalized**: the normalized font size. The font size is set to font_size/output_height. The value range is `[0,1]`. If the font size calculated by the system based on the normalization method is greater than 1024, 1024 is used.
	//
	// 	- **BorderWidthNormalized**: the normalized value of the text border width. The normalized value is calculated based on the size of the text, that is, "BorderWidth/FontSize". The value range is `[0,1]`. If the value calculated based on the normalization method exceeds 16, 16 is used. The default value is 0.
	//
	// 	- **FontName**: the font name. For more information about the value, see **Font description**. The default value is KaiTi.
	//
	// 	- **BorderColor**: the color of the text border. Valid values: 0x000000 to 0xffffff. By default, this parameter is not set. In this case, the color of the text border is transparent.
	//
	// 	- **Text**: the content of the text. By default, this parameter is not set. In this case, the text contains no content.
	//
	// 	- **Color**: the color of the text. The default value is 0xff0000, which indicates that the text is in red.
	//
	// The value is a JSON string. Use upper camel case for field names.
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
