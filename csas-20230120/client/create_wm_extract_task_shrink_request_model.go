// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmExtractTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsvControlShrink(v string) *CreateWmExtractTaskShrinkRequest
	GetCsvControlShrink() *string
	SetDocumentIsCapture(v bool) *CreateWmExtractTaskShrinkRequest
	GetDocumentIsCapture() *bool
	SetFileUrl(v string) *CreateWmExtractTaskShrinkRequest
	GetFileUrl() *string
	SetFilename(v string) *CreateWmExtractTaskShrinkRequest
	GetFilename() *string
	SetImageExtractParamsOpenApiShrink(v string) *CreateWmExtractTaskShrinkRequest
	GetImageExtractParamsOpenApiShrink() *string
	SetIsClientEmbed(v bool) *CreateWmExtractTaskShrinkRequest
	GetIsClientEmbed() *bool
	SetVideoIsLong(v bool) *CreateWmExtractTaskShrinkRequest
	GetVideoIsLong() *bool
	SetVideoSpeed(v string) *CreateWmExtractTaskShrinkRequest
	GetVideoSpeed() *string
	SetWmInfoSize(v int64) *CreateWmExtractTaskShrinkRequest
	GetWmInfoSize() *int64
	SetWmType(v string) *CreateWmExtractTaskShrinkRequest
	GetWmType() *string
}

type CreateWmExtractTaskShrinkRequest struct {
	// CSV watermark control parameters. These parameters must be consistent with those used during embedding. Otherwise, the extraction fails.
	CsvControlShrink *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	// Document watermark parameter. Specifies whether the file from which the watermark is to be extracted is a screenshot of a document with a background watermark. The service determines whether to use the document background watermark extraction logic based on whether the file is an image file. Therefore, you do not need to set this parameter by default. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	DocumentIsCapture *bool `json:"DocumentIsCapture,omitempty" xml:"DocumentIsCapture,omitempty"`
	// The URL used to download the file to be fetched. The URL must support public network access.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The name of the file from which the watermark is to be extracted. The backend determines and verifies the file type based on the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Image extraction parameters.
	ImageExtractParamsOpenApiShrink *string `json:"ImageExtractParamsOpenApi,omitempty" xml:"ImageExtractParamsOpenApi,omitempty"`
	// Audio watermark parameter. Specifies whether the watermark is embedded by using the client SDK. Default value: false. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	IsClientEmbed *bool `json:"IsClientEmbed,omitempty" xml:"IsClientEmbed,omitempty"`
	// Video watermark parameter. Specifies whether to use the long video watermark SDK. Default value: false. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// Long video watermark parameter. The video playback speed, represented as a floating-point string. Default value: 1. This value represents the playback speed used when adding the watermark, or the timeline stretching speed after the watermark is added.
	//
	// example:
	//
	// 1
	VideoSpeed *string `json:"VideoSpeed,omitempty" xml:"VideoSpeed,omitempty"`
	// The bit width of the watermark information capacity. Default value: 32. This parameter must be consistent during embedding and extraction. For example, if a 40-bit SDK is used for embedding, you must also set this value to 40 for extraction.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// The watermark type. Valid values:
	//
	// - **PureWebappInvisible**: Web page watermark.
	//
	// - **PureAppInvisible**: App watermark.
	//
	// - **PureScreenInvisible**: Screen watermark.
	//
	// - **PureDocument**: Document watermark.
	//
	// - **PureImage**: Image watermark.
	//
	// - **PureAudio**: Audio watermark.
	//
	// - **PureVideo**: Video watermark.
	//
	// - **AigcWebappInvisible**: AIGC web page watermark.
	//
	// - **AigcAppInvisible**: AIGC app watermark.
	//
	// - **AigcScreenInvisible**: AIGC screen watermark.
	//
	// - **AigcDocument**: AIGC document watermark.
	//
	// - **AigcImage**: AIGC image watermark.
	//
	// - **AigcAudio**: AIGC audio watermark.
	//
	// - **AigcVideo**: AIGC video watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmExtractTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmExtractTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskShrinkRequest) GetCsvControlShrink() *string {
	return s.CsvControlShrink
}

func (s *CreateWmExtractTaskShrinkRequest) GetDocumentIsCapture() *bool {
	return s.DocumentIsCapture
}

func (s *CreateWmExtractTaskShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWmExtractTaskShrinkRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateWmExtractTaskShrinkRequest) GetImageExtractParamsOpenApiShrink() *string {
	return s.ImageExtractParamsOpenApiShrink
}

func (s *CreateWmExtractTaskShrinkRequest) GetIsClientEmbed() *bool {
	return s.IsClientEmbed
}

func (s *CreateWmExtractTaskShrinkRequest) GetVideoIsLong() *bool {
	return s.VideoIsLong
}

func (s *CreateWmExtractTaskShrinkRequest) GetVideoSpeed() *string {
	return s.VideoSpeed
}

func (s *CreateWmExtractTaskShrinkRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmExtractTaskShrinkRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmExtractTaskShrinkRequest) SetCsvControlShrink(v string) *CreateWmExtractTaskShrinkRequest {
	s.CsvControlShrink = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetDocumentIsCapture(v bool) *CreateWmExtractTaskShrinkRequest {
	s.DocumentIsCapture = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetFileUrl(v string) *CreateWmExtractTaskShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetFilename(v string) *CreateWmExtractTaskShrinkRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetImageExtractParamsOpenApiShrink(v string) *CreateWmExtractTaskShrinkRequest {
	s.ImageExtractParamsOpenApiShrink = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetIsClientEmbed(v bool) *CreateWmExtractTaskShrinkRequest {
	s.IsClientEmbed = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetVideoIsLong(v bool) *CreateWmExtractTaskShrinkRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetVideoSpeed(v string) *CreateWmExtractTaskShrinkRequest {
	s.VideoSpeed = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetWmInfoSize(v int64) *CreateWmExtractTaskShrinkRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetWmType(v string) *CreateWmExtractTaskShrinkRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
