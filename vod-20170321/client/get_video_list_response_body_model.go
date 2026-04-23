// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVideoListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *GetVideoListResponseBody
	GetTotal() *int32
	SetVideoList(v *GetVideoListResponseBodyVideoList) *GetVideoListResponseBody
	GetVideoList() *GetVideoListResponseBodyVideoList
}

type GetVideoListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of media files returned.
	//
	// example:
	//
	// 100
	Total     *int32                             `json:"Total,omitempty" xml:"Total,omitempty"`
	VideoList *GetVideoListResponseBodyVideoList `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Struct"`
}

func (s GetVideoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetVideoListResponseBody) GetVideoList() *GetVideoListResponseBodyVideoList {
	return s.VideoList
}

func (s *GetVideoListResponseBody) SetRequestId(v string) *GetVideoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoListResponseBody) SetTotal(v int32) *GetVideoListResponseBody {
	s.Total = &v
	return s
}

func (s *GetVideoListResponseBody) SetVideoList(v *GetVideoListResponseBodyVideoList) *GetVideoListResponseBody {
	s.VideoList = v
	return s
}

func (s *GetVideoListResponseBody) Validate() error {
	if s.VideoList != nil {
		if err := s.VideoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoListResponseBodyVideoList struct {
	Video []*GetVideoListResponseBodyVideoListVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Repeated"`
}

func (s GetVideoListResponseBodyVideoList) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyVideoList) GetVideo() []*GetVideoListResponseBodyVideoListVideo {
	return s.Video
}

func (s *GetVideoListResponseBodyVideoList) SetVideo(v []*GetVideoListResponseBodyVideoListVideo) *GetVideoListResponseBodyVideoList {
	s.Video = v
	return s
}

func (s *GetVideoListResponseBodyVideoList) Validate() error {
	if s.Video != nil {
		for _, item := range s.Video {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoListResponseBodyVideoListVideo struct {
	AppId            *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId           *int64   `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string  `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL         *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime     *string  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration         *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModificationTime *string  `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId       *string                                          `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	RestoreExpiration *string                                          `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	RestoreStatus     *string                                          `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	Size              *int64                                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots         *GetVideoListResponseBodyVideoListVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Status            *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageClass      *string                                          `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	StorageLocation   *string                                          `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags              *string                                          `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title             *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData          *string                                          `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId           *string                                          `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoListResponseBodyVideoListVideo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponseBodyVideoListVideo) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyVideoListVideo) GetAppId() *string {
	return s.AppId
}

func (s *GetVideoListResponseBodyVideoListVideo) GetCateId() *int64 {
	return s.CateId
}

func (s *GetVideoListResponseBodyVideoListVideo) GetCateName() *string {
	return s.CateName
}

func (s *GetVideoListResponseBodyVideoListVideo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetVideoListResponseBodyVideoListVideo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetVideoListResponseBodyVideoListVideo) GetDescription() *string {
	return s.Description
}

func (s *GetVideoListResponseBodyVideoListVideo) GetDuration() *float32 {
	return s.Duration
}

func (s *GetVideoListResponseBodyVideoListVideo) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetVideoListResponseBodyVideoListVideo) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetVideoListResponseBodyVideoListVideo) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *GetVideoListResponseBodyVideoListVideo) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *GetVideoListResponseBodyVideoListVideo) GetSize() *int64 {
	return s.Size
}

func (s *GetVideoListResponseBodyVideoListVideo) GetSnapshots() *GetVideoListResponseBodyVideoListVideoSnapshots {
	return s.Snapshots
}

func (s *GetVideoListResponseBodyVideoListVideo) GetStatus() *string {
	return s.Status
}

func (s *GetVideoListResponseBodyVideoListVideo) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetVideoListResponseBodyVideoListVideo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetVideoListResponseBodyVideoListVideo) GetTags() *string {
	return s.Tags
}

func (s *GetVideoListResponseBodyVideoListVideo) GetTitle() *string {
	return s.Title
}

func (s *GetVideoListResponseBodyVideoListVideo) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoListResponseBodyVideoListVideo) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoListResponseBodyVideoListVideo) SetAppId(v string) *GetVideoListResponseBodyVideoListVideo {
	s.AppId = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCateId(v int64) *GetVideoListResponseBodyVideoListVideo {
	s.CateId = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCateName(v string) *GetVideoListResponseBodyVideoListVideo {
	s.CateName = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCoverURL(v string) *GetVideoListResponseBodyVideoListVideo {
	s.CoverURL = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCreationTime(v string) *GetVideoListResponseBodyVideoListVideo {
	s.CreationTime = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetDescription(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Description = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetDuration(v float32) *GetVideoListResponseBodyVideoListVideo {
	s.Duration = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetModificationTime(v string) *GetVideoListResponseBodyVideoListVideo {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetReferenceId(v string) *GetVideoListResponseBodyVideoListVideo {
	s.ReferenceId = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetRestoreExpiration(v string) *GetVideoListResponseBodyVideoListVideo {
	s.RestoreExpiration = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetRestoreStatus(v string) *GetVideoListResponseBodyVideoListVideo {
	s.RestoreStatus = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetSize(v int64) *GetVideoListResponseBodyVideoListVideo {
	s.Size = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetSnapshots(v *GetVideoListResponseBodyVideoListVideoSnapshots) *GetVideoListResponseBodyVideoListVideo {
	s.Snapshots = v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetStatus(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Status = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetStorageClass(v string) *GetVideoListResponseBodyVideoListVideo {
	s.StorageClass = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetStorageLocation(v string) *GetVideoListResponseBodyVideoListVideo {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetTags(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Tags = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetTitle(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Title = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetUserData(v string) *GetVideoListResponseBodyVideoListVideo {
	s.UserData = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetVideoId(v string) *GetVideoListResponseBodyVideoListVideo {
	s.VideoId = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) Validate() error {
	if s.Snapshots != nil {
		if err := s.Snapshots.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoListResponseBodyVideoListVideoSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s GetVideoListResponseBodyVideoListVideoSnapshots) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponseBodyVideoListVideoSnapshots) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyVideoListVideoSnapshots) GetSnapshot() []*string {
	return s.Snapshot
}

func (s *GetVideoListResponseBodyVideoListVideoSnapshots) SetSnapshot(v []*string) *GetVideoListResponseBodyVideoListVideoSnapshots {
	s.Snapshot = v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideoSnapshots) Validate() error {
	return dara.Validate(s)
}
