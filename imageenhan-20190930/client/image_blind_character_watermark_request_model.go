// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBlindCharacterWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionType(v string) *ImageBlindCharacterWatermarkRequest
	GetFunctionType() *string
	SetOriginImageURL(v string) *ImageBlindCharacterWatermarkRequest
	GetOriginImageURL() *string
	SetOutputFileType(v string) *ImageBlindCharacterWatermarkRequest
	GetOutputFileType() *string
	SetQualityFactor(v int32) *ImageBlindCharacterWatermarkRequest
	GetQualityFactor() *int32
	SetText(v string) *ImageBlindCharacterWatermarkRequest
	GetText() *string
	SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkRequest
	GetWatermarkImageURL() *string
}

type ImageBlindCharacterWatermarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// encode_text
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
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
	QualityFactor *int32  `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	Text          *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// https://viapi-doc.oss-cn-shanghai.aliyuncs.com/imageenhan/xxxxx.jpg
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindCharacterWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ImageBlindCharacterWatermarkRequest) GetOriginImageURL() *string {
	return s.OriginImageURL
}

func (s *ImageBlindCharacterWatermarkRequest) GetOutputFileType() *string {
	return s.OutputFileType
}

func (s *ImageBlindCharacterWatermarkRequest) GetQualityFactor() *int32 {
	return s.QualityFactor
}

func (s *ImageBlindCharacterWatermarkRequest) GetText() *string {
	return s.Text
}

func (s *ImageBlindCharacterWatermarkRequest) GetWatermarkImageURL() *string {
	return s.WatermarkImageURL
}

func (s *ImageBlindCharacterWatermarkRequest) SetFunctionType(v string) *ImageBlindCharacterWatermarkRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetOriginImageURL(v string) *ImageBlindCharacterWatermarkRequest {
	s.OriginImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetOutputFileType(v string) *ImageBlindCharacterWatermarkRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetQualityFactor(v int32) *ImageBlindCharacterWatermarkRequest {
	s.QualityFactor = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetText(v string) *ImageBlindCharacterWatermarkRequest {
	s.Text = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkRequest {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
