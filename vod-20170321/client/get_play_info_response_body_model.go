// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlayInfoList(v *GetPlayInfoResponseBodyPlayInfoList) *GetPlayInfoResponseBody
	GetPlayInfoList() *GetPlayInfoResponseBodyPlayInfoList
	SetRequestId(v string) *GetPlayInfoResponseBody
	GetRequestId() *string
	SetVideoBase(v *GetPlayInfoResponseBodyVideoBase) *GetPlayInfoResponseBody
	GetVideoBase() *GetPlayInfoResponseBodyVideoBase
}

type GetPlayInfoResponseBody struct {
	// The information about the audio or video stream.
	PlayInfoList *GetPlayInfoResponseBodyPlayInfoList `json:"PlayInfoList,omitempty" xml:"PlayInfoList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F552E596-967D-5500-842F-17E6364****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The basic information about the audio or video file.
	VideoBase *GetPlayInfoResponseBodyVideoBase `json:"VideoBase,omitempty" xml:"VideoBase,omitempty" type:"Struct"`
}

func (s GetPlayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBody) GetPlayInfoList() *GetPlayInfoResponseBodyPlayInfoList {
	return s.PlayInfoList
}

func (s *GetPlayInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlayInfoResponseBody) GetVideoBase() *GetPlayInfoResponseBodyVideoBase {
	return s.VideoBase
}

func (s *GetPlayInfoResponseBody) SetPlayInfoList(v *GetPlayInfoResponseBodyPlayInfoList) *GetPlayInfoResponseBody {
	s.PlayInfoList = v
	return s
}

func (s *GetPlayInfoResponseBody) SetRequestId(v string) *GetPlayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlayInfoResponseBody) SetVideoBase(v *GetPlayInfoResponseBodyVideoBase) *GetPlayInfoResponseBody {
	s.VideoBase = v
	return s
}

func (s *GetPlayInfoResponseBody) Validate() error {
	if s.PlayInfoList != nil {
		if err := s.PlayInfoList.Validate(); err != nil {
			return err
		}
	}
	if s.VideoBase != nil {
		if err := s.VideoBase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPlayInfoResponseBodyPlayInfoList struct {
	PlayInfo []*GetPlayInfoResponseBodyPlayInfoListPlayInfo `json:"PlayInfo,omitempty" xml:"PlayInfo,omitempty" type:"Repeated"`
}

func (s GetPlayInfoResponseBodyPlayInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBodyPlayInfoList) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyPlayInfoList) GetPlayInfo() []*GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	return s.PlayInfo
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetPlayInfo(v []*GetPlayInfoResponseBodyPlayInfoListPlayInfo) *GetPlayInfoResponseBodyPlayInfoList {
	s.PlayInfo = v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoList) Validate() error {
	if s.PlayInfo != nil {
		for _, item := range s.PlayInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPlayInfoResponseBodyPlayInfoListPlayInfo struct {
	// The color depth. This value is an integer.
	//
	// example:
	//
	// 8
	BitDepth *int32 `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	// The bitrate of the media stream. Unit: Kbit/s.
	//
	// example:
	//
	// 450.878
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The encoding type. The possible values are:
	//
	// - H264
	//
	// - H265
	//
	// example:
	//
	// H264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The time when the audio or video stream was created. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-18T07:37:15Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The quality of the video stream. Valid values:
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
	// LD
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
	// 1
	Encrypt *int64 `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// The encryption type of the media stream. Valid values:
	//
	// 	- **License**: decryption on local devices
	//
	// >  If the encryption type is **License**, only ApsaraVideo Player SDK can be used to play videos.
	//
	// example:
	//
	// License
	EncryptMode *string `json:"EncryptMode,omitempty" xml:"EncryptMode,omitempty"`
	// The encryption type of the media stream. Valid values:
	//
	// 	- **AliyunVoDEncryption**: Alibaba Cloud proprietary cryptography
	//
	// 	- **HLSEncryption**: HTTP-Live-Streaming (HLS) encryption
	//
	// >  If the encryption type is AliyunVoDEncryption, only ApsaraVideo Player SDK can be used to play videos.
	//
	// example:
	//
	// AliyunVoDEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The format of the media stream.
	//
	// 	- If the media file is a video file, the valid values are **mp4*	- and **m3u8**.
	//
	// 	- If the media asset is an audio-only file, the value is **mp3**.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the media stream. Unit: frames per second.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The HDR type of the media stream. Valid values:
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
	// HLG
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The height of the media stream. Unit: pixels.
	//
	// example:
	//
	// 640
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The custom watermark information of the copyright watermark. This parameter is returned if you set `JobType` to `2`.
	//
	// example:
	//
	// CopyrightMarkTest
	JobExt *string `json:"JobExt,omitempty" xml:"JobExt,omitempty"`
	// The job ID for transcoding the media stream. This ID uniquely identifies a media stream.
	//
	// example:
	//
	// 80e9c6580e754a798c3c19c59b16****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the digital watermark. Valid values:
	//
	// 	- **1**: user-tracing watermark
	//
	// 	- **2**: copyright watermark
	//
	// example:
	//
	// 2
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the audio or video file was last updated. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-20T06:32:19Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The transcoding type. Valid values:
	//
	// 	- **0**: regular transcoding
	//
	// 	- **1.0**: Narrowband HD™ 1.0 transcoding
	//
	// 	- **2.0**: Narrowband HD™ 2.0 transcoding
	//
	// example:
	//
	// 0
	NarrowBandType *string `json:"NarrowBandType,omitempty" xml:"NarrowBandType,omitempty"`
	// The playback URL of the video stream.
	PlayURL *string `json:"PlayURL,omitempty" xml:"PlayURL,omitempty"`
	// The size of the media stream. Unit: bytes.
	//
	// example:
	//
	// 418112
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The specifications of transcoded audio and video streams. For more information about the valid values, see [Output specifications](~~124671#section-6bv-l0g-opq~~).
	//
	// example:
	//
	// H264.LD
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The status of the audio or video stream. Valid values:
	//
	// 	- **Normal**: The latest transcoded stream in each quality and format is in the Normal status.
	//
	// 	- **Invisible**: If multiple streams are transcoded in the same quality and format, the latest transcoded stream is in the Normal status and other streams are in the Invisible status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the media stream. If the media stream is a video stream, the value is **video**. If the media stream is an audio-only stream, the value is **audio**.
	//
	// example:
	//
	// video
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the watermark that is associated with the media stream.
	//
	// example:
	//
	// dgfn26457856****
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	// The width of the media stream. Unit: pixels.
	//
	// example:
	//
	// 360
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetPlayInfoResponseBodyPlayInfoListPlayInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBodyPlayInfoListPlayInfo) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetBitDepth() *int32 {
	return s.BitDepth
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetCodecName() *string {
	return s.CodecName
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetDefinition() *string {
	return s.Definition
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetEncrypt() *int64 {
	return s.Encrypt
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetEncryptMode() *string {
	return s.EncryptMode
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetEncryptType() *string {
	return s.EncryptType
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetFormat() *string {
	return s.Format
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetFps() *string {
	return s.Fps
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetHDRType() *string {
	return s.HDRType
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetHeight() *int64 {
	return s.Height
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetJobExt() *string {
	return s.JobExt
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetJobId() *string {
	return s.JobId
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetJobType() *int32 {
	return s.JobType
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetNarrowBandType() *string {
	return s.NarrowBandType
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetPlayURL() *string {
	return s.PlayURL
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetSize() *int64 {
	return s.Size
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetSpecification() *string {
	return s.Specification
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetStreamType() *string {
	return s.StreamType
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) GetWidth() *int64 {
	return s.Width
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetBitDepth(v int32) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.BitDepth = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetBitrate(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Bitrate = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetCodecName(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.CodecName = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetCreationTime(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.CreationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetDefinition(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Definition = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetDuration(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Duration = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetEncrypt(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Encrypt = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetEncryptMode(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.EncryptMode = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetEncryptType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.EncryptType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetFormat(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Format = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetFps(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Fps = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetHDRType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.HDRType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetHeight(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Height = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetJobExt(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.JobExt = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetJobId(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.JobId = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetJobType(v int32) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.JobType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetModificationTime(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.ModificationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetNarrowBandType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.NarrowBandType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetPlayURL(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.PlayURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetSize(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Size = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetSpecification(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Specification = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetStatus(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Status = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetStreamType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.StreamType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetWatermarkId(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.WatermarkId = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetWidth(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Width = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) Validate() error {
	return dara.Validate(s)
}

type GetPlayInfoResponseBodyVideoBase struct {
	// The thumbnail URL of the audio or video file.
	//
	// example:
	//
	// http://example.aliyundoc.com/sample.jpg?auth_key=2333232-atb****
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the audio or video file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-06-26T06:38:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the masked live comment data.
	//
	// example:
	//
	// http://example.aliyundoc.com/****?auth_key=abdf2123-6783232****
	DanMuURL *string `json:"DanMuURL,omitempty" xml:"DanMuURL,omitempty"`
	// The duration of the audio or video file. Unit: seconds.
	//
	// example:
	//
	// 3.1667
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The type of the media file. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The status of the media file. For more information about the value range and description, see the [Status](~~52839#title-vqg-8cz-7p8~~) table.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the audio file. Valid values:
	//
	// 	- **Standard**: All media resources are stored as Standard objects.
	//
	// 	- **IA**: All media resources are stored as IA objects.
	//
	// 	- **Archive**: All media resources are stored as Archive objects.
	//
	// 	- **ColdArchive**: All media resources are stored as Cold Archive objects.
	//
	// 	- **SourceIA**: Only the source files are IA objects.
	//
	// 	- **SourceArchive**: Only the source files are Archive objects.
	//
	// 	- **SourceColdArchive**: Only the source file is stored as a Cold Archive object.
	//
	// 	- **Changing**: The storage class of the video file is being changed.
	//
	// 	- **SourceChanging**: The storage class of the source file is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The title of the audio or video file.
	//
	// example:
	//
	// ApsaraVideo VOD
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// 93ab850b4f654b6e91d24d81d44****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetPlayInfoResponseBodyVideoBase) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponseBodyVideoBase) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyVideoBase) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetPlayInfoResponseBodyVideoBase) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetPlayInfoResponseBodyVideoBase) GetDanMuURL() *string {
	return s.DanMuURL
}

func (s *GetPlayInfoResponseBodyVideoBase) GetDuration() *string {
	return s.Duration
}

func (s *GetPlayInfoResponseBodyVideoBase) GetMediaType() *string {
	return s.MediaType
}

func (s *GetPlayInfoResponseBodyVideoBase) GetStatus() *string {
	return s.Status
}

func (s *GetPlayInfoResponseBodyVideoBase) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetPlayInfoResponseBodyVideoBase) GetTitle() *string {
	return s.Title
}

func (s *GetPlayInfoResponseBodyVideoBase) GetVideoId() *string {
	return s.VideoId
}

func (s *GetPlayInfoResponseBodyVideoBase) SetCoverURL(v string) *GetPlayInfoResponseBodyVideoBase {
	s.CoverURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetCreationTime(v string) *GetPlayInfoResponseBodyVideoBase {
	s.CreationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetDanMuURL(v string) *GetPlayInfoResponseBodyVideoBase {
	s.DanMuURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetDuration(v string) *GetPlayInfoResponseBodyVideoBase {
	s.Duration = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetMediaType(v string) *GetPlayInfoResponseBodyVideoBase {
	s.MediaType = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetStatus(v string) *GetPlayInfoResponseBodyVideoBase {
	s.Status = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetStorageClass(v string) *GetPlayInfoResponseBodyVideoBase {
	s.StorageClass = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetTitle(v string) *GetPlayInfoResponseBodyVideoBase {
	s.Title = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetVideoId(v string) *GetPlayInfoResponseBodyVideoBase {
	s.VideoId = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) Validate() error {
	return dara.Validate(s)
}
