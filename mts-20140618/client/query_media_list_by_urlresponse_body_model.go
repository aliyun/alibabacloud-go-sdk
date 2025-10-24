// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaListByURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaList(v *QueryMediaListByURLResponseBodyMediaList) *QueryMediaListByURLResponseBody
	GetMediaList() *QueryMediaListByURLResponseBodyMediaList
	SetNonExistFileURLs(v *QueryMediaListByURLResponseBodyNonExistFileURLs) *QueryMediaListByURLResponseBody
	GetNonExistFileURLs() *QueryMediaListByURLResponseBodyNonExistFileURLs
	SetRequestId(v string) *QueryMediaListByURLResponseBody
	GetRequestId() *string
}

type QueryMediaListByURLResponseBody struct {
	// The list of media files.
	MediaList *QueryMediaListByURLResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Struct"`
	// The IDs of the media files that do not exist. This parameter is not returned if all specified media files exist.
	NonExistFileURLs *QueryMediaListByURLResponseBodyNonExistFileURLs `json:"NonExistFileURLs,omitempty" xml:"NonExistFileURLs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1C8A0AEB-4321-485B-B4CB-DA4E9E6C9B42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaListByURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBody) GetMediaList() *QueryMediaListByURLResponseBodyMediaList {
	return s.MediaList
}

func (s *QueryMediaListByURLResponseBody) GetNonExistFileURLs() *QueryMediaListByURLResponseBodyNonExistFileURLs {
	return s.NonExistFileURLs
}

func (s *QueryMediaListByURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaListByURLResponseBody) SetMediaList(v *QueryMediaListByURLResponseBodyMediaList) *QueryMediaListByURLResponseBody {
	s.MediaList = v
	return s
}

func (s *QueryMediaListByURLResponseBody) SetNonExistFileURLs(v *QueryMediaListByURLResponseBodyNonExistFileURLs) *QueryMediaListByURLResponseBody {
	s.NonExistFileURLs = v
	return s
}

func (s *QueryMediaListByURLResponseBody) SetRequestId(v string) *QueryMediaListByURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaListByURLResponseBody) Validate() error {
	if s.MediaList != nil {
		if err := s.MediaList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistFileURLs != nil {
		if err := s.NonExistFileURLs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListByURLResponseBodyMediaList struct {
	Media []*QueryMediaListByURLResponseBodyMediaListMedia `json:"Media,omitempty" xml:"Media,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaList) GetMedia() []*QueryMediaListByURLResponseBodyMediaListMedia {
	return s.Media
}

func (s *QueryMediaListByURLResponseBodyMediaList) SetMedia(v []*QueryMediaListByURLResponseBodyMediaListMedia) *QueryMediaListByURLResponseBodyMediaList {
	s.Media = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMedia struct {
	// The bitrate.
	//
	// example:
	//
	// 593.192
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The ID of the category to which the media file belongs.
	//
	// example:
	//
	// 123
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
	// 2021-07-14T13:05:00Z
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
	// 79.204000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The details of the input file.
	File *QueryMediaListByURLResponseBodyMediaListMediaFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
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
	// 15.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the queried media file.
	//
	// example:
	//
	// 360
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// 52d7e98b05e648199612290bb819****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The media information.
	MediaInfo *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// The playlist.
	PlayList *QueryMediaListByURLResponseBodyMediaListMediaPlayList `json:"PlayList,omitempty" xml:"PlayList,omitempty" type:"Struct"`
	// The publishing status of the media file. Valid values:
	//
	// - **Initiated**: The media file is in the initial state.
	//
	// - **UnPublish**: The media file has not been published, and the playback permission on the OSS object is Private.
	//
	// - **Published**: The media file has been published, and the playback permission on the OSS object is Default.
	//
	// - **Deleted**: The file is deleted.
	//
	// example:
	//
	// Published
	PublishState *string `json:"PublishState,omitempty" xml:"PublishState,omitempty"`
	// The IDs of the media workflow execution instances.
	RunIdList *QueryMediaListByURLResponseBodyMediaListMediaRunIdList `json:"RunIdList,omitempty" xml:"RunIdList,omitempty" type:"Struct"`
	// The size of the file.
	//
	// example:
	//
	// 5872904
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The list of snapshots.
	SnapshotList *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList `json:"SnapshotList,omitempty" xml:"SnapshotList,omitempty" type:"Struct"`
	// The list of video summaries.
	SummaryList *QueryMediaListByURLResponseBodyMediaListMediaSummaryList `json:"SummaryList,omitempty" xml:"SummaryList,omitempty" type:"Struct"`
	// The tags of the media file.
	Tags *QueryMediaListByURLResponseBodyMediaListMediaTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The title.
	//
	// example:
	//
	// kled.mp4
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The width.
	//
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMedia) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMedia) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetCateId() *int64 {
	return s.CateId
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetCensorState() *string {
	return s.CensorState
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetCoverURL() *string {
	return s.CoverURL
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetDescription() *string {
	return s.Description
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetFile() *QueryMediaListByURLResponseBodyMediaListMediaFile {
	return s.File
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetFormat() *string {
	return s.Format
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetMediaInfo() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo {
	return s.MediaInfo
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetPlayList() *QueryMediaListByURLResponseBodyMediaListMediaPlayList {
	return s.PlayList
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetPublishState() *string {
	return s.PublishState
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetRunIdList() *QueryMediaListByURLResponseBodyMediaListMediaRunIdList {
	return s.RunIdList
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetSize() *string {
	return s.Size
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetSnapshotList() *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList {
	return s.SnapshotList
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetSummaryList() *QueryMediaListByURLResponseBodyMediaListMediaSummaryList {
	return s.SummaryList
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetTags() *QueryMediaListByURLResponseBodyMediaListMediaTags {
	return s.Tags
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetTitle() *string {
	return s.Title
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetBitrate(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetCateId(v int64) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.CateId = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetCensorState(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.CensorState = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetCoverURL(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.CoverURL = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetCreationTime(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetDescription(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Description = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetDuration(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Duration = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetFile(v *QueryMediaListByURLResponseBodyMediaListMediaFile) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.File = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetFormat(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Format = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetFps(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Fps = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetHeight(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Height = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetMediaId(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.MediaId = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetMediaInfo(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.MediaInfo = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetPlayList(v *QueryMediaListByURLResponseBodyMediaListMediaPlayList) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.PlayList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetPublishState(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.PublishState = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetRunIdList(v *QueryMediaListByURLResponseBodyMediaListMediaRunIdList) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.RunIdList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetSize(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Size = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetSnapshotList(v *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.SnapshotList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetSummaryList(v *QueryMediaListByURLResponseBodyMediaListMediaSummaryList) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.SummaryList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetTags(v *QueryMediaListByURLResponseBodyMediaListMediaTags) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Tags = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetTitle(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Title = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) SetWidth(v string) *QueryMediaListByURLResponseBodyMediaListMedia {
	s.Width = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMedia) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaFile struct {
	// The status of the media file. Valid values:
	//
	// 	- **Normal**: The file is normal.
	//
	// 	- **Deleted**: The file is deleted.
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

func (s QueryMediaListByURLResponseBodyMediaListMediaFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaFile) SetState(v string) *QueryMediaListByURLResponseBodyMediaListMediaFile {
	s.State = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaFile) SetURL(v string) *QueryMediaListByURLResponseBodyMediaListMediaFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfo struct {
	// The format information.
	Format *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	// The stream information.
	Streams *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) GetFormat() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	return s.Format
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) GetStreams() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams {
	return s.Streams
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) SetFormat(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo {
	s.Format = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) SetStreams(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo {
	s.Streams = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfo) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat struct {
	// The bitrate.
	//
	// example:
	//
	// 593.192
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration.
	//
	// example:
	//
	// 79.204000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The full name of the encoding format.
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
	// 0
	NumPrograms *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	// The total number of media streams.
	//
	// example:
	//
	// 2
	NumStreams *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	// The size.
	//
	// example:
	//
	// 5872904
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetSize() *string {
	return s.Size
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetBitrate(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetDuration(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.Duration = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetFormatLongName(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.FormatLongName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetFormatName(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.FormatName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetNumPrograms(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.NumPrograms = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetNumStreams(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.NumStreams = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetSize(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.Size = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) SetStartTime(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat {
	s.StartTime = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoFormat) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams struct {
	// The list of audio streams.
	AudioStreamList *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	// The list of subtitle streams.
	SubtitleStreamList *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	// The list of video streams.
	VideoStreamList *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) GetAudioStreamList() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) GetSubtitleStreamList() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) GetVideoStreamList() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) SetAudioStreamList(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams {
	s.AudioStreamList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) SetSubtitleStreamList(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) SetVideoStreamList(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams {
	s.VideoStreamList = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreams) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList struct {
	AudioStream []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) GetAudioStream() []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) SetAudioStream(v []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream struct {
	// The bitrate.
	//
	// example:
	//
	// 76.356
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
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values: H264, mov, aac, avc, and mpeg.
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
	// The duration of the media file.
	//
	// example:
	//
	// 79.203265
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language. For more information, see [FFmpeg language definition](https://www.ffmpeg.org/ffmpeg-all.html#Metadata).
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

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetBitrate(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetChannels(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecName(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecTag(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetDuration(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetIndex(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetLang(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetNumFrames(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetSamplerate(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetStartTime(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) SetTimebase(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList struct {
	SubtitleStream []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) GetSubtitleStream() []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) SetSubtitleStream(v []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream struct {
	// The sequence number of the subtitle stream. The value indicates the position of the subtitle stream in all subtitle streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language. For more information, see [FFmpeg language definition](https://www.ffmpeg.org/ffmpeg-all.html#Metadata).
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList struct {
	VideoStream []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) GetVideoStream() []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) SetVideoStream(v []*QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream struct {
	// The average frame rate.
	//
	// example:
	//
	// 15.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate.
	//
	// example:
	//
	// 512.701
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// H.264/AVC/MPEG-4 AVC/MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values: H264, mov, aac, avc, and mpeg.
	//
	// example:
	//
	// H264
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
	// 1/30
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR) of the video stream.
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration.
	//
	// example:
	//
	// 79.200000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate.
	//
	// example:
	//
	// 15.0
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
	// 360
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
	// 31
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network bandwidth consumption.
	NetworkCost *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	// The total number of frames.
	//
	// example:
	//
	// 12
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format of the video stream.
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
	// The rotation angle of the video.
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
	// 0.046029
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/15360
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The former number in the video resolution. The number indicates the video width and cannot be negative.
	//
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetNetworkCost() *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetRotate() *string {
	return s.Rotate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetBitrate(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecName(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecTag(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetDar(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetDuration(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetFps(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetHeight(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetIndex(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetLang(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetLevel(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetNetworkCost(v *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetNumFrames(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetPixFmt(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetProfile(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetRotate(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Rotate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetSar(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetStartTime(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetTimebase(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) SetWidth(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost struct {
	// The average bitrate of the video stream.
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

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaMediaInfoStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaPlayList struct {
	Play []*QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay `json:"Play,omitempty" xml:"Play,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaPlayList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaPlayList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayList) GetPlay() []*QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	return s.Play
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayList) SetPlay(v []*QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) *QueryMediaListByURLResponseBodyMediaListMediaPlayList {
	s.Play = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay struct {
	// The name of the workflow activity.
	//
	// example:
	//
	// test name
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The bitrate of the media file.
	//
	// example:
	//
	// 25.067
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration.
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
	// 1
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// The playback file.
	File *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The encoding format of the media file. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
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
	// 10
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the workflow that generates the playback file.
	//
	// example:
	//
	// 6cc3aa66d1cb4bb2adf14e726c0a****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the workflow that generates the playback file.
	//
	// example:
	//
	// example-mediaworkflow-****
	MediaWorkflowName *string `json:"MediaWorkflowName,omitempty" xml:"MediaWorkflowName,omitempty"`
	// The size.
	//
	// example:
	//
	// 100
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The width.
	//
	// example:
	//
	// 11
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetActivityName() *string {
	return s.ActivityName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetEncryption() *string {
	return s.Encryption
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetFile() *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile {
	return s.File
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetFormat() *string {
	return s.Format
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetSize() *string {
	return s.Size
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetActivityName(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.ActivityName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetBitrate(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetDuration(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Duration = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetEncryption(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Encryption = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetFile(v *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.File = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetFormat(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Format = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetFps(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Fps = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetHeight(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Height = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetMediaWorkflowId(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetMediaWorkflowName(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.MediaWorkflowName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetSize(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Size = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) SetWidth(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay {
	s.Width = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlay) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile struct {
	// The status of the media file. Valid values:
	//
	// 	- **Normal**: The file is normal.
	//
	// 	- **Deleted**: The file is deleted.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The OSS URL of the playback file.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com//example-****.mp4l-test/in/1.mp4
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) SetState(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile {
	s.State = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) SetURL(v string) *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaPlayListPlayFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaRunIdList struct {
	RunId []*string `json:"RunId,omitempty" xml:"RunId,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaRunIdList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaRunIdList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaRunIdList) GetRunId() []*string {
	return s.RunId
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaRunIdList) SetRunId(v []*string) *QueryMediaListByURLResponseBodyMediaListMediaRunIdList {
	s.RunId = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaRunIdList) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaSnapshotList struct {
	Snapshot []*QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSnapshotList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSnapshotList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList) GetSnapshot() []*QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	return s.Snapshot
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList) SetSnapshot(v []*QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList {
	s.Snapshot = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot struct {
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
	// 3
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The snapshot.
	File *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
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
	// - **Single**: a single snapshot
	//
	// - **Sequence**: snapshots in sequence
	//
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GetActivityName() *string {
	return s.ActivityName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GetCount() *string {
	return s.Count
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GetFile() *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile {
	return s.File
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) GetType() *string {
	return s.Type
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) SetActivityName(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	s.ActivityName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) SetCount(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	s.Count = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) SetFile(v *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	s.File = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) SetMediaWorkflowId(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) SetMediaWorkflowName(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	s.MediaWorkflowName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) SetType(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot {
	s.Type = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshot) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile struct {
	// The status of the file. Valid values:
	//
	// - **Normal**: The file is normal.
	//
	// - **Deleted**: The file is deleted.
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

func (s QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) SetState(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile {
	s.State = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) SetURL(v string) *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSnapshotListSnapshotFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaSummaryList struct {
	Summary []*QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSummaryList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSummaryList) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryList) GetSummary() []*QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary {
	return s.Summary
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryList) SetSummary(v []*QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) *QueryMediaListByURLResponseBodyMediaListMediaSummaryList {
	s.Summary = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryList) Validate() error {
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

type QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary struct {
	// The name of the workflow activity.
	//
	// example:
	//
	// example-activity-****
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The information about the input file.
	File *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
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
	// Gif
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) GetActivityName() *string {
	return s.ActivityName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) GetFile() *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile {
	return s.File
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) GetType() *string {
	return s.Type
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) SetActivityName(v string) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary {
	s.ActivityName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) SetFile(v *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary {
	s.File = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) SetMediaWorkflowId(v string) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) SetMediaWorkflowName(v string) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary {
	s.MediaWorkflowName = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) SetType(v string) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary {
	s.Type = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummary) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile struct {
	// The status of the media file. Valid values:
	//
	// 	- **Normal**: The file is normal.
	//
	// 	- **Deleted**: The file is deleted.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The OSS URL of the input file.
	//
	// example:
	//
	// http://example-bucket-****.o
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) GetState() *string {
	return s.State
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) GetURL() *string {
	return s.URL
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) SetState(v string) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile {
	s.State = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) SetURL(v string) *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile {
	s.URL = &v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaSummaryListSummaryFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyMediaListMediaTags struct {
	Tag []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyMediaListMediaTags) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyMediaListMediaTags) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaTags) GetTag() []*string {
	return s.Tag
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaTags) SetTag(v []*string) *QueryMediaListByURLResponseBodyMediaListMediaTags {
	s.Tag = v
	return s
}

func (s *QueryMediaListByURLResponseBodyMediaListMediaTags) Validate() error {
	return dara.Validate(s)
}

type QueryMediaListByURLResponseBodyNonExistFileURLs struct {
	FileURL []*string `json:"FileURL,omitempty" xml:"FileURL,omitempty" type:"Repeated"`
}

func (s QueryMediaListByURLResponseBodyNonExistFileURLs) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponseBodyNonExistFileURLs) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponseBodyNonExistFileURLs) GetFileURL() []*string {
	return s.FileURL
}

func (s *QueryMediaListByURLResponseBodyNonExistFileURLs) SetFileURL(v []*string) *QueryMediaListByURLResponseBodyNonExistFileURLs {
	s.FileURL = v
	return s
}

func (s *QueryMediaListByURLResponseBodyNonExistFileURLs) Validate() error {
	return dara.Validate(s)
}
