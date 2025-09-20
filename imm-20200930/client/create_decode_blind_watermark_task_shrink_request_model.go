// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDecodeBlindWatermarkTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageQuality(v int32) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetImageQuality() *int32
	SetModel(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetModel() *string
	SetNotificationShrink(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetNotificationShrink() *string
	SetOriginalImageURI(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetOriginalImageURI() *string
	SetProjectName(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetSourceURI() *string
	SetStrengthLevel(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetStrengthLevel() *string
	SetTargetURI(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetTargetURI() *string
	SetWatermarkType(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest
	GetWatermarkType() *string
}

type CreateDecodeBlindWatermarkTaskShrinkRequest struct {
	// The quality of the output image. This parameter is also available in the earlier DecodeBlindWatermark operation.
	//
	// Higher image quality indicates a larger image size and higher watermark resolution quality.
	//
	// example:
	//
	// 90
	ImageQuality *int32 `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	// The watermark algorithm model. This parameter is also available in the earlier DecodeBlindWatermark operation. Valid values: FFT, FFT_FULL, DWT, and DWT_IBG. Default value: FFT.
	//
	// If this parameter is left empty, the CreateDecodeBlindWatermarkTask operation is called. Otherwise, the earlier DecodeBlindWatermark operation is called.
	//
	// example:
	//
	// FFT
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The OSS URI of the image before the blind watermark is added. This parameter is also available in the earlier DecodeBlindWatermark operation.
	//
	// Do not specify this parameter when you set the Model parameter to DWT or DWT_IBG.
	//
	// Specify the OSS URI in the `oss://<bucket>/<object>` format, where `<bucket>` is the name of the bucket in the same region as the current project and `<object>` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://imm-test/testcases/watermarktestbefore.jpg
	OriginalImageURI *string `json:"OriginalImageURI,omitempty" xml:"OriginalImageURI,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// >  The project specified in the request must match the one in the EncodeBlindWatermark request to encode the blind watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the image.
	//
	// Specify the OSS URI in the `oss://<bucket>/<object>` format, where `<bucket>` is the name of the bucket in the same region as the current project and `<object>` is the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://target/sampleobject.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The watermark strength level. The higher the strength level, the more resistant the watermarked image is to attacks, but the more the image is distorted. Valid values: low, medium, and high. Default value: low.
	//
	// example:
	//
	// low
	StrengthLevel *string `json:"StrengthLevel,omitempty" xml:"StrengthLevel,omitempty"`
	// The OSS URI of the output image. This parameter is also available in the earlier DecodeBlindWatermark operation.
	//
	// Specify the OSS URI in the `oss://<bucket>/<object>` format, where `<bucket>` is the name of the bucket in the same region as the current project and `<object>` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://target/targetobject.jpg
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The type of the watermark. Valid value: text.
	//
	// No image watermarks are supported.
	//
	// example:
	//
	// text
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s CreateDecodeBlindWatermarkTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDecodeBlindWatermarkTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetImageQuality() *int32 {
	return s.ImageQuality
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetOriginalImageURI() *string {
	return s.OriginalImageURI
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetStrengthLevel() *string {
	return s.StrengthLevel
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetImageQuality(v int32) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.ImageQuality = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetModel(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.Model = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetNotificationShrink(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetOriginalImageURI(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.OriginalImageURI = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetProjectName(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetSourceURI(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetStrengthLevel(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.StrengthLevel = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetTargetURI(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) SetWatermarkType(v string) *CreateDecodeBlindWatermarkTaskShrinkRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
