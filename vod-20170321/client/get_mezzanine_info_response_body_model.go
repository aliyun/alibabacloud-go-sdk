// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMezzanineInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMezzanine(v *GetMezzanineInfoResponseBodyMezzanine) *GetMezzanineInfoResponseBody
	GetMezzanine() *GetMezzanineInfoResponseBodyMezzanine
	SetRequestId(v string) *GetMezzanineInfoResponseBody
	GetRequestId() *string
}

type GetMezzanineInfoResponseBody struct {
	// The information about the source file.
	Mezzanine *GetMezzanineInfoResponseBodyMezzanine `json:"Mezzanine,omitempty" xml:"Mezzanine,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMezzanineInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBody) GetMezzanine() *GetMezzanineInfoResponseBodyMezzanine {
	return s.Mezzanine
}

func (s *GetMezzanineInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMezzanineInfoResponseBody) SetMezzanine(v *GetMezzanineInfoResponseBodyMezzanine) *GetMezzanineInfoResponseBody {
	s.Mezzanine = v
	return s
}

func (s *GetMezzanineInfoResponseBody) SetRequestId(v string) *GetMezzanineInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMezzanineInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMezzanineInfoResponseBodyMezzanine struct {
	// The codec time base.
	AudioStreamList []*GetMezzanineInfoResponseBodyMezzanineAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Repeated"`
	// The bitrate of the file. Unit: Kbit/s.
	//
	// example:
	//
	// 771.2280
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The duration of the file. Unit: seconds.
	//
	// example:
	//
	// 42.4930
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 27ffc438-164h67f57ef-0005-6884-51a-1****.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The URL of the file.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-shanghai.aliyuncs.com/27ffc438-164h67f57ef-0005-6884-51a-1****.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The frame rate of the file. Unit: frames per second.
	//
	// example:
	//
	// 25.0000
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the file. Unit: pixel.
	//
	// example:
	//
	// 540
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The type of the mezzanine file URL. Valid values:
	//
	// - **oss**: OSS URL
	//
	// - **cdn*	- (default): CDN URL
	//
	// > If you specify an OSS URL for the video stream, the video stream must be in the MP4 format.
	//
	// example:
	//
	// oss
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The preprocess status od the media.
	//
	// example:
	//
	// UnPreprocess
	PreprocessStatus *string `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	// The period of time in which the object remains in the restored state.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the audio or video file. Valid values:
	//
	// 	- **Processing**
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The size of the file. Unit: byte.
	//
	// example:
	//
	// 4096477
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the file. Valid values:
	//
	// 	- **Uploading**: The file is being uploaded. This is the initial status.
	//
	// 	- **Normal**: The file is uploaded.
	//
	// 	- **UploadFail**: The file fails to be uploaded.
	//
	// 	- **Deleted**: The file is deleted.
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
	// 	- **SourceColdArchive**: Only the source files are Cold Archive objects.
	//
	// 	- **Changing**: The storage class of the audio file is being changed.
	//
	// 	- **SourceChanging**: The storage class of the source file is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 1f1a6fc03ca04814031b8a6559e****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// The HDR type of the video stream.
	VideoStreamList []*GetMezzanineInfoResponseBodyMezzanineVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Repeated"`
	// The width of the file. Unit: pixel.
	//
	// example:
	//
	// 960
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMezzanineInfoResponseBodyMezzanine) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoResponseBodyMezzanine) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetAudioStreamList() []*GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	return s.AudioStreamList
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetDuration() *string {
	return s.Duration
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetFileName() *string {
	return s.FileName
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetFileURL() *string {
	return s.FileURL
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetFps() *string {
	return s.Fps
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetHeight() *int64 {
	return s.Height
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetOutputType() *string {
	return s.OutputType
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetPreprocessStatus() *string {
	return s.PreprocessStatus
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetSize() *int64 {
	return s.Size
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetStatus() *string {
	return s.Status
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetVideoId() *string {
	return s.VideoId
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetVideoStreamList() []*GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	return s.VideoStreamList
}

func (s *GetMezzanineInfoResponseBodyMezzanine) GetWidth() *int64 {
	return s.Width
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetAudioStreamList(v []*GetMezzanineInfoResponseBodyMezzanineAudioStreamList) *GetMezzanineInfoResponseBodyMezzanine {
	s.AudioStreamList = v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetBitrate(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Bitrate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetCreationTime(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.CreationTime = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetDuration(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Duration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetFileName(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.FileName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetFileURL(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.FileURL = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetFps(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Fps = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetHeight(v int64) *GetMezzanineInfoResponseBodyMezzanine {
	s.Height = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetOutputType(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.OutputType = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetPreprocessStatus(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.PreprocessStatus = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetRestoreExpiration(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.RestoreExpiration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetRestoreStatus(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.RestoreStatus = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetSize(v int64) *GetMezzanineInfoResponseBodyMezzanine {
	s.Size = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetStatus(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Status = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetStorageClass(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.StorageClass = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetVideoId(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.VideoId = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetVideoStreamList(v []*GetMezzanineInfoResponseBodyMezzanineVideoStreamList) *GetMezzanineInfoResponseBodyMezzanine {
	s.VideoStreamList = v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetWidth(v int64) *GetMezzanineInfoResponseBodyMezzanine {
	s.Width = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) Validate() error {
	return dara.Validate(s)
}

type GetMezzanineInfoResponseBodyMezzanineAudioStreamList struct {
	// The bitrate.
	//
	// example:
	//
	// 62.885
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The output layout of the sound channels. Valid values:
	//
	// 	- **mono**
	//
	// 	- **stereo**
	//
	// example:
	//
	// mono
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 1
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// mp4a
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration of the audio file.
	//
	// example:
	//
	// 3.227574
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 1
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The sampling format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sampling rate of the audio stream.
	//
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The beginning of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base of the audio stream.
	//
	// example:
	//
	// 0.000000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetMezzanineInfoResponseBodyMezzanineAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetChannels() *string {
	return s.Channels
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetDuration() *string {
	return s.Duration
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetIndex() *string {
	return s.Index
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetLang() *string {
	return s.Lang
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetBitrate(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Bitrate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetChannelLayout(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetChannels(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Channels = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecLongName(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecLongName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecName(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecTag(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecTag = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecTagString(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecTagString = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecTimeBase(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetDuration(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Duration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetIndex(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Index = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetLang(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Lang = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetNumFrames(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.NumFrames = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetSampleFmt(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.SampleFmt = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetSampleRate(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.SampleRate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetStartTime(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.StartTime = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetTimebase(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Timebase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) Validate() error {
	return dara.Validate(s)
}

type GetMezzanineInfoResponseBodyMezzanineVideoStreamList struct {
	// The average frame rate.
	//
	// example:
	//
	// 30.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format.
	//
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x31637661
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// avc1
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/60
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR) of the video stream.
	//
	// example:
	//
	// 0:1
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration of the audio file.
	//
	// example:
	//
	// 3.166667
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate of the output file.
	//
	// example:
	//
	// 30.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The HDR type of the video stream.
	//
	// example:
	//
	// HDR
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// Indicates whether the video stream contains B-frames.
	//
	// example:
	//
	// 0
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The height of the video stream.
	//
	// example:
	//
	// 320
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video stream. The value indicates the position of the video stream in all video streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The codec level.
	//
	// example:
	//
	// 30
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 0
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The codec profile.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The rotation angle of the video. Valid values: **[0,360)**.
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR) of the video stream.
	//
	// example:
	//
	// 0:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The beginning of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base of the audio stream.
	//
	// example:
	//
	// 0.000000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The width of the video in pixels.
	//
	// example:
	//
	// 568
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMezzanineInfoResponseBodyMezzanineVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetDar() *string {
	return s.Dar
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetDuration() *string {
	return s.Duration
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetFps() *string {
	return s.Fps
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetHDRType() *string {
	return s.HDRType
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetHeight() *string {
	return s.Height
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetIndex() *string {
	return s.Index
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetLang() *string {
	return s.Lang
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetLevel() *string {
	return s.Level
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetProfile() *string {
	return s.Profile
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetRotate() *string {
	return s.Rotate
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetSar() *string {
	return s.Sar
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GetWidth() *string {
	return s.Width
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetAvgFPS(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.AvgFPS = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetBitrate(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Bitrate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecLongName(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecLongName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecName(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecTag(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecTag = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecTagString(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecTagString = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecTimeBase(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetDar(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Dar = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetDuration(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Duration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetFps(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Fps = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetHDRType(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.HDRType = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetHasBFrames(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.HasBFrames = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetHeight(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Height = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetIndex(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Index = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetLang(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Lang = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetLevel(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Level = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetNumFrames(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.NumFrames = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetPixFmt(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.PixFmt = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetProfile(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Profile = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetRotate(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Rotate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetSar(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Sar = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetStartTime(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.StartTime = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetTimebase(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Timebase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetWidth(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Width = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) Validate() error {
	return dara.Validate(s)
}
