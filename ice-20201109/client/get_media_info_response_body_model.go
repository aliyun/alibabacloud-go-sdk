// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfo(v *GetMediaInfoResponseBodyMediaInfo) *GetMediaInfoResponseBody
	GetMediaInfo() *GetMediaInfoResponseBodyMediaInfo
	SetRequestId(v string) *GetMediaInfoResponseBody
	GetRequestId() *string
}

type GetMediaInfoResponseBody struct {
	// The information about the media asset.
	MediaInfo *GetMediaInfoResponseBodyMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FDE2411-DB8D-4A9A-875B-275798F14A5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBody) GetMediaInfo() *GetMediaInfoResponseBodyMediaInfo {
	return s.MediaInfo
}

func (s *GetMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaInfoResponseBody) SetMediaInfo(v *GetMediaInfoResponseBodyMediaInfo) *GetMediaInfoResponseBody {
	s.MediaInfo = v
	return s
}

func (s *GetMediaInfoResponseBody) SetRequestId(v string) *GetMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaInfoResponseBody) Validate() error {
	if s.MediaInfo != nil {
		if err := s.MediaInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaInfoResponseBodyMediaInfo struct {
	// The original AI analysis data.
	AiRoughData *GetMediaInfoResponseBodyMediaInfoAiRoughData `json:"AiRoughData,omitempty" xml:"AiRoughData,omitempty" type:"Struct"`
	// The file information.
	FileInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information about the media asset.
	MediaBasicInfo *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfo) GetAiRoughData() *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	return s.AiRoughData
}

func (s *GetMediaInfoResponseBodyMediaInfo) GetFileInfoList() []*GetMediaInfoResponseBodyMediaInfoFileInfoList {
	return s.FileInfoList
}

func (s *GetMediaInfoResponseBodyMediaInfo) GetMediaBasicInfo() *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *GetMediaInfoResponseBodyMediaInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetAiRoughData(v *GetMediaInfoResponseBodyMediaInfoAiRoughData) *GetMediaInfoResponseBodyMediaInfo {
	s.AiRoughData = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetFileInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoList) *GetMediaInfoResponseBodyMediaInfo {
	s.FileInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetMediaBasicInfo(v *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) *GetMediaInfoResponseBodyMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetMediaId(v string) *GetMediaInfoResponseBodyMediaInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) Validate() error {
	if s.AiRoughData != nil {
		if err := s.AiRoughData.Validate(); err != nil {
			return err
		}
	}
	if s.FileInfoList != nil {
		for _, item := range s.FileInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MediaBasicInfo != nil {
		if err := s.MediaBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaInfoResponseBodyMediaInfoAiRoughData struct {
	// The AI category. Valid values:
	//
	// 	- Life
	//
	// 	- Good-looking
	//
	// 	- Cute pets
	//
	// 	- News
	//
	// 	- Ads
	//
	// 	- Environmental resources
	//
	// 	- Automobile
	AiCategory *string `json:"AiCategory,omitempty" xml:"AiCategory,omitempty"`
	// The ID of the AI task.
	//
	// example:
	//
	// ****483915d4f2cd8ac20b48fb04****
	AiJobId *string `json:"AiJobId,omitempty" xml:"AiJobId,omitempty"`
	// The analysis result.
	//
	// example:
	//
	// https://sample-bucket.cn-shanghai.aliyuncs.com/result.json
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The storage type. This parameter indicates the library in which the analysis data is stored. Valid values:
	//
	// 	- TEXT: the text library.
	//
	// example:
	//
	// TEXT
	SaveType *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
	// The information about the tagging job.
	StandardSmartTagJob *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob `json:"StandardSmartTagJob,omitempty" xml:"StandardSmartTagJob,omitempty" type:"Struct"`
	// The analysis status. Valid values:
	//
	// 	- Analyzing
	//
	// 	- AnalyzeSuccess
	//
	// 	- AnalyzeFailed
	//
	// 	- Saving
	//
	// 	- SaveSuccess
	//
	// 	- SaveFailed
	//
	// 	- Deleting
	//
	// 	- DeleteSuccess
	//
	// 	- DeleteFailed
	//
	// example:
	//
	// Analyzing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughData) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughData) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) GetAiCategory() *string {
	return s.AiCategory
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) GetAiJobId() *string {
	return s.AiJobId
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) GetResult() *string {
	return s.Result
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) GetSaveType() *string {
	return s.SaveType
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) GetStandardSmartTagJob() *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob {
	return s.StandardSmartTagJob
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) GetStatus() *string {
	return s.Status
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) SetAiCategory(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	s.AiCategory = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) SetAiJobId(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	s.AiJobId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) SetResult(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	s.Result = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) SetSaveType(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	s.SaveType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) SetStandardSmartTagJob(v *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	s.StandardSmartTagJob = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) SetStatus(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughData {
	s.Status = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughData) Validate() error {
	if s.StandardSmartTagJob != nil {
		if err := s.StandardSmartTagJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob struct {
	// The ID of the AI task.
	//
	// example:
	//
	// ****483915d4f2cd8ac20b48fb04****
	AiJobId *string `json:"AiJobId,omitempty" xml:"AiJobId,omitempty"`
	// The URL of the tagging result.
	//
	// example:
	//
	// http://xx.oss-cn-shanghai.aliyuncs.com/result2.txt
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// The recognized tags.
	Results []*GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The analysis status. Valid values:
	//
	// 	- Analyzing
	//
	// 	- AnalyzeSuccess
	//
	// 	- AnalyzeFailed
	//
	// example:
	//
	// Analyzing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) GetAiJobId() *string {
	return s.AiJobId
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) GetResults() []*GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults {
	return s.Results
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) GetStatus() *string {
	return s.Status
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) SetAiJobId(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob {
	s.AiJobId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) SetResultUrl(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob {
	s.ResultUrl = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) SetResults(v []*GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob {
	s.Results = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) SetStatus(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob {
	s.Status = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJob) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults struct {
	// The result data. The value is a JSON string. For information about the data structures of different data types<props="china">, see [Description of the Results parameter](https://help.aliyun.com/zh/ims/developer-reference/api-ice-2020-11-09-querysmarttagjob?spm=a2c4g.11186623.0.0.521d48b7KfapOL#api-detail-40).
	//
	// example:
	//
	// {"autoChapters": [...]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The tagging type. Valid values:
	//
	// 	- NLP: natural language processing (NLP)-based tagging
	//
	// example:
	//
	// NLP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) GetData() *string {
	return s.Data
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) GetType() *string {
	return s.Type
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) SetData(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults {
	s.Data = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) SetType(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults {
	s.Type = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataStandardSmartTagJobResults) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoResponseBodyMediaInfoFileInfoList struct {
	// The information about the audio tracks. A media asset may have multiple audio tracks.
	AudioStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic information about the file, including the duration and size.
	FileBasicInfo *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The information about the subtitle tracks. A media asset may have multiple subtitle tracks.
	SubtitleStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
	// The information about the video tracks. A media asset may have multiple video tracks.
	VideoStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) GetAudioStreamInfoList() []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) GetFileBasicInfo() *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) GetSubtitleStreamInfoList() []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	return s.SubtitleStreamInfoList
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) GetVideoStreamInfoList() []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetAudioStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetFileBasicInfo(v *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetSubtitleStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetVideoStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) Validate() error {
	if s.AudioStreamInfoList != nil {
		for _, item := range s.AudioStreamInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleStreamInfoList != nil {
		for _, item := range s.SubtitleStreamInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoStreamInfoList != nil {
		for _, item := range s.VideoStreamInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList struct {
	// The bitrate.
	//
	// example:
	//
	// 127.794
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The output layout of sound channels.
	//
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the codec.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the codec.
	//
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the codec.
	//
	// example:
	//
	// mp4a
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the codec.
	//
	// example:
	//
	// 1/24000
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration.
	//
	// example:
	//
	// 16.200998
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The audio frame rate.
	//
	// example:
	//
	// 8
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The sequence number of the audio track.
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
	// The number of frames.
	//
	// example:
	//
	// 10
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The codec profile.
	//
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The sampling format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/44100
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannels(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetFps(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetNumFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetProfile(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleRate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo struct {
	// The bitrate.
	//
	// example:
	//
	// 1132.68
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the file was created.
	//
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The duration.
	//
	// example:
	//
	// 216.206667
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// example:
	//
	// example.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 30611502
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file status.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// The file type.
	//
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The OSS URL of the file.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The container format.
	//
	// example:
	//
	// mov,mp4,m4a,3gp,3g2,mj2
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height.
	//
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region in which the file is stored.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The width.
	//
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetCreateTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileSize(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileStatus(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileType(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileUrl(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFormatName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetHeight(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetModifiedTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetRegion(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetWidth(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList struct {
	// The full name of the codec.
	//
	// example:
	//
	// SubRip Text
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	//
	// example:
	//
	// srt
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the codec.
	//
	// example:
	//
	// unicode
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the codec.
	//
	// example:
	//
	// unicode
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the codec.
	//
	// example:
	//
	// 29.97
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the subtitle track.
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
	// The start time.
	//
	// example:
	//
	// 0
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 30
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList struct {
	// The average video frame rate.
	//
	// example:
	//
	// 24.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate.
	//
	// example:
	//
	// 1001.594
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the codec.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	//
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the codec.
	//
	// example:
	//
	// 0x0000
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the codec.
	//
	// example:
	//
	// [0][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the codec.
	//
	// example:
	//
	// 1/48
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR).
	//
	// example:
	//
	// 0:1
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration.
	//
	// example:
	//
	// 216.206706
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The video frame rate.
	//
	// example:
	//
	// 24.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Indicates whether the video track contains bidirectional frames (B-frames).
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The height.
	//
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video track.
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
	// 5184
	NbFrames *string `json:"Nb_frames,omitempty" xml:"Nb_frames,omitempty"`
	// The number of frames.
	//
	// example:
	//
	// 5184
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
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The rotation angle.
	//
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	//
	// example:
	//
	// 0:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.081706
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/12288
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The width.
	//
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNbFrames() *string {
	return s.NbFrames
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDar(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetFps(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHeight(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLevel(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNbFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNumFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetPixFmt(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetProfile(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetRotate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetSar(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetWidth(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoResponseBodyMediaInfoMediaBasicInfo struct {
	// The service to which the media asset belongs.
	//
	// example:
	//
	// ICE
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The business type.
	//
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 3048
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// cateName
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The category.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created.
	//
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the media asset was deleted.
	//
	// example:
	//
	// 2020-12-26T04:11:15Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The content description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input URL of the media asset in another service.
	//
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags.
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time when the media asset was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The custom ID. The ID can be 6 to 64 characters in length and can contain only letters, digits, hyphens (-), and underscores (_). The ID is unique among users.
	//
	// example:
	//
	// 123-1234
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The snapshots.
	//
	// example:
	//
	// [
	//
	//     "http://example-bucket.oss-cn-shanghai.aliyuncs.com/snapshot-00001.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>",
	//
	//     "http://example-bucket.oss-cn-shanghai.aliyuncs.com/snapshot-00002.jpg?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>",
	//
	//     "http://example-bucket.oss-cn-shanghai.aliyuncs.com/snapshot-00003.jpg?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>"
	//
	// ]
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source.
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The sprite.
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The resource status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The upload source of the media asset.
	//
	// example:
	//
	// general
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// The user data.
	//
	// example:
	//
	// userDataTest
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCateName() *string {
	return s.CateName
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetUploadSource() *string {
	return s.UploadSource
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBiz(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBusinessType(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCateId(v int64) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CateId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCateName(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CateName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCategory(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCoverURL(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCreateTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDeletedTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDescription(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetInputURL(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaId(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaTags(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaType(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetModifiedTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetReferenceId(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ReferenceId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSnapshots(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSource(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSpriteImages(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetStatus(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetTitle(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUploadSource(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UploadSource = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUserData(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
