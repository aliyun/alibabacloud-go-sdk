// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaBase(v *GetPlayInfoResponseBodyMediaBase) *GetPlayInfoResponseBody
	GetMediaBase() *GetPlayInfoResponseBodyMediaBase
	SetPlayInfoList(v []*GetPlayInfoResponseBodyPlayInfoList) *GetPlayInfoResponseBody
	GetPlayInfoList() []*GetPlayInfoResponseBodyPlayInfoList
	SetRequestId(v string) *GetPlayInfoResponseBody
	GetRequestId() *string
}

type GetPlayInfoResponseBody struct {
	// The information about the media asset.
	MediaBase *GetPlayInfoResponseBodyMediaBase `json:"MediaBase,omitempty" xml:"MediaBase,omitempty" type:"Struct"`
	// The information about the audio or video stream.
	PlayInfoList []*GetPlayInfoResponseBodyPlayInfoList `json:"PlayInfoList,omitempty" xml:"PlayInfoList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPlayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBody) GetMediaBase() *GetPlayInfoResponseBodyMediaBase {
	return s.MediaBase
}

func (s *GetPlayInfoResponseBody) GetPlayInfoList() []*GetPlayInfoResponseBodyPlayInfoList {
	return s.PlayInfoList
}

func (s *GetPlayInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlayInfoResponseBody) SetMediaBase(v *GetPlayInfoResponseBodyMediaBase) *GetPlayInfoResponseBody {
	s.MediaBase = v
	return s
}

func (s *GetPlayInfoResponseBody) SetPlayInfoList(v []*GetPlayInfoResponseBodyPlayInfoList) *GetPlayInfoResponseBody {
	s.PlayInfoList = v
	return s
}

func (s *GetPlayInfoResponseBody) SetRequestId(v string) *GetPlayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlayInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPlayInfoResponseBodyMediaBase struct {
	// The category ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [Intelligent Media Services (IMS) console](https://ims.console.aliyun.com) and choose **Media Asset Management*	- > **Category Management*	- to view the category ID.
	//
	// 	- View the value of the CateId parameter returned by the AddCategory operation that you called to create a category.
	//
	// 	- View the value of the CateId parameter returned by the GetCategories operation that you called to query a category.
	//
	// example:
	//
	// 4220
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// https://***.oss-cn-shanghai.aliyuncs.com/cover/281c64d6-b5fb-4c57-97cd-84da56a8b151_large_cover_url.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created.
	//
	// example:
	//
	// 2021-09-22T10:07:31+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The content description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 2eea77a61c7b4ddd95bec34a6f65b***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags.
	//
	// 	- Up to 16 tags are supported.
	//
	// 	- Multiple tags are separated by commas (,).
	//
	// 	- Each tag can be up to 32 bytes in length.
	//
	// 	- The value is encoded in UTF-8.
	//
	// example:
	//
	// test,ccc
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset. Valid values:
	//
	// video audio
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The resource status. Valid values:
	//
	// Init: the initial state, which indicates that the source file is not ready.
	//
	// Preparing: The source file is being prepared. For example, the file is being uploaded or edited.
	//
	// PrepareFail: The source file failed to be prepared. For example, the information of the source file failed to be obtained.
	//
	// Normal: The source file is ready.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	//
	// example:
	//
	// testTitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetPlayInfoResponseBodyMediaBase) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBodyMediaBase) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyMediaBase) GetCateId() *int64 {
	return s.CateId
}

func (s *GetPlayInfoResponseBodyMediaBase) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetPlayInfoResponseBodyMediaBase) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetPlayInfoResponseBodyMediaBase) GetDescription() *string {
	return s.Description
}

func (s *GetPlayInfoResponseBodyMediaBase) GetMediaId() *string {
	return s.MediaId
}

func (s *GetPlayInfoResponseBodyMediaBase) GetMediaTags() *string {
	return s.MediaTags
}

func (s *GetPlayInfoResponseBodyMediaBase) GetMediaType() *string {
	return s.MediaType
}

func (s *GetPlayInfoResponseBodyMediaBase) GetStatus() *string {
	return s.Status
}

func (s *GetPlayInfoResponseBodyMediaBase) GetTitle() *string {
	return s.Title
}

func (s *GetPlayInfoResponseBodyMediaBase) SetCateId(v int64) *GetPlayInfoResponseBodyMediaBase {
	s.CateId = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetCoverURL(v string) *GetPlayInfoResponseBodyMediaBase {
	s.CoverURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetCreationTime(v string) *GetPlayInfoResponseBodyMediaBase {
	s.CreationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetDescription(v string) *GetPlayInfoResponseBodyMediaBase {
	s.Description = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetMediaId(v string) *GetPlayInfoResponseBodyMediaBase {
	s.MediaId = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetMediaTags(v string) *GetPlayInfoResponseBodyMediaBase {
	s.MediaTags = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetMediaType(v string) *GetPlayInfoResponseBodyMediaBase {
	s.MediaType = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetStatus(v string) *GetPlayInfoResponseBodyMediaBase {
	s.Status = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) SetTitle(v string) *GetPlayInfoResponseBodyMediaBase {
	s.Title = &v
	return s
}

func (s *GetPlayInfoResponseBodyMediaBase) Validate() error {
	return dara.Validate(s)
}

type GetPlayInfoResponseBodyPlayInfoList struct {
	// The color depth.
	//
	// example:
	//
	// 8
	BitDepth *int32 `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	// The bitrate of the media stream. Unit: Kbit/s.
	//
	// example:
	//
	// 20
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the media stream was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-10T02:28:49Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The quality of the media stream. Valid values:
	//
	// 	- **FD**: low definition
	//
	// 	- **LD**: standard definition
	//
	// 	- **SD**: high definition
	//
	// 	- **HD**: ultra-high definition
	//
	// 	- **OD**: original definition
	//
	// 	- **2K**
	//
	// 	- **4K**
	//
	// 	- **SQ**: standard sound quality
	//
	// 	- **HQ**: high sound quality
	//
	// 	- **AUTO**: adaptive bitrate
	//
	// example:
	//
	// HD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The duration of the media stream. Unit: seconds.
	//
	// example:
	//
	// 9.0464
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether the media stream is encrypted. Valid values:
	//
	// 	- **0**: The media stream is not encrypted.
	//
	// 	- **1**: The media stream is encrypted.
	//
	// example:
	//
	// 0
	Encrypt *int64 `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// The encryption type of the media stream. Valid values:
	//
	// 	- **AliyunVoDEncryption**: Alibaba Cloud proprietary cryptography
	//
	// 	- **HLSEncryption**: HTTP Live Streaming (HLS) encryption
	//
	// >  If the encryption type is AliyunVoDEncryption, only ApsaraVideo Player SDK can be used to play videos.
	//
	// example:
	//
	// AliyunVoDEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The OSS URL of the file.
	//
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/sv/43a68ee9-181809b6aba/43a68ee9-181809b6aba.mpeg
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The format of the media stream.
	//
	// 	- If the media asset is a video file, the valid values are **mp4*	- and **m3u8**.
	//
	// 	- If the media asset is an audio-only file, the value is **mp3**.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the media stream. Unit: frames per second (FPS).
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The high dynamic range (HDR) type of the media stream. Valid values:
	//
	// 	- HDR
	//
	// 	- HDR10
	//
	// 	- HLG
	//
	// 	- DolbyVision
	//
	// 	- HDRVivid
	//
	// 	- SDR+
	//
	// example:
	//
	// HDR
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The height of the media stream. Unit: pixels.
	//
	// example:
	//
	// 1080
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 36c9d38e70bf43ed9f7f8f48d6356***
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The time when the media stream was updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-13T11:39:41.714+08:00
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The type of Narrowband HD™ transcoding. Valid values:
	//
	// 	- **0**: standard transcoding
	//
	// 	- **1.0**: Narrowband HD™ 1.0 transcoding
	//
	// 	- **2.0**: Narrowband HD™ 2.0 transcoding
	//
	// This parameter is returned only when a definition that is available in the built-in Narrowband HD™ 1.0 transcoding template is specified. For more information, see the [Definition parameter in TranscodeTemplate](https://help.aliyun.com/document_detail/52839.html) table.
	//
	// example:
	//
	// 0
	NarrowBandType *string `json:"NarrowBandType,omitempty" xml:"NarrowBandType,omitempty"`
	// The playback URL of the media stream.
	//
	// example:
	//
	// https://***.aliyuncdn.com/sv/756bee1-17f980f0945/756bee1-17f980f0945.mp4
	PlayURL *string `json:"PlayURL,omitempty" xml:"PlayURL,omitempty"`
	// The size of the media stream. Unit: bytes.
	//
	// example:
	//
	// 418112
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the media stream. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Invisible**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the media stream, which are used to identify the transcoding type.
	//
	// example:
	//
	// "{\\"ims.audioServiceType\\": \\"AudioEnhancement\\"}"
	StreamTags *string `json:"StreamTags,omitempty" xml:"StreamTags,omitempty"`
	// The type of the media stream. If the media stream is a video stream, the value is **video**. If the media stream is an audio-only stream, the value is **audio**.
	//
	// example:
	//
	// video
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The type of the transcoding template. Valid values:
	//
	// 	- Normal: standard transcoding
	//
	// 	- AudioTranscode: audio transcoding
	//
	// 	- Remux: container format conversion
	//
	// 	- NarrowBandV1: Narrowband HD™ 1.0
	//
	// 	- NarrowBandV2: Narrowband HD™ 2.0
	//
	// 	- UHD: audio and video enhancement (ultra-high definition)
	//
	// example:
	//
	// Normal
	TransTemplateType *string `json:"TransTemplateType,omitempty" xml:"TransTemplateType,omitempty"`
	// The ID of the watermark that is associated with the media stream.
	//
	// example:
	//
	// 5bed88672b1e2520ead228935ed51***
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	// The width of the media stream. Unit: pixels.
	//
	// example:
	//
	// 1024
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetPlayInfoResponseBodyPlayInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBodyPlayInfoList) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetBitDepth() *int32 {
	return s.BitDepth
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetDefinition() *string {
	return s.Definition
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetEncrypt() *int64 {
	return s.Encrypt
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetEncryptType() *string {
	return s.EncryptType
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetFileURL() *string {
	return s.FileURL
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetFormat() *string {
	return s.Format
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetHDRType() *string {
	return s.HDRType
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetHeight() *int64 {
	return s.Height
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetJobId() *string {
	return s.JobId
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetNarrowBandType() *string {
	return s.NarrowBandType
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetPlayURL() *string {
	return s.PlayURL
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetSize() *int64 {
	return s.Size
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetStreamTags() *string {
	return s.StreamTags
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetStreamType() *string {
	return s.StreamType
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetTransTemplateType() *string {
	return s.TransTemplateType
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetWidth() *int64 {
	return s.Width
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetBitDepth(v int32) *GetPlayInfoResponseBodyPlayInfoList {
	s.BitDepth = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetBitrate(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetCreationTime(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetDefinition(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.Definition = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetDuration(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.Duration = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetEncrypt(v int64) *GetPlayInfoResponseBodyPlayInfoList {
	s.Encrypt = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetEncryptType(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.EncryptType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetFileURL(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.FileURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetFormat(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.Format = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetFps(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.Fps = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetHDRType(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.HDRType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetHeight(v int64) *GetPlayInfoResponseBodyPlayInfoList {
	s.Height = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetJobId(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.JobId = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetModificationTime(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.ModificationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetNarrowBandType(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.NarrowBandType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetPlayURL(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.PlayURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetSize(v int64) *GetPlayInfoResponseBodyPlayInfoList {
	s.Size = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetStatus(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.Status = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetStreamTags(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.StreamTags = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetStreamType(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.StreamType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetTransTemplateType(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.TransTemplateType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetWatermarkId(v string) *GetPlayInfoResponseBodyPlayInfoList {
	s.WatermarkId = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetWidth(v int64) *GetPlayInfoResponseBodyPlayInfoList {
	s.Width = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) Validate() error {
	return dara.Validate(s)
}
