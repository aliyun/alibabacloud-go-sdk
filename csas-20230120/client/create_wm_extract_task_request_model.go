// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmExtractTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsvControl(v *CreateWmExtractTaskRequestCsvControl) *CreateWmExtractTaskRequest
	GetCsvControl() *CreateWmExtractTaskRequestCsvControl
	SetDocumentIsCapture(v bool) *CreateWmExtractTaskRequest
	GetDocumentIsCapture() *bool
	SetFileUrl(v string) *CreateWmExtractTaskRequest
	GetFileUrl() *string
	SetFilename(v string) *CreateWmExtractTaskRequest
	GetFilename() *string
	SetIsClientEmbed(v bool) *CreateWmExtractTaskRequest
	GetIsClientEmbed() *bool
	SetVideoIsLong(v bool) *CreateWmExtractTaskRequest
	GetVideoIsLong() *bool
	SetVideoSpeed(v string) *CreateWmExtractTaskRequest
	GetVideoSpeed() *string
	SetWmInfoSize(v int64) *CreateWmExtractTaskRequest
	GetWmInfoSize() *int64
	SetWmType(v string) *CreateWmExtractTaskRequest
	GetWmType() *string
}

type CreateWmExtractTaskRequest struct {
	// The CSV watermark control parameter. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	CsvControl *CreateWmExtractTaskRequestCsvControl `json:"CsvControl,omitempty" xml:"CsvControl,omitempty" type:"Struct"`
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

func (s CreateWmExtractTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskRequest) GetCsvControl() *CreateWmExtractTaskRequestCsvControl {
	return s.CsvControl
}

func (s *CreateWmExtractTaskRequest) GetDocumentIsCapture() *bool {
	return s.DocumentIsCapture
}

func (s *CreateWmExtractTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWmExtractTaskRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateWmExtractTaskRequest) GetIsClientEmbed() *bool {
	return s.IsClientEmbed
}

func (s *CreateWmExtractTaskRequest) GetVideoIsLong() *bool {
	return s.VideoIsLong
}

func (s *CreateWmExtractTaskRequest) GetVideoSpeed() *string {
	return s.VideoSpeed
}

func (s *CreateWmExtractTaskRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmExtractTaskRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmExtractTaskRequest) SetCsvControl(v *CreateWmExtractTaskRequestCsvControl) *CreateWmExtractTaskRequest {
	s.CsvControl = v
	return s
}

func (s *CreateWmExtractTaskRequest) SetDocumentIsCapture(v bool) *CreateWmExtractTaskRequest {
	s.DocumentIsCapture = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetFileUrl(v string) *CreateWmExtractTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetFilename(v string) *CreateWmExtractTaskRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetIsClientEmbed(v bool) *CreateWmExtractTaskRequest {
	s.IsClientEmbed = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetVideoIsLong(v bool) *CreateWmExtractTaskRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetVideoSpeed(v string) *CreateWmExtractTaskRequest {
	s.VideoSpeed = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetWmInfoSize(v int64) *CreateWmExtractTaskRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetWmType(v string) *CreateWmExtractTaskRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmExtractTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateWmExtractTaskRequestCsvControl struct {
	// The timestamp watermark parameter that specifies how much information a single timestamp holds. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	//
	// example:
	//
	// 1
	EmbedBitsNumberInEachTime *int64 `json:"EmbedBitsNumberInEachTime,omitempty" xml:"EmbedBitsNumberInEachTime,omitempty"`
	// The lossy embedding control parameter that specifies columns to be modified You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	//
	// example:
	//
	// 1
	EmbedColumn *int64 `json:"EmbedColumn,omitempty" xml:"EmbedColumn,omitempty"`
	// The lossy embedding control parameter that specifies the modification precision. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	//
	// example:
	//
	// 0
	EmbedPrecision *int64 `json:"EmbedPrecision,omitempty" xml:"EmbedPrecision,omitempty"`
	// The timestamp watermark parameter that specifies the embedding position of the timestamp watermarks. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	//
	// example:
	//
	// Min
	EmbedTimePosition *string `json:"EmbedTimePosition,omitempty" xml:"EmbedTimePosition,omitempty"`
	// The CSV watermark embedding method. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	//
	// example:
	//
	// lossless_row_shift_embed
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The timestamp watermark parameter that specifies the timestamp format. You must keep the value of this parameter consistent for watermark embedding and watermark extraction. Otherwise, the extraction fails.
	//
	// example:
	//
	// Year-Mon-Day Hour:Min:Sec.MilSec
	TimeFormat *string `json:"TimeFormat,omitempty" xml:"TimeFormat,omitempty"`
}

func (s CreateWmExtractTaskRequestCsvControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmExtractTaskRequestCsvControl) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskRequestCsvControl) GetEmbedBitsNumberInEachTime() *int64 {
	return s.EmbedBitsNumberInEachTime
}

func (s *CreateWmExtractTaskRequestCsvControl) GetEmbedColumn() *int64 {
	return s.EmbedColumn
}

func (s *CreateWmExtractTaskRequestCsvControl) GetEmbedPrecision() *int64 {
	return s.EmbedPrecision
}

func (s *CreateWmExtractTaskRequestCsvControl) GetEmbedTimePosition() *string {
	return s.EmbedTimePosition
}

func (s *CreateWmExtractTaskRequestCsvControl) GetMethod() *string {
	return s.Method
}

func (s *CreateWmExtractTaskRequestCsvControl) GetTimeFormat() *string {
	return s.TimeFormat
}

func (s *CreateWmExtractTaskRequestCsvControl) SetEmbedBitsNumberInEachTime(v int64) *CreateWmExtractTaskRequestCsvControl {
	s.EmbedBitsNumberInEachTime = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetEmbedColumn(v int64) *CreateWmExtractTaskRequestCsvControl {
	s.EmbedColumn = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetEmbedPrecision(v int64) *CreateWmExtractTaskRequestCsvControl {
	s.EmbedPrecision = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetEmbedTimePosition(v string) *CreateWmExtractTaskRequestCsvControl {
	s.EmbedTimePosition = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetMethod(v string) *CreateWmExtractTaskRequestCsvControl {
	s.Method = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetTimeFormat(v string) *CreateWmExtractTaskRequestCsvControl {
	s.TimeFormat = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) Validate() error {
	return dara.Validate(s)
}
