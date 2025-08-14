// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoListResponseBody
	GetCode() *string
	SetMediaList(v []*GetVideoListResponseBodyMediaList) *GetVideoListResponseBody
	GetMediaList() []*GetVideoListResponseBodyMediaList
	SetRequestId(v string) *GetVideoListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetVideoListResponseBody
	GetSuccess() *string
	SetTotal(v int64) *GetVideoListResponseBody
	GetTotal() *int64
}

type GetVideoListResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the audio and video files.
	MediaList []*GetVideoListResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of audio and video files that meet the conditions.
	//
	// example:
	//
	// 163
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetVideoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoListResponseBody) GetMediaList() []*GetVideoListResponseBodyMediaList {
	return s.MediaList
}

func (s *GetVideoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetVideoListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetVideoListResponseBody) SetCode(v string) *GetVideoListResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoListResponseBody) SetMediaList(v []*GetVideoListResponseBodyMediaList) *GetVideoListResponseBody {
	s.MediaList = v
	return s
}

func (s *GetVideoListResponseBody) SetRequestId(v string) *GetVideoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoListResponseBody) SetSuccess(v string) *GetVideoListResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoListResponseBody) SetTotal(v int64) *GetVideoListResponseBody {
	s.Total = &v
	return s
}

func (s *GetVideoListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoListResponseBodyMediaList struct {
	// The ID of the application. Default value: app-1000000.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 3679
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// http://example.aliyundoc.com/snapshot/****.jpg?auth_key=1498476426-0-0-f00b9455c49a423ce69cf4e27333****
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The time when the audio or video file was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the audio or video file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration. Unit: seconds.
	//
	// example:
	//
	// 135.6
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the audio or video file.
	//
	// example:
	//
	// 1c6ce34007d571ed94667630a6bc****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The time when the audio or video file was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:16:50Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The size of the source file. Unit: bytes.
	//
	// example:
	//
	// 10897890
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The array of video snapshot URLs.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The status of the video.
	//
	// Valid values:
	//
	// 	- PrepareFail: The file is abnormal.
	//
	// 	- UploadFail: The video failed to be uploaded.
	//
	// 	- UploadSucc: The video is uploaded.
	//
	// 	- Transcoding: The video is being transcoded.
	//
	// 	- TranscodeFail: The video failed to be transcoded.
	//
	// 	- ProduceFail: The video failed to be produced.
	//
	// 	- Normal: The video is normal.
	//
	// 	- Uploading: The video is being uploaded.
	//
	// 	- Preparing: The file is being generated.
	//
	// 	- Blocked: The video is blocked.
	//
	// 	- checking: The video is being reviewed.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage address.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the audio or video file.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the audio or video file.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetVideoListResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyMediaList) GetAppId() *string {
	return s.AppId
}

func (s *GetVideoListResponseBodyMediaList) GetCateId() *int64 {
	return s.CateId
}

func (s *GetVideoListResponseBodyMediaList) GetCateName() *string {
	return s.CateName
}

func (s *GetVideoListResponseBodyMediaList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *GetVideoListResponseBodyMediaList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetVideoListResponseBodyMediaList) GetDescription() *string {
	return s.Description
}

func (s *GetVideoListResponseBodyMediaList) GetDuration() *float32 {
	return s.Duration
}

func (s *GetVideoListResponseBodyMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *GetVideoListResponseBodyMediaList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetVideoListResponseBodyMediaList) GetSize() *int64 {
	return s.Size
}

func (s *GetVideoListResponseBodyMediaList) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *GetVideoListResponseBodyMediaList) GetStatus() *string {
	return s.Status
}

func (s *GetVideoListResponseBodyMediaList) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetVideoListResponseBodyMediaList) GetTags() *string {
	return s.Tags
}

func (s *GetVideoListResponseBodyMediaList) GetTitle() *string {
	return s.Title
}

func (s *GetVideoListResponseBodyMediaList) SetAppId(v string) *GetVideoListResponseBodyMediaList {
	s.AppId = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetCateId(v int64) *GetVideoListResponseBodyMediaList {
	s.CateId = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetCateName(v string) *GetVideoListResponseBodyMediaList {
	s.CateName = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetCoverUrl(v string) *GetVideoListResponseBodyMediaList {
	s.CoverUrl = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetCreationTime(v string) *GetVideoListResponseBodyMediaList {
	s.CreationTime = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetDescription(v string) *GetVideoListResponseBodyMediaList {
	s.Description = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetDuration(v float32) *GetVideoListResponseBodyMediaList {
	s.Duration = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetMediaId(v string) *GetVideoListResponseBodyMediaList {
	s.MediaId = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetModificationTime(v string) *GetVideoListResponseBodyMediaList {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetSize(v int64) *GetVideoListResponseBodyMediaList {
	s.Size = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetSnapshots(v []*string) *GetVideoListResponseBodyMediaList {
	s.Snapshots = v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetStatus(v string) *GetVideoListResponseBodyMediaList {
	s.Status = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetStorageLocation(v string) *GetVideoListResponseBodyMediaList {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetTags(v string) *GetVideoListResponseBodyMediaList {
	s.Tags = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) SetTitle(v string) *GetVideoListResponseBodyMediaList {
	s.Title = &v
	return s
}

func (s *GetVideoListResponseBodyMediaList) Validate() error {
	return dara.Validate(s)
}
