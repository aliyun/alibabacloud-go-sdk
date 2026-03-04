// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAssetMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfo(v *GetYikeAssetMediaInfoResponseBodyMediaInfo) *GetYikeAssetMediaInfoResponseBody
	GetMediaInfo() *GetYikeAssetMediaInfoResponseBodyMediaInfo
	SetRequestId(v string) *GetYikeAssetMediaInfoResponseBody
	GetRequestId() *string
}

type GetYikeAssetMediaInfoResponseBody struct {
	MediaInfo *GetYikeAssetMediaInfoResponseBodyMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBody) GetMediaInfo() *GetYikeAssetMediaInfoResponseBodyMediaInfo {
	return s.MediaInfo
}

func (s *GetYikeAssetMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeAssetMediaInfoResponseBody) SetMediaInfo(v *GetYikeAssetMediaInfoResponseBodyMediaInfo) *GetYikeAssetMediaInfoResponseBody {
	s.MediaInfo = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBody) SetRequestId(v string) *GetYikeAssetMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBody) Validate() error {
	if s.MediaInfo != nil {
		if err := s.MediaInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeAssetMediaInfoResponseBodyMediaInfo struct {
	BizData        *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData        `json:"BizData,omitempty" xml:"BizData,omitempty" type:"Struct"`
	FileInfoList   []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	MediaBasicInfo *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfo) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) GetBizData() *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	return s.BizData
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) GetFileInfoList() []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList {
	return s.FileInfoList
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) GetMediaBasicInfo() *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) SetBizData(v *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) *GetYikeAssetMediaInfoResponseBodyMediaInfo {
	s.BizData = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) SetFileInfoList(v []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) *GetYikeAssetMediaInfoResponseBodyMediaInfo {
	s.FileInfoList = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) SetMediaBasicInfo(v *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) *GetYikeAssetMediaInfoResponseBodyMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) SetMediaId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfo {
	s.MediaId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfo) Validate() error {
	if s.BizData != nil {
		if err := s.BizData.Validate(); err != nil {
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

type GetYikeAssetMediaInfoResponseBodyMediaInfoBizData struct {
	// example:
	//
	// Label
	AuditBlockedLabel *string `json:"AuditBlockedLabel,omitempty" xml:"AuditBlockedLabel,omitempty"`
	// example:
	//
	// Status
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// ID
	CreationJobId *string `json:"CreationJobId,omitempty" xml:"CreationJobId,omitempty"`
	// example:
	//
	// fd-CBMEJaa2fA
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// 1
	IsFavorite *string `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// example:
	//
	// 1
	IsLogicalDeleted *string `json:"IsLogicalDeleted,omitempty" xml:"IsLogicalDeleted,omitempty"`
	// example:
	//
	// SubType
	MediaAssetSubType *string `json:"MediaAssetSubType,omitempty" xml:"MediaAssetSubType,omitempty"`
	// example:
	//
	// Type
	MediaAssetType *string `json:"MediaAssetType,omitempty" xml:"MediaAssetType,omitempty"`
	// example:
	//
	// ID
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	Prompt       *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// f4a26390f02371f0a1f4e6e7c758****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// SourceName
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// example:
	//
	// MainBody
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetAuditBlockedLabel() *string {
	return s.AuditBlockedLabel
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetCreationJobId() *string {
	return s.CreationJobId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetFolderId() *string {
	return s.FolderId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetIsFavorite() *string {
	return s.IsFavorite
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetIsLogicalDeleted() *string {
	return s.IsLogicalDeleted
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetMediaAssetSubType() *string {
	return s.MediaAssetSubType
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetMediaAssetType() *string {
	return s.MediaAssetType
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetProductionId() *string {
	return s.ProductionId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetPrompt() *string {
	return s.Prompt
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetSourceId() *string {
	return s.SourceId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetSourceName() *string {
	return s.SourceName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) GetSourceType() *string {
	return s.SourceType
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetAuditBlockedLabel(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.AuditBlockedLabel = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetAuditStatus(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.AuditStatus = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetCreationJobId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.CreationJobId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetFolderId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.FolderId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetIsFavorite(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.IsFavorite = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetIsLogicalDeleted(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.IsLogicalDeleted = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetMediaAssetSubType(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.MediaAssetSubType = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetMediaAssetType(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.MediaAssetType = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetProductionId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.ProductionId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetPrompt(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.Prompt = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetSourceId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.SourceId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetSourceName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.SourceName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) SetSourceType(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData {
	s.SourceType = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoBizData) Validate() error {
	return dara.Validate(s)
}

type GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList struct {
	AudioStreamInfoList    []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList    `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	FileBasicInfo          *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo            `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	SubtitleStreamInfoList []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
	VideoStreamInfoList    []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList    `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) GetAudioStreamInfoList() []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) GetFileBasicInfo() *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) GetSubtitleStreamInfoList() []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	return s.SubtitleStreamInfoList
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) GetVideoStreamInfoList() []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) SetAudioStreamInfoList(v []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) SetFileBasicInfo(v *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) SetSubtitleStreamInfoList(v []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) SetVideoStreamInfoList(v []*GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoList) Validate() error {
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

type GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList struct {
	// example:
	//
	// 192.0
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// mp4a
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 16.2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 10
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 2026-02-04T02:13:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1/44100
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetBitrate(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannels(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTag(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetDuration(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetFps(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetIndex(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetLang(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetNumFrames(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetProfile(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleRate(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetStartTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetTimebase(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo struct {
	// example:
	//
	// 20
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 16.2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// example.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 27007
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// mp4
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetBitrate(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetCreateTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetDuration(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileSize(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileStatus(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileType(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileUrl(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFormatName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetHeight(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetModifiedTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetRegion(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetWidth(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList struct {
	// example:
	//
	// SubRip Text
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// srt
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// unicode
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// unicode
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 29.97
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 2025-03-07T01:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 30
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetDuration(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetIndex(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetLang(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList struct {
	// example:
	//
	// 24.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// example:
	//
	// 20
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// 0x0000
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// [0][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 1/48
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 0:1
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// example:
	//
	// 216.206706
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 24.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 2
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 30
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 5184
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// example:
	//
	// 0:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// example:
	//
	// 0.081706
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1/12288
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetBitrate(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTag(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDar(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDuration(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetFps(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHeight(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetIndex(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLang(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLevel(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNumFrames(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetPixFmt(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetProfile(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetRotate(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetSar(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetStartTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetTimebase(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetWidth(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo struct {
	// example:
	//
	// AiSaas
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 3048
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// example:
	//
	// cateName
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// example:
	//
	// category
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/01e1271d-ff4f-4689-9c20-e1df81486859_open_live_cover.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:15Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// tag1
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 123-1234
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// example:
	//
	// []
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ThumbURL240P
	ThumbURL240P *string `json:"ThumbURL240P,omitempty" xml:"ThumbURL240P,omitempty"`
	// example:
	//
	// ThumbURLWebp
	ThumbURLWebp *string `json:"ThumbURLWebp,omitempty" xml:"ThumbURLWebp,omitempty"`
	// example:
	//
	// v6l82k_176822379****.jpeg
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// general
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// example:
	//
	// UserData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCateName() *string {
	return s.CateName
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetThumbURL240P() *string {
	return s.ThumbURL240P
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetThumbURLWebp() *string {
	return s.ThumbURLWebp
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetUploadSource() *string {
	return s.UploadSource
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBiz(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBusinessType(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCateId(v int64) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CateId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCateName(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CateName = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCategory(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCoverURL(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCreateTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDeletedTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDescription(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetInputURL(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaTags(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaType(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetModifiedTime(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetReferenceId(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ReferenceId = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSnapshots(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSource(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSpriteImages(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetStatus(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetThumbURL240P(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ThumbURL240P = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetThumbURLWebp(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ThumbURLWebp = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetTitle(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUploadSource(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UploadSource = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUserData(v string) *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponseBodyMediaInfoMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
