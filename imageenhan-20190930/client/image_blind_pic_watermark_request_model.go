// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBlindPicWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionType(v string) *ImageBlindPicWatermarkRequest
	GetFunctionType() *string
	SetLogoURL(v string) *ImageBlindPicWatermarkRequest
	GetLogoURL() *string
	SetOriginImageURL(v string) *ImageBlindPicWatermarkRequest
	GetOriginImageURL() *string
	SetOutputFileType(v string) *ImageBlindPicWatermarkRequest
	GetOutputFileType() *string
	SetQualityFactor(v int32) *ImageBlindPicWatermarkRequest
	GetQualityFactor() *int32
	SetWatermarkImageURL(v string) *ImageBlindPicWatermarkRequest
	GetWatermarkImageURL() *string
}

type ImageBlindPicWatermarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// encode_pic
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxxxx.jpg
	LogoURL *string `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxxxx.jpg
	OriginImageURL *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
	// example:
	//
	// jpg
	OutputFileType *string `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	// example:
	//
	// 100
	QualityFactor *int32 `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	// example:
	//
	// https://viapi-doc.oss-cn-shanghai.aliyuncs.com/imageenhan/xxxxx.jpg
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindPicWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindPicWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ImageBlindPicWatermarkRequest) GetLogoURL() *string {
	return s.LogoURL
}

func (s *ImageBlindPicWatermarkRequest) GetOriginImageURL() *string {
	return s.OriginImageURL
}

func (s *ImageBlindPicWatermarkRequest) GetOutputFileType() *string {
	return s.OutputFileType
}

func (s *ImageBlindPicWatermarkRequest) GetQualityFactor() *int32 {
	return s.QualityFactor
}

func (s *ImageBlindPicWatermarkRequest) GetWatermarkImageURL() *string {
	return s.WatermarkImageURL
}

func (s *ImageBlindPicWatermarkRequest) SetFunctionType(v string) *ImageBlindPicWatermarkRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetLogoURL(v string) *ImageBlindPicWatermarkRequest {
	s.LogoURL = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetOriginImageURL(v string) *ImageBlindPicWatermarkRequest {
	s.OriginImageURL = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetOutputFileType(v string) *ImageBlindPicWatermarkRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetQualityFactor(v int32) *ImageBlindPicWatermarkRequest {
	s.QualityFactor = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkRequest {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
