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
	// The CSV watermark control parameter. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	CsvControlShrink *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	// The document watermark parameter that specifies whether the file to be extracted is a screenshot of a document with a background watermark added. The system determines whether to use the extraction logic for document background watermarks based on whether the file to be extracted is an image file. By default, you do not need to configure this parameter. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DocumentIsCapture *bool `json:"DocumentIsCapture,omitempty" xml:"DocumentIsCapture,omitempty"`
	// The URL used to download the file to be extracted. The URL must be accessible over the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The name of the file to be extracted. The system needs to check the file type based on the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-****.pdf
	Filename      *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	IsClientEmbed *bool   `json:"IsClientEmbed,omitempty" xml:"IsClientEmbed,omitempty"`
	// The watermark parameter for videos that specifies whether to use the long video watermark SDK. Default value: false. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// The watermark parameter for long videos that specifies the video speed factor. The value can be a floating-point number or a string. Default value: 1. This parameter indicates the speed at which a watermark is added or the time-stretching rate for videos after a watermark is added.
	//
	// example:
	//
	// 1
	VideoSpeed *string `json:"VideoSpeed,omitempty" xml:"VideoSpeed,omitempty"`
	// The watermark information size. Default value: 32. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. For example, if a 40-bit watermark is used for watermark embedding, you must set this parameter to 40 for watermark extraction.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// The watermark type. Valid values:
	//
	// 	- **PureWebappInvisible**: web page watermark
	//
	// 	- **PureAppInvisible**: app watermark
	//
	// 	- **PureScreenInvisible**: screen watermark
	//
	// 	- **PureDocument**: document watermark
	//
	// 	- **PureImage**: image watermark
	//
	// 	- **PureAudio**: audio watermark
	//
	// 	- **PureVideo**: video watermark
	//
	// 	- **AigcWebappInvisible**: artificial intelligence generated content (AIGC)-based webpage watermark
	//
	// 	- **AigcAppInvisible**: AIGC-based app watermark
	//
	// 	- **AigcScreenInvisible**: AIGC-based screen watermark
	//
	// 	- **AigcDocument**: AIGC-based document watermark
	//
	// 	- **AigcImage**: AIGC-based image watermark
	//
	// 	- **AigcAudio**: AIGC-based audio watermark
	//
	// 	- **AigcVideo**: AIGC-based video watermark
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
