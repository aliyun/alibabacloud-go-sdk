// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmEmbedTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioControlShrink(v string) *CreateWmEmbedTaskShrinkRequest
	GetAudioControlShrink() *string
	SetCsvControlShrink(v string) *CreateWmEmbedTaskShrinkRequest
	GetCsvControlShrink() *string
	SetDocumentControlShrink(v string) *CreateWmEmbedTaskShrinkRequest
	GetDocumentControlShrink() *string
	SetFileUrl(v string) *CreateWmEmbedTaskShrinkRequest
	GetFileUrl() *string
	SetFilename(v string) *CreateWmEmbedTaskShrinkRequest
	GetFilename() *string
	SetImageControlShrink(v string) *CreateWmEmbedTaskShrinkRequest
	GetImageControlShrink() *string
	SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskShrinkRequest
	GetImageEmbedJpegQuality() *int64
	SetImageEmbedLevel(v int64) *CreateWmEmbedTaskShrinkRequest
	GetImageEmbedLevel() *int64
	SetInvisibleEnable(v bool) *CreateWmEmbedTaskShrinkRequest
	GetInvisibleEnable() *bool
	SetVideoBitrate(v string) *CreateWmEmbedTaskShrinkRequest
	GetVideoBitrate() *string
	SetVideoControlShrink(v string) *CreateWmEmbedTaskShrinkRequest
	GetVideoControlShrink() *string
	SetVideoIsLong(v bool) *CreateWmEmbedTaskShrinkRequest
	GetVideoIsLong() *bool
	SetWmInfoBytesB64(v string) *CreateWmEmbedTaskShrinkRequest
	GetWmInfoBytesB64() *string
	SetWmInfoSize(v int64) *CreateWmEmbedTaskShrinkRequest
	GetWmInfoSize() *int64
	SetWmInfoUint(v string) *CreateWmEmbedTaskShrinkRequest
	GetWmInfoUint() *string
	SetWmType(v string) *CreateWmEmbedTaskShrinkRequest
	GetWmType() *string
}

type CreateWmEmbedTaskShrinkRequest struct {
	// The audio control parameters.
	AudioControlShrink *string `json:"AudioControl,omitempty" xml:"AudioControl,omitempty"`
	// The CSV watermark embedding control parameters.
	CsvControlShrink *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	// The document watermark control parameters.
	DocumentControlShrink *string `json:"DocumentControl,omitempty" xml:"DocumentControl,omitempty"`
	// The URL for downloading the file to be embedded. The URL must be active for public network access.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/abc****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The name of the file to be embedded. The backend validates the file type based on the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// The image watermark control parameters.
	ImageControlShrink *string `json:"ImageControl,omitempty" xml:"ImageControl,omitempty"`
	// The image watermark parameter that specifies the expected JPEG compression quality factor of the output image. Default value: 95. Valid values: 1 to 100.
	//
	// example:
	//
	// 95
	ImageEmbedJpegQuality *int64 `json:"ImageEmbedJpegQuality,omitempty" xml:"ImageEmbedJpegQuality,omitempty"`
	// The image watermark parameter. A larger value indicates higher robustness but lower visual quality. Default value: 2. Valid values: 0 to 4.
	//
	// example:
	//
	// 2
	ImageEmbedLevel *int64 `json:"ImageEmbedLevel,omitempty" xml:"ImageEmbedLevel,omitempty"`
	// Specifies whether to enable invisible watermark embedding. Default value: true.
	InvisibleEnable *bool `json:"InvisibleEnable,omitempty" xml:"InvisibleEnable,omitempty"`
	// The short video watermark parameter that specifies the video bitrate. By default, the video bitrate is automatically obtained. You can use this parameter to forcibly specify the bitrate used during extraction. Typically, you do not need to set this parameter.
	//
	// example:
	//
	// 3000k
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The video control parameters.
	VideoControlShrink *string `json:"VideoControl,omitempty" xml:"VideoControl,omitempty"`
	// Video watermark parameter. Specifies whether to use the long video watermark SDK. Valid values:
	//
	// - **true**: The long video watermark SDK is used.
	//
	// - **false**: The long video watermark SDK is not used.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// The watermark information in Base64-encoded string format. If this parameter is set, WmInfoUint cannot be set.
	//
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// The bit width of the watermark information capacity. Default value: 32. This parameter must be consistent between embedding and extraction. For example, if the 40-bit SDK is used for embedding, set this parameter to 40 during extraction as well.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// The watermark information in decimal number format. If this parameter is set, WmInfoBytesB64 cannot be set.
	//
	// example:
	//
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// The watermark type. Valid values:
	//
	// - **PureDocument**: document watermark.
	//
	// - **PureImage**: image watermark.
	//
	// - **PureAudio**: audio watermark.
	//
	// - **PureVideo**: video watermark.
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

func (s CreateWmEmbedTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskShrinkRequest) GetAudioControlShrink() *string {
	return s.AudioControlShrink
}

func (s *CreateWmEmbedTaskShrinkRequest) GetCsvControlShrink() *string {
	return s.CsvControlShrink
}

func (s *CreateWmEmbedTaskShrinkRequest) GetDocumentControlShrink() *string {
	return s.DocumentControlShrink
}

func (s *CreateWmEmbedTaskShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWmEmbedTaskShrinkRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateWmEmbedTaskShrinkRequest) GetImageControlShrink() *string {
	return s.ImageControlShrink
}

func (s *CreateWmEmbedTaskShrinkRequest) GetImageEmbedJpegQuality() *int64 {
	return s.ImageEmbedJpegQuality
}

func (s *CreateWmEmbedTaskShrinkRequest) GetImageEmbedLevel() *int64 {
	return s.ImageEmbedLevel
}

func (s *CreateWmEmbedTaskShrinkRequest) GetInvisibleEnable() *bool {
	return s.InvisibleEnable
}

func (s *CreateWmEmbedTaskShrinkRequest) GetVideoBitrate() *string {
	return s.VideoBitrate
}

func (s *CreateWmEmbedTaskShrinkRequest) GetVideoControlShrink() *string {
	return s.VideoControlShrink
}

func (s *CreateWmEmbedTaskShrinkRequest) GetVideoIsLong() *bool {
	return s.VideoIsLong
}

func (s *CreateWmEmbedTaskShrinkRequest) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *CreateWmEmbedTaskShrinkRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmEmbedTaskShrinkRequest) GetWmInfoUint() *string {
	return s.WmInfoUint
}

func (s *CreateWmEmbedTaskShrinkRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmEmbedTaskShrinkRequest) SetAudioControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.AudioControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetCsvControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.CsvControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetDocumentControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.DocumentControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetFileUrl(v string) *CreateWmEmbedTaskShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetFilename(v string) *CreateWmEmbedTaskShrinkRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetImageControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.ImageControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskShrinkRequest {
	s.ImageEmbedJpegQuality = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetImageEmbedLevel(v int64) *CreateWmEmbedTaskShrinkRequest {
	s.ImageEmbedLevel = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetInvisibleEnable(v bool) *CreateWmEmbedTaskShrinkRequest {
	s.InvisibleEnable = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetVideoBitrate(v string) *CreateWmEmbedTaskShrinkRequest {
	s.VideoBitrate = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetVideoControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.VideoControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetVideoIsLong(v bool) *CreateWmEmbedTaskShrinkRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmInfoBytesB64(v string) *CreateWmEmbedTaskShrinkRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmInfoSize(v int64) *CreateWmEmbedTaskShrinkRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmInfoUint(v string) *CreateWmEmbedTaskShrinkRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmType(v string) *CreateWmEmbedTaskShrinkRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
