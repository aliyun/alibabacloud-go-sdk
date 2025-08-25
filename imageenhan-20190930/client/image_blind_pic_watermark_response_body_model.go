// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBlindPicWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ImageBlindPicWatermarkResponseBodyData) *ImageBlindPicWatermarkResponseBody
	GetData() *ImageBlindPicWatermarkResponseBodyData
	SetRequestId(v string) *ImageBlindPicWatermarkResponseBody
	GetRequestId() *string
}

type ImageBlindPicWatermarkResponseBody struct {
	Data *ImageBlindPicWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DE7869E4-0ACE-4C02-8B98-540B49F49205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageBlindPicWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindPicWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponseBody) GetData() *ImageBlindPicWatermarkResponseBodyData {
	return s.Data
}

func (s *ImageBlindPicWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageBlindPicWatermarkResponseBody) SetData(v *ImageBlindPicWatermarkResponseBodyData) *ImageBlindPicWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *ImageBlindPicWatermarkResponseBody) SetRequestId(v string) *ImageBlindPicWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageBlindPicWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImageBlindPicWatermarkResponseBodyData struct {
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty%2F2020_03_04%2F61f544a1a5004c88a2bf29452db494e9.jpeg?OSSAccessKeyId=LTAI4Fmdm1gQonFLrghJ****&Expires=158340****&Signature=Heet1ivG0xFP3YlO6usvd0pmrH****
	LogoURL *string `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty%2F2020_03_04%2F61f544a1a5004c88a2bf29452db494e9.jpeg?OSSAccessKeyId=LTAI4Fmdm1gQonFLrghJ****&Expires=158340****&Signature=Heet1ivG0xFP3YlO6usvd0pmrH****
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindPicWatermarkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindPicWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponseBodyData) GetLogoURL() *string {
	return s.LogoURL
}

func (s *ImageBlindPicWatermarkResponseBodyData) GetWatermarkImageURL() *string {
	return s.WatermarkImageURL
}

func (s *ImageBlindPicWatermarkResponseBodyData) SetLogoURL(v string) *ImageBlindPicWatermarkResponseBodyData {
	s.LogoURL = &v
	return s
}

func (s *ImageBlindPicWatermarkResponseBodyData) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkResponseBodyData {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindPicWatermarkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
