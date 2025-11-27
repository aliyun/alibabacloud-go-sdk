// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDecodeBlindWatermarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageQuality(v int32) *CreateDecodeBlindWatermarkTaskRequest
	GetImageQuality() *int32
	SetModel(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetModel() *string
	SetNotification(v *Notification) *CreateDecodeBlindWatermarkTaskRequest
	GetNotification() *Notification
	SetOriginalImageURI(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetOriginalImageURI() *string
	SetProjectName(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetSourceURI() *string
	SetStrengthLevel(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetStrengthLevel() *string
	SetTargetURI(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetTargetURI() *string
	SetWatermarkType(v string) *CreateDecodeBlindWatermarkTaskRequest
	GetWatermarkType() *string
}

type CreateDecodeBlindWatermarkTaskRequest struct {
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
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	// The level of watermark extraction. A higher level indicates a longer time and a higher quality. Valid values:
	//
	// 	- low
	//
	// 	- medium
	//
	// 	- high
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

func (s CreateDecodeBlindWatermarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDecodeBlindWatermarkTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetImageQuality() *int32 {
	return s.ImageQuality
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetModel() *string {
	return s.Model
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetOriginalImageURI() *string {
	return s.OriginalImageURI
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetStrengthLevel() *string {
	return s.StrengthLevel
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateDecodeBlindWatermarkTaskRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetImageQuality(v int32) *CreateDecodeBlindWatermarkTaskRequest {
	s.ImageQuality = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetModel(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.Model = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetNotification(v *Notification) *CreateDecodeBlindWatermarkTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetOriginalImageURI(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.OriginalImageURI = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetProjectName(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetSourceURI(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetStrengthLevel(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.StrengthLevel = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetTargetURI(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) SetWatermarkType(v string) *CreateDecodeBlindWatermarkTaskRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskRequest) Validate() error {
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}
