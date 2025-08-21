// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmEmbedTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetVideoBitrate(v string) *CreateWmEmbedTaskShrinkRequest
	GetVideoBitrate() *string
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
	CsvControlShrink      *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	DocumentControlShrink *string `json:"DocumentControl,omitempty" xml:"DocumentControl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/abc****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc****.pdf
	Filename           *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	ImageControlShrink *string `json:"ImageControl,omitempty" xml:"ImageControl,omitempty"`
	// example:
	//
	// 95
	ImageEmbedJpegQuality *int64 `json:"ImageEmbedJpegQuality,omitempty" xml:"ImageEmbedJpegQuality,omitempty"`
	// example:
	//
	// 2
	ImageEmbedLevel *int64 `json:"ImageEmbedLevel,omitempty" xml:"ImageEmbedLevel,omitempty"`
	// example:
	//
	// 3000k
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// example:
	//
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
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

func (s *CreateWmEmbedTaskShrinkRequest) GetVideoBitrate() *string {
	return s.VideoBitrate
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

func (s *CreateWmEmbedTaskShrinkRequest) SetVideoBitrate(v string) *CreateWmEmbedTaskShrinkRequest {
	s.VideoBitrate = &v
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
