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
	// A parameter from the earlier DecodeBlindWatermark API. It specifies the quality of the output image. The default value is 90. The value must be in the range of 70 to 100.
	//
	// A higher quality results in a larger image size and better watermark parsing quality.
	//
	// example:
	//
	// 90
	ImageQuality *int32 `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	// A parameter from the earlier DecodeBlindWatermark API. It specifies the watermark algorithm model. Valid values include FFT, FFT_FULL, DWT, and DWT_IBG. The default value is FFT.
	//
	// If this parameter is left empty, the new version of the blind watermarking feature is used. Otherwise, the earlier version is used.
	//
	// example:
	//
	// FFT
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// A parameter from the earlier DecodeBlindWatermark API. It specifies the OSS URI of the image before the blind watermark was added.
	//
	// This parameter is not required when Model is set to DWT or DWT_IBG.
	//
	// The OSS URI must be in the `oss://<bucket>/<object>` format. `<bucket>` is the name of the OSS bucket that is in the same region as the current project. `<object>` is the full path of the file, including the file name extension.
	//
	// example:
	//
	// oss://imm-test/testcases/watermarktestbefore.jpg
	OriginalImageURI *string `json:"OriginalImageURI,omitempty" xml:"OriginalImageURI,omitempty"`
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// 	Notice: The project name must be the same as the one used to add the blind watermark with the [EncodeBlindWatermark](https://help.aliyun.com/document_detail/2743655.html) operation. Otherwise, the watermark cannot be extracted.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Object Storage Service (OSS) URI of the image.
	//
	// The OSS URI must be in the `oss://<bucket>/<object>` format. `<bucket>` is the name of the OSS bucket that is in the same region as the current project. `<object>` is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://target/sampleobject.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The watermark extraction level, which controls the extraction precision. A higher level indicates a longer processing time and a better extraction effect. Valid values:
	//
	// - low
	//
	// - medium
	//
	// - high
	//
	// example:
	//
	// low
	StrengthLevel *string `json:"StrengthLevel,omitempty" xml:"StrengthLevel,omitempty"`
	// A parameter from the earlier DecodeBlindWatermark API. It specifies the OSS URI where the image is saved after the blind watermark is parsed.
	//
	// The OSS URI must be in the `oss://<bucket>/<object>` format. `<bucket>` is the name of the OSS bucket that is in the same region as the current project. `<object>` is the full path of the file, including the file name extension.
	//
	// example:
	//
	// oss://target/targetobject.jpg
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The type of the embedded watermark. Valid value: text
	//
	// (Image watermarks are not supported. Therefore, this parameter can only be set to text.)
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
