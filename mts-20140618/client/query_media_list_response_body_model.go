// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaList(v *QueryMediaListResponseBodyMediaList) *QueryMediaListResponseBody
	GetMediaList() *QueryMediaListResponseBodyMediaList
	SetNonExistMediaIds(v *QueryMediaListResponseBodyNonExistMediaIds) *QueryMediaListResponseBody
	GetNonExistMediaIds() *QueryMediaListResponseBodyNonExistMediaIds
	SetRequestId(v string) *QueryMediaListResponseBody
	GetRequestId() *string
}

type QueryMediaListResponseBody struct {
	// The list of media files.
	MediaList *QueryMediaListResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Struct"`
	// The IDs of the media files that do not exist. This parameter is not returned when all specified media files exist.
	NonExistMediaIds *QueryMediaListResponseBodyNonExistMediaIds `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 283DC68C-146F-4489-A2A1-2F88F1472A56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBody) GetMediaList() *QueryMediaListResponseBodyMediaList {
	return s.MediaList
}

func (s *QueryMediaListResponseBody) GetNonExistMediaIds() *QueryMediaListResponseBodyNonExistMediaIds {
	return s.NonExistMediaIds
}

func (s *QueryMediaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaListResponseBody) SetMediaList(v *QueryMediaListResponseBodyMediaList) *QueryMediaListResponseBody {
	s.MediaList = v
	return s
}

func (s *QueryMediaListResponseBody) SetNonExistMediaIds(v *QueryMediaListResponseBodyNonExistMediaIds) *QueryMediaListResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *QueryMediaListResponseBody) SetRequestId(v string) *QueryMediaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaListResponseBody) Validate() error {
	if s.MediaList != nil {
		if err := s.MediaList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistMediaIds != nil {
		if err := s.NonExistMediaIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaList struct {
	Media []*QueryMediaListResponseBodyMediaListMedia `json:"Media,omitempty" xml:"Media,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaList) GetMedia() []*QueryMediaListResponseBodyMediaListMedia {
	return s.Media
}

func (s *QueryMediaListResponseBodyMediaList) SetMedia(v []*QueryMediaListResponseBodyMediaListMedia) *QueryMediaListResponseBodyMediaList {
	s.Media = v
	return s
}

func (s *QueryMediaListResponseBodyMediaList) Validate() error {
	if s.Media != nil {
		for _, item := range s.Media {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMedia struct {
	// The bitrate.
	//
	// example:
	//
	// 2659.326
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The ID of the category to which the media file belongs.
	//
	// example:
	//
	// 1
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The review status of the media file. Valid values:
	//
	// 	- **Initiated**: The media file is uploaded but not reviewed.
	//
	// 	- **Pass**: The media file is uploaded and passes the review.
	//
	// example:
	//
	// Initiated
	CensorState *string `json:"CensorState,omitempty" xml:"CensorState,omitempty"`
	// The OSS URL of the thumbnail.
	//
	// example:
	//
	// http://example-bucket1-****.oss-cn-hangzhou.aliyuncs.com//example-****.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media file was created.
	//
	// example:
	//
	// 2016-09-14T08:30:33Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// This is description ****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration.
	//
	// example:
	//
	// 7.965000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The details of the input file.
	File *QueryMediaListResponseBodyMediaListMediaFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The encoding format. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mov
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate.
	//
	// example:
	//
	// 25.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the media file.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// 3e1cd21131a94525be55acf65888****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The media information.
	MediaInfo *QueryMediaListResponseBodyMediaListMediaMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// The playlist.
	PlayList *QueryMediaListResponseBodyMediaListMediaPlayList `json:"PlayList,omitempty" xml:"PlayList,omitempty" type:"Struct"`
	// The publishing status of the media file. Valid values:
	//
	// - **Initiated**: The media file is in the initial state.
	//
	// - **UnPublish**: The media file has not been published, and the playback permission on the OSS object is Private.
	//
	// - **Published**: The media file has been published, and the playback permission on the OSS object is Default.
	//
	// - **Deleted**: The media file has been deleted.
	//
	// example:
	//
	// Published
	PublishState *string `json:"PublishState,omitempty" xml:"PublishState,omitempty"`
	// The ID of the instance.
	RunIdList *QueryMediaListResponseBodyMediaListMediaRunIdList `json:"RunIdList,omitempty" xml:"RunIdList,omitempty" type:"Struct"`
	// The size of the file.
	//
	// example:
	//
	// 2647692
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The list of snapshots.
	SnapshotList *QueryMediaListResponseBodyMediaListMediaSnapshotList `json:"SnapshotList,omitempty" xml:"SnapshotList,omitempty" type:"Struct"`
	// The list of video summaries.
	SummaryList *QueryMediaListResponseBodyMediaListMediaSummaryList `json:"SummaryList,omitempty" xml:"SummaryList,omitempty" type:"Struct"`
	// The tags of the media file.
	Tags *QueryMediaListResponseBodyMediaListMediaTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The title.
	//
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMedia) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMedia) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetCateId() *int64 {
	return s.CateId
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetCensorState() *string {
	return s.CensorState
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetCoverURL() *string {
	return s.CoverURL
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetDescription() *string {
	return s.Description
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetFile() *QueryMediaListResponseBodyMediaListMediaFile {
	return s.File
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetFormat() *string {
	return s.Format
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetMediaInfo() *QueryMediaListResponseBodyMediaListMediaMediaInfo {
	return s.MediaInfo
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetPlayList() *QueryMediaListResponseBodyMediaListMediaPlayList {
	return s.PlayList
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetPublishState() *string {
	return s.PublishState
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetRunIdList() *QueryMediaListResponseBodyMediaListMediaRunIdList {
	return s.RunIdList
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetSize() *string {
	return s.Size
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetSnapshotList() *QueryMediaListResponseBodyMediaListMediaSnapshotList {
	return s.SnapshotList
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetSummaryList() *QueryMediaListResponseBodyMediaListMediaSummaryList {
	return s.SummaryList
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetTags() *QueryMediaListResponseBodyMediaListMediaTags {
	return s.Tags
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetTitle() *string {
	return s.Title
}

func (s *QueryMediaListResponseBodyMediaListMedia) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetBitrate(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetCateId(v int64) *QueryMediaListResponseBodyMediaListMedia {
	s.CateId = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetCensorState(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.CensorState = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetCoverURL(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.CoverURL = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetCreationTime(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetDescription(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Description = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetDuration(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Duration = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetFile(v *QueryMediaListResponseBodyMediaListMediaFile) *QueryMediaListResponseBodyMediaListMedia {
	s.File = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetFormat(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Format = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetFps(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Fps = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetHeight(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Height = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetMediaId(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.MediaId = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetMediaInfo(v *QueryMediaListResponseBodyMediaListMediaMediaInfo) *QueryMediaListResponseBodyMediaListMedia {
	s.MediaInfo = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetPlayList(v *QueryMediaListResponseBodyMediaListMediaPlayList) *QueryMediaListResponseBodyMediaListMedia {
	s.PlayList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetPublishState(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.PublishState = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetRunIdList(v *QueryMediaListResponseBodyMediaListMediaRunIdList) *QueryMediaListResponseBodyMediaListMedia {
	s.RunIdList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetSize(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Size = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetSnapshotList(v *QueryMediaListResponseBodyMediaListMediaSnapshotList) *QueryMediaListResponseBodyMediaListMedia {
	s.SnapshotList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetSummaryList(v *QueryMediaListResponseBodyMediaListMediaSummaryList) *QueryMediaListResponseBodyMediaListMedia {
	s.SummaryList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetTags(v *QueryMediaListResponseBodyMediaListMediaTags) *QueryMediaListResponseBodyMediaListMedia {
	s.Tags = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetTitle(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Title = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) SetWidth(v string) *QueryMediaListResponseBodyMediaListMedia {
	s.Width = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMedia) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.MediaInfo != nil {
		if err := s.MediaInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PlayList != nil {
		if err := s.PlayList.Validate(); err != nil {
			return err
		}
	}
	if s.RunIdList != nil {
		if err := s.RunIdList.Validate(); err != nil {
			return err
		}
	}
	if s.SnapshotList != nil {
		if err := s.SnapshotList.Validate(); err != nil {
			return err
		}
	}
	if s.SummaryList != nil {
		if err := s.SummaryList.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaFile struct {
	// The status of the input file. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Deleted**: deleted
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The OSS URL of the input file.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com//example-****.mp4
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListResponseBodyMediaListMediaFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListResponseBodyMediaListMediaFile) SetState(v string) *QueryMediaListResponseBodyMediaListMediaFile {
	s.State = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaFile) SetURL(v string) *QueryMediaListResponseBodyMediaListMediaFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaMediaInfo struct {
	// The format information.
	Format *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	// The stream information.
	Streams *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfo) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfo) GetFormat() *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	return s.Format
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfo) GetStreams() *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams {
	return s.Streams
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfo) SetFormat(v *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) *QueryMediaListResponseBodyMediaListMediaMediaInfo {
	s.Format = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfo) SetStreams(v *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) *QueryMediaListResponseBodyMediaListMediaMediaInfo {
	s.Streams = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfo) Validate() error {
	if s.Format != nil {
		if err := s.Format.Validate(); err != nil {
			return err
		}
	}
	if s.Streams != nil {
		if err := s.Streams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoFormat struct {
	// The bitrate.
	//
	// example:
	//
	// 2659.326
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The total duration.
	//
	// example:
	//
	// 7.965000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The full name of the container format.
	//
	// example:
	//
	// QuickTime/MOV
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	// The short name of the container format. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mov
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The total number of program streams.
	//
	// example:
	//
	// 2
	NumPrograms *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	// The total number of media streams.
	//
	// example:
	//
	// 2
	NumStreams *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	// The size of the file.
	//
	// example:
	//
	// 2647692
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetSize() *string {
	return s.Size
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetBitrate(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetDuration(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.Duration = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetFormatLongName(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.FormatLongName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetFormatName(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.FormatName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetNumPrograms(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.NumPrograms = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetNumStreams(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.NumStreams = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetSize(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.Size = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) SetStartTime(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat {
	s.StartTime = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoFormat) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreams struct {
	// The list of audio streams.
	AudioStreamList *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	// The list of subtitle streams.
	SubtitleStreamList *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	// The list of video streams.
	VideoStreamList *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) GetAudioStreamList() *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) GetSubtitleStreamList() *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) GetVideoStreamList() *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) SetAudioStreamList(v *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams {
	s.AudioStreamList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) SetSubtitleStreamList(v *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) SetVideoStreamList(v *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams {
	s.VideoStreamList = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreams) Validate() error {
	if s.AudioStreamList != nil {
		if err := s.AudioStreamList.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleStreamList != nil {
		if err := s.SubtitleStreamList.Validate(); err != nil {
			return err
		}
	}
	if s.VideoStreamList != nil {
		if err := s.VideoStreamList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList struct {
	AudioStream []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) GetAudioStream() []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) SetAudioStream(v []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) Validate() error {
	if s.AudioStream != nil {
		for _, item := range s.AudioStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream struct {
	// The bitrate.
	//
	// example:
	//
	// 160.008
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The output layout of the sound channels.
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
	// The full name of the encoding format.
	//
	// example:
	//
	// AAC(Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values: H264, mov, aac, avc, and mpeg.
	//
	// example:
	//
	// mov
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
	// The duration.
	//
	// example:
	//
	// 182.591995
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html#Metadata).
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 100
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
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
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
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

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetBitrate(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetChannels(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecName(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecTag(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetDuration(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetIndex(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetLang(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetNumFrames(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetSamplerate(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetStartTime(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetTimebase(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList struct {
	SubtitleStream []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) GetSubtitleStream() []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) SetSubtitleStream(v []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) Validate() error {
	if s.SubtitleStream != nil {
		for _, item := range s.SubtitleStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream struct {
	// The sequence number of the subtitle stream. The value indicates the position of the subtitle stream in all subtitle streams.
	//
	// example:
	//
	// 3
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html#Metadata).
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList struct {
	VideoStream []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) GetVideoStream() []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) SetVideoStream(v []*QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) Validate() error {
	if s.VideoStream != nil {
		for _, item := range s.VideoStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream struct {
	// The average frame rate.
	//
	// example:
	//
	// 29.97003
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate.
	//
	// example:
	//
	// 2659.326
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// QuickTime/MOV
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mov
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
	// 1001/60000
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR).
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration.
	//
	// example:
	//
	// 182.683000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate.
	//
	// example:
	//
	// 29.97003
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Indicates whether the video stream contains bidirectional frames (B-frames). A value of **1*	- indicates that the video stream contains B-frames. A value of **2*	- indicates that the video stream does not contain B-frames.
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The latter number in the video resolution. The number indicates the video height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video stream. The value indicates the position of the video stream in all video streams.
	//
	// example:
	//
	// 5
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html#Metadata).
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The codec level.
	//
	// example:
	//
	// 40
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network bandwidth consumption.
	NetworkCost *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	// The total number of frames.
	//
	// example:
	//
	// 12
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
	// The video rotation angle.
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	//
	// example:
	//
	// 1:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
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
	// 1/30000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The former number in the video resolution. The number indicates the video width.
	//
	// example:
	//
	// 100
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetNetworkCost() *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetRotate() *string {
	return s.Rotate
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetBitrate(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecName(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecTag(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetDar(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetDuration(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetFps(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetHeight(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetIndex(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetLang(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetLevel(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetNetworkCost(v *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetNumFrames(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetPixFmt(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetProfile(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetRotate(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Rotate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetSar(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetStartTime(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetTimebase(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetWidth(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost struct {
	// The average bitrate.
	//
	// example:
	//
	// 2659.326
	AvgBitrate *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	// The maximum bandwidth that was consumed.
	//
	// example:
	//
	// 100
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	// The amount of preload time.
	//
	// example:
	//
	// 0.01
	PreloadTime *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaPlayList struct {
	Play []*QueryMediaListResponseBodyMediaListMediaPlayListPlay `json:"Play,omitempty" xml:"Play,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaPlayList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaPlayList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayList) GetPlay() []*QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	return s.Play
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayList) SetPlay(v []*QueryMediaListResponseBodyMediaListMediaPlayListPlay) *QueryMediaListResponseBodyMediaListMediaPlayList {
	s.Play = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayList) Validate() error {
	if s.Play != nil {
		for _, item := range s.Play {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaPlayListPlay struct {
	// The name of the workflow activity.
	//
	// example:
	//
	// example-activity-****
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The bitrate of the media file.
	//
	// example:
	//
	// 2659.326
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the media file.
	//
	// example:
	//
	// 7.965000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether the media file is encrypted. Valid values:
	//
	// 	- **0**: The media file is not encrypted.
	//
	// 	- **1**: The media file is encrypted.
	//
	// example:
	//
	// 0
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// The playback file.
	File *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The encoding format of the media file. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the media file.
	//
	// example:
	//
	// 25.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the workflow that generates the playback file.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e91d24d****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the workflow that generates the playback file.
	//
	// example:
	//
	// example-mediaworkflow-****
	MediaWorkflowName *string `json:"MediaWorkflowName,omitempty" xml:"MediaWorkflowName,omitempty"`
	// The size of the media file.
	//
	// example:
	//
	// 2647692
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The width of the media file.
	//
	// example:
	//
	// 760
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaPlayListPlay) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaPlayListPlay) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetActivityName() *string {
	return s.ActivityName
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetEncryption() *string {
	return s.Encryption
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetFile() *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile {
	return s.File
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetFormat() *string {
	return s.Format
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetSize() *string {
	return s.Size
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetActivityName(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.ActivityName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetBitrate(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetDuration(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Duration = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetEncryption(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Encryption = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetFile(v *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.File = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetFormat(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Format = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetFps(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Fps = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetHeight(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Height = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetMediaWorkflowId(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetMediaWorkflowName(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.MediaWorkflowName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetSize(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Size = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) SetWidth(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlay {
	s.Width = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlay) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaPlayListPlayFile struct {
	// The status of the file. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Deleted**: deleted
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The Object Storage Service (OSS) URL of the output file.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com//example-****.mp4
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) SetState(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile {
	s.State = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) SetURL(v string) *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaPlayListPlayFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaRunIdList struct {
	RunId []*string `json:"RunId,omitempty" xml:"RunId,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaRunIdList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaRunIdList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaRunIdList) GetRunId() []*string {
	return s.RunId
}

func (s *QueryMediaListResponseBodyMediaListMediaRunIdList) SetRunId(v []*string) *QueryMediaListResponseBodyMediaListMediaRunIdList {
	s.RunId = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaRunIdList) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaSnapshotList struct {
	Snapshot []*QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaSnapshotList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaSnapshotList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotList) GetSnapshot() []*QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	return s.Snapshot
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotList) SetSnapshot(v []*QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) *QueryMediaListResponseBodyMediaListMediaSnapshotList {
	s.Snapshot = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotList) Validate() error {
	if s.Snapshot != nil {
		for _, item := range s.Snapshot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot struct {
	// The name of the workflow activity that generates the snapshot.
	//
	// example:
	//
	// example-activity1-****
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The number of snapshots. This parameter is valid only when the value of the **Type*	- parameter is **Sequence**.
	//
	// example:
	//
	// 5
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The snapshot.
	File *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The ID of the workflow that generates the snapshot.
	//
	// example:
	//
	// 6cc3aa66d1cb4bb2adf14e726c0a****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the workflow that generates the snapshot.
	//
	// example:
	//
	// example-workflow-****
	MediaWorkflowName *string `json:"MediaWorkflowName,omitempty" xml:"MediaWorkflowName,omitempty"`
	// The type of the snapshot. Valid values:
	//
	// - **Single**
	//
	// - **Sequence**
	//
	// example:
	//
	// Sequence
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GetActivityName() *string {
	return s.ActivityName
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GetCount() *string {
	return s.Count
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GetFile() *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile {
	return s.File
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) GetType() *string {
	return s.Type
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) SetActivityName(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	s.ActivityName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) SetCount(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	s.Count = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) SetFile(v *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	s.File = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) SetMediaWorkflowId(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) SetMediaWorkflowName(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	s.MediaWorkflowName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) SetType(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot {
	s.Type = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshot) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile struct {
	// The status of the file. Valid values:
	//
	// - **Normal**: normal
	//
	// - **Deleted**: deleted
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The OSS URL of the snapshot.
	//
	// example:
	//
	// http://example1-bucket1-****.oss-cn-hangzhou.aliyuncs.com//example111-****.png
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) SetState(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile {
	s.State = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) SetURL(v string) *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSnapshotListSnapshotFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaSummaryList struct {
	Summary []*QueryMediaListResponseBodyMediaListMediaSummaryListSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaSummaryList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaSummaryList) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryList) GetSummary() []*QueryMediaListResponseBodyMediaListMediaSummaryListSummary {
	return s.Summary
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryList) SetSummary(v []*QueryMediaListResponseBodyMediaListMediaSummaryListSummary) *QueryMediaListResponseBodyMediaListMediaSummaryList {
	s.Summary = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryList) Validate() error {
	if s.Summary != nil {
		for _, item := range s.Summary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaSummaryListSummary struct {
	// The name of the workflow activity.
	//
	// example:
	//
	// example-activity-****
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The information about the input file.
	File *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The ID of the workflow that generates the summary.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e91d24d****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the workflow that generates the summary.
	//
	// example:
	//
	// example-mediaworkflow-****
	MediaWorkflowName *string `json:"MediaWorkflowName,omitempty" xml:"MediaWorkflowName,omitempty"`
	// The type of the summary. Valid values:
	//
	// 	- **Video**: video
	//
	// 	- **Gif**: dynamic image
	//
	// example:
	//
	// Video
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaSummaryListSummary) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaSummaryListSummary) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) GetActivityName() *string {
	return s.ActivityName
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) GetFile() *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile {
	return s.File
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) GetType() *string {
	return s.Type
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) SetActivityName(v string) *QueryMediaListResponseBodyMediaListMediaSummaryListSummary {
	s.ActivityName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) SetFile(v *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) *QueryMediaListResponseBodyMediaListMediaSummaryListSummary {
	s.File = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) SetMediaWorkflowId(v string) *QueryMediaListResponseBodyMediaListMediaSummaryListSummary {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) SetMediaWorkflowName(v string) *QueryMediaListResponseBodyMediaListMediaSummaryListSummary {
	s.MediaWorkflowName = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) SetType(v string) *QueryMediaListResponseBodyMediaListMediaSummaryListSummary {
	s.Type = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummary) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile struct {
	// The status of the file. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Deleted**: deleted
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The OSS URL of the input file.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com//example-****.mp4
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) SetState(v string) *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile {
	s.State = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) SetURL(v string) *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaSummaryListSummaryFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyMediaListMediaTags struct {
	Tag []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyMediaListMediaTags) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyMediaListMediaTags) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyMediaListMediaTags) GetTag() []*string {
	return s.Tag
}

func (s *QueryMediaListResponseBodyMediaListMediaTags) SetTag(v []*string) *QueryMediaListResponseBodyMediaListMediaTags {
	s.Tag = v
	return s
}

func (s *QueryMediaListResponseBodyMediaListMediaTags) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListResponseBodyNonExistMediaIds struct {
	MediaId []*string `json:"MediaId,omitempty" xml:"MediaId,omitempty" type:"Repeated"`
}

func (s QueryMediaListResponseBodyNonExistMediaIds) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponseBodyNonExistMediaIds) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponseBodyNonExistMediaIds) GetMediaId() []*string {
	return s.MediaId
}

func (s *QueryMediaListResponseBodyNonExistMediaIds) SetMediaId(v []*string) *QueryMediaListResponseBodyNonExistMediaIds {
	s.MediaId = v
	return s
}

func (s *QueryMediaListResponseBodyNonExistMediaIds) Validate() error {
	return dara.Validate(s)
}
