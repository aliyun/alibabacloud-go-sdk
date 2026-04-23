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
	BitDepth         *int32  `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	Bitrate          *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecName        *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Definition       *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Duration         *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Encrypt          *int64  `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	EncryptMode      *string `json:"EncryptMode,omitempty" xml:"EncryptMode,omitempty"`
	EncryptType      *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Format           *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Fps              *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	HDRType          *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	Height           *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	JobExt           *string `json:"JobExt,omitempty" xml:"JobExt,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType          *int32  `json:"JobType,omitempty" xml:"JobType,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	NarrowBandType   *string `json:"NarrowBandType,omitempty" xml:"NarrowBandType,omitempty"`
	PlayURL          *string `json:"PlayURL,omitempty" xml:"PlayURL,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Specification    *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StreamType       *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	WatermarkId      *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	Width            *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
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
