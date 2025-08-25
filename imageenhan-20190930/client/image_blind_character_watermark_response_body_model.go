// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBlindCharacterWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ImageBlindCharacterWatermarkResponseBodyData) *ImageBlindCharacterWatermarkResponseBody
	GetData() *ImageBlindCharacterWatermarkResponseBodyData
	SetRequestId(v string) *ImageBlindCharacterWatermarkResponseBody
	GetRequestId() *string
}

type ImageBlindCharacterWatermarkResponseBody struct {
	Data *ImageBlindCharacterWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2457E1ED-9C76-4386-B9A2-7E41B7D6E849
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageBlindCharacterWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponseBody) GetData() *ImageBlindCharacterWatermarkResponseBodyData {
	return s.Data
}

func (s *ImageBlindCharacterWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageBlindCharacterWatermarkResponseBody) SetData(v *ImageBlindCharacterWatermarkResponseBodyData) *ImageBlindCharacterWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBody) SetRequestId(v string) *ImageBlindCharacterWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImageBlindCharacterWatermarkResponseBodyData struct {
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty%2F2020_03_04%2F61f544a1a5004c88a2bf29452db494e9.jpeg?OSSAccessKeyId=LTAI4Fmdm1gQonFLrghJ****&Expires=158340****&Signature=Heet1ivG0xFP3YlO6usvd0pmrH****
	TextImageURL *string `json:"TextImageURL,omitempty" xml:"TextImageURL,omitempty"`
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty%2F2020_03_04%2F61f544a1a5004c88a2bf29452db494e9.jpeg?OSSAccessKeyId=LTAI4Fmdm1gQonFLrghJ****&Expires=158340****&Signature=Heet1ivG0xFP3YlO6usvd0pmrH****
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) GetTextImageURL() *string {
	return s.TextImageURL
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) GetWatermarkImageURL() *string {
	return s.WatermarkImageURL
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) SetTextImageURL(v string) *ImageBlindCharacterWatermarkResponseBodyData {
	s.TextImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkResponseBodyData {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
