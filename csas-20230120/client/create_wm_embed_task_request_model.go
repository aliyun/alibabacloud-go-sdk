// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmEmbedTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioControl(v *CreateWmEmbedTaskRequestAudioControl) *CreateWmEmbedTaskRequest
	GetAudioControl() *CreateWmEmbedTaskRequestAudioControl
	SetCsvControl(v *CreateWmEmbedTaskRequestCsvControl) *CreateWmEmbedTaskRequest
	GetCsvControl() *CreateWmEmbedTaskRequestCsvControl
	SetDocumentControl(v *CreateWmEmbedTaskRequestDocumentControl) *CreateWmEmbedTaskRequest
	GetDocumentControl() *CreateWmEmbedTaskRequestDocumentControl
	SetFileUrl(v string) *CreateWmEmbedTaskRequest
	GetFileUrl() *string
	SetFilename(v string) *CreateWmEmbedTaskRequest
	GetFilename() *string
	SetImageControl(v *CreateWmEmbedTaskRequestImageControl) *CreateWmEmbedTaskRequest
	GetImageControl() *CreateWmEmbedTaskRequestImageControl
	SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskRequest
	GetImageEmbedJpegQuality() *int64
	SetImageEmbedLevel(v int64) *CreateWmEmbedTaskRequest
	GetImageEmbedLevel() *int64
	SetInvisibleEnable(v bool) *CreateWmEmbedTaskRequest
	GetInvisibleEnable() *bool
	SetVideoBitrate(v string) *CreateWmEmbedTaskRequest
	GetVideoBitrate() *string
	SetVideoControl(v *CreateWmEmbedTaskRequestVideoControl) *CreateWmEmbedTaskRequest
	GetVideoControl() *CreateWmEmbedTaskRequestVideoControl
	SetVideoIsLong(v bool) *CreateWmEmbedTaskRequest
	GetVideoIsLong() *bool
	SetWmInfoBytesB64(v string) *CreateWmEmbedTaskRequest
	GetWmInfoBytesB64() *string
	SetWmInfoSize(v int64) *CreateWmEmbedTaskRequest
	GetWmInfoSize() *int64
	SetWmInfoUint(v string) *CreateWmEmbedTaskRequest
	GetWmInfoUint() *string
	SetWmType(v string) *CreateWmEmbedTaskRequest
	GetWmType() *string
}

type CreateWmEmbedTaskRequest struct {
	// The audio control parameters.
	AudioControl *CreateWmEmbedTaskRequestAudioControl `json:"AudioControl,omitempty" xml:"AudioControl,omitempty" type:"Struct"`
	// The CSV watermark embedding control parameters.
	CsvControl *CreateWmEmbedTaskRequestCsvControl `json:"CsvControl,omitempty" xml:"CsvControl,omitempty" type:"Struct"`
	// The document watermark control parameters.
	DocumentControl *CreateWmEmbedTaskRequestDocumentControl `json:"DocumentControl,omitempty" xml:"DocumentControl,omitempty" type:"Struct"`
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
	ImageControl *CreateWmEmbedTaskRequestImageControl `json:"ImageControl,omitempty" xml:"ImageControl,omitempty" type:"Struct"`
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
	VideoControl *CreateWmEmbedTaskRequestVideoControl `json:"VideoControl,omitempty" xml:"VideoControl,omitempty" type:"Struct"`
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

func (s CreateWmEmbedTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequest) GetAudioControl() *CreateWmEmbedTaskRequestAudioControl {
	return s.AudioControl
}

func (s *CreateWmEmbedTaskRequest) GetCsvControl() *CreateWmEmbedTaskRequestCsvControl {
	return s.CsvControl
}

func (s *CreateWmEmbedTaskRequest) GetDocumentControl() *CreateWmEmbedTaskRequestDocumentControl {
	return s.DocumentControl
}

func (s *CreateWmEmbedTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWmEmbedTaskRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateWmEmbedTaskRequest) GetImageControl() *CreateWmEmbedTaskRequestImageControl {
	return s.ImageControl
}

func (s *CreateWmEmbedTaskRequest) GetImageEmbedJpegQuality() *int64 {
	return s.ImageEmbedJpegQuality
}

func (s *CreateWmEmbedTaskRequest) GetImageEmbedLevel() *int64 {
	return s.ImageEmbedLevel
}

func (s *CreateWmEmbedTaskRequest) GetInvisibleEnable() *bool {
	return s.InvisibleEnable
}

func (s *CreateWmEmbedTaskRequest) GetVideoBitrate() *string {
	return s.VideoBitrate
}

func (s *CreateWmEmbedTaskRequest) GetVideoControl() *CreateWmEmbedTaskRequestVideoControl {
	return s.VideoControl
}

func (s *CreateWmEmbedTaskRequest) GetVideoIsLong() *bool {
	return s.VideoIsLong
}

func (s *CreateWmEmbedTaskRequest) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *CreateWmEmbedTaskRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmEmbedTaskRequest) GetWmInfoUint() *string {
	return s.WmInfoUint
}

func (s *CreateWmEmbedTaskRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmEmbedTaskRequest) SetAudioControl(v *CreateWmEmbedTaskRequestAudioControl) *CreateWmEmbedTaskRequest {
	s.AudioControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetCsvControl(v *CreateWmEmbedTaskRequestCsvControl) *CreateWmEmbedTaskRequest {
	s.CsvControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetDocumentControl(v *CreateWmEmbedTaskRequestDocumentControl) *CreateWmEmbedTaskRequest {
	s.DocumentControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetFileUrl(v string) *CreateWmEmbedTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetFilename(v string) *CreateWmEmbedTaskRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageControl(v *CreateWmEmbedTaskRequestImageControl) *CreateWmEmbedTaskRequest {
	s.ImageControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskRequest {
	s.ImageEmbedJpegQuality = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageEmbedLevel(v int64) *CreateWmEmbedTaskRequest {
	s.ImageEmbedLevel = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetInvisibleEnable(v bool) *CreateWmEmbedTaskRequest {
	s.InvisibleEnable = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoBitrate(v string) *CreateWmEmbedTaskRequest {
	s.VideoBitrate = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoControl(v *CreateWmEmbedTaskRequestVideoControl) *CreateWmEmbedTaskRequest {
	s.VideoControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoIsLong(v bool) *CreateWmEmbedTaskRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoBytesB64(v string) *CreateWmEmbedTaskRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoSize(v int64) *CreateWmEmbedTaskRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoUint(v string) *CreateWmEmbedTaskRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmType(v string) *CreateWmEmbedTaskRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) Validate() error {
	if s.AudioControl != nil {
		if err := s.AudioControl.Validate(); err != nil {
			return err
		}
	}
	if s.CsvControl != nil {
		if err := s.CsvControl.Validate(); err != nil {
			return err
		}
	}
	if s.DocumentControl != nil {
		if err := s.DocumentControl.Validate(); err != nil {
			return err
		}
	}
	if s.ImageControl != nil {
		if err := s.ImageControl.Validate(); err != nil {
			return err
		}
	}
	if s.VideoControl != nil {
		if err := s.VideoControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestAudioControl struct {
	// The control parameters for audio metadata.
	MetadataControl *CreateWmEmbedTaskRequestAudioControlMetadataControl `json:"MetadataControl,omitempty" xml:"MetadataControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestAudioControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestAudioControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestAudioControl) GetMetadataControl() *CreateWmEmbedTaskRequestAudioControlMetadataControl {
	return s.MetadataControl
}

func (s *CreateWmEmbedTaskRequestAudioControl) SetMetadataControl(v *CreateWmEmbedTaskRequestAudioControlMetadataControl) *CreateWmEmbedTaskRequestAudioControl {
	s.MetadataControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestAudioControl) Validate() error {
	if s.MetadataControl != nil {
		if err := s.MetadataControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestAudioControlMetadataControl struct {
	// Specifies whether to enable this feature.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The metadata in Base64 format. Encode the following string in Base64 format: AIGC={"Label":"1","ContentProducer":"AXXXX","ProduceID":"BXXXX,"ReservedCode1":"CXXX","ContentPropagator":"DXXX","PropagateID":"EXXX","ReservedCode2":"FXXXX"}. Note: 1. The "AIGC=" prefix is required. Otherwise, the metadata cannot be added. The prefix differs from that of image metadata. 2. The Base64 encoding must be in standard format with padding.
	//
	// example:
	//
	// QUlHQz17IkxhYmVsIjoiMSIsIkNvbnRlbnRQcm9kdWNlciI6IkFYWFhYIiwiUHJvZHVjZUlEIjoiQlhYWFgsIlJlc2VydmVkQ29kZTEiOiJDWFhYIiwiQ29udGVudFByb3BhZ2F0b3IiOiJEWFhYIiwiUHJvcGFnYXRlSUQiOiJFWFhYIiwiUmVzZXJ2ZWRDb2RlMiI6IkZYWFhYIn0=
	XmpKvBase64 *string `json:"XmpKvBase64,omitempty" xml:"XmpKvBase64,omitempty"`
}

func (s CreateWmEmbedTaskRequestAudioControlMetadataControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestAudioControlMetadataControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestAudioControlMetadataControl) GetEnable() *bool {
	return s.Enable
}

func (s *CreateWmEmbedTaskRequestAudioControlMetadataControl) GetXmpKvBase64() *string {
	return s.XmpKvBase64
}

func (s *CreateWmEmbedTaskRequestAudioControlMetadataControl) SetEnable(v bool) *CreateWmEmbedTaskRequestAudioControlMetadataControl {
	s.Enable = &v
	return s
}

func (s *CreateWmEmbedTaskRequestAudioControlMetadataControl) SetXmpKvBase64(v string) *CreateWmEmbedTaskRequestAudioControlMetadataControl {
	s.XmpKvBase64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequestAudioControlMetadataControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestCsvControl struct {
	// The timestamp watermark information bit width. Specifies how much information a single timestamp can contain. A larger value theoretically reduces the number of rows required for extraction, but increases the time modification magnitude. The magnitude range is 2^n, where n is this parameter value.
	//
	// example:
	//
	// 2
	EmbedBitsNumberInEachTime *int64 `json:"EmbedBitsNumberInEachTime,omitempty" xml:"EmbedBitsNumberInEachTime,omitempty"`
	// The column to embed the watermark. We recommend that you use a string content column. Counting starts from 1.
	//
	// example:
	//
	// 1
	EmbedColumn *int64 `json:"EmbedColumn,omitempty" xml:"EmbedColumn,omitempty"`
	// The zero-width character watermark parameter that specifies the embedding density. Valid values: a floating-point number between 0 and 1. 0 indicates that only the first row is embedded. 1 indicates that all rows are embedded.
	//
	// example:
	//
	// 1
	EmbedDensity *string `json:"EmbedDensity,omitempty" xml:"EmbedDensity,omitempty"`
	// The modification precision, which indicates the magnitude of modification as a power of 10. For example, 0 indicates a modification precision of 10^0 (the ones place), -1 indicates the first decimal place, and 1 indicates the tens place. If the float data does not have this precision, no modification is made.
	//
	// example:
	//
	// -1
	EmbedPrecision *int64 `json:"EmbedPrecision,omitempty" xml:"EmbedPrecision,omitempty"`
	// The timestamp watermark parameter that specifies the watermark embedding position. Valid values: Min (minute), Sec (second), and MilSec (millisecond). Select one of the three. The algorithm modifies the data at the selected position.
	//
	// example:
	//
	// Sec
	EmbedTimePosition *string `json:"EmbedTimePosition,omitempty" xml:"EmbedTimePosition,omitempty"`
	// The watermark embedding method.
	//
	// example:
	//
	// lossy_zero_width_embed
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The timestamp watermark parameter that specifies the format string for parsing timestamps in the CSV file. For example, if the timestamp in the CSV file is similar to "2023-10-15 13:20:59:342", the corresponding format string is "Year-Mon-Day Hour:Min:Sec.MilSec". The watermark output retains this format after embedding. If the format is incorrect, this method cannot be used. Year, month, day, hour, minute, second, and millisecond in the format string must follow the specified notation. Delimiters must be single non-alphabetic English characters, typically ":", "/", "-", or " " (space). "T" and "Z" are also supported as delimiters. Other time formats are not currently supported.
	//
	// example:
	//
	// Hour:Min:Sec
	TimeFormat *string `json:"TimeFormat,omitempty" xml:"TimeFormat,omitempty"`
}

func (s CreateWmEmbedTaskRequestCsvControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestCsvControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedBitsNumberInEachTime() *int64 {
	return s.EmbedBitsNumberInEachTime
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedColumn() *int64 {
	return s.EmbedColumn
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedDensity() *string {
	return s.EmbedDensity
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedPrecision() *int64 {
	return s.EmbedPrecision
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedTimePosition() *string {
	return s.EmbedTimePosition
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetMethod() *string {
	return s.Method
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetTimeFormat() *string {
	return s.TimeFormat
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedBitsNumberInEachTime(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedBitsNumberInEachTime = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedColumn(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedColumn = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedDensity(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedDensity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedPrecision(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedPrecision = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedTimePosition(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedTimePosition = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetMethod(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.Method = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetTimeFormat(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.TimeFormat = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestDocumentControl struct {
	// The background watermark control parameters.
	BackgroundControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControl `json:"BackgroundControl,omitempty" xml:"BackgroundControl,omitempty" type:"Struct"`
	// Specifies whether to enable component invisible watermark. The component invisible watermark can resist document addition, deletion, modification, save-as (same format), and full-select copy from docx to a new docx document. It cannot resist format conversion attacks. Valid values:
	//
	// example:
	//
	// true
	InvisibleAntiAllCopy *bool `json:"InvisibleAntiAllCopy,omitempty" xml:"InvisibleAntiAllCopy,omitempty"`
	// Specifies whether to enable zero-width character invisible watermark. The zero-width character invisible watermark can resist document addition, deletion, modification, save-as (same format), partial text copy-paste, and CopytoTxt attacks. It cannot resist format conversion toPDF attacks. Valid values:
	//
	// example:
	//
	// true
	InvisibleAntiTextCopy *bool `json:"InvisibleAntiTextCopy,omitempty" xml:"InvisibleAntiTextCopy,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControl) GetBackgroundControl() *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	return s.BackgroundControl
}

func (s *CreateWmEmbedTaskRequestDocumentControl) GetInvisibleAntiAllCopy() *bool {
	return s.InvisibleAntiAllCopy
}

func (s *CreateWmEmbedTaskRequestDocumentControl) GetInvisibleAntiTextCopy() *bool {
	return s.InvisibleAntiTextCopy
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetBackgroundControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) *CreateWmEmbedTaskRequestDocumentControl {
	s.BackgroundControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetInvisibleAntiAllCopy(v bool) *CreateWmEmbedTaskRequestDocumentControl {
	s.InvisibleAntiAllCopy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetInvisibleAntiTextCopy(v bool) *CreateWmEmbedTaskRequestDocumentControl {
	s.InvisibleAntiTextCopy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) Validate() error {
	if s.BackgroundControl != nil {
		if err := s.BackgroundControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControl struct {
	// Specifies whether to add a background invisible watermark. Valid values:
	//
	// example:
	//
	// true
	BgAddInvisible *bool `json:"BgAddInvisible,omitempty" xml:"BgAddInvisible,omitempty"`
	// Specifies whether to enable the background visible watermark. Valid values:
	//
	// example:
	//
	// true
	BgAddVisible *bool `json:"BgAddVisible,omitempty" xml:"BgAddVisible,omitempty"`
	// The background invisible watermark control parameters.
	BgInvisibleControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl `json:"BgInvisibleControl,omitempty" xml:"BgInvisibleControl,omitempty" type:"Struct"`
	// The background visible watermark control parameters.
	BgVisibleControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl `json:"BgVisibleControl,omitempty" xml:"BgVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgAddInvisible() *bool {
	return s.BgAddInvisible
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgAddVisible() *bool {
	return s.BgAddVisible
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgInvisibleControl() *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl {
	return s.BgInvisibleControl
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgVisibleControl() *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	return s.BgVisibleControl
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgAddInvisible(v bool) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgAddInvisible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgAddVisible(v bool) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgAddVisible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgInvisibleControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgInvisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgVisibleControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) Validate() error {
	if s.BgInvisibleControl != nil {
		if err := s.BgInvisibleControl.Validate(); err != nil {
			return err
		}
	}
	if s.BgVisibleControl != nil {
		if err := s.BgVisibleControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl struct {
	// The opacity parameter of the background invisible watermark. Valid values: 1 to 13. A larger value indicates less transparency.
	//
	// example:
	//
	// 10
	Opacity *int64 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) GetOpacity() *int64 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) SetOpacity(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl struct {
	// The counterclockwise rotation angle of the visible watermark text. Valid values: 1 to 360.
	//
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// The font color of the visible watermark text. The format is 0xFFFFFF RGB color format. For example, 0x000000 indicates black.
	//
	// example:
	//
	// 0x000000
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font size of the visible watermark text. A larger value indicates a larger font.
	//
	// example:
	//
	// 30
	FontSize *int64 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Takes effect when Mode is set to repeat. Specifies the number of times the visible watermark repeats horizontally.
	//
	// example:
	//
	// 3
	HorizontalNumber *int64 `json:"HorizontalNumber,omitempty" xml:"HorizontalNumber,omitempty"`
	// The background visible watermark mode. Valid values:
	//
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The opacity parameter of the visible watermark. Valid values: 1 to 255. A larger value indicates less transparency.
	//
	// example:
	//
	// 100
	Opacity *int64 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// Takes effect when Mode is set to pos. Controls the horizontal position of the visible watermark, with the lower-left corner as the origin. When the value is between 0 and 1, it represents proportional control. When the value is greater than 1, it represents precise pixel position control.
	//
	// example:
	//
	// 0.5
	PosX *string `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// Takes effect when Mode is set to pos. Controls the vertical position of the visible watermark, with the lower-left corner as the origin. When the value is between 0 and 1, it represents proportional control. When the value is greater than 1, it represents precise pixel position control.
	//
	// example:
	//
	// 0.5
	PosY *string `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// Takes effect when Mode is set to repeat. Specifies the number of times the visible watermark repeats vertically.
	//
	// example:
	//
	// 3
	VerticalNumber *int64 `json:"VerticalNumber,omitempty" xml:"VerticalNumber,omitempty"`
	// The background visible watermark text. The format is a UTF-8 string.
	//
	// example:
	//
	// hello ****
	VisibleText *string `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetFontColor() *string {
	return s.FontColor
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetFontSize() *int64 {
	return s.FontSize
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetHorizontalNumber() *int64 {
	return s.HorizontalNumber
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetOpacity() *int64 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetPosX() *string {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetPosY() *string {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetVerticalNumber() *int64 {
	return s.VerticalNumber
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetVisibleText() *string {
	return s.VisibleText
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetFontColor(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetFontSize(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetHorizontalNumber(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.HorizontalNumber = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetOpacity(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetPosX(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetPosY(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetVerticalNumber(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.VerticalNumber = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetVisibleText(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.VisibleText = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControl struct {
	// The control parameters for logo watermarks.
	LogoVisibleControl *CreateWmEmbedTaskRequestImageControlLogoVisibleControl `json:"LogoVisibleControl,omitempty" xml:"LogoVisibleControl,omitempty" type:"Struct"`
	// The metadata control parameters. This parameter takes effect when WmType is set to PureImage or AigcImage.
	MetadataControl *CreateWmEmbedTaskRequestImageControlMetadataControl `json:"MetadataControl,omitempty" xml:"MetadataControl,omitempty" type:"Struct"`
	// The control parameters for image text watermarks.
	TextVisibleControl *CreateWmEmbedTaskRequestImageControlTextVisibleControl `json:"TextVisibleControl,omitempty" xml:"TextVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestImageControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControl) GetLogoVisibleControl() *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	return s.LogoVisibleControl
}

func (s *CreateWmEmbedTaskRequestImageControl) GetMetadataControl() *CreateWmEmbedTaskRequestImageControlMetadataControl {
	return s.MetadataControl
}

func (s *CreateWmEmbedTaskRequestImageControl) GetTextVisibleControl() *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	return s.TextVisibleControl
}

func (s *CreateWmEmbedTaskRequestImageControl) SetLogoVisibleControl(v *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) *CreateWmEmbedTaskRequestImageControl {
	s.LogoVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControl) SetMetadataControl(v *CreateWmEmbedTaskRequestImageControlMetadataControl) *CreateWmEmbedTaskRequestImageControl {
	s.MetadataControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControl) SetTextVisibleControl(v *CreateWmEmbedTaskRequestImageControlTextVisibleControl) *CreateWmEmbedTaskRequestImageControl {
	s.TextVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControl) Validate() error {
	if s.LogoVisibleControl != nil {
		if err := s.LogoVisibleControl.Validate(); err != nil {
			return err
		}
	}
	if s.MetadataControl != nil {
		if err := s.MetadataControl.Validate(); err != nil {
			return err
		}
	}
	if s.TextVisibleControl != nil {
		if err := s.TextVisibleControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestImageControlLogoVisibleControl struct {
	// The clockwise rotation angle of the logo watermark. Valid values: 1 to 360.
	//
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// Specifies whether to enable enhanced visible watermarking. After this feature is enabled, the logo is processed so that information embedded in the logo can be extracted.
	//
	// example:
	//
	// false
	Enhance *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// The logo watermark in Base64 format. The logo file is a PNG image converted to Base64 format.
	//
	// example:
	//
	// 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
	LogoBase64 *string `json:"LogoBase64,omitempty" xml:"LogoBase64,omitempty"`
	// This parameter takes effect when Mode is set to top-left, top-right, bottom-left, or bottom-right. The margin settings.
	Margin *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	// The display mode of the logo watermark. Valid values:
	//
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The opacity of the logo watermark. Valid values: 1 to 255. A larger value indicates lower transparency.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// The horizontal anchor point of the logo watermark. Valid values: 0 to 1. When (PosAx, PosAy) is set to (0, 0), the watermark is drawn with the upper-left corner as the anchor point. When the value is 0.5, the watermark is drawn at the center. When (PosAx, PosAy) is set to (1, 1), the watermark is drawn at the lower-right corner.
	//
	// example:
	//
	// 0
	PosAx *float32 `json:"PosAx,omitempty" xml:"PosAx,omitempty"`
	// The vertical anchor point of the logo watermark. Valid values: 0 to 1. When (PosAx, PosAy) is set to (0, 0), the watermark is drawn with the upper-left corner as the anchor point. When the value is 0.5, the watermark is drawn at the center. When (PosAx, PosAy) is set to (1, 1), the watermark is drawn at the lower-right corner.
	//
	// example:
	//
	// 0
	PosAy *float32 `json:"PosAy,omitempty" xml:"PosAy,omitempty"`
	// This parameter takes effect when Mode is set to pos. Specifies the horizontal position of the visible watermark in pixels, with the upper-left corner as the origin.
	//
	// example:
	//
	// 0
	PosX *int64 `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// This parameter takes effect when Mode is set to pos. Specifies the vertical position of the visible watermark in pixels, with the upper-left corner as the origin.
	//
	// example:
	//
	// 0
	PosY *int64 `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// This parameter takes effect when Mode is set to repeat. Specifies the horizontal spacing for tiled visible watermarks.
	//
	// example:
	//
	// 30
	SpaceX *int64 `json:"SpaceX,omitempty" xml:"SpaceX,omitempty"`
	// This parameter takes effect when Mode is set to repeat. Specifies the vertical spacing for tiled visible watermarks.
	//
	// example:
	//
	// 30
	SpaceY *int64 `json:"SpaceY,omitempty" xml:"SpaceY,omitempty"`
	// The visibility. Valid values:
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetEnhance() *bool {
	return s.Enhance
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetLogoBase64() *string {
	return s.LogoBase64
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetMargin() *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosAx() *float32 {
	return s.PosAx
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosAy() *float32 {
	return s.PosAy
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosX() *int64 {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosY() *int64 {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetSpaceX() *int64 {
	return s.SpaceX
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetSpaceY() *int64 {
	return s.SpaceY
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetEnhance(v bool) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Enhance = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetLogoBase64(v string) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.LogoBase64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetMargin(v *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetOpacity(v int32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosAx(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosAx = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosAy(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosAy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosX(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosY(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetSpaceX(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.SpaceX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetSpaceY(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.SpaceY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetVisible(v bool) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) Validate() error {
	if s.Margin != nil {
		if err := s.Margin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin struct {
	// This parameter takes effect when Mode is set to bottom-left or bottom-right. The bottom margin.
	//
	// example:
	//
	// 0
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// This parameter takes effect when Mode is set to top-left or bottom-left. The left margin.
	//
	// example:
	//
	// 0
	Left *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// This parameter takes effect when Mode is set to top-right or bottom-right. The right margin.
	//
	// example:
	//
	// 0
	Right *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	// This parameter takes effect when Mode is set to top-left or top-right. The top margin.
	//
	// example:
	//
	// 0
	Top *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetBottom() *float32 {
	return s.Bottom
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetLeft() *float32 {
	return s.Left
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetRight() *float32 {
	return s.Right
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetTop() *float32 {
	return s.Top
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetBottom(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetLeft(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Left = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetRight(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetTop(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Top = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlMetadataControl struct {
	// Specifies whether to enable this feature.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The metadata in Base64 format. Encode the following string in Base64 format: AIGC:{"Label":"1","ContentProducer":"AXXXX","ProduceID":"BXXXX,"ReservedCode1":"CXXX","ContentPropagator":"DXXX","PropagateID":"EXXX","ReservedCode2":"FXXXX"}. Note: 1. The "AIGC:" prefix is required. Otherwise, the metadata cannot be added. The format differs from that of audio and video metadata. 2. The Base64 encoding must be in standard format with padding.
	//
	// example:
	//
	// QUlHQzp7IkxhYmVsIjoiMSIsIkNvbnRlbnRQcm9kdWNlciI6IkFYWFhYIiwiUHJvZHVjZUlEIjoiQlhYWFgsIlJlc2VydmVkQ29kZTEiOiJDWFhYIiwiQ29udGVudFByb3BhZ2F0b3IiOiJEWFhYIiwiUHJvcGFnYXRlSUQiOiJFWFhYIiwiUmVzZXJ2ZWRDb2RlMiI6IkZYWFhYIn0=
	XmpKvBase64 *string `json:"XmpKvBase64,omitempty" xml:"XmpKvBase64,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlMetadataControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlMetadataControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) GetEnable() *bool {
	return s.Enable
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) GetXmpKvBase64() *string {
	return s.XmpKvBase64
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) SetEnable(v bool) *CreateWmEmbedTaskRequestImageControlMetadataControl {
	s.Enable = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) SetXmpKvBase64(v string) *CreateWmEmbedTaskRequestImageControlMetadataControl {
	s.XmpKvBase64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlTextVisibleControl struct {
	// The clockwise rotation angle of the text watermark. Valid values: 0 to 360.
	//
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// The font color of the text watermark. The format is 0xFFFFFF or #FFFFFF RGB color format. For example, 0x000000 or #000000 indicates black.
	//
	// example:
	//
	// #FF0000
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font size of the text watermark. A larger value indicates a larger font.
	//
	// example:
	//
	// 30
	FontSize *int64 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// This parameter takes effect when Mode is set to top-left, top-right, bottom-left, or bottom-right. The margin settings.
	Margin *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	// The display mode of the text watermark. Valid values:
	//
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The opacity of the text watermark. Valid values: 1 to 255. A larger value indicates lower transparency.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// The horizontal anchor point of the text watermark.
	//
	// Valid values: 0 to 1. When (PosAx, PosAy) is set to (0, 0), the text is drawn with the upper-left corner as the anchor point. When the value is 0.5, the text is drawn at the center point. When (PosAx, PosAy) is set to (1, 1), the text is drawn with the lower-right corner as the anchor point.
	//
	// example:
	//
	// 0
	PosAx *float32 `json:"PosAx,omitempty" xml:"PosAx,omitempty"`
	// The vertical anchor point of the text watermark.
	//
	// Valid values: 0 to 1. When (PosAx, PosAy) is set to (0, 0), the text is drawn with the upper-left corner as the anchor point. When the value is 0.5, the text is drawn from the center point. When (PosAx, PosAy) is set to (1, 1), the text is drawn with the lower-right corner as the anchor point.
	//
	// example:
	//
	// 0
	PosAy *float32 `json:"PosAy,omitempty" xml:"PosAy,omitempty"`
	// This parameter takes effect when Mode is set to pos. Specifies the horizontal position of the text watermark in pixels, with the upper-left corner as the origin.
	//
	// example:
	//
	// 0
	PosX *int64 `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// This parameter takes effect when Mode is set to pos. Specifies the vertical position of the text watermark in pixels, with the upper-left corner as the origin.
	//
	// example:
	//
	// 0
	PosY *int64 `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// This parameter takes effect when Mode is set to repeat. Specifies the horizontal spacing for tiled text watermarks.
	//
	// example:
	//
	// 30
	SpaceX *int64 `json:"SpaceX,omitempty" xml:"SpaceX,omitempty"`
	// This parameter takes effect when Mode is set to repeat. Specifies the vertical spacing for tiled text watermarks.
	//
	// example:
	//
	// 0
	SpaceY *int64 `json:"SpaceY,omitempty" xml:"SpaceY,omitempty"`
	// The visibility. Valid values:
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
	// The text watermark content. The format is a UTF-8 string.
	//
	// example:
	//
	// WatermarkText
	VisibleText *string `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetFontColor() *string {
	return s.FontColor
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetFontSize() *int64 {
	return s.FontSize
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetMargin() *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosAx() *float32 {
	return s.PosAx
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosAy() *float32 {
	return s.PosAy
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosX() *int64 {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosY() *int64 {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetSpaceX() *int64 {
	return s.SpaceX
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetSpaceY() *int64 {
	return s.SpaceY
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetVisibleText() *string {
	return s.VisibleText
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetFontColor(v string) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetFontSize(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetMargin(v *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetOpacity(v int32) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosAx(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosAx = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosAy(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosAy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosX(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosY(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetSpaceX(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.SpaceX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetSpaceY(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.SpaceY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetVisible(v bool) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetVisibleText(v string) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.VisibleText = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) Validate() error {
	if s.Margin != nil {
		if err := s.Margin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin struct {
	// This parameter takes effect when Mode is set to bottom-left or bottom-right. The bottom margin.
	//
	// example:
	//
	// 0
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// This parameter takes effect when Mode is set to top-left or bottom-left. The left margin.
	//
	// example:
	//
	// 0
	Left *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// This parameter takes effect when Mode is set to top-right or bottom-right. The right margin.
	//
	// example:
	//
	// 0
	Right *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	// This parameter takes effect when Mode is set to top-left or top-right. The top margin.
	//
	// example:
	//
	// 0
	Top *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetBottom() *float32 {
	return s.Bottom
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetLeft() *float32 {
	return s.Left
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetRight() *float32 {
	return s.Right
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetTop() *float32 {
	return s.Top
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetBottom(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetLeft(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Left = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetRight(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetTop(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Top = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestVideoControl struct {
	// The metadata control parameters.
	MetadataControl *CreateWmEmbedTaskRequestVideoControlMetadataControl `json:"MetadataControl,omitempty" xml:"MetadataControl,omitempty" type:"Struct"`
	// The control parameters for video text watermarks.
	TextVisibleControl *CreateWmEmbedTaskRequestVideoControlTextVisibleControl `json:"TextVisibleControl,omitempty" xml:"TextVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestVideoControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestVideoControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestVideoControl) GetMetadataControl() *CreateWmEmbedTaskRequestVideoControlMetadataControl {
	return s.MetadataControl
}

func (s *CreateWmEmbedTaskRequestVideoControl) GetTextVisibleControl() *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	return s.TextVisibleControl
}

func (s *CreateWmEmbedTaskRequestVideoControl) SetMetadataControl(v *CreateWmEmbedTaskRequestVideoControlMetadataControl) *CreateWmEmbedTaskRequestVideoControl {
	s.MetadataControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControl) SetTextVisibleControl(v *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) *CreateWmEmbedTaskRequestVideoControl {
	s.TextVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControl) Validate() error {
	if s.MetadataControl != nil {
		if err := s.MetadataControl.Validate(); err != nil {
			return err
		}
	}
	if s.TextVisibleControl != nil {
		if err := s.TextVisibleControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestVideoControlMetadataControl struct {
	// Specifies whether to enable this feature.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The metadata in Base64 format. Encode the following string in Base64 format: AIGC={"Label":"1","ContentProducer":"AXXXX","ProduceID":"BXXXX,"ReservedCode1":"CXXX","ContentPropagator":"DXXX","PropagateID":"EXXX","ReservedCode2":"FXXXX"}. Note: 1. The "AIGC=" prefix is required. Otherwise, the metadata cannot be added. The prefix differs from that of image metadata. 2. The Base64 encoding must be in standard format with padding.
	//
	// example:
	//
	// QUlHQz17IkxhYmVsIjoiMSIsIkNvbnRlbnRQcm9kdWNlciI6IkFYWFhYIiwiUHJvZHVjZUlEIjoiQlhYWFgsIlJlc2VydmVkQ29kZTEiOiJDWFhYIiwiQ29udGVudFByb3BhZ2F0b3IiOiJEWFhYIiwiUHJvcGFnYXRlSUQiOiJFWFhYIiwiUmVzZXJ2ZWRDb2RlMiI6IkZYWFhYIn0=
	XmpKvBase64 *string `json:"XmpKvBase64,omitempty" xml:"XmpKvBase64,omitempty"`
}

func (s CreateWmEmbedTaskRequestVideoControlMetadataControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestVideoControlMetadataControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestVideoControlMetadataControl) GetEnable() *bool {
	return s.Enable
}

func (s *CreateWmEmbedTaskRequestVideoControlMetadataControl) GetXmpKvBase64() *string {
	return s.XmpKvBase64
}

func (s *CreateWmEmbedTaskRequestVideoControlMetadataControl) SetEnable(v bool) *CreateWmEmbedTaskRequestVideoControlMetadataControl {
	s.Enable = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlMetadataControl) SetXmpKvBase64(v string) *CreateWmEmbedTaskRequestVideoControlMetadataControl {
	s.XmpKvBase64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlMetadataControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestVideoControlTextVisibleControl struct {
	// The font color of the text watermark. The format is 0xFFFFFF or #FFFFFF RGB color format.
	//
	// example:
	//
	// #FF0000
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font size. Valid values: **0*	- to **72**.
	//
	// example:
	//
	// 30
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// This parameter takes effect when Mode is set to top-left, top-right, bottom-left, or bottom-right. The margin settings.
	Margin *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	// The display mode of the text watermark. Valid values:
	//
	// - **pos**: fixed position, with the upper-left corner as the origin.
	//
	// - **bottom-right**: lower-left mode.
	//
	// example:
	//
	// bottom-right
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The opacity of the text watermark. Valid values: 1 to 255. A larger value indicates lower transparency.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// This parameter takes effect when Mode is set to pos. Specifies the horizontal position of the visible watermark in pixels, with the upper-left corner as the origin.
	//
	// example:
	//
	// 10
	PosX *int32 `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// This parameter takes effect when Mode is set to pos. Specifies the vertical position of the visible watermark in pixels, with the upper-left corner as the origin.
	//
	// example:
	//
	// 10
	PosY *int32 `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// The visibility. Valid values:
	//
	// example:
	//
	// True
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
	// The text watermark content. The format is a UTF-8 string.
	//
	// example:
	//
	// WatermarkTest
	VisibleText *string `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
}

func (s CreateWmEmbedTaskRequestVideoControlTextVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetFontColor() *string {
	return s.FontColor
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetFontSize() *int32 {
	return s.FontSize
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetMargin() *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetPosX() *int32 {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetPosY() *int32 {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) GetVisibleText() *string {
	return s.VisibleText
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetFontColor(v string) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetFontSize(v int32) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetMargin(v *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetOpacity(v int32) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetPosX(v int32) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetPosY(v int32) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetVisible(v bool) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) SetVisibleText(v string) *CreateWmEmbedTaskRequestVideoControlTextVisibleControl {
	s.VisibleText = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControl) Validate() error {
	if s.Margin != nil {
		if err := s.Margin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin struct {
	// This parameter takes effect when Mode is set to bottom-left or bottom-right. The bottom margin.
	//
	// example:
	//
	// 10
	Bottom *int32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// This parameter takes effect when Mode is set to top-right or bottom-right. The right margin.
	//
	// example:
	//
	// 10
	Right *int32 `json:"Right,omitempty" xml:"Right,omitempty"`
}

func (s CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) GetBottom() *int32 {
	return s.Bottom
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) GetRight() *int32 {
	return s.Right
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) SetBottom(v int32) *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) SetRight(v int32) *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmEmbedTaskRequestVideoControlTextVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}
