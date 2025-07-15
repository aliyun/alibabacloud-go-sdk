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
	// The information about the subtitle component. The value must be a JSON string. This parameter contains the following fields:
	//
	// >  This parameter is required if you set ComponentType to caption.
	//
	// 	- **SizeNormalized**: the normalized value of the font size. The value of this field equals the font size divided by the output height. Valid values: `0 to 1`. The maximum font size is 1,024, even if the font size calculated based on this field is greater than 1,024.
	//
	// 	- **BorderWidthNormalized**: the normalized value of the border width. The value of this field equals the border width divided by the font size. Valid values: `0 to 1`. Default value: 0. The maximum border width is 16, even if the border width calculated based on this field is greater than 16.
	//
	// 	- **FontName**: the font name. Default value: KaiTi. For more information about the valid values, see **Fonts used in a production studio**.
	//
	// 	- **BorderColor**: the color of the text border. Valid values: 0x000000 to 0xffffff. By default, this parameter is left empty. In this case, the color of the text border is transparent.
	//
	// 	- **LocationId**: the channel ID of the source subtitles.
	//
	// 	- **SourceLan**: the source language of the subtitles in the video. Valid values: en (English), cn (Chinese), es (Spanish), and ru (Russian). Default value: cn.
	//
	// 	- **TargetLan**: the target language of the subtitles in the video. If you do not specify this field, speech recognition is used. If you specify this field, translation is used. Valid values: en (English), cn (Chinese), es (Spanish), and ru (Russian).
	//
	// 	- **ShowSourceLan**: specifies whether to display the source language. A value of true specifies that the source language is displayed. A value of false specifies that the source language is not displayed. Default value: false.
	//
	// 	- **Truncation**: specifies whether to allow subtitle truncation. A value of true specifies that the subtitles can be truncated. A value of false specifies that the subtitles cannot be truncated. Default value: false.
	//
	// 	- **SourceLanPerLineWordCount**: the number of words displayed in each line of the source language. This field takes effect only if you set Truncation to true. Default value: 20.
	//
	// 	- **TargetLanPerLineWordCount**: the number of words displayed in each line of the target language. This field takes effect only if you set Truncation to true. Default value: 20.
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
	// The ID of the component. If the component was added by calling the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation, check the value of the response parameter ComponentId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The information about the component layer, such as the size and layout, The value must be a JSON string. This parameter contains the following fields:
	//
	// 	- **HeightNormalized**: the normalized value of the height of the component layer.
	//
	// 	- **WidthNormalized**: the normalized value of the width of the component layer.
	//
	// 	- **PositionNormalized**: the normalized value of the position of the component layer.
	//
	// 	- **PositionRefer**: the reference coordinates of the component layer.
	//
	// example:
	//
	// {"HeightNormalized":"1","PositionRefer":"topRight","WidthNormalized":"0","PositionNormalized":["0.1","0.2"]}
	ComponentLayer *string `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty"`
	// The name of the component. By default, the name is the ID of the component.
	//
	// example:
	//
	// text01
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The type of the component. Valid values:
	//
	// 	- **text**: text component. The TextLayerContent parameter is required if you set ComponentType to text.
	//
	// 	- **image**: image component. The ImageLayerContent parameter is required if you set ComponentType to image.
	//
	// 	- **caption**: subtitle component. The CaptionLayerContent parameter is required if you set ComponentType to caption.
	//
	// example:
	//
	// text
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The display effect for the component. Valid values:
	//
	// 	- **none*	- (default)
	//
	// 	- **animateH**: horizontal scrolling
	//
	// 	- **animateV**: vertical scrolling
	//
	// example:
	//
	// animateV
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The information about the image component. The value must be a JSON string.
	//
	// >  This parameter is required if you set ComponentType to image.
	//
	// The MaterialId field specifies the ID of the material from the media asset library.
	//
	// example:
	//
	// {"MaterialId":"6cf724c6ebfd4a59b5b3cec6f10d5ecf"}
	ImageLayerContent *string `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the text component. The value must be a JSON string. This parameter contains the following fields:
	//
	// >  This parameter is required if you set ComponentType to text.
	//
	// 	- **SizeNormalized**: the normalized value of the font size. The value of this field equals the font size divided by the output height. Valid values: `0 to 1`. The maximum font size is 1,024, even if the font size calculated based on this field is greater than 1,024.
	//
	// 	- **BorderWidthNormalized**: the normalized value of the border width. The value of this field equals the border width divided by the font size. Valid values: `0 to 1`. Default value: 0. The maximum border width is 16, even if the border width calculated based on this field is greater than 16.
	//
	// 	- **FontName**: the font name. Default value: KaiTi. For more information about the valid values, see **Fonts used in a production studio**.
	//
	// 	- **BorderColor**: the color of the text border. Valid values: 0x000000 to 0xffffff. By default, this parameter is left empty. In this case, the color of the text border is transparent.
	//
	// 	- **Text**: the content of the text. By default, this parameter is left empty. In this case, the text contains no content.
	//
	// 	- **Color**: the color of the text. The default value is 0xff0000, which indicates that the text is in red.
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
