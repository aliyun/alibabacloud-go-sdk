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
	// The basic information about the media asset.
	MediaBase *GetPlayInfoResponseBodyMediaBase `json:"MediaBase,omitempty" xml:"MediaBase,omitempty" type:"Struct"`
	// A list of audio or video playback streams.
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
	if s.MediaBase != nil {
		if err := s.MediaBase.Validate(); err != nil {
			return err
		}
	}
	if s.PlayInfoList != nil {
		for _, item := range s.PlayInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPlayInfoResponseBodyMediaBase struct {
	// The category ID. You can obtain the category ID in one of the following ways:
	//
	// - Log on to the [IMS console](https://ims.console.aliyun.com) and choose **media asset management*	- > **category management*	- to view the category ID.
	//
	// - The create category operation returns the category ID in the `CateId` parameter.
	//
	// - The get category operation returns the category ID in the `CateId` parameter.
	//
	// example:
	//
	// 4220
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The cover URL.
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
	// The description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// 2eea77a61c7b4ddd95bec34a6f65b***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags.
	//
	// - You can add up to 16 tags.
	//
	// - Separate multiple tags with commas (,).
	//
	// - The maximum length of a tag is 32 bytes.
	//
	// - Tags must be UTF-8 encoded.
	//
	// example:
	//
	// test,ccc
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media file. Valid values:
	//
	// `video`: A video file. `audio`: An audio-only file.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The status of the media asset. Valid values:
	//
	// - `Init`: The source file is not ready.
	//
	// - `Preparing`: The source file is being prepared. This process may involve uploading or compositing.
	//
	// - `PrepareFail`: Preparation of the source file failed. For example, the system failed to retrieve the source file metadata.
	//
	// - `Normal`: The source file is ready.
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
	// The color bit depth.
	//
	// example:
	//
	// 8
	BitDepth *int32 `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	// The bitrate of the media stream in Kbit/s.
	//
	// example:
	//
	// 20
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The creation time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-10T02:28:49Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The definition of the video stream. Valid values:
	//
	// - **FD**: fluent
	//
	// - **LD**: standard definition
	//
	// - **SD**: high definition
	//
	// - **HD**: ultra-high definition
	//
	// - **OD**: original
	//
	// - **2K**
	//
	// - **4K**
	//
	// - **SQ**: standard-quality audio
	//
	// - **HQ**: high-quality audio
	//
	// - **AUTO**: adaptive bitrate
	//
	// example:
	//
	// HD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The duration of the media stream in seconds.
	//
	// example:
	//
	// 9.0464
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether the media stream is encrypted. Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	// example:
	//
	// 0
	Encrypt *int64 `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// The encryption type of the media stream. Valid values:
	//
	// - **AliyunVoDEncryption**: Alibaba Cloud VoD Encryption.
	//
	// - **HLSEncryption**: HLS standard encryption.
	//
	// > If a stream is encrypted with **AliyunVoDEncryption**, you can play it only with the Alibaba Cloud Player SDK.
	//
	// example:
	//
	// AliyunVoDEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The OSS file URL.
	//
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/sv/43a68ee9-181809b6aba/43a68ee9-181809b6aba.mpeg
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The format of the media stream.
	//
	// - For video streams, valid values are **mp4*	- and **m3u8**.
	//
	// - For audio-only streams, the value is **mp3**.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the media stream in frames per second.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The High Dynamic Range (HDR) type of the media stream. Valid values:
	//
	// - HDR
	//
	// - HDR10
	//
	// - HLG
	//
	// - DolbyVision
	//
	// - HDRVivid
	//
	// - SDR+
	//
	// example:
	//
	// HDR
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The height of the media stream in pixels.
	//
	// example:
	//
	// 1080
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 36c9d38e70bf43ed9f7f8f48d6356***
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The last modification time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-13T11:39:41.714+08:00
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The Narrowband HD type. Valid values:
	//
	// - **0**: regular.
	//
	// - **1.0**: Narrowband HD 1.0.
	//
	// - **2.0**: Narrowband HD 2.0.
	//
	// This parameter applies only if a definition is configured in the built-in transcoding template for Narrowband HD 1.0. For more information, see [Configure transcoding templates - Definition](https://help.aliyun.com/document_detail/52839.html).
	//
	// example:
	//
	// 0
	NarrowBandType *string `json:"NarrowBandType,omitempty" xml:"NarrowBandType,omitempty"`
	// The playback URL of the video stream.
	//
	// example:
	//
	// https://***.aliyuncdn.com/sv/756bee1-17f980f0945/756bee1-17f980f0945.mp4
	PlayURL *string `json:"PlayURL,omitempty" xml:"PlayURL,omitempty"`
	// The size of the media stream in bytes.
	//
	// example:
	//
	// 418112
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The media stream status. Valid values:
	//
	// - **Normal**: The stream is available.
	//
	// - **Invisible**: The stream is not visible.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The stream tags, which are used to identify the transcoding type.
	//
	// example:
	//
	// "{\\"ims.audioServiceType\\": \\"AudioEnhancement\\"}"
	StreamTags *string `json:"StreamTags,omitempty" xml:"StreamTags,omitempty"`
	// The type of the media stream. The value is **video*	- for video streams or **audio*	- for audio-only streams.
	//
	// example:
	//
	// video
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The type of the transcoding template. Valid values:
	//
	// - `Normal`: regular transcoding
	//
	// - `AudioTranscode`: audio transcoding
	//
	// - `Remux`: remuxing
	//
	// - `NarrowBandV1`: Narrowband HD 1.0
	//
	// - `NarrowBandV2`: Narrowband HD 2.0
	//
	// - `UHD`: audio and video enhancement (ultra-high definition)
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
	// The width of the media stream in pixels.
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
