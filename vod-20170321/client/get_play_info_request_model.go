// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionType(v string) *GetPlayInfoRequest
	GetAdditionType() *string
	SetAuthTimeout(v int64) *GetPlayInfoRequest
	GetAuthTimeout() *int64
	SetCodecName(v string) *GetPlayInfoRequest
	GetCodecName() *string
	SetDefinition(v string) *GetPlayInfoRequest
	GetDefinition() *string
	SetDigitalWatermarkType(v string) *GetPlayInfoRequest
	GetDigitalWatermarkType() *string
	SetFormats(v string) *GetPlayInfoRequest
	GetFormats() *string
	SetOutputType(v string) *GetPlayInfoRequest
	GetOutputType() *string
	SetPlayConfig(v string) *GetPlayInfoRequest
	GetPlayConfig() *string
	SetReAuthInfo(v string) *GetPlayInfoRequest
	GetReAuthInfo() *string
	SetReferenceId(v string) *GetPlayInfoRequest
	GetReferenceId() *string
	SetResultType(v string) *GetPlayInfoRequest
	GetResultType() *string
	SetStreamType(v string) *GetPlayInfoRequest
	GetStreamType() *string
	SetTrace(v string) *GetPlayInfoRequest
	GetTrace() *string
	SetVideoId(v string) *GetPlayInfoRequest
	GetVideoId() *string
}

type GetPlayInfoRequest struct {
	// Obtains the URL of the China-accessible bullet screen mask data. Valid values: **danmu**.
	//
	// > This parameter takes effect only when `outputType` is set to `cdn`.
	//
	// example:
	//
	// danmu
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	// The validity period of the playback URL. Unit: seconds.
	//
	// - If OutputType is set to **cdn**:
	//
	//     - The playback URL expires periodically only when URL authentication is enabled. Otherwise, the URL is permanently valid. For information about how to enable and configure URL authentication, refer to [URL authentication](https://help.aliyun.com/document_detail/86090.html).
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If this parameter is not specified, the default validity period configured in URL authentication is used.
	//
	// - If OutputType is set to **oss**:
	//
	//     - The playback URL expires periodically only when the storage permission is private. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: To reduce security risks to the origin server, when audio or video files are stored in an ApsaraVideo VOD system bucket, the maximum value is **604800*	- (7 days). When audio or video files are stored in your own OSS bucket, the maximum value is **129600*	- (36 hours). If the maximum value does not meet your requirements, set OutputType to **cdn*	- and configure URL authentication to set a longer validity period.
	//
	//     - Default value: If this parameter is not specified, the default value is **3600**.
	//
	// example:
	//
	// 1800
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// example:
	//
	// H264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The definition of the video stream. Separate multiple definitions with commas (,). Valid values:
	//
	// - **FD**: low definition.
	//
	// - **LD**: standard definition.
	//
	// - **SD**: high definition.
	//
	// - **HD**: ultra-high definition.
	//
	// - **OD**: original definition.
	//
	// - **2K**: 2K.
	//
	// - **4K**: 4K.
	//
	// - **SQ**: standard sound quality.
	//
	// - **HQ**: high sound quality.
	//
	// - **AUTO**: adaptive bitrate streaming.
	//
	// > - By default, streams of all definitions are returned.
	//
	// > - When generating tracing watermarks, this parameter is required and must be consistent with the definition configured during tracing watermark transcoding.
	//
	// > - The AUTO definition is returned only when transcoding packaging is configured in the transcoding template. For more information, refer to [PackageSetting: transcoding packaging settings](~~52839#title-4fk-cg8-gzx~~).
	//
	// example:
	//
	// LD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The digital watermarking type. Valid values:
	//
	// - TraceMark: tracing watermark.
	//
	// - CopyrightMark: copyright watermark.
	//
	// example:
	//
	// TraceMark
	DigitalWatermarkType *string `json:"DigitalWatermarkType,omitempty" xml:"DigitalWatermarkType,omitempty"`
	// The media stream format. Separate multiple formats with commas (,). Valid values:
	//
	// - **mp4**
	//
	// - **m3u8**
	//
	// - **mp3**
	//
	// - **flv**
	//
	// - **mpd**
	//
	//
	// > - By default, streams in all formats are returned.
	//
	// > - The mpd format is returned only when the `dash` container format is configured in the transcoding template. For more information, refer to [Container: container format](~~52839#title-7rr-3hj-gy5~~).
	//
	// example:
	//
	// mp4,m3u8
	Formats *string `json:"Formats,omitempty" xml:"Formats,omitempty"`
	// The type of the output URL. Valid values:
	//
	// - **oss**: back-to-origin URL.
	//
	// - **cdn*	- (default): accelerated URL.
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The custom playback settings. The value is a JSON string that supports specifying domain name playback settings. For details about parameter construction, refer to [PlayConfig](~~86952#section-9g7-s9b-v7z~~).
	//
	// > - If PlayConfig is not set or `PlayDomain` within it is not set, the operation uses the default domain name configured in ApsaraVideo VOD. If no default domain name is configured, the most recently modified domain name is used as the playback domain name based on reverse chronological order of modification time. To prevent an unexpected domain name from being returned, set a default playback domain name. Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management*	- > **Storage*	- > **Manage*	- > **Domain names that perform origin fetch from this storage address*	- to set the default playback domain name.
	//
	// > - When the `EncryptType` parameter in PlayConfig is set to `AliyunVoDEncryption`, the playback URL of the privately encrypted stream is not returned by default to ensure video security. To return the playback URL of the privately encrypted stream, set the `ResultType` parameter to `Multiple`.
	//
	// example:
	//
	// {"PlayDomain":"vod.test_domain","XForwardedFor":"yqCD7Fp1uqChoVj/sl/p5Q==","PreviewTime":"20","MtsHlsUriToken":"yqCD7Fp1uqChoVjslp5Q"}
	PlayConfig *string `json:"PlayConfig,omitempty" xml:"PlayConfig,omitempty"`
	// The CDN reauthentication parameter. The value is a JSON string. When type A signing is enabled for URL authentication, you can use this parameter to set the `uid` and `rand` of the authentication URL. For more information, refer to [Type A signing](https://help.aliyun.com/document_detail/2249352.html).
	//
	// example:
	//
	// {"uid":"12345","rand":"abckljd"}
	ReAuthInfo *string `json:"ReAuthInfo,omitempty" xml:"ReAuthInfo,omitempty"`
	// The custom ID. Only lowercase letters, uppercase letters, digits, hyphens, and underscores are supported. The length is 6 to 64 characters. The ID is unique per user.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The type of the returned data. Valid values:
	//
	// - **Single*	- (default): returns only the latest transcoded stream for each definition and format.
	//
	// - **Multiple**: returns all transcoded streams for each definition and format.
	//
	// example:
	//
	// Single
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The media stream type. Separate multiple types with commas (,). Valid values:
	//
	// - **video**: video.
	//
	// - **audio**: audio.
	//
	// By default, streams of all types are returned.
	//
	// example:
	//
	// video
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The custom digital watermarking settings.
	//
	// - When `DigitalWatermarkType` is set to `TraceMark`, pass in this parameter to set the tracing watermark information for the video and return the video stream that contains the watermark information. Only English letters, digits, and Chinese characters are supported. A maximum of 1024 characters are supported.
	//
	// - When `DigitalWatermarkType` is set to `CopyrightMark`, `Trace` corresponds to the **watermark text*	- configured when the watermark template was created. Pass in this parameter to query and return the video stream with the specified watermark text.
	//
	// example:
	//
	// test mark
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// The audio or video ID. Only a single audio or video ID is supported. You can obtain the ID by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When uploading audio or video files by calling the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// 93ab850b4f654b6e91d24d81d44****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetPlayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPlayInfoRequest) GetAdditionType() *string {
	return s.AdditionType
}

func (s *GetPlayInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetPlayInfoRequest) GetCodecName() *string {
	return s.CodecName
}

func (s *GetPlayInfoRequest) GetDefinition() *string {
	return s.Definition
}

func (s *GetPlayInfoRequest) GetDigitalWatermarkType() *string {
	return s.DigitalWatermarkType
}

func (s *GetPlayInfoRequest) GetFormats() *string {
	return s.Formats
}

func (s *GetPlayInfoRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetPlayInfoRequest) GetPlayConfig() *string {
	return s.PlayConfig
}

func (s *GetPlayInfoRequest) GetReAuthInfo() *string {
	return s.ReAuthInfo
}

func (s *GetPlayInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetPlayInfoRequest) GetResultType() *string {
	return s.ResultType
}

func (s *GetPlayInfoRequest) GetStreamType() *string {
	return s.StreamType
}

func (s *GetPlayInfoRequest) GetTrace() *string {
	return s.Trace
}

func (s *GetPlayInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetPlayInfoRequest) SetAdditionType(v string) *GetPlayInfoRequest {
	s.AdditionType = &v
	return s
}

func (s *GetPlayInfoRequest) SetAuthTimeout(v int64) *GetPlayInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetPlayInfoRequest) SetCodecName(v string) *GetPlayInfoRequest {
	s.CodecName = &v
	return s
}

func (s *GetPlayInfoRequest) SetDefinition(v string) *GetPlayInfoRequest {
	s.Definition = &v
	return s
}

func (s *GetPlayInfoRequest) SetDigitalWatermarkType(v string) *GetPlayInfoRequest {
	s.DigitalWatermarkType = &v
	return s
}

func (s *GetPlayInfoRequest) SetFormats(v string) *GetPlayInfoRequest {
	s.Formats = &v
	return s
}

func (s *GetPlayInfoRequest) SetOutputType(v string) *GetPlayInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetPlayInfoRequest) SetPlayConfig(v string) *GetPlayInfoRequest {
	s.PlayConfig = &v
	return s
}

func (s *GetPlayInfoRequest) SetReAuthInfo(v string) *GetPlayInfoRequest {
	s.ReAuthInfo = &v
	return s
}

func (s *GetPlayInfoRequest) SetReferenceId(v string) *GetPlayInfoRequest {
	s.ReferenceId = &v
	return s
}

func (s *GetPlayInfoRequest) SetResultType(v string) *GetPlayInfoRequest {
	s.ResultType = &v
	return s
}

func (s *GetPlayInfoRequest) SetStreamType(v string) *GetPlayInfoRequest {
	s.StreamType = &v
	return s
}

func (s *GetPlayInfoRequest) SetTrace(v string) *GetPlayInfoRequest {
	s.Trace = &v
	return s
}

func (s *GetPlayInfoRequest) SetVideoId(v string) *GetPlayInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetPlayInfoRequest) Validate() error {
	return dara.Validate(s)
}
