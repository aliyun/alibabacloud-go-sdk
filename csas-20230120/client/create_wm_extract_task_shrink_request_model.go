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
	// The CSV watermark control parameters. These must be consistent with the parameters used during embedding. Otherwise, extraction fails.
	CsvControlShrink *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	// The document watermark parameter that specifies whether the file to be extracted is a screenshot of a document with a background watermark. The service determines whether to use the document background watermark extraction logic based on whether the file is an image file. Therefore, this parameter does not need to be set by default. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	DocumentIsCapture *bool `json:"DocumentIsCapture,omitempty" xml:"DocumentIsCapture,omitempty"`
	// The URL used to download the file from which the watermark is to be fetched. The URL must be accessible over the public network access.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The name of the file from which the watermark is to be extracted. The backend determines and validates the file type based on the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// The image extraction parameters.
	ImageExtractParamsOpenApiShrink *string `json:"ImageExtractParamsOpenApi,omitempty" xml:"ImageExtractParamsOpenApi,omitempty"`
	// The audio watermark parameter that specifies whether the watermark was embedded by the client SDK. Default value: false. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	IsClientEmbed *bool `json:"IsClientEmbed,omitempty" xml:"IsClientEmbed,omitempty"`
	// The video watermark parameter that specifies whether to use the long video watermark SDK. Default value: false. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// The long video watermark parameter that specifies the video playback speed as a floating-point string. Default value: 1, which indicates the playback speed used when the watermark was added, or the speed at which the video timeline was stretched after the watermark was added.
	//
	// example:
	//
	// 1
	VideoSpeed *string `json:"VideoSpeed,omitempty" xml:"VideoSpeed,omitempty"`
	// The bit width of the watermark information capacity. Default value: 32. This parameter must be consistent between embedding and extraction. For example, if the 40-bit SDK was used for embedding, set this value to 40 for extraction.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// The watermark type. Valid values:
	//
	// - **PureWebappInvisible**: web page watermark.
	//
	// - **PureAppInvisible**: app watermark.
	//
	// - **PureScreenInvisible**: screen watermark.
	//
	// - **PureDocument**: document watermark.
	//
	// - **PureImage**: image watermark.
	//
	// - **PureAudio**: audio watermark.
	//
	// - **PureVideo**: video watermark.
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
