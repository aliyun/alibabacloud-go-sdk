// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iImageBlindPicWatermarkAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionType(v string) *ImageBlindPicWatermarkAdvanceRequest
	GetFunctionType() *string
	SetLogoURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest
	GetLogoURLObject() io.Reader
	SetOriginImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest
	GetOriginImageURLObject() io.Reader
	SetOutputFileType(v string) *ImageBlindPicWatermarkAdvanceRequest
	GetOutputFileType() *string
	SetQualityFactor(v int32) *ImageBlindPicWatermarkAdvanceRequest
	GetQualityFactor() *int32
	SetWatermarkImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest
	GetWatermarkImageURLObject() io.Reader
}

type ImageBlindPicWatermarkAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// encode_pic
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxxxx.jpg
	LogoURLObject io.Reader `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
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
	QualityFactor *int32 `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	// example:
	//
	// https://viapi-doc.oss-cn-shanghai.aliyuncs.com/imageenhan/xxxxx.jpg
	WatermarkImageURLObject io.Reader `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindPicWatermarkAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindPicWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkAdvanceRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ImageBlindPicWatermarkAdvanceRequest) GetLogoURLObject() io.Reader {
	return s.LogoURLObject
}

func (s *ImageBlindPicWatermarkAdvanceRequest) GetOriginImageURLObject() io.Reader {
	return s.OriginImageURLObject
}

func (s *ImageBlindPicWatermarkAdvanceRequest) GetOutputFileType() *string {
	return s.OutputFileType
}

func (s *ImageBlindPicWatermarkAdvanceRequest) GetQualityFactor() *int32 {
	return s.QualityFactor
}

func (s *ImageBlindPicWatermarkAdvanceRequest) GetWatermarkImageURLObject() io.Reader {
	return s.WatermarkImageURLObject
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetFunctionType(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetLogoURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.LogoURLObject = v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetOriginImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.OriginImageURLObject = v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetOutputFileType(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetQualityFactor(v int32) *ImageBlindPicWatermarkAdvanceRequest {
	s.QualityFactor = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetWatermarkImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.WatermarkImageURLObject = v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
