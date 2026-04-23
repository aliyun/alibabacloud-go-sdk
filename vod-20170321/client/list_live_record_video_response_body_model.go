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
	AppName         *string                                                                 `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName      *string                                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PlaylistId      *string                                                                 `json:"PlaylistId,omitempty" xml:"PlaylistId,omitempty"`
	RecordEndTime   *string                                                                 `json:"RecordEndTime,omitempty" xml:"RecordEndTime,omitempty"`
	RecordStartTime *string                                                                 `json:"RecordStartTime,omitempty" xml:"RecordStartTime,omitempty"`
	StreamName      *string                                                                 `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	Video           *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
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
	CateId          *int32                                                                           `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName        *string                                                                          `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL        *string                                                                          `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime    *string                                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration        *float32                                                                         `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModifyTime      *string                                                                          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Size            *int64                                                                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots       *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Status          *string                                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            *string                                                                          `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateGroupId *string                                                                          `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	Title           *string                                                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId         *string                                                                          `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
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
