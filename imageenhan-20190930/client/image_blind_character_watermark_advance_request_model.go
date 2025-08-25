// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iImageBlindCharacterWatermarkAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionType(v string) *ImageBlindCharacterWatermarkAdvanceRequest
	GetFunctionType() *string
	SetOriginImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest
	GetOriginImageURLObject() io.Reader
	SetOutputFileType(v string) *ImageBlindCharacterWatermarkAdvanceRequest
	GetOutputFileType() *string
	SetQualityFactor(v int32) *ImageBlindCharacterWatermarkAdvanceRequest
	GetQualityFactor() *int32
	SetText(v string) *ImageBlindCharacterWatermarkAdvanceRequest
	GetText() *string
	SetWatermarkImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest
	GetWatermarkImageURLObject() io.Reader
}

type ImageBlindCharacterWatermarkAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// encode_text
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxxxx.jpg
	OriginImageURLObject io.Reader `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
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
	WatermarkImageURLObject io.Reader `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindCharacterWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) GetOriginImageURLObject() io.Reader {
	return s.OriginImageURLObject
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) GetOutputFileType() *string {
	return s.OutputFileType
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) GetQualityFactor() *int32 {
	return s.QualityFactor
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) GetText() *string {
	return s.Text
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) GetWatermarkImageURLObject() io.Reader {
	return s.WatermarkImageURLObject
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetFunctionType(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetOriginImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.OriginImageURLObject = v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetOutputFileType(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetQualityFactor(v int32) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.QualityFactor = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetText(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.Text = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetWatermarkImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.WatermarkImageURLObject = v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
