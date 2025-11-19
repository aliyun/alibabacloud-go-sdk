// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveRecordVideoList(v *ListLiveRecordVideoResponseBodyLiveRecordVideoList) *ListLiveRecordVideoResponseBody
	GetLiveRecordVideoList() *ListLiveRecordVideoResponseBodyLiveRecordVideoList
	SetRequestId(v string) *ListLiveRecordVideoResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListLiveRecordVideoResponseBody
	GetTotal() *int32
}

type ListLiveRecordVideoResponseBody struct {
	// The list of videos.
	LiveRecordVideoList *ListLiveRecordVideoResponseBodyLiveRecordVideoList `json:"LiveRecordVideoList,omitempty" xml:"LiveRecordVideoList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-****-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of videos.
	//
	// example:
	//
	// 123
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLiveRecordVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBody) GetLiveRecordVideoList() *ListLiveRecordVideoResponseBodyLiveRecordVideoList {
	return s.LiveRecordVideoList
}

func (s *ListLiveRecordVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRecordVideoResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListLiveRecordVideoResponseBody) SetLiveRecordVideoList(v *ListLiveRecordVideoResponseBodyLiveRecordVideoList) *ListLiveRecordVideoResponseBody {
	s.LiveRecordVideoList = v
	return s
}

func (s *ListLiveRecordVideoResponseBody) SetRequestId(v string) *ListLiveRecordVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBody) SetTotal(v int32) *ListLiveRecordVideoResponseBody {
	s.Total = &v
	return s
}

func (s *ListLiveRecordVideoResponseBody) Validate() error {
	if s.LiveRecordVideoList != nil {
		if err := s.LiveRecordVideoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoList struct {
	LiveRecordVideo []*ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo `json:"LiveRecordVideo,omitempty" xml:"LiveRecordVideo,omitempty" type:"Repeated"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoList) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoList) GetLiveRecordVideo() []*ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	return s.LiveRecordVideo
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoList) SetLiveRecordVideo(v []*ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) *ListLiveRecordVideoResponseBodyLiveRecordVideoList {
	s.LiveRecordVideo = v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoList) Validate() error {
	if s.LiveRecordVideo != nil {
		for _, item := range s.LiveRecordVideo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo struct {
	// The name of the app.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the playlist.
	//
	// example:
	//
	// ****
	PlaylistId *string `json:"PlaylistId,omitempty" xml:"PlaylistId,omitempty"`
	// The recording end time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-08T08:44:56Z
	RecordEndTime *string `json:"RecordEndTime,omitempty" xml:"RecordEndTime,omitempty"`
	// The recording start time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-08T07:40:56Z
	RecordStartTime *string `json:"RecordStartTime,omitempty" xml:"RecordStartTime,omitempty"`
	// The name of the live-to-VOD file.
	//
	// example:
	//
	// live-test
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The information about the live-to-VOD file.
	Video *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetAppName() *string {
	return s.AppName
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetDomainName() *string {
	return s.DomainName
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetPlaylistId() *string {
	return s.PlaylistId
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetRecordEndTime() *string {
	return s.RecordEndTime
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetRecordStartTime() *string {
	return s.RecordStartTime
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetStreamName() *string {
	return s.StreamName
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GetVideo() *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	return s.Video
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetAppName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.AppName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetDomainName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.DomainName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetPlaylistId(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.PlaylistId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetRecordEndTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.RecordEndTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetRecordStartTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.RecordStartTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetStreamName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.StreamName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetVideo(v *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.Video = v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) Validate() error {
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo struct {
	// The ID of the video category.
	//
	// example:
	//
	// 78
	CateId *int32 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category of the video.
	//
	// example:
	//
	// Category name
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The URL of the video thumbnail.
	//
	// example:
	//
	// https://example.aliyundoc.com/coversample.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the audio or video file was created. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-08T07:40:56Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the video file.
	//
	// example:
	//
	// Description of the ApsaraVideo VOD video
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the video file. Unit: seconds.
	//
	// example:
	//
	// 135.6
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the video was updated. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-08T09:40:56Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The size of the source video file. Unit: bytes.
	//
	// example:
	//
	// 10897890
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The array of video snapshot URLs.
	Snapshots *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The status of the video. Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **UploadFail**
	//
	// 	- **UploadSuccess**
	//
	// 	- **Transcoding**
	//
	// 	- **TranscodeFail**
	//
	// 	- **Blocked**
	//
	// 	- **Normal**: The video is normal.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the video. Multiple tags are separated with commas (,).
	//
	// example:
	//
	// tag1, tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// 1
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the video.
	//
	// example:
	//
	// Title of the ApsaraVideo VOD video
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 93ab850b4f6f*****54b6e91d24d81d4
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetCateId() *int32 {
	return s.CateId
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetCateName() *string {
	return s.CateName
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetDescription() *string {
	return s.Description
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetDuration() *float32 {
	return s.Duration
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetSize() *int64 {
	return s.Size
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetSnapshots() *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots {
	return s.Snapshots
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetStatus() *string {
	return s.Status
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetTags() *string {
	return s.Tags
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetTitle() *string {
	return s.Title
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GetVideoId() *string {
	return s.VideoId
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCateId(v int32) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CateId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCateName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CateName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCoverURL(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CoverURL = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCreationTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CreationTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetDescription(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Description = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetDuration(v float32) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Duration = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetModifyTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.ModifyTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetSize(v int64) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Size = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetSnapshots(v *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Snapshots = v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetStatus(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Status = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetTags(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Tags = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetTemplateGroupId(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.TemplateGroupId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetTitle(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Title = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetVideoId(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.VideoId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) Validate() error {
	if s.Snapshots != nil {
		if err := s.Snapshots.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) GetSnapshot() []*string {
	return s.Snapshot
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) SetSnapshot(v []*string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots {
	s.Snapshot = v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) Validate() error {
	return dara.Validate(s)
}
